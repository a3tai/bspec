# POS: Positioning Document Type Specification

**Document Type Code:** POS
**Document Type Name:** Positioning
**Domain:** Market & Environment
**Version:** 1.0.0
**Status:** Draft
**Last Updated:** 2025-09-28

## Purpose and Scope

The Positioning document defines how the organization wants to be perceived in the market relative to competitors and alternatives. It articulates the unique value proposition and market position.

## Document Metadata Schema

```yaml
---
id: POS-{positioning-strategy-identifier}
title: "Market Positioning Strategy"
type: POS
status: Draft|Review|Accepted|Deprecated
version: 1.0.0
owner: Marketing-Lead|Strategy-Lead
stakeholders: [sales-team, product-team, leadership-team]
domain: market
priority: high
scope: global
horizon: medium
visibility: internal

depends_on: [SEG-*, CMP-*, VAL-*]
enables: [GTM-*, BMC-*, REV-*]

success_criteria:
  - "Positioning is distinctive and memorable"
  - "Target customers understand and value positioning"
  - "Positioning drives consistent messaging"
  - "Market perception aligns with intended positioning"

assumptions:
  - "Target segments value our positioning"
  - "Positioning is sustainable and defensible"
  - "Organization can deliver on positioning promise"

review_cycle: quarterly
---
```

## Content Structure Template

```markdown
# Market Positioning Strategy

## Overview
*Summary of how we position ourselves in the market*

## Positioning Framework

### Positioning Statement
*Formal positioning statement*

**Format**: "For [target customer] who [need/problem], [product/service] is the [category] that [unique benefit] unlike [competition] because [reason to believe]."

**Our Positioning Statement**:
*Complete positioning statement following the format above*

### Positioning Elements

**Target Customer**
- Primary audience for positioning
- Specific customer segments
- Decision-making units
- Influencer groups

**Category Definition**
- Market category we compete in
- Category boundaries
- Category evolution
- Our role in category development

**Unique Value Proposition**
- Primary benefit we deliver
- Emotional benefits
- Functional benefits
- Social benefits

**Competitive Differentiation**
- How we're different from alternatives
- Comparison framework
- Competitive advantages
- Sustainability of differences

**Reason to Believe**
- Evidence supporting our claims
- Proof points and validation
- Customer testimonials
- Third-party endorsements

## Positioning Strategy

### Strategic Positioning Choice
*How we choose to compete*

**Positioning Approach**
- Leader positioning
- Challenger positioning
- Niche specialist
- Innovation pioneer

**Value Positioning**
- Premium positioning
- Value positioning
- Cost leadership
- Performance leadership

**Emotional Positioning**
- Brand personality
- Emotional benefits
- Cultural values
- Aspirational elements

### Positioning Architecture

**Core Message**
- Primary message platform
- Key value proposition
- Supporting messages
- Proof points

**Message Hierarchy**
- Primary message
- Secondary messages
- Supporting evidence
- Call to action

**Audience Adaptation**
- Message variations by segment
- Channel-specific messaging
- Stakeholder customization
- Use case positioning

## Competitive Positioning

### Competitive Context
*How we position relative to competitors*

**Competitive Frame**
- Who we compare against
- Comparison criteria
- Competitive messaging
- Differentiation themes

**Head-to-Head Positioning**
- Direct competitor comparisons
- Feature/benefit battles
- Price/value positioning
- Quality/performance positioning

**Alternative Positioning**
- Against status quo
- Against manual processes
- Against different solutions
- Against do-nothing option

### Positioning Map
*Visual representation of market positioning*

**Positioning Dimensions**
- Primary competitive axis
- Secondary differentiation axis
- Market positioning grid
- White space opportunities

## Market Perception

### Current Perception
*How market currently views us*

**Brand Awareness**
- Aided awareness levels
- Unaided awareness levels
- Brand recognition
- Consideration set inclusion

**Brand Associations**
- Positive associations
- Negative associations
- Strength of associations
- Unique associations

**Perception vs. Reality**
- Gaps between perception and positioning
- Misperceptions to correct
- Strengths to amplify
- Weaknesses to address

### Perception Building Strategy
*How we shape market perception*

**Messaging Strategy**
- Key messages by audience
- Message frequency and reach
- Message consistency
- Message evolution

**Channel Strategy**
- Owned media approach
- Earned media strategy
- Paid media investment
- Partner channel messaging

**Experience Strategy**
- Customer experience alignment
- Touchpoint consistency
- Proof point creation
- Advocacy development

## Positioning Implementation

### Internal Alignment
*Ensuring organization delivers on positioning*

**Leadership Alignment**
- Executive understanding
- Strategic commitment
- Resource allocation
- Performance metrics

**Employee Alignment**
- Internal communication
- Training programs
- Performance criteria
- Cultural integration

**Operational Alignment**
- Product development
- Service delivery
- Customer experience
- Quality standards

### External Execution

**Marketing Execution**
- Campaign development
- Content strategy
- Event positioning
- Public relations

**Sales Execution**
- Sales messaging
- Competitive battle cards
- Objection handling
- Customer presentations

**Partnership Execution**
- Channel partner training
- Co-marketing approach
- Ecosystem positioning
- Alliance messaging

## Positioning Evolution

### Market Response
*How market is responding to our positioning*

### Positioning Adjustments
*Refinements based on market feedback*

### Future Positioning
*How positioning may evolve*

### Competitive Response
*How competitors are responding to our positioning*

## Validation
*Evidence that positioning is working effectively*
```

