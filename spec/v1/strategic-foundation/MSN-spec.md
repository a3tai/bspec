# MSN: Mission Document Type Specification

**Document Type Code:** MSN
**Document Type Name:** Mission Statement
**Domain:** Strategic Foundation
**Version:** 1.0.0
**Status:** Draft
**Last Updated:** 2025-09-28

## Purpose and Scope

The Mission document defines why the organization exists and its core purpose in the world. It articulates the fundamental reason the organization was created and what it aims to accomplish for its stakeholders.

## Document Metadata Schema

```yaml
---
id: MSN-{organization-identifier}
title: "Mission Statement"
type: MSN
status: Draft|Review|Accepted|Deprecated
version: 1.0.0
owner: CEO|Founder
stakeholders: [all-employees, board, investors]
domain: strategic
priority: critical
scope: global
horizon: long
visibility: public

# Mission documents typically enable other strategic documents
enables: [VSN-*, VAL-*, STR-*]

# Mission rarely depends on other documents (foundational)
depends_on: []

success_criteria:
  - "All employees can recite and explain the mission"
  - "Mission guides major strategic decisions"
  - "Stakeholders understand organizational purpose"
  - "Mission differentiates from competitors"

assumptions:
  - "Organization has identified core beneficiaries"
  - "Founding team alignment on fundamental purpose"
  - "Purpose is sustainable and meaningful"

review_cycle: annually
---
```

## Content Structure Template

```markdown
# Mission Statement

## Overview
*Brief statement of the organization's fundamental purpose and reason for existence. This should be 1-2 sentences that clearly communicate why the organization exists.*

## Mission Statement
*The formal mission statement (typically 1-3 sentences that can be memorized)*

## Core Purpose
*Detailed explanation of the organization's reason for being*

### Primary Beneficiaries
- Who we serve
- How we create value for them
- Why they matter to our purpose

### Fundamental Value Creation
- What value we create in the world
- How we make things better
- What would be lost if we didn't exist

### Scope and Boundaries
- What we do and don't do
- Where we operate
- Limits of our mission

## Mission Evolution
*How the mission has evolved and why*

## Living the Mission
*How the mission translates into daily operations and decisions*

## Validation
*Evidence that the mission is being fulfilled*
```

## Quality Standards

### Bronze Level (Minimum Viable)
- [ ] Clear, one-sentence mission statement exists
- [ ] Primary beneficiaries are identified
- [ ] Basic value creation is described
- [ ] All required metadata fields completed

### Silver Level (Investment Ready)
- [ ] Mission statement is memorable and distinctive
- [ ] Detailed beneficiary analysis completed
- [ ] Value creation is quantifiable
- [ ] Mission guides documented decisions
- [ ] Stakeholder feedback incorporated

### Gold Level (Operational Excellence)
- [ ] Mission evolution history documented
- [ ] Regular mission review process established
- [ ] Mission measurement framework implemented
- [ ] Cultural integration practices documented
- [ ] Mission validation evidence collected

## Common Pitfalls

1. **Too Broad**
   - Problem: "Make the world better" tells us nothing actionable
   - Solution: Focus on specific beneficiaries and value creation

2. **Too Narrow**
   - Problem: Limiting mission to current products rather than purpose
   - Solution: Think about fundamental purpose, not current solutions

3. **Jargon Heavy**
   - Problem: Using corporate speak that employees can't understand or remember
   - Solution: Use simple, clear language that anyone can understand

4. **Profit-Focused**
   - Problem: Confusing mission (why we exist) with business model (how we make money)
   - Solution: Focus on value creation, not value capture

## Success Patterns

1. **Customer-Centric**
   - Clear identification of who benefits and how
   - Specific understanding of customer needs and pain points

2. **Memorable**
   - Short enough to remember, meaningful enough to inspire
   - Employees can recite and explain the mission

3. **Distinctive**
   - Differentiates from competitors and alternatives
   - Unique perspective on value creation

4. **Actionable**
   - Provides decision-making guidance for employees
   - Clear implications for product and service decisions

## Industry Variations

### Software/SaaS
- Often focuses on enabling customer success or transformation
- Emphasis on empowering users or solving business problems
- Technology as means, not end

### Physical Products
- Usually emphasizes improving daily life or experiences
- Focus on quality, accessibility, or innovation
- Tangible impact on customer lives

### Services
- Typically centers on solving specific problems or challenges
- Professional expertise as value driver
- Relationship-based value creation

### Nonprofit
- Always includes social impact and beneficiary outcomes
- Clear social change or improvement goal
- Mission often defines success metrics

## Relationship Guidelines

### Typical Dependencies
- **None** (Mission is foundational)

### Typical Enablements
- **VSN (Vision)**: Mission provides foundation for future aspirations
- **VAL (Values)**: Mission informs behavioral principles
- **STR (Strategy)**: Mission guides strategic approach
- **PUR (Purpose)**: Mission may inform broader social purpose

### Common Conflicts
- **Multiple competing missions** in organization
- **Mission-strategy misalignment**
- **Mission-culture gaps**

## Implementation Guidelines

### Creation Process
1. **Stakeholder Interviews**: Understand founder/leadership intent
2. **Beneficiary Analysis**: Identify who truly benefits from organization
3. **Value Creation Mapping**: Understand unique value provided
4. **Draft and Iterate**: Create multiple versions and test

### Review Process
1. **Annual Review**: Assess mission relevance and effectiveness
2. **Strategic Planning**: Use mission to guide strategic decisions
3. **Performance Assessment**: Measure progress toward mission fulfillment
4. **Cultural Integration**: Ensure mission lives in daily operations

### Measurement Approaches
- **Employee Understanding**: Survey employee ability to explain mission
- **Decision Alignment**: Track strategic decisions against mission
- **Stakeholder Recognition**: Measure external understanding of purpose
- **Value Creation Metrics**: Quantify mission fulfillment

## Document Relationships

This document type commonly relates to:

- **Enables**: VSN (Vision), VAL (Values), STR (Strategy), PUR (Purpose)
- **Informs**: All strategic planning documents
- **Guides**: Product development, hiring, partnership decisions
- **Measures**: Organization performance and culture health

## Validation Checklist

- [ ] Mission statement is clear and memorable
- [ ] Primary beneficiaries are specifically identified
- [ ] Value creation is unique and meaningful
- [ ] Mission provides decision-making guidance
- [ ] All stakeholders understand and support mission
- [ ] Mission differentiates from competitors
- [ ] Evidence of mission fulfillment exists
- [ ] Regular review process established