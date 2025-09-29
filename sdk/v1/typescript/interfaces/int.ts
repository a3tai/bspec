/**
 * Integration Specification Document Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 */

import { BaseBSpecDocument } from '../base';

/**
 * Integration Specification (INT)
 *
 * Domain: Product & Service
 * Purpose: The Integration Specification defines detailed technical and business requirements for connecting systems, applications, and services. It ensures reliable, secure, and performant data exchange and functionality sharing between integrated components while maintaining system integrity and compliance.
 */
export interface INTDocument extends BaseBSpecDocument {
  type: "INT";
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

export default INTDocument;
