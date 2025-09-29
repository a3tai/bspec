/**
 * Budget Document Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 */

import { BaseBSpecDocument } from '../base';

/**
 * Budget (BUD)
 *
 * Domain: Financial & Investment
 * Purpose: The Budget document defines systematic resource allocation and spending plans that translate strategic objectives into financial commitments. It establishes budgeting frameworks that ensure disciplined resource management, performance accountability, and financial control.
 */
export interface BUDDocument extends BaseBSpecDocument {
  type: "BUD";
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

export default BUDDocument;
