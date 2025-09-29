/**
 * Generated BSpec Document Type Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 * Generated at: 2025-09-28T15:14:00.156948
 * Generator: typescript-generator v1.0.0
 */

import { BaseBSpecDocument } from '../base';

/**
 * Key Resources (RES)
 * Domain: business-model
 *
 * Critical assets required for the business model
 */
export interface RESDocument extends BaseBSpecDocument {
  /** Document type identifier */
  readonly type: 'RES';

  /** Domain classification */
  readonly domain: 'business-model';




}