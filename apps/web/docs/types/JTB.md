---
title: "JTB: Jobs-to-be-Done"
description: "BSpec JTB document type specification"
---

# JTB: Jobs-to-be-Done

**Domain**: Customer Value

The Jobs-to-be-Done document defines the specific outcomes customers hire products or services to achieve. It captures the functional, emotional, and social jobs that drive customer behavior and innovation opportunities.

## Metadata Schema

```yaml
---
id: JTB-{identifier}
type: JTB
status: Draft|Review|Accepted|Deprecated
version: 1.0.0
domain: customer-value
priority: critical|high|medium|low

depends_on: []
enables: []
---
```

## Quality Levels

| Level | Requirements |
|-------|--------------|
| Bronze | Job statement clearly defined with situation-motivation-outcome, Job categories (functional, emotional, social) identified, Basic job process and desired outcomes documented |
| Silver | Comprehensive job context and constraints analysis, Detailed job process with steps and workflow, Current solution evaluation and gap analysis |
| Gold | Full documentation with validation evidence |

## Relationships

**Depends on**: PER

**Enables**: USE, PRD, SVC

---

- [Back to Document Types](/docs/document-types)
- [Domain: Customer Value](/docs/domains/customer-value)
