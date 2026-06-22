# BSpec Audit To-Do

_Last updated: 2026-05-28_

This file tracks remediation items from [`docs/audit.md`](./audit.md) as actionable work items.

## P0 — Must-Fix (Blocking)

- [x] Fix YAML type-code collisions (`BPO`, `MCH`, `CIN`, `IFL`, `PSP`).
- [x] Remove phantom relationship references in `depends_on`/`enables` (e.g., `BMC`, `GTM`, `SUC`, ...).
- [x] Add `LICENSE`, `CONTRIBUTING`, and `ATTRIBUTION` top-level governance files (root `ATTRIBUTION.md` added; contributor and license terms remain present).
- [x] Remove un-attributed NPS/Net Promoter references or add proper trademark attribution with `source_frameworks`.
- [x] Handle VPC-derived `PAI` / `GAI` / `JTB` attribution to Strategyzer Value Proposition Canvas.
- [x] Reconcile and regenerate web docs/type index pages from spec source so document inventory reflects newly added types.

## P1 — Must Fix Before Broad Distribution

- [x] Mark all Strategyzer Business Model Canvas-derived document types as attribution+ShareAlike and align `source_frameworks` consistently.
- [ ] Complete `BSpec` trademark search/clearance workflow for the standard name.
- [x] Add explicit trademark notices for additional named frameworks where they are reused (Stage-Gate, etc.).
- [x] Add attribution framework for Empathy Map (`EMP`) and add a source-framework declaration.
- [x] Add attribution for Moore positioning template in `POS`.
- [x] Add attribution for Tuckman stages in `TEA`.
- [x] Add attribution for Balanced Scorecard structure in `MET`.

## P2 — Legal/Accuracy and Quality

- [x] Resolve remaining legal/regulatory clarity items:
  - [x] SOX certification/control wording in `AUD` and filing scope.
  - [x] Clarify SOX scope and applicability.
  - [x] Clarify HIPAA scope as healthcare-only.
  - [x] Use "Right to Erasure" wording in privacy sections.
  - [x] Clarify FIN 48 versus ASC 740-10 in `TAX`.
  - [x] Distinguish PCI-DSS as contractual/commercial obligation.
  - [x] Clarify SOC 2 as attestation (not a certification).
  - [x] Remove/clarify mischaracterized legal requirements in `COM`/`LEG` overlap.
- [x] Add explicit `attribution_required` + `source_frameworks` for remaining borrowed named methods from the audit list.
- [x] Carry trademark notices for ITIL/COBIT/CMMI where names are used. (`REQ` and `QUA` now include explicit trademark notices and `attribution_required` entries.)

## P3 — Naming / Structural Cleanup

- [x] Add explicit `VAL` vs `VPR` separation (Value Propositions not Organizational Values).
- [x] Decide and standardize `STO` naming semantics, or rename. (`STO` kept as namespace-compatible alias of User Stories; explicit naming note added.)
- [x] Address `JTB` vs JTBD nomenclature tradeoff. (`JTB` retained as namespace code with explicit JTBD naming note.)
- [x] Standardize `THR` vs `THY` boundary and usage (`THR` explicitly excludes THY intent and states it maps to theory of change).
- [x] Harmonize `MOT` and `THY` framing if rename is desired.
- [x] Resolve `LEA` meaning ambiguity between Learning vs Leadership (explicitly fixed to Learning Organization).

## Overlap/Redundancy Workstream

- [x] Split/align `POS` vs `BPO` (positioning roles) boundaries.
- [x] Split/align `CHN` vs `MCH` roles (distribution vs media/marketing mix).
- [x] Clarify `CUS` / `CJM` / `FEE` / `SUP` boundaries and add missing Customer Success document.
- [x] Deconflict `PRD` / `FEA` / `REQ` / `USE` / `STO` overlaps and scope language.
- [x] Deconflict `ARC` / `SYS` / `INF` / `DEV` overlaps and domain boundaries.
- [x] Deconflict `KNO` / `LRN` / `LEA` / `WIS` into a two-document learning structure or equivalent (bounded in place; explicit ownership boundaries are now documented).
- [x] Clarify `DEC` vs `HYP` decision status semantics (state transitions and scope boundaries now explicit).
- [x] Clarify `INN` / `RND` / `EXP` / `IGN` innovation stack responsibilities.
- [x] Clarify `POL` / `CTL` / `COM` / `GOV` governance stack boundaries.
- [x] Clarify `STR` / `OBJ` / `MET` objective/measurement interface.
- [x] Clarify `VLU` / `FIN` / `FOR` overlap in valuation and forecasting models.

## Gaps to Fill

- [x] Add a first-class Customer Value Proposition document type (placeholder: `VPR`) and re-point value-proposition references.
- [x] Add Sales Strategy / Sales Operations document (`SAL`) and align references.
- [x] Add People/HR Strategy document (`PPL` or equivalent) for talent, hiring, compensation, performance.
- [x] Add explicit Data Privacy document beyond shared `SEC`/`DAT` content.
- [x] Add M&A / Corporate Development type and dependencies.
- [x] Add pricing lifecycle doc for discounting, deal desk, contracting policy.
- [x] Add Product Lifecycle / Sunset / EOL planning type.
- [x] Add Crisis Management / Business Continuity (`CRI` / `BCR`) docs or replace references.
- [x] Add Stakeholder Map (`STA`).
- [x] Add SWOT synthesis type and naming.
- [x] Add missing framework templates where consistently used (SWOT/Porter/Wardley/COSO).
