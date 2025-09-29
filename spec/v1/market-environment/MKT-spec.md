# MKT: Market Definition Document Type Specification

**Document Type Code:** MKT
**Document Type Name:** Market Definition
**Domain:** Market & Environment
**Version:** 1.0.0
**Status:** Draft
**Last Updated:** 2025-09-28

## Purpose and Scope

The Market Definition document establishes the boundaries, size, and characteristics of the addressable market. It defines TAM (Total Addressable Market), SAM (Serviceable Addressable Market), and SOM (Serviceable Obtainable Market).

## Document Metadata Schema

```yaml
---
id: MKT-{market-identifier}
title: "Market Definition and Sizing"
type: MKT
status: Draft|Review|Accepted|Deprecated
version: 1.0.0
owner: Strategy-Lead|Market-Research
stakeholders: [leadership-team, product-team, sales-team, investors]
domain: market
priority: critical
scope: global
horizon: medium
visibility: internal

depends_on: [MSN-*, STR-*]
enables: [SEG-*, CMP-*, GTM-*, FIN-*]

success_criteria:
  - "Market size estimates are research-backed"
  - "Market boundaries are clearly defined"
  - "Growth rates and trends are validated"
  - "Addressable market is realistic and achievable"

assumptions:
  - "Market research sources are reliable"
  - "Market growth projections are realistic"
  - "Competitive dynamics are well understood"

review_cycle: quarterly
---
```

## Content Structure Template

```markdown
# Market Definition and Sizing

## Overview
*Summary of the market we're addressing and its key characteristics*

## Market Framework

### Market Definition
*Precise boundaries of our addressable market*

**Problem-Based Definition**
- Core problem we solve
- Variations of the problem
- Adjacent problems
- Problem severity and frequency

**Solution-Based Definition**
- Category of solutions
- Alternative approaches
- Substitute products/services
- Solution evolution trends

**Customer-Based Definition**
- Who experiences the problem
- Decision-making units
- Buying processes
- Usage patterns

### Market Sizing

### Total Addressable Market (TAM)
**Definition**: *Total market demand if every potential customer bought*

**Calculation Method**:
- Top-down approach: [methodology]
- Bottom-up approach: [methodology]
- Value-theory approach: [methodology]

**Size Estimate**: $X billion
**Data Sources**: [primary and secondary research]
**Confidence Level**: High/Medium/Low
**Validation**: [cross-referencing methods]

### Serviceable Addressable Market (SAM)
**Definition**: *Portion of TAM we can realistically serve*

**Geographic Constraints**
- Markets we can enter
- Regulatory limitations
- Operational reach
- Cultural factors

**Customer Constraints**
- Segments we can serve
- Use cases we address
- Price points we target
- Channel requirements

**Size Estimate**: $X million
**Percentage of TAM**: X%

### Serviceable Obtainable Market (SOM)
**Definition**: *Realistic market share we can capture*

**Competitive Analysis**
- Market share distribution
- Competitive advantages
- Entry barriers
- Time to market

**Organizational Capacity**
- Resources available
- Execution capabilities
- Growth constraints
- Investment requirements

**Size Estimate**: $X million
**Percentage of SAM**: X%
**Timeline to Capture**: X years

## Market Characteristics

### Market Maturity
*Stage of market development*

**Market Stage**
- Emerging/Growth/Mature/Declining
- Characteristics of current stage
- Timeline to next stage
- Implications for strategy

**Customer Sophistication**
- Awareness of problem
- Understanding of solutions
- Buying experience
- Decision criteria evolution

**Solution Evolution**
- Technology advancement
- Feature development
- Price compression
- Quality improvements

### Market Dynamics

**Growth Drivers**
- Factors accelerating market growth
- Quantified impact on market size
- Sustainability of growth factors
- Geographic variations

**Market Constraints**
- Factors limiting market growth
- Regulatory barriers
- Economic constraints
- Cultural resistance

**Seasonality and Cycles**
- Predictable market patterns
- Economic cycle sensitivity
- Seasonal variations
- Planning implications

### Customer Behavior

**Buying Patterns**
- Purchase frequency
- Decision timelines
- Buying criteria
- Price sensitivity

**Usage Patterns**
- Consumption frequency
- Usage intensity
- Adoption curves
- Renewal behaviors

**Value Realization**
- Time to value
- Success metrics
- ROI expectations
- Satisfaction drivers

## Market Segmentation Preview

### Primary Segments
*High-level customer groupings*

### Segment Characteristics
*Key differentiators between segments*

### Segment Priorities
*Which segments offer best opportunities*

## Market Research Foundation

### Primary Research
*Direct customer and market research*
- Customer interviews
- Surveys and quantitative research
- Focus groups
- Market testing

### Secondary Research
*Published sources and analysis*
- Industry reports
- Government data
- Academic research
- Analyst reports

### Data Quality Assessment
*Reliability and validity of sources*
- Source credibility
- Data recency
- Sample sizes
- Methodology assessment

## Market Evolution

### Historical Trends
*How the market has developed*
- Size growth over time
- Key inflection points
- Technology adoption
- Competitive evolution

### Future Projections
*Expected market development*
- Growth rate projections
- Emerging segments
- Technology impact
- Disruption potential

### Scenario Planning
*Alternative market futures*
- Optimistic scenario
- Most likely scenario
- Pessimistic scenario
- Black swan events

## Validation
*Evidence supporting market size and characteristics*
```

