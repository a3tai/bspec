/**
 * Strategic Objectives Document Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 */

import { BaseBSpecDocument } from '../base';

/**
 * Strategic Objectives (OBJ)
 *
 * Domain: Strategic Foundation
 * Purpose: The Objectives document sets specific, measurable goals with timeframes that advance the organization toward its vision. These are typically structured as OKRs (Objectives and Key Results) or similar goal-setting frameworks.
 */
export interface OBJDocument extends BaseBSpecDocument {
  type: "OBJ";
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

export default OBJDocument;
