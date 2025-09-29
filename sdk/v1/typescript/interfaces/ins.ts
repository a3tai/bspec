/**
 * Insurance Document Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 */

import { BaseBSpecDocument } from '../base';

/**
 * Insurance (INS)
 *
 * Domain: Risk & Governance
 * Purpose: The Insurance document defines systematic approaches to managing insurance programs that transfer financial risks and protect organizational assets, operations, and stakeholders. It establishes insurance frameworks that optimize risk transfer, minimize total cost of risk, and ensure adequate protection against potential losses.
 */
export interface INSDocument extends BaseBSpecDocument {
  type: "INS";
  domain: "Risk & Governance";

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

export default INSDocument;
