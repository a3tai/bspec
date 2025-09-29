/**
 * Vision Statement Document Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 */

import { BaseBSpecDocument } from '../base';

/**
 * Vision Statement (VSN)
 *
 * Domain: Strategic Foundation
 * Purpose: The Vision document articulates the future state the organization aims to create—both for itself and the world. It describes what success looks like at a meaningful time horizon (typically 3-10 years).
 */
export interface VSNDocument extends BaseBSpecDocument {
  type: "VSN";
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

export default VSNDocument;
