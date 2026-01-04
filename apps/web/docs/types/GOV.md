---
title: "GOV: Governance"
description: "BSpec GOV document type specification"
---

# GOV: Governance

**Domain**: Risk & Governance

The Governance document defines systematic approaches to corporate governance that ensure effective oversight, accountability, and decision-making throughout the organization. It establishes governance frameworks that protect stakeholder interests, promote ethical behavior, and drive sustainable business performance.

## Metadata Schema

```yaml
---
id: GOV-{identifier}
type: GOV
status: Draft|Review|Accepted|Deprecated
version: 1.0.0
domain: risk-governance
priority: critical|high|medium|low

depends_on: []
enables: []
---
```

## Quality Levels

| Level | Requirements |
|-------|--------------|
| Bronze | Basic governance structure with board and management oversight, Simple governance policies and procedures, Basic stakeholder communication and reporting |
| Silver | Comprehensive governance framework with effective board oversight, Structured governance policies with compliance monitoring, Active stakeholder engagement and transparent communication |
| Gold | Full documentation with validation evidence |

## Relationships

**Depends on**: STR, RSK, COM

**Enables**: REP, AUD, PER

---

- [Back to Document Types](/docs/document-types)
- [Domain: Risk & Governance](/docs/domains/risk-governance)
