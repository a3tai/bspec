/**
 * BSpec TypeScript SDK - Core Types
 * Generated from BSpec v1.0.0 JSON SDK
 */

/**
 * Main BSpec data structure containing the entire specification
 */
export interface BSpecData {
  metadata: BSpecMetadata;
  files: Record<string, BSpecFile>;
  directory_structure: string[][];
  document_index: BSpecDocumentIndex[];
  domains: BSpecDomain[];
  statistics: BSpecStatistics;
  document_types: BSpecDocumentType[];
  yaml_schema: BSpecYAMLSchema;
  conformance_levels: BSpecConformanceLevel[];
}

/**
 * Metadata about the BSpec generation
 */
export interface BSpecMetadata {
  bspec_version: string;
  generated_at: string;
  generator: string;
  source_directory: string;
}

/**
 * Represents a file in the BSpec specification
 */
export interface BSpecFile {
  path: string;
  name: string;
  type: 'markdown' | 'yaml' | 'json' | 'text' | 'binary' | 'error';
  extension: string;
  size?: number;
  modified?: string;
  has_frontmatter?: boolean;
  frontmatter?: Record<string, any>;
  content?: string;
  data?: any;
  error?: string;
  yaml_error?: string;
  json_error?: string;
}

/**
 * Document index entry with frontmatter metadata
 */
export interface BSpecDocumentIndex {
  file_path: string;
  frontmatter: Record<string, any>;
  title: string;
  type?: string;
  domain?: string;
}

/**
 * Business domain containing document types
 */
export interface BSpecDomain {
  name: string;
  description: string;
  document_types: string[];
  document_count: number;
  documents?: BSpecDomainDocument[];
}

/**
 * Document within a domain
 */
export interface BSpecDomainDocument {
  file_path: string;
  title: string;
  type: string;
}

/**
 * Statistics about the BSpec specification
 */
export interface BSpecStatistics {
  total_files: number;
  total_documents: number;
  total_domains: number;
  total_document_types: number;
  markdown_files: number;
  yaml_files: number;
  other_files: number;
}

/**
 * BSpec document type definition
 */
export interface BSpecDocumentType {
  code: string;
  name: string;
  purpose: string;
  domain: string;
}

/**
 * YAML schema definition for BSpec documents
 */
export interface BSpecYAMLSchema {
  required_fields: string[];
  optional_fields: string[];
}

/**
 * Conformance level definition
 */
export interface BSpecConformanceLevel {
  name: string;
  description: string;
}

/**
 * Document status types
 */
export type DocumentStatus = 'Draft' | 'Review' | 'Accepted' | 'Deprecated' | 'Archived';

/**
 * Business domain names
 */
export type BusinessDomainName =
  | 'Strategic Foundation'
  | 'Market & Environment'
  | 'Customer & Value'
  | 'Product & Service'
  | 'Business Model'
  | 'Operations & Execution'
  | 'Technology & Data'
  | 'Financial & Investment'
  | 'Risk & Governance'
  | 'Growth & Innovation'
  | 'Brand & Marketing'
  | 'Learning & Decisions';

/**
 * Conformance level names
 */
export type ConformanceLevelName = 'Bronze' | 'Silver' | 'Gold';

/**
 * Document type codes (all 82 types)
 */
export type DocumentTypeCode =
  // Strategic Foundation
  | 'MSN' | 'VSN' | 'VAL' | 'STR' | 'OBJ' | 'MOT' | 'PUR' | 'THY'
  // Market & Environment
  | 'MKT' | 'SEG' | 'CMP' | 'POS' | 'TRN' | 'ECO' | 'OPP' | 'THR' | 'REG' | 'MAC'
  // Customer & Value
  | 'PER' | 'JTB' | 'CJM' | 'USE' | 'STO' | 'PAI' | 'GAI' | 'EMP' | 'FEE' | 'INT' | 'SUR' | 'BEH'
  // Product & Service
  | 'PRD' | 'SVC' | 'FEA' | 'ROD' | 'REQ' | 'QUA' | 'UXD' | 'SUP'
  // Business Model
  | 'REV' | 'PRI' | 'CHN' | 'CST' | 'KPT' | 'KRS' | 'KAC' | 'VST' | 'CUS'
  // Operations & Execution
  | 'PRO' | 'CAP' | 'SLA' | 'OPS' | 'WFL' | 'ORG' | 'ROL' | 'TEA' | 'SKI' | 'POL' | 'VND' | 'FAC'
  // Technology & Data
  | 'ARC' | 'SYS' | 'DAT' | 'API' | 'INF' | 'SEC' | 'DEV' | 'ANA'
  // Financial & Investment
  | 'FIN' | 'BUD' | 'FOR' | 'FND' | 'INV' | 'VLU' | 'MET' | 'REP' | 'AUD' | 'TAX'
  // Risk & Governance
  | 'RSK' | 'CTL' | 'COM' | 'GOV' | 'ETH' | 'INC' | 'LEG' | 'INS'
  // Growth & Innovation
  | 'INN' | 'LEA' | 'EXP' | 'IGN' | 'RND' | 'KNO' | 'ADT' | 'FUT'
  // Brand & Marketing
  | 'BRD' | 'VID' | 'MSG' | 'TON' | 'CNT' | 'SEO' | 'CAM' | 'SOC' | 'INF' | 'PRF'
  // Learning & Decisions
  | 'DEC' | 'LRN' | 'RET' | 'HYP' | 'WIS';