## Quality Standards

### Bronze Level (Minimum Viable)
- [ ] Clear market definition with boundaries
- [ ] TAM, SAM, and SOM estimates provided
- [ ] Basic market characteristics described
- [ ] All required metadata fields completed

### Silver Level (Investment Ready)
- [ ] Multiple calculation methods for market sizing
- [ ] Market research foundation documented
- [ ] Market dynamics and trends analyzed
- [ ] Segmentation preview provided
- [ ] Data quality assessment included
- [ ] Market evolution timeline outlined

### Gold Level (Operational Excellence)
- [ ] Regular market size validation and updates
- [ ] Comprehensive primary and secondary research
- [ ] Scenario planning for market futures
- [ ] Customer behavior analysis detailed
- [ ] Market intelligence system established
- [ ] Cross-validation with competitive analysis

## Common Pitfalls

1. **Over-estimation bias**
   - Problem: Inflating market size with unrealistic assumptions
   - Solution: Provide measurable evidence and multiple validation methods

2. **Market mirage**
   - Problem: Defining market around current solution rather than customer problem
   - Solution: Start with customer problem and expand outward

3. **Static thinking**
   - Problem: Treating market as fixed rather than evolving ecosystem
   - Solution: Include market evolution and scenario planning

4. **Research laziness**
   - Problem: Relying on outdated or poor-quality market data
   - Solution: Invest in current, high-quality research and validation

## Success Patterns

1. **Problem-centric**
   - Market defined by customer problem, not current solutions
   - Multiple problem variations considered

2. **Multi-method validation**
   - Market size confirmed through multiple calculation approaches
   - Cross-validation between sources and methodologies

3. **Dynamic perspective**
   - Understanding of market evolution and future scenarios
   - Regular updates based on new information

4. **Customer-validated**
   - Market understanding confirmed through direct customer research
   - Customer behavior patterns documented and analyzed

## Industry Variations

### Software/SaaS
- Digital market boundaries and user metrics
- Platform economics and network effects
- Subscription model implications
- Developer ecosystem considerations

### Physical Products
- Manufacturing and distribution constraints
- Supply chain considerations
- Geographic market boundaries
- Physical infrastructure requirements

### Services
- Service delivery capacity constraints
- Professional service market dynamics
- Local vs. global service markets
- Human capital requirements

### B2B Markets
- Enterprise buying cycles and decision units
- Procurement processes and requirements
- Industry-specific regulations
- Channel partner dynamics

## Relationship Guidelines

### Typical Dependencies
- **MSN (Mission)**: Market definition aligns with organizational mission
- **STR (Strategy)**: Market scope supports strategic objectives

### Typical Enablements
- **SEG (Segments)**: Market size enables segment prioritization
- **CMP (Competitive Analysis)**: Market definition frames competitive landscape
- **GTM (Go-to-Market)**: Market understanding drives market entry strategy
- **FIN (Financial Planning)**: Market size informs revenue projections

### Common Conflicts
- **Market-capability misalignment** between market opportunity and organizational capability
- **Market-resource gaps** between market requirements and available resources
- **Market-strategy disconnects** between market realities and strategic assumptions

## Implementation Guidelines

### Creation Process
1. **Problem Definition**: Start with clear customer problem definition
2. **Market Boundary Setting**: Define precise market boundaries and scope
3. **Research Design**: Plan comprehensive primary and secondary research
4. **Data Collection**: Gather market data from multiple reliable sources
5. **Size Calculation**: Apply multiple methodologies for market sizing
6. **Validation**: Cross-validate findings through multiple approaches

### Review Process
1. **Quarterly Market Review**: Assess market size and growth assumptions
2. **Annual Deep Analysis**: Comprehensive market analysis update
3. **Event-Based Updates**: Update based on significant market changes
4. **Research Validation**: Regular validation through customer research

### Measurement Approaches
- **Market Growth Tracking**: Monitor actual market growth vs. projections
- **Share Analysis**: Track our market share evolution
- **Competitive Benchmarking**: Compare market understanding with competitors
- **Customer Validation**: Regular customer research to validate market assumptions

## Document Relationships

This document type commonly relates to:

- **Depends on**: MSN (Mission), STR (Strategy)
- **Enables**: SEG (Segments), CMP (Competitive Analysis), GTM (Go-to-Market), FIN (Financial Planning)
- **Informs**: Product development priorities, investment decisions, resource allocation
- **Guides**: Market entry strategies, competitive positioning, segment prioritization

## Validation Checklist

- [ ] Market boundaries are clearly and precisely defined
- [ ] TAM, SAM, and SOM calculations are methodology-based
- [ ] Multiple validation approaches confirm market size
- [ ] Market characteristics and dynamics are analyzed
- [ ] Customer behavior patterns are documented
- [ ] Market research foundation is comprehensive
- [ ] Market evolution and trends are projected
- [ ] Scenario planning addresses market uncertainty
- [ ] Data quality assessment validates research sources
- [ ] Market definition enables strategic decision-making
- [ ] Regular review and update process established
- [ ] Cross-validation with other market documents completed