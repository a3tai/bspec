/**
 * BSpec Relationship Analyzer
 *
 * Analyzes document relationships, dependencies, and business knowledge graph structure
 * Provides insights into business coherence and completeness
 */

import type { ParsedBSpecDocument } from './parser';
import type { BusinessDomain, DocumentType, ConformanceLevel } from '../types.js';

export interface RelationshipGraph {
  nodes: DocumentNode[];
  edges: RelationshipEdge[];
  clusters: DomainCluster[];
  metrics: GraphMetrics;
}

export interface DocumentNode {
  id: string;
  type: DocumentType;
  title: string;
  domain: BusinessDomain;
  status: string;
  priority?: string;
  centrality: number;
  inDegree: number;
  outDegree: number;
}

export interface RelationshipEdge {
  source: string;
  target: string;
  type: RelationshipType;
  strength: number;
}

export type RelationshipType = 'depends_on' | 'enables' | 'conflicts_with' | 'related' | 'supersedes';

export interface DomainCluster {
  domain: BusinessDomain;
  documents: string[];
  completeness: number;
  connections: number;
  criticality: 'high' | 'medium' | 'low';
}

export interface GraphMetrics {
  totalDocuments: number;
  totalRelationships: number;
  avgConnectionsPerDocument: number;
  clusteringCoefficient: number;
  longestDependencyChain: number;
  conformanceLevel: ConformanceLevel;
  completenessScore: number;
}

export interface ImpactAnalysis {
  documentId: string;
  directlyAffected: string[];
  transitivelyAffected: string[];
  totalImpact: number;
  riskLevel: 'low' | 'medium' | 'high' | 'critical';
}

export interface DependencyChain {
  chain: string[];
  length: number;
  isCircular: boolean;
  riskFactors: string[];
}

/**
 * Build a comprehensive relationship graph from document collection
 */
export function buildRelationshipGraph(documents: ParsedBSpecDocument[]): RelationshipGraph {
  const nodes = documents.map(doc => createDocumentNode(doc, documents));
  const edges = extractRelationshipEdges(documents);
  const clusters = analyzeDomainClusters(documents);
  const metrics = calculateGraphMetrics(nodes, edges, documents);

  return {
    nodes,
    edges,
    clusters,
    metrics
  };
}

/**
 * Create a document node with centrality and connection metrics
 */
function createDocumentNode(document: ParsedBSpecDocument, allDocuments: ParsedBSpecDocument[]): DocumentNode {
  const incomingConnections = allDocuments.filter(doc =>
    doc.depends_on?.includes(document.id) ||
    doc.enables?.includes(document.id) ||
    doc.related?.includes(document.id)
  ).length;

  const outgoingConnections = [
    ...(document.depends_on || []),
    ...(document.enables || []),
    ...(document.related || [])
  ].length;

  return {
    id: document.id,
    type: document.type,
    title: document.title,
    domain: document.domain || getDomainForType(document.type),
    status: document.status,
    priority: document.priority,
    centrality: calculateCentrality(document, allDocuments),
    inDegree: incomingConnections,
    outDegree: outgoingConnections
  };
}

/**
 * Extract all relationship edges between documents
 */
function extractRelationshipEdges(documents: ParsedBSpecDocument[]): RelationshipEdge[] {
  const edges: RelationshipEdge[] = [];

  documents.forEach(doc => {
    // Dependencies
    doc.depends_on?.forEach(targetId => {
      edges.push({
        source: doc.id,
        target: targetId,
        type: 'depends_on',
        strength: 1.0
      });
    });

    // Enablements
    doc.enables?.forEach(targetId => {
      edges.push({
        source: doc.id,
        target: targetId,
        type: 'enables',
        strength: 0.8
      });
    });

    // Conflicts
    doc.conflicts_with?.forEach(targetId => {
      edges.push({
        source: doc.id,
        target: targetId,
        type: 'conflicts_with',
        strength: 0.9
      });
    });

    // Related documents
    doc.related?.forEach(targetId => {
      edges.push({
        source: doc.id,
        target: targetId,
        type: 'related',
        strength: 0.5
      });
    });

    // Supersedes
    if (doc.supersedes) {
      edges.push({
        source: doc.id,
        target: doc.supersedes,
        type: 'supersedes',
        strength: 1.0
      });
    }
  });

  return edges;
}

