# PFO: Porter's Five Forces Document Type Specification

**Document Type Code:** PFO
**Document Type Name:** Porter's Five Forces
**Domain:** Market & Environment
**Version:** 1.0.0
**Status:** Draft
**Last Updated:** 2025-09-28

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

# Porter's Five Forces — {Industry or Segment}

## Executive Summary

**Market Segment:** {Primary segment under analysis}<br>
**Review Date:** {YYYY-MM-DD}<br>
**Primary Dynamic:** {Most material competitive pressure}<br>
**Recommended Positioning Impact:** {Key strategic implication}

## Competitive Force Assessment

### Supplier Power
- **Concentration:** {High/Medium/Low}
- **Switching Costs:** {High/Medium/Low}
- **Differentiation of Inputs:** {High/Medium/Low}
- **Mitigation:** {How supplier power is reduced}

### Buyer Power
- **Concentration:** {High/Medium/Low}
- **Price Sensitivity:** {High/Medium/Low}
- **Information Access:** {Low/Medium/High}
- **Mitigation:** {How buyer power is reduced}

### Threat of New Entrants
- **Barriers:** {High/Medium/Low}
- **Capital Requirements:** {High/Medium/Low}
- **Regulatory Barriers:** {High/Medium/Low}
- **Mitigation:** {Defensive/Offensive responses}

### Threat of Substitutes
- **Substitute Strength:** {High/Medium/Low}
- **Switching Attractiveness:** {Low/Medium/High}
- **Price/Quality Differential:** {Favorable/Unfavorable}
- **Mitigation:** {How value and lock-in are protected}

### Competitive Rivalry
- **Intensity:** {High/Medium/Low}
- **Competitor Concentration:** {High/Medium/Low}
- **Cycle Length:** {Short/Medium/Long}
- **Mitigation:** {Differentiation, defensibility, speed}

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

- **Owner Review Cadence:** {Monthly/Quarterly}
- **Approval:** {Leadership/Strategy Council}
- **Escalation:** {What changes trigger strategy review}

## Quality Standard

- Force assessments have evidence-backed rationale.
- Force scores are consistent across reporting periods.
- Strategic implications are explicit and actionable.
