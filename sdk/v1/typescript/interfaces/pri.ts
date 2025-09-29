/**
 * Pricing Strategy Document Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 */

import { BaseBSpecDocument } from '../base';

/**
 * Pricing Strategy (PRI)
 *
 * Domain: Business Model
 * Purpose: The Pricing Strategy defines systematic approaches to product and service pricing that optimize value capture while supporting competitive positioning and business objectives. It establishes pricing frameworks that balance customer value, market dynamics, and financial goals.
 */
export interface PRIDocument extends BaseBSpecDocument {
  type: "PRI";
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

export default PRIDocument;
