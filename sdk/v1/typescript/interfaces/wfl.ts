/**
 * Workflow Specification Document Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 */

import { BaseBSpecDocument } from '../base';

/**
 * Workflow Specification (WFL)
 *
 * Domain: Operations & Execution
 * Purpose: The Workflow Specification defines systematic approaches to designing, implementing, and managing business workflows that automate and optimize operational processes. It establishes workflow frameworks that ensure efficient execution, robust exception handling, and continuous improvement.
 */
export interface WFLDocument extends BaseBSpecDocument {
  type: "WFL";
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

export default WFLDocument;
