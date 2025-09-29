/**
 * Surveys Document Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 */

import { BaseBSpecDocument } from '../base';

/**
 * Surveys (SUR)
 *
 * Domain: Customer & Value
 * Purpose: The Surveys document captures quantitative customer research through structured questionnaires that provide statistical insights into customer attitudes, behaviors, preferences, and satisfaction levels.
 */
export interface SURDocument extends BaseBSpecDocument {
  type: "SUR";
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

export default SURDocument;
