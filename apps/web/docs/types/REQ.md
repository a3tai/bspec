---
title: "REQ: Requirements Specification"
description: "BSpec REQ document type specification"
---

# REQ: Requirements Specification

**Domain**: Product Service

The Requirements Specification defines detailed functional and non-functional requirements for systems, features, and capabilities. It provides comprehensive, testable specifications that guide development and serve as the basis for validation and acceptance testing.

## Metadata Schema

```yaml
---
id: REQ-{identifier}
type: REQ
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
| Bronze | Requirements overview with purpose, scope, and stakeholder identification, Core functional requirements with basic acceptance criteria, Key non-functional requirements for performance and security |
| Silver | Comprehensive stakeholder analysis with interests and constraints, Detailed functional requirements across all system areas, Complete non-functional requirements covering all quality attributes |
| Gold | Full documentation with validation evidence |

## Relationships

**Depends on**: PRD, FEA, USE

**Enables**: QUA, UXD, PER

---

- [Back to Document Types](/docs/document-types)
- [Domain: Product Service](/docs/domains/product-service)
