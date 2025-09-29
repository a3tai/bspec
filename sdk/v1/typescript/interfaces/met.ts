/**
 * Metrics Document Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 */

import { BaseBSpecDocument } from '../base';

/**
 * Metrics (MET)
 *
 * Domain: Financial & Investment
 * Purpose: The Metrics document defines systematic approaches to measuring, monitoring, and managing business performance through key performance indicators, financial metrics, and operational measures. It establishes measurement frameworks that enable data-driven decision making, performance accountability, and continuous improvement.
 */
export interface METDocument extends BaseBSpecDocument {
  type: "MET";
  domain: "Financial & Investment";

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

export default METDocument;
