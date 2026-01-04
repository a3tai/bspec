---
title: "INC: Incidents"
description: "BSpec INC document type specification"
---

# INC: Incidents

**Domain**: Risk & Governance

The Incidents document defines systematic approaches to identifying, responding to, and managing incidents that disrupt business operations, threaten stakeholder safety, or impact organizational objectives. It establishes incident management frameworks that enable rapid response, effective recovery, and organizational learning from disruptive events.

## Metadata Schema

```yaml
---
id: INC-{identifier}
type: INC
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
| Bronze | Basic incident response procedures with team assignments, Simple incident classification and escalation processes, Basic investigation and documentation procedures |
| Silver | Comprehensive incident management framework with crisis management, Structured investigation processes with lessons learned integration, Effective recovery procedures with stakeholder communication |
| Gold | Full documentation with validation evidence |

## Relationships

**Depends on**: RSK, OPS, CTL

**Enables**: BCR, CRI, LEA

---

- [Back to Document Types](/docs/document-types)
- [Domain: Risk & Governance](/docs/domains/risk-governance)
