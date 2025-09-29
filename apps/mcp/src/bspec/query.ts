/**
 * BSpec Query Engine
 *
 * Advanced search and filtering capabilities for BSpec documents
 * Supports complex queries across all document types and relationships
 */

import type { ParsedBSpecDocument } from './parser';
import type { BusinessDomain, DocumentType, DocumentStatus, Priority } from '../types.js';

export interface QueryOptions {
  // Document type filters
  types?: DocumentType[];
  domains?: BusinessDomain[];
  statuses?: DocumentStatus[];
  priorities?: Priority[];

  // Text search
  textSearch?: string;
  searchFields?: ('title' | 'content' | 'id' | 'tags')[];

  // Relationship filters
  dependsOn?: string[];
  enables?: string[];
  relatedTo?: string[];
  conflictsWith?: string[];

  // Temporal filters
  createdAfter?: string;
  createdBefore?: string;
  updatedAfter?: string;
  updatedBefore?: string;

  // Ownership filters
  owner?: string;
  stakeholders?: string[];
  reviewers?: string[];

  // Business context filters
  implementationStatus?: string[];
  hasSuccessCriteria?: boolean;
  hasRisks?: boolean;
  hasMetrics?: boolean;

  // Result options
  sortBy?: SortOption;
  sortOrder?: 'asc' | 'desc';
  limit?: number;
  offset?: number;
  includeContent?: boolean;
}

export type SortOption =
  | 'created'
  | 'updated'
  | 'title'
  | 'type'
  | 'priority'
  | 'status'
  | 'relationships'
  | 'relevance';

export interface QueryResult {
  documents: ParsedBSpecDocument[];
  totalCount: number;
  executionTimeMs: number;
  appliedFilters: string[];
  suggestions?: string[];
}

export interface SearchSuggestion {
  type: 'document' | 'relationship' | 'domain' | 'tag';
  value: string;
  count: number;
  relevance: number;
}

/**
 * Main query engine for BSpec documents
 */
export class BSpecQueryEngine {
  private documents: ParsedBSpecDocument[];
  private indexCache: Map<string, any> = new Map();

  constructor(documents: ParsedBSpecDocument[]) {
    this.documents = documents;
    this.buildSearchIndexes();
  }

  /**
   * Execute a query against the document collection
   */
  query(options: QueryOptions = {}): QueryResult {
    const startTime = Date.now();
    let filteredDocs = [...this.documents];
    const appliedFilters: string[] = [];

    // Apply filters in order of selectivity (most selective first)
    filteredDocs = this.applyTypeFilters(filteredDocs, options, appliedFilters);
    filteredDocs = this.applyStatusFilters(filteredDocs, options, appliedFilters);
    filteredDocs = this.applyRelationshipFilters(filteredDocs, options, appliedFilters);
    filteredDocs = this.applyTemporalFilters(filteredDocs, options, appliedFilters);
    filteredDocs = this.applyOwnershipFilters(filteredDocs, options, appliedFilters);
    filteredDocs = this.applyBusinessContextFilters(filteredDocs, options, appliedFilters);
    filteredDocs = this.applyTextSearch(filteredDocs, options, appliedFilters);

    // Sort results
    filteredDocs = this.sortDocuments(filteredDocs, options);

    // Apply pagination
    const totalCount = filteredDocs.length;
    const offset = options.offset || 0;
    const limit = options.limit || totalCount;
    filteredDocs = filteredDocs.slice(offset, offset + limit);

    // Optionally exclude content for lighter responses
    if (!options.includeContent) {
      filteredDocs = filteredDocs.map(doc => ({ ...doc, content: '' }));
    }

    const executionTimeMs = Date.now() - startTime;

    return {
      documents: filteredDocs,
      totalCount,
      executionTimeMs,
      appliedFilters,
      suggestions: this.generateSuggestions(options, filteredDocs)
    };
  }

  /**
   * Find documents similar to a given document
   */
  findSimilarDocuments(documentId: string, limit: number = 5): ParsedBSpecDocument[] {
    const sourceDoc = this.documents.find(doc => doc.id === documentId);
    if (!sourceDoc) return [];

    const similarities = this.documents
      .filter(doc => doc.id !== documentId)
      .map(doc => ({
        document: doc,
        similarity: this.calculateSimilarity(sourceDoc, doc)
      }))
      .sort((a, b) => b.similarity - a.similarity)
      .slice(0, limit);

    return similarities.map(s => s.document);
  }

  /**
   * Get document recommendations based on current collection
   */
  getDocumentRecommendations(): {
    missingTypes: DocumentType[];
    incompleteChains: string[];
    opportunities: string[];
  } {
    const presentTypes = new Set(this.documents.map(doc => doc.type));
    const allTypes = this.getAllDocumentTypes();
    const missingTypes = allTypes.filter(type => !presentTypes.has(type));

    const incompleteChains = this.findIncompleteRelationshipChains();
    const opportunities = this.identifyOpportunities();

    return {
      missingTypes,
      incompleteChains,
      opportunities
    };
  }

