/**
 * Analytics Document Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 */

import { BaseBSpecDocument } from '../base';

/**
 * Analytics (ANA)
 *
 * Domain: Technology & Data
 * Purpose: The Analytics document defines systematic approaches to designing, implementing, and managing analytics capabilities that enable data-driven decision making through business intelligence, advanced analytics, and strategic insights. It establishes analytics frameworks that ensure business value, data quality, and user adoption.
 */
export interface ANADocument extends BaseBSpecDocument {
  type: "ANA";
  domain: "Technology & Data";

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

export default ANADocument;
