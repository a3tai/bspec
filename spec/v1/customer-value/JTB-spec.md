# JTB: Jobs-to-be-Done Document Type Specification

**Document Type Code:** JTB
**Document Type Name:** Jobs-to-be-Done
**Domain:** Customer & Value
**Version:** 1.0.0
**Status:** Draft
**Last Updated:** 2025-09-28

## Abstract

This specification defines the Jobs-to-be-Done document type within the BSpec 1.0 Universal Business Specification Standard. It establishes normative requirements, structured templates, and implementation guidance for organizations documenting jobs-to-be-done within the customer-value domain. This specification enables systematic, machine-readable documentation that supports strategic planning, operational execution, and organizational alignment.

## Purpose and Scope

The Jobs-to-be-Done document defines the specific outcomes customers hire products or services to achieve. It captures the functional, emotional, and social jobs that drive customer behavior and innovation opportunities.

## Document Metadata Schema

```yaml
---
id: JTB-{job-identifier}
title: "Job-to-be-Done: [Job Description]"
type: JTB
status: Draft|Review|Accepted|Deprecated
version: 1.0.0
owner: Product-Strategy|Customer-Research
stakeholders: [product-team, design-team, marketing-team]
domain: customer
priority: critical
scope: global
horizon: medium
visibility: internal

depends_on: [PER-*]
enables: [USE-*, PRD-*, SVC-*]

success_criteria:
  - "Job is validated through customer research"
  - "Job drives product development priorities"
  - "Solution effectively serves the job"
  - "Job measurement shows progress"

assumptions:
  - "Job represents significant customer need"
  - "Customers actively seek solutions for this job"
  - "Job is underserved by existing solutions"

research_sources:
  - "Customer interviews focusing on job context"
  - "Observational research of job performance"
  - "Outcome-driven innovation research"

validation_method: "Jobs-to-be-Done interview methodology"
review_cycle: quarterly
---
```

## Content Structure Template

```markdown
# Job-to-be-Done: [Job Description]

## Overview
*Brief description of the job customers are trying to obtain done*

## Job Definition

### Job Statement
**Primary Job**: *"When I [situation], I want to [motivation], so I can [outcome]"*

**Job Categories**:
- **Functional Job**: The practical task being accomplished
- **Emotional Job**: The feeling the customer wants to achieve
- **Social Job**: How the customer wants to be perceived by others

### Job Scope
**Job Boundaries**
- What's included in this job
- What's excluded from this job
- Related but separate jobs
- Job hierarchy and relationships

**Job Universality**
- Who experiences this job
- Geographic variations
- Cultural differences
- Market segment variations

## Job Context

### Job Circumstances
**Situational Context**
- When this job arises
- Where it typically occurs
- Environmental factors
- Timing considerations

**Trigger Events**
- What causes the job to arise
- Frequency of occurrence
- Urgency levels
- Predictability patterns

### Job Constraints
**Time Constraints**
- Time availability for job performance
- Deadline pressures
- Sequence dependencies
- Scheduling limitations

**Resource Constraints**
- Budget limitations
- Tool availability
- Skill requirements
- Access restrictions

## Job Process

### Job Steps
**Core Job Process**
1. **[Step 1]**: [Description of what happens]
   - Input requirements
   - Activities performed
   - Decision points
   - Output generated

2. **[Step 2]**: [Description of what happens]
   - Input requirements
   - Activities performed
   - Decision points
   - Output generated

### Job Workflow
**Preparation Phase**
- Setup activities
- Resource gathering
- Stakeholder alignment
- Environment preparation

**Execution Phase**
- Core activities
- Decision making
- Problem solving
- Progress monitoring

**Completion Phase**
- Validation activities
- Communication of results
- Documentation
- Follow-up actions

## Desired Outcomes

### Primary Outcomes
**Functional Outcomes**
- Specific results achieved
- Performance metrics
- Quality standards
- Efficiency measures

**Emotional Outcomes**
- Feelings experienced
- Confidence levels
- Satisfaction measures
- Stress reduction

**Social Outcomes**
- Reputation effects
- Relationship impacts
- Status implications
- Recognition received

### Success Criteria
**Objective Measures**
- Quantifiable results
- Performance benchmarks
- Quality indicators
- Time-based measures

**Subjective Measures**
- Satisfaction levels
- Confidence feelings
- Pride in accomplishment
- Sense of control

## Current Solutions

### Existing Approaches
**Category 1: [Solution Type]**
- Description of approach
- How it addresses the job
- Strengths and limitations
- Usage patterns

**Category 2: [Solution Type]**
- Description of approach
- How it addresses the job
- Strengths and limitations
- Usage patterns

### Solution Evaluation
**Outcome Achievement**
- How well each solution delivers desired outcomes
- Gap analysis and unmet needs
- Customer satisfaction assessment
- Market adoption patterns

## Pain Points and Frustrations

### Job Performance Challenges
**Execution Difficulties**
- Process complexity
- Skill requirements
- Tool limitations
- Resource constraints

**Quality Issues**
- Accuracy problems
- Consistency challenges
- Reliability concerns
- Standard variations

### Emotional Frustrations
**Stress and Anxiety**
- Uncertainty about outcomes
- Performance pressure
- Risk concerns
- Complexity overwhelm

**Confidence Issues**
- Skill inadequacy feelings
- Decision uncertainty
- Result unpredictability
- Competence concerns

## Opportunity Analysis

### Underserved Outcomes
**Functional Gaps**
- Unaddressed outcomes
- Performance limitations
- Quality issues
- Efficiency opportunities

**Emotional Gaps**
- Confidence building
- Stress reduction
- Satisfaction enhancement
- Pride creation

### Market Opportunities
**Solution Innovation**
- New approach possibilities
- Technology enablement
- Process redesign
- Experience improvement

**Value Creation**
- Outcome enhancement
- Efficiency improvement
- Quality advancement
- Cost reduction

## Our Solution Fit

### Job Addressing Approach
**How We Serve the Job**
- Our solution's approach
- Job step coverage
- Outcome delivery
- Value proposition

**Competitive Advantages**
- Unique capabilities
- Superior outcomes
- Better experience
- Distinctive value

### Solution Performance
**Outcome Achievement**
- How well we deliver desired outcomes
- Customer success metrics
- Performance benchmarks
- Satisfaction levels

## Measurement Framework

### Job Success Metrics
**Outcome Metrics**
- Functional outcome measures
- Emotional outcome indicators
- Social outcome assessments
- Overall success rates

**Process Metrics**
- Job completion time
- Effort requirements
- Error rates
- Satisfaction scores

### Research Methods
**Customer Research**
- Jobs interview methodology
- Outcome-driven innovation surveys
- Observational research
- Usage analytics

**Validation Approaches**
- Customer success tracking
- Satisfaction measurement
- Behavioral analysis
- Competitive benchmarking

## Validation
*Evidence that job definition is accurate and our solution serves it effectively*
```

