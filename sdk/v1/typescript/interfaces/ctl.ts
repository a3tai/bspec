/**
 * Controls Document Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 */

import { BaseBSpecDocument } from '../base';

/**
 * Controls (CTL)
 *
 * Domain: Risk & Governance
 * Purpose: The Controls document defines systematic approaches to designing, implementing, and operating internal controls that mitigate business risks, ensure compliance, and support reliable business operations. It establishes control frameworks that provide reasonable assurance for achieving business objectives while managing risk exposure.
 */
export interface CTLDocument extends BaseBSpecDocument {
  type: "CTL";
  domain: "Risk & Governance";

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

export default CTLDocument;
