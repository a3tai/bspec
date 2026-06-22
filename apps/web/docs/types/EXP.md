---
title: "EXP: Experimentation"
description: "BSpec EXP document type specification"
---

# EXP: Experimentation

::: tip Document Type
**Code**: EXP<br>
**Name**: Experimentation<br>
**Domain**: Growth & Innovation
:::

## Abstract

This specification defines the Experimentation document type within the BSpec 1.0 Universal Business Specification Standard. It establishes normative requirements, structured templates, and implementation guidance for organizations documenting experimentation within the growth-innovation domain. This specification enables systematic, machine-readable documentation that supports strategic planning, operational execution, and organizational alignment.

## Framework and Attribution

Experimentation methods explicitly referenced in this template include Lean Startup and Design Thinking.
Use these frameworks under their respective usage terms and add notice where frameworks are published as named methods.

## Purpose and Scope

The Experimentation document governs controlled validation of priority hypotheses before deeper productization. It is the short-cycle mechanism feeding `IGN` and `FUT`, while `INN` and `RND` provide strategic direction and technical foundation.

## Document Metadata Schema

```yaml
---
id: EXP-{experimentation-area}
title: "Experimentation — {Experimentation Area or Framework}"
type: EXP
status: Draft|Review|Approved|Active|Deprecated
attribution_required: true
source_frameworks:
  - "Eric Ries - Lean Startup"
  - "IDEO / d.school - Design Thinking"
version: 1.0.0
owner: Experimentation-Team|Data-Science-Team|Product-Team
stakeholders: [experimentation-team, product-teams, data-science, executives]
domain: growth-innovation
priority: Critical|High|Medium|Low
scope: experimentation-framework
horizon: tactical
visibility: internal

depends_on: [INN-*,LEA-*,DAT-*,ANA-*]
enables: [PRD-*,SVC-*,STR-*,IGN-*]

experimentation_methodology: Scientific|Lean-startup|Design-thinking|Hybrid
hypothesis_framework: Falsifiable|Testable|Measurable|Time-bound
validation_approach: Quantitative|Qualitative|Mixed-methods
risk_tolerance: Conservative|Balanced|Aggressive

success_criteria:
  - "Experimentation framework enables rapid hypothesis testing and validation"
  - "Experimental insights drive evidence-based business decisions"
  - "Experimentation culture reduces risk and accelerates learning"
  - "Experimental processes scale across organization and initiatives"

assumptions:
  - "Organization values evidence-based decision-making over intuition"
  - "Leadership supports experimentation and tolerates failures"
  - "Data infrastructure supports experimental measurement and analysis"

constraints: [Resource and time constraints]
standards: [Experimental design and statistical standards]
review_cycle: monthly
---
```

## Content Structure Template

