---
title: "SLA: Service Level Agreement"
description: "BSpec SLA document type specification"
---

# SLA: Service Level Agreement

**Domain**: Operations & Execution

The Service Level Agreement defines systematic approaches to establishing, measuring, and managing service level commitments that ensure consistent service delivery and customer satisfaction. It establishes performance frameworks that optimize service quality and accountability.

## Metadata Schema

```yaml
---
id: SLA-{identifier}
type: SLA
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
| Bronze | SLA overview with purpose, parties, service scope, and business context, Service definition with description, classification, and availability, Basic service level commitments with availability and performance targets |
| Silver | Comprehensive measurement and monitoring with KPIs, systems, and reporting, Roles and responsibilities with clear provider, user, and shared accountability, Exception handling with planned and unplanned exception procedures |
| Gold | Full documentation with validation evidence |

## Relationships

**Depends on**: SVC, PRO, CAP

**Enables**: QUA, CUS, SUP

---

- [Back to Document Types](/docs/document-types)
- [Domain: Operations & Execution](/docs/domains/operations-execution)
