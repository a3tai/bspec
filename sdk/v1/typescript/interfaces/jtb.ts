/**
 * Jobs-to-be-Done Document Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 */

import { BaseBSpecDocument } from '../base';

/**
 * Jobs-to-be-Done (JTB)
 *
 * Domain: Customer & Value
 * Purpose: The Jobs-to-be-Done document defines the specific outcomes customers hire products or services to achieve. It captures the functional, emotional, and social jobs that drive customer behavior and innovation opportunities.
 */
export interface JTBDocument extends BaseBSpecDocument {
  type: "JTB";
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

export default JTBDocument;
