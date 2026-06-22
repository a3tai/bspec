---
title: "SKI: Skills Framework"
description: "BSpec SKI document type specification"
---

# SKI: Skills Framework

::: tip Document Type
**Code**: SKI<br>
**Name**: Skills Framework<br>
**Domain**: Operations & Execution
:::

## Abstract

This specification defines the Skills Framework document type within the BSpec 1.0 Universal Business Specification Standard. It establishes normative requirements, structured templates, and implementation guidance for organizations documenting skills framework within the operations-execution domain. This specification enables systematic, machine-readable documentation that supports strategic planning, operational execution, and organizational alignment.

## Purpose and Scope

The Skills Framework defines systematic approaches to identifying, assessing, and developing organizational skills and competencies that enable strategic execution and competitive advantage. It establishes skill frameworks that optimize talent development, career planning, and organizational capability building.

## Document Metadata Schema

```yaml
---
id: SKI-{skill-category}
title: "Skills — {Skill Category or Framework Name}"
type: SKI
status: Draft|Review|Approved|Active|Deprecated
version: 1.0.0
owner: Skills-Owner|HR-Team
stakeholders: [hr-team, learning-development-team, managers, employees]
domain: operations
priority: Critical|High|Medium|Low
scope: skills-management
horizon: strategic
visibility: internal

depends_on: [ROL-*,TEA-*,ORG-*,CAP-*]
enables: [PER-*,QUA-*,DEV-*]

framework_scope: Technical|Functional|Behavioral|Leadership
skill_categories: [List of skill categories]
proficiency_levels: Basic|Intermediate|Advanced|Expert
assessment_method: Self|Peer|Manager|360|Competency-test

success_criteria:
  - "Skills framework enables effective talent development"
  - "Skill assessment is accurate and consistent"
  - "Skill gaps are identified and addressed"
  - "Skills align with business strategy and needs"

assumptions:
  - "Skill requirements are clearly defined and validated"
  - "Assessment methods are reliable and consistent"
  - "Development resources and programs are available"

constraints: [Resource and development constraints]
standards: [Skills management and development standards]
review_cycle: annually
---
```

## Content Structure Template

```markdown
# Skills — {Skill Category or Framework Name}

## Executive Summary
**Purpose:** {Brief description of skill framework purpose}
**Scope:** {Areas and domains covered}
**Application:** {How this framework is used in organization}

## Skills Framework

### Framework Structure
- **Categories:** {Major skill categories}
- **Levels:** {Proficiency levels used}
- **Assessment Method:** {How skills are evaluated}
- **Development Path:** {How skills are developed}

### Skill Categories

#### Technical Skills
    ```yaml
    technical_skills:
      programming:
        - skill: &#123;Programming Language&#125;
          description: &#123;Skill description&#125;
          levels: [Beginner, Intermediate, Advanced, Expert]
          assessment: &#123;Assessment method&#125;

      systems:
        - skill: &#123;System/Platform&#125;
          description: &#123;Skill description&#125;
          levels: [Basic, Proficient, Advanced, Master]
          assessment: &#123;Assessment method&#125;
    ```

#### Functional Skills
    ```yaml
    functional_skills:
      domain_expertise:
        - skill: &#123;Domain Knowledge&#125;
          description: &#123;Skill description&#125;
          levels: [Novice, Competent, Proficient, Expert]
          assessment: &#123;Assessment method&#125;

      process_skills:
        - skill: &#123;Process/Methodology&#125;
          description: &#123;Skill description&#125;
          levels: [Awareness, Working, Fluent, Master]
          assessment: &#123;Assessment method&#125;
    ```

#### Behavioral Skills
    ```yaml
    behavioral_skills:
      leadership:
        - skill: &#123;Leadership Competency&#125;
          description: &#123;Skill description&#125;
          levels: [Developing, Competent, Strong, Exemplary]
          assessment: &#123;Assessment method&#125;

      collaboration:
        - skill: &#123;Collaboration Skill&#125;
          description: &#123;Skill description&#125;
          levels: [Basic, Effective, Advanced, Exceptional]
          assessment: &#123;Assessment method&#125;
    ```

## Proficiency Levels

### Level Definitions
| Level | Description | Characteristics | Assessment Criteria |
|-------|-------------|-----------------|-------------------|
| **Beginner** | {Level description} | {Key characteristics} | {How to assess} |
| **Intermediate** | {Level description} | {Key characteristics} | {How to assess} |
| **Advanced** | {Level description} | {Key characteristics} | {How to assess} |
| **Expert** | {Level description} | {Key characteristics} | {How to assess} |

### Progression Criteria
- **Advancement Requirements:** {What's needed to move to next level}
- **Time Investment:** {Expected time to advance}
- **Assessment Methods:** {How progression is evaluated}

## Skills Assessment

### Assessment Framework
    ```yaml
    assessment_methods:
      self_assessment:
        frequency: &#123;How often&#125;
        tool: &#123;Assessment tool/process&#125;
        calibration: &#123;How accuracy is ensured&#125;

      peer_assessment:
        frequency: &#123;How often&#125;
        process: &#123;360-degree or targeted&#125;
        validation: &#123;How results are validated&#125;

      manager_assessment:
        frequency: &#123;How often&#125;
        criteria: &#123;Assessment criteria&#125;
        documentation: &#123;How results are recorded&#125;
    ```

### Skill Inventory
- **Current State Analysis:** {Organization's current skill levels}
- **Gap Analysis:** {Identified skill gaps}
- **Demand Forecasting:** {Future skill requirements}

## Development Framework

### Learning Pathways
    ```yaml
    development_paths:
      skill_category:
        - level: &#123;Target level&#125;
          methods: [Training, Mentoring, Projects, Certification]
          resources: [Course, Book, Platform, Expert]
          timeline: &#123;Expected duration&#125;
          validation: &#123;How to confirm achievement&#125;
    ```

### Development Methods
- **Formal Training:** {Structured learning programs}
- **On-the-job Learning:** {Experiential learning opportunities}
- **Mentoring:** {Mentorship and coaching programs}
- **Self-directed Learning:** {Individual learning resources}

### Career Integration
- **Role Mapping:** {How skills map to roles}
- **Career Paths:** {Skill progression for career advancement}
- **Specialization Tracks:** {Deep expertise development}

## Skills Planning

### Organizational Planning
- **Strategic Skills:** {Skills critical for business strategy}
- **Core Skills:** {Essential skills for operations}
- **Emerging Skills:** {Skills needed for future}

### Individual Planning
    ```yaml
    individual_development:
      current_skills:
        - skill: &#123;skill-name&#125;
          level: &#123;current-level&#125;
          evidence: &#123;Supporting evidence&#125;

      target_skills:
        - skill: &#123;skill-name&#125;
          target_level: &#123;desired-level&#125;
          timeline: &#123;target-date&#125;
          development_plan: &#123;Learning approach&#125;
    ```

### Team Planning
- **Skill Distribution:** {Current team skill profile}
- **Capability Gaps:** {Team skill gaps}
- **Development Priorities:** {Team learning priorities}

## Quality Assurance

### Validation Methods
- **Competency Testing:** {Formal skill validation}
- **Portfolio Review:** {Work product assessment}
- **Practical Application:** {Real-world demonstration}

### Continuous Improvement
- **Framework Updates:** {How framework evolves}
- **Calibration Process:** {Ensuring consistent assessment}
- **Effectiveness Measurement:** {Framework success metrics}

## Validation
*Evidence that skills framework enables effective talent development, accurate assessment, and strategic alignment*
```