```markdown
# Experimentation — {Experimentation Area or Framework}

## Executive Summary
**Purpose:** {Brief description of experimentation scope and objectives}
**Experimentation Methodology:** {Scientific|Lean-startup|Design-thinking|Hybrid}
**Hypothesis Framework:** {Falsifiable|Testable|Measurable|Time-bound}
**Validation Approach:** {Quantitative|Qualitative|Mixed-methods}

## Experimentation Framework

### Experimentation Philosophy
- **Evidence-Based Decisions:** {Decisions based on experimental evidence rather than opinions}
- **Rapid Learning:** {Fast experimentation cycles to accelerate learning and adaptation}
- **Fail Fast, Learn Faster:** {Embracing failure as learning opportunity and iteration}
- **Systematic Testing:** {Structured approach to hypothesis generation and testing}

### Experimentation Strategy
    ```yaml
    experimentation_methodology:
      experimental_design: [randomized_controlled_trials, a_b_testing, multivariate_testing]
      hypothesis_framework: &#123;Structured approach to hypothesis formulation and testing&#125;
      measurement_protocols: &#123;Standardized measurement and data collection procedures&#125;
      validation_criteria: &#123;Clear criteria for experiment success and failure&#125;

    experimentation_governance:
      oversight_structure: &#123;Experimentation governance and quality assurance&#125;
      ethics_framework: &#123;Ethical guidelines for experimentation&#125;
      resource_allocation: &#123;Resource allocation for experimentation activities&#125;
    ```

### Experimentation Scope
- **Product Experimentation:** {Testing product features and user experiences}
- **Marketing Experimentation:** {Testing marketing messages and campaigns}
- **Business Model Experimentation:** {Testing business model assumptions and approaches}
- **Operational Experimentation:** {Testing operational processes and improvements}

## Experimental Design and Methodology

### Hypothesis Development
    ```yaml
    hypothesis_framework:
      hypothesis_generation:
        assumption_identification: &#123;Systematic identification of business assumptions&#125;
        risk_prioritization: &#123;Prioritizing assumptions by risk and impact&#125;
        hypothesis_formulation: &#123;Structured hypothesis formulation process&#125;
        testability_assessment: &#123;Assessment of hypothesis testability and feasibility&#125;

      hypothesis_structure:
        if_then_format: &#123;Clear if-then hypothesis structure&#125;
        measurable_outcomes: &#123;Specific, measurable predicted outcomes&#125;
        success_criteria: &#123;Clear criteria for hypothesis validation or refutation&#125;
        time_bounds: &#123;Specific time frames for hypothesis testing&#125;

      hypothesis_evaluation:
        feasibility_assessment: &#123;Technical and resource feasibility evaluation&#125;
        risk_assessment: &#123;Risk evaluation for experimental interventions&#125;
        impact_estimation: &#123;Estimated impact of hypothesis validation&#125;
        priority_ranking: &#123;Prioritization of hypotheses for testing&#125;
    ```

### Experimental Design
- **Randomized Controlled Trials:** {Gold standard experimental design with control groups}
- **A/B Testing:** {Simple two-variant testing for digital experiments}
- **Multivariate Testing:** {Testing multiple variables simultaneously}
- **Quasi-Experimental Design:** {Experimental design when randomization isn't possible}

### Measurement and Analytics
    ```yaml
    measurement_framework:
      metric_selection:
        primary_metrics: &#123;Key metrics that determine experiment success&#125;
        secondary_metrics: &#123;Additional metrics for comprehensive evaluation&#125;
        guardrail_metrics: &#123;Metrics to ensure experiments don't cause harm&#125;
        leading_indicators: &#123;Early indicators of experiment direction&#125;

      data_collection:
        data_sources: &#123;Sources of experimental data and measurement&#125;
        collection_methods: &#123;Methods for collecting experimental data&#125;
        data_quality: &#123;Data quality standards and validation procedures&#125;
        real_time_monitoring: &#123;Real-time monitoring of experiment progress&#125;

      statistical_analysis:
        analysis_methods: &#123;Statistical methods for experiment analysis&#125;
        significance_testing: &#123;Statistical significance testing procedures&#125;
        effect_size_calculation: &#123;Measurement of practical significance&#125;
        confidence_intervals: &#123;Confidence interval calculation and interpretation&#125;
    ```

## Experiment Portfolio Management

### Portfolio Framework
    ```yaml
    experiment_portfolio:
      portfolio_strategy:
        experiment_allocation: &#123;Resource allocation across experiment types&#125;
        risk_distribution: &#123;Distribution of experiments across risk levels&#125;
        time_horizon_mix: &#123;Mix of short-term and long-term experiments&#125;
        domain_coverage: &#123;Experiment coverage across business domains&#125;

      experiment_classification:
        - experiment_name: &#123;Experiment identification and naming&#125;
          hypothesis: &#123;Core hypothesis being tested&#125;
          methodology: &#123;Experimental methodology and design&#125;
          success_metrics: &#123;Primary and secondary success metrics&#125;
          duration: &#123;Experiment duration and timeline&#125;
          resource_cost: &#123;Resource requirements and costs&#125;
          status: &#123;planned|running|completed|cancelled&#125;
          results: &#123;Experiment outcomes and learnings&#125;

      portfolio_optimization:
        experiment_prioritization: &#123;Framework for prioritizing experiments&#125;
        resource_optimization: &#123;Optimization of experimental resource allocation&#125;
        learning_maximization: &#123;Maximizing learning per experiment dollar&#125;
        portfolio_balance: &#123;Balancing portfolio across dimensions&#125;
    ```

### Experiment Lifecycle
- **Planning Phase:** {Experiment design, resource planning, and approval}
- **Setup Phase:** {Experiment implementation and infrastructure preparation}
- **Execution Phase:** {Running experiment and monitoring progress}
- **Analysis Phase:** {Data analysis and insight generation}

### Portfolio Performance
    ```yaml
    portfolio_performance:
      experiment_metrics:
        experiment_velocity: &#123;Number of experiments completed per period&#125;
        success_rate: &#123;Percentage of experiments achieving success criteria&#125;
        learning_rate: &#123;Learning generated per experiment&#125;
        cost_efficiency: &#123;Cost per experiment and cost per insight&#125;

      business_impact:
        decision_influence: &#123;Experiments that influenced business decisions&#125;
        implementation_rate: &#123;Rate of experiment insights implementation&#125;
        business_value: &#123;Business value generated from experiments&#125;
        risk_mitigation: &#123;Risks avoided through experimental validation&#125;

      capability_metrics:
        experimentation_maturity: &#123;Organizational experimentation maturity&#125;
        experiment_quality: &#123;Quality of experimental design and execution&#125;
        learning_integration: &#123;Integration of learnings into business&#125;
        culture_adoption: &#123;Adoption of experimentation culture&#125;
    ```

## Experimentation Process and Operations

### Experiment Process
    ```yaml
    experimentation_process:
      experiment_planning:
        hypothesis_development: &#123;Systematic hypothesis development process&#125;
        experimental_design: &#123;Rigorous experimental design procedures&#125;
        resource_planning: &#123;Resource and timeline planning&#125;
        risk_assessment: &#123;Experimental risk assessment and mitigation&#125;

      experiment_execution:
        implementation_protocols: &#123;Standardized experiment implementation&#125;
        monitoring_procedures: &#123;Real-time experiment monitoring&#125;
        quality_control: &#123;Quality control during experiment execution&#125;
        issue_escalation: &#123;Procedures for handling experiment issues&#125;

      analysis_reporting:
        data_analysis: &#123;Systematic data analysis procedures&#125;
        insight_generation: &#123;Process for generating insights from results&#125;
        reporting_standards: &#123;Standardized experiment reporting&#125;
        decision_support: &#123;Supporting decision-making with results&#125;

      learning_integration:
        insight_capture: &#123;Systematic capture of experimental insights&#125;
        knowledge_sharing: &#123;Sharing learnings across organization&#125;
        decision_frameworks: &#123;Frameworks for acting on experimental results&#125;
        iteration_planning: &#123;Planning follow-up experiments and iterations&#125;
    ```

### Experimentation Tools and Platforms
- **Experiment Management:** {Platforms for managing experiment lifecycle}
- **A/B Testing Tools:** {Tools for digital A/B testing and optimization}
- **Analytics Platforms:** {Analytics tools for experiment measurement}
- **Collaboration Tools:** {Tools for experiment collaboration and sharing}

### Quality Assurance
    ```yaml
    quality_framework:
      design_quality:
        statistical_rigor: &#123;Statistical rigor in experimental design&#125;
        bias_mitigation: &#123;Identification and mitigation of experimental bias&#125;
        validity_threats: &#123;Assessment and mitigation of validity threats&#125;
        replication_potential: &#123;Design for experiment replication&#125;

      execution_quality:
        protocol_adherence: &#123;Adherence to experimental protocols&#125;
        data_integrity: &#123;Maintaining data quality and integrity&#125;
        contamination_prevention: &#123;Preventing experimental contamination&#125;
        timing_precision: &#123;Precise timing and duration management&#125;

      analysis_quality:
        statistical_accuracy: &#123;Accurate statistical analysis&#125;
        interpretation_validity: &#123;Valid interpretation of results&#125;
        confidence_assessment: &#123;Appropriate confidence in conclusions&#125;
        limitation_acknowledgment: &#123;Recognition of experiment limitations&#125;
    ```

## Experimentation Culture and Capabilities

### Culture Development
    ```yaml
    experimentation_culture:
      cultural_values:
        evidence_over_opinion: &#123;Valuing evidence over intuition and opinion&#125;
        learning_from_failure: &#123;Celebrating learning from failed experiments&#125;
        curiosity_encouragement: &#123;Encouraging curiosity and questioning&#125;
        iterative_improvement: &#123;Embracing iterative improvement approach&#125;

      behavioral_norms:
        hypothesis_thinking: &#123;Thinking in terms of testable hypotheses&#125;
        data_driven_decisions: &#123;Making decisions based on experimental data&#125;
        experimentation_default: &#123;Defaulting to experimentation for validation&#125;
        rapid_iteration: &#123;Embracing rapid experimentation cycles&#125;

      organizational_support:
        experimentation_time: &#123;Dedicated time for experimentation activities&#125;
        failure_tolerance: &#123;Tolerance for experimental failures and learning&#125;
        resource_access: &#123;Access to experimentation tools and platforms&#125;
        skill_development: &#123;Development of experimentation skills and capabilities&#125;
    ```

### Capability Building
- **Experimental Design Skills:** {Training in experimental design and methodology}
- **Statistical Analysis:** {Statistical analysis and interpretation skills}
- **Data Science Integration:** {Integration with data science capabilities}
- **Business Application:** {Applying experimentation to business contexts}

### Centers of Excellence
    ```yaml
    experimentation_coe:
      center_structure:
        experimentation_experts: &#123;Core team of experimentation experts&#125;
        domain_specialists: &#123;Domain specialists for different business areas&#125;
        methodology_advisors: &#123;Advisors for experimental methodology&#125;
        platform_specialists: &#123;Specialists in experimentation platforms&#125;

      center_services:
        design_consultation: &#123;Consultation on experimental design&#125;
        methodology_training: &#123;Training in experimentation methods&#125;
        platform_support: &#123;Support for experimentation platforms&#125;
        quality_review: &#123;Quality review of experimental designs&#125;

      knowledge_management:
        best_practices: &#123;Best practices for experimentation&#125;
        methodology_guides: &#123;Guides for experimental methodology&#125;
        case_studies: &#123;Case studies of successful experiments&#125;
        lesson_databases: &#123;Databases of experimental learnings&#125;
    ```

## Results Integration and Decision Making

### Insight Generation
    ```yaml
    insight_generation:
      analysis_framework:
        quantitative_analysis: &#123;Statistical analysis of experimental data&#125;
        qualitative_insights: &#123;Qualitative insights from experimental observations&#125;
        pattern_recognition: &#123;Recognition of patterns across experiments&#125;
        causal_inference: &#123;Drawing causal inferences from experimental results&#125;

      insight_synthesis:
        result_interpretation: &#123;Interpretation of experimental results&#125;
        business_implications: &#123;Business implications of experimental findings&#125;
        action_recommendations: &#123;Specific recommendations based on results&#125;
        confidence_assessment: &#123;Assessment of confidence in conclusions&#125;

      insight_communication:
        stakeholder_reporting: &#123;Reporting results to relevant stakeholders&#125;
        visualization_standards: &#123;Standards for result visualization&#125;
        narrative_construction: &#123;Constructing compelling narratives from data&#125;
        decision_support: &#123;Supporting decision-making with insights&#125;
    ```

### Decision Integration
- **Decision Frameworks:** {Frameworks for integrating experimental results into decisions}
- **Action Planning:** {Planning implementation actions based on experiment results}
- **Risk Assessment:** {Assessing risks of implementing experimental insights}
- **Success Tracking:** {Tracking success of implemented experimental insights}

### Scaling and Implementation
    ```yaml
    scaling_framework:
      validation_scaling:
        pilot_validation: &#123;Validating results through pilot implementations&#125;
        gradual_rollout: &#123;Gradual scaling of successful experiments&#125;
        monitoring_systems: &#123;Monitoring systems for scaled implementations&#125;
        adjustment_mechanisms: &#123;Mechanisms for adjusting based on scaling results&#125;

      implementation_support:
        change_management: &#123;Change management for experimental implementations&#125;
        training_support: &#123;Training for implementing experimental changes&#125;
        resource_allocation: &#123;Resource allocation for scaling implementations&#125;
        success_measurement: &#123;Measuring success of scaled implementations&#125;

      learning_loops:
        implementation_learning: &#123;Learning from implementation experiences&#125;
        feedback_integration: &#123;Integrating feedback into future experiments&#125;
        continuous_improvement: &#123;Continuous improvement of experimentation&#125;
        knowledge_updating: &#123;Updating knowledge based on implementation results&#125;
    ```

## Validation
*Evidence that experimentation enables rapid hypothesis testing, drives evidence-based decisions, and scales across organization*

## Quality Standards

### Bronze Level (Minimum Viable)
- [ ] Basic experimentation framework with simple A/B testing capabilities
- [ ] Simple hypothesis development and testing processes
- [ ] Basic experimental measurement and analysis
- [ ] All required metadata fields completed

### Silver Level (Investment Ready)
- [ ] Comprehensive experimentation framework with rigorous experimental design
- [ ] Structured experiment portfolio management with quality assurance
- [ ] Active experimentation culture with systematic insight integration
- [ ] Regular experimentation effectiveness measurement and improvement

### Gold Level (Operational Excellence)
- [ ] Advanced experimentation capabilities with sophisticated statistical methods
- [ ] Comprehensive experimentation ecosystem with centers of excellence
- [ ] Experimentation excellence with industry leadership and innovation
- [ ] Strategic experimentation driving business transformation and competitive advantage

## Common Pitfalls

1. **HiPPO decisions**: Highest paid person's opinion overriding experimental evidence
2. **Experiment aversion**: Fear of failure preventing experimentation
3. **Analysis paralysis**: Over-analyzing experiments instead of acting on results
4. **Confirmation bias**: Designing experiments to confirm preconceptions

## Success Patterns

1. **Evidence-based culture**: Culture that values experimental evidence over opinions
2. **Rapid iteration**: Fast experimentation cycles enabling rapid learning
3. **Systematic approach**: Structured and rigorous approach to experimentation
4. **Learning integration**: Effective integration of experimental learnings into business

## Relationship Guidelines

### Typical Dependencies
- **INN (Innovation)**: Innovation strategy drives experimentation priorities
- **LEA (Learning Organization)**: Learning capabilities enable experimental learning
- **DAT (Data)**: Data infrastructure enables experimental measurement
- **ANA (Analytics)**: Analytics capabilities enable experimental analysis

### Typical Enablements
- **PRD (Products)**: Experimentation enables product validation and optimization
- **SVC (Services)**: Experimentation drives service testing and improvement
- **STR (Strategy)**: Experimental insights inform strategic decisions
- **IGN (Insights)**: Experimentation generates valuable business insights

## Document Relationships

This document type commonly relates to:

- **Depends on**: INN (Innovation), LEA (Learning Organization), DAT (Data), ANA (Analytics)
- **Enables**: PRD (Products), SVC (Services), STR (Strategy), IGN (Insights)
- **Informs**: Product development, marketing optimization, business model validation
- **Guides**: Hypothesis testing, evidence-based decisions, risk mitigation

## Validation Checklist

- [ ] Executive summary with clear purpose, methodology, hypothesis framework, and validation approach
- [ ] Experimentation framework with philosophy, strategy, and scope definition
- [ ] Experimental design with hypothesis development, design methods, and measurement
- [ ] Portfolio management with framework, lifecycle, and performance measurement
- [ ] Process and operations with procedures, tools, and quality assurance
- [ ] Culture and capabilities with development, building, and centers of excellence
- [ ] Results integration with insight generation, decision integration, and scaling
- [ ] Validation evidence of rapid testing, evidence-based decisions, and organizational scaling


---

## Related Documents

- [Back to Document Types](/docs/document-types)
- [Domain: Growth & Innovation](/docs/domains/growth-innovation)
- [Full Specification](/spec/v1-0-0)
