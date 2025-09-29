/**
 * Generated BSpec Document Type Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 * Generated at: 2025-09-28T15:14:00.157470
 * Generator: typescript-generator v1.0.0
 */

import { BaseBSpecDocument } from '../base';

/**
 * Lifetime Value (LTV)
 * Domain: business-model
 *
 * Customer lifetime value models and drivers
 */
export interface LTVDocument extends BaseBSpecDocument {
  /** Document type identifier */
  readonly type: 'LTV';

  /** Domain classification */
  readonly domain: 'business-model';




}