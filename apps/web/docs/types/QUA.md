---
title: "QUA: Quality Specification"
description: "BSpec QUA document type specification"
---

# QUA: Quality Specification

**Domain**: Product Service

The Quality Specification defines comprehensive quality standards, metrics, and assurance processes for products, services, and systems. It establishes quality frameworks that ensure consistent delivery of value while meeting stakeholder expectations and regulatory requirements.

## Metadata Schema

```yaml
---
id: QUA-{identifier}
type: QUA
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
| Bronze | Quality overview with purpose, scope, and philosophy, Basic quality framework with model and objectives, Core quality standards for internal and external quality |
| Silver | Comprehensive quality metrics across product, process, and customer dimensions, Quality assurance processes with planning, control, and improvement, Quality testing strategy with approach, planning, and execution |
| Gold | Full documentation with validation evidence |

## Relationships

**Depends on**: REQ, FEA, ROD

**Enables**: PER, UXD, SUP

---

- [Back to Document Types](/docs/document-types)
- [Domain: Product Service](/docs/domains/product-service)
