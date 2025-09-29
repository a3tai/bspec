/**
 * Generated BSpec Document Type Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 * Generated at: 2025-09-28T15:14:00.157751
 * Generator: typescript-generator v1.0.0
 */

import { BaseBSpecDocument } from '../base';

/**
 * Processes (PRC)
 * Domain: operations
 *
 * Repeatable workflows that drive business outcomes
 */
export interface PRCDocument extends BaseBSpecDocument {
  /** Document type identifier */
  readonly type: 'PRC';

  /** Domain classification */
  readonly domain: 'operations';




}