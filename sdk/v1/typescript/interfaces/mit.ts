/**
 * Generated BSpec Document Type Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v1.0.0 specification
 * Generated at: 2025-09-28T15:14:00.160974
 * Generator: typescript-generator v1.0.0
 */

import { BaseBSpecDocument } from '../base';

/**
 * Mitigations (MIT)
 * Domain: risk
 *
 * Risk response strategies and controls
 */
export interface MITDocument extends BaseBSpecDocument {
  /** Document type identifier */
  readonly type: 'MIT';

  /** Domain classification */
  readonly domain: 'risk';




  /** Risk management documents require mitigation planning */
  readonly risk_management: true;
}