  /**
   * Get aggregated statistics about the document collection
   */
  getCollectionStats(): {
    byType: Record<DocumentType, number>;
    byDomain: Record<BusinessDomain, number>;
    byStatus: Record<DocumentStatus, number>;
    relationshipStats: {
      totalDependencies: number;
      totalEnablements: number;
      totalConflicts: number;
      avgRelationshipsPerDoc: number;
    };
    temporalStats: {
      oldestDocument: string;
      newestDocument: string;
      avgAge: number;
    };
  } {
    const byType = this.groupBy(this.documents, doc => doc.type);
    const byDomain = this.groupBy(this.documents, doc => doc.domain || 'strategic');
    const byStatus = this.groupBy(this.documents, doc => doc.status);

    const relationshipStats = this.calculateRelationshipStats();
    const temporalStats = this.calculateTemporalStats();

    return {
      byType,
      byDomain,
      byStatus,
      relationshipStats,
      temporalStats
    };
  }

  /**
   * Private helper methods
   */

  private buildSearchIndexes(): void {
    // Build text search index
    const textIndex = new Map<string, Set<string>>();

    this.documents.forEach(doc => {
      const searchableText = [
        doc.title,
        doc.content,
        doc.id,
        ...(doc.tags || [])
      ].join(' ').toLowerCase();

      const words = searchableText.split(/\s+/).filter(word => word.length > 2);

      words.forEach(word => {
        if (!textIndex.has(word)) {
          textIndex.set(word, new Set());
        }
        textIndex.get(word)!.add(doc.id);
      });
    });

    this.indexCache.set('textIndex', textIndex);

    // Build relationship index
    const relationshipIndex = new Map<string, {
      dependents: Set<string>;
      enables: Set<string>;
      relatedTo: Set<string>;
    }>();

    this.documents.forEach(doc => {
      relationshipIndex.set(doc.id, {
        dependents: new Set(),
        enables: new Set(),
        relatedTo: new Set()
      });
    });

    this.documents.forEach(doc => {
      doc.depends_on?.forEach(depId => {
        relationshipIndex.get(depId)?.dependents.add(doc.id);
      });

      doc.enables?.forEach(enabledId => {
        relationshipIndex.get(doc.id)?.enables.add(enabledId);
      });

      doc.related?.forEach(relatedId => {
        relationshipIndex.get(doc.id)?.relatedTo.add(relatedId);
        relationshipIndex.get(relatedId)?.relatedTo.add(doc.id);
      });
    });

    this.indexCache.set('relationshipIndex', relationshipIndex);
  }

  private applyTypeFilters(
    docs: ParsedBSpecDocument[],
    options: QueryOptions,
    appliedFilters: string[]
  ): ParsedBSpecDocument[] {
    if (options.types && options.types.length > 0) {
      appliedFilters.push(`types: ${options.types.join(', ')}`);
      docs = docs.filter(doc => options.types!.includes(doc.type));
    }

    if (options.domains && options.domains.length > 0) {
      appliedFilters.push(`domains: ${options.domains.join(', ')}`);
      docs = docs.filter(doc =>
        options.domains!.includes(doc.domain || this.getDomainForType(doc.type))
      );
    }

    return docs;
  }

  private applyStatusFilters(
    docs: ParsedBSpecDocument[],
    options: QueryOptions,
    appliedFilters: string[]
  ): ParsedBSpecDocument[] {
    if (options.statuses && options.statuses.length > 0) {
      appliedFilters.push(`statuses: ${options.statuses.join(', ')}`);
      docs = docs.filter(doc => options.statuses!.includes(doc.status));
    }

    if (options.priorities && options.priorities.length > 0) {
      appliedFilters.push(`priorities: ${options.priorities.join(', ')}`);
      docs = docs.filter(doc => doc.priority && options.priorities!.includes(doc.priority));
    }

    return docs;
  }

  private applyRelationshipFilters(
    docs: ParsedBSpecDocument[],
    options: QueryOptions,
    appliedFilters: string[]
  ): ParsedBSpecDocument[] {
    if (options.dependsOn && options.dependsOn.length > 0) {
      appliedFilters.push(`depends on: ${options.dependsOn.join(', ')}`);
      docs = docs.filter(doc =>
        doc.depends_on?.some(depId => options.dependsOn!.includes(depId))
      );
    }

    if (options.enables && options.enables.length > 0) {
      appliedFilters.push(`enables: ${options.enables.join(', ')}`);
      docs = docs.filter(doc =>
        doc.enables?.some(enabledId => options.enables!.includes(enabledId))
      );
    }

    if (options.relatedTo && options.relatedTo.length > 0) {
      appliedFilters.push(`related to: ${options.relatedTo.join(', ')}`);
      docs = docs.filter(doc =>
        doc.related?.some(relatedId => options.relatedTo!.includes(relatedId))
      );
    }

    return docs;
  }

