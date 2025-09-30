# EMP: Empathy Maps Document Type Specification

**Document Type Code:** EMP
**Document Type Name:** Empathy Maps
**Domain:** Customer & Value
**Version:** 1.0.0
**Status:** Draft
**Last Updated:** 2025-09-28

## Abstract

This specification defines the Empathy Maps document type within the BSpec 1.0 Universal Business Specification Standard. It establishes normative requirements, structured templates, and implementation guidance for organizations documenting empathy maps within the customer-value domain. This specification enables systematic, machine-readable documentation that supports strategic planning, operational execution, and organizational alignment.

## Purpose and Scope

The Empathy Maps document captures deep understanding of customer thoughts, feelings, behaviors, and environment. It provides holistic view of customer experience and emotional context that drives behavior.

## Document Metadata Schema

```yaml
---
id: EMP-{empathy-map-identifier}
title: "Empathy Map: [Persona/Scenario Name]"
type: EMP
status: Draft|Review|Accepted|Deprecated
version: 1.0.0
owner: Design-Research|Customer-Experience
stakeholders: [design-team, product-team, marketing-team]
domain: customer
priority: Medium|High
scope: persona|scenario
horizon: current
visibility: internal

depends_on: [PER-*, JTB-*]
enables: [CJM-*, UXD-*, REL-*]

persona_focus: [Primary persona this map represents]
scenario_context: [Specific scenario or situation]
research_depth: High|Medium|Basic

success_criteria:
  - "Empathy map reflects authentic customer experience"
  - "Insights drive design and experience decisions"
  - "Team develops genuine customer empathy"
  - "Customer validation confirms accuracy"

assumptions:
  - "Customer experience is observable and mappable"
  - "Empathy drives better product decisions"
  - "Research provides sufficient insight depth"

research_methodology:
  - "Customer interviews and observation"
  - "Journey mapping sessions"
  - "Day-in-the-life studies"
  - "Emotional response research"

observation_hours: X # Total customer observation time
interview_count: X # Number of customer interviews
review_cycle: quarterly
---
```

## Content Structure Template

