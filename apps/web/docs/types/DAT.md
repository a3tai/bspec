---
title: "DAT: Data Models"
description: "BSpec DAT document type specification"
---

# DAT: Data Models

**Domain**: Technology & Data

The Data Models document defines systematic approaches to designing, governing, and managing data structures that support business capabilities through coherent data architecture, quality assurance, and strategic data management. It establishes data frameworks that ensure consistency, integrity, and value creation from organizational data assets.

## Metadata Schema

```yaml
---
id: DAT-{identifier}
type: DAT
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
| Bronze | Basic data model with core entities and relationships, Simple data governance and ownership definition, Basic data quality and security measures |
| Silver | Comprehensive logical and physical data model, Detailed data governance framework with quality controls, Structured data integration and lifecycle management |
| Gold | Full documentation with validation evidence |

## Relationships

**Depends on**: SYS, ARC, REQ

**Enables**: API, ANA, QUA

---

- [Back to Document Types](/docs/document-types)
- [Domain: Technology & Data](/docs/domains/technology-data)
