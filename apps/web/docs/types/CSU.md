---
title: "CSU: Customer Success"
description: "BSpec CSU document type specification"
---

# CSU: Customer Success

::: tip Document Type
**Code**: CSU<br>
**Name**: Customer Success<br>
**Domain**: Customer & Value
:::

## Abstract

This specification defines the Customer Success document type within the BSpec 1.0 Universal Business Specification Standard. It establishes a proactive model for onboarding, adoption, retention, and expansion management after sale.

## Purpose and Scope

Customer Success defines how the organization drives customer value realization and long-term retention after purchase. It is responsible for proactive lifecycle outcomes: onboarding, adoption, expansion, risk mitigation, and customer advocacy.

## Document Metadata Schema

```yaml
---
id: CSU-{success-program}
title: "Customer Success Plan — [Program Name]"
type: CSU
status: Draft|Review|Approved|Active|Deprecated
attribution_required: false
version: 1.0.0
owner: Customer-Success-Lead|Customer-Support|Sales-Operations
stakeholders: [customer-success-team, support, sales, product, finance]
domain: customer-value
priority: High|Medium|Low
scope: lifecycle
horizon: tactical
visibility: internal

depends_on: [PER-*,CJM-*,FEE-*,SUP-*,PRI-*]
enables: [SUP-*,FEE-*,MET-*,PRD-*]

success_goals: [onboarding, adoption, retention, expansion]
health_model: usage|value|engagement|outcome
customer_segments: [segment identifiers]
cohorts: [new, expansion, churn-risk, strategic]
review_cycle: monthly
---
```

## Framework and Relationships

- **Depends on Customer Insights:** Persona and journey context from `PER-*`, `CJM-*`, and feedback programs.
- **Coordinates With:** Support (`SUP-*`) for service operations, Product (`PRD-*`) for delivery commitments, and Pricing (`PRI-*`) for expansion mechanics.

## Content Structure Template

### Customer Lifecycle Model

- **Onboarding**
  - Enrollment criteria and acceptance process
  - Setup and enablement milestones
  - Handover checkpoints from Sales
- **Adoption**
  - Feature adoption targets
  - Training and enablement assets
  - User competency milestones
- **Retention**
  - Usage/engagement risk detection
  - Escalation and rescue playbooks
  - Retention commitments and governance
- **Expansion**
  - Expansion triggers and qualification
  - Cross-sell/up-sell opportunity tracking
  - Value-based expansion governance

### Operational Model

- **Segment Programs:** Tailored success motion per persona and segment.
- **Health Framework:** Leading indicators, lagging outcomes, and trigger thresholds.
- **Success Plans:** Named outcomes, owners, cadence, and confidence levels.
- **Escalation and Governance:** Clear ownership for blockers, exceptions, and renewal risks.

### Metrics and Signals

- **North Star Health Signals:** Adoption velocity, time-to-value, renewal intent.
- **Experience Signals:** Product usage consistency and service interaction quality.
- **Business Signals:** Retention, net expansion, churn probability, referral propensity.

## Validation Checklist

- [ ] Onboarding pathways are defined per customer segment.
- [ ] Expansion criteria are based on measurable outcomes.
- [ ] Customer risk detection and rescue controls are explicit.
- [ ] Dependencies to support and product commitments are traceable.
- [ ] Health metrics and governance cadence are defined.


---

## Related Documents

- [Back to Document Types](/docs/document-types)
- [Domain: Customer & Value](/docs/domains/customer-value)
- [Full Specification](/spec/v1-0-0)