```markdown
# Empathy Map: [Persona/Scenario Name]

## Overview
*Summary of the customer empathy map and its purpose*

## Empathy Map Context

### Subject Definition
**Primary Persona**: [Persona name and key characteristics]
**Scenario Focus**: [Specific situation or context being mapped]
**Time Frame**: [Duration or period covered by map]
**Environmental Context**: [Physical and social environment]

### Research Foundation
**Research Methods Used**
- Customer interviews (X conducted)
- Observational studies (X hours)
- Day-in-the-life shadowing
- Digital behavior tracking
- Survey emotional response data

**Research Participants**
- Number of customers studied: X
- Demographic representation
- Geographic distribution
- Experience level variation
- Role and responsibility diversity

**Research Quality**
- Depth of engagement with each participant
- Triangulation across multiple sources
- Validation through follow-up research
- Recency and relevance of data

## THINKS (Cognitive Layer)

### Beliefs and Assumptions
**Core Beliefs**
- Fundamental beliefs about work/life/industry
- Assumptions about how things should work
- Mental models and frameworks used
- Deeply held convictions and principles

**Situational Beliefs**
- Beliefs about current situation
- Assumptions about available options
- Expectations about outcomes
- Predictions about future states

**Self-Beliefs**
- Beliefs about own capabilities
- Assumptions about personal limitations
- Confidence levels and self-efficacy
- Identity and role perceptions

### Thought Patterns
**Decision-Making Thoughts**
- How they analyze options
- What criteria they consider important
- Risk assessment patterns
- Trade-off evaluation processes

**Problem-Solving Thoughts**
- How they approach challenges
- What resources they consider
- Systematic vs. intuitive approaches
- Learning and adaptation patterns

**Planning Thoughts**
- How they think about the future
- Goal-setting mental processes
- Timeline and milestone thinking
- Resource allocation considerations

### Mental Preoccupations
**Constant Concerns**
- Recurring thoughts and worries
- Background mental processing
- Persistent questions and uncertainties
- Ongoing evaluation and monitoring

**Situational Focus**
- Current priorities and attention
- Immediate concerns and considerations
- Context-specific thought patterns
- Task-related mental focus

## FEELS (Emotional Layer)

### Primary Emotions
**Dominant Emotional States**
- Most frequent emotional experiences
- Baseline emotional patterns
- Characteristic moods and feelings
- Emotional stability or volatility

**Emotional Triggers**
- Events that cause strong reactions
- Situational factors that shift emotions
- Relationship dynamics that affect feelings
- Success and failure emotional responses

**Emotional Progression**
- How emotions change over time
- Emotional journey patterns
- Recovery and resilience patterns
- Emotional learning and adaptation

### Emotional Needs
**Security Needs**
- Need for predictability and control
- Safety and risk management concerns
- Stability and continuity requirements
- Trust and reliability expectations

**Achievement Needs**
- Need for accomplishment and success
- Recognition and validation desires
- Progress and improvement motivations
- Mastery and competence aspirations

**Connection Needs**
- Relationship and belonging requirements
- Social recognition and status needs
- Community and support desires
- Communication and understanding needs

### Emotional Barriers
**Fear-Based Barriers**
- Specific fears and anxieties
- Risk aversion patterns
- Uncertainty discomfort
- Change resistance factors

**Confidence Barriers**
- Self-doubt and insecurity
- Competence concerns
- Performance anxiety
- Comparison and inadequacy feelings

**Social Barriers**
- Relationship concerns
- Status and reputation worries
- Social acceptance fears
- Judgment and criticism sensitivity

## SEES (Environmental Layer)

### Physical Environment
**Workspace/Environment**
- Physical space characteristics
- Layout and organization patterns
- Tools and equipment present
- Environmental conditions and constraints

**Visual Information**
- Information displays and dashboards
- Documents and reference materials
- Digital interfaces and screens
- Visual cues and indicators

**Social Environment**
- People present and their roles
- Team dynamics and interactions
- Organizational culture manifestations
- Communication patterns observed

### Information Environment
**Information Sources**
- Primary information channels used
- Trusted sources and authorities
- Peer and colleague communications
- Formal and informal information flows

**Information Quality**
- Accuracy and reliability perceptions
- Timeliness and currency concerns
- Completeness and detail expectations
- Accessibility and usability experiences

**Information Overload**
- Volume and complexity challenges
- Filtering and prioritization needs
- Signal vs. noise problems
- Attention and focus difficulties

### Market Environment
**Competitive Landscape**
- Competitor actions and responses
- Market changes and trends
- Industry dynamics and forces
- Economic conditions and pressures

**Stakeholder Environment**
- Customer behavior and expectations
- Partner actions and communications
- Regulatory and compliance factors
- Investor and board pressures

## SAYS (Communication Layer)

### Verbal Communication
**Frequent Phrases**
- Common expressions and language used
- Professional jargon and terminology
- Emotional expressions and reactions
- Requests and question patterns

**Communication Style**
- Formal vs. informal communication
- Direct vs. indirect expression
- Assertive vs. collaborative approach
- Detailed vs. high-level preference

**Value Expressions**
- How they articulate what matters
- Language used to describe success
- Expressions of priorities and concerns
- Communication of needs and desires

### Written Communication
**Documentation Patterns**
- Types of written communication created
- Formality and structure preferences
- Detail level and comprehensiveness
- Organization and presentation style

**Digital Communication**
- Email and messaging patterns
- Social media expression style
- Platform-specific communication
- Digital interaction preferences

### Influence Communication
**Persuasion Patterns**
- How they try to influence others
- Arguments and evidence they use
- Emotional appeals they make
- Logical structures they employ

**Feedback Expression**
- How they give and receive feedback
- Criticism and praise patterns
- Suggestion and improvement communication
- Conflict resolution communication

## DOES (Behavioral Layer)

### Observable Actions
**Routine Behaviors**
- Daily habits and patterns
- Workflow and process behaviors
- Decision-making actions
- Problem-solving behaviors

**Interaction Behaviors**
- How they engage with others
- Communication initiation patterns
- Collaboration and teamwork actions
- Leadership and followership behaviors

**Tool Usage Behaviors**
- How they use technology and tools
- Adaptation and learning behaviors
- Efficiency and optimization actions
- Troubleshooting and enable-seeking

### Behavioral Patterns
**Work Patterns**
- Productivity rhythms and cycles
- Focus and attention patterns
- Break and recovery behaviors
- Task switching and prioritization

**Learning Patterns**
- How they acquire new information
- Skill development approaches
- Knowledge sharing behaviors
- Expertise building actions

**Stress Response Patterns**
- How they react under pressure
- Coping and adaptation mechanisms
- Support-seeking behaviors
- Recovery and resilience actions

### Behavioral Contradictions
**Stated vs. Actual Behavior**
- Gaps between intentions and actions
- Inconsistencies in behavior patterns
- Context-dependent behavioral changes
- Unconscious vs. conscious behaviors

**Social vs. Private Behavior**
- Public presentation vs. private reality
- Role-based behavioral variations
- Social influence on behavior
- Authentic vs. performed behaviors

## PAINS (Challenges Layer)

### Frustrations and Obstacles
**Daily Frustrations**
- Recurring irritations and annoyances
- Process inefficiencies experienced
- Tool and technology limitations
- Communication and coordination problems

**Strategic Obstacles**
- Long-term challenges and barriers
- Systemic problems encountered
- Resource and capability constraints
- Organizational and structural impediments

**Emotional Frustrations**
- Stress and anxiety sources
- Confidence and competence challenges
- Relationship and social difficulties
- Recognition and appreciation gaps

### Fears and Anxieties
**Performance Fears**
- Fear of failure or mistakes
- Anxiety about meeting expectations
- Concerns about competence
- Worry about judgment and evaluation

**Change Fears**
- Resistance to new approaches
- Uncertainty about future states
- Loss and disruption concerns
- Adaptation and learning anxieties

**Social Fears**
- Relationship and interpersonal concerns
- Status and reputation worries
- Acceptance and belonging fears
- Conflict and confrontation avoidance

## GAINS (Aspirations Layer)

### Desires and Wants
**Achievement Desires**
- Success and accomplishment goals
- Recognition and reward aspirations
- Progress and improvement wants
- Mastery and expertise ambitions

**Experience Desires**
- Quality of experience improvements
- Ease and convenience wants
- Satisfaction and fulfillment goals
- Enjoyment and engagement aspirations

**Relationship Desires**
- Connection and belonging wants
- Trust and reliability aspirations
- Collaboration and support goals
- Communication and understanding desires

### Success Measures
**Personal Success Indicators**
- How they define personal success
- Metrics and milestones they track
- Recognition and validation needs
- Growth and development measures

**Professional Success Indicators**
- Career advancement criteria
- Performance measurement standards
- Industry recognition goals
- Expertise and reputation building

**Organizational Success Indicators**
- Team and department success metrics
- Company performance criteria
- Stakeholder satisfaction measures
- Market and competitive success

## Empathy Synthesis

### Key Insights
**Emotional Core**
- Central emotional experiences
- Driving emotional needs
- Emotional barriers to address
- Emotional opportunities to create

**Behavioral Drivers**
- What motivates action
- What inhibits or prevents action
- Decision-making drivers
- Change and adaptation factors

**Value Creation Opportunities**
- Unmet needs to address
- Pain points to solve
- Gains to amplify
- Experiences to improve

### Design Implications
**Experience Design**
- How insights inform experience design
- Emotional considerations for design
- Behavioral design opportunities
- Environmental design factors

**Product Design**
- Feature and capability implications
- User interface considerations
- Interaction design insights
- Performance requirements

**Service Design**
- Support and service implications
- Communication design needs
- Relationship design opportunities
- Touchpoint optimization

### Strategy Implications
**Positioning Strategy**
- How empathy informs positioning
- Emotional positioning opportunities
- Value proposition refinement
- Competitive differentiation

**Communication Strategy**
- Message development insights
- Channel selection guidance
- Tone and style implications
- Content strategy direction

## Validation and Evolution

### Empathy Validation
**Customer Feedback**
- Direct validation from customers
- Accuracy confirmation methods
- Insight validation approaches
- Correction and refinement feedback

**Behavioral Validation**
- Observed behavior confirmation
- Action pattern verification
- Decision-making validation
- Outcome achievement confirmation

### Empathy Evolution
**Change Indicators**
- Signals that empathy map needs updating
- Environmental changes affecting customers
- Behavioral pattern evolution
- Emotional state shifts

**Update Process**
- Regular empathy map review cycle
- Research refresh methodology
- Validation and correction process
- Stakeholder update communication

## Usage Guidelines

### Team Application
**Design Process Integration**
- How teams use empathy maps in design
- Decision-making guidance
- Evaluation criteria development
- Solution validation approaches

**Stakeholder Alignment**
- Empathy sharing with stakeholders
- Common understanding development
- Priority setting guidance
- Investment decision support

### Empathy Maintenance
**Living Document Approach**
- Regular updates and validation
- New insight integration
- Accuracy maintenance
- Relevance preservation

**Cross-Functional Sharing**
- Empathy map distribution
- Training and education
- Reference and application
- Continuous reinforcement

## Validation
*Evidence that empathy map accurately represents customer experience*
```

