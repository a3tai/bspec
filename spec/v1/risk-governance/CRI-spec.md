# CRI: Crisis Management Document Type Specification

**Document Type Code:** CRI
**Document Type Name:** Crisis Management
**Domain:** Risk & Governance
**Version:** 1.0.0
**Status:** Draft
**Last Updated:** 2026-05-27

## Abstract

This specification defines the Crisis Management document type within the BSpec 1.0 Universal Business Specification Standard. It addresses acute, high-impact events that threaten operations, reputation, or continuity.

## Purpose and Scope

Crisis Management defines severity-classified response playbooks, executive communication, and command structure for high-impact disruptions that go beyond routine incidents.

## Document Metadata Schema

```yaml
---
id: CRI-{crisis-program}
title: "Crisis Management — [Program Name]"
type: CRI
status: Draft|Review|Approved|Active|Deprecated
attribution_required: false
version: 1.0.0
owner: Crisis-Lead|COO|Leadership-Office
stakeholders: [leadership-team, operations, legal, communications, finance]
domain: risk-governance
priority: Critical|High|Medium|Low
horizon: tactical
visibility: internal

depends_on: [INC-*,OPS-*,RSK-*]
enables: [BCR-*]

severity_levels: [observe|warn|major|critical]
activation_thresholds: [revenue, safety, regulatory, reputational]
incident_command: [incident-commander, communications-lead, legal-lead, restoration-lead]
review_cycle: annual
---
```

## Crisis Framework

- **Detection and Escalation:** Trigger conditions and activation flow.
- **Command Structure:** Roles, decision authority, and communication cadence.
- **Mitigation Actions:** Containment, mitigation, and stabilization priorities.
- **Recovery and Reconstitution:** Return-to-normal plan and post-incident review.

## Validation Checklist

- [ ] Severity thresholds and activation criteria are explicit.
- [ ] Crisis command chain and authority are documented.
- [ ] External and internal communication templates are defined.
- [ ] Closure and lessons-learned process is documented.
