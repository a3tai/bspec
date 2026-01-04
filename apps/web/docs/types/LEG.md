---
title: "LEG: Legal"
description: "BSpec LEG document type specification"
---

# LEG: Legal

**Domain**: Risk & Governance

The Legal document defines systematic approaches to managing legal affairs, protecting legal interests, and ensuring compliance with applicable laws and regulations. It establishes legal frameworks that mitigate legal risks, manage contracts and disputes, protect intellectual property, and support business operations within appropriate legal boundaries.

## Metadata Schema

```yaml
---
id: LEG-{identifier}
type: LEG
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
| Bronze | Basic legal framework with essential contract and compliance management, Simple litigation management and dispute resolution procedures, Basic intellectual property protection and regulatory compliance |
| Silver | Comprehensive legal framework with systematic risk management, Structured contract lifecycle management and dispute prevention, Effective IP portfolio management and regulatory compliance programs |
| Gold | Full documentation with validation evidence |

## Relationships

**Depends on**: COM, RSK, GOV

**Enables**: CTL, REP, AUD

---

- [Back to Document Types](/docs/document-types)
- [Domain: Risk & Governance](/docs/domains/risk-governance)
