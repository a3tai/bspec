/**
 * Facilities Management Document Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 */

import { BaseBSpecDocument } from '../base';

/**
 * Facilities Management (FAC)
 *
 * Domain: Operations & Execution
 * Purpose: The Facilities Management document defines systematic approaches to planning, operating, and maintaining physical facilities that support business operations through efficient space utilization, operational excellence, and employee productivity. It establishes facility frameworks that optimize cost, safety, and performance.
 */
export interface FACDocument extends BaseBSpecDocument {
  type: "FAC";
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

export default FACDocument;
