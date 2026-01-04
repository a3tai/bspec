---
title: "USE: Use Cases"
description: "BSpec USE document type specification"
---

# USE: Use Cases

**Domain**: Customer Value

The Use Cases document describes specific scenarios where customers apply the solution. It details the context, actors, preconditions, and steps involved in successful solution usage.

## Metadata Schema

```yaml
---
id: USE-{identifier}
type: USE
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
| Bronze | Use case summary with basic details and scope, Primary actors and stakeholders identified, Main success scenario documented |
| Silver | Comprehensive use case scenario with preconditions and postconditions, Alternative and exception flows documented, Business rules and success metrics defined |
| Gold | Full documentation with validation evidence |

## Relationships

**Depends on**: JTB, PER

**Enables**: REQ, STO, PRD

---

- [Back to Document Types](/docs/document-types)
- [Domain: Customer Value](/docs/domains/customer-value)
