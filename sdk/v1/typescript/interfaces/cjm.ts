/**
 * Customer Journey Map Document Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 */

import { BaseBSpecDocument } from '../base';

/**
 * Customer Journey Map (CJM)
 *
 * Domain: Customer & Value
 * Purpose: The Customer Journey Map document visualizes the end-to-end customer experience from awareness to advocacy. It identifies touchpoints, emotions, pain points, and opportunities across the entire customer lifecycle.
 */
export interface CJMDocument extends BaseBSpecDocument {
  type: "CJM";
  domain: "Customer & Value";

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

export default CJMDocument;
