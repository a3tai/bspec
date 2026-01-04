---
title: "IGN: Insight Generation"
description: "BSpec IGN document type specification"
---

# IGN: Insight Generation

**Domain**: Growth & Innovation

The Insight Generation document defines systematic approaches to transforming information and experience into actionable business insights. It establishes insight frameworks that convert data, observations, and knowledge into strategic intelligence that drives better decision-making, innovation, and competitive advantage.

## Metadata Schema

```yaml
---
id: IGN-{identifier}
type: IGN
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
| Bronze | Basic insight generation with simple analytical capabilities, Simple data collection and basic analysis processes, Basic insight documentation and sharing |
| Silver | Comprehensive insight generation with advanced analytical capabilities, Structured insight management with quality assurance and validation, Active insight application with decision support systems |
| Gold | Full documentation with validation evidence |

## Relationships

**Depends on**: DAT, ANA, LEA

**Enables**: STR, INN, FUT

---

- [Back to Document Types](/docs/document-types)
- [Domain: Growth & Innovation](/docs/domains/growth-innovation)
