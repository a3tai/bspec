---
title: "CAM: Marketing Campaign"
description: "BSpec CAM document type specification"
---

# CAM: Marketing Campaign

**Domain**: Brand & Marketing

The Marketing Campaign document defines integrated marketing initiatives that achieve specific business objectives through coordinated messaging, creative execution, and multi-channel activation. It establishes campaign frameworks that drive awareness, engagement, and conversion while maintaining brand consistency and maximizing return on marketing investment.

## Metadata Schema

```yaml
---
id: CAM-{identifier}
type: CAM
status: Draft|Review|Accepted|Deprecated
version: 1.0.0
domain: brand-marketing
priority: critical|high|medium|low

depends_on: []
enables: []
---
```

## Quality Levels

| Level | Requirements |
|-------|--------------|
| Bronze | Basic campaign strategy with clear objectives and target audience, Simple creative concept and multi-channel execution plan, Basic performance tracking and measurement framework |
| Silver | Comprehensive campaign strategy with integrated channel approach, Structured content creation with performance optimization, Active campaign management with real-time optimization |
| Gold | Full documentation with validation evidence |

## Relationships

**Depends on**: BRD, MSG, POS

**Enables**: LED, CON, BRA

---

- [Back to Document Types](/docs/document-types)
- [Domain: Brand & Marketing](/docs/domains/brand-marketing)
