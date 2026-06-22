---
title: "MET: Metrics"
description: "BSpec MET document type specification"
---

# MET: Metrics

::: tip Document Type
**Code**: MET<br>
**Name**: Metrics<br>
**Domain**: Financial & Investment
:::

## Abstract

This specification defines the Metrics document type within the BSpec 1.0 Universal Business Specification Standard. It establishes normative requirements, structured templates, and implementation guidance for organizations documenting metrics within the financial-investment domain. This specification enables systematic, machine-readable documentation that supports strategic planning, operational execution, and organizational alignment.

## Framework and Attribution

Balanced Scorecard structure elements in this document are attributed to Kaplan & Norton and must be used according to their publication terms and any applicable trademark policy.

## Purpose and Scope

The Metrics document defines systematic approaches to measuring, monitoring, and managing business performance through key performance indicators, financial metrics, and operational measures. It establishes measurement frameworks that enable data-driven decision making, performance accountability, and continuous improvement.

## Scope Boundary

MET owns metric definitions, reporting architectures, and measurement cadences. It converts `OBJ` commitments into measurable signals and dashboards, but does **not** own the underlying forecasts (`FOR`) or the full financial projection model (`FIN`).

## Document Metadata Schema

```yaml
---
id: MET-{metric-category}
title: "Metrics — {Metric Category or Dashboard}"
type: MET
status: Draft|Review|Approved|Active|Deprecated
attribution_required: true
source_frameworks:
  - "Kaplan & Norton - Balanced Scorecard"
version: 1.0.0
owner: Finance-Team|Analytics-Team
stakeholders: [finance-team, business-units, executives, managers]
domain: financial
priority: Critical|High|Medium|Low
scope: performance-measurement
horizon: tactical
visibility: internal

depends_on: [FIN-*,BUD-*,STR-*,OBJ-*]
enables: [REP-*,PER-*,QUA-*,GOV-*]

metric_type: Financial|Operational|Customer|Employee|Quality
measurement_level: Strategic|Tactical|Operational
reporting_frequency: Real-time|Daily|Weekly|Monthly|Quarterly
automation_level: Manual|Semi-automated|Fully-automated

success_criteria:
  - "Metrics provide actionable insights for business performance"
  - "Measurement system drives accountability and improvement"
  - "Metrics are accurate, timely, and accessible to stakeholders"
  - "Performance measurement supports strategic objectives"

assumptions:
  - "Data sources are reliable and measurement systems are accurate"
  - "Stakeholders understand and act on metric insights"
  - "Measurement processes are efficient and cost-effective"

constraints: [Data availability and system constraints]
standards: [Performance measurement and reporting standards]
review_cycle: monthly
---
```

## Content Structure Template

