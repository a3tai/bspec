/**
 * Governance Document Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 */

import { BaseBSpecDocument } from '../base';

/**
 * Governance (GOV)
 *
 * Domain: Risk & Governance
 * Purpose: The Governance document defines systematic approaches to corporate governance that ensure effective oversight, accountability, and decision-making throughout the organization. It establishes governance frameworks that protect stakeholder interests, promote ethical behavior, and drive sustainable business performance.
 */
export interface GOVDocument extends BaseBSpecDocument {
  type: "GOV";
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

export default GOVDocument;
