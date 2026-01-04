---
title: "STO: User Stories"
description: "BSpec STO document type specification"
---

# STO: User Stories

**Domain**: Customer Value

The User Stories document captures individual requirements from the user perspective using the standard "As a... I want...

## Metadata Schema

```yaml
---
id: STO-{identifier}
type: STO
status: Draft|Review|Accepted|Deprecated
version: 1.0.0
domain: customer-value
priority: critical|high|medium|low

depends_on: []
enables: []
---
```

## Quality Levels

| Level | Requirements |
|-------|--------------|
| Bronze | Story statement with As a/I want/So that format, Basic acceptance criteria with Given/When/Then scenarios, Story details including type, epic, and component |
| Silver | Comprehensive acceptance criteria with functional and non-functional aspects, Complete definition of done with development, quality, and documentation criteria, Detailed technical considerations with complexity and risk assessment |
| Gold | Full documentation with validation evidence |

## Relationships

**Depends on**: USE, PER

**Enables**: REQ, TSK

---

- [Back to Document Types](/docs/document-types)
- [Domain: Customer Value](/docs/domains/customer-value)
