---
title: "KRS: Key Resources"
description: "BSpec KRS document type specification"
---

# KRS: Key Resources

**Domain**: Business Model

The Key Resources defines systematic identification, development, and management of critical organizational assets that create competitive advantage and enable business strategy execution. It establishes resource frameworks that optimize asset utilization, investment allocation, and capability building.

## Metadata Schema

```yaml
---
id: KRS-{identifier}
type: KRS
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
| Bronze | Resource overview with purpose, strategic importance, and value creation role, Resource strategy with context, philosophy, and investment approach, Basic resource inventory with major resource categories and assets |
| Silver | Comprehensive resource assessment with criticality analysis and utilization evaluation, Resource development strategy with capability building and acquisition approaches, Resource management with planning, allocation, and performance systems |
| Gold | Full documentation with validation evidence |

## Relationships

**Depends on**: STR, KAC, CAP

**Enables**: PER, QUA, SVC

---

- [Back to Document Types](/docs/document-types)
- [Domain: Business Model](/docs/domains/business-model)
