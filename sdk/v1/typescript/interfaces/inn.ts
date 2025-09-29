/**
 * Innovation Strategy Document Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 */

import { BaseBSpecDocument } from '../base';

/**
 * Innovation Strategy (INN)
 *
 * Domain: Growth & Innovation
 * Purpose: The Innovation Strategy document defines systematic approaches to identifying, developing, and scaling innovation opportunities that create competitive advantage and drive sustainable growth. It establishes innovation frameworks that transform innovation from random creativity into systematic business capability.
 */
export interface INNDocument extends BaseBSpecDocument {
  type: "INN";
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

export default INNDocument;
