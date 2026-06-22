---
title: "PFO: Porter's Five Forces"
description: "BSpec PFO document type specification"
---

# PFO: Porter's Five Forces

::: tip Document Type
**Code**: PFO<br>
**Name**: Porter's Five Forces<br>
**Domain**: Market & Environment
:::

## Abstract

This specification defines the Porter's Five Forces document type within the BSpec 1.0 Universal Business Specification Standard. It establishes a structured assessment of competitive pressure across suppliers, customers, substitutes, entrants, and rivalry, and links those findings to positioning and strategy execution.

## Framework and Attribution

This document references **Porter's Five Forces** for competitive dynamics assessment.

## Purpose and Scope

Use this template to provide a defensible competitive-pressure assessment that supports strategic choices. It should be updated as market structure changes and as strategic priorities evolve.

## Document Metadata Schema

```yaml
---
id: PFO-{market-context}
title: "Porter's Five Forces — {Industry or Market Segment}"
type: PFO
status: Draft|Review|Accepted|Deprecated
attribution_required: true
source_frameworks:
  - "Michael E. Porter - Five Forces Framework"
version: 1.0.0
owner: Strategy-Team|Market-Analyst
stakeholders: [leadership-team, strategy-team, market-team, sales-team]
domain: market
priority: High|Medium|Low
scope: competitive-structure
horizon: current|near-term
visibility: internal

depends_on: [CMP-*,TRN-*]
enables: [POS-*,STR-*]

force_analysis: Supplier Power|Buyer Power|Threat of New Entrants|Threat of Substitutes|Rivalry Intensity
assessment_horizon: Current|12Months|24Months|36Months

success_criteria:
  - "Forces are consistently measured and revisited on schedule"
  - "Strategic recommendations are tied to force changes"
  - "Findings are reflected in positioning and portfolio decisions"

assumptions:
  - "Competitor and market intelligence is current"
  - "Strategic implications are mapped to concrete decisions"
  - "Cross-functional owners are engaged"

constraints: [Data availability, intelligence reliability, market volatility]
standards: [Competitive analysis and strategy standards]
review_cycle: quarterly
---
```

## Content Structure Template

# Porter's Five Forces — &#123;Industry or Segment&#125;

## Executive Summary

**Market Segment:** &#123;Primary segment under analysis&#125;<br>
**Review Date:** &#123;YYYY-MM-DD&#125;<br>
**Primary Dynamic:** &#123;Most material competitive pressure&#125;<br>
**Recommended Positioning Impact:** &#123;Key strategic implication&#125;

## Competitive Force Assessment

### Supplier Power
- **Concentration:** &#123;High/Medium/Low&#125;
- **Switching Costs:** &#123;High/Medium/Low&#125;
- **Differentiation of Inputs:** &#123;High/Medium/Low&#125;
- **Mitigation:** &#123;How supplier power is reduced&#125;

### Buyer Power
- **Concentration:** &#123;High/Medium/Low&#125;
- **Price Sensitivity:** &#123;High/Medium/Low&#125;
- **Information Access:** &#123;Low/Medium/High&#125;
- **Mitigation:** &#123;How buyer power is reduced&#125;

### Threat of New Entrants
- **Barriers:** &#123;High/Medium/Low&#125;
- **Capital Requirements:** &#123;High/Medium/Low&#125;
- **Regulatory Barriers:** &#123;High/Medium/Low&#125;
- **Mitigation:** &#123;Defensive/Offensive responses&#125;

### Threat of Substitutes
- **Substitute Strength:** &#123;High/Medium/Low&#125;
- **Switching Attractiveness:** &#123;Low/Medium/High&#125;
- **Price/Quality Differential:** &#123;Favorable/Unfavorable&#125;
- **Mitigation:** &#123;How value and lock-in are protected&#125;

### Competitive Rivalry
- **Intensity:** &#123;High/Medium/Low&#125;
- **Competitor Concentration:** &#123;High/Medium/Low&#125;
- **Cycle Length:** &#123;Short/Medium/Long&#125;
- **Mitigation:** &#123;Differentiation, defensibility, speed&#125;

## Heat Map and Priorities

1. **Force**
2. **Current Score (1-5)**
3. **Trend (Worsening | Stable | Improving)**
4. **Strategic Countermeasure**

## Strategic Linkage

- `POS` impact: how competitive forces shift positioning choices.
- `STR` impact: where this analysis changes capability and go-to-market priorities.
- `OPP` impact: how market opportunities are screened against force pressure.

## Governance

- **Owner Review Cadence:** &#123;Monthly/Quarterly&#125;
- **Approval:** &#123;Leadership/Strategy Council&#125;
- **Escalation:** &#123;What changes trigger strategy review&#125;

## Quality Standard

- Force assessments have evidence-backed rationale.
- Force scores are consistent across reporting periods.
- Strategic implications are explicit and actionable.


---

## Related Documents

- [Back to Document Types](/docs/document-types)
- [Domain: Market & Environment](/docs/domains/market-environment)
- [Full Specification](/spec/v1-0-0)
