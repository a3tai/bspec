---
title: "OPS: Operations Manual"
description: "BSpec OPS document type specification"
---

# OPS: Operations Manual

**Domain**: Operations & Execution

The Operations Manual defines systematic approaches to day-to-day operational management that ensure consistent service delivery, efficient resource utilization, and continuous operational excellence. It establishes operational frameworks that optimize performance and reliability.

## Metadata Schema

```yaml
---
id: OPS-{identifier}
type: OPS
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
| Bronze | Operations overview with purpose, scope, audience, and authority, Operational context with business alignment and operating model, Basic operational structure with teams, roles, and responsibilities |
| Silver | Comprehensive service delivery operations with portfolio, processes, and quality assurance, Standard operating procedures with daily, weekly, and monthly operations, Monitoring and performance management with KPIs, systems, and reporting |
| Gold | Full documentation with validation evidence |

## Relationships

**Depends on**: PRO, SLA, CAP

**Enables**: PER, QUA, MON

---

- [Back to Document Types](/docs/document-types)
- [Domain: Operations & Execution](/docs/domains/operations-execution)
