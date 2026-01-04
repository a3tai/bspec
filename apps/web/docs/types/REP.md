---
title: "REP: Reporting"
description: "BSpec REP document type specification"
---

# REP: Reporting

**Domain**: Financial & Investment

The Reporting document defines systematic approaches to financial and business reporting that provide stakeholders with accurate, timely, and relevant information for decision making. It establishes reporting frameworks that ensure regulatory compliance, transparency, and effective communication of business performance.

## Metadata Schema

```yaml
---
id: REP-{identifier}
type: REP
status: Draft|Review|Accepted|Deprecated
version: 1.0.0
domain: financial-investment
priority: critical|high|medium|low

depends_on: []
enables: []
---
```

## Quality Levels

| Level | Requirements |
|-------|--------------|
| Bronze | Basic financial statements with manual preparation, Simple management reports with key metrics, Basic regulatory compliance and filing |
| Silver | Comprehensive financial reporting with detailed analysis, Automated management reporting with dashboards, Full regulatory compliance with quality controls |
| Gold | Full documentation with validation evidence |

## Relationships

**Depends on**: MET, FIN, BUD

**Enables**: GOV, COM, INV

---

- [Back to Document Types](/docs/document-types)
- [Domain: Financial & Investment](/docs/domains/financial-investment)
