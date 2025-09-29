/**
 * Market Definition Document Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 */

import { BaseBSpecDocument } from '../base';

/**
 * Market Definition (MKT)
 *
 * Domain: Market & Environment
 * Purpose: The Market Definition document establishes the boundaries, size, and characteristics of the addressable market. It defines TAM (Total Addressable Market), SAM (Serviceable Addressable Market), and SOM (Serviceable Obtainable Market).
 */
export interface MKTDocument extends BaseBSpecDocument {
  type: "MKT";
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

export default MKTDocument;
