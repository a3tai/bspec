---
title: "VND: Vendor Management"
description: "BSpec VND document type specification"
---

# VND: Vendor Management

**Domain**: Operations & Execution

The Vendor Management document defines systematic approaches to selecting, managing, and optimizing vendor relationships that deliver business value through effective partnership, performance management, and risk mitigation. It establishes vendor frameworks that ensure quality service delivery, cost optimization, and strategic alignment.

## Metadata Schema

```yaml
---
id: VND-{identifier}
type: VND
status: Draft|Review|Accepted|Deprecated
version: 1.0.0
domain: operations-execution
priority: critical|high|medium|low

depends_on: []
enables: []
---
```

## Quality Levels

| Level | Requirements |
|-------|--------------|
| Bronze | Basic vendor information and contact details, Clear service definitions and basic performance expectations, Simple contract and pricing framework |
| Silver | Comprehensive vendor assessment and capability mapping, Detailed performance management framework with KPIs, Structured risk management and contingency planning |
| Gold | Full documentation with validation evidence |

## Relationships

**Depends on**: PRO, POL, SLA

**Enables**: PER, QUA, CST

---

- [Back to Document Types](/docs/document-types)
- [Domain: Operations & Execution](/docs/domains/operations-execution)
