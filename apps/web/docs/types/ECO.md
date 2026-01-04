---
title: "ECO: Ecosystem"
description: "BSpec ECO document type specification"
---

# ECO: Ecosystem

**Domain**: Market Environment

The Ecosystem document maps the network of partners, suppliers, distributors, and other stakeholders that create value around the organization. It analyzes ecosystem dynamics and partnership strategies.

## Metadata Schema

```yaml
---
id: ECO-{identifier}
type: ECO
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
| Bronze | Core ecosystem participants identified, Basic partnership strategy outlined, Key dependencies and risks assessed |
| Silver | Comprehensive ecosystem mapping with dynamics analysis, Partner portfolio strategy with prioritization criteria, Value creation mechanisms quantified |
| Gold | Full documentation with validation evidence |

## Relationships

**Depends on**: MKT, STR

**Enables**: PRT, CHN, VND

---

- [Back to Document Types](/docs/document-types)
- [Domain: Market Environment](/docs/domains/market-environment)
