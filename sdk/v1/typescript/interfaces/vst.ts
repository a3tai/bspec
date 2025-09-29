/**
 * Value Stream Document Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 */

import { BaseBSpecDocument } from '../base';

/**
 * Value Stream (VST)
 *
 * Domain: Business Model
 * Purpose: The Value Stream defines systematic analysis and optimization of end-to-end value creation processes that deliver customer value. It establishes value flow frameworks that eliminate waste, optimize performance, and align organizational activities with customer value realization.
 */
export interface VSTDocument extends BaseBSpecDocument {
  type: "VST";
  domain: "Business Model";

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

export default VSTDocument;
