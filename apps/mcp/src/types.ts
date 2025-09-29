/**
 * Local type definitions for BSpec MCP server
 * These replace the external @bspec/sdk imports to keep everything bundled
 */

// Document Types (82 total)
export type DocumentType =
  // Strategic Foundation (8)
  | 'MSN' | 'VSN' | 'VAL' | 'STR' | 'OBJ' | 'MOT' | 'PUR' | 'THY'
  // Market & Environment (10)
  | 'MKT' | 'SEG' | 'CMP' | 'POS' | 'TRN' | 'ECO' | 'OPP' | 'THR' | 'REG' | 'MAC'
  // Customer & Value (12)
  | 'PER' | 'JTB' | 'CJM' | 'USE' | 'STO' | 'PAI' | 'GAI' | 'EMP' | 'FEE' | 'INT' | 'SUR' | 'BEH'
  // Product & Service (10)
  | 'PRD' | 'SVC' | 'FEA' | 'ROD' | 'REQ' | 'QUA' | 'UXD' | 'PER' | 'INT' | 'SUP'
  // Business Model (12)
  | 'BMC' | 'REV' | 'PRC' | 'CST' | 'CHN' | 'REL' | 'RES' | 'ACT' | 'PRT' | 'UNT' | 'LTV' | 'CAC'
  // Operations & Execution (12)
  | 'PRC' | 'WFL' | 'ORG' | 'ROL' | 'TEA' | 'SKI' | 'POL' | 'SLA' | 'VND' | 'FAC' | 'TOO' | 'CAP'
  // Technology & Data (8)
  | 'ARC' | 'SYS' | 'DAT' | 'API' | 'INF' | 'SEC' | 'DEV' | 'ANA'
  // Financial & Investment (10)
  | 'FIN' | 'BUD' | 'FOR' | 'FND' | 'INV' | 'VAL' | 'MET' | 'REP' | 'AUD' | 'TAX'
  // Risk & Governance (8)
  | 'RSK' | 'MIT' | 'CMP' | 'GVN' | 'CTL' | 'CRI' | 'ETH' | 'STA'
  // Growth & Innovation (8)
  | 'GTM' | 'GRW' | 'SCL' | 'EXP' | 'INN' | 'RND' | 'ACQ' | 'EXP'
  // Learning & Decisions (6)
  | 'DEC' | 'LRN' | 'RET' | 'HYP' | 'KNO' | 'WIS';

// Business Domains (11 total)
export type BusinessDomain =
  | 'strategic'
  | 'market'
  | 'customer'
  | 'product'
  | 'model'
  | 'operations'
  | 'technology'
  | 'financial'
  | 'risk'
  | 'growth'
  | 'learning';

// Document Status
export type DocumentStatus =
  | 'Draft'
  | 'Review'
  | 'Accepted'
  | 'Deprecated';

// Priority Levels
export type Priority =
  | 'critical'
  | 'high'
  | 'medium'
  | 'low';

// Conformance Levels
export type ConformanceLevel =
  | 'bronze'
  | 'silver'
  | 'gold';

// Base BSpec Document Interface
export interface BaseBSpecDocument {
  // Core Identity
  id: string;
  title: string;
  type: DocumentType;
  status: DocumentStatus;
  version: string;

  // Ownership & Responsibility
  owner: string;
  stakeholders?: string[];
  reviewers?: string[];

  // Temporal Metadata
  created: string;
  updated: string;
  review_cycle?: 'monthly' | 'quarterly' | 'annually';

  // Relationship Graph
  depends_on?: string[];
  enables?: string[];
  conflicts_with?: string[];
  related?: string[];

  // Business Context
  domain: BusinessDomain;
  priority?: Priority;
  success_criteria?: string[];

  // Validation & Measurement
  assumptions?: string[];
  constraints?: string[];
  risks?: string[];
  metrics?: string[];

  // Additional optional properties
  supersedes?: string;
  tags?: string[];
}