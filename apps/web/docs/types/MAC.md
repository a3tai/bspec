---
title: "MAC: Macro Environment"
description: "BSpec MAC document type specification"
---

# MAC: Macro Environment

**Domain**: Market Environment

The Macro Environment document analyzes broad economic, political, social, and technological factors that influence the business environment. It examines PESTEL factors and their implications.

## Metadata Schema

```yaml
---
id: MAC-{identifier}
type: MAC
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
| Bronze | PESTEL factors identified and analyzed, Basic regional analysis for major markets, Key macro trends and implications documented |
| Silver | Comprehensive PESTEL analysis with interconnections, Scenario planning with strategic implications, Regional analysis covering all major markets |
| Gold | Full documentation with validation evidence |

## Relationships

**Depends on**: STR

**Enables**: TRN, THR, OPP

---

- [Back to Document Types](/docs/document-types)
- [Domain: Market Environment](/docs/domains/market-environment)
