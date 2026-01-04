---
title: "INS: Insurance"
description: "BSpec INS document type specification"
---

# INS: Insurance

**Domain**: Risk & Governance

The Insurance document defines systematic approaches to managing insurance programs that transfer financial risks and protect organizational assets, operations, and stakeholders. It establishes insurance frameworks that optimize risk transfer, minimize total cost of risk, and ensure adequate protection against potential losses.

## Metadata Schema

```yaml
---
id: INS-{identifier}
type: INS
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
| Bronze | Basic insurance program with essential coverage types, Simple claims management and vendor relationships, Basic compliance with regulatory and contractual requirements |
| Silver | Comprehensive insurance program with systematic risk transfer, Structured claims management with recovery optimization, Effective vendor management and performance monitoring |
| Gold | Full documentation with validation evidence |

## Relationships

**Depends on**: RSK, FIN, LEG

**Enables**: CTL, COM, REP

---

- [Back to Document Types](/docs/document-types)
- [Domain: Risk & Governance](/docs/domains/risk-governance)
