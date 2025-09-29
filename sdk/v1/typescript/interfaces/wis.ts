/**
 * Wisdom Synthesis Document Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 */

import { BaseBSpecDocument } from '../base';

/**
 * Wisdom Synthesis (WIS)
 *
 * Domain: Learning & Decisions
 * Purpose: The Wisdom Synthesis document captures the practical wisdom, judgment principles, and synthesized insights that guide organizational decision-making and action. It establishes wisdom frameworks that combine knowledge, experience, and judgment to create actionable guidance for complex situations, ethical dilemmas, and strategic choices.
 */
export interface WISDocument extends BaseBSpecDocument {
  type: "WIS";
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

export default WISDocument;
