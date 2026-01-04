---
title: "FIN: Financial Model"
description: "BSpec FIN document type specification"
---

# FIN: Financial Model

**Domain**: Financial & Investment

The Financial Model document defines comprehensive financial projections and planning models that forecast business performance through detailed P&L, balance sheet, and cash flow analysis. It establishes financial frameworks that enable strategic planning, investment decisions, and performance management.

## Metadata Schema

```yaml
---
id: FIN-{identifier}
type: FIN
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
| Bronze | Basic three-statement financial model, Simple revenue and cost projections, Basic scenario analysis capability |
| Silver | Comprehensive integrated financial model, Detailed revenue and cost modeling with drivers, Multiple scenario analysis and sensitivity testing |
| Gold | Full documentation with validation evidence |

## Relationships

**Depends on**: REV, BUD, FOR

**Enables**: FND, INV, MET

---

- [Back to Document Types](/docs/document-types)
- [Domain: Financial & Investment](/docs/domains/financial-investment)
