/**
 * Future Planning Document Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 */

import { BaseBSpecDocument } from '../base';

/**
 * Future Planning (FUT)
 *
 * Domain: Growth & Innovation
 * Purpose: The Future Planning document defines systematic approaches to understanding and preparing for future possibilities through scenario development and strategic planning. It establishes future planning frameworks that enable organizations to navigate uncertainty and build capabilities for multiple possible futures.
 */
export interface FUTDocument extends BaseBSpecDocument {
  type: "FUT";
  domain: "Growth & Innovation";

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

export default FUTDocument;