```markdown
# Metrics — {Metric Category or Dashboard}

## Executive Summary
**Purpose:** {Brief description of metrics scope and objectives}
**Measurement Focus:** {Financial|Operational|Customer|Employee|Quality}
**Reporting Level:** {Strategic|Tactical|Operational}
**Update Frequency:** {Real-time|Daily|Weekly|Monthly|Quarterly}

## Metrics Framework

### Performance Philosophy
- **Measurement Purpose:** {How metrics drive business performance}
- **Data-Driven Culture:** {Culture of measurement and accountability}
- **Continuous Improvement:** {Metrics-driven improvement process}
- **Stakeholder Value:** {How metrics create stakeholder value}

### Metrics Strategy
    ```yaml
    measurement_approach:
      balanced_scorecard: &#123;Financial, customer, process, learning perspectives&#125;
      okr_framework: &#123;Objectives and key results methodology&#125;
      kpi_hierarchy: &#123;Strategic, tactical, and operational KPI levels&#125;

    metrics_governance:
      metric_ownership: &#123;Clear accountability for metric performance&#125;
      data_quality: &#123;Data accuracy and reliability standards&#125;
      reporting_standards: &#123;Consistent measurement and reporting&#125;
      review_process: &#123;Regular metric review and optimization&#125;
    ```

### Measurement Categories
- **Financial Metrics:** {Revenue, profitability, cash flow, and efficiency metrics}
- **Operational Metrics:** {Process efficiency, quality, and productivity metrics}
- **Customer Metrics:** {Satisfaction, retention, and value metrics}
- **Employee Metrics:** {Engagement, productivity, and development metrics}

## Financial Metrics

### Revenue Metrics
    ```yaml
    revenue_metrics:
      growth_metrics:
        - metric: &#123;Revenue Growth Rate&#125;
          calculation: &#123;(Current Period - Prior Period) / Prior Period&#125;
          target: &#123;Target growth percentage&#125;
          frequency: &#123;Monthly|Quarterly|Annual&#125;

        - metric: &#123;Monthly Recurring Revenue (MRR)&#125;
          calculation: &#123;Sum of monthly subscription revenue&#125;
          segments: [new_business, expansion, churn]
          trend: &#123;MRR growth trend analysis&#125;

      quality_metrics:
        - metric: &#123;Revenue per Customer&#125;
          calculation: &#123;Total Revenue / Number of Customers&#125;
          benchmark: &#123;Industry or historical benchmark&#125;

        - metric: &#123;Customer Lifetime Value&#125;
          calculation: &#123;Average Revenue per Customer / Churn Rate&#125;
          cohort_analysis: &#123;CLV by customer cohort&#125;
    ```

### Profitability Metrics
- **Gross Margin:** {Revenue minus cost of goods sold as percentage}
- **EBITDA Margin:** {EBITDA as percentage of revenue}
- **Net Profit Margin:** {Net income as percentage of revenue}
- **Return on Assets:** {Net income divided by average total assets}
- **Return on Equity:** {Net income divided by average shareholders' equity}

### Cash Flow Metrics
    ```yaml
    cash_flow_metrics:
      operational_cash_flow:
        operating_cash_ratio: &#123;Operating cash flow / Current liabilities&#125;
        cash_conversion_cycle: &#123;Days sales outstanding + Days inventory - Days payable&#125;
        free_cash_flow: &#123;Operating cash flow - Capital expenditures&#125;

      liquidity_metrics:
        current_ratio: &#123;Current assets / Current liabilities&#125;
        quick_ratio: &#123;(Current assets - Inventory) / Current liabilities&#125;
        cash_ratio: &#123;Cash and equivalents / Current liabilities&#125;

      efficiency_metrics:
        asset_turnover: &#123;Revenue / Average total assets&#125;
        receivables_turnover: &#123;Revenue / Average accounts receivable&#125;
        inventory_turnover: &#123;Cost of goods sold / Average inventory&#125;
    ```

## Operational Metrics

### Productivity Metrics
    ```yaml
    productivity_metrics:
      efficiency_measures:
        - metric: &#123;Revenue per Employee&#125;
          calculation: &#123;Total revenue / Number of employees&#125;
          benchmark: &#123;Industry benchmark&#125;
          trend: &#123;Productivity trend analysis&#125;

        - metric: &#123;Output per Hour&#125;
          calculation: &#123;Units produced / Total hours worked&#125;
          quality_adjustment: &#123;Quality-adjusted productivity&#125;

      process_metrics:
        - metric: &#123;Cycle Time&#125;
          measurement: &#123;Time from start to completion&#125;
          target: &#123;Target cycle time&#125;
          improvement: &#123;Cycle time reduction initiatives&#125;

        - metric: &#123;First Pass Yield&#125;
          calculation: &#123;Units passing first time / Total units&#125;
          quality_cost: &#123;Cost of quality failures&#125;
    ```

### Quality Metrics
- **Defect Rate:** {Number of defects per unit produced}
- **Customer Complaints:** {Number and resolution time of complaints}
- **Rework Rate:** {Percentage of work requiring rework}
- **Error Rate:** {Frequency of errors in processes}

### Capacity Metrics
    ```yaml
    capacity_metrics:
      utilization:
        - metric: &#123;Capacity Utilization&#125;
          calculation: &#123;Actual output / Maximum possible output&#125;
          target: &#123;Target utilization percentage&#125;

      throughput:
        - metric: &#123;Units per Period&#125;
          measurement: &#123;Production or service delivery rate&#125;
          bottleneck_analysis: &#123;Capacity constraint identification&#125;

      scalability:
        - metric: &#123;Scalability Index&#125;
          measurement: &#123;Output increase / Resource increase&#125;
          efficiency_curve: &#123;Economies of scale analysis&#125;
    ```

## Customer Metrics

### Acquisition Metrics
    ```yaml
    customer_metrics:
      acquisition:
        - metric: &#123;Customer Acquisition Cost (CAC)&#125;
          calculation: &#123;Sales and marketing costs / New customers acquired&#125;
          channel_breakdown: &#123;CAC by acquisition channel&#125;
          payback_period: &#123;Time to recover acquisition cost&#125;

        - metric: &#123;Conversion Rate&#125;
          calculation: &#123;Customers acquired / Prospects contacted&#125;
          funnel_analysis: &#123;Conversion at each stage&#125;

      retention:
        - metric: &#123;Customer Retention Rate&#125;
          calculation: &#123;(Customers at end - New customers) / Customers at start&#125;
          cohort_analysis: &#123;Retention by customer cohort&#125;

        - metric: &#123;Churn Rate&#125;
          calculation: &#123;Customers lost / Total customers at start&#125;
          voluntary_involuntary: &#123;Voluntary vs involuntary churn&#125;
    ```

### Satisfaction Metrics
- **Net Promoter Score (NPS):** {Likelihood to recommend measurement}
- **Customer Satisfaction Score:** {Overall satisfaction rating}
- **Customer Effort Score:** {Ease of doing business measurement}
- **Support Ticket Resolution:** {Time and quality of issue resolution}

### Value Metrics
    ```yaml
    value_metrics:
      economic_value:
        - metric: &#123;Customer Lifetime Value (CLV)&#125;
          calculation: &#123;Average revenue per customer * Gross margin * Lifetime&#125;
          clv_cac_ratio: &#123;CLV to CAC ratio&#125;

        - metric: &#123;Revenue per Customer&#125;
          measurement: &#123;Total revenue / Number of customers&#125;
          growth_rate: &#123;Revenue per customer growth&#125;

      engagement:
        - metric: &#123;Product Usage&#125;
          measurement: &#123;Feature usage and adoption rates&#125;
          stickiness: &#123;User engagement and retention&#125;
    ```

## Employee Metrics

### Engagement Metrics
    ```yaml
    employee_metrics:
      engagement:
        - metric: &#123;Employee Engagement Score&#125;
          measurement: &#123;Survey-based engagement measurement&#125;
          drivers: [recognition, development, autonomy, purpose]

        - metric: &#123;Employee Net Promoter Score&#125;
          calculation: &#123;Likelihood to recommend as workplace&#125;
          benchmark: &#123;Industry engagement benchmarks&#125;

      retention:
        - metric: &#123;Employee Turnover Rate&#125;
          calculation: &#123;Employees leaving / Average employees&#125;
          voluntary_involuntary: &#123;Voluntary vs involuntary turnover&#125;
          cost_of_turnover: &#123;Financial impact of turnover&#125;

      development:
        - metric: &#123;Training Hours per Employee&#125;
          measurement: &#123;Total training hours / Number of employees&#125;
          roi_training: &#123;Return on training investment&#125;
    ```

### Performance Metrics
- **Goal Achievement Rate:** {Percentage of employees meeting objectives}
- **Performance Rating Distribution:** {Performance review score distribution}
- **Promotion Rate:** {Internal promotion and advancement rates}
- **Skills Development:** {Skill acquisition and certification progress}

## Metrics Dashboard and Reporting

### Dashboard Design
    ```yaml
    dashboard_framework:
      executive_dashboard:
        kpis: [revenue_growth, profitability, cash_flow, customer_satisfaction]
        frequency: &#123;Daily updates with weekly executive review&#125;
        format: &#123;High-level summary with exception reporting&#125;

      operational_dashboard:
        metrics: [productivity, quality, capacity, employee_engagement]
        frequency: &#123;Real-time or daily updates&#125;
        interactivity: &#123;Drill-down and filtering capabilities&#125;

      financial_dashboard:
        reports: [p&l, balance_sheet, cash_flow, budget_variance]
        frequency: &#123;Monthly close with quarterly forecasts&#125;
        analytics: &#123;Trend analysis and variance reporting&#125;
    ```

### Reporting Structure
- **Strategic Reports:** {Board and executive-level performance summaries}
- **Management Reports:** {Department and functional area performance}
- **Operational Reports:** {Day-to-day operational performance monitoring}
- **Exception Reports:** {Alerts and notifications for performance issues}

### Data Visualization
    ```yaml
    visualization:
      chart_types:
        trend_analysis: [line_charts, area_charts]
        performance_comparison: [bar_charts, column_charts]
        composition_analysis: [pie_charts, stacked_charts]
        correlation_analysis: [scatter_plots, bubble_charts]

      design_principles:
        clarity: &#123;Clear and unambiguous visual presentation&#125;
        simplicity: &#123;Minimal and focused information display&#125;
        consistency: &#123;Consistent design and formatting&#125;
        interactivity: &#123;User-friendly navigation and exploration&#125;
    ```

## Performance Management

### Target Setting
    ```yaml
    target_management:
      target_methodology:
        historical_performance: &#123;Historical trend-based targets&#125;
        benchmarking: &#123;Industry and peer comparison targets&#125;
        strategic_goals: &#123;Strategy-driven target setting&#125;

      target_types:
        threshold: &#123;Minimum acceptable performance level&#125;
        target: &#123;Expected performance level&#125;
        stretch: &#123;Aspirational performance level&#125;

      calibration:
        difficulty_assessment: &#123;Target achievability analysis&#125;
        resource_alignment: &#123;Resource allocation to support targets&#125;
        risk_adjustment: &#123;Risk-adjusted target setting&#125;
    ```

### Performance Review
- **Variance Analysis:** {Actual vs. target performance analysis}
- **Root Cause Analysis:** {Understanding performance drivers and barriers}
- **Corrective Actions:** {Performance improvement initiatives}
- **Trend Analysis:** {Long-term performance pattern analysis}

### Continuous Improvement
    ```yaml
    improvement_process:
      metric_optimization:
        metric_relevance: &#123;Regular review of metric relevance&#125;
        measurement_accuracy: &#123;Data quality and accuracy improvement&#125;
        reporting_efficiency: &#123;Streamlined reporting processes&#125;

      performance_enhancement:
        best_practice_sharing: &#123;Cross-functional learning&#125;
        process_improvement: &#123;Performance-driven process optimization&#125;
        innovation_metrics: &#123;Metrics to drive innovation&#125;
    ```

## Data Quality and Governance

### Data Management
    ```yaml
    data_governance:
      data_sources:
        system_integration: &#123;ERP, CRM, and operational system integration&#125;
        data_validation: &#123;Automated data quality checks&#125;
        master_data: &#123;Single source of truth for key entities&#125;

      data_quality:
        accuracy: &#123;Data accuracy measurement and improvement&#125;
        completeness: &#123;Data completeness monitoring&#125;
        timeliness: &#123;Data freshness and update frequency&#125;
        consistency: &#123;Data consistency across systems&#125;

      data_security:
        access_control: &#123;Role-based access to sensitive metrics&#125;
        audit_trail: &#123;Data access and modification tracking&#125;
        compliance: &#123;Regulatory compliance for data handling&#125;
    ```

### Automation and Technology
- **Automated Data Collection:** {Automated metric calculation and reporting}
- **Real-Time Monitoring:** {Real-time performance tracking and alerting}
- **Predictive Analytics:** {Predictive performance modeling}
- **Self-Service Analytics:** {User-friendly metric exploration tools}

## Validation
*Evidence that metrics provide actionable insights, drive accountability, and support strategic objectives*

## Quality Standards

### Bronze Level (Minimum Viable)
- [ ] Basic financial and operational metrics with manual reporting
- [ ] Simple dashboard with key performance indicators
- [ ] Monthly performance review and variance analysis
- [ ] All required metadata fields completed

### Silver Level (Investment Ready)
- [ ] Comprehensive metric framework with balanced scorecard approach
- [ ] Automated data collection and reporting with quality controls
- [ ] Interactive dashboards with drill-down capabilities
- [ ] Regular performance management and improvement processes

### Gold Level (Operational Excellence)
- [ ] Advanced analytics with predictive and prescriptive metrics
- [ ] Real-time performance monitoring with automated alerting
- [ ] Sophisticated data governance with self-service capabilities
- [ ] Strategic performance optimization with continuous improvement culture

## Common Pitfalls

1. **Metric overload**: Too many metrics without focus on key drivers
2. **Poor data quality**: Inaccurate or untimely data undermining decision making
3. **Lack of action**: Measuring without taking action on insights
4. **Gaming behaviors**: Metrics driving counterproductive behaviors

## Success Patterns

1. **Balanced measurement**: Comprehensive metrics covering all key performance dimensions
2. **Actionable insights**: Metrics that directly drive decision making and improvement actions
3. **Cultural integration**: Measurement embedded in daily operations and decision processes
4. **Continuous evolution**: Regular refinement and optimization of measurement approach

## Relationship Guidelines

### Typical Dependencies
- **FIN (Financial Model)**: Financial models drive financial metric definitions and targets
- **BUD (Budget)**: Budget targets drive performance measurement and variance analysis
- **STR (Strategy)**: Strategic objectives drive KPI selection and target setting
- **OBJ (Objectives)**: Business objectives drive metric prioritization and focus

### Typical Enablements
- **REP (Reporting)**: Metrics drive reporting requirements and dashboard design
- **PER (Performance)**: Performance measurement drives improvement initiatives
- **QUA (Quality)**: Quality metrics drive quality management and improvement
- **GOV (Governance)**: Performance metrics drive governance and accountability processes

## Document Relationships

This document type commonly relates to:

- **Depends on**: FIN (Financial Model), BUD (Budget), STR (Strategy), OBJ (Objectives)
- **Enables**: REP (Reporting), PER (Performance), QUA (Quality), GOV (Governance)
- **Informs**: Performance management, decision making, strategic planning
- **Guides**: Data collection, reporting design, improvement initiatives

## Validation Checklist

- [ ] Executive summary with clear purpose, measurement focus, reporting level, and frequency
- [ ] Metrics framework with performance philosophy, strategy, and measurement categories
- [ ] Financial metrics with revenue, profitability, and cash flow measurements
- [ ] Operational metrics with productivity, quality, and capacity measurements
- [ ] Customer metrics with acquisition, satisfaction, and value measurements
- [ ] Employee metrics with engagement, retention, and performance measurements
- [ ] Dashboard and reporting with design framework, structure, and visualization
- [ ] Performance management with target setting, review, and continuous improvement
- [ ] Data quality and governance with management, automation, and technology
- [ ] Validation evidence of actionable insights, accountability, and strategic objective support


---

## Related Documents

- [Back to Document Types](/docs/document-types)
- [Domain: Financial & Investment](/docs/domains/financial-investment)
- [Full Specification](/spec/v1-0-0)
