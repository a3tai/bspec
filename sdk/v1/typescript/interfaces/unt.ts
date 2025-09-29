/**
 * Generated BSpec Document Type Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 * Generated at: 2025-09-28T15:14:00.157287
 * Generator: typescript-generator v1.0.0
 */

import { BaseBSpecDocument } from '../base';

/**
 * Unit Economics (UNT)
 * Domain: business-model
 *
 * Per-customer or per-unit financial metrics
 */
export interface UNTDocument extends BaseBSpecDocument {
  /** Document type identifier */
  readonly type: 'UNT';

  /** Domain classification */
  readonly domain: 'business-model';




}