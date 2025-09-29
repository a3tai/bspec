/**
 * Positioning Document Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 */

import { BaseBSpecDocument } from '../base';

/**
 * Positioning (POS)
 *
 * Domain: Market & Environment
 * Purpose: The Positioning document defines how the organization wants to be perceived in the market relative to competitors and alternatives. It articulates the unique value proposition and market position.
 */
export interface POSDocument extends BaseBSpecDocument {
  type: "POS";
  domain: "Market & Environment";

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

export default POSDocument;
