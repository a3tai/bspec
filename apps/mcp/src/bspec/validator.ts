/**
 * BSpec Document Validator
 *
 * Provides comprehensive validation of BSpec documents using existing TypeScript interfaces
 * Validates structure, relationships, and business rules
 */

import type {
  BaseBSpecDocument,
  DocumentType,
  BusinessDomain,
  DocumentStatus,
  Priority,
  ConformanceLevel
} from '../types.js';
import type { ParsedBSpecDocument } from './parser';

export interface ValidationResult {
  isValid: boolean;
  errors: ValidationError[];
  warnings: ValidationWarning[];
  conformanceLevel?: ConformanceLevel;
}

export interface ValidationError {
  code: string;
  message: string;
  field?: string;
  severity: 'error' | 'warning';
}

export interface ValidationWarning {
  code: string;
  message: string;
  field?: string;
  recommendation?: string;
}

export interface ValidationOptions {
  /** Validate relationships against provided document collection */
  validateRelationships?: boolean;
  /** Documents to validate relationships against */
  documentCollection?: ParsedBSpecDocument[];
  /** Target conformance level to validate against */
  targetConformanceLevel?: ConformanceLevel;
  /** Industry profile for specialized validation */
  industryProfile?: string;
}

/**
 * Comprehensive BSpec document validation
 */
export function validateBSpecDocument(
  document: ParsedBSpecDocument,
  options: ValidationOptions = {}
): ValidationResult {
  const errors: ValidationError[] = [];
  const warnings: ValidationWarning[] = [];

  // Core structure validation
  validateCoreStructure(document, errors);

  // Domain-specific validation
  validateDomainSpecific(document, errors, warnings);

  // Relationship validation
  if (options.validateRelationships && options.documentCollection) {
    validateRelationships(document, options.documentCollection, errors, warnings);
  }

  // Content validation
  validateContent(document, errors, warnings);

  // Conformance level assessment
  const conformanceLevel = assessConformanceLevel(document, options.documentCollection || []);

  return {
    isValid: errors.length === 0,
    errors: [...errors, ...warnings.map(w => ({ ...w, severity: 'warning' as const }))],
    warnings,
    conformanceLevel
  };
}

/**
 * Validate core document structure required for all BSpec documents
 */
function validateCoreStructure(document: ParsedBSpecDocument, errors: ValidationError[]): void {
  // ID format validation
  if (!document.id.match(/^[A-Z]{3}-[a-z0-9-]+$/)) {
    errors.push({
      code: 'INVALID_ID_FORMAT',
      message: 'Document ID must follow format: TYPE-kebab-case-name',
      field: 'id',
      severity: 'error'
    });
  }

  // Type consistency check
  const expectedTypePrefix = document.id.split('-')[0];
  if (document.type !== expectedTypePrefix) {
    errors.push({
      code: 'TYPE_ID_MISMATCH',
      message: `Document type "${document.type}" does not match ID prefix "${expectedTypePrefix}"`,
      field: 'type',
      severity: 'error'
    });
  }

  // Version format validation
  if (!document.version.match(/^\d+\.\d+\.\d+$/)) {
    errors.push({
      code: 'INVALID_VERSION_FORMAT',
      message: 'Version must follow semantic versioning (x.y.z)',
      field: 'version',
      severity: 'error'
    });
  }

  // Date validation
  if (new Date(document.created) > new Date(document.updated)) {
    errors.push({
      code: 'INVALID_DATE_SEQUENCE',
      message: 'Created date cannot be after updated date',
      field: 'updated',
      severity: 'error'
    });
  }

  // Required stakeholder information
  if (!document.owner) {
    errors.push({
      code: 'MISSING_OWNER',
      message: 'Document must have an owner',
      field: 'owner',
      severity: 'error'
    });
  }
}

/**
 * Validate domain-specific requirements
 */