  private applyTemporalFilters(
    docs: ParsedBSpecDocument[],
    options: QueryOptions,
    appliedFilters: string[]
  ): ParsedBSpecDocument[] {
    if (options.createdAfter) {
      appliedFilters.push(`created after: ${options.createdAfter}`);
      docs = docs.filter(doc => doc.created >= options.createdAfter!);
    }

    if (options.createdBefore) {
      appliedFilters.push(`created before: ${options.createdBefore}`);
      docs = docs.filter(doc => doc.created <= options.createdBefore!);
    }

    if (options.updatedAfter) {
      appliedFilters.push(`updated after: ${options.updatedAfter}`);
      docs = docs.filter(doc => doc.updated >= options.updatedAfter!);
    }

    if (options.updatedBefore) {
      appliedFilters.push(`updated before: ${options.updatedBefore}`);
      docs = docs.filter(doc => doc.updated <= options.updatedBefore!);
    }

    return docs;
  }

  private applyOwnershipFilters(
    docs: ParsedBSpecDocument[],
    options: QueryOptions,
    appliedFilters: string[]
  ): ParsedBSpecDocument[] {
    if (options.owner) {
      appliedFilters.push(`owner: ${options.owner}`);
      docs = docs.filter(doc => doc.owner === options.owner);
    }

    if (options.stakeholders && options.stakeholders.length > 0) {
      appliedFilters.push(`stakeholders: ${options.stakeholders.join(', ')}`);
      docs = docs.filter(doc =>
        doc.stakeholders?.some(stakeholder => options.stakeholders!.includes(stakeholder))
      );
    }

    return docs;
  }

  private applyBusinessContextFilters(
    docs: ParsedBSpecDocument[],
    options: QueryOptions,
    appliedFilters: string[]
  ): ParsedBSpecDocument[] {
    if (options.hasSuccessCriteria !== undefined) {
      appliedFilters.push(`has success criteria: ${options.hasSuccessCriteria}`);
      docs = docs.filter(doc =>
        !!doc.success_criteria?.length === options.hasSuccessCriteria
      );
    }

    if (options.hasRisks !== undefined) {
      appliedFilters.push(`has risks: ${options.hasRisks}`);
      docs = docs.filter(doc => !!doc.risks?.length === options.hasRisks);
    }

    if (options.hasMetrics !== undefined) {
      appliedFilters.push(`has metrics: ${options.hasMetrics}`);
      docs = docs.filter(doc => !!doc.metrics?.length === options.hasMetrics);
    }

    return docs;
  }

  private applyTextSearch(
    docs: ParsedBSpecDocument[],
    options: QueryOptions,
    appliedFilters: string[]
  ): ParsedBSpecDocument[] {
    if (!options.textSearch) return docs;

    appliedFilters.push(`text search: "${options.textSearch}"`);
    const searchTerm = options.textSearch.toLowerCase();
    const searchFields = options.searchFields || ['title', 'content', 'id', 'tags'];

    return docs.filter(doc => {
      return searchFields.some(field => {
        switch (field) {
          case 'title':
            return doc.title.toLowerCase().includes(searchTerm);
          case 'content':
            return doc.content.toLowerCase().includes(searchTerm);
          case 'id':
            return doc.id.toLowerCase().includes(searchTerm);
          case 'tags':
            return doc.tags?.some(tag => tag.toLowerCase().includes(searchTerm));
          default:
            return false;
        }
      });
    });
  }

  private sortDocuments(docs: ParsedBSpecDocument[], options: QueryOptions): ParsedBSpecDocument[] {
    if (!options.sortBy) return docs;

    const sortOrder = options.sortOrder || 'asc';
    const multiplier = sortOrder === 'asc' ? 1 : -1;

    return docs.sort((a, b) => {
      let comparison = 0;

      switch (options.sortBy) {
        case 'created':
          comparison = a.created.localeCompare(b.created);
          break;
        case 'updated':
          comparison = a.updated.localeCompare(b.updated);
          break;
        case 'title':
          comparison = a.title.localeCompare(b.title);
          break;
        case 'type':
          comparison = a.type.localeCompare(b.type);
          break;
        case 'status':
          comparison = a.status.localeCompare(b.status);
          break;
        case 'relationships':
          const aRelCount = (a.depends_on?.length || 0) + (a.enables?.length || 0);
          const bRelCount = (b.depends_on?.length || 0) + (b.enables?.length || 0);
          comparison = aRelCount - bRelCount;
          break;
        default:
          comparison = 0;
      }

      return comparison * multiplier;
    });
  }

