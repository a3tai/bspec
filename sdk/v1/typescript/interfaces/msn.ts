/**
 * Mission Statement Document Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 */

import { BaseBSpecDocument } from '../base';

/**
 * Mission Statement (MSN)
 *
 * Domain: Strategic Foundation
 * Purpose: The Mission document defines why the organization exists and its core purpose in the world. It articulates the fundamental reason the organization was created and what it aims to accomplish for its stakeholders.
 */
export interface MSNDocument extends BaseBSpecDocument {
  type: "MSN";
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

export default MSNDocument;
