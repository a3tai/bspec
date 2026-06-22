# SAL: Sales Strategy Document Type Specification

**Document Type Code:** SAL
**Document Type Name:** Sales Strategy
**Domain:** Brand & Marketing
**Version:** 1.0.0
**Status:** Draft
**Last Updated:** 2025-09-28

## Abstract

This specification defines the Sales Strategy document type within the BSpec 1.0 Universal Business Specification Standard. It establishes a structured operating view of the sales process, from lead intake and qualification through close execution and account expansion.

## Purpose and Scope

The Sales Strategy document defines how the organization acquires, converts, and retains customers through repeatable sales motions. It establishes strategy, process design, and performance governance for sales planning, demand generation handoff, and commercial execution.

## Document Metadata Schema

```yaml
---
id: SAL-{sales-context}
title: "Sales Strategy — {Sales Context}"
type: SAL
status: Draft|Review|Approved|Active|Deprecated
attribution_required: false
version: 1.0.0
owner: Sales-Lead|Commercial-Operations
stakeholders: [sales-team, marketing-team, customer-success-team, finance-team]
domain: brand-marketing
priority: Critical|High|Medium|Low
horizon: tactical
visibility: internal

depends_on: [CJM-*,PER-*,POS-*,VPR-*]
enables: [REV-*,PRI-*,SUP-*,FEE-*,ROD-*]

sales_model: Pipeline|Enterprise|Mid-Market|SMB|Self-serve|Hybrid
funnel_stage: Awareness|Consideration|Evaluation|Decision|Onboarding|Expansion
channel_priority: [inbound, outbound, partner, referral, direct]
quota_cycle: Annual|Quarterly|Monthly
quota_type: Revenue|Units|New-logos|Retention
assumptions: [Sales team capacity, lead quality, pricing clarity]
review_cycle: quarterly
---
```

## Sales Motion

- **Sales Motion Type:** {How the sales process is executed by segment/customer type}
- **Segment Prioritization:** {Priority sequencing for target segments}
- **Account Strategy:** {Top-down vs bottom-up account planning}
- **Channel Model:** {Channel selection, handoff, and ownership}
- **Commercial Governance:** {Quota, forecasting, and accountability model}

## Pipeline Structure

### Lead Management
- **Lead Sources:** {Inbound, outbound, partner, ecosystem}
- **Qualification Standard:** {Lead scoring and qualification criteria}
- **Routing Rules:** {Lead handoff and ownership flow}
- **Disqualification Criteria:** {When and why opportunities are declined}

### Opportunity Management
- **Opportunity Stages:** {Defined sales stage transitions and entry/exit criteria}
- **Milestone Criteria:** {What advances opportunities through each stage}
- **Risk Triggers:** {Signals indicating pause, escalation, or rework}
- **Closure Requirements:** {Contract, legal, delivery, and onboarding readiness}

### Commercial Execution
- **Deal Design:** {Packaging, commitments, and delivery promises}
- **Pricing & Terms:** {Commercial terms governance and variance process}
- **Proposal Governance:** {Quality gates for proposals and SOWs}
- **Negotiation Playbook:** {Concession rules and approval thresholds}

## Sales Enablement Interfaces

Sales Strategy is the handoff anchor between marketing and execution:

- **Marketing Interface:** Campaign and messaging alignment with lead quality and conversion criteria.
- **Product Interface:** Requirements for onboarding, implementation capacity, and feature commitments.
- **Support Interface:** Handoff of expectations and escalation paths post-sale.
- **Finance Interface:** Forecast and revenue recognition assumptions tied to contract structure.

## Performance Management

### Metrics and KPIs
- **Pipeline Health:** {MQL → SQL conversion, stage progression}
- **Win Quality:** {Win rate by segment/channel/segment and cycle length}
- **Revenue Velocity:** {Velocity-to-close and booking conversion}
- **Forecast Accuracy:** {Forecast-to-actual variance and reasons}

### Coaching and Improvement
- **Field Coaching:** {Cadence, shadowing, and playbook updates}
- **Failure Analysis:** {Lost deal analysis and pattern fixes}
- **Capacity Planning:** {Headcount and ramp planning}
- **Expansion Strategy:** {Upsell, cross-sell, and retention sequencing}

## Governance

- **Sales Calendar:** {Cadence for planning, forecasting, and reviews}
- **Accountability Model:** {Roles, responsibilities, and handoff ownership}
- **Systems and Tooling:** {CRM, sequencing, and reporting stack requirements}
- **Escalation Protocol:** {Commercial exceptions and risk escalations}

## Relationship Guidelines

### Typical Dependencies
- **POS (Positioning)**: Sales messaging and differentiation framing
- **CJM (Customer Journey)**: Journey-informed touchpoints and conversion moments
- **VPR (Value Proposition)**: Sales narrative and value proof

### Typical Enablements
- **REV (Revenue)**: Converts sales strategy into revenue model assumptions
- **PRI (Pricing)**: Aligns commercial terms and discount governance
- **SUP (Support)**: Supports onboarding and customer success transfer
- **REP (Reporting)**: Drives sales planning and execution reporting

## Validation Checklist

- [ ] Sales stages, qualifications, and handoff rules are explicitly defined
- [ ] Channel strategy and segment focus are codified
- [ ] Pricing, approval, and concession governance is explicit
- [ ] Pipeline performance metrics are owned and reviewed
- [ ] Downstream dependencies and enablements are documented
