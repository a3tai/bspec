/**
 * Architecture Document Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 */

import { BaseBSpecDocument } from '../base';

/**
 * Architecture (ARC)
 *
 * Domain: {Technical domain or system boundary}
 * Purpose: The Architecture document defines systematic approaches to designing and documenting technical architecture that enables business capabilities through coherent technology decisions, quality attribute optimization, and strategic alignment. It establishes architectural frameworks that guide technology evolution and ensure scalable, secure, and maintainable systems.
 */
export interface ARCDocument extends BaseBSpecDocument {
  type: "ARC";
  domain: "{Technical domain or system boundary}";

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

export default ARCDocument;