## Quality Standards

### Bronze Level (Minimum Viable)
- [ ] Job statement clearly defined with situation-motivation-outcome
- [ ] Job categories (functional, emotional, social) identified
- [ ] Basic job process and desired outcomes documented
- [ ] All required metadata fields completed

### Silver Level (Investment Ready)
- [ ] Comprehensive job context and constraints analysis
- [ ] Detailed job process with steps and workflow
- [ ] Current solution evaluation and gap analysis
- [ ] Pain points and opportunity identification
- [ ] Our solution fit assessment
- [ ] Research validation with customer interviews

### Gold Level (Operational Excellence)
- [ ] Dynamic job tracking with outcome measurement
- [ ] Advanced job analytics and pattern recognition
- [ ] Continuous job evolution monitoring
- [ ] Solution performance optimization based on job insights
- [ ] Cross-functional job utilization in development
- [ ] Competitive job analysis and positioning

## Common Pitfalls

1. **Solution confusion**: Describing existing solutions rather than underlying jobs
2. **Feature fixation**: Focusing on product features rather than customer outcomes
3. **Process obsession**: Detailing how customers work rather than what they're trying to achieve
4. **Single dimension**: Ignoring emotional and social jobs in favor of functional jobs

## Success Patterns

1. **Outcome-focused**: Clear articulation of what customers want to achieve
2. **Context-rich**: Deep understanding of when and why jobs arise
3. **Research-validated**: Based on direct customer research and observation
4. **Innovation-driving**: Insights that guide product development and improvement

## Relationship Guidelines

### Typical Dependencies
- **PER (Personas)**: Customer personas provide context for job performers

### Typical Enablements
- **USE (Use Cases)**: Jobs inform specific usage scenarios
- **PRD (Product Requirements)**: Jobs drive product feature priorities
- **SVC (Service Design)**: Jobs guide service offering design

## Document Relationships

This document type commonly relates to:

- **Depends on**: PER (Personas)
- **Enables**: USE (Use Cases), PRD (Product Requirements), SVC (Service Design)
- **Informs**: Product strategy, feature development, value proposition design
- **Guides**: Innovation priorities, user experience design, competitive positioning

## Validation Checklist

- [ ] Job statement follows "When I... I want to... so I can..." format
- [ ] Functional, emotional, and social job dimensions documented
- [ ] Job context and trigger events clearly identified
- [ ] Job process steps and workflow mapped
- [ ] Desired outcomes and success criteria defined
- [ ] Current solutions evaluated for job performance
- [ ] Pain points and frustrations documented
- [ ] Opportunity analysis identifies innovation areas
- [ ] Our solution fit and competitive advantages assessed
- [ ] Measurement framework established with metrics
- [ ] Research validation confirms job accuracy
- [ ] Regular review process ensures job evolution tracking