/**
 * User Experience Design Document Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 */

import { BaseBSpecDocument } from '../base';

/**
 * User Experience Design (UXD)
 *
 * Domain: Product & Service
 * Purpose: The User Experience Design document defines comprehensive user experience strategies, interaction patterns, and design specifications that ensure products and services deliver intuitive, accessible, and delightful user experiences aligned with user needs and business objectives.
 */
export interface UXDDocument extends BaseBSpecDocument {
  type: "UXD";
  domain: "Product & Service";

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

export default UXDDocument;
