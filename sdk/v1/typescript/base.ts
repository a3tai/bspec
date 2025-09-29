/**
 * BSpec TypeScript SDK - Base Interfaces
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 * Generated at: 2025-09-29T01:32:41.030729
 * Generator: typescript-generator-v1.0.0
 */

// Base document interface
export interface BaseBSpecDocument {
  id: string;
  title: string;
  type: DocumentType;
  status: DocumentStatus;
  version: string;
  owner: string;
  domain: BusinessDomain;
  depends_on?: string[];
  enables?: string[];
  related?: string[];
  stakeholders?: string[];
  priority?: Priority;
  review_cycle?: string;
  constraints?: string[];
  assumptions?: string[];
  success_criteria?: string[];
}

// Document types
export type DocumentType = "PRD" | "QUA" | "ROD" | "UXD" | "REQ" | "SUP" | "FEA" | "SVC" | "INT" | "PSP" | "SYS" | "DEV" | "DAT" | "ANA" | "SEC" | "ARC" | "INF" | "API" | "PRI" | "VST" | "KRS" | "KPT" | "REV" | "CST" | "CUS" | "CHN" | "KAC" | "LEA" | "ADT" | "EXP" | "INN" | "RND" | "FUT" | "IGN" | "PAI" | "BEH" | "CJM" | "PER" | "CIN" | "STO" | "EMP" | "GAI" | "FEE" | "USE" | "SUR" | "JTB" | "LRN" | "DEC" | "WIS" | "HYP" | "RET" | "KNO" | "MSN" | "OBJ" | "STR" | "VSN" | "PUR" | "VAL" | "MOT" | "THY" | "CTL" | "LEG" | "GOV" | "COM" | "INC" | "INS" | "RSK" | "ETH" | "OPP" | "ECO" | "SEG" | "MKT" | "POS" | "REG" | "MAC" | "THR" | "CMP" | "TRN" | "TON" | "VID" | "BRD" | "CAM" | "CNT" | "PRF" | "MSG" | "IFL" | "MCH" | "SEO" | "SOC" | "BPO" | "REP" | "FOR" | "TAX" | "FIN" | "FND" | "INV" | "AUD" | "MET" | "BUD" | "VLU" | "POL" | "SKI" | "VND" | "PRO" | "CAP" | "WFL" | "FAC" | "SLA" | "ROL" | "OPS" | "ORG" | "TEA";

// Document status
export type DocumentStatus = "Draft" | "Review" | "Approved" | "Active" | "Deprecated";

// Business domains
export type BusinessDomain = "Product & Service" | "Technology & Data" | "{Business domain or bounded context}" | "{Security domain or control area}" | "{Technical domain or system boundary}" | "Business Model" | "Growth & Innovation" | "Customer & Value" | "Learning & Decisions" | "Strategic Foundation" | "Risk & Governance" | "Market & Environment" | "Brand & Marketing" | "Financial & Investment" | "Operations & Execution";

// Priority levels
export type Priority = "Critical" | "High" | "Medium" | "Low";

// Conformance levels
export type ConformanceLevel = "bronze" | "silver" | "gold";

// Changelog entry
export interface ChangelogEntry {
  version: string;
  date: string;
  changes: string[];
  author: string;
}

// BSpec metadata
export interface BSpecMetadata {
  version: string;
  generated_at: string;
  generator: string;
  total_document_types: number;
  total_domains: number;
}
