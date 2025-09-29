/**
 * Data Models Document Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 */

import { BaseBSpecDocument } from '../base';

/**
 * Data Models (DAT)
 *
 * Domain: {Business domain or bounded context}
 * Purpose: The Data Models document defines systematic approaches to designing, governing, and managing data structures that support business capabilities through coherent data architecture, quality assurance, and strategic data management. It establishes data frameworks that ensure consistency, integrity, and value creation from organizational data assets.
 */
export interface DATDocument extends BaseBSpecDocument {
  type: "DAT";
  domain: "{Business domain or bounded context}";

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

export default DATDocument;
