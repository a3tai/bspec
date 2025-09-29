/**
 * Competitive Analysis Document Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 */

import { BaseBSpecDocument } from '../base';

/**
 * Competitive Analysis (CMP)
 *
 * Domain: Market & Environment
 * Purpose: The Competitive Analysis document maps the competitive landscape, analyzes key competitors, and identifies competitive threats and opportunities. It provides intelligence for strategic positioning and tactical responses.
 */
export interface CMPDocument extends BaseBSpecDocument {
  type: "CMP";
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

export default CMPDocument;
