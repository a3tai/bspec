/**
 * Role Definition Document Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 */

import { BaseBSpecDocument } from '../base';

/**
 * Role Definition (ROL)
 *
 * Domain: Operations & Execution
 * Purpose: The Role Definition defines systematic approaches to designing and documenting organizational roles that clarify responsibilities, authorities, and requirements. It establishes role frameworks that enable effective hiring, performance management, and career development.
 */
export interface ROLDocument extends BaseBSpecDocument {
  type: "ROL";
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

export default ROLDocument;
