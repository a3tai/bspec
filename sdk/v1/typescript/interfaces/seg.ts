/**
 * Market Segments Document Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 */

import { BaseBSpecDocument } from '../base';

/**
 * Market Segments (SEG)
 *
 * Domain: Market & Environment
 * Purpose: The Market Segments document identifies and analyzes distinct customer groups within the broader market. It defines segmentation criteria, profiles key segments, and prioritizes target segments.
 */
export interface SEGDocument extends BaseBSpecDocument {
  type: "SEG";
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

export default SEGDocument;