## Quality Standards

### Bronze Level (Minimum Viable)
- [ ] Six empathy layers documented (THINKS, FEELS, SEES, SAYS, DOES, PAINS/GAINS)
- [ ] Primary persona and scenario context clearly defined
- [ ] Basic research foundation with customer interviews/observation
- [ ] All required metadata fields completed

### Silver Level (Investment Ready)
- [ ] Comprehensive empathy mapping across all six layers with detailed insights
- [ ] Research foundation with multiple validation methods and quality assessment
- [ ] Empathy synthesis with key insights and design implications
- [ ] Strategy implications for positioning and communication
- [ ] Usage guidelines for team application and stakeholder alignment
- [ ] Validation and evolution framework

### Gold Level (Operational Excellence)
- [ ] Dynamic empathy tracking with real-time customer insight updates
- [ ] Advanced empathy analytics with behavioral pattern recognition
- [ ] Continuous empathy validation through ongoing customer engagement
- [ ] Empathy-driven design and strategy optimization
- [ ] Cross-functional empathy coordination and decision integration
- [ ] Automated empathy evolution monitoring and insight generation

## Common Pitfalls

1. **Assumption-based**: Creating empathy maps based on internal assumptions rather than customer research
2. **Superficial mapping**: Filling sections with obvious or generic insights rather than deep understanding
3. **Static documentation**: Creating empathy maps once without ongoing validation and updates
4. **Single perspective**: Mapping from company viewpoint rather than authentic customer experience

