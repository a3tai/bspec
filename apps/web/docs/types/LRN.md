---
title: "LRN: Learning Records"
description: "BSpec LRN document type specification"
---

# LRN: Learning Records

**Domain**: Learning & Decisions

The Learning Records document captures key discoveries, insights, and knowledge gained through organizational activities, experiments, and experiences. It establishes learning frameworks that preserve institutional knowledge, accelerate organizational learning, and enable evidence-based decision-making and continuous improvement.

## Metadata Schema

```yaml
---
id: LRN-{identifier}
type: LRN
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
| Bronze | Basic learning documentation with key insights and supporting evidence, Simple sharing mechanisms and basic application processes, Basic quality assurance and learning governance |
| Silver | Comprehensive learning framework with systematic discovery and validation, Structured application processes with impact measurement, Active learning governance with quality standards and knowledge management |
| Gold | Full documentation with validation evidence |

## Relationships

**Depends on**: EXP, HYP, DEC

**Enables**: STR, PRO, POL

---

- [Back to Document Types](/docs/document-types)
- [Domain: Learning & Decisions](/docs/domains/learning-decisions)
