---
title: "CTL: Controls"
description: "BSpec CTL document type specification"
---

# CTL: Controls

**Domain**: Risk & Governance

The Controls document defines systematic approaches to designing, implementing, and operating internal controls that mitigate business risks, ensure compliance, and support reliable business operations. It establishes control frameworks that provide reasonable assurance for achieving business objectives while managing risk exposure.

## Metadata Schema

```yaml
---
id: CTL-{identifier}
type: CTL
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
| Bronze | Basic control documentation with key controls identified, Simple control testing and monitoring procedures, Basic control owner assignment and accountability |
| Silver | Comprehensive control framework with systematic design, Structured control testing with effectiveness assessment, Regular control monitoring with performance metrics |
| Gold | Full documentation with validation evidence |

## Relationships

**Depends on**: RSK, PRO, SYS

**Enables**: AUD, GOV, QUA

---

- [Back to Document Types](/docs/document-types)
- [Domain: Risk & Governance](/docs/domains/risk-governance)
