/**
 * Key Resources Document Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 */

import { BaseBSpecDocument } from '../base';

/**
 * Key Resources (KRS)
 *
 * Domain: Business Model
 * Purpose: The Key Resources defines systematic identification, development, and management of critical organizational assets that create competitive advantage and enable business strategy execution. It establishes resource frameworks that optimize asset utilization, investment allocation, and capability building.
 */
export interface KRSDocument extends BaseBSpecDocument {
  type: "KRS";
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

export default KRSDocument;
