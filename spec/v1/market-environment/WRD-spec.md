# WRD: Wardley Mapping Document Type Specification

**Document Type Code:** WRD
**Document Type Name:** Wardley Mapping
**Domain:** Market & Environment
**Version:** 1.0.0
**Status:** Draft
**Last Updated:** 2025-09-28

## Abstract

This specification defines the Wardley Mapping document type within the BSpec 1.0 Universal Business Specification Standard. It is intended to map value chains against competitive/market evolution to improve strategy sequencing, investing timing, and execution planning.

## Framework and Attribution

This document references the **Wardley Mapping** method and must retain source attribution where method-specific outputs are used.

## Purpose and Scope

Use this template to represent value components, user needs, and evolutionary movement in a way that informs investment and operational decisions.

## Document Metadata Schema

```yaml
---
id: WRD-{strategic-context}
title: "Wardley Map — {Business or Capability}"
type: WRD
status: Draft|Review|Accepted|Deprecated
attribution_required: true
source_frameworks:
  - "Simon Wardley - Wardley Mapping"
version: 1.0.0
owner: Strategy-Team|Product-Team
stakeholders: [executives, strategy-team, product-team, engineering]
domain: market
priority: High|Medium|Low
scope: strategy-mapping
horizon: current|future
visibility: internal

depends_on: [CMP-*,MKT-*]
enables: [STR-*,POS-*,VST-*]

mapping_scope: Strategy|Product|Capability
evolution_focus: Genesis|Custom|Product|Commodity
mapping_horizon: 6-12Months|12-36Months|36+Months

success_criteria:
  - "Map outputs are tied to explicit strategic decisions"
  - "Evolution states are updated with observed signals"
  - "Decisions include sequencing and risk-aware timing"

assumptions:
  - "User need chain is identifiable"
  - "Market evolution signals are being monitored"
  - "Execution teams can act on map outputs"

constraints: [Method skill requirements, model staleness]
standards: [Strategy mapping and scenario analysis]
review_cycle: quarterly
---
```

## Content Structure Template

# Wardley Map — {Business or Capability}

## Executive Summary

**Map Focus:** {User need chain or capability}<br>
**Current Evolution:** {Observed state}<br>
**Decisions Enabled:** {Primary decisions}

## Strategic Context

- **Users:** {Primary users impacted by mapped capabilities}
- **Value Chain Nodes:** {Core components under map}
- **Competitive View:** {Current competitive positioning against mapped components}
- **Regulatory/Infrastructure Constraints:** {Constraint notes}

## Evolution Analysis

For each major component, document:

- **Component:** {Node name}
- **Evolution stage:** {Genesis/Custom/Product/Commodity}
- **Movement direction:** {Improving / Evolving / Declining}
- **Signal quality:** {High/Medium/Low}
- **Action:** {Invest / Build / Buy / Partner / Abandon}

## Strategic Actions

- **Immediate Actions:** {Actions within current quarter}
- **Mid-term Actions:** {Actions in 6-12 months}
- **Option Management:** {What options to preserve now}
- **Risk Controls:** {Risks if execution is delayed}

## Link to Strategy

- `STR` impact: updates positioning and sequence decisions.
- `KAC` / `CAP` impact: aligns capability investments.
- `VST` impact: prioritizes which value functions to modernize.

## Quality Standard

- Value chain components are mapped consistently.
- Evolution states are revised with new evidence.
- Decisions include ownership and timing.
