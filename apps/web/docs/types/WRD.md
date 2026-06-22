---
title: "WRD: Wardley Mapping"
description: "BSpec WRD document type specification"
---

# WRD: Wardley Mapping

::: tip Document Type
**Code**: WRD<br>
**Name**: Wardley Mapping<br>
**Domain**: Market & Environment
:::

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

# Wardley Map — &#123;Business or Capability&#125;

## Executive Summary

**Map Focus:** &#123;User need chain or capability&#125;<br>
**Current Evolution:** &#123;Observed state&#125;<br>
**Decisions Enabled:** &#123;Primary decisions&#125;

## Strategic Context

- **Users:** &#123;Primary users impacted by mapped capabilities&#125;
- **Value Chain Nodes:** &#123;Core components under map&#125;
- **Competitive View:** &#123;Current competitive positioning against mapped components&#125;
- **Regulatory/Infrastructure Constraints:** &#123;Constraint notes&#125;

## Evolution Analysis

For each major component, document:

- **Component:** &#123;Node name&#125;
- **Evolution stage:** &#123;Genesis/Custom/Product/Commodity&#125;
- **Movement direction:** &#123;Improving / Evolving / Declining&#125;
- **Signal quality:** &#123;High/Medium/Low&#125;
- **Action:** &#123;Invest / Build / Buy / Partner / Abandon&#125;

## Strategic Actions

- **Immediate Actions:** &#123;Actions within current quarter&#125;
- **Mid-term Actions:** &#123;Actions in 6-12 months&#125;
- **Option Management:** &#123;What options to preserve now&#125;
- **Risk Controls:** &#123;Risks if execution is delayed&#125;

## Link to Strategy

- `STR` impact: updates positioning and sequence decisions.
- `KAC` / `CAP` impact: aligns capability investments.
- `VST` impact: prioritizes which value functions to modernize.

## Quality Standard

- Value chain components are mapped consistently.
- Evolution states are revised with new evidence.
- Decisions include ownership and timing.


---

## Related Documents

- [Back to Document Types](/docs/document-types)
- [Domain: Market & Environment](/docs/domains/market-environment)
- [Full Specification](/spec/v1-0-0)
