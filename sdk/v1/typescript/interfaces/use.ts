/**
 * Use Cases Document Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 */

import { BaseBSpecDocument } from '../base';

/**
 * Use Cases (USE)
 *
 * Domain: Customer & Value
 * Purpose: The Use Cases document describes specific scenarios where customers apply the solution. It details the context, actors, preconditions, and steps involved in successful solution usage.
 */
export interface USEDocument extends BaseBSpecDocument {
  type: "USE";
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

export default USEDocument;
