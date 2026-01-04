---
title: "THR: Threats"
description: "BSpec THR document type specification"
---

# THR: Threats

**Domain**: Market Environment

The Threats document identifies external risks to market position and business model. It analyzes potential threats from competitors, market changes, and environmental factors.

## Metadata Schema

```yaml
---
id: THR-{identifier}
type: THR
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
| Bronze | Key threats identified and categorized, Basic probability and impact assessment, Early warning indicators defined |
| Silver | Comprehensive threat analysis with interactions, Detailed impact assessment and vulnerability analysis, Mitigation strategies and response plans |
| Gold | Full documentation with validation evidence |

## Relationships

**Depends on**: CMP, TRN, MAC

**Enables**: RSK, MIT, STR

---

- [Back to Document Types](/docs/document-types)
- [Domain: Market Environment](/docs/domains/market-environment)
