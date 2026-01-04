---
title: "KAC: Key Activities"
description: "BSpec KAC document type specification"
---

# KAC: Key Activities

**Domain**: Business Model

The Key Activities defines systematic approaches to critical business activities that create customer value, drive competitive advantage, and enable business model execution. It establishes activity frameworks that optimize operational excellence, resource utilization, and performance management.

## Metadata Schema

```yaml
---
id: KAC-{identifier}
type: KAC
status: Draft|Review|Accepted|Deprecated
version: 1.0.0
domain: business-model
priority: critical|high|medium|low

depends_on: []
enables: []
---
```

## Quality Levels

| Level | Requirements |
|-------|--------------|
| Bronze | Activity overview with purpose, strategic role, and value creation, Activity strategy with context, philosophy, and focus areas, Basic activity portfolio with production, platform, and problem-solving categories |
| Silver | Comprehensive activity analysis with value stream analysis and process mapping, Performance management with metrics, optimization, and improvement approaches, Capability requirements with development strategies and gap analysis |
| Gold | Full documentation with validation evidence |

## Relationships

**Depends on**: STR, KRS, VAL

**Enables**: PER, QUA, PRO

---

- [Back to Document Types](/docs/document-types)
- [Domain: Business Model](/docs/domains/business-model)
