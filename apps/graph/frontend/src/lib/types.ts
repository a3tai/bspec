export type Domain = {
  id: string;
  slug: string;
  name: string;
  description: string;
  documentCount: number;
  outgoingEdgeCount: number;
  incomingEdgeCount: number;
};

export type RelationshipRef = {
  code?: string;
  pattern: string;
  documentId?: string;
};

export type DocumentNode = {
  id: string;
  kind: string;
  documentId: string;
  code: string;
  name: string;
  domainSlug: string;
  domainName: string;
  status: string;
  version: string;
  lastUpdated: string;
  owner: string;
  sourcePath: string;
  abstract: string;
  purpose: string;
  body: string;
  relationshipGuidance: string;
  relationships: Record<string, RelationshipRef[]>;
  outgoingCount: number;
};

export type Edge = {
  id: string;
  source: string;
  sourceCode: string;
  sourceDocumentId: string;
  sourceName: string;
  sourceDomainSlug: string;
  target: string;
  targetCode: string;
  targetDocumentId: string;
  targetName: string;
  targetDomainSlug: string;
  targetPattern: string;
  field: string;
  label: string;
  fact: string;
};

export type NamedCount = {
  name: string;
  count: number;
};

export type DomainCoverageItem = {
  domainSlug: string;
  domainName: string;
  count: number;
  silverTarget: number;
  goldTarget: number;
  silverMissing: number;
  goldMissing: number;
};

export type ImplementationIssue = {
  severity: string;
  kind: string;
  documentId: string;
  document: string;
  sourcePath: string;
  message: string;
};

export type ImplementationSummary = {
  name: string;
  description: string;
  author: string;
  formatVersion: string;
  bspecVersion: string;
  created: string;
  documentCount: number;
  domainCount: number;
  relationshipCount: number;
  brokenReferenceCount: number;
  externalReferenceCount: number;
  missingRequiredFieldCount: number;
  invalidTypeCount: number;
  invalidStatusCount: number;
  duplicateIdCount: number;
  orphanDocumentCount: number;
  acceptedCount: number;
  draftCount: number;
  conformanceLevel: string;
  nextConformanceLevel: string;
  conformanceProgress: number;
  missingBronzeTypes: string[];
  missingSilverDomains: string[];
  statusCounts: NamedCount[];
  ownerCounts: NamedCount[];
  typeCounts: NamedCount[];
  domainCoverage: DomainCoverageItem[];
  issues: ImplementationIssue[];
};

export type GraphData = {
  schemaVersion: number;
  sourceKind: string;
  title: string;
  description: string;
  generatedAt: string;
  version: string;
  groupId: string;
  sourceRoot: string;
  manifest?: Record<string, unknown>;
  implementation?: ImplementationSummary | null;
  nodeLabelSingular: string;
  nodeLabelPlural: string;
  domainHeading: string;
  graphitiCoreEpisodes: number;
  graphitiRelationshipRuleEpisodes: number;
  domains: Domain[];
  documents: DocumentNode[];
  edges: Edge[];
  relationCounts: { field: string; label: string; count: number }[];
  domainSeries: number[];
};

export type Mode = 'overview' | 'graph' | 'documents' | 'assistant' | 'issues' | 'facts' | 'episodes';
export type InspectorMode = 'context' | 'assistant';
export type DocumentScope = 'all' | 'issues' | 'drafts' | 'accepted' | 'orphans';
export type ChatRole = 'user' | 'assistant';

export type ChatGrounding = {
  id: string;
  kind: string;
  title: string;
  sourcePath: string;
  snippet: string;
};

export type ChatSuggestion = {
  label: string;
  prompt: string;
};

export type ChatResponse = {
  answer: string;
  draftTitle: string;
  draftDocumentId: string;
  draftDocumentType: string;
  draftMarkdown: string;
  canSaveDraft: boolean;
  grounding: ChatGrounding[];
  suggestedActions: ChatSuggestion[];
  graphitiEpisode: Record<string, unknown>;
};

export type AssistantMessage = {
  id: string;
  role: ChatRole;
  content: string;
  draftTitle?: string;
  draftDocumentId?: string;
  draftDocumentType?: string;
  draftMarkdown?: string;
  canSaveDraft?: boolean;
  grounding?: ChatGrounding[];
  suggestions?: ChatSuggestion[];
  graphitiEpisode?: Record<string, unknown>;
};

export const emptyImplementation: ImplementationSummary = {
  name: '',
  description: '',
  author: '',
  formatVersion: '',
  bspecVersion: '',
  created: '',
  documentCount: 0,
  domainCount: 0,
  relationshipCount: 0,
  brokenReferenceCount: 0,
  externalReferenceCount: 0,
  missingRequiredFieldCount: 0,
  invalidTypeCount: 0,
  invalidStatusCount: 0,
  duplicateIdCount: 0,
  orphanDocumentCount: 0,
  acceptedCount: 0,
  draftCount: 0,
  conformanceLevel: 'Not loaded',
  nextConformanceLevel: 'Bronze',
  conformanceProgress: 0,
  missingBronzeTypes: [],
  missingSilverDomains: [],
  statusCounts: [],
  ownerCounts: [],
  typeCounts: [],
  domainCoverage: [],
  issues: [],
};

export const emptyData: GraphData = {
  schemaVersion: 1,
  sourceKind: 'empty',
  title: 'BSpec Implementation Explorer',
  description: '',
  generatedAt: '',
  version: '',
  groupId: '',
  sourceRoot: '',
  nodeLabelSingular: 'Document',
  nodeLabelPlural: 'Documents',
  domainHeading: 'Implementation Domains',
  graphitiCoreEpisodes: 0,
  graphitiRelationshipRuleEpisodes: 0,
  domains: [],
  documents: [],
  edges: [],
  relationCounts: [],
  domainSeries: [],
};
