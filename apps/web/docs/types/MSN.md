---
title: "MSN: Mission Statement"
description: "BSpec MSN document type specification"
---

# MSN: Mission Statement

**Domain**: Strategic Foundation

This specification SHALL define the requirements and structure for Mission Statement documents that articulate an organization's fundamental purpose and reason for existence. Organizations MUST use this specification to document their core mission in a manner that enables strategic alignment, decision-making guidance, and stakeholder communication.

## Metadata Schema

```yaml
---
id: MSN-{identifier}
type: MSN
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
| Bronze | Clear, one-sentence mission statement exists, Primary beneficiaries are identified, Basic value creation is described |
| Silver | Mission statement is memorable and distinctive, Detailed beneficiary analysis completed, Value creation is quantifiable |
| Gold | Full documentation with validation evidence |

## Relationships

**Depends on**: Varies by organization

**Enables**: VSN, VAL, STR

---

- [Back to Document Types](/docs/document-types)
- [Domain: Strategic Foundation](/docs/domains/strategic-foundation)
