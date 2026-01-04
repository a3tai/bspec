---
title: "INT: Integration Specification"
description: "BSpec INT document type specification"
---

# INT: Integration Specification

**Domain**: Product Service

The Integration Specification defines detailed technical and business requirements for connecting systems, applications, and services. It ensures reliable, secure, and performant data exchange and functionality sharing between integrated components while maintaining system integrity and compliance.

## Metadata Schema

```yaml
---
id: INT-{identifier}
type: INT
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
| Bronze | Integration overview with purpose, scope, and business value, Business context with problem definition and stakeholder identification, Basic technical specification with API and data exchange format |
| Silver | System integration architecture with comprehensive system overview and integration patterns, Detailed technical specification with complete API, data, and security specifications, Integration implementation approach with development and deployment strategy |
| Gold | Full documentation with validation evidence |

## Relationships

**Depends on**: FEA, REQ, QUA

**Enables**: SUP, PER, UXD

---

- [Back to Document Types](/docs/document-types)
- [Domain: Product Service](/docs/domains/product-service)