## Quality Standards

### Bronze Level (Minimum Viable)
- [ ] Basic skill categories with clear definitions
- [ ] Simple proficiency levels (3-4 levels)
- [ ] Basic assessment approach
- [ ] All required metadata fields completed

### Silver Level (Investment Ready)
- [ ] Comprehensive skill framework with detailed competencies
- [ ] Well-defined progression criteria and assessment methods
- [ ] Individual development planning integration
- [ ] Learning pathway definition

### Gold Level (Operational Excellence)
- [ ] Advanced competency modeling with behavioral indicators
- [ ] Sophisticated assessment and validation framework
- [ ] Integration with career development and succession planning
- [ ] Data-driven skills optimization and forecasting

## Common Pitfalls

1. **Overly complex framework**: Too many skills or levels making framework unwieldy
2. **Inconsistent assessment**: Subjective or inconsistent skill evaluation
3. **Static framework**: Skills framework that doesn't evolve with business needs
4. **Poor integration**: Framework not connected to roles, careers, or development

## Success Patterns

1. **Business-aligned framework**: Skills directly tied to business strategy with regular strategic alignment
2. **Development-focused**: Clear pathways for skill development with integrated learning support
3. **Data-driven optimization**: Regular measurement and improvement with evidence-based investment decisions
4. **Practical application**: Framework designed for practical use with clear, actionable guidance

## Relationship Guidelines

### Typical Dependencies
- **ROL (Role Definition)**: Role requirements drive skill framework design and application
- **TEA (Team Structure)**: Team composition informs skill requirements and development priorities
- **ORG (Organization Structure)**: Organizational design drives skill framework scope and structure
- **CAP (Capabilities)**: Organizational capabilities determine required skills and competencies

### Typical Enablements
- **PER (Performance Specification)**: Skill development drives performance improvement and achievement
- **QUA (Quality Specification)**: Skill competencies drive quality standards and capabilities
- **DEV (Development)**: Skills framework enables development programs and career planning
- **Succession Planning**: Skill assessment enables succession and talent planning

## Document Relationships

This document type commonly relates to:

- **Depends on**: ROL (Role Definition), TEA (Team Structure), ORG (Organization Structure), CAP (Capabilities)
- **Enables**: PER (Performance Specification), QUA (Quality Specification), DEV (Development), SUC (Succession Planning)
- **Informs**: Hiring strategy, training programs, career development
- **Guides**: Skills assessment, development planning, talent management

## Validation Checklist

- [ ] Executive summary with clear purpose, scope, and application
- [ ] Skills framework with structure, categories, and comprehensive skill definitions
- [ ] Proficiency levels with clear definitions, characteristics, and progression criteria
- [ ] Skills assessment with framework, methods, and inventory analysis
- [ ] Development framework with learning pathways, methods, and career integration
- [ ] Skills planning with organizational, individual, and team planning approaches
- [ ] Quality assurance with validation methods and continuous improvement processes
- [ ] Validation evidence of framework effectiveness, assessment accuracy, and development success


---

## Related Documents

- [Back to Document Types](/docs/document-types)
- [Domain: Operations & Execution](/docs/domains/operations-execution)
- [Full Specification](/spec/v1-0-0)
