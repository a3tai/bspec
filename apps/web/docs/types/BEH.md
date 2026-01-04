---
title: "BEH: Behaviors"
description: "BSpec BEH document type specification"
---

# BEH: Behaviors

**Domain**: Customer Value

The Behaviors document analyzes customer usage patterns, behavioral data, and interaction analytics to understand how customers actually use products and services, revealing gaps between stated preferences and actual behavior.

## Metadata Schema

```yaml
---
id: BEH-{identifier}
type: BEH
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
| Bronze | Analysis framework with research objectives and data collection approach, Basic usage behavior analysis with activity and engagement metrics, Pain point identification with friction analysis |
| Silver | Comprehensive behavioral patterns analysis by segment and persona validation, Detailed engagement and retention analysis with churn behavior insights, Success pattern analysis with high-performance behaviors and value realization |
| Gold | Full documentation with validation evidence |

## Relationships

**Depends on**: PER, USE

**Enables**: REQ, UXD, GAI

---

- [Back to Document Types](/docs/document-types)
- [Domain: Customer Value](/docs/domains/customer-value)
