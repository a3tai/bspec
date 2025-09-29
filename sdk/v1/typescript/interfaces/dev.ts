/**
 * Development Document Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 */

import { BaseBSpecDocument } from '../base';

/**
 * Development (DEV)
 *
 * Domain: Technology & Data
 * Purpose: The Development document defines systematic approaches to software development practices, methodologies, and team processes that enable efficient delivery of high-quality software solutions. It establishes development frameworks that ensure consistency, quality, and collaborative effectiveness in software creation.
 */
export interface DEVDocument extends BaseBSpecDocument {
  type: "DEV";
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

export default DEVDocument;
