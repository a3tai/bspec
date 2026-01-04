---
title: "VAL: Organizational Values"
description: "BSpec VAL document type specification"
---

# VAL: Organizational Values

**Domain**: Strategic Foundation

The Values document defines the guiding principles that shape culture, guide decisions, and determine how the organization behaves. Values are the non-negotiable beliefs that influence every action.

## Metadata Schema

```yaml
---
id: VAL-{identifier}
type: VAL
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
| Bronze | 3-7 core values clearly defined, Basic behavioral descriptions for each value, Values origin story documented |
| Silver | Detailed behavioral indicators for each value, Decision frameworks that incorporate values, Anti-patterns clearly identified |
| Gold | Full documentation with validation evidence |

## Relationships

**Depends on**: MSN

**Enables**: ORG, POL, ETH

---

- [Back to Document Types](/docs/document-types)
- [Domain: Strategic Foundation](/docs/domains/strategic-foundation)
