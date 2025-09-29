/**
 * Key Activities Document Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 */

import { BaseBSpecDocument } from '../base';

/**
 * Key Activities (KAC)
 *
 * Domain: Business Model
 * Purpose: The Key Activities defines systematic approaches to critical business activities that create customer value, drive competitive advantage, and enable business model execution. It establishes activity frameworks that optimize operational excellence, resource utilization, and performance management.
 */
export interface KACDocument extends BaseBSpecDocument {
  type: "KAC";
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

export default KACDocument;
