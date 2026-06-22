---
title: "COS: COSO Enterprise Risk Management"
description: "BSpec COS document type specification"
---

# COS: COSO Enterprise Risk Management

::: tip Document Type
**Code**: COS<br>
**Name**: COSO Enterprise Risk Management<br>
**Domain**: Risk & Governance
:::

## Abstract

This specification defines a COSO-based risk governance document within the BSpec 1.0 Universal Business Specification Standard. It aligns enterprise risk management around governance, strategy, performance, review, and information systems.

## Framework and Attribution

This template references **COSO ERM** terminology and structure for enterprise-wide risk governance alignment.

## Purpose and Scope

Use this document when an organization needs a structured ERM framing across governance bodies, risk appetite, and control assurance.

## Document Metadata Schema

```yaml
---
id: COS-{erm-scope}
title: "COSO ERM Framework — {Organization or Business Unit}"
type: COS
status: Draft|Review|Accepted|Deprecated
attribution_required: true
source_frameworks:
  - "COSO - Enterprise Risk Management (ERM)"
version: 1.0.0
owner: Chief-Risk-Officer|Internal-Audit
stakeholders: [board, executives, risk-team, business-unit-leads]
domain: risk-governance
priority: Critical|High|Medium|Low
scope: erm-framework
horizon: strategic
visibility: internal

depends_on: [RSK-*,GOV-*,COM-*]
enables: [CTL-*,AUD-*,LEG-*]

erm_components: Governance|Strategy|Risk_Profile|Performance|Review|Information_Communication
risk_appetite_owner: Board|Risk-Committee|Executive-Committee

success_criteria:
  - "ERM structure is traceable from board intent to operational controls"
  - "Risk appetite and tolerance are consistently applied"
  - "Risk events and actions are monitored through a shared language"

assumptions:
  - "Risk governance has board and executive sponsorship"
  - "ERM metrics are measured consistently"
  - "Control owners are accountable for outcomes"

constraints: [Maturity limits, data integration, accountability gaps]
standards: [COSO ERM and internal governance requirements]
review_cycle: quarterly
---
```

## Content Structure Template

# COSO ERM Framework — &#123;Organization or Unit&#125;

## Executive Summary

**Scope:** &#123;Business unit or enterprise scope&#125;<br>
**Governance Owner:** &#123;Board / Risk Committee / CEO&#125;<br>
**Risk Appetite Status:** &#123;Defined / Draft / Needs Review&#125;

## ERM Governance

### Governance & Oversight
- **Board role**
- **Management role**
- **Committee cadence**
- **Escalation paths**

### Strategic Alignment
- **Strategy risks**
- **Risk appetite and tolerance statements**
- **Objective linkage**

## Core ERM Components

For each component, document:

- **Component:** &#123;Governance, Strategy, Risk Assessment, Control Activities, Information & Communication, Monitoring&#125;
- **Current State:** &#123;Current baseline&#125;
- **Gaps:** &#123;Known gaps and controls&#125;
- **Actions:** &#123;Mitigation or improvement steps&#125;

## Risk Profile Map

- **Top Enterprise Risks**
- **Root Causes**
- **Control Coverage**
- **Residual Exposure**

## Integration Links

- `RSK` impact: source risk identification and prioritization.
- `COM` impact: compliance obligations and legal requirements.
- `AUD` impact: assurance and testing scope.
- `CTL` impact: control design and operating effectiveness.

## Quality Standard

- ERM outputs are reviewed at governance cadence.
- Risks and responses are connected to owners and evidence.
- Framework language is consistent across units.


---

## Related Documents

- [Back to Document Types](/docs/document-types)
- [Domain: Risk & Governance](/docs/domains/risk-governance)
- [Full Specification](/spec/v1-0-0)
