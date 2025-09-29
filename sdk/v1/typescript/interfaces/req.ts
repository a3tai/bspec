/**
 * Requirements Specification Document Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 */

import { BaseBSpecDocument } from '../base';

/**
 * Requirements Specification (REQ)
 *
 * Domain: Product & Service
 * Purpose: The Requirements Specification defines detailed functional and non-functional requirements for systems, features, and capabilities. It provides comprehensive, testable specifications that guide development and serve as the basis for validation and acceptance testing.
 */
export interface REQDocument extends BaseBSpecDocument {
  type: "REQ";
  domain: "Product & Service";

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

export default REQDocument;
