/**
 * Service Specification Document Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 */

import { BaseBSpecDocument } from '../base';

/**
 * Service Specification (SVC)
 *
 * Domain: Product & Service
 * Purpose: The Service Specification defines comprehensive requirements for services, including delivery models, quality standards, and operational requirements. It ensures services are designed to deliver customer value effectively and efficiently.
 */
export interface SVCDocument extends BaseBSpecDocument {
  type: "SVC";
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

export default SVCDocument;
