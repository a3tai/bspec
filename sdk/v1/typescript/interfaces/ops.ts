/**
 * Operations Manual Document Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 */

import { BaseBSpecDocument } from '../base';

/**
 * Operations Manual (OPS)
 *
 * Domain: Operations & Execution
 * Purpose: The Operations Manual defines systematic approaches to day-to-day operational management that ensure consistent service delivery, efficient resource utilization, and continuous operational excellence. It establishes operational frameworks that optimize performance and reliability.
 */
export interface OPSDocument extends BaseBSpecDocument {
  type: "OPS";
  domain: "Operations & Execution";

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

export default OPSDocument;
