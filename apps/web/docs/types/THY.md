---
title: "THY: Theory of Change"
description: "BSpec THY document type specification"
---

# THY: Theory of Change

**Domain**: Strategic Foundation

The Theory of Change document maps the logic model connecting the organization's activities to its intended outcomes. It explains how and why specific actions will lead to desired changes.

## Metadata Schema

```yaml
---
id: THY-{identifier}
type: THY
status: Draft|Review|Accepted|Deprecated
version: 1.0.0
domain: strategic-foundation
priority: critical|high|medium|low

depends_on: []
enables: []
---
```

## Quality Levels

| Level | Requirements |
|-------|--------------|
| Bronze | Clear logic model from activities to outcomes, Ultimate goal and intended outcomes defined, Core activities and outputs identified |
| Silver | Detailed causal assumptions articulated, Evidence base supporting theory provided, Measurement framework for outcomes |
| Gold | Full documentation with validation evidence |

## Relationships

**Depends on**: PUR, MSN, STR

**Enables**: MET, EXP, LRN

---

- [Back to Document Types](/docs/document-types)
- [Domain: Strategic Foundation](/docs/domains/strategic-foundation)
