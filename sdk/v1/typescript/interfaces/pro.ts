/**
 * Process Specification Document Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 */

import { BaseBSpecDocument } from '../base';

/**
 * Process Specification (PRO)
 *
 * Domain: Operations & Execution
 * Purpose: The Process Specification defines systematic approaches to executing business activities through documented, repeatable, and optimized processes. It establishes process frameworks that ensure consistent execution, continuous improvement, and strategic alignment.
 */
export interface PRODocument extends BaseBSpecDocument {
  type: "PRO";
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

export default PRODocument;
