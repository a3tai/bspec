/**
 * Messaging Framework Document Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 */

import { BaseBSpecDocument } from '../base';

/**
 * Messaging Framework (MSG)
 *
 * Domain: Brand & Marketing
 * Purpose: The Messaging Framework document defines what the brand says and how it says it to different audiences across various touchpoints. It establishes messaging architecture that ensures consistent communication, builds brand recognition, and drives desired customer actions through strategic message development and adaptation.
 */
export interface MSGDocument extends BaseBSpecDocument {
  type: "MSG";
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

export default MSGDocument;
