---
title: "HYP: Hypothesis Management"
description: "BSpec HYP document type specification"
---

# HYP: Hypothesis Management

**Domain**: Learning & Decisions

The Hypothesis Management document captures testable assumptions and beliefs about business, customers, markets, and solutions that guide organizational decision-making and learning. It establishes hypothesis frameworks that enable evidence-based validation, systematic experimentation, and continuous learning through iterative hypothesis development and testing.

## Metadata Schema

```yaml
---
id: HYP-{identifier}
type: HYP
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
| Bronze | Basic hypothesis formulation with simple testing and validation approaches, Simple tracking and basic learning capture from hypothesis work, Basic hypothesis governance and quality standards |
| Silver | Comprehensive hypothesis framework with systematic development and testing, Structured lifecycle management with portfolio approach and learning integration, Active hypothesis governance with quality standards and continuous improvement |
| Gold | Full documentation with validation evidence |

## Relationships

**Depends on**: THY, STR, CUS

**Enables**: EXP, LRN, DEC

---

- [Back to Document Types](/docs/document-types)
- [Domain: Learning & Decisions](/docs/domains/learning-decisions)