/**
 * Analyze domain clusters for completeness and connectivity
 */
function analyzeDomainClusters(documents: ParsedBSpecDocument[]): DomainCluster[] {
  const domainGroups = groupDocumentsByDomain(documents);

  return Object.entries(domainGroups).map(([domain, docs]) => {
    const connections = countCrossClusterConnections(docs, documents);
    const completeness = calculateDomainCompleteness(domain as BusinessDomain, docs);

    return {
      domain: domain as BusinessDomain,
      documents: docs.map(d => d.id),
      completeness,
      connections,
      criticality: assessDomainCriticality(domain as BusinessDomain, completeness, connections)
    };
  });
}

/**
 * Calculate comprehensive graph metrics
 */
function calculateGraphMetrics(
  nodes: DocumentNode[],
  edges: RelationshipEdge[],
  documents: ParsedBSpecDocument[]
): GraphMetrics {
  const totalDocuments = nodes.length;
  const totalRelationships = edges.length;
  const avgConnectionsPerDocument = totalRelationships / totalDocuments;

  return {
    totalDocuments,
    totalRelationships,
    avgConnectionsPerDocument,
    clusteringCoefficient: calculateClusteringCoefficient(nodes, edges),
    longestDependencyChain: findLongestDependencyChain(documents),
    conformanceLevel: assessConformanceLevel(totalDocuments),
    completenessScore: calculateCompletenessScore(documents)
  };
}

/**
 * Analyze impact of changing a specific document
 */
export function analyzeDocumentImpact(
  documentId: string,
  documents: ParsedBSpecDocument[]
): ImpactAnalysis {
  const directlyAffected = findDirectlyAffectedDocuments(documentId, documents);
  const transitivelyAffected = findTransitivelyAffectedDocuments(documentId, documents);

  const totalImpact = directlyAffected.length + transitivelyAffected.length;
  const riskLevel = assessRiskLevel(totalImpact, documents.length);

  return {
    documentId,
    directlyAffected,
    transitivelyAffected,
    totalImpact,
    riskLevel
  };
}

/**
 * Find all dependency chains in the document collection
 */
export function findDependencyChains(documents: ParsedBSpecDocument[]): DependencyChain[] {
  const chains: DependencyChain[] = [];
  const visited = new Set<string>();

  documents.forEach(doc => {
    if (!visited.has(doc.id)) {
      const chain = traceDependencyChain(doc.id, documents, new Set());
      if (chain.length > 1) {
        chains.push({
          chain,
          length: chain.length,
          isCircular: chain[0] === chain[chain.length - 1],
          riskFactors: assessChainRisks(chain, documents)
        });
      }
      chain.forEach(id => visited.add(id));
    }
  });

  return chains.sort((a, b) => b.length - a.length);
}

/**
 * Get critical path analysis for business implementation
 */
export function getCriticalPath(documents: ParsedBSpecDocument[]): {
  path: string[];
  estimatedDuration: number;
  blockers: string[];
  recommendations: string[];
} {
  const strategicDocs = documents.filter(doc =>
    ['MSN', 'VSN', 'VAL', 'STR'].includes(doc.type)
  );

  const operationalDocs = documents.filter(doc =>
    doc.domain === 'operations' && doc.status === 'Accepted'
  );

  const path = [...strategicDocs.map(d => d.id), ...operationalDocs.map(d => d.id)];

  return {
    path,
    estimatedDuration: estimateImplementationDuration(path, documents),
    blockers: findImplementationBlockers(documents),
    recommendations: generateImplementationRecommendations(documents)
  };
}

/**
 * Utility Functions
 */

function calculateCentrality(document: ParsedBSpecDocument, allDocuments: ParsedBSpecDocument[]): number {
  const connections = allDocuments.filter(doc =>
    doc.depends_on?.includes(document.id) ||
    doc.enables?.includes(document.id) ||
    document.depends_on?.includes(doc.id) ||
    document.enables?.includes(doc.id)
  ).length;

  return connections / Math.max(allDocuments.length - 1, 1);
}

