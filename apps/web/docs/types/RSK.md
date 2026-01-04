---
title: "RSK: Risks"
description: "BSpec RSK document type specification"
---

# RSK: Risks

**Domain**: Risk & Governance

The Risks document defines systematic approaches to identifying, assessing, and managing business risks that could impact organizational objectives, operations, and stakeholder value. It establishes risk management frameworks that enable proactive risk mitigation, informed decision making, and resilient business operations.

## Metadata Schema

```yaml
---
id: RSK-{identifier}
type: RSK
status: Draft|Review|Accepted|Deprecated
version: 1.0.0
domain: risk-governance
priority: critical|high|medium|low

depends_on: []
enables: []
---
```

## Quality Levels

| Level | Requirements |
|-------|--------------|
| Bronze | Basic risk register with identification and assessment, Simple risk monitoring and reporting, Basic risk mitigation and control processes |
| Silver | Comprehensive risk framework with systematic processes, Structured risk assessment with quantitative and qualitative methods, Effective risk mitigation with performance monitoring |
| Gold | Full documentation with validation evidence |

## Relationships

**Depends on**: STR, OBJ, OPS

**Enables**: CTL, COM, GOV

---

- [Back to Document Types](/docs/document-types)
- [Domain: Risk & Governance](/docs/domains/risk-governance)
