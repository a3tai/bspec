/**
 * Key Partnerships Document Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 */

import { BaseBSpecDocument } from '../base';

/**
 * Key Partnerships (KPT)
 *
 * Domain: Business Model
 * Purpose: The Key Partnerships defines systematic approaches to strategic alliances and partnerships that create mutual value and competitive advantages. It establishes partnership frameworks that optimize collaboration, resource sharing, and market access while minimizing risks and costs.
 */
export interface KPTDocument extends BaseBSpecDocument {
  type: "KPT";
  domain: "Business Model";

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

export default KPTDocument;
