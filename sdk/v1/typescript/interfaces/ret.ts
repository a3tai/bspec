/**
 * Retrospective Analysis Document Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 */

import { BaseBSpecDocument } from '../base';

/**
 * Retrospective Analysis (RET)
 *
 * Domain: Learning & Decisions
 * Purpose: The Retrospective Analysis document captures structured reflection on completed projects, initiatives, or time periods to identify successes, failures, lessons learned, and improvement opportunities. It establishes retrospective frameworks that promote continuous learning, team development, and organizational improvement through systematic reflection and analysis.
 */
export interface RETDocument extends BaseBSpecDocument {
  type: "RET";
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

export default RETDocument;
