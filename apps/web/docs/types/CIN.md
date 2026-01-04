---
title: "CIN: Interviews"
description: "BSpec CIN document type specification"
---

# CIN: Interviews

**Domain**: Customer Value

The Interviews document captures structured customer research conversations that provide deep insights into needs, behaviors, motivations, and experiences. It documents qualitative research findings that inform product and strategy decisions.

## Metadata Schema

```yaml
---
id: CIN-{identifier}
type: CIN
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
| Bronze | Research objectives clearly defined with primary questions and learning goals, Interview methodology documented with structure and participant selection, Basic interview findings with key themes and insights |
| Silver | Comprehensive research design with detailed methodology and quality assessment, Rich interview findings with detailed themes, quotes, and evidence, Segment-specific insights and persona validation |
| Gold | Full documentation with validation evidence |

## Relationships

**Depends on**: PER, JTB

**Enables**: EMP, REQ, USE

---

- [Back to Document Types](/docs/document-types)
- [Domain: Customer Value](/docs/domains/customer-value)
