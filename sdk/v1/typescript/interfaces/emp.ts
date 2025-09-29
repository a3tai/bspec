/**
 * Empathy Maps Document Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 */

import { BaseBSpecDocument } from '../base';

/**
 * Empathy Maps (EMP)
 *
 * Domain: Customer & Value
 * Purpose: The Empathy Maps document captures deep understanding of customer thoughts, feelings, behaviors, and environment. It provides holistic view of customer experience and emotional context that drives behavior.
 */
export interface EMPDocument extends BaseBSpecDocument {
  type: "EMP";
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

export default EMPDocument;
