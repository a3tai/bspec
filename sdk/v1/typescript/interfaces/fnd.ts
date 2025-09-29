/**
 * Funding Document Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 */

import { BaseBSpecDocument } from '../base';

/**
 * Funding (FND)
 *
 * Domain: Financial & Investment
 * Purpose: The Funding document defines systematic approaches to raising capital and securing financial resources to support business operations, growth initiatives, and strategic investments. It establishes funding frameworks that optimize capital structure, minimize cost of capital, and align funding strategies with business objectives.
 */
export interface FNDDocument extends BaseBSpecDocument {
  type: "FND";
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

export default FNDDocument;
