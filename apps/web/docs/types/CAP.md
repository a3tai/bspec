---
title: "CAP: Capability Specification"
description: "BSpec CAP document type specification"
---

# CAP: Capability Specification

**Domain**: Operations & Execution

The Capability Specification defines systematic approaches to building and managing organizational capabilities that create competitive advantage and enable business strategy execution. It establishes capability frameworks that optimize development, performance, and continuous improvement.

## Metadata Schema

```yaml
---
id: CAP-{identifier}
type: CAP
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
| Bronze | Capability overview with purpose, definition, and strategic value, Capability context with strategic alignment and classification, Basic capability definition with components and boundaries |
| Silver | Comprehensive current state assessment with maturity, strengths, and gaps analysis, Target state vision with future requirements and success criteria, Capability development plan with strategy, components, and implementation approach |
| Gold | Full documentation with validation evidence |

## Relationships

**Depends on**: STR, KRS, KAC

**Enables**: PER, QUA, SVC

---

- [Back to Document Types](/docs/document-types)
- [Domain: Operations & Execution](/docs/domains/operations-execution)
