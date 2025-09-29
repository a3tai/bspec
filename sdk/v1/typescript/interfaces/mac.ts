/**
 * Macro Environment Document Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 */

import { BaseBSpecDocument } from '../base';

/**
 * Macro Environment (MAC)
 *
 * Domain: Market & Environment
 * Purpose: The Macro Environment document analyzes broad economic, political, social, and technological factors that influence the business environment. It examines PESTEL factors and their implications.
 */
export interface MACDocument extends BaseBSpecDocument {
  type: "MAC";
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

export default MACDocument;
