/**
 * Personas Document Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 */

import { BaseBSpecDocument } from '../base';

/**
 * Personas (PER)
 *
 * Domain: Customer & Value
 * Purpose: The Personas document creates detailed archetypal representations of key customer segments. Personas humanize customer data and provide shared understanding of who the organization serves.
 */
export interface PERDocument extends BaseBSpecDocument {
  type: "PER";
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

export default PERDocument;
