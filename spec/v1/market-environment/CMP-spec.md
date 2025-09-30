# CMP: Competitive Analysis Document Type Specification

**Document Type Code:** CMP
**Document Type Name:** Competitive Analysis
**Domain:** Market & Environment
**Version:** 1.0.0
**Status:** Draft
**Last Updated:** 2025-09-28

## Abstract

This specification defines the Competitive Analysis document type within the BSpec 1.0 Universal Business Specification Standard. It establishes normative requirements, structured templates, and implementation guidance for organizations documenting competitive analysis within the market-environment domain. This specification enables systematic, machine-readable documentation that supports strategic planning, operational execution, and organizational alignment.

## Purpose and Scope

The Competitive Analysis document maps the competitive landscape, analyzes key competitors, and identifies competitive threats and opportunities. It provides intelligence for strategic positioning and tactical responses.

## Document Metadata Schema

```yaml
---
id: CMP-{competitive-landscape-identifier}
title: "Competitive Analysis"
type: CMP
status: Draft|Review|Accepted|Deprecated
version: 1.0.0
owner: Strategy-Lead|Product-Marketing
stakeholders: [leadership-team, sales-team, product-team]
domain: market
priority: high
scope: global
horizon: short
visibility: internal

depends_on: [MKT-*, SEG-*]
enables: [POS-*, MOT-*, STR-*]

success_criteria:
  - "Competitive intelligence is current and actionable"
  - "Analysis drives positioning and product decisions"
  - "Competitive threats are identified early"
  - "Sales team is equipped for competitive battles"

assumptions:
  - "Competitive information sources are reliable"
  - "Competitive landscape analysis is comprehensive"
  - "Competitive intelligence is updated regularly"

review_cycle: monthly
---
```

## Content Structure Template