function validateDomainSpecific(
  document: ParsedBSpecDocument,
  errors: ValidationError[],
  warnings: ValidationWarning[]
): void {
  const domainRules = getDomainValidationRules(document.type);

  // Check if document domain matches its type
  if (document.domain && document.domain !== domainRules.expectedDomain) {
    errors.push({
      code: 'DOMAIN_TYPE_MISMATCH',
      message: `Document type ${document.type} should be in ${domainRules.expectedDomain} domain, not ${document.domain}`,
      field: 'domain',
      severity: 'error'
    });
  }

  // Validate required fields for document type
  for (const requiredField of domainRules.requiredFields) {
    if (!document[requiredField as keyof ParsedBSpecDocument]) {
      errors.push({
        code: 'MISSING_REQUIRED_FIELD',
        message: `Field "${requiredField}" is required for ${document.type} documents`,
        field: requiredField,
        severity: 'error'
      });
    }
  }

  // Validate expected relationships
  for (const expectedRelation of domainRules.expectedRelationships) {
    if (!document[expectedRelation.field as keyof ParsedBSpecDocument]) {
      warnings.push({
        code: 'MISSING_EXPECTED_RELATIONSHIP',
        message: `${document.type} documents typically ${expectedRelation.description}`,
        field: expectedRelation.field,
        recommendation: expectedRelation.recommendation
      });
    }
  }
}

/**
 * Validate document relationships and dependencies
 */
function validateRelationships(
  document: ParsedBSpecDocument,
  documentCollection: ParsedBSpecDocument[],
  errors: ValidationError[],
  warnings: ValidationWarning[]
): void {
  const docIds = new Set(documentCollection.map(d => d.id));

  // Validate dependencies exist
  if (document.depends_on) {
    for (const depId of document.depends_on) {
      if (!docIds.has(depId)) {
        errors.push({
          code: 'MISSING_DEPENDENCY',
          message: `Dependency "${depId}" not found in document collection`,
          field: 'depends_on',
          severity: 'error'
        });
      }
    }
  }

  // Validate enablements are valid
  if (document.enables) {
    for (const enabledId of document.enables) {
      if (!docIds.has(enabledId)) {
        warnings.push({
          code: 'MISSING_ENABLED_DOCUMENT',
          message: `Enabled document "${enabledId}" not found in collection`,
          field: 'enables'
        });
      }
    }
  }

  // Check for circular dependencies
  if (document.depends_on) {
    const circularDeps = findCircularDependencies(document, documentCollection);
    if (circularDeps.length > 0) {
      errors.push({
        code: 'CIRCULAR_DEPENDENCY',
        message: `Circular dependency detected: ${circularDeps.join(' -> ')}`,
        field: 'depends_on',
        severity: 'error'
      });
    }
  }

  // Validate conflicts
  if (document.conflicts_with) {
    for (const conflictId of document.conflicts_with) {
      const conflictDoc = documentCollection.find(d => d.id === conflictId);
      if (conflictDoc && conflictDoc.status === 'Accepted') {
        warnings.push({
          code: 'ACTIVE_CONFLICT',
          message: `Document conflicts with accepted document "${conflictId}"`,
          field: 'conflicts_with',
          recommendation: 'Consider updating conflict resolution or document status'
        });
      }
    }
  }
}

/**
 * Validate document content quality and completeness
 */
function validateContent(
  document: ParsedBSpecDocument,
  errors: ValidationError[],
  warnings: ValidationWarning[]
): void {
  // Content length check
  if (document.content.length < 100) {
    warnings.push({
      code: 'MINIMAL_CONTENT',
      message: 'Document content is quite short, consider adding more detail',
      recommendation: 'Add more context, examples, or implementation details'
    });
  }

  // Check for placeholder content
  const placeholderPatterns = [
    /TODO/gi,
    /TBD/gi,
    /\[placeholder\]/gi,
    /\[insert\s+\w+\]/gi
  ];

  for (const pattern of placeholderPatterns) {
    if (pattern.test(document.content)) {
      warnings.push({
        code: 'PLACEHOLDER_CONTENT',
        message: 'Document contains placeholder content',
        recommendation: 'Replace placeholder content with actual information'
      });
      break;
    }
  }

  // Success criteria validation
  if (document.success_criteria && document.success_criteria.length === 0) {
    warnings.push({
      code: 'MISSING_SUCCESS_CRITERIA',
      message: 'Document has no success criteria defined',
      field: 'success_criteria',
      recommendation: 'Add measurable success criteria for validation'
    });
  }
}

/**
 * Assess conformance level based on document collection
 */
function assessConformanceLevel(
  document: ParsedBSpecDocument,
  documentCollection: ParsedBSpecDocument[]
): ConformanceLevel {
  const docCount = documentCollection.length;

  if (docCount >= 45) return 'gold';
  if (docCount >= 25) return 'silver';
  if (docCount >= 12) return 'bronze';

  return 'bronze'; // Default minimum
}

