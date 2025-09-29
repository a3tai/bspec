/**
 * Ecosystem Document Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 */

import { BaseBSpecDocument } from '../base';

/**
 * Ecosystem (ECO)
 *
 * Domain: Market & Environment
 * Purpose: The Ecosystem document maps the network of partners, suppliers, distributors, and other stakeholders that create value around the organization. It analyzes ecosystem dynamics and partnership strategies.
 */
export interface ECODocument extends BaseBSpecDocument {
  type: "ECO";
  domain: "Market & Environment";

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

export default ECODocument;
