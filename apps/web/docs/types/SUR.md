---
title: "SUR: Surveys"
description: "BSpec SUR document type specification"
---

# SUR: Surveys

**Domain**: Customer Value

The Surveys document captures quantitative customer research through structured questionnaires that provide statistical insights into customer attitudes, behaviors, preferences, and satisfaction levels.

## Metadata Schema

```yaml
---
id: SUR-{identifier}
type: SUR
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
| Bronze | Survey design with clear objectives and methodology, Target population and sampling strategy defined, Basic survey results with response metrics and key findings |
| Silver | Comprehensive survey implementation with distribution strategy and response management, Detailed survey results with statistical analysis and segmentation insights, Competitive intelligence and market position analysis |
| Gold | Full documentation with validation evidence |

## Relationships

**Depends on**: PER, SEG

**Enables**: FEE, MET, CMP

---

- [Back to Document Types](/docs/document-types)
- [Domain: Customer Value](/docs/domains/customer-value)
