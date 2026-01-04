---
title: "FOR: Forecasts"
description: "BSpec FOR document type specification"
---

# FOR: Forecasts

**Domain**: Financial & Investment

The Forecasts document defines forward-looking financial predictions and scenarios that anticipate future business performance through analytical modeling and trend analysis. It establishes forecasting frameworks that enable strategic planning, risk assessment, and informed decision making.

## Metadata Schema

```yaml
---
id: FOR-{identifier}
type: FOR
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
| Bronze | Basic forecasting models with simple trend analysis, Regular forecast updates and accuracy tracking, Simple scenario analysis and risk consideration |
| Silver | Comprehensive forecasting framework with multiple models, Advanced statistical methods and validation processes, Detailed scenario planning and stress testing |
| Gold | Full documentation with validation evidence |

## Relationships

**Depends on**: FIN, BUD, TRN

**Enables**: STR, OBJ, INV

---

- [Back to Document Types](/docs/document-types)
- [Domain: Financial & Investment](/docs/domains/financial-investment)
