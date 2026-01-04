---
title: "TEA: Team Structure"
description: "BSpec TEA document type specification"
---

# TEA: Team Structure

**Domain**: Operations & Execution

The Team Structure defines systematic approaches to organizing and managing teams that deliver business outcomes through effective collaboration, clear accountability, and continuous improvement. It establishes team frameworks that optimize performance, engagement, and value delivery.

## Metadata Schema

```yaml
---
id: TEA-{identifier}
type: TEA
status: Draft|Review|Accepted|Deprecated
version: 1.0.0
domain: operations-execution
priority: critical|high|medium|low

depends_on: []
enables: []
---
```

## Quality Levels

| Level | Requirements |
|-------|--------------|
| Bronze | Clear team purpose and basic composition, Defined roles and responsibilities, Basic communication and meeting structure |
| Silver | Comprehensive team charter and operating model, Detailed skill matrix and development planning, Performance framework with metrics and objectives |
| Gold | Full documentation with validation evidence |

## Relationships

**Depends on**: ORG, ROL, SKI

**Enables**: PER, QUA, COL

---

- [Back to Document Types](/docs/document-types)
- [Domain: Operations & Execution](/docs/domains/operations-execution)
