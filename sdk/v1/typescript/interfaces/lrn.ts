/**
 * Learning Records Document Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 */

import { BaseBSpecDocument } from '../base';

/**
 * Learning Records (LRN)
 *
 * Domain: Learning & Decisions
 * Purpose: The Learning Records document captures key discoveries, insights, and knowledge gained through organizational activities, experiments, and experiences. It establishes learning frameworks that preserve institutional knowledge, accelerate organizational learning, and enable evidence-based decision-making and continuous improvement.
 */
export interface LRNDocument extends BaseBSpecDocument {
  type: "LRN";
  domain: "Learning & Decisions";

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

export default LRNDocument;
