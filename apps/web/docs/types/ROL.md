---
title: "ROL: Role Definition"
description: "BSpec ROL document type specification"
---

# ROL: Role Definition

**Domain**: Operations & Execution

The Role Definition defines systematic approaches to designing and documenting organizational roles that clarify responsibilities, authorities, and requirements. It establishes role frameworks that enable effective hiring, performance management, and career development.

## Metadata Schema

```yaml
---
id: ROL-{identifier}
type: ROL
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
| Bronze | Clear role title and reporting relationship, Basic responsibilities and requirements, Essential qualifications and experience |
| Silver | Detailed responsibility breakdown with decision authority, Comprehensive skill and competency requirements, Performance metrics and success criteria |
| Gold | Full documentation with validation evidence |

## Relationships

**Depends on**: ORG, TEA, SKI

**Enables**: PER, QUA, DEV

---

- [Back to Document Types](/docs/document-types)
- [Domain: Operations & Execution](/docs/domains/operations-execution)
