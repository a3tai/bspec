/**
 * Incidents Document Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 */

import { BaseBSpecDocument } from '../base';

/**
 * Incidents (INC)
 *
 * Domain: Risk & Governance
 * Purpose: The Incidents document defines systematic approaches to identifying, responding to, and managing incidents that disrupt business operations, threaten stakeholder safety, or impact organizational objectives. It establishes incident management frameworks that enable rapid response, effective recovery, and organizational learning from disruptive events.
 */
export interface INCDocument extends BaseBSpecDocument {
  type: "INC";
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

export default INCDocument;