  private calculateSimilarity(doc1: ParsedBSpecDocument, doc2: ParsedBSpecDocument): number {
    let similarity = 0;

    // Domain similarity
    if (doc1.domain === doc2.domain) similarity += 0.3;

    // Type similarity (same domain family)
    if (this.getDocumentFamily(doc1.type) === this.getDocumentFamily(doc2.type)) {
      similarity += 0.2;
    }

    // Relationship similarity
    const commonDeps = this.getCommonElements(doc1.depends_on || [], doc2.depends_on || []);
    const commonEnables = this.getCommonElements(doc1.enables || [], doc2.enables || []);
    similarity += (commonDeps.length + commonEnables.length) * 0.1;

    // Content similarity (simplified)
    const content1Words = new Set(doc1.content.toLowerCase().split(/\s+/));
    const content2Words = new Set(doc2.content.toLowerCase().split(/\s+/));
    const commonWords = this.getCommonElements(Array.from(content1Words), Array.from(content2Words));
    similarity += Math.min(commonWords.length / 100, 0.3);

    return Math.min(similarity, 1.0);
  }

  private generateSuggestions(options: QueryOptions, results: ParsedBSpecDocument[]): string[] {
    const suggestions: string[] = [];

    if (results.length === 0) {
      suggestions.push('Try removing some filters to get more results');
      suggestions.push('Check document type spellings (MSN, VSN, STR, etc.)');
    }

    if (results.length > 100) {
      suggestions.push('Consider adding more specific filters to narrow results');
    }

    return suggestions;
  }

  private getDomainForType(type: DocumentType): BusinessDomain {
    // Implementation similar to analyzer.ts
    const domainMap: Record<string, BusinessDomain> = {
      'MSN': 'strategic', 'VSN': 'strategic', 'VAL': 'strategic', 'STR': 'strategic',
      'MKT': 'market', 'SEG': 'market', 'CMP': 'market', 'POS': 'market',
      // ... other mappings
    };
    return domainMap[type] || 'strategic';
  }

  private getDocumentFamily(type: DocumentType): string {
    const families: Record<string, string[]> = {
      'strategic': ['MSN', 'VSN', 'VAL', 'STR', 'OBJ', 'MOT', 'PUR', 'THY'],
      'market': ['MKT', 'SEG', 'CMP', 'POS', 'TRN', 'ECO', 'OPP', 'THR', 'REG', 'MAC'],
      // ... other families
    };

    for (const [family, types] of Object.entries(families)) {
      if (types.includes(type)) return family;
    }
    return 'unknown';
  }

  private getCommonElements<T>(arr1: T[], arr2: T[]): T[] {
    return arr1.filter(item => arr2.includes(item));
  }

  private getAllDocumentTypes(): DocumentType[] {
    return [
      'MSN', 'VSN', 'VAL', 'STR', 'OBJ', 'MOT', 'PUR', 'THY',
      'MKT', 'SEG', 'CMP', 'POS', 'TRN', 'ECO', 'OPP', 'THR', 'REG', 'MAC',
      // ... all 82 types
    ] as DocumentType[];
  }

  private findIncompleteRelationshipChains(): string[] {
    // Analyze relationship chains and find missing links
    return [];
  }

  private identifyOpportunities(): string[] {
    // Identify opportunities for new documents or relationships
    return [];
  }

  private groupBy<T, K extends string | number>(
    items: T[],
    keyFn: (item: T) => K
  ): Record<K, number> {
    const result = {} as Record<K, number>;
    items.forEach(item => {
      const key = keyFn(item);
      result[key] = (result[key] || 0) + 1;
    });
    return result;
  }

  private calculateRelationshipStats() {
    const totalDependencies = this.documents.reduce((sum, doc) => sum + (doc.depends_on?.length || 0), 0);
    const totalEnablements = this.documents.reduce((sum, doc) => sum + (doc.enables?.length || 0), 0);
    const totalConflicts = this.documents.reduce((sum, doc) => sum + (doc.conflicts_with?.length || 0), 0);
    const totalRelationships = totalDependencies + totalEnablements + totalConflicts;

    return {
      totalDependencies,
      totalEnablements,
      totalConflicts,
      avgRelationshipsPerDoc: totalRelationships / this.documents.length
    };
  }

  private calculateTemporalStats() {
    const dates = this.documents.map(doc => doc.created).sort();
    const ages = this.documents.map(doc =>
      Math.floor((Date.now() - new Date(doc.created).getTime()) / (1000 * 60 * 60 * 24))
    );

    return {
      oldestDocument: dates[0],
      newestDocument: dates[dates.length - 1],
      avgAge: ages.reduce((sum, age) => sum + age, 0) / ages.length
    };
  }
}