/**
 * Find circular dependencies in document relationships
 */
function findCircularDependencies(
  document: ParsedBSpecDocument,
  documentCollection: ParsedBSpecDocument[]
): string[] {
  const visited = new Set<string>();
  const path: string[] = [];

  function dfs(docId: string): string[] | null {
    if (path.includes(docId)) {
      return path.slice(path.indexOf(docId)).concat(docId);
    }

    if (visited.has(docId)) {
      return null;
    }

    visited.add(docId);
    path.push(docId);

    const doc = documentCollection.find(d => d.id === docId);
    if (doc?.depends_on) {
      for (const depId of doc.depends_on) {
        const cycle = dfs(depId);
        if (cycle) return cycle;
      }
    }

    path.pop();
    return null;
  }

  return dfs(document.id) || [];
}

/**
 * Get validation rules for specific document types
 */
function getDomainValidationRules(type: DocumentType): {
  expectedDomain: BusinessDomain;
  requiredFields: string[];
  expectedRelationships: Array<{
    field: string;
    description: string;
    recommendation: string;
  }>;
} {
  const domainMap: Record<string, BusinessDomain> = {
    // Strategic Foundation
    'MSN': 'strategic', 'VSN': 'strategic', 'VAL': 'strategic', 'STR': 'strategic',
    'OBJ': 'strategic', 'MOT': 'strategic', 'PUR': 'strategic', 'THY': 'strategic',
    // Market & Environment
    'MKT': 'market', 'SEG': 'market', 'CMP': 'market', 'POS': 'market',
    'TRN': 'market', 'ECO': 'market', 'OPP': 'market', 'THR': 'market',
    'REG': 'market', 'MAC': 'market',
    // Customer & Value
    'PER': 'customer', 'JTB': 'customer', 'CJM': 'customer', 'USE': 'customer',
    'STO': 'customer', 'PAI': 'customer', 'GAI': 'customer', 'EMP': 'customer',
    'FEE': 'customer', 'INT': 'customer', 'SUR': 'customer', 'BEH': 'customer'
  };

  return {
    expectedDomain: domainMap[type] || 'strategic',
    requiredFields: ['id', 'title', 'type', 'status', 'version', 'owner'],
    expectedRelationships: getExpectedRelationships(type)
  };
}

/**
 * Get expected relationships for document types
 */
function getExpectedRelationships(type: DocumentType): Array<{
  field: string;
  description: string;
  recommendation: string;
}> {
  const relationships: Record<string, Array<{ field: string; description: string; recommendation: string; }>> = {
    'STR': [{
      field: 'depends_on',
      description: 'depend on Mission (MSN) and Vision (VSN)',
      recommendation: 'Link to MSN and VSN documents'
    }],
    'OBJ': [{
      field: 'depends_on',
      description: 'depend on Strategy (STR)',
      recommendation: 'Link to relevant STR document'
    }],
    'PER': [{
      field: 'enables',
      description: 'enable Customer Journey Maps (CJM)',
      recommendation: 'Link to CJM documents that use these personas'
    }]
  };

  return relationships[type] || [];
}

/**
 * Validate a collection of documents for overall consistency
 */
export function validateDocumentCollection(
  documents: ParsedBSpecDocument[],
  options: ValidationOptions = {}
): {
  overallValid: boolean;
  documentResults: Array<{ document: ParsedBSpecDocument; validation: ValidationResult; }>;
  collectionErrors: ValidationError[];
  conformanceLevel: ConformanceLevel;
} {
  const documentResults = documents.map(document => ({
    document,
    validation: validateBSpecDocument(document, { ...options, documentCollection: documents })
  }));

  const collectionErrors: ValidationError[] = [];

  // Check for duplicate IDs
  const idCounts = new Map<string, number>();
  documents.forEach(doc => {
    idCounts.set(doc.id, (idCounts.get(doc.id) || 0) + 1);
  });

  idCounts.forEach((count, id) => {
    if (count > 1) {
      collectionErrors.push({
        code: 'DUPLICATE_DOCUMENT_ID',
        message: `Duplicate document ID found: ${id}`,
        severity: 'error'
      });
    }
  });

  const overallValid = documentResults.every(r => r.validation.isValid) && collectionErrors.length === 0;
  const conformanceLevel = assessConformanceLevel(documents[0], documents);

  return {
    overallValid,
    documentResults,
    collectionErrors,
    conformanceLevel
  };
}