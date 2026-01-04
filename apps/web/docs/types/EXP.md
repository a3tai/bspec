---
title: "EXP: Experimentation"
description: "BSpec EXP document type specification"
---

# EXP: Experimentation

**Domain**: Growth & Innovation

The Experimentation document defines systematic approaches to testing assumptions and validating new ideas through structured experimentation. It establishes experimentation frameworks that transform hypothesis testing from ad-hoc activities into systematic business capabilities that drive evidence-based decision-making and rapid learning.

## Metadata Schema

```yaml
---
id: EXP-{identifier}
type: EXP
status: Draft|Review|Accepted|Deprecated
version: 1.0.0
domain: growth-innovation
priority: critical|high|medium|low

depends_on: []
enables: []
---
```

## Quality Levels

| Level | Requirements |
|-------|--------------|
| Bronze | Basic experimentation framework with simple A/B testing capabilities, Simple hypothesis development and testing processes, Basic experimental measurement and analysis |
| Silver | Comprehensive experimentation framework with rigorous experimental design, Structured experiment portfolio management with quality assurance, Active experimentation culture with systematic insight integration |
| Gold | Full documentation with validation evidence |

## Relationships

**Depends on**: INN, LEA, DAT

**Enables**: PRD, SVC, STR

---

- [Back to Document Types](/docs/document-types)
- [Domain: Growth & Innovation](/docs/domains/growth-innovation)
