/**
 * Capability Specification Document Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 */

import { BaseBSpecDocument } from '../base';

/**
 * Capability Specification (CAP)
 *
 * Domain: Operations & Execution
 * Purpose: The Capability Specification defines systematic approaches to building and managing organizational capabilities that create competitive advantage and enable business strategy execution. It establishes capability frameworks that optimize development, performance, and continuous improvement.
 */
export interface CAPDocument extends BaseBSpecDocument {
  type: "CAP";
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

export default CAPDocument;
