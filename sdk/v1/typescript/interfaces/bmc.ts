/**
 * Generated BSpec Document Type Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 * Generated at: 2025-09-28T15:14:00.155917
 * Generator: typescript-generator v1.0.0
 */

import { BaseBSpecDocument } from '../base';

/**
 * Business Model Canvas (BMC)
 * Domain: business-model
 *
 * Complete business model overview
 */
export interface BMCDocument extends BaseBSpecDocument {
  /** Document type identifier */
  readonly type: 'BMC';

  /** Domain classification */
  readonly domain: 'business-model';




}