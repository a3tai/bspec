---
title: "COM: Compliance"
description: "BSpec COM document type specification"
---

# COM: Compliance

**Domain**: Risk & Governance

The Compliance document defines systematic approaches to ensuring adherence to laws, regulations, policies, and standards that govern business operations. It establishes compliance frameworks that prevent violations, manage regulatory risks, and maintain organizational integrity and reputation.

## Metadata Schema

```yaml
---
id: COM-{identifier}
type: COM
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
| Bronze | Basic compliance program with key regulatory requirements, Simple compliance monitoring and incident response, Basic compliance training and policy framework |
| Silver | Comprehensive compliance framework with systematic monitoring, Structured compliance risk management and mitigation, Effective compliance training and culture development |
| Gold | Full documentation with validation evidence |

## Relationships

**Depends on**: RSK, CTL, REG

**Enables**: AUD, GOV, REP

---

- [Back to Document Types](/docs/document-types)
- [Domain: Risk & Governance](/docs/domains/risk-governance)