## Quality Standards

### Bronze Level (Minimum Viable)
- [ ] Clear positioning statement defined
- [ ] Basic positioning elements identified
- [ ] Target customer and competitive frame established
- [ ] All required metadata fields completed

### Silver Level (Investment Ready)
- [ ] Comprehensive positioning strategy with competitive context
- [ ] Market perception analysis and building strategy
- [ ] Positioning implementation plan for internal and external execution
- [ ] Message hierarchy and audience adaptation
- [ ] Evidence and proof points supporting positioning
- [ ] Initial market validation of positioning effectiveness

### Gold Level (Operational Excellence)
- [ ] Dynamic positioning evolution based on market feedback
- [ ] Comprehensive market perception tracking
- [ ] Advanced positioning measurement and optimization
- [ ] Full organizational alignment on positioning delivery
- [ ] Competitive response monitoring and adaptation
- [ ] Integrated positioning across all customer touchpoints

## Common Pitfalls

1. **Me-too positioning**
   - Problem: Copying successful competitors rather than differentiating
   - Solution: Focus on unique value and authentic differentiation

2. **Feature positioning**
   - Problem: Leading with features rather than customer benefits
   - Solution: Start with customer outcomes and value delivered

3. **Internal focus**
   - Problem: Positioning based on internal perspective rather than customer perception
   - Solution: Ground positioning in customer research and market validation

4. **Positioning drift**
   - Problem: Inconsistent messaging that dilutes market perception
   - Solution: Maintain consistent positioning across all touchpoints

## Success Patterns

1. **Customer-relevant**
   - Positioning that matters to target customers' decision criteria
   - Based on real customer needs and preferences

2. **Distinctive differentiation**
   - Clear and valuable differences from alternatives
   - Sustainable competitive advantages

3. **Authentic delivery**
   - Organization can actually deliver on positioning promise
   - Alignment between positioning and capabilities

4. **Consistent execution**
   - Unified positioning across all customer touchpoints
   - Coordinated internal and external messaging

## Industry Variations

### Software/SaaS
- Technology leadership positioning
- Platform and integration capabilities
- Scalability and performance benefits
- Developer experience focus

### Physical Products
- Quality and reliability positioning
- Design and aesthetics focus
- Performance and durability benefits
- Manufacturing excellence

### Services
- Expertise and experience positioning
- Relationship and trust focus
- Outcome and results orientation
- Cultural and values alignment

### B2B Markets
- Industry expertise positioning
- Integration and compatibility focus
- Compliance and security emphasis
- Partnership and ecosystem benefits

## Relationship Guidelines

### Typical Dependencies
- **SEG (Segments)**: Target segments inform positioning strategy
- **CMP (Competitive Analysis)**: Competitive landscape shapes differentiation
- **VAL (Values)**: Organizational values support authentic positioning

### Typical Enablements
- **GTM (Go-to-Market)**: Positioning drives market entry strategy
- **BMC (Brand Management)**: Positioning guides brand development
- **REV (Revenue Model)**: Positioning influences pricing and packaging

### Common Conflicts
- **Differentiation-capability gaps** between positioning claims and delivery
- **Market perception** vs. intended positioning
- **Competitive response** undermining positioning effectiveness

## Implementation Guidelines

### Creation Process
1. **Customer Research**: Understand target customer decision criteria
2. **Competitive Analysis**: Map competitive landscape and white space
3. **Value Proposition Development**: Define unique value and benefits
4. **Positioning Framework**: Create comprehensive positioning strategy
5. **Message Development**: Build messaging hierarchy and proof points
6. **Market Validation**: Test positioning with target customers

### Review Process
1. **Quarterly Positioning Review**: Assess positioning effectiveness and market response
2. **Annual Positioning Strategy**: Comprehensive positioning strategy review
3. **Market Perception Tracking**: Regular measurement of market perception
4. **Competitive Response Analysis**: Monitor and respond to competitive positioning

### Measurement Approaches
- **Brand Tracking**: Monitor brand awareness, perception, and associations
- **Message Testing**: Test message effectiveness and resonance
- **Competitive Positioning**: Track relative positioning vs. competitors
- **Customer Feedback**: Gather direct feedback on positioning effectiveness

## Document Relationships

This document type commonly relates to:

- **Depends on**: SEG (Segments), CMP (Competitive Analysis), VAL (Values)
- **Enables**: GTM (Go-to-Market), BMC (Brand Management), REV (Revenue Model)
- **Informs**: Marketing strategy, sales enablement, product development
- **Guides**: Messaging strategy, competitive strategy, brand development

## Validation Checklist

- [ ] Positioning statement is clear, compelling, and distinctive
- [ ] Target customer and competitive frame are well-defined
- [ ] Unique value proposition is authentic and deliverable
- [ ] Competitive differentiation is sustainable and meaningful
- [ ] Evidence and proof points support positioning claims
- [ ] Message hierarchy addresses different audiences and use cases
- [ ] Internal alignment ensures organizational delivery on positioning
- [ ] External execution plan covers all customer touchpoints
- [ ] Market perception tracking measures positioning effectiveness
- [ ] Competitive response is monitored and addressed
- [ ] Positioning evolution process adapts to market changes
- [ ] Customer validation confirms positioning resonance