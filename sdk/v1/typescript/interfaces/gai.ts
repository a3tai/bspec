/**
 * Gains Document Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 */

import { BaseBSpecDocument } from '../base';

/**
 * Gains (GAI)
 *
 * Domain: Customer & Value
 * Purpose: The Gains document identifies and analyzes the positive outcomes, benefits, and value that customers achieve or seek. It captures the upside potential that motivates customer behavior and creates opportunities for value delivery.
 */
export interface GAIDocument extends BaseBSpecDocument {
  type: "GAI";
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

export default GAIDocument;
