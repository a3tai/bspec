/**
 * Product Requirements Document Document Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 */

import { BaseBSpecDocument } from '../base';

/**
 * Product Requirements Document (PRD)
 *
 * Domain: Product & Service
 * Purpose: The Product Requirements Document defines comprehensive requirements for products, bridging customer needs with technical implementation. It serves as the authoritative specification for product development teams and stakeholders.
 */
export interface PRDDocument extends BaseBSpecDocument {
  type: "PRD";
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

export default PRDDocument;
