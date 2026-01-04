---
title: "OBJ: Strategic Objectives"
description: "BSpec OBJ document type specification"
---

# OBJ: Strategic Objectives

**Domain**: Strategic Foundation

The Objectives document sets specific, measurable goals with timeframes that advance the organization toward its vision. These are typically structured as OKRs (Objectives and Key Results) or similar goal-setting frameworks.

## Metadata Schema

```yaml
---
id: OBJ-{identifier}
type: OBJ
status: Draft|Review|Accepted|Deprecated
version: 1.0.0
domain: strategic-foundation
priority: critical|high|medium|low

depends_on: []
enables: []
---
```

## Quality Levels

| Level | Requirements |
|-------|--------------|
| Bronze | 3-5 clear strategic objectives defined, Each objective has 2-3 measurable key results, Time horizons and owners specified |
| Silver | Strategic rationale clear for each objective, Baselines and targets quantified, Cascade framework to departments |
| Gold | Full documentation with validation evidence |

## Relationships

**Depends on**: STR, VSN

**Enables**: MET, PRC, GTM

---

- [Back to Document Types](/docs/document-types)
- [Domain: Strategic Foundation](/docs/domains/strategic-foundation)
