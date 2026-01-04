---
title: "SEC: Security"
description: "BSpec SEC document type specification"
---

# SEC: Security

**Domain**: Technology & Data

The Security document defines systematic approaches to designing, implementing, and managing security controls that protect business assets through comprehensive risk management, compliance adherence, and threat mitigation. It establishes security frameworks that ensure confidentiality, integrity, and availability of organizational resources.

## Metadata Schema

```yaml
---
id: SEC-{identifier}
type: SEC
status: Draft|Review|Accepted|Deprecated
version: 1.0.0
domain: technology-data
priority: critical|high|medium|low

depends_on: []
enables: []
---
```

## Quality Levels

| Level | Requirements |
|-------|--------------|
| Bronze | Basic security controls with essential protections, Simple identity and access management, Basic incident response capabilities |
| Silver | Comprehensive security framework with defense in depth, Advanced identity and access management with automation, Structured security operations with monitoring and response |
| Gold | Full documentation with validation evidence |

## Relationships

**Depends on**: ARC, INF, POL

**Enables**: PER, QUA, SLA

---

- [Back to Document Types](/docs/document-types)
- [Domain: Technology & Data](/docs/domains/technology-data)
