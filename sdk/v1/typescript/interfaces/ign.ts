/**
 * Insight Generation Document Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 */

import { BaseBSpecDocument } from '../base';

/**
 * Insight Generation (IGN)
 *
 * Domain: Growth & Innovation
 * Purpose: The Insight Generation document defines systematic approaches to transforming information and experience into actionable business insights. It establishes insight frameworks that convert data, observations, and knowledge into strategic intelligence that drives better decision-making, innovation, and competitive advantage.
 */
export interface IGNDocument extends BaseBSpecDocument {
  type: "IGN";
  domain: "Growth & Innovation";

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

export default IGNDocument;
