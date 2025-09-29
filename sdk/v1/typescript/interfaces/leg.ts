/**
 * Legal Document Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 */

import { BaseBSpecDocument } from '../base';

/**
 * Legal (LEG)
 *
 * Domain: Risk & Governance
 * Purpose: The Legal document defines systematic approaches to managing legal affairs, protecting legal interests, and ensuring compliance with applicable laws and regulations. It establishes legal frameworks that mitigate legal risks, manage contracts and disputes, protect intellectual property, and support business operations within appropriate legal boundaries.
 */
export interface LEGDocument extends BaseBSpecDocument {
  type: "LEG";
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

export default LEGDocument;
