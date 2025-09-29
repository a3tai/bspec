/**
 * Financial Model Document Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 */

import { BaseBSpecDocument } from '../base';

/**
 * Financial Model (FIN)
 *
 * Domain: Financial & Investment
 * Purpose: The Financial Model document defines comprehensive financial projections and planning models that forecast business performance through detailed P&L, balance sheet, and cash flow analysis. It establishes financial frameworks that enable strategic planning, investment decisions, and performance management.
 */
export interface FINDocument extends BaseBSpecDocument {
  type: "FIN";
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

export default FINDocument;
