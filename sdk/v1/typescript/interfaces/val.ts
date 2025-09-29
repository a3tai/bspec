/**
 * Organizational Values Document Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 */

import { BaseBSpecDocument } from '../base';

/**
 * Organizational Values (VAL)
 *
 * Domain: Strategic Foundation
 * Purpose: The Values document defines the guiding principles that shape culture, guide decisions, and determine how the organization behaves. Values are the non-negotiable beliefs that influence every action.
 */
export interface VALDocument extends BaseBSpecDocument {
  type: "VAL";
  domain: "Strategic Foundation";

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

export default VALDocument;
