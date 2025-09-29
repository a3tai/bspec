# TON: Tone of Voice Document Type Specification

**Document Type Code:** TON
**Document Type Name:** Tone of Voice
**Domain:** Brand & Marketing
**Version:** 1.0.0
**Status:** Draft
**Last Updated:** 2025-09-28

## Purpose and Scope

The Tone of Voice document defines how the brand sounds in all communications, establishing voice characteristics that express brand personality and create consistent customer experiences. It provides guidelines that ensure authentic, recognizable, and engaging brand communication across all touchpoints and channels.

## Document Metadata Schema

```yaml
---
id: TON-{voice-area}
title: "Tone of Voice — {Voice Area or Application}"
type: TON
status: Draft|Review|Approved|Active|Deprecated
version: 1.0.0
owner: Brand-Manager|Communications-Team|Content-Team
stakeholders: [brand-team, marketing-team, content-team, customer-service]
domain: brand-marketing
priority: Critical|High|Medium|Low
scope: tone-of-voice
horizon: tactical
visibility: internal

depends_on: [BRD-*, MSG-*, PER-*, CUS-*]
enables: [CNT-*, CAM-*, SOC-*, SUP-*]

voice_characteristics: [3-5 primary personality traits]
communication_style: [formal, casual, technical, conversational, etc.]
language_guidelines: [vocabulary to use and avoid]
cultural_considerations: [how tone adapts across cultures/regions]
channel_variations: [how tone changes across communication channels]
voice_evolution: [how tone may develop as brand matures]

success_criteria:
  - "Tone of voice creates authentic and recognizable brand personality"
  - "Voice guidelines ensure consistent communication across channels"
  - "Tone effectively connects with target audiences and builds relationships"
  - "Voice framework enables scalable and adaptable brand communication"

assumptions:
  - "Brand personality and archetype clearly defined and validated"
  - "Target audience communication preferences well understood"
  - "Brand values and positioning provide clear voice direction"

constraints: [Communication style and cultural constraints]
standards: [Brand communication and voice standards]
review_cycle: annually
---
```

## Content Structure Template

```markdown
# Tone of Voice — {Voice Area or Application}

## Executive Summary
**Purpose:** {Brief description of tone of voice scope and objectives}
**Voice Characteristics:** {3-5 primary personality traits}
**Communication Style:** {Formal, casual, technical, conversational, etc.}
**Language Guidelines:** {Vocabulary to use and avoid}

## Tone of Voice Framework

### Voice Philosophy
- **Brand Personality Expression:** {How voice expresses brand personality and values}
- **Authentic Communication:** {Creating genuine and authentic brand conversations}
- **Relationship Building:** {Using voice to build customer relationships and trust}
- **Consistent Experience:** {Maintaining consistent voice across all interactions}

### Voice Foundation
```yaml
voice_strategy:
  personality_alignment: {How voice aligns with brand personality and archetype}
  audience_connection: {How voice connects with target audience preferences}
  differentiation_approach: {How voice differentiates from competitors}
  relationship_objectives: {Type of relationship voice aims to create}

voice_characteristics:
  primary_traits: [3-5 core personality traits defining voice]
  communication_style: {Overall approach to communication and interaction}
  emotional_tone: {Emotional qualities and feelings conveyed}
  expertise_level: {How knowledgeable and authoritative voice appears}

voice_principles:
  authenticity_standards: {Requirements for genuine and authentic communication}
  consistency_guidelines: {Maintaining consistent voice across contexts}
  adaptability_framework: {How voice adapts while maintaining core identity}
  evolution_pathway: {How voice can evolve as brand matures}
```

### Voice Identity
- **Voice Personality:** {Human characteristics that define how brand sounds}
- **Communication Approach:** {General approach to customer interaction}
- **Relationship Style:** {Type of relationship brand wants with customers}
- **Expertise Positioning:** {How brand positions its knowledge and authority}

## Voice Characteristics Definition

### Personality Dimensions
```yaml
voice_personality:
  formality_spectrum:
    position: {formal, semi-formal, casual, very-casual}
    rationale: {Why this level of formality serves brand and audience}
    examples: {Examples of appropriate formality level}
    boundaries: {Limits of acceptable formality range}

  authority_spectrum:
    position: {authoritative, confident, humble, supportive}
    rationale: {Why this authority level aligns with brand positioning}
    examples: {Examples of appropriate authority expression}
    boundaries: {Limits of acceptable authority range}

  emotion_spectrum:
    position: {rational, balanced, emotional, passionate}
    rationale: {Why this emotional level connects with audience}
    examples: {Examples of appropriate emotional expression}
    boundaries: {Limits of acceptable emotional range}

  complexity_spectrum:
    position: {simple, accessible, sophisticated, technical}
    rationale: {Why this complexity level serves audience needs}
    examples: {Examples of appropriate complexity level}
    boundaries: {Limits of acceptable complexity range}
