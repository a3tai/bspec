---
title: "SEG: Market Segments"
description: "BSpec SEG document type specification"
---

# SEG: Market Segments

**Domain**: Market Environment

The Market Segments document identifies and analyzes distinct customer groups within the broader market. It defines segmentation criteria, profiles key segments, and prioritizes target segments.

## Metadata Schema

```yaml
---
id: SEG-{identifier}
type: SEG
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
| Bronze | Clear segmentation criteria defined, Primary segments identified and profiled, Basic size and value estimates provided |
| Silver | Comprehensive segment profiles with demographics, psychographics, and behavior, Segment prioritization with evaluation criteria, Go-to-market implications by segment |
| Gold | Full documentation with validation evidence |

## Relationships

**Depends on**: MKT, PER

**Enables**: POS, GTM, REV

---

- [Back to Document Types](/docs/document-types)
- [Domain: Market Environment](/docs/domains/market-environment)
