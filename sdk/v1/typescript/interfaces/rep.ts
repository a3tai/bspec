/**
 * Reporting Document Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 */

import { BaseBSpecDocument } from '../base';

/**
 * Reporting (REP)
 *
 * Domain: Financial & Investment
 * Purpose: The Reporting document defines systematic approaches to financial and business reporting that provide stakeholders with accurate, timely, and relevant information for decision making. It establishes reporting frameworks that ensure regulatory compliance, transparency, and effective communication of business performance.
 */
export interface REPDocument extends BaseBSpecDocument {
  type: "REP";
  domain: "Financial & Investment";

  // Content sections
  executive_summary?: string;
  framework?: any;
  implementation?: any;
  validation?: any;
  quality_standards?: {
    bronze?: string[];
    silver?: string[];
    gold?: string[];
  };

  // Domain-specific fields can be added here
  [key: string]: any;
}

export default REPDocument;
