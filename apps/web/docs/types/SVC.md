---
title: "SVC: Service Specification"
description: "BSpec SVC document type specification"
---

# SVC: Service Specification

**Domain**: Product Service

The Service Specification defines comprehensive requirements for services, including delivery models, quality standards, and operational requirements. It ensures services are designed to deliver customer value effectively and efficiently.

## Metadata Schema

```yaml
---
id: SVC-{identifier}
type: SVC
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
| Bronze | Service overview with purpose, scope, and target customers, Basic service value proposition and description, High-level delivery model and operational requirements |
| Silver | Comprehensive service delivery model with channels and processes, Service quality framework with standards and performance metrics, Risk management with mitigation strategies |
| Gold | Full documentation with validation evidence |

## Relationships

**Depends on**: CJM, CAP, PRO

**Enables**: SLA, PER, SUP

---

- [Back to Document Types](/docs/document-types)
- [Domain: Product Service](/docs/domains/product-service)