function getDomainForType(type: DocumentType): BusinessDomain {
  const domainMap: Record<string, BusinessDomain> = {
    'MSN': 'strategic', 'VSN': 'strategic', 'VAL': 'strategic', 'STR': 'strategic',
    'MKT': 'market', 'SEG': 'market', 'CMP': 'market', 'POS': 'market',
    'PER': 'customer', 'JTB': 'customer', 'CJM': 'customer', 'USE': 'customer',
    'PRD': 'product', 'SVC': 'product', 'FEA': 'product', 'REQ': 'product',
    'BMC': 'model', 'REV': 'model', 'PRC': 'model',
    'ORG': 'operations', 'WFL': 'operations', 'ROL': 'operations',
    'ARC': 'technology', 'SYS': 'technology', 'DAT': 'technology',
    'FIN': 'financial', 'BUD': 'financial', 'FOR': 'financial',
    'RSK': 'risk', 'MIT': 'risk', 'GVN': 'risk',
    'GTM': 'growth', 'GRW': 'growth', 'SCL': 'growth'
  };

  return domainMap[type] || 'strategic';
}

function groupDocumentsByDomain(documents: ParsedBSpecDocument[]): Record<string, ParsedBSpecDocument[]> {
  return documents.reduce((groups, doc) => {
    const domain = doc.domain || getDomainForType(doc.type);
    groups[domain] = groups[domain] || [];
    groups[domain].push(doc);
    return groups;
  }, {} as Record<string, ParsedBSpecDocument[]>);
}

function countCrossClusterConnections(clusterDocs: ParsedBSpecDocument[], allDocs: ParsedBSpecDocument[]): number {
  const clusterIds = new Set(clusterDocs.map(d => d.id));

  return clusterDocs.reduce((count, doc) => {
    const externalConnections = [
      ...(doc.depends_on || []),
      ...(doc.enables || []),
      ...(doc.related || [])
    ].filter(id => !clusterIds.has(id));

    return count + externalConnections.length;
  }, 0);
}

function calculateDomainCompleteness(domain: BusinessDomain, documents: ParsedBSpecDocument[]): number {
  const expectedTypesForDomain = getExpectedTypesForDomain(domain);
  const presentTypes = new Set(documents.map(d => d.type));
  const coverage = expectedTypesForDomain.filter(type => presentTypes.has(type)).length;

  return coverage / expectedTypesForDomain.length;
}

function getExpectedTypesForDomain(domain: BusinessDomain): DocumentType[] {
  const domainTypes: Record<BusinessDomain, DocumentType[]> = {
    'strategic': ['MSN', 'VSN', 'VAL', 'STR', 'OBJ', 'MOT', 'PUR', 'THY'],
    'market': ['MKT', 'SEG', 'CMP', 'POS', 'TRN', 'ECO', 'OPP', 'THR', 'REG', 'MAC'],
    'customer': ['PER', 'JTB', 'CJM', 'USE', 'STO', 'PAI', 'GAI', 'EMP', 'FEE', 'INT', 'SUR', 'BEH'],
    'product': ['PRD', 'SVC', 'FEA', 'ROD', 'REQ', 'QUA', 'UXD', 'PER', 'INT', 'SUP'],
    'model': ['BMC', 'REV', 'PRC', 'CST', 'CHN', 'REL', 'RES', 'ACT', 'PRT', 'UNT', 'LTV', 'CAC'],
    'operations': ['PRC', 'WFL', 'ORG', 'ROL', 'TEA', 'SKI', 'POL', 'SLA', 'VND', 'FAC', 'TOO', 'CAP'],
    'technology': ['ARC', 'SYS', 'DAT', 'API', 'INF', 'SEC', 'DEV', 'ANA'],
    'financial': ['FIN', 'BUD', 'FOR', 'FND', 'INV', 'VAL', 'MET', 'REP', 'AUD', 'TAX'],
    'risk': ['RSK', 'MIT', 'CMP', 'GVN', 'CTL', 'CRI', 'ETH', 'STA'],
    'growth': ['GTM', 'GRW', 'SCL', 'EXP', 'INN', 'RND', 'ACQ', 'EXP'],
    'learning': ['DEC', 'LRN', 'RET', 'HYP', 'KNO', 'WIS']
  };

  return domainTypes[domain] || [];
}

function assessDomainCriticality(
  domain: BusinessDomain,
  completeness: number,
  connections: number
): 'high' | 'medium' | 'low' {
  if (domain === 'strategic' || (completeness > 0.8 && connections > 10)) return 'high';
  if (completeness > 0.5 && connections > 5) return 'medium';
  return 'low';
}

