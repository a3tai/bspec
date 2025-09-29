/**
 * BSpec TypeScript SDK - Main BSpec Class
 * Generated from BSpec v1.0.0 JSON SDK
 */

import type {
  BSpecData,
  BSpecDomain,
  BSpecDocumentType,
  BSpecFile,
  BSpecStatistics,
  BusinessDomainName,
  DocumentTypeCode,
  ConformanceLevelName
} from './types';

/**
 * Main BSpec class for working with BSpec data
 */
export class BSpec {
  private data: BSpecData;

  constructor(data: BSpecData) {
    this.data = data;
  }

  /**
   * Create BSpec instance from JSON data
   */
  static fromJSON(jsonData: BSpecData): BSpec {
    return new BSpec(jsonData);
  }

  /**
   * Get BSpec version
   */
  get version(): string {
    return this.data.metadata.bspec_version;
  }

  /**
   * Get generation metadata
   */
  get metadata() {
    return this.data.metadata;
  }

  /**
   * Get specification statistics
   */
  get statistics(): BSpecStatistics {
    return this.data.statistics;
  }

  /**
   * Get all business domains
   */
  get domains(): BSpecDomain[] {
    return this.data.domains;
  }

  /**
   * Get all document types
   */
  get documentTypes(): BSpecDocumentType[] {
    return this.data.document_types;
  }

  /**
   * Get all files in the specification
   */
  get files(): Record<string, BSpecFile> {
    return this.data.files;
  }

  /**
   * Get conformance levels
   */
  get conformanceLevels() {
    return this.data.conformance_levels;
  }

  /**
   * Get YAML schema definition
   */
  get yamlSchema() {
    return this.data.yaml_schema;
  }

  /**
   * Get directory structure
   */
  get directoryStructure(): string[][] {
    return this.data.directory_structure;
  }

  /**
   * Get document index
   */
  get documentIndex() {
    return this.data.document_index;
  }

  /**
   * Get a specific domain by name
   */
  getDomain(name: BusinessDomainName): BSpecDomain | undefined {
    return this.data.domains.find(domain => domain.name === name);
  }

  /**
   * Get a specific document type by code
   */
  getDocumentType(code: DocumentTypeCode): BSpecDocumentType | undefined {
    return this.data.document_types.find(type => type.code === code);
  }

  /**
   * Get document types for a specific domain
   */
  getDocumentTypesForDomain(domainName: BusinessDomainName): BSpecDocumentType[] {
    const domain = this.getDomain(domainName);
    if (!domain) return [];

    return this.data.document_types.filter(type =>
      domain.document_types.includes(type.code)
    );
  }

  /**
   * Get a specific file by path
   */
  getFile(path: string): BSpecFile | undefined {
    return this.data.files[path];
  }

  /**
   * Get all files of a specific type
   */
  getFilesByType(type: BSpecFile['type']): BSpecFile[] {
    return Object.values(this.data.files).filter(file => file.type === type);
  }

  /**
   * Get markdown files with frontmatter
   */
  getDocumentFiles(): BSpecFile[] {
    return Object.values(this.data.files).filter(file =>
      file.type === 'markdown' && file.has_frontmatter
    );
  }

  /**
   * Search for document types by name or purpose
   */
  searchDocumentTypes(query: string): BSpecDocumentType[] {
    const lowercaseQuery = query.toLowerCase();
    return this.data.document_types.filter(type =>
      type.name.toLowerCase().includes(lowercaseQuery) ||
      type.purpose.toLowerCase().includes(lowercaseQuery) ||
      type.code.toLowerCase().includes(lowercaseQuery)
    );
  }

  /**
   * Get domains containing a specific document type
   */
  getDomainsForDocumentType(code: DocumentTypeCode): BSpecDomain[] {
    return this.data.domains.filter(domain =>
      domain.document_types.includes(code)
    );
  }

  /**
   * Get conformance level by name
   */
  getConformanceLevel(name: ConformanceLevelName) {
    return this.data.conformance_levels.find(level => level.name === name);
  }

  /**
   * Validate that all required YAML fields are present
   */
  getRequiredFields(): string[] {
    return this.data.yaml_schema.required_fields;
  }

  /**
   * Get all optional YAML fields
   */
  getOptionalFields(): string[] {
    return this.data.yaml_schema.optional_fields;
  }

  /**
   * Get summary information about the specification
   */
  getSummary() {
    return {
      version: this.version,
      statistics: this.statistics,
      domains: this.domains.map(domain => ({
        name: domain.name,
        description: domain.description,
        documentCount: domain.document_count
      })),
      conformanceLevels: this.conformanceLevels.map(level => level.name),
      generatedAt: this.metadata.generated_at,
      generator: this.metadata.generator
    };
  }

  /**
   * Export to JSON
   */
  toJSON(): BSpecData {
    return this.data;
  }
}

/**
 * Load BSpec from JSON file content
 */
export function loadBSpec(jsonContent: string): BSpec {
  const data = JSON.parse(jsonContent) as BSpecData;
  return BSpec.fromJSON(data);
}

/**
 * Default export for convenience
 */
export default BSpec;