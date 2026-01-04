---
title: "FEE: Feedback"
description: "BSpec FEE document type specification"
---

# FEE: Feedback

**Domain**: Customer Value

The Feedback document captures, analyzes, and manages customer input, reviews, and satisfaction data. It provides systematic approach to collecting, processing, and acting on customer feedback across all touchpoints.

## Metadata Schema

```yaml
---
id: FEE-{identifier}
type: FEE
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
| Bronze | Feedback framework with collection strategy and methods, Basic feedback analysis with quantitative and qualitative insights, Key findings with positive themes and improvement opportunities |
| Silver | Comprehensive feedback collection results with demographics and channel analysis, Detailed feedback analysis including sentiment analysis and segmentation, Action planning with immediate, short-term, and long-term initiatives |
| Gold | Full documentation with validation evidence |

## Relationships

**Depends on**: CJM, SUP

**Enables**: REQ, PRD, SVC

---

- [Back to Document Types](/docs/document-types)
- [Domain: Customer Value](/docs/domains/customer-value)
