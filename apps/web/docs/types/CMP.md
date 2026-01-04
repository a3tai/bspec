---
title: "CMP: Competitive Analysis"
description: "BSpec CMP document type specification"
---

# CMP: Competitive Analysis

**Domain**: Market Environment

The Competitive Analysis document maps the competitive landscape, analyzes key competitors, and identifies competitive threats and opportunities. It provides intelligence for strategic positioning and tactical responses.

## Metadata Schema

```yaml
---
id: CMP-{identifier}
type: CMP
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
| Bronze | Key competitors identified and profiled, Basic competitive positioning analysis, Competitive threats and opportunities identified |
| Silver | Comprehensive competitor profiles with financial and strategic analysis, Competitive dynamics and market structure analysis, Win/loss analysis with customer feedback |
| Gold | Full documentation with validation evidence |

## Relationships

**Depends on**: MKT, SEG

**Enables**: POS, MOT, STR

---

- [Back to Document Types](/docs/document-types)
- [Domain: Market Environment](/docs/domains/market-environment)
