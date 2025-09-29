/**
 * Brand Positioning Document Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 */

import { BaseBSpecDocument } from '../base';

/**
 * Brand Positioning (BPO)
 *
 * Domain: Brand & Marketing
 * Purpose: The Brand Positioning document defines how the brand occupies a distinctive position in customers' minds relative to competitors. It establishes positioning frameworks that create clear differentiation, build competitive advantage, and guide all marketing communications to ensure consistent market perception and customer preference.
 */
export interface BPODocument extends BaseBSpecDocument {
  type: "BPO";
  domain: "Brand & Marketing";

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

export default BPODocument;
