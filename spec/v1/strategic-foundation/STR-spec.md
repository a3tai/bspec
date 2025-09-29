# STR: Strategy Document Type Specification

**Document Type Code:** STR
**Document Type Name:** Business Strategy
**Domain:** Strategic Foundation
**Version:** 1.0.0
**Status:** Draft
**Last Updated:** 2025-09-28

## Purpose and Scope

The Strategy document defines how the organization will achieve its vision and compete in its chosen markets. It articulates the key choices about where to play, how to win, and what capabilities to build.

## Document Metadata Schema

```yaml
---
id: STR-{strategy-identifier}
title: "Business Strategy"
type: STR
status: Draft|Review|Accepted|Deprecated
version: 1.0.0
owner: CEO|Strategy-Lead
stakeholders: [leadership-team, board, department-heads]
domain: strategic
priority: critical
scope: global
horizon: medium
visibility: internal

depends_on: [MSN-*, VSN-*, VAL-*, MKT-*, CMP-*]
enables: [OBJ-*, GTM-*, GRW-*, BMC-*]

success_criteria:
  - "Strategy clearly differentiates from competitors"
  - "Resource allocation aligns with strategic priorities"
  - "All departments understand strategic direction"
  - "Strategy guides major decisions and investments"

assumptions:
  - "Market analysis supports strategic choices"
  - "Organizational capabilities match strategic requirements"
  - "Competitive positioning is sustainable"

review_cycle: quarterly
---
```

## Content Structure Template

```markdown
# Business Strategy

## Overview
*Brief summary of strategic approach and key differentiators*

## Strategic Framework

### Where to Play
*Market and customer choices*

**Target Markets**
- Primary market segments
- Geographic markets
- Market timing and sequencing

**Customer Focus**
- Primary customer segments
- Customer needs being served
- Customer acquisition strategy

**Product/Service Scope**
- Core offerings
- Adjacent opportunities
- Boundaries and what we won't do

### How to Win
*Competitive approach and differentiation*

**Value Proposition**
- Unique value we deliver
- Why customers choose us over alternatives
- Price-value positioning

**Competitive Advantages**
- Core differentiators
- Sustainable moats
- Barriers to competitor entry

**Business Model Logic**
- How we create value
- How we capture value
- How we sustain value creation

### Capabilities Required
*What we must be excellent at*

**Core Capabilities**
- Must-have organizational strengths
- Distinctive competencies
- Capability development priorities

**Strategic Investments**
- Technology investments
- Talent and hiring priorities
- Infrastructure requirements
- Partnership strategies

## Strategic Choices

### Market Strategy
*How we approach our chosen markets*
- Market entry strategies
- Market development approach
- Competitive positioning
- Market share objectives

### Product Strategy
*How we develop and evolve our offerings*
- Product development priorities
- Innovation approach
- Quality and differentiation strategy
- Product lifecycle management

### Customer Strategy
*How we acquire, serve, and retain customers*
- Customer acquisition strategy
- Customer experience strategy
- Customer success and retention
- Customer expansion strategy

### Operational Strategy
*How we deliver value efficiently*
- Operating model design
- Process optimization priorities
- Technology and automation strategy
- Quality and excellence approach

### Financial Strategy
*How we fund and measure success*
- Capital allocation priorities
- Revenue growth strategy
- Profitability targets
- Investment and funding approach

## Strategic Initiatives
*Key projects and programs to execute strategy*

### Year 1 Priorities
*Most critical strategic initiatives*

### Multi-Year Programs
*Longer-term strategic investments*

### Resource Requirements
*What we need to execute successfully*

## Risk and Contingencies
*Strategic risks and alternative scenarios*

### Key Strategic Risks
*Threats to strategic success*

### Contingency Plans
*Alternative approaches if assumptions prove wrong*

### Success Metrics
*How we'll measure strategic progress*

## Validation
*Evidence that strategy is working and assumptions are valid*
```

## Quality Standards

### Bronze Level (Minimum Viable)
- [ ] Clear where-to-play choices defined
- [ ] Basic how-to-win approach articulated
- [ ] Core capabilities identified
- [ ] All required metadata fields completed

### Silver Level (Investment Ready)
- [ ] Comprehensive strategic framework
- [ ] Detailed competitive analysis integrated
- [ ] Strategic initiatives mapped and prioritized
- [ ] Resource requirements quantified
- [ ] Risk assessment and contingency planning
- [ ] Success metrics defined

