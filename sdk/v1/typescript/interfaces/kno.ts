/**
 * Knowledge Management Document Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 */

import { BaseBSpecDocument } from '../base';

/**
 * Knowledge Management (KNO)
 *
 * Domain: Learning & Decisions
 * Purpose: The Knowledge Management document defines systematic approaches to capturing, organizing, sharing, and leveraging organizational knowledge assets. It establishes knowledge frameworks that preserve institutional knowledge, accelerate learning, enable innovation, and create competitive advantage through effective knowledge creation, retention, and application.
 */
export interface KNODocument extends BaseBSpecDocument {
  type: "KNO";
  domain: "Learning & Decisions";

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

export default KNODocument;
