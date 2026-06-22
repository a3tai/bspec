# SWO: SWOT Synthesis Document Type Specification

**Document Type Code:** SWO
**Document Type Name:** SWOT Synthesis
**Domain:** Market & Environment
**Version:** 1.0.0
**Status:** Draft
**Last Updated:** 2026-05-27

## Abstract

This specification defines the SWOT Synthesis document type within the BSpec 1.0 Universal Business Specification Standard. It synthesizes internal and external analyses into strategy-forming artifacts.

## Purpose and Scope

SWOT Synthesis converts Opportunity and Threat inputs into consolidated strategic implications and option sets, and connects them to strategy and action priorities.

## Document Metadata Schema

```yaml
---
id: SWO-{market-context}
title: "SWOT Synthesis — [Market Context]"
type: SWO
status: Draft|Review|Approved|Active|Deprecated
attribution_required: false
version: 1.0.0
owner: Strategy-Lead|Strategy-Team
stakeholders: [leadership-team, strategy-team, marketing-team]
domain: market
priority: Medium|High|Critical
horizon: strategic
visibility: internal

depends_on: [OPP-*,THR-*,MKT-*]
enables: [STR-*,POS-*,OPP-*]

analysis_horizon: current|mid-term|long-term
method_notes: [qualitative|quantitative|mixed]
review_cycle: quarterly
---
```

## Framework

- **Strengths:** Internal capabilities and advantages.
- **Weaknesses:** Internal limitations requiring mitigation.
- **Opportunities:** External openings and timing windows.
- **Threats:** External risks and downside scenarios.

## Strategic Synthesis

- **Option Generation:** Strategic alternatives from strengths/opportunities and defense responses.
- **Prioritization:** Trade-offs by impact and feasibility.
- **Action Plan:** Priority initiatives and ownership.
- **Feedback Loop:** Update with new market and operational data.

## Validation Checklist

- [ ] Inputs are explicitly linked to OPP and Threat analysis.
- [ ] Strategic implications are converted to concrete decisions.
- [ ] Ownership and review cycle are specified.
- [ ] Assumptions and confidence are recorded.
