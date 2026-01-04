---
title: "PSP: Performance Specification"
description: "BSpec PSP document type specification"
---

# PSP: Performance Specification

**Domain**: Product Service

The Performance Specification defines comprehensive performance requirements, targets, and measurement frameworks for systems, applications, and services. It ensures optimal system performance that meets user expectations, business objectives, and technical scalability requirements.

## Metadata Schema

```yaml
---
id: PSP-{identifier}
type: PSP
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
| Bronze | Performance overview with purpose, scope, and business impact, Basic performance objectives and requirements, Core performance metrics and measurement approach |
| Silver | Comprehensive performance requirements across response time, throughput, and scalability, Performance measurement framework with metrics, tools, and baselines, Performance testing strategy with planning and execution approach |
| Gold | Full documentation with validation evidence |

## Relationships

**Depends on**: REQ, FEA, QUA

**Enables**: SUP, INT, UXD

---

- [Back to Document Types](/docs/document-types)
- [Domain: Product Service](/docs/domains/product-service)
