---
title: "SUP: Support Specification"
description: "BSpec SUP document type specification"
---

# SUP: Support Specification

**Domain**: Product Service

The Support Specification defines comprehensive customer support strategies, processes, and service standards that ensure exceptional customer experience throughout the product lifecycle. It establishes support frameworks that maximize customer success while optimizing operational efficiency and business outcomes.

## Metadata Schema

```yaml
---
id: SUP-{identifier}
type: SUP
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
| Bronze | Support overview with purpose, scope, and philosophy, Basic support strategy with objectives and principles, Core support service definition with tiers and channels |
| Silver | Comprehensive support process design with issue management and escalation, Support team structure with roles, responsibilities, and training, Service level agreements with response and resolution targets |
| Gold | Full documentation with validation evidence |

## Relationships

**Depends on**: SVC, PER, INT

**Enables**: QUA, UXD, REQ

---

- [Back to Document Types](/docs/document-types)
- [Domain: Product Service](/docs/domains/product-service)
