/**
 * Generated BSpec Document Type Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 * Generated at: 2025-09-28T15:14:00.161777
 * Generator: typescript-generator v1.0.0
 */

import { BaseBSpecDocument } from '../base';

/**
 * Growth Model (GRW)
 * Domain: growth
 *
 * How the business scales customers and revenue
 */
export interface GRWDocument extends BaseBSpecDocument {
  /** Document type identifier */
  readonly type: 'GRW';

  /** Domain classification */
  readonly domain: 'growth';




}