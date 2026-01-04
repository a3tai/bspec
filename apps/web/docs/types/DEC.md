---
title: "DEC: Decision Records"
description: "BSpec DEC document type specification"
---

# DEC: Decision Records

**Domain**: Learning & Decisions

The Decision Records document captures strategic and operational decisions made within the organization, including context, rationale, alternatives considered, and outcomes. It establishes decision frameworks that preserve institutional knowledge, enable accountability, and provide historical context for future decision-making.

## Metadata Schema

```yaml
---
id: DEC-{identifier}
type: DEC
status: Draft|Review|Accepted|Deprecated
version: 1.0.0
domain: learning-decisions
priority: critical|high|medium|low

depends_on: []
enables: []
---
```

## Quality Levels

| Level | Requirements |
|-------|--------------|
| Bronze | Basic decision documentation with context and rationale, Simple implementation tracking and basic accountability, Basic decision review and learning capture process |
| Silver | Comprehensive decision framework with structured analysis and evaluation, Structured implementation planning with monitoring and tracking, Active decision governance with quality standards and accountability |
| Gold | Full documentation with validation evidence |

## Relationships

**Depends on**: STR, OBJ, THY

**Enables**: ROL, PRO, POL

---

- [Back to Document Types](/docs/document-types)
- [Domain: Learning & Decisions](/docs/domains/learning-decisions)
