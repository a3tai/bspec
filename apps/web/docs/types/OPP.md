---
title: "OPP: Opportunities"
description: "BSpec OPP document type specification"
---

# OPP: Opportunities

**Domain**: Market Environment

The Opportunities document identifies market gaps, growth potential, and strategic opportunities available to the organization. It analyzes opportunity attractiveness and prioritizes pursuit strategies.

## Metadata Schema

```yaml
---
id: OPP-{identifier}
type: OPP
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
| Bronze | Key opportunities identified and described, Basic sizing and business case provided, Prioritization criteria established |
| Silver | Comprehensive opportunity analysis with validation, Detailed business cases and ROI projections, Pursuit strategies and resource requirements |
| Gold | Full documentation with validation evidence |

## Relationships

**Depends on**: MKT, TRN, CMP

**Enables**: STR, INN, EXP

---

- [Back to Document Types](/docs/document-types)
- [Domain: Market Environment](/docs/domains/market-environment)
