/**
 * Regulatory Environment Document Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 */

import { BaseBSpecDocument } from '../base';

/**
 * Regulatory Environment (REG)
 *
 * Domain: Market & Environment
 * Purpose: The Regulatory Environment document analyzes laws, regulations, and compliance requirements affecting the business. It tracks regulatory changes and assesses compliance obligations.
 */
export interface REGDocument extends BaseBSpecDocument {
  type: "REG";
  domain: "Market & Environment";

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

export default REGDocument;
