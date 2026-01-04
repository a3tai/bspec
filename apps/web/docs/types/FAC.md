---
title: "FAC: Facilities Management"
description: "BSpec FAC document type specification"
---

# FAC: Facilities Management

**Domain**: Operations & Execution

The Facilities Management document defines systematic approaches to planning, operating, and maintaining physical facilities that support business operations through efficient space utilization, operational excellence, and employee productivity. It establishes facility frameworks that optimize cost, safety, and performance.

## Metadata Schema

```yaml
---
id: FAC-{identifier}
type: FAC
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
| Bronze | Basic facility information and space allocation, Essential safety and security procedures, Simple cost tracking and budget management |
| Silver | Comprehensive facility operations and management framework, Detailed safety, security, and compliance procedures, Structured cost management and optimization |
| Gold | Full documentation with validation evidence |

## Relationships

**Depends on**: ORG, TEA, POL

**Enables**: PER, QUA, CST

---

- [Back to Document Types](/docs/document-types)
- [Domain: Operations & Execution](/docs/domains/operations-execution)