```markdown
# Competitive Analysis

## Overview
*Summary of competitive landscape and key strategic implications*

## Competitive Framework

### Competition Definition
*How we define and categorize competitors*

**Direct Competitors**
- organizations solving same problem with similar approach
- Head-to-head competitive situations
- Feature-to-feature comparisons
- Same customer segments

**Indirect Competitors**
- Alternative solutions to same problem
- Different approach, same outcome
- Partial overlap in functionality
- Cross-category competition

**Substitute Solutions**
- Non-technology alternatives
- Manual processes
- In-house solutions
- Status quo/do nothing

**Future Competitors**
- organizations likely to enter market
- Adjacent market players
- Technology disruptors
- Platform players

### Competitive Landscape Map
*Visual representation of competitive space*

**Market Position Mapping**
- Price vs. feature positioning
- Market share vs. growth rate
- Customer satisfaction vs. market presence
- Innovation vs. execution

**Competitive Groups**
- Strategic groups analysis
- Mobility barriers between groups
- Group dynamics and rivalry
- White space identification

## Competitor Profiles

### Competitor 1: [Company Name]
**Company Overview**
- Founding date and history
- Size and scale metrics
- Geographic presence
- Corporate structure

**Business Model**
- Revenue streams
- Pricing strategy
- Customer acquisition model
- Partnership approach

**Product/Service Offering**
- Core products/services
- Feature comparison
- Quality assessment
- Innovation pipeline

**Market Position**
- Market share
- Customer segments served
- Brand positioning
- Competitive messaging

**Strengths**
- Key competitive advantages
- Unique capabilities
- Market leadership areas
- Customer loyalty factors

**Weaknesses**
- Competitive vulnerabilities
- Customer complaints
- Product/service gaps
- Execution challenges

**Strategy Analysis**
- Strategic priorities
- Investment areas
- Growth initiatives
- Market expansion plans

**Financial Performance**
- Revenue and growth
- Profitability
- Funding and investments
- Financial health indicators

**Competitive Behavior**
- Response patterns
- Pricing behavior
- Partnership strategies
- Innovation approach

**Threat Assessment**
- Threat level: High/Medium/Low
- Specific threat areas
- Timeline for competitive impact
- Required defensive actions

### Competitor 2: [Company Name]
*[Same structure as Competitor 1]*

### [Continue for each major competitor]

## Competitive Dynamics

### Market Structure
*Competitive intensity and dynamics*

**Rivalry Intensity**
- Number of competitors
- Market growth rate
- Product differentiation
- Exit barriers

**Competitive Moves**
- Recent strategic initiatives
- Pricing changes
- Product launches
- Market expansion

**Collaboration Patterns**
- Industry partnerships
- Standards participation
- Ecosystem relationships
- Co-opetition dynamics

### Competitive Intelligence

**Information Sources**
- Public information sources
- Customer feedback
- Partner intelligence
- Market research
- Sales team insights

**Intelligence Gathering**
- Systematic monitoring processes
- Competitive research methods
- Information validation
- Update frequency

**Early Warning Systems**
- Indicators of competitive threats
- Monitoring dashboards
- Alert mechanisms
- Response protocols

## Competitive Positioning

### Our Competitive Position
*How we stack up against competitors*

**Competitive Advantages**
- Areas where we lead
- Unique differentiators
- Sustainable advantages
- Customer preference drivers

**Competitive Gaps**
- Areas where we lag
- Feature disparities
- Market presence gaps
- Execution challenges

**Competitive Parity**
- Areas of similar capability
- Table stakes features
- Commoditized aspects
- Standards compliance

### Win/Loss Analysis
*Understanding competitive outcomes*

**Win Factors**
- Reasons we win deals
- Competitive advantages in action
- Customer decision criteria
- Successful positioning themes

**Loss Factors**
- Reasons we lose deals
- Competitive disadvantages
- Missing capabilities
- Positioning weaknesses

**Deal Intelligence**
- Competitive encounter frequency
- Win rates by competitor
- Deal size and type patterns
- Sales cycle impacts

## Strategic Implications

### Competitive Strategy
*How competitive landscape shapes our strategy*

**Competitive Response Strategy**
- How we respond to competitive moves
- Proactive competitive initiatives
- Defensive strategies
- Competitive positioning

**Product Strategy Implications**
- Feature development priorities
- Innovation requirements
- Quality benchmarks
- Roadmap adjustments

**Go-to-Market Implications**
- Competitive messaging
- Sales enablement needs
- Channel strategy adjustments
- Pricing considerations

### Future Scenarios

**Competitive Evolution**
- Expected landscape changes
- New entrant possibilities
- Consolidation potential
- Disruption scenarios

**Strategic Preparations**
- Capabilities to build
- Partnerships to develop
- Market positions to strengthen
- Threats to mitigate

## Validation
*Evidence that competitive analysis is accurate and actionable*
```

## Quality Standards

### Bronze Level (Minimum Viable)
- [ ] Key competitors identified and profiled
- [ ] Basic competitive positioning analysis
- [ ] Competitive threats and opportunities identified
- [ ] All required metadata fields completed

### Silver Level (Investment Ready)
- [ ] Comprehensive competitor profiles with financial and strategic analysis
- [ ] Competitive dynamics and market structure analysis
- [ ] Win/loss analysis with customer feedback
- [ ] Competitive intelligence gathering process
- [ ] Strategic implications for product and go-to-market
- [ ] Early warning systems for competitive threats

### Gold Level (Operational Excellence)
- [ ] Real-time competitive intelligence system
- [ ] Advanced competitive scenario planning
- [ ] Competitive response playbooks
- [ ] Sales team competitive enablement program
- [ ] Regular competitive strategy updates
- [ ] Competitive advantage measurement and tracking

## Common Pitfalls

1. **Competitor fixation**
   - Problem: Focusing too much on competitors rather than customers
   - Solution: Balance competitive analysis with customer needs focus

