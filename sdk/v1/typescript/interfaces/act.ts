/**
 * Generated BSpec Document Type Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 * Generated at: 2025-09-28T15:14:00.157061
 * Generator: typescript-generator v1.0.0
 */

import { BaseBSpecDocument } from '../base';

/**
 * Key Activities (ACT)
 * Domain: business-model
 *
 * Essential processes for value creation
 */
export interface ACTDocument extends BaseBSpecDocument {
  /** Document type identifier */
  readonly type: 'ACT';

  /** Domain classification */
  readonly domain: 'business-model';




}