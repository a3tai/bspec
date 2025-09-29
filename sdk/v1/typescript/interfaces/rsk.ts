/**
 * Risks Document Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 */

import { BaseBSpecDocument } from '../base';

/**
 * Risks (RSK)
 *
 * Domain: Risk & Governance
 * Purpose: The Risks document defines systematic approaches to identifying, assessing, and managing business risks that could impact organizational objectives, operations, and stakeholder value. It establishes risk management frameworks that enable proactive risk mitigation, informed decision making, and resilient business operations.
 */
export interface RSKDocument extends BaseBSpecDocument {
  type: "RSK";
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

export default RSKDocument;
