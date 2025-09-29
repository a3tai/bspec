/**
 * Skills Framework Document Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 */

import { BaseBSpecDocument } from '../base';

/**
 * Skills Framework (SKI)
 *
 * Domain: Operations & Execution
 * Purpose: The Skills Framework defines systematic approaches to identifying, assessing, and developing organizational skills and competencies that enable strategic execution and competitive advantage. It establishes skill frameworks that optimize talent development, career planning, and organizational capability building.
 */
export interface SKIDocument extends BaseBSpecDocument {
  type: "SKI";
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

export default SKIDocument;
