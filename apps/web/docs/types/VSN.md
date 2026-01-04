---
title: "VSN: Vision Statement"
description: "BSpec VSN document type specification"
---

# VSN: Vision Statement

**Domain**: Strategic Foundation

The Vision document articulates the future state the organization aims to create—both for itself and the world. It describes what success looks like at a meaningful time horizon (typically 3-10 years).

## Metadata Schema

```yaml
---
id: VSN-{identifier}
type: VSN
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
| Bronze | Clear vision statement exists (1-3 sentences), Future state description provided, Basic success metrics identified |
| Silver | Compelling and inspirational vision, Quantitative and qualitative success metrics, Detailed timeline with phases |
| Gold | Full documentation with validation evidence |

## Relationships

**Depends on**: MSN, VAL

**Enables**: STR, OBJ, GTM

---

- [Back to Document Types](/docs/document-types)
- [Domain: Strategic Foundation](/docs/domains/strategic-foundation)
