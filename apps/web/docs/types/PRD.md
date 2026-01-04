---
title: "PRD: Product Requirements Document"
description: "BSpec PRD document type specification"
---

# PRD: Product Requirements Document

**Domain**: Product Service

The Product Requirements Document defines comprehensive requirements for products, bridging customer needs with technical implementation. It serves as the authoritative specification for product development teams and stakeholders.

## Metadata Schema

```yaml
---
id: PRD-{identifier}
type: PRD
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
| Bronze | Basic problem statement and solution overview, High-level feature list with priorities, Basic success metrics identified |
| Silver | Detailed market research and competitive analysis, Comprehensive functional and non-functional requirements, Risk assessment with mitigation strategies |
| Gold | Full documentation with validation evidence |

## Relationships

**Depends on**: PER, JTB, GAI

**Enables**: FEA, REQ, UXD

---

- [Back to Document Types](/docs/document-types)
- [Domain: Product Service](/docs/domains/product-service)
