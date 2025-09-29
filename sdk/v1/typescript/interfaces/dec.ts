/**
 * Decision Records Document Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 */

import { BaseBSpecDocument } from '../base';

/**
 * Decision Records (DEC)
 *
 * Domain: Learning & Decisions
 * Purpose: The Decision Records document captures strategic and operational decisions made within the organization, including context, rationale, alternatives considered, and outcomes. It establishes decision frameworks that preserve institutional knowledge, enable accountability, and provide historical context for future decision-making.
 */
export interface DECDocument extends BaseBSpecDocument {
  type: "DEC";
  domain: "Learning & Decisions";

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

export default DECDocument;
