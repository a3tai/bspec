---
title: "PER: Personas"
description: "BSpec PER document type specification"
---

# PER: Personas

**Domain**: Customer Value

The Personas document creates detailed archetypal representations of key customer segments. Personas humanize customer data and provide shared understanding of who the organization serves.

## Metadata Schema

```yaml
---
id: PER-{identifier}
type: PER
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
| Bronze | Key personas identified and categorized, Basic demographic and behavioral data documented, Core pain points and motivations captured |
| Silver | Comprehensive persona research with quantitative validation, Detailed behavioral patterns and decision-making process, Jobs-to-be-done connection established |
| Gold | Full documentation with validation evidence |

## Relationships

**Depends on**: SEG

**Enables**: JTB, USE, CJM

---

- [Back to Document Types](/docs/document-types)
- [Domain: Customer Value](/docs/domains/customer-value)
