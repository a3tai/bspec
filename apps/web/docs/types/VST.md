---
title: "VST: Value Stream"
description: "BSpec VST document type specification"
---

# VST: Value Stream

**Domain**: Business Model

The Value Stream defines systematic analysis and optimization of end-to-end value creation processes that deliver customer value. It establishes value flow frameworks that eliminate waste, optimize performance, and align organizational activities with customer value realization.

## Metadata Schema

```yaml
---
id: VST-{identifier}
type: VST
status: Draft|Review|Accepted|Deprecated
version: 1.0.0
domain: business-model
priority: critical|high|medium|low

depends_on: []
enables: []
---
```

## Quality Levels

| Level | Requirements |
|-------|--------------|
| Bronze | Value stream overview with purpose, customer value, and strategic importance, Value stream strategy with alignment, proposition, and competitive context, Basic value stream definition with scope, flow, and customer integration |
| Silver | Comprehensive current state analysis with process mapping and waste identification, Future state design with vision, redesign, and technology enablement, Value stream metrics across flow, quality, financial, and innovation dimensions |
| Gold | Full documentation with validation evidence |

## Relationships

**Depends on**: KAC, CJM, PRO

**Enables**: PER, QUA, CUS

---

- [Back to Document Types](/docs/document-types)
- [Domain: Business Model](/docs/domains/business-model)
