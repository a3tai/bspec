/**
 * Opportunities Document Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 */

import { BaseBSpecDocument } from '../base';

/**
 * Opportunities (OPP)
 *
 * Domain: Market & Environment
 * Purpose: The Opportunities document identifies market gaps, growth potential, and strategic opportunities available to the organization. It analyzes opportunity attractiveness and prioritizes pursuit strategies.
 */
export interface OPPDocument extends BaseBSpecDocument {
  type: "OPP";
  domain: "Market & Environment";

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

export default OPPDocument;
