/**
 * Tone of Voice Document Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 */

import { BaseBSpecDocument } from '../base';

/**
 * Tone of Voice (TON)
 *
 * Domain: Brand & Marketing
 * Purpose: The Tone of Voice document defines how the brand sounds in all communications, establishing voice characteristics that express brand personality and create consistent customer experiences. It provides guidelines that ensure authentic, recognizable, and engaging brand communication across all touchpoints and channels.
 */
export interface TONDocument extends BaseBSpecDocument {
  type: "TON";
  domain: "Brand & Marketing";

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

export default TONDocument;
