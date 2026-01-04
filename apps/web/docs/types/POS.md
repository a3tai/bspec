---
title: "POS: Positioning"
description: "BSpec POS document type specification"
---

# POS: Positioning

**Domain**: Market Environment

The Positioning document defines how the organization wants to be perceived in the market relative to competitors and alternatives. It articulates the unique value proposition and market position.

## Metadata Schema

```yaml
---
id: POS-{identifier}
type: POS
status: Draft|Review|Accepted|Deprecated
version: 1.0.0
domain: market-environment
priority: critical|high|medium|low

depends_on: []
enables: []
---
```

## Quality Levels

| Level | Requirements |
|-------|--------------|
| Bronze | Clear positioning statement defined, Basic positioning elements identified, Target customer and competitive frame established |
| Silver | Comprehensive positioning strategy with competitive context, Market perception analysis and building strategy, Positioning implementation plan for internal and external execution |
| Gold | Full documentation with validation evidence |

## Relationships

**Depends on**: SEG, CMP, VAL

**Enables**: GTM, BMC, REV

---

- [Back to Document Types](/docs/document-types)
- [Domain: Market Environment](/docs/domains/market-environment)
