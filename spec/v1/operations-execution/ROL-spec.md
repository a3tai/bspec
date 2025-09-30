# ROL: Role Definition Document Type Specification

**Document Type Code:** ROL
**Document Type Name:** Role Definition
**Domain:** Operations & Execution
**Version:** 1.0.0
**Status:** Draft
**Last Updated:** 2025-09-28

## Abstract

This specification defines the Role Definition document type within the BSpec 1.0 Universal Business Specification Standard. It establishes normative requirements, structured templates, and implementation guidance for organizations documenting role definition within the operations-execution domain. This specification enables systematic, machine-readable documentation that supports strategic planning, operational execution, and organizational alignment.

## Purpose and Scope

The Role Definition defines systematic approaches to designing and documenting organizational roles that clarify responsibilities, authorities, and requirements. It establishes role frameworks that enable effective hiring, performance management, and career development.

## Document Metadata Schema

```yaml
---
id: ROL-{role-name}
title: "Role — {Role Title}"
type: ROL
status: Draft|Review|Approved|Active|Deprecated
version: 1.0.0
owner: Role-Owner|HR-Team
stakeholders: [hr-team, hiring-managers, team-leads, employees]
domain: operations
priority: Critical|High|Medium|Low
scope: role-management
horizon: operational
visibility: internal

depends_on: [ORG-*, TEA-*, SKI-*, CAP-*]
enables: [PER-*, QUA-*, DEV-*, SUC-*]

role_level: IC|Manager|Director|VP|C-Level
role_type: Full-time|Part-time|Contract|Consultant
career_track: Individual-Contributor|Management|Leadership
employment_type: Permanent|Temporary|Project-based

success_criteria:
  - "Role definition enables effective hiring and onboarding"
  - "Role responsibilities and authorities are clearly defined"
  - "Role requirements align with organizational needs"
  - "Role supports career development and progression"

assumptions:
  - "Role requirements are clearly understood and validated"
  - "Organizational context and structure are defined"
  - "Skill frameworks and competency models exist"

constraints: [Budget and organizational constraints]
standards: [HR and role management standards]
review_cycle: annually
---
```

## Content Structure Template

```markdown
# Role — {Role Title}

## Executive Summary
**Purpose:** {Brief description of role purpose and value}
**Level:** {IC|Manager|Director|VP|C-Level}
**Reports To:** {Manager role}
**Team:** {Primary team assignment}

## Role Overview

### Core Purpose
- **Mission:** {Why this role exists}
- **Impact:** {How this role contributes to business success}
- **Scope:** {Boundaries of responsibility and authority}

### Position Context
- **Department:** {Primary organizational unit}
- **Function:** {Primary business function}
- **Career Track:** {Individual Contributor|Management|Leadership}
- **Employment Type:** {Full-time|Part-time|Contract|Consultant}

## Responsibilities

### Primary Responsibilities
1. **{Responsibility Category}**
   - {Specific responsibility 1}
   - {Specific responsibility 2}
   - {Specific responsibility 3}

2. **{Responsibility Category}**
   - {Specific responsibility 1}
   - {Specific responsibility 2}

### Key Activities
| Activity | Frequency | Time Allocation | Priority |
|----------|-----------|-----------------|----------|
| {Activity} | {Daily|Weekly|Monthly} | {%} | {High|Medium|Low} |

### Decision Authority
```yaml
decisions:
  autonomous:
    - {Decision type the role can make independently}

  collaborative:
    - decision: {Decision requiring input}
      stakeholders: [role1, role2]

  escalated:
    - decision: {Decision requiring approval}
      approver: {role}
```

## Requirements

### Experience Requirements
- **Years of Experience:** {X-Y years}
- **Industry Experience:** {Specific industry requirements}
- **Functional Experience:** {Specific functional requirements}
- **Leadership Experience:** {Management/leadership requirements}

### Technical Skills
```yaml
required_skills:
  - skill: {skill-name}
    proficiency: {Beginner|Intermediate|Advanced|Expert}
    critical: {yes|no}

preferred_skills:
  - skill: {skill-name}
    proficiency: {level}
    value_add: {benefit-description}
