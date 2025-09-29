/**
 * Pain Points Document Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 */

import { BaseBSpecDocument } from '../base';

/**
 * Pain Points (PAI)
 *
 * Domain: Customer & Value
 * Purpose: The Pain Points document identifies and analyzes customer problems, frustrations, and obstacles. It captures the negative experiences that drive customers to seek solutions and creates opportunities for value creation.
 */
export interface PAIDocument extends BaseBSpecDocument {
  type: "PAI";
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

export default PAIDocument;
