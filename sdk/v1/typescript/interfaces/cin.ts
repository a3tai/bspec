/**
 * Interviews Document Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 */

import { BaseBSpecDocument } from '../base';

/**
 * Interviews (CIN)
 *
 * Domain: Customer & Value
 * Purpose: The Interviews document captures structured customer research conversations that provide deep insights into needs, behaviors, motivations, and experiences. It documents qualitative research findings that inform product and strategy decisions.
 */
export interface CINDocument extends BaseBSpecDocument {
  type: "CIN";
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

export default CINDocument;
