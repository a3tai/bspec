/**
 * Threats Document Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 */

import { BaseBSpecDocument } from '../base';

/**
 * Threats (THR)
 *
 * Domain: Market & Environment
 * Purpose: The Threats document identifies external risks to market position and business model. It analyzes potential threats from competitors, market changes, and environmental factors.
 */
export interface THRDocument extends BaseBSpecDocument {
  type: "THR";
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

export default THRDocument;