## Success Patterns

1. **Research-grounded**: Based on actual customer interviews, observation, and behavioral data
2. **Emotionally authentic**: Captures genuine emotional experience and psychological drivers
3. **Actionable insights**: Provides specific guidance for design, product, and strategy decisions
4. **Continuously validated**: Regularly updated and confirmed through ongoing customer engagement

## Relationship Guidelines

### Typical Dependencies
- **PER (Personas)**: Personas provide the foundation for empathy mapping subjects
- **JTB (Jobs-to-be-Done)**: Jobs provide context for customer goals and outcomes

### Typical Enablements
- **CJM (Customer Journey Maps)**: Empathy insights inform journey emotional states
- **UXD (User Experience Design)**: Empathy guides experience design decisions
- **REL (Relationships)**: Empathy informs relationship management strategies

## Document Relationships

This document type commonly relates to:

- **Depends on**: PER (Personas), JTB (Jobs-to-be-Done)
- **Enables**: CJM (Customer Journey Maps), UXD (User Experience Design), REL (Relationships)
- **Informs**: Product design, user experience strategy, customer communication
- **Guides**: Feature development, service design, relationship building

## Validation Checklist

- [ ] Six empathy layers (THINKS, FEELS, SEES, SAYS, DOES, PAINS/GAINS) comprehensively documented
- [ ] Subject definition with persona focus and scenario context
- [ ] Research foundation with multiple methods and quality assessment
- [ ] Cognitive layer with beliefs, thought patterns, and mental preoccupations
- [ ] Emotional layer with emotions, needs, and barriers
- [ ] Environmental layer covering physical, information, and market context
- [ ] Communication layer with verbal, written, and influence patterns
- [ ] Behavioral layer with observable actions, patterns, and contradictions
- [ ] Challenge and aspiration layers with detailed analysis
- [ ] Empathy synthesis with key insights and implications
- [ ] Validation and evolution framework with ongoing accuracy confirmation
- [ ] Usage guidelines for team application and maintenance