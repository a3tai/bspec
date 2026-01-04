---
title: "EMP: Empathy Maps"
description: "BSpec EMP document type specification"
---

# EMP: Empathy Maps

**Domain**: Customer Value

The Empathy Maps document captures deep understanding of customer thoughts, feelings, behaviors, and environment. It provides holistic view of customer experience and emotional context that drives behavior.

## Metadata Schema

```yaml
---
id: EMP-{identifier}
type: EMP
status: Draft|Review|Accepted|Deprecated
version: 1.0.0
domain: customer-value
priority: critical|high|medium|low

depends_on: []
enables: []
---
```

## Quality Levels

| Level | Requirements |
|-------|--------------|
| Bronze | Six empathy layers documented (THINKS, FEELS, SEES, SAYS, DOES, PAINS/GAINS), Primary persona and scenario context clearly defined, Basic research foundation with customer interviews/observation |
| Silver | Comprehensive empathy mapping across all six layers with detailed insights, Research foundation with multiple validation methods and quality assessment, Empathy synthesis with key insights and design implications |
| Gold | Full documentation with validation evidence |

## Relationships

**Depends on**: PER, JTB

**Enables**: CJM, UXD, REL

---

- [Back to Document Types](/docs/document-types)
- [Domain: Customer Value](/docs/domains/customer-value)
