# VCH: Value Chain Analysis Document Type Specification

**Document Type Code:** VCH
**Document Type Name:** Value Chain Analysis
**Domain:** Market & Environment
**Version:** 1.0.0
**Status:** Draft
**Last Updated:** 2025-09-28

## Abstract

This specification defines the Value Chain Analysis document type within the BSpec 1.0 Universal Business Specification Standard. It captures primary and support activities, identifies value leak points, and links chain-level insights into strategic priorities.

## Framework and Attribution

This document references **Porter Value Chain analysis** methods for structuring internal and ecosystem value flows.

## Purpose and Scope

Use this template to produce a repeatable map of value-driving activities, dependencies, and optimization opportunities across inbound logistics, operations, and customer-facing delivery.

## Document Metadata Schema

```yaml
---
id: VCH-{value-chain-area}
title: "Value Chain Analysis — {Business or Product Line}"
type: VCH
status: Draft|Review|Accepted|Deprecated
attribution_required: true
source_frameworks:
  - "Michael E. Porter - Value Chain"
version: 1.0.0
owner: Strategy-Team|Operations-Team
stakeholders: [executives, operations-team, product-team, finance-team]
domain: market
priority: High|Medium|Low
scope: value-chain-architecture
horizon: current|strategic
visibility: internal

depends_on: [STR-*,KAC-*]
enables: [VST-*,CUS-*,POS-*]

chain_scope: End-to-end|Customer-facing|Back-office|Partner-led
analysis_horizon: Current|12Months|24Months
maturity_level: Emerging|Defined|Managed|Optimizing

success_criteria:
  - "Value drivers and bottlenecks are explicitly mapped"
  - "Activity-level opportunities are ranked and owned"
  - "Delivery quality and margin outcomes improve over cycles"

assumptions:
  - "Activity boundaries are consistently defined"
  - "Cost and performance inputs are measurable"
  - "Owners align on chain-level priorities"

constraints: [Data quality, system fragmentation, change readiness]
standards: [Value chain and operational analysis standards]
review_cycle: quarterly
---
```

## Content Structure Template

# Value Chain Analysis — {Business or Product Line}

## Executive Summary

**Value Chain Scope:** {Primary chain being analyzed}<br>
**Primary Bottleneck:** {Most critical constraint}<br>
**Top Value Lever:** {Highest impact opportunity}

## Activity Mapping

### Primary Activities
- Inbound logistics
- Operations
- Outbound logistics
- Marketing and sales
- Service

### Support Activities
- Firm infrastructure
- Human resources
- Technology development
- Procurement

## Value Measurement

- **Cost to Deliver** by activity
- **Lead time** by activity
- **Quality defects / rework rates**
- **Customer outcome correlation** by stage

## Improvement Opportunities

For each opportunity, document:

- **Activity:** {Which value step}
- **Constraint:** {Observed weakness}
- **Expected Impact:** {Margin, speed, quality, risk}
- **Owner:** {Team owner}
- **Evidence:** {Source data}
- **Timeline:** {Implementation window}

## Strategic Linkage

- `VST` impact: links activity changes to end-to-end customer value delivery.
- `FIN` impact: supports cost and margin decomposition.
- `STR` impact: informs defensibility and capability investments.

## Quality Standard

- Activity map has consistent boundaries across reporting periods.
- Improvement opportunities include owners, timelines, and expected outcomes.
- Dependencies are explicit and auditable.
