---
title: "REG: Regulatory Environment"
description: "BSpec REG document type specification"
---

# REG: Regulatory Environment

**Domain**: Market Environment

The Regulatory Environment document analyzes laws, regulations, and compliance requirements affecting the business. It tracks regulatory changes and assesses compliance obligations.

## Metadata Schema

```yaml
---
id: REG-{identifier}
type: REG
status: Draft|Review|Accepted|Deprecated
version: 1.0.0
domain: market-environment
priority: critical|high|medium|low

depends_on: []
enables: []
---
```

## Quality Levels

| Level | Requirements |
|-------|--------------|
| Bronze | Core regulatory requirements identified, Basic compliance obligations documented, Key regulatory bodies and contacts established |
| Silver | Comprehensive regulatory landscape mapping, Detailed compliance framework and controls, Regulatory change monitoring system |
| Gold | Full documentation with validation evidence |

## Relationships

**Depends on**: MKT, STR

**Enables**: CMP, RSK, POL

---

- [Back to Document Types](/docs/document-types)
- [Domain: Market Environment](/docs/domains/market-environment)
