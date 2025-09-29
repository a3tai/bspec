/**
 * APIs Document Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 */

import { BaseBSpecDocument } from '../base';

/**
 * APIs (API)
 *
 * Domain: Technology & Data
 * Purpose: The APIs document defines systematic approaches to designing, implementing, and managing application programming interfaces that enable business capabilities through effective integration, developer experience, and scalable service delivery. It establishes API frameworks that ensure security, performance, and strategic alignment.
 */
export interface APIDocument extends BaseBSpecDocument {
  type: "API";
  domain: "Technology & Data";

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

export default APIDocument;