```

### Competencies
- **Core Competencies:** {Essential behavioral competencies}
- **Leadership Competencies:** {Leadership and management competencies}
- **Functional Competencies:** {Role-specific competencies}

## Performance Framework

### Success Metrics
```yaml
individual_kpis:
  - metric: {metric-name}
    target: {target-value}
    measurement: {how-measured}
    frequency: {measurement-frequency}

team_contributions:
  - metric: {team-metric}
    contribution: {how-role-contributes}
```

### Career Development
- **Career Progression:** {Advancement opportunities}
- **Skill Development:** {Required skill development}
- **Learning Opportunities:** {Training and development resources}

## Working Relationships

### Internal Stakeholders
| Stakeholder | Relationship Type | Interaction Frequency | Purpose |
|-------------|-------------------|----------------------|---------|
| {Role/Team} | {Reports to|Collaborates|Supports} | {Daily|Weekly|Monthly} | {Purpose} |

### External Stakeholders
- **Customers:** {Customer interaction requirements}
- **Partners:** {Partner relationship management}
- **Vendors:** {Supplier relationship requirements}

## Work Environment

### Work Arrangements
- **Location:** {Office|Remote|Hybrid}
- **Travel Requirements:** {Travel expectations}
- **Work Schedule:** {Schedule flexibility and requirements}

### Tools and Resources
- **Technology:** {Required tools and systems}
- **Budget Authority:** {Financial authority limits}
- **Resource Access:** {Available resources and support}

## Validation
*Evidence that role definition enables effective hiring, clear performance expectations, and career development*
```

## Quality Standards

### Bronze Level (Minimum Viable)
- [ ] Clear role title and reporting relationship
- [ ] Basic responsibilities and requirements
- [ ] Essential qualifications and experience
- [ ] All required metadata fields completed

### Silver Level (Investment Ready)
- [ ] Detailed responsibility breakdown with decision authority
- [ ] Comprehensive skill and competency requirements
- [ ] Performance metrics and success criteria
- [ ] Career development framework

### Gold Level (Operational Excellence)
- [ ] Strategic role contribution and business impact
- [ ] Advanced competency modeling and assessment
- [ ] Comprehensive stakeholder relationship mapping
- [ ] Dynamic role evolution and adaptation framework

## Common Pitfalls

1. **Vague responsibilities**: Unclear or overlapping responsibilities
2. **Unrealistic requirements**: Excessive or conflicting skill requirements
3. **Poor performance definition**: Lack of clear success criteria
4. **Static role design**: Roles that don't evolve with business needs

## Success Patterns

1. **Clear value proposition**: Well-defined role impact with alignment to business objectives
2. **Growth-oriented design**: Clear career progression with skill building pathways
3. **Stakeholder integration**: Well-defined working relationships with clear collaboration expectations
4. **Performance-driven framework**: Measurable success criteria with regular feedback and development

## Relationship Guidelines

### Typical Dependencies
- **ORG (Organization Structure)**: Organizational design drives role structure and reporting
- **TEA (Team Structure)**: Team composition informs role design and collaboration
- **SKI (Skills Framework)**: Skill requirements drive role competency definitions
- **CAP (Capabilities)**: Organizational capabilities determine role requirements

### Typical Enablements
- **PER (Performance Specification)**: Role clarity drives performance achievement
- **QUA (Quality Specification)**: Role competencies drive quality standards
- **DEV (Development)**: Role definitions enable development and training programs
- **SUC (Succession Planning)**: Role clarity enables succession and career planning

## Document Relationships

This document type commonly relates to:

- **Depends on**: ORG (Organization Structure), TEA (Team Structure), SKI (Skills Framework), CAP (Capabilities)
- **Enables**: PER (Performance Specification), QUA (Quality Specification), DEV (Development), SUC (Succession Planning)
- **Informs**: Hiring strategy, performance management, compensation design
- **Guides**: Recruitment processes, onboarding programs, career development

## Validation Checklist

- [ ] Executive summary with clear purpose, level, reporting relationship, and team assignment
- [ ] Role overview with core purpose, impact, scope, and position context
- [ ] Responsibilities with primary duties, key activities, and decision authority framework
- [ ] Requirements with experience, technical skills, and competency specifications
- [ ] Performance framework with success metrics and career development pathways
- [ ] Working relationships with internal and external stakeholder mapping
- [ ] Work environment with arrangements, tools, and resource requirements
- [ ] Validation evidence of role effectiveness, hiring success, and performance clarity