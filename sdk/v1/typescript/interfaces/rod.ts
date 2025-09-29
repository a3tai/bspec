/**
 * Roadmap Document Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 */

import { BaseBSpecDocument } from '../base';

/**
 * Roadmap (ROD)
 *
 * Domain: Product & Service
 * Purpose: The Roadmap document defines strategic product and technology direction over multiple time horizons, aligning development efforts with business objectives while managing resource constraints and market dynamics. It provides stakeholders with visibility into planned evolution and investment priorities.
 */
export interface RODDocument extends BaseBSpecDocument {
  type: "ROD";
  domain: "Product & Service";

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

export default RODDocument;
