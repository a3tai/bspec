/**
 * Organizational Purpose Document Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 */

import { BaseBSpecDocument } from '../base';

/**
 * Organizational Purpose (PUR)
 *
 * Domain: Strategic Foundation
 * Purpose: The Purpose document articulates the organization's social impact and stakeholder value beyond profit. It defines the broader positive change the organization creates in the world.
 */
export interface PURDocument extends BaseBSpecDocument {
  type: "PUR";
  domain: "Strategic Foundation";

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

export default PURDocument;
