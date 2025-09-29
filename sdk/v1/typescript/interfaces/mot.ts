/**
 * Competitive Moats Document Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 */

import { BaseBSpecDocument } from '../base';

/**
 * Competitive Moats (MOT)
 *
 * Domain: Strategic Foundation
 * Purpose: The Moats document identifies and analyzes the competitive advantages that protect the organization's market position. Moats are sustainable advantages that make it difficult for competitors to replicate success.
 */
export interface MOTDocument extends BaseBSpecDocument {
  type: "MOT";
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

export default MOTDocument;
