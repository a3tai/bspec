---
title: "INV: Investment"
description: "BSpec INV document type specification"
---

# INV: Investment

**Domain**: Financial & Investment

The Investment document defines systematic approaches to capital allocation and investment decisions that optimize return on investment while managing risk and supporting strategic business objectives. It establishes investment frameworks that ensure disciplined capital deployment, rigorous evaluation, and performance accountability.

## Metadata Schema

```yaml
---
id: INV-{identifier}
type: INV
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
| Bronze | Basic investment analysis with financial metrics, Simple risk assessment and approval process, Basic performance monitoring and reporting |
| Silver | Comprehensive investment framework with portfolio management, Detailed financial and strategic analysis with scenario modeling, Structured due diligence and alternative evaluation |
| Gold | Full documentation with validation evidence |

## Relationships

**Depends on**: FIN, STR, VAL

**Enables**: MET, REP, RSK

---

- [Back to Document Types](/docs/document-types)
- [Domain: Financial & Investment](/docs/domains/financial-investment)
