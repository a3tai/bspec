---
title: "REV: Revenue Model"
description: "BSpec REV document type specification"
---

# REV: Revenue Model

**Domain**: Business Model

The Revenue Model defines how organizations create, capture, and optimize revenue streams. It establishes value exchange mechanisms that align customer value with business profitability while ensuring sustainable and scalable monetization strategies.

## Metadata Schema

```yaml
---
id: REV-{identifier}
type: REV
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
| Bronze | Revenue overview with purpose, strategic importance, and value exchange, Revenue stream definition with characteristics and value creation logic, Basic revenue mechanics with pricing and payment structure |
| Silver | Comprehensive customer value analysis with segments, value proposition, and CLV, Market analysis with opportunity assessment and competitive positioning, Financial projections with forecasting, unit economics, and dependencies |
| Gold | Full documentation with validation evidence |

## Relationships

**Depends on**: VSN, VAL, PER

**Enables**: CST, PRO, CAP

---

- [Back to Document Types](/docs/document-types)
- [Domain: Business Model](/docs/domains/business-model)
