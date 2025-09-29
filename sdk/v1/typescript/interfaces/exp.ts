/**
 * Experimentation Document Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 */

import { BaseBSpecDocument } from '../base';

/**
 * Experimentation (EXP)
 *
 * Domain: Growth & Innovation
 * Purpose: The Experimentation document defines systematic approaches to testing assumptions and validating new ideas through structured experimentation. It establishes experimentation frameworks that transform hypothesis testing from ad-hoc activities into systematic business capabilities that drive evidence-based decision-making and rapid learning.
 */
export interface EXPDocument extends BaseBSpecDocument {
  type: "EXP";
  domain: "Growth & Innovation";

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

export default EXPDocument;
