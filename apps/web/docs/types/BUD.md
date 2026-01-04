---
title: "BUD: Budget"
description: "BSpec BUD document type specification"
---

# BUD: Budget

**Domain**: Financial & Investment

The Budget document defines systematic resource allocation and spending plans that translate strategic objectives into financial commitments. It establishes budgeting frameworks that ensure disciplined resource management, performance accountability, and financial control.

## Metadata Schema

```yaml
---
id: BUD-{identifier}
type: BUD
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
| Bronze | Basic operating budget with revenue and expense projections, Simple budget approval and monitoring process, Basic variance reporting and analysis |
| Silver | Comprehensive budget framework with detailed resource allocation, Structured budget process with clear controls and approval workflows, Regular variance analysis and performance monitoring |
| Gold | Full documentation with validation evidence |

## Relationships

**Depends on**: FIN, STR, OBJ

**Enables**: MET, REP, CTL

---

- [Back to Document Types](/docs/document-types)
- [Domain: Financial & Investment](/docs/domains/financial-investment)
