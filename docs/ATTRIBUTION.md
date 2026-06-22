# Attribution and Legal Compliance Notes

This repository inherits external names, trademarks, standards, and frameworks through these documents. Use this file as the compliance index before adding new content.

For the repository-level policy, including code vs documentation licensing split, see [ATTRIBUTION.md](../ATTRIBUTION.md).

## Trademarked Framework Names/Marks

- **Net Promoter**, **Net Promoter System**, **Net Promoter Score**
  - Use exact marks: **Net Promoter®**, **NPS®**, **Net Promoter System®**, **Net Promoter Score℠**.
  - Commercial use of the marks requires prior written consent per the Bain licensing terms.
  - If included in a document, add a trademark notice near first mention and set `attribution_required: true` with `source_frameworks` entry.

- **ITIL**, **COBIT**, **CMMI**, **Stage-Gate**, **Three Horizons**, **Balanced Scorecard**, and similar named frameworks
  - Treat as trademarked product/model names where applicable.
  - Add naming and source attribution in the using document.

- **Bain**, **Porter**, **Wardley**, **COSO**, **Lean Startup**, **Design Thinking**, **Value Proposition Canvas**, **Jobs-to-be-Done**, **Business Model Canvas**
  - Treat as rights-managed named frameworks unless the cited source states otherwise.
  - Add explicit framework attribution at first mention and keep `attribution_required: true`.

## Copyrighted / License-Sensitive Frameworks

- **Empathy Map (Dave Gray / XPLANE / Gamestorming)**: typically CC BY-SA 4.0 in public reuse.
- **Business Model Canvas (Strategyzer / Osterwalder / Pigneur)**: CC BY-SA 3.0 applies; ShareAlike applies to derivative work.
- **Value Proposition Canvas (Strategyzer)**: not open by default for software or adaptation reuse; requires explicit permission where included.
- **Jobs-to-be-Done**: non-trademark concept in many contexts; separate outcome framework attributions should still be documented when used as a named method.

## BSpec Trademark Clearance Workflow

Before publishing or branding any BSpec distribution as a named standard or product, complete this checklist:

- Confirm trademark owner and registration status in:
  - USPTO TESS (United States)
  - EUIPO eSearch (European Union)
  - WIPO Global Brand Database (international collision checks)
- Confirm intended product/service classes (branding, SaaS, training, docs, certifications).
- Confirm search notes include:
  - Search terms used
  - Classes and jurisdictions reviewed
  - Clearance decision and legal counsel sign-off
- Record clearance status in the repository (e.g., `docs/ATTRIBUTION.md` and changelog notes).
- If conflicts are identified, either:
  - Rebrand references (e.g., `BSpec` to a non-branded alias), or
  - Obtain explicit permission before external commercialization/marketing.

## Regulatory / Legal Notes in Content

- Distinguish certifications/assessments (e.g., SOC 2 is an attestation, not a statutory regulation).
- Mark licensing/contract instruments as contractual (e.g., PCI DSS obligations).
- State jurisdiction clearly when referencing law or regulation.

## Operational Rule

- Any spec document using or adapting an external framework, standard, or trademarked measure must:
  1. set `attribution_required: true`
  2. list sources in `source_frameworks`
  3. include a short attribution note in the markdown body near first mention
