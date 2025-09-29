/**
 * Generated BSpec Document Type Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 * Generated at: 2025-09-28T15:14:00.161342
 * Generator: typescript-generator v1.0.0
 */

import { BaseBSpecDocument } from '../base';

/**
 * Crisis Management (CRI)
 * Domain: risk
 *
 * Crisis response and business continuity plans
 */
export interface CRIDocument extends BaseBSpecDocument {
  /** Document type identifier */
  readonly type: 'CRI';

  /** Domain classification */
  readonly domain: 'risk';




}