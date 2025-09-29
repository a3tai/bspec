/**
 * Hypothesis Management Document Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 */

import { BaseBSpecDocument } from '../base';

/**
 * Hypothesis Management (HYP)
 *
 * Domain: Learning & Decisions
 * Purpose: The Hypothesis Management document captures testable assumptions and beliefs about business, customers, markets, and solutions that guide organizational decision-making and learning. It establishes hypothesis frameworks that enable evidence-based validation, systematic experimentation, and continuous learning through iterative hypothesis development and testing.
 */
export interface HYPDocument extends BaseBSpecDocument {
  type: "HYP";
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

export default HYPDocument;
