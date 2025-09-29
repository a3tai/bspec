# THY: Theory of Change Document Type Specification

**Document Type Code:** THY
**Document Type Name:** Theory of Change
**Domain:** Strategic Foundation
**Version:** 1.0.0
**Status:** Draft
**Last Updated:** 2025-09-28

## Purpose and Scope

The Theory of Change document maps the logic model connecting the organization's activities to its intended outcomes. It explains how and why specific actions will lead to desired changes.

## Document Metadata Schema

```yaml
---
id: THY-{change-model-identifier}
title: "Theory of Change"
type: THY
status: Draft|Review|Accepted|Deprecated
version: 1.0.0
owner: Impact-Lead|Strategy-Lead
stakeholders: [leadership-team, impact-team, board]
domain: strategic
priority: medium
scope: global
horizon: long
visibility: internal

depends_on: [PUR-*, MSN-*, STR-*]
enables: [MET-*, EXP-*, LRN-*]

success_criteria:
  - "Logic model is clear and testable"
  - "Assumptions are explicit and validated"
  - "Outcomes are measurable"
  - "Theory guides program design"

assumptions:
  - "Activities will produce intended outputs"
  - "Outputs will lead to desired outcomes"
  - "External factors support change process"

review_cycle: annually
---
```

## Content Structure Template

```markdown
# Theory of Change

## Overview
*Summary of how our work creates intended change*

## Change Logic Model

### Ultimate Goal
*Long-term change we want to see in the world*

### Intended Outcomes
*Specific changes for beneficiaries*

**Short-term Outcomes (0-1 year)**
- Immediate changes for beneficiaries
- Initial behavior changes
- Awareness and knowledge gains
- Skill development

**Medium-term Outcomes (1-3 years)**
- Sustained behavior changes
- System improvements
- Relationship changes
- Capacity building

**Long-term Outcomes (3+ years)**
- Lasting impact
- System transformation
- Cultural change
- Sustainable improvements

### Activities and Outputs
*What we do and what we produce*

**Core Activities**
- Primary interventions
- Key programs and services
- Engagement strategies
- Delivery mechanisms

**Direct Outputs**
- Products and services delivered
- People reached
- Events and touchpoints
- Content and resources created

### Input Requirements
*What we need to create change*

**Human Resources**
- Staff and expertise required
- Partner capabilities
- Volunteer engagement
- Leadership commitment

**Financial Resources**
- Funding requirements
- Cost structure
- Revenue generation
- Resource allocation

**Physical Resources**
- Infrastructure needs
- Technology requirements
- Material resources
- Location access

## Causal Assumptions

### Activity-Output Links
*Why our activities will produce intended outputs*

### Output-Outcome Links
*Why our outputs will lead to desired outcomes*

### Outcome-Impact Links
*Why outcomes will create lasting change*

### External Factors
*Conditions beyond our control that affect success*

## Evidence Base

### Research Foundation
*Academic and practical evidence supporting our theory*

### Historical Precedents
*Similar successful interventions*

### Pilot Results
*Evidence from our own testing*

### Expert Validation
*Review by subject matter experts*

## Measurement Framework

### Outcome Indicators
*How we measure progress toward outcomes*

### Data Collection Methods
*How we gather evidence*

### Evaluation Timeline
*When we assess progress*

### Learning Integration
*How we use evidence to improve*

## Theory Testing

### Key Hypotheses
*Testable assumptions in our logic model*

### Validation Methods
*How we test our assumptions*

### Adaptation Triggers
*When we modify our theory*

### Learning Culture
*How we integrate new evidence*

## Risk Factors

### Theory Risks
*Threats to our logic model*

### Implementation Risks
*Challenges in executing activities*

### External Risks
*Environmental factors that could interfere*

### Mitigation Strategies
*How we address risks*

## Validation
*Evidence that our theory of change is sound and working*
```

## Quality Standards

### Bronze Level (Minimum Viable)
- [ ] Clear logic model from activities to outcomes
- [ ] Ultimate goal and intended outcomes defined
- [ ] Core activities and outputs identified
- [ ] All required metadata fields completed

### Silver Level (Investment Ready)
- [ ] Detailed causal assumptions articulated
- [ ] Evidence base supporting theory provided
- [ ] Measurement framework for outcomes
- [ ] Key hypotheses and validation methods
- [ ] Risk assessment and mitigation strategies
- [ ] Input requirements clearly specified

