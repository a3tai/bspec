---
title: "STA: Stakeholder Map"
description: "BSpec STA document type specification"
---

# STA: Stakeholder Map

::: tip Document Type
**Code**: STA<br>
**Name**: Stakeholder Map<br>
**Domain**: Risk & Governance
:::

## Abstract

This specification defines the Stakeholder Map document type within the BSpec 1.0 Universal Business Specification Standard. It identifies influence and impact stakeholders, and records active engagement requirements.

## Purpose and Scope

Stakeholder Map establishes who can materially influence outcomes, operations, compliance posture, funding outcomes, and execution feasibility. It is used for strategy, governance, ethics, and risk communication planning.

## Document Metadata Schema

```yaml
---
id: STA-{stakeholder-set}
title: "Stakeholder Map — [Stakeholder Set Name]"
type: STA
status: Draft|Review|Approved|Active|Deprecated
attribution_required: false
version: 1.0.0
owner: Governance-Lead|Leadership-Team
stakeholders: [governance-team, legal, leadership-team, investor-relations]
domain: risk-governance
priority: High|Medium|Low
horizon: strategic
visibility: internal

depends_on: [PUR-*,VAL-*]
enables: [GOV-*,COM-*,ETH-*,LEG-*]

stakeholder_groups: [internal, external, regulatory, customer, partner, investor]
influence_scale: low|medium|high|critical
engagement_rhythm: quarterly|monthly|ad-hoc
review_cycle: annual
---
```

## Structure

- **Stakeholder Inventory:** Named stakeholders and accountable owners.
- **Interest and Influence Matrix:** Power-interest and risk-exposure mapping.
- **Engagement Model:** Who is consulted, informed, governed, and escalated.
- **Grievance and feedback channels:** Formal and informal communication channels.

## Validation Checklist

- [ ] Critical stakeholders and owners are all represented.
- [ ] Influence and urgency are explicitly rated.
- [ ] Engagement cadence and escalation paths are documented.
- [ ] Cross-domain linkages to risk, compliance, strategy, and ethics are explicit.


---

## Related Documents

- [Back to Document Types](/docs/document-types)
- [Domain: Risk & Governance](/docs/domains/risk-governance)
- [Full Specification](/spec/v1-0-0)
