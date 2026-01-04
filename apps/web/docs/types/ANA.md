---
title: "ANA: Analytics"
description: "BSpec ANA document type specification"
---

# ANA: Analytics

**Domain**: Technology & Data

The Analytics document defines systematic approaches to designing, implementing, and managing analytics capabilities that enable data-driven decision making through business intelligence, advanced analytics, and strategic insights. It establishes analytics frameworks that ensure business value, data quality, and user adoption.

## Metadata Schema

```yaml
---
id: ANA-{identifier}
type: ANA
status: Draft|Review|Accepted|Deprecated
version: 1.0.0
domain: technology-data
priority: critical|high|medium|low

depends_on: []
enables: []
---
```

## Quality Levels

| Level | Requirements |
|-------|--------------|
| Bronze | Basic reporting and dashboard capabilities, Simple data integration and basic data quality, Manual analytics processes with limited self-service |
| Silver | Comprehensive BI platform with self-service capabilities, Structured data governance and quality management, Advanced analytics capabilities with some automation |
| Gold | Full documentation with validation evidence |

## Relationships

**Depends on**: DAT, SYS, INF

**Enables**: PER, QUA, GOV

---

- [Back to Document Types](/docs/document-types)
- [Domain: Technology & Data](/docs/domains/technology-data)
