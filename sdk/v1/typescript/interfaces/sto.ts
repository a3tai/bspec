/**
 * User Stories Document Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 */

import { BaseBSpecDocument } from '../base';

/**
 * User Stories (STO)
 *
 * Domain: Customer & Value
 * Purpose: The User Stories document captures individual requirements from the user perspective using the standard "As a... I want... So that..." format. Stories decompose features into implementable units of customer value.
 */
export interface STODocument extends BaseBSpecDocument {
  type: "STO";
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

export default STODocument;
