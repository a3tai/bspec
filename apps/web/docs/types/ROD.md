---
title: "ROD: Roadmap"
description: "BSpec ROD document type specification"
---

# ROD: Roadmap

**Domain**: Product Service

The Roadmap document defines strategic product and technology direction over multiple time horizons, aligning development efforts with business objectives while managing resource constraints and market dynamics. It provides stakeholders with visibility into planned evolution and investment priorities.

## Metadata Schema

```yaml
---
id: ROD-{identifier}
type: ROD
status: Draft|Review|Accepted|Deprecated
version: 1.0.0
domain: product-service
priority: critical|high|medium|low

depends_on: []
enables: []
---
```

## Quality Levels

| Level | Requirements |
|-------|--------------|
| Bronze | Roadmap overview with purpose, scope, and time horizon, Strategic context with objectives and market conditions, Basic planning horizons with near-term focus areas |
| Silver | Comprehensive roadmap structure with detailed planning horizons and initiative prioritization, Feature and capability evolution with current state assessment and development plan, Technology evolution strategy with investments and risk management |
| Gold | Full documentation with validation evidence |

## Relationships

**Depends on**: STR, OBJ, PRD

**Enables**: FEA, REQ, QUA

---

- [Back to Document Types](/docs/document-types)
- [Domain: Product Service](/docs/domains/product-service)
