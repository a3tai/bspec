---
title: "MET: Metrics"
description: "BSpec MET document type specification"
---

# MET: Metrics

**Domain**: Financial & Investment

The Metrics document defines systematic approaches to measuring, monitoring, and managing business performance through key performance indicators, financial metrics, and operational measures. It establishes measurement frameworks that enable data-driven decision making, performance accountability, and continuous improvement.

## Metadata Schema

```yaml
---
id: MET-{identifier}
type: MET
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
| Bronze | Basic financial and operational metrics with manual reporting, Simple dashboard with key performance indicators, Monthly performance review and variance analysis |
| Silver | Comprehensive metric framework with balanced scorecard approach, Automated data collection and reporting with quality controls, Interactive dashboards with drill-down capabilities |
| Gold | Full documentation with validation evidence |

## Relationships

**Depends on**: FIN, BUD, STR

**Enables**: REP, PER, QUA

---

- [Back to Document Types](/docs/document-types)
- [Domain: Financial & Investment](/docs/domains/financial-investment)
