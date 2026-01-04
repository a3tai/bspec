---
title: "UXD: User Experience Design"
description: "BSpec UXD document type specification"
---

# UXD: User Experience Design

**Domain**: Product Service

The User Experience Design document defines comprehensive user experience strategies, interaction patterns, and design specifications that ensure products and services deliver intuitive, accessible, and delightful user experiences aligned with user needs and business objectives.

## Metadata Schema

```yaml
---
id: UXD-{identifier}
type: UXD
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
| Bronze | Design overview with purpose, scope, and vision, User research foundation with understanding and insights, Basic experience strategy with principles and goals |
| Silver | Comprehensive user journey design with current and future state mapping, Information architecture with content strategy and navigation design, Interaction design with principles, patterns, and microinteractions |
| Gold | Full documentation with validation evidence |

## Relationships

**Depends on**: PER, CJM, USE

**Enables**: FEA, REQ, QUA

---

- [Back to Document Types](/docs/document-types)
- [Domain: Product Service](/docs/domains/product-service)