2. **Static analysis**
   - Problem: Treating competitive landscape as unchanging
   - Solution: Implement dynamic monitoring and regular updates

3. **Information bias**
   - Problem: Relying on outdated or biased competitive information
   - Solution: Use multiple sources and validate information regularly

4. **Feature obsession**
   - Problem: Competing on features rather than customer value
   - Solution: Focus on customer outcomes and value delivery

## Success Patterns

1. **Customer-centric competition**
   - Understanding how competitors impact customer choice
   - Focus on customer decision criteria and value

2. **Dynamic monitoring**
   - Regular updates and competitive intelligence gathering
   - Early warning systems for competitive changes

3. **Actionable insights**
   - Analysis that drives specific strategic and tactical decisions
   - Clear implications for product, marketing, and sales

4. **Balanced perspective**
   - Honest assessment of both competitive strengths and weaknesses
   - Recognition of competitive threats and opportunities

## Industry Variations

### Software/SaaS
- Feature comparison matrices
- Technology stack analysis
- Integration capabilities
- Platform ecosystem competition

### Physical Products
- Manufacturing capabilities
- Distribution networks
- Quality and reliability
- Supply chain advantages

### Services
- Service delivery models
- Expertise and talent
- Customer relationships
- Geographic coverage

### B2B Markets
- Enterprise sales capabilities
- Partner ecosystems
- Compliance and security
- Industry expertise

## Relationship Guidelines

### Typical Dependencies
- **MKT (Market Definition)**: Market boundaries frame competitive landscape
- **SEG (Segments)**: Segment analysis provides competitive context

### Typical Enablements
- **POS (Positioning)**: Competitive analysis drives positioning strategy
- **MOT (Moats)**: Competitive gaps identify moat opportunities
- **STR (Strategy)**: Competitive landscape shapes strategic choices

### Common Conflicts
- **Analysis-action balance** between thorough analysis and timely decisions
- **Competitor focus** vs. customer focus
- **Competitive response** vs. strategic consistency

## Implementation Guidelines

### Creation Process
1. **Competitor Identification**: Define and categorize competitive landscape
2. **Intelligence Gathering**: Collect information from multiple sources
3. **Competitor Profiling**: Create comprehensive competitor profiles
4. **Competitive Analysis**: Analyze competitive dynamics and positioning
5. **Strategic Assessment**: Determine strategic implications and responses
6. **Validation**: Validate analysis through customer and market feedback

### Review Process
1. **Monthly Intelligence Updates**: Regular competitive intelligence updates
2. **Quarterly Landscape Review**: Comprehensive competitive landscape assessment
3. **Annual Strategy Review**: Deep competitive strategy analysis
4. **Event-Based Analysis**: Analysis triggered by competitive moves

### Measurement Approaches
- **Competitive Win Rates**: Track win/loss rates against specific competitors
- **Market Share**: Monitor relative market share evolution
- **Customer Perception**: Measure customer perception of competitive position
- **Response Effectiveness**: Assess effectiveness of competitive responses

## Document Relationships

This document type commonly relates to:

- **Depends on**: MKT (Market Definition), SEG (Segments)
- **Enables**: POS (Positioning), MOT (Moats), STR (Strategy)
- **Informs**: Product development, marketing strategy, sales enablement
- **Guides**: Strategic planning, investment priorities, partnership decisions

## Validation Checklist

- [ ] Competitive landscape is comprehensively mapped
- [ ] Key competitors are thoroughly profiled
- [ ] Competitive dynamics and market structure analyzed
- [ ] Our competitive position is honestly assessed
- [ ] Win/loss analysis provides actionable insights
- [ ] Strategic implications for product and go-to-market identified
- [ ] Competitive intelligence gathering process established
- [ ] Early warning systems for competitive threats operational
- [ ] Sales team equipped with competitive intelligence
- [ ] Regular review and update process implemented
- [ ] Customer feedback validates competitive assessment
- [ ] Competitive response strategies prepared