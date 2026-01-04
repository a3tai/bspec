---
title: "CST: Cost Structure"
description: "BSpec CST document type specification"
---

# CST: Cost Structure

**Domain**: Business Model

The Cost Structure defines systematic analysis and optimization of organizational costs to support business strategy and profitability. It establishes cost frameworks that enable efficient resource allocation, competitive positioning, and sustainable business operations.

## Metadata Schema

```yaml
---
id: CST-{identifier}
type: CST
status: Draft|Review|Accepted|Deprecated
version: 1.0.0
domain: business-model
priority: critical|high|medium|low

depends_on: []
enables: []
---
```

## Quality Levels

| Level | Requirements |
|-------|--------------|
| Bronze | Cost structure overview with purpose, scope, and strategic context, Cost categories and classification with primary categories and behavior analysis, Basic cost driver analysis with primary drivers and relationships |
| Silver | Comprehensive activity-based costing with activity analysis and resource consumption, Cost structure optimization with reduction opportunities and flexibility strategies, Unit economics with cost analysis, break-even analysis, and benchmarking |
| Gold | Full documentation with validation evidence |

## Relationships

**Depends on**: REV, PRI, CHN

**Enables**: KPT, KRS, KAC

---

- [Back to Document Types](/docs/document-types)
- [Domain: Business Model](/docs/domains/business-model)
