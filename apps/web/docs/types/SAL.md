---
title: "SAL: Sales Strategy"
description: "BSpec SAL document type specification"
---

# SAL: Sales Strategy

::: tip Document Type
**Code**: SAL<br>
**Name**: Sales Strategy<br>
**Domain**: Brand & Marketing
:::

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

- **Sales Motion Type:** &#123;How the sales process is executed by segment/customer type&#125;
- **Segment Prioritization:** &#123;Priority sequencing for target segments&#125;
- **Account Strategy:** &#123;Top-down vs bottom-up account planning&#125;
- **Channel Model:** &#123;Channel selection, handoff, and ownership&#125;
- **Commercial Governance:** &#123;Quota, forecasting, and accountability model&#125;

## Pipeline Structure

### Lead Management
- **Lead Sources:** &#123;Inbound, outbound, partner, ecosystem&#125;
- **Qualification Standard:** &#123;Lead scoring and qualification criteria&#125;
- **Routing Rules:** &#123;Lead handoff and ownership flow&#125;
- **Disqualification Criteria:** &#123;When and why opportunities are declined&#125;

### Opportunity Management
- **Opportunity Stages:** &#123;Defined sales stage transitions and entry/exit criteria&#125;
- **Milestone Criteria:** &#123;What advances opportunities through each stage&#125;
- **Risk Triggers:** &#123;Signals indicating pause, escalation, or rework&#125;
- **Closure Requirements:** &#123;Contract, legal, delivery, and onboarding readiness&#125;

### Commercial Execution
- **Deal Design:** &#123;Packaging, commitments, and delivery promises&#125;
- **Pricing & Terms:** &#123;Commercial terms governance and variance process&#125;
- **Proposal Governance:** &#123;Quality gates for proposals and SOWs&#125;
- **Negotiation Playbook:** &#123;Concession rules and approval thresholds&#125;

## Sales Enablement Interfaces

Sales Strategy is the handoff anchor between marketing and execution:

- **Marketing Interface:** Campaign and messaging alignment with lead quality and conversion criteria.
- **Product Interface:** Requirements for onboarding, implementation capacity, and feature commitments.
- **Support Interface:** Handoff of expectations and escalation paths post-sale.
- **Finance Interface:** Forecast and revenue recognition assumptions tied to contract structure.

## Performance Management

### Metrics and KPIs
- **Pipeline Health:** &#123;MQL → SQL conversion, stage progression&#125;
- **Win Quality:** &#123;Win rate by segment/channel/segment and cycle length&#125;
- **Revenue Velocity:** &#123;Velocity-to-close and booking conversion&#125;
- **Forecast Accuracy:** &#123;Forecast-to-actual variance and reasons&#125;

### Coaching and Improvement
- **Field Coaching:** &#123;Cadence, shadowing, and playbook updates&#125;
- **Failure Analysis:** &#123;Lost deal analysis and pattern fixes&#125;
- **Capacity Planning:** &#123;Headcount and ramp planning&#125;
- **Expansion Strategy:** &#123;Upsell, cross-sell, and retention sequencing&#125;

## Governance

- **Sales Calendar:** &#123;Cadence for planning, forecasting, and reviews&#125;
- **Accountability Model:** &#123;Roles, responsibilities, and handoff ownership&#125;
- **Systems and Tooling:** &#123;CRM, sequencing, and reporting stack requirements&#125;
- **Escalation Protocol:** &#123;Commercial exceptions and risk escalations&#125;

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


---

## Related Documents

- [Back to Document Types](/docs/document-types)
- [Domain: Brand & Marketing](/docs/domains/brand-marketing)
- [Full Specification](/spec/v1-0-0)
