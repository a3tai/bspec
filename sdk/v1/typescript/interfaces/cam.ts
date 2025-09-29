/**
 * Marketing Campaign Document Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 */

import { BaseBSpecDocument } from '../base';

/**
 * Marketing Campaign (CAM)
 *
 * Domain: Brand & Marketing
 * Purpose: The Marketing Campaign document defines integrated marketing initiatives that achieve specific business objectives through coordinated messaging, creative execution, and multi-channel activation. It establishes campaign frameworks that drive awareness, engagement, and conversion while maintaining brand consistency and maximizing return on marketing investment.
 */
export interface CAMDocument extends BaseBSpecDocument {
  type: "CAM";
  domain: "Brand & Marketing";

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

export default CAMDocument;
