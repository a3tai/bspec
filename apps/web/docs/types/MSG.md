---
title: "MSG: Messaging Framework"
description: "BSpec MSG document type specification"
---

# MSG: Messaging Framework

**Domain**: Brand & Marketing

The Messaging Framework document defines what the brand says and how it says it to different audiences across various touchpoints. It establishes messaging architecture that ensures consistent communication, builds brand recognition, and drives desired customer actions through strategic message development and adaptation.

## Metadata Schema

```yaml
---
id: MSG-{identifier}
type: MSG
status: Draft|Review|Accepted|Deprecated
version: 1.0.0
domain: brand-marketing
priority: critical|high|medium|low

depends_on: []
enables: []
---
```

## Quality Levels

| Level | Requirements |
|-------|--------------|
| Bronze | Basic messaging framework with core messages and key points, Simple audience adaptation and channel guidelines, Basic message testing and validation processes |
| Silver | Comprehensive messaging framework with detailed architecture, Structured audience segmentation with channel-specific adaptation, Active message testing with performance measurement and optimization |
| Gold | Full documentation with validation evidence |

## Relationships

**Depends on**: BRD, POS, CUS

**Enables**: CNT, CAM, TON

---

- [Back to Document Types](/docs/document-types)
- [Domain: Brand & Marketing](/docs/domains/brand-marketing)