### Gold Level (Operational Excellence)
- [ ] Regular theory testing and validation
- [ ] Evidence-based theory refinement process
- [ ] Comprehensive data collection system
- [ ] Learning integration and adaptation protocols
- [ ] External expert validation obtained
- [ ] Theory-driven program design implemented

## Common Pitfalls

1. **Linear Thinking**
   - Problem: Assuming simple cause-effect relationships in complex systems
   - Solution: Account for feedback loops, multiple pathways, and system complexity

2. **Activity Focus**
   - Problem: Emphasizing what we do rather than change we create
   - Solution: Focus on outcomes and impact, not just activities

3. **Assumption Blindness**
   - Problem: Not making underlying beliefs explicit and testable
   - Solution: Document and test all key assumptions regularly

4. **Static Model**
   - Problem: Treating theory as fixed rather than learning-driven
   - Solution: Build in adaptation mechanisms and regular updates

## Success Patterns

1. **Evidence-Based**
   - Grounded in research and validated through experience
   - Continuous evidence collection and integration

2. **Testable Assumptions**
   - Clear hypotheses that can be validated or disproven
   - Regular testing and refinement

3. **Learning-Oriented**
   - Designed to generate evidence and adapt based on results
   - Culture of continuous improvement

4. **Stakeholder-Informed**
   - Developed with input from beneficiaries and experts
   - Regular stakeholder feedback integration

## Industry Variations

### Software/SaaS
- User adoption and behavior change theories
- Technology adoption curves and network effects
- Digital transformation change models
- Platform ecosystem development

### Physical Products
- Consumer behavior and adoption patterns
- Market penetration and diffusion theories
- Quality improvement and satisfaction models
- Supply chain impact theories

### Services
- Service delivery and client change models
- Professional development and capacity building
- Relationship building and trust theories
- Knowledge transfer and application

### Nonprofit
- Social change and systems transformation
- Community development and empowerment
- Behavior change and social norms
- Policy change and advocacy theories

## Relationship Guidelines

### Typical Dependencies
- **PUR (Purpose)**: Theory shows how purpose is achieved
- **MSN (Mission)**: Theory supports mission fulfillment
- **STR (Strategy)**: Theory aligns with strategic approach

### Typical Enablements
- **MET (Metrics)**: Theory defines what to measure
- **EXP (Experiments)**: Theory guides experimentation
- **LRN (Learnings)**: Theory provides learning framework

### Common Conflicts
- **Theory-practice gaps** in implementation
- **Competing theories** within organization
- **Theory-resource misalignment**

## Implementation Guidelines

### Creation Process
1. **Outcome Definition**: Start with desired long-term change
2. **Backward Mapping**: Work backwards from outcomes to activities
3. **Assumption Identification**: Make causal beliefs explicit
4. **Evidence Gathering**: Collect supporting research and data
5. **Stakeholder Validation**: Test theory with beneficiaries and experts

### Review Process
1. **Annual Theory Review**: Assess theory validity and relevance
2. **Quarterly Progress Assessment**: Track outcome indicators
3. **Assumption Testing**: Regularly validate key assumptions
4. **Evidence Integration**: Incorporate new learning into theory

### Measurement Approaches
- **Outcome Tracking**: Monitor progress toward intended outcomes
- **Assumption Testing**: Validate causal relationships
- **Process Evaluation**: Assess activity implementation
- **Impact Assessment**: Measure ultimate goal achievement

## Document Relationships

This document type commonly relates to:

- **Depends on**: PUR (Purpose), MSN (Mission), STR (Strategy)
- **Enables**: MET (Metrics), EXP (Experiments), LRN (Learnings)
- **Informs**: Program design, evaluation frameworks, impact measurement
- **Guides**: Resource allocation, activity prioritization, adaptation decisions

## Validation Checklist

- [ ] Clear logic model from activities to ultimate goal
- [ ] Short, medium, and long-term outcomes defined
- [ ] Core activities and direct outputs specified
- [ ] Input requirements (human, financial, physical) identified
- [ ] Causal assumptions explicitly stated
- [ ] Evidence base supporting theory documented
- [ ] Outcome indicators and measurement methods defined
- [ ] Key hypotheses and validation approaches outlined
- [ ] Risk factors and mitigation strategies identified
- [ ] External factors and conditions acknowledged
- [ ] Theory testing and adaptation protocols established
- [ ] Regular review and updating process implemented