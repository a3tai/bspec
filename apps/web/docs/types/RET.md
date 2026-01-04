---
title: "RET: Retrospective Analysis"
description: "BSpec RET document type specification"
---

# RET: Retrospective Analysis

**Domain**: Learning & Decisions

The Retrospective Analysis document captures structured reflection on completed projects, initiatives, or time periods to identify successes, failures, lessons learned, and improvement opportunities. It establishes retrospective frameworks that promote continuous learning, team development, and organizational improvement through systematic reflection and analysis.

## Metadata Schema

```yaml
---
id: RET-{identifier}
type: RET
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
| Bronze | Basic retrospective process with regular reflection and simple action planning, Simple documentation and basic follow-through on improvements, Basic retrospective facilitation and team participation |
| Silver | Comprehensive retrospective framework with structured analysis and improvement planning, Structured implementation tracking with measurement and accountability, Active retrospective culture with quality standards and continuous enhancement |
| Gold | Full documentation with validation evidence |

## Relationships

**Depends on**: PRO, PRJ, OBJ

**Enables**: LRN, PRO, POL

---

- [Back to Document Types](/docs/document-types)
- [Domain: Learning & Decisions](/docs/domains/learning-decisions)
