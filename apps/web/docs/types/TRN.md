---
title: "TRN: Trends"
description: "BSpec TRN document type specification"
---

# TRN: Trends

**Domain**: Market Environment

The Trends document identifies and analyzes market forces and changes shaping the industry. It examines technology trends, social shifts, regulatory changes, and other factors that could impact the business.

## Metadata Schema

```yaml
---
id: TRN-{identifier}
type: TRN
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
| Bronze | Key trends identified and described, Basic impact assessment for each trend, Trend sources and validation documented |
| Silver | Comprehensive trend analysis with driving forces and evidence, Trend interaction analysis showing convergent and conflicting patterns, Strategic implications for business model and innovation |
| Gold | Full documentation with validation evidence |

## Relationships

**Depends on**: MKT, MAC

**Enables**: STR, INN, OPP

---

- [Back to Document Types](/docs/document-types)
- [Domain: Market Environment](/docs/domains/market-environment)
