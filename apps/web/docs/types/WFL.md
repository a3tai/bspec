---
title: "WFL: Workflow Specification"
description: "BSpec WFL document type specification"
---

# WFL: Workflow Specification

**Domain**: Operations & Execution

The Workflow Specification defines systematic approaches to designing, implementing, and managing business workflows that automate and optimize operational processes. It establishes workflow frameworks that ensure efficient execution, robust exception handling, and continuous improvement.

## Metadata Schema

```yaml
---
id: WFL-{identifier}
type: WFL
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
| Bronze | Workflow overview with purpose, business context, and value creation, Workflow definition with scope, classification, and business rules, Basic workflow design with triggers, steps, and paths |
| Silver | Comprehensive workflow execution with model, task management, and error handling, Automation and technology with strategy, integration, and engine specification, Performance and monitoring with metrics, framework, and analytics |
| Gold | Full documentation with validation evidence |

## Relationships

**Depends on**: PRO, CAP, SVC

**Enables**: PER, QUA, MON

---

- [Back to Document Types](/docs/document-types)
- [Domain: Operations & Execution](/docs/domains/operations-execution)
