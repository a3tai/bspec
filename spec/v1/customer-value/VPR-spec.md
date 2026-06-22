# VPR: Value Proposition Document Type Specification

**Document Type Code:** VPR
**Document Type Name:** Value Proposition
**Domain:** Customer & Value
**Version:** 1.0.0
**Status:** Draft
**Last Updated:** 2025-09-28

## Abstract

This specification defines the Value Proposition document type within the BSpec 1.0 Universal Business Specification Standard. It establishes a first-class customer-value artifact for expressing the outcomes, pains, and gains combination that justifies buying a product, service, or solution.

## Framework and Attribution

VPR is the central synthesis layer for the Value Proposition Canvas family and should preserve attribution when structural language and workflow are derived from Strategyzer materials:

- Strategyzer / Value Proposition Canvas (CC BY-SA terms differ from reuse guidance depending on commercial packaging and software usage context).
- Use explicit licensing and attribution language wherever this artifact is published or redistributed.

## Purpose and Scope

The Value Proposition document converts customer insight into a clearly testable proposition statement. It includes pains, gains, gain creators, and gain alleviators and defines the boundary between what customers need and what the organization can credibly deliver.

## Document Metadata Schema

```yaml
---
id: VPR-{value-proposition-identifier}
title: "Value Proposition — [Value Proposition Name]"
type: VPR
status: Draft|Review|Accepted|Deprecated
attribution_required: true
source_frameworks:
  - "Strategyzer - Value Proposition Canvas"
version: 1.0.0
owner: Product-Strategy|Marketing-Team
stakeholders: [product-team, sales-team, marketing-team, leadership-team]
domain: customer
priority: High|Medium|Low
horizon: current
visibility: internal

depends_on: [PAI-*,GAI-*,JTB-*,PER-*]
enables: [POS-*,REV-*,PRI-*,BPO-*,MSG-*]

value_statement: [customer, outcome, and proof structure]
customer_jobs: [jobs addressed in value proposition]
pain_points: [primary pains]
gain_points: [primary gains]
key_benefits: [business and customer outcomes]
proof_sources: [data, test results, reference cases]

success_criteria:
  - "Value proposition is clearly differentiated and testable"
  - "Pain and gain assumptions are validated"
  - "Proposition is supported by evidence and outcomes"
  - "Execution alignment is visible across product and market artifacts"

assumptions:
  - "Customer outcomes are measurable and meaningful"
  - "Organization can reasonably deliver the value promised"
  - "Market fit is validated with real customer behavior"

constraints: [Market, channel, and operational constraints]
standards: [Customer value and market communication standards]
review_cycle: quarterly
---
```

## Content Structure Template

```markdown
# Value Proposition: [Value Proposition Name]

## Executive Summary
*Short statement of the customer outcome and why it matters now*

## Value Proposition
**Core Benefit:** [what outcome changes for the customer]
**Supporting Benefits:** [secondary and adjacent outcomes]
**Exclusions:** [what this proposition does not cover]

## Customer Segment
- Segment: [customer group]
- Primary Jobs: [job statements]
- Core Pains: [high-impact pain points]
- Desired Gains: [high-impact gains]

## Problem and Outcome Hypothesis
- [ ] Problem statement is specific and observable
- [ ] Outcome is measurable
- [ ] Current alternatives are mapped
- [ ] Positioning is defensible

## Value Proposition Design
### Pain Reduction
- Identify the top pains reduced by this proposition
- State measurable deltas (time, effort, cost, risk)
- Define target customer expectations

### Gain Creation
- Define gains created or amplified
- State proof points and evidence sources
- Define who benefits and how they will measure success

## Delivery Fit
- Internal capability match
- Product/service scope alignment
- Delivery dependencies (team, systems, channels)
- Risk controls and downside cases

## Validation and Measurement
### Evidence Plan
- Test plans and success metrics
- Learning checkpoints and falsification criteria
- Decision triggers and revision conditions

### KPI Set
- Adoption and conversion proxy signals
- Retention and expansion indicators
- Advocacy and satisfaction signals

## Delivery and Positioning References
- **POS**: Positioning statement and competition framing
- **CUS/PRD/FEA**: Journey and feature grounding
- **REV/PRI**: Revenue and pricing alignment
- **MSG/CAM**: Messaging and campaign alignment
```

## Validation Checklist

- [ ] Value proposition statement is clear and differentiated
- [ ] Jobs, pains, and gains are mapped and consistent
- [ ] Claims are tied to evidence and measurable outcomes
- [ ] Dependencies and execution scope are complete
- [ ] Positioning and sales enablement references are explicit
