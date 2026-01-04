---
title: "MKT: Market Definition"
description: "BSpec MKT document type specification"
---

# MKT: Market Definition

**Domain**: Market Environment

The Market Definition document establishes the boundaries, size, and characteristics of the addressable market. It defines TAM (Total Addressable Market), SAM (Serviceable Addressable Market), and SOM (Serviceable Obtainable Market).

## Metadata Schema

```yaml
---
id: MKT-{identifier}
type: MKT
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
| Bronze | Clear market definition with boundaries, TAM, SAM, and SOM estimates provided, Basic market characteristics described |
| Silver | Multiple calculation methods for market sizing, Market research foundation documented, Market dynamics and trends analyzed |
| Gold | Full documentation with validation evidence |

## Relationships

**Depends on**: MSN, STR

**Enables**: SEG, CMP, GTM

---

- [Back to Document Types](/docs/document-types)
- [Domain: Market Environment](/docs/domains/market-environment)
