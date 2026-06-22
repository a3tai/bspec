---
title: "EOL: End-of-Life and Retirement"
description: "BSpec EOL document type specification"
---

# EOL: End-of-Life and Retirement

::: tip Document Type
**Code**: EOL<br>
**Name**: End-of-Life and Retirement<br>
**Domain**: Product & Service
:::

## Abstract

This specification defines the End-of-Life and Retirement document type within the BSpec 1.0 Universal Business Specification Standard. It governs product/service deprecation, support transitions, and migration obligations.

## Purpose and Scope

End-of-Life and Retirement manages planned retirement of products/services, including customer communication, migration strategy, and sunset operations across support, compliance, and security.

## Document Metadata Schema

```yaml
---
id: EOL-{product-service}
title: "End-of-Life Plan — [Product or Service Name]"
type: EOL
status: Draft|Review|Approved|Active|Deprecated
attribution_required: false
version: 1.0.0
owner: Product-Management|Support|Engineering
stakeholders: [product-team, customer-success, support, security, legal]
domain: product-service
priority: High|Medium|Low
horizon: tactical
visibility: internal

depends_on: [ROD-*,PRD-*,SVC-*]
enables: [SUP-*,REQ-*,SEC-*]

deprecation_type: minor|major|platform|category
sunset_horizon: [60|90|120|custom] # days
migration_paths: [manual|assisted|automated|parallel]
support_tiers: [L1|L2|L3]
review_cycle: quarterly
---
```

## Lifecycle Design

- **Lifecycle Trigger:** Criteria for initiating retirement planning.
- **Impact Surface:** Customers, contracts, integrations, and reporting obligations.
- **Migration Plan:** Transition milestones, migration support, and fallback paths.
- **Communication Plan:** Multi-channel communication by audience.

## Operational Controls

- **Parallel Support:** Overlap period and deactivation rules.
- **Security and Data Retention:** Retention, deletion, and compliance alignment.
- **Support Exit Criteria:** Known issue handling and closure rules.
- **Post-EOL Review:** Lessons and template reuse.

## Validation Checklist

- [ ] End-of-life decision rationale is explicitly documented.
- [ ] Customer impact and migration commitments are quantified.
- [ ] Support and security handoffs are complete.
- [ ] Closure conditions are defined and executable.


---

## Related Documents

- [Back to Document Types](/docs/document-types)
- [Domain: Product & Service](/docs/domains/product-service)
- [Full Specification](/spec/v1-0-0)
