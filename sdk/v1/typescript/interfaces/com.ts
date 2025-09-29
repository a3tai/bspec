/**
 * Compliance Document Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 */

import { BaseBSpecDocument } from '../base';

/**
 * Compliance (COM)
 *
 * Domain: Risk & Governance
 * Purpose: The Compliance document defines systematic approaches to ensuring adherence to laws, regulations, policies, and standards that govern business operations. It establishes compliance frameworks that prevent violations, manage regulatory risks, and maintain organizational integrity and reputation.
 */
export interface COMDocument extends BaseBSpecDocument {
  type: "COM";
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

export default COMDocument;
