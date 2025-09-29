/**
 * Investment Document Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 */

import { BaseBSpecDocument } from '../base';

/**
 * Investment (INV)
 *
 * Domain: Financial & Investment
 * Purpose: The Investment document defines systematic approaches to capital allocation and investment decisions that optimize return on investment while managing risk and supporting strategic business objectives. It establishes investment frameworks that ensure disciplined capital deployment, rigorous evaluation, and performance accountability.
 */
export interface INVDocument extends BaseBSpecDocument {
  type: "INV";
  domain: "Financial & Investment";

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

export default INVDocument;