function calculateClusteringCoefficient(nodes: DocumentNode[], edges: RelationshipEdge[]): number {
  // Simplified clustering coefficient calculation
  const avgConnections = nodes.reduce((sum, node) => sum + node.inDegree + node.outDegree, 0) / nodes.length;
  return Math.min(avgConnections / nodes.length, 1.0);
}

function findLongestDependencyChain(documents: ParsedBSpecDocument[]): number {
  return Math.max(...documents.map(doc => traceDependencyChain(doc.id, documents, new Set()).length));
}

function traceDependencyChain(docId: string, documents: ParsedBSpecDocument[], visited: Set<string>): string[] {
  if (visited.has(docId)) return [docId]; // Circular reference

  const doc = documents.find(d => d.id === docId);
  if (!doc || !doc.depends_on || doc.depends_on.length === 0) {
    return [docId];
  }

  visited.add(docId);
  let longestChain = [docId];

  for (const depId of doc.depends_on) {
    const depChain = traceDependencyChain(depId, documents, new Set(visited));
    if (depChain.length + 1 > longestChain.length) {
      longestChain = [docId, ...depChain];
    }
  }

  return longestChain;
}

function assessConformanceLevel(documentCount: number): ConformanceLevel {
  if (documentCount >= 45) return 'gold';
  if (documentCount >= 25) return 'silver';
  return 'bronze';
}

function calculateCompletenessScore(documents: ParsedBSpecDocument[]): number {
  const totalPossibleTypes = 82;
  const presentTypes = new Set(documents.map(d => d.type)).size;
  return presentTypes / totalPossibleTypes;
}

function findDirectlyAffectedDocuments(documentId: string, documents: ParsedBSpecDocument[]): string[] {
  return documents
    .filter(doc => doc.depends_on?.includes(documentId))
    .map(doc => doc.id);
}

function findTransitivelyAffectedDocuments(documentId: string, documents: ParsedBSpecDocument[]): string[] {
  const directlyAffected = findDirectlyAffectedDocuments(documentId, documents);
  const transitivelyAffected = new Set<string>();

  function traverse(id: string) {
    const affected = findDirectlyAffectedDocuments(id, documents);
    affected.forEach(affectedId => {
      if (!transitivelyAffected.has(affectedId) && !directlyAffected.includes(affectedId)) {
        transitivelyAffected.add(affectedId);
        traverse(affectedId);
      }
    });
  }

  directlyAffected.forEach(traverse);
  return Array.from(transitivelyAffected);
}

function assessRiskLevel(impactCount: number, totalDocuments: number): 'low' | 'medium' | 'high' | 'critical' {
  const impactRatio = impactCount / totalDocuments;
  if (impactRatio > 0.5) return 'critical';
  if (impactRatio > 0.3) return 'high';
  if (impactRatio > 0.1) return 'medium';
  return 'low';
}

function assessChainRisks(chain: string[], documents: ParsedBSpecDocument[]): string[] {
  const risks: string[] = [];

  if (chain.length > 10) risks.push('Very long dependency chain');
  if (chain[0] === chain[chain.length - 1]) risks.push('Circular dependency detected');

  const chainDocs = documents.filter(doc => chain.includes(doc.id));
  const draftCount = chainDocs.filter(doc => doc.status === 'Draft').length;

  if (draftCount > chain.length * 0.5) {
    risks.push('Many draft documents in chain');
  }

  return risks;
}

function estimateImplementationDuration(path: string[], documents: ParsedBSpecDocument[]): number {
  // Simplified duration estimation based on document complexity
  return path.length * 2; // weeks
}

function findImplementationBlockers(documents: ParsedBSpecDocument[]): string[] {
  return documents
    .filter(doc => doc.status === 'Draft' && (doc.depends_on?.length || 0) > 3)
    .map(doc => doc.id);
}

function generateImplementationRecommendations(documents: ParsedBSpecDocument[]): string[] {
  const recommendations: string[] = [];

  const strategicCount = documents.filter(doc => doc.domain === 'strategic').length;
  if (strategicCount < 4) {
    recommendations.push('Complete strategic foundation documents (MSN, VSN, VAL, STR) first');
  }

  const operationalCount = documents.filter(doc => doc.domain === 'operations').length;
  if (operationalCount < 6) {
    recommendations.push('Define core operational processes and organizational structure');
  }

  return recommendations;
}