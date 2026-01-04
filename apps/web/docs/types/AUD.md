---
title: "AUD: Audit"
description: "BSpec AUD document type specification"
---

# AUD: Audit

**Domain**: Financial & Investment

The Audit document defines systematic approaches to internal and external audit processes that provide independent assurance on financial reporting, internal controls, and operational effectiveness. It establishes audit frameworks that ensure compliance, risk management, and continuous improvement in business operations.

## Metadata Schema

```yaml
---
id: AUD-{identifier}
type: AUD
status: Draft|Review|Accepted|Deprecated
version: 1.0.0
domain: financial-investment
priority: critical|high|medium|low

depends_on: []
enables: []
---
```

## Quality Levels

| Level | Requirements |
|-------|--------------|
| Bronze | Basic internal audit function with annual planning, Simple audit procedures and documentation, Basic external audit coordination and support |
| Silver | Comprehensive risk-based audit program, Structured audit methodology with quality controls, Effective SOX 404 compliance and reporting |
| Gold | Full documentation with validation evidence |

## Relationships

**Depends on**: REP, CTL, RSK

**Enables**: GOV, QUA, PER

---

- [Back to Document Types](/docs/document-types)
- [Domain: Financial & Investment](/docs/domains/financial-investment)
