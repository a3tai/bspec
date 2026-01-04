---
title: "CJM: Customer Journey Map"
description: "BSpec CJM document type specification"
---

# CJM: Customer Journey Map

**Domain**: Customer Value

The Customer Journey Map document visualizes the end-to-end customer experience from awareness to advocacy. It identifies touchpoints, emotions, pain points, and opportunities across the entire customer lifecycle.

## Metadata Schema

```yaml
---
id: CJM-{identifier}
type: CJM
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
| Bronze | Key journey stages identified and documented, Major touchpoints and pain points captured, Basic emotional journey mapped |
| Silver | Comprehensive 7-stage journey mapping (Awareness to Advocacy), Detailed touchpoint analysis with experience quality metrics, Emotional journey arc with pain point impact assessment |
| Gold | Full documentation with validation evidence |

## Relationships

**Depends on**: PER, JTB

**Enables**: USE, REL, SUP

---

- [Back to Document Types](/docs/document-types)
- [Domain: Customer Value](/docs/domains/customer-value)
