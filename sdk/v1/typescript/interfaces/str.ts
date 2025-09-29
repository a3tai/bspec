/**
 * Business Strategy Document Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 */

import { BaseBSpecDocument } from '../base';

/**
 * Business Strategy (STR)
 *
 * Domain: Strategic Foundation
 * Purpose: The Strategy document defines how the organization will achieve its vision and compete in its chosen markets. It articulates the key choices about where to play, how to win, and what capabilities to build.
 */
export interface STRDocument extends BaseBSpecDocument {
  type: "STR";
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

export default STRDocument;
