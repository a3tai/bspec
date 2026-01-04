---
title: "GAI: Gains"
description: "BSpec GAI document type specification"
---

# GAI: Gains

**Domain**: Customer Value

The Gains document identifies and analyzes the positive outcomes, benefits, and value that customers achieve or seek. It captures the upside potential that motivates customer behavior and creates opportunities for value delivery.

## Metadata Schema

```yaml
---
id: GAI-{identifier}
type: GAI
status: Draft|Review|Accepted|Deprecated
version: 1.0.0
domain: customer-value
priority: critical|high|medium|low

depends_on: []
enables: []
---
```

## Quality Levels

| Level | Requirements |
|-------|--------------|
| Bronze | Gain statement clearly defined with core benefit and value, Gain categories (functional, emotional, financial, social) identified, Basic customer context and gain recipients documented |
| Silver | Comprehensive gain analysis with significance, frequency, and magnitude, Detailed value creation mechanisms and gain delivery process, Customer value perception with recognition and communication patterns |
| Gold | Full documentation with validation evidence |

## Relationships

**Depends on**: PAI, JTB, PER

**Enables**: VAL, REV, POS

---

- [Back to Document Types](/docs/document-types)
- [Domain: Customer Value](/docs/domains/customer-value)
