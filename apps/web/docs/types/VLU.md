---
title: "VLU: Valuation"
description: "BSpec VLU document type specification"
---

# VLU: Valuation

**Domain**: Financial & Investment

The Valuation document defines systematic approaches to determining business value, asset worth, and enterprise valuation through multiple methodologies and analytical frameworks. It establishes valuation standards that support investment decisions, transaction analysis, and strategic planning with objective, defensible value assessments.

## Metadata Schema

```yaml
---
id: VLU-{identifier}
type: VLU
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
| Bronze | Basic valuation using single approach, Simple comparable analysis with market multiples, Basic sensitivity analysis and documentation |
| Silver | Multiple valuation approaches with appropriate weighting, Comprehensive market analysis and comparable selection, Detailed sensitivity analysis and scenario modeling |
| Gold | Full documentation with validation evidence |

## Relationships

**Depends on**: FIN, FOR, MKT

**Enables**: FND, INV, STR

---

- [Back to Document Types](/docs/document-types)
- [Domain: Financial & Investment](/docs/domains/financial-investment)
