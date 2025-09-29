/**
 * Theory of Change Document Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 */

import { BaseBSpecDocument } from '../base';

/**
 * Theory of Change (THY)
 *
 * Domain: Strategic Foundation
 * Purpose: The Theory of Change document maps the logic model connecting the organization's activities to its intended outcomes. It explains how and why specific actions will lead to desired changes.
 */
export interface THYDocument extends BaseBSpecDocument {
  type: "THY";
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

export default THYDocument;
