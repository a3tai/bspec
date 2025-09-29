/**
 * Adaptation and Agility Document Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 */

import { BaseBSpecDocument } from '../base';

/**
 * Adaptation and Agility (ADT)
 *
 * Domain: Growth & Innovation
 * Purpose: The Adaptation and Agility document defines systematic approaches to building organizational capabilities that enable rapid response to changing conditions and emerging opportunities. It establishes agility frameworks that transform organizations into adaptive entities that thrive in uncertainty and complexity.
 */
export interface ADTDocument extends BaseBSpecDocument {
  type: "ADT";
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

export default ADTDocument;
