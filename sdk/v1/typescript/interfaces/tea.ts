/**
 * Team Structure Document Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 */

import { BaseBSpecDocument } from '../base';

/**
 * Team Structure (TEA)
 *
 * Domain: Operations & Execution
 * Purpose: The Team Structure defines systematic approaches to organizing and managing teams that deliver business outcomes through effective collaboration, clear accountability, and continuous improvement. It establishes team frameworks that optimize performance, engagement, and value delivery.
 */
export interface TEADocument extends BaseBSpecDocument {
  type: "TEA";
  domain: "Operations & Execution";

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

export default TEADocument;