```

### Character Traits
- **Primary Traits:** {3-4 main personality characteristics that define voice}
- **Supporting Traits:** {Additional characteristics that add depth to voice}
- **Situational Traits:** {How voice characteristics adapt to different contexts}
- **Anti-Traits:** {Characteristics that conflict with brand voice}

### Voice Archetypes
```yaml
archetype_expression:
  chosen_archetype: {Brand archetype that guides voice development}
  voice_manifestation: {How archetype personality appears in communication}
  communication_patterns: {Typical ways archetype communicates}
  relationship_approach: {How archetype builds relationships with others}

  archetype_examples:
    - situation: {specific communication context}
      archetype_response: {how archetype would naturally respond}
      voice_application: {how brand voice expresses archetype response}
      tone_guidance: {specific tone guidance for this situation}

  archetype_boundaries:
    authentic_expression: {Staying true to archetype while being authentic}
    situational_flexibility: {Adapting archetype expression to context}
    evolution_potential: {How archetype expression can develop over time}
    consistency_maintenance: {Maintaining archetype consistency across channels}
```

## Communication Guidelines

### Language and Vocabulary
```yaml
language_framework:
  preferred_vocabulary:
    - category: {industry_terms, emotional_words, action_words}
      words_phrases: [specific words and phrases to use]
      usage_context: {when and how to use these words}
      effectiveness_rationale: {why these words work for brand}

  avoided_vocabulary:
    - category: {jargon, negative_words, competitor_language}
      words_phrases: [specific words and phrases to avoid]
      avoidance_rationale: {why these words don't work for brand}
      alternative_suggestions: {better alternatives to use instead}

  industry_language:
    technical_terms: {Approach to using industry-specific terminology}
    jargon_policy: {When technical jargon is appropriate vs. plain language}
    explanation_strategy: {How to explain complex concepts clearly}
    accessibility_standards: {Ensuring language is accessible to all audiences}
```

### Writing Style Guidelines
- **Sentence Structure:** {Preferred sentence length and complexity}
- **Paragraph Length:** {Guidelines for paragraph structure and length}
- **Punctuation Style:** {Specific punctuation preferences and standards}
- **Formatting Approach:** {How to format text for clarity and personality}

### Grammar and Usage
```yaml
grammar_standards:
  style_preferences:
    contraction_usage: {When to use contractions for conversational tone}
    person_perspective: {First person, second person, third person usage}
    active_passive_voice: {Preference for active vs. passive voice}
    tense_consistency: {Maintaining consistent verb tenses}

  punctuation_guidelines:
    comma_style: {Oxford comma usage and other comma guidelines}
    quotation_marks: {When and how to use quotation marks}
    exclamation_points: {Guidelines for excitement and emphasis}
    ellipses_dashes: {Using ellipses and dashes for pause and emphasis}

  capitalization_rules:
    title_case: {When to use title case vs. sentence case}
    brand_terms: {Capitalization of brand-specific terms}
    emphasis_capitalization: {Using capitalization for emphasis}
    consistency_standards: {Maintaining consistent capitalization}
```

## Channel-Specific Applications

### Digital Channel Adaptations
```yaml
digital_voice:
  website_communication:
    voice_adaptation: {How voice adapts for website content}
    formality_level: {Appropriate formality for web content}
    engagement_approach: {How to engage visitors through voice}
    conversion_optimization: {Using voice to drive desired actions}

  social_media_voice:
    - platform: {linkedin, twitter, instagram, facebook}
      voice_characteristics: {How voice adapts to platform culture}
      content_approach: {Approach to content and engagement}
      interaction_style: {How to interact with followers and comments}
      hashtag_personality: {Personality in hashtag and emoji usage}

  email_communication:
    subject_line_voice: {Voice characteristics in email subjects}
    body_content_tone: {Tone for email body content}
    call_to_action_style: {Voice in calls to action}
    signature_personality: {Personality in email signatures}
```

### Traditional Channel Adaptations
- **Print Materials:** {Voice adaptation for brochures, flyers, advertisements}
- **Broadcast Media:** {Voice for radio, television, and podcast content}
- **Presentations:** {Voice for speaking engagements and presentations}
- **Packaging:** {Voice on product packaging and labeling}

### Interactive Channel Guidelines
```yaml
interactive_voice:
  customer_service:
    support_tone: {Voice for customer support interactions}
    problem_resolution: {Tone when addressing customer problems}
    satisfaction_follow_up: {Voice for satisfaction and follow-up}
    escalation_management: {Tone for handling escalated situations}

  sales_conversations:
    discovery_tone: {Voice during customer discovery conversations}
    presentation_style: {Tone for sales presentations and demos}
    objection_handling: {Voice when addressing customer concerns}
    closing_approach: {Tone for closing and commitment conversations}

  community_engagement:
    forum_participation: {Voice for forum and community participation}
    user_generated_content: {Tone when featuring customer content}
    feedback_response: {Voice when responding to customer feedback}
    crisis_communication: {Tone during challenging situations}
```

## Implementation and Consistency

### Voice Implementation Strategy
```yaml
implementation_approach:
  rollout_planning:
    phase_strategy: {Phased approach to implementing voice guidelines}
    priority_channels: {Most important channels for initial implementation}
    training_requirements: {Training needed for effective voice implementation}
    resource_development: {Resources needed to support implementation}

  stakeholder_enablement:
    voice_training: {Training stakeholders on voice characteristics and usage}
    example_library: {Library of voice examples and templates}
    review_processes: {Processes for reviewing voice implementation}
    feedback_mechanisms: {Collecting feedback on voice effectiveness}

  quality_assurance:
    consistency_monitoring: {Ensuring consistent voice across channels}
    effectiveness_measurement: {Measuring voice effectiveness and impact}
    compliance_auditing: {Auditing compliance with voice guidelines}
    improvement_identification: {Identifying opportunities for voice improvement}
```

### Consistency Framework
- **Review Processes:** {How voice usage is reviewed and approved}
- **Training Programs:** {Training team members on voice guidelines}
- **Template Development:** {Creating templates that embody brand voice}
- **Feedback Systems:** {Systems for improving voice implementation}

### Voice Evolution Management
```yaml
voice_evolution:
  evolution_triggers:
    brand_maturity: {How voice evolves as brand matures}
    audience_changes: {Adapting voice to changing audience preferences}
    market_evolution: {Voice adaptation to market and industry changes}
    competitive_response: {Voice adjustments based on competitive landscape}

  evolution_process:
    assessment_methodology: {How to assess need for voice evolution}
    testing_approach: {Testing voice changes before full implementation}
    implementation_strategy: {Rolling out voice evolution across channels}
    change_management: {Managing stakeholder transition to new voice}

  version_control:
    voice_versioning: {Managing different versions of voice guidelines}
    historical_tracking: {Tracking voice evolution over time}
    rollback_procedures: {Procedures for reverting voice changes if needed}
    archive_management: {Managing archived voice guidelines and examples}
```

## Measurement and Optimization

### Voice Performance Metrics
```yaml
voice_metrics:
  recognition_metrics:
    - metric: {Voice Recognition Rate}
      measurement: {Ability of audience to recognize brand voice}
      assessment: {Testing voice recognition in blind studies}

    - metric: {Voice Consistency Score}
      measurement: {Consistency of voice across different channels}
      tracking: {Regular assessment of voice consistency}

  engagement_metrics:
    - metric: {Communication Engagement}
      measurement: {How well voice drives audience engagement}
      channels: {Engagement measurement across different channels}

    - metric: {Relationship Quality}
      assessment: {How voice contributes to customer relationship quality}
      indicators: {Indicators of strong voice-driven relationships}

  effectiveness_metrics:
    message_clarity: {How well voice communicates intended messages}
    audience_connection: {How effectively voice connects with target audience}
    brand_alignment: {How well voice aligns with overall brand strategy}
    competitive_differentiation: {How voice differentiates from competitors}
```

### Optimization Process
- **Performance Analysis:** {Regular analysis of voice performance and effectiveness}
- **Audience Feedback:** {Collecting and analyzing audience feedback on voice}
- **Competitive Benchmarking:** {Comparing voice effectiveness against competitors}
- **Continuous Refinement:** {Process for continuously improving voice guidelines}

### Voice Research and Testing
```yaml
voice_research:
  audience_testing:
    voice_preference_research: {Understanding audience voice preferences}
    message_clarity_testing: {Testing how clearly voice communicates}
    emotional_response_measurement: {Measuring emotional response to voice}
    relationship_perception_assessment: {How voice affects relationship perception}

  competitive_analysis:
    voice_landscape_analysis: {Analysis of competitive voice characteristics}
    differentiation_assessment: {How brand voice differs from competitors}
    opportunity_identification: {Voice opportunities in competitive landscape}
    trend_monitoring: {Monitoring voice trends in industry and market}

  effectiveness_measurement:
    communication_impact: {Measuring impact of voice on communication effectiveness}
    business_outcome_correlation: {Correlating voice with business outcomes}
    channel_performance: {Voice performance across different channels}
    roi_assessment: {Return on investment in voice development and implementation}
```

## Validation
*Evidence that tone of voice creates authentic personality, ensures consistent communication, and effectively connects with audiences*

## Quality Standards

### Bronze Level (Minimum Viable)
- [ ] Basic tone of voice with core characteristics and simple guidelines
- [ ] Basic language guidelines and communication standards
- [ ] Simple channel adaptations and usage examples
- [ ] All required metadata fields completed

### Silver Level (Investment Ready)
- [ ] Comprehensive tone of voice with detailed characteristics and guidelines
- [ ] Structured channel adaptations with implementation strategy
- [ ] Active voice measurement with performance tracking and optimization
- [ ] Regular voice review and evolution planning

### Gold Level (Operational Excellence)
- [ ] Advanced voice capabilities with sophisticated guidelines and adaptations
- [ ] Comprehensive voice ecosystem with integrated implementation and governance
- [ ] Voice excellence with industry recognition and differentiation
- [ ] Strategic voice management driving brand recognition and customer connection

## Common Pitfalls

1. **Generic voice**: Voice characteristics that don't meaningfully differentiate brand
2. **Inconsistent application**: Voice not consistently applied across channels and interactions
3. **Audience mismatch**: Voice that doesn't resonate with target audience preferences
4. **Artificial personality**: Voice that feels forced or inauthentic rather than natural

## Success Patterns

1. **Authentic expression**: Voice that genuinely reflects brand personality and values
2. **Audience alignment**: Voice that resonates with target audience communication preferences
3. **Consistent implementation**: Voice consistently applied across all channels and touchpoints
4. **Relationship building**: Voice that effectively builds customer relationships and trust

## Relationship Guidelines

### Typical Dependencies
- **BRD (Brand Strategy)**: Brand personality and archetype guide voice development
- **MSG (Messaging)**: Message framework influences voice characteristics
- **PER (Personas)**: Customer personas inform voice preferences and adaptation
- **CUS (Customer)**: Customer insights drive voice design and optimization

### Typical Enablements
- **CNT (Content)**: Voice guidelines enable consistent content tone and style
- **CAM (Campaigns)**: Voice framework guides campaign communication style
- **SOC (Social Media)**: Voice standards enable effective social media communication
- **SUP (Customer Support)**: Voice guidelines enable consistent support interactions

## Document Relationships

This document type commonly relates to:

- **Depends on**: BRD (Brand Strategy), MSG (Messaging), PER (Personas), CUS (Customer)
- **Enables**: CNT (Content), CAM (Campaigns), SOC (Social Media), SUP (Customer Support)
- **Informs**: Content creation, customer communications, brand training
- **Guides**: Communication style, customer interactions, brand expression

## Validation Checklist

- [ ] Executive summary with clear purpose, characteristics, style, and guidelines
- [ ] Voice framework with philosophy, foundation, and identity definition
- [ ] Characteristics definition with personality dimensions, traits, and archetypes
- [ ] Communication guidelines with language, style, and grammar standards
- [ ] Channel applications with digital, traditional, and interactive adaptations
- [ ] Implementation and consistency with strategy, framework, and evolution management
- [ ] Measurement and optimization with metrics, optimization, and research processes
- [ ] Validation evidence of authentic personality, consistent communication, and audience connection