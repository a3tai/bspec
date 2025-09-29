/**
 * Learning Organization Document Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 */

import { BaseBSpecDocument } from '../base';

/**
 * Learning Organization (LEA)
 *
 * Domain: Growth & Innovation
 * Purpose: The Learning Organization document defines systematic approaches to building organizational capabilities for continuous learning and adaptation. It establishes learning frameworks that transform organizations into adaptive entities that learn faster than their environment changes, creating sustainable competitive advantage through superior learning capabilities.
 */
export interface LEADocument extends BaseBSpecDocument {
  type: "LEA";
  domain: "Growth & Innovation";

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

export default LEADocument;
