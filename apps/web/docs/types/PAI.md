---
title: "PAI: Pain Points"
description: "BSpec PAI document type specification"
---

# PAI: Pain Points

**Domain**: Customer Value

The Pain Points document identifies and analyzes customer problems, frustrations, and obstacles. It captures the negative experiences that drive customers to seek solutions and creates opportunities for value creation.

## Metadata Schema

```yaml
---
id: PAI-{identifier}
type: PAI
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
| Bronze | Pain point clearly defined with core problem statement, Basic customer context and affected personas identified, Pain severity and frequency assessment documented |
| Silver | Comprehensive pain analysis with root cause identification, Detailed customer impact assessment across multiple dimensions, Current solutions and workarounds analysis |
| Gold | Full documentation with validation evidence |

## Relationships

**Depends on**: PER, JTB

**Enables**: GAI, OPP, REQ

---

- [Back to Document Types](/docs/document-types)
- [Domain: Customer Value](/docs/domains/customer-value)
