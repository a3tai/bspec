/**
 * Infrastructure Document Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 */

import { BaseBSpecDocument } from '../base';

/**
 * Infrastructure (INF)
 *
 * Domain: Technology & Data
 * Purpose: The Infrastructure document defines systematic approaches to designing, deploying, and managing technology infrastructure that supports business operations through reliable, secure, and scalable platforms. It establishes infrastructure frameworks that ensure availability, performance, and cost optimization.
 */
export interface INFDocument extends BaseBSpecDocument {
  type: "INF";
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

export default INFDocument;
