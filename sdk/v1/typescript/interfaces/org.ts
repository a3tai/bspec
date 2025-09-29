/**
 * Organization Structure Document Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 */

import { BaseBSpecDocument } from '../base';

/**
 * Organization Structure (ORG)
 *
 * Domain: Operations & Execution
 * Purpose: The Organization Structure defines systematic approaches to designing and managing organizational hierarchies, reporting relationships, and team structures that enable effective execution and coordination. It establishes organizational frameworks that optimize authority, accountability, and communication.
 */
export interface ORGDocument extends BaseBSpecDocument {
  type: "ORG";
  domain: "Operations & Execution";

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

export default ORGDocument;
