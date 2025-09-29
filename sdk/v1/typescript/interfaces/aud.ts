/**
 * Audit Document Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 */

import { BaseBSpecDocument } from '../base';

/**
 * Audit (AUD)
 *
 * Domain: Financial & Investment
 * Purpose: The Audit document defines systematic approaches to internal and external audit processes that provide independent assurance on financial reporting, internal controls, and operational effectiveness. It establishes audit frameworks that ensure compliance, risk management, and continuous improvement in business operations.
 */
export interface AUDDocument extends BaseBSpecDocument {
  type: "AUD";
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

export default AUDDocument;
