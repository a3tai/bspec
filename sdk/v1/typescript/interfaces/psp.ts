/**
 * Performance Specification Document Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 */

import { BaseBSpecDocument } from '../base';

/**
 * Performance Specification (PSP)
 *
 * Domain: Product & Service
 * Purpose: The Performance Specification defines comprehensive performance requirements, targets, and measurement frameworks for systems, applications, and services. It ensures optimal system performance that meets user expectations, business objectives, and technical scalability requirements.
 */
export interface PSPDocument extends BaseBSpecDocument {
  type: "PSP";
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

export default PSPDocument;