### Gold Level (Operational Excellence)
- [ ] Regular strategy review and updating process
- [ ] Strategy execution tracking system
- [ ] Cross-functional strategy alignment
- [ ] Strategy communication and cascading process
- [ ] Strategy-based decision framework implemented
- [ ] Strategy evolution history documented

## Common Pitfalls

1. **Everything to Everyone**
   - Problem: Trying to serve all markets and customers
   - Solution: Make explicit choices about focus areas and trade-offs

2. **Tactics vs Strategy**
   - Problem: Confusing activities with strategic choices
   - Solution: Focus on fundamental choices about markets, customers, and capabilities

3. **Copying Competitors**
   - Problem: Following industry best practices instead of differentiating
   - Solution: Develop unique positioning based on distinctive capabilities

4. **Resource Blindness**
   - Problem: Strategy that ignores organizational capabilities and constraints
   - Solution: Align strategy with available and buildable capabilities

## Success Patterns

1. **Clear Trade-offs**
   - Explicit choices about what not to do
   - Focus on specific markets and customer segments

2. **Differentiated Approach**
   - Strategy that creates unique market position
   - Sustainable competitive advantages

3. **Capability-Informed**
   - Strategy that builds on and develops core strengths
   - Realistic assessment of organizational capabilities

4. **Measurable Progress**
   - Clear metrics for strategic success
   - Regular tracking and adjustment

## Industry Variations

### Software/SaaS
- Platform vs point solution strategy
- Horizontal vs vertical market approach
- Freemium vs premium positioning
- API-first vs integrated approach

### Physical Products
- Cost leadership vs differentiation
- Innovation vs operational excellence
- Direct vs channel distribution
- Global vs local market strategy

### Services
- Specialization vs generalization
- High-touch vs scalable delivery
- Expert vs process-driven approach
- Local vs national vs global reach

### Nonprofit
- Direct service vs advocacy strategy
- Grassroots vs institutional approach
- Broad vs focused mission scope
- Volunteer vs professional model

## Relationship Guidelines

### Typical Dependencies
- **MSN (Mission)**: Strategy must align with organizational purpose
- **VSN (Vision)**: Strategy enables vision achievement
- **VAL (Values)**: Strategy reflects organizational values
- **MKT (Market Definition)**: Strategy based on market understanding
- **CMP (Competitive Analysis)**: Strategy informed by competitive landscape

### Typical Enablements
- **OBJ (Objectives)**: Strategy informs goal setting
- **GTM (Go-to-Market)**: Strategy guides market approach
- **GRW (Growth Model)**: Strategy defines growth approach
- **BMC (Business Model Canvas)**: Strategy informs business model

### Common Conflicts
- **Multiple competing strategies** within organization
- **Strategy-capability misalignment**
- **Strategy-resource conflicts**

## Implementation Guidelines

### Creation Process
1. **Environmental Analysis**: Understand market, competition, and trends
2. **Capability Assessment**: Evaluate organizational strengths and gaps
3. **Strategic Option Generation**: Develop alternative strategic approaches
4. **Strategic Choice**: Select preferred strategy based on analysis
5. **Implementation Planning**: Define initiatives and resource requirements

### Review Process
1. **Quarterly Reviews**: Assess strategy execution progress
2. **Annual Strategy Refresh**: Update strategy based on learning and changes
3. **Strategic Decision Testing**: Evaluate major decisions against strategy
4. **Strategy Communication**: Ensure organization understands and executes strategy

### Measurement Approaches
- **Strategic KPI Tracking**: Monitor key strategic metrics
- **Initiative Progress**: Track strategic initiative execution
- **Competitive Position**: Assess competitive standing over time
- **Resource Allocation**: Measure alignment of resources with strategy

## Document Relationships

This document type commonly relates to:

- **Depends on**: MSN (Mission), VSN (Vision), VAL (Values), MKT (Market Definition), CMP (Competitive Analysis)
- **Enables**: OBJ (Objectives), GTM (Go-to-Market), GRW (Growth Model), BMC (Business Model Canvas)
- **Informs**: All operational and functional strategies
- **Guides**: Resource allocation, investment decisions, organizational development

## Validation Checklist

- [ ] Clear where-to-play choices defined
- [ ] Compelling how-to-win approach articulated
- [ ] Core capabilities identified and development planned
- [ ] Strategic initiatives mapped to strategy
- [ ] Resource requirements understood and planned
- [ ] Key strategic risks identified and mitigated
- [ ] Success metrics defined and trackable
- [ ] Strategy differentiates from competitors
- [ ] Strategy aligns with organizational capabilities
- [ ] Leadership team aligned on strategic direction
- [ ] Strategy guides major decisions and investments
- [ ] Regular review and updating process established