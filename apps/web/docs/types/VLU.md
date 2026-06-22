---
title: "VLU: Valuation"
description: "BSpec VLU document type specification"
---

# VLU: Valuation

::: tip Document Type
**Code**: VLU<br>
**Name**: Valuation<br>
**Domain**: Financial & Investment
:::

## Abstract

This specification defines the Valuation document type within the BSpec 1.0 Universal Business Specification Standard. It establishes normative requirements, structured templates, and implementation guidance for organizations documenting valuation within the financial-investment domain. This specification enables systematic, machine-readable documentation that supports strategic planning, operational execution, and organizational alignment.

## Purpose and Scope

The Valuation document defines systematic approaches to determining business value, asset worth, and enterprise valuation through multiple methodologies and analytical frameworks. It establishes valuation standards that support investment decisions, transaction analysis, and strategic planning with objective, defensible value assessments.

## Scope Boundary

VLU owns valuation methodology selection and value-conclusion logic (DCF, market and asset approaches, discounts/premiums). It does **not** own integrated operating financial plans (`FIN`) or routine periodic forecasting outputs (`FOR`).

## Document Metadata Schema

```yaml
---
id: VLU-{valuation-purpose}
title: "Valuation — {Valuation Purpose or Asset}"
type: VLU
status: Draft|Review|Approved|Active|Deprecated
version: 1.0.0
owner: Finance-Team|Valuation-Analyst
stakeholders: [finance-team, executive-team, investors, board]
domain: financial
priority: Critical|High|Medium|Low
scope: valuation-analysis
horizon: strategic
visibility: confidential

depends_on: [FIN-*,FOR-*,MKT-*]
enables: [FND-*,INV-*,STR-*]

valuation_purpose: Transaction|Investment|Reporting|Tax|Litigation|Strategic
valuation_scope: Enterprise|Equity|Asset|Division|Minority-interest
valuation_standard: Fair-value|Fair-market-value|Investment-value|Liquidation-value
methodology_primary: DCF|Market|Asset-based|Hybrid

success_criteria:
  - "Valuation analysis provides accurate and defensible value estimates"
  - "Valuation methodology is appropriate for purpose and industry"
  - "Valuation process follows professional standards and best practices"
  - "Valuation conclusion supports informed decision making"

assumptions:
  - "Financial information and projections are accurate and reliable"
  - "Market data and comparable information is available and relevant"
  - "Valuation methodology is appropriate for business and context"

constraints: [Data availability and methodology constraints]
standards: [Valuation and appraisal standards]
review_cycle: quarterly
---
```

## Content Structure Template

```markdown
# Valuation — {Valuation Purpose or Asset}

## Executive Summary
**Purpose:** {Brief description of valuation purpose and scope}
**Valuation Date:** {As-of date for valuation analysis}
**Valuation Standard:** {Fair-value|Fair-market-value|Investment-value}
**Value Conclusion:** {Valuation range and point estimate}

## Valuation Framework

### Valuation Purpose and Scope
- **Valuation Purpose:** {Transaction, investment, reporting, or other purpose}
- **Valuation Subject:** {Business, assets, or interest being valued}
- **Standard of Value:** {Fair value, fair market value, or investment value}
- **Premise of Value:** {Going concern or liquidation basis}

### Valuation Standards
    ```yaml
    valuation_approach:
      professional_standards: [USPAP, IVS, ASA, AICPA]
      regulatory_compliance: [SEC, FASB, IRS, regulatory_requirements]
      quality_control: &#123;Valuation quality assurance processes&#125;

    valuation_methodology:
      approach_selection: &#123;Rationale for methodology selection&#125;
      weight_allocation: &#123;Relative weight given to each approach&#125;
      reconciliation: &#123;How multiple approaches are reconciled&#125;
    ```

### Business Overview
- **Business Description:** {Nature of business and operations}
- **Industry Analysis:** {Industry characteristics and outlook}
- **Competitive Position:** {Market position and competitive advantages}
- **Financial Performance:** {Historical and projected financial performance}

## Valuation Approaches

### Income Approach
    ```yaml
    income_approach:
      discounted_cash_flow:
        projection_period: &#123;Forecast period for cash flow projections&#125;
        terminal_value: &#123;Terminal value calculation method&#125;
        discount_rate: &#123;Weighted average cost of capital&#125;

        cash_flow_projections:
          revenue_growth: &#123;Revenue growth assumptions&#125;
          margin_assumptions: &#123;Profit margin projections&#125;
          capex_requirements: &#123;Capital expenditure needs&#125;
          working_capital: &#123;Working capital changes&#125;

        discount_rate_analysis:
          risk_free_rate: &#123;Government bond yield&#125;
          equity_risk_premium: &#123;Market risk premium&#125;
          beta: &#123;Company-specific risk factor&#125;
          debt_cost: &#123;After-tax cost of debt&#125;
          capital_structure: &#123;Target debt-to-equity ratio&#125;

      capitalization_method:
        normalized_earnings: &#123;Sustainable earnings level&#125;
        capitalization_rate: &#123;Required rate of return&#125;
        growth_assumptions: &#123;Long-term growth expectations&#125;
    ```

### Market Approach
- **Comparable Company Analysis:** {Public company trading multiples}
- **Precedent Transaction Analysis:** {Recent transaction multiples}
- **Market Data Sources:** {Industry databases and market information}
- **Multiple Selection:** {Appropriate valuation multiples}

### Asset Approach
    ```yaml
    asset_approach:
      net_asset_value:
        book_value_adjustments: &#123;Fair value adjustments to book value&#125;
        asset_revaluations: [real_estate, equipment, intangibles]
        liability_adjustments: &#123;Fair value of liabilities&#125;

      liquidation_value:
        orderly_liquidation: &#123;Systematic asset liquidation value&#125;
        forced_liquidation: &#123;Distressed sale value&#125;
        liquidation_costs: &#123;Transaction and disposal costs&#125;

      replacement_cost:
        asset_reproduction: &#123;Cost to reproduce identical assets&#125;
        asset_replacement: &#123;Cost to replace with equivalent assets&#125;
        functional_obsolescence: &#123;Loss due to functional obsolescence&#125;
        economic_obsolescence: &#123;Loss due to economic factors&#125;
    ```

## Market Analysis

### Industry Analysis
    ```yaml
    industry_assessment:
      industry_overview:
        size_and_growth: &#123;Market size and growth characteristics&#125;
        life_cycle_stage: &#123;Industry maturity and development stage&#125;
        competitive_dynamics: &#123;Competition level and structure&#125;

      industry_outlook:
        growth_drivers: [driver1, driver2, driver3]
        risk_factors: [risk1, risk2, risk3]
        regulatory_environment: &#123;Regulatory impact on industry&#125;

      key_success_factors:
        competitive_advantages: [advantage1, advantage2]
        critical_resources: [resource1, resource2]
        industry_barriers: [barrier1, barrier2]
    ```

### Comparable Analysis
- **Public Company Comparables:** {Listed organizations with similar business models}
- **Transaction Comparables:** {Recent M&A transactions in industry}
- **Size and Scale Adjustments:** {Adjustments for size differences}
- **Control Premium Analysis:** {Control vs. minority interest premiums}

### Market Multiples
    ```yaml
    valuation_multiples:
      revenue_multiples:
        - multiple: &#123;Price/Revenue&#125;
          median: &#123;Market median multiple&#125;
          range: &#123;Multiple range for comparables&#125;
          application: &#123;Application to subject company&#125;

      earnings_multiples:
        - multiple: &#123;Price/EBITDA&#125;
          median: &#123;Market median multiple&#125;
          adjustment: &#123;Adjustments for differences&#125;
          implied_value: &#123;Implied value calculation&#125;

      book_value_multiples:
        - multiple: &#123;Price/Book Value&#125;
          industry_average: &#123;Industry average multiple&#125;
          premium_discount: &#123;Premium or discount rationale&#125;
    ```

## Discount and Premium Analysis

### Discount Factors
    ```yaml
    valuation_discounts:
      marketability_discount:
        lack_of_marketability: &#123;Discount for illiquid investment&#125;
        restricted_stock_studies: &#123;Market evidence of discounts&#125;
        pre_ipo_studies: &#123;Private transaction evidence&#125;

      minority_discount:
        lack_of_control: &#123;Discount for minority position&#125;
        minority_oppression: &#123;Additional minority discounts&#125;
        voting_vs_nonvoting: &#123;Difference in voting rights&#125;

      key_person_discount:
        management_dependency: &#123;Discount for key person risk&#125;
        succession_planning: &#123;Management transition risks&#125;
        operational_continuity: &#123;Business continuity risks&#125;
    ```

### Premium Factors
- **Control Premium:** {Premium for controlling interest}
- **Strategic Premium:** {Premium for strategic value}
- **Synergy Premium:** {Premium for acquisition synergies}
- **Scarcity Premium:** {Premium for unique assets or market position}

### Size Premium Analysis
    ```yaml
    size_analysis:
      size_premium:
        company_size: &#123;Revenue, assets, market cap size metrics&#125;
        size_premium_data: &#123;Historical size premium evidence&#125;
        specific_company_risk: &#123;Company-specific risk factors&#125;

      small_company_factors:
        limited_access_capital: &#123;Capital access constraints&#125;
        customer_concentration: &#123;Customer diversification risks&#125;
        management_depth: &#123;Management team capabilities&#125;
    ```

## Sensitivity Analysis

### Key Variable Analysis
    ```yaml
    sensitivity_testing:
      revenue_sensitivity:
        base_case: &#123;Base revenue growth assumption&#125;
        upside_case: &#123;Optimistic revenue scenario&#125;
        downside_case: &#123;Conservative revenue scenario&#125;
        value_impact: &#123;Impact on valuation conclusion&#125;

      margin_sensitivity:
        margin_assumptions: [gross_margin, ebitda_margin, net_margin]
        improvement_scenarios: &#123;Operational improvement cases&#125;
        compression_scenarios: &#123;Competitive pressure cases&#125;

      discount_rate_sensitivity:
        base_discount_rate: &#123;Base case discount rate&#125;
        rate_range: &#123;Range of discount rates tested&#125;
        value_sensitivity: &#123;Valuation sensitivity to rate changes&#125;
    ```

### Scenario Analysis
- **Base Case Scenario:** {Most likely business performance scenario}
- **Upside Scenario:** {Optimistic business performance scenario}
- **Downside Scenario:** {Conservative business performance scenario}
- **Stress Test Scenarios:** {Extreme adverse condition scenarios}

### Monte Carlo Analysis
    ```yaml
    monte_carlo:
      variable_distributions: &#123;Probability distributions for key variables&#125;
      correlation_assumptions: &#123;Variable correlation relationships&#125;
      simulation_results: &#123;Distribution of valuation outcomes&#125;
      risk_metrics: &#123;Value at risk and confidence intervals&#125;
    ```

## Valuation Reconciliation

### Approach Weighting
    ```yaml
    valuation_reconciliation:
      approach_weights:
        income_approach: &#123;Weight assigned to income approach&#125;
        market_approach: &#123;Weight assigned to market approach&#125;
        asset_approach: &#123;Weight assigned to asset approach&#125;

      weighting_rationale:
        data_reliability: &#123;Quality and reliability of approach data&#125;
        methodology_appropriateness: &#123;Suitability for valuation purpose&#125;
        market_acceptance: &#123;Market acceptance of approach&#125;

      reconciliation_process:
        approach_comparison: &#123;Comparison of approach results&#125;
        outlier_analysis: &#123;Analysis of significant differences&#125;
        final_conclusion: &#123;Final value conclusion and range&#125;
    ```

### Value Conclusion
- **Valuation Range:** {Range of reasonable value estimates}
- **Point Estimate:** {Single point value conclusion}
- **Confidence Level:** {Confidence in valuation conclusion}
- **Key Value Drivers:** {Primary factors driving value conclusion}

## Special Considerations

### Complex Securities
    ```yaml
    complex_valuations:
      preferred_equity:
        dividend_rights: &#123;Dividend preferences and rates&#125;
        liquidation_preferences: &#123;Liquidation priority and multiples&#125;
        conversion_features: &#123;Conversion terms and triggers&#125;

      warrants_options:
        exercise_terms: &#123;Strike price and exercise period&#125;
        valuation_model: &#123;Black-Scholes or binomial model&#125;
        volatility_assumptions: &#123;Share price volatility estimates&#125;

      convertible_debt:
        debt_component: &#123;Straight debt value&#125;
        equity_component: &#123;Conversion option value&#125;
        credit_risk: &#123;Default risk and credit spreads&#125;
    ```

### Special Situations
- **Distressed Company Valuation:** {Valuation under financial distress}
- **Start-up Valuation:** {Valuation of early-stage organizations}
- **Intangible Asset Valuation:** {Valuation of intellectual property and intangibles}
- **International Valuation:** {Cross-border valuation considerations}

## Quality Control and Review

### Valuation Review Process
    ```yaml
    quality_control:
      technical_review:
        methodology_review: &#123;Review of valuation approaches&#125;
        calculation_review: &#123;Review of calculations and models&#125;
        assumption_review: &#123;Review of key assumptions&#125;

      independent_review:
        peer_review: &#123;Independent expert review&#125;
        technical_standards: &#123;Compliance with professional standards&#125;
        documentation_review: &#123;Review of supporting documentation&#125;

      client_review:
        draft_presentation: &#123;Initial findings presentation&#125;
        feedback_incorporation: &#123;Client feedback and revisions&#125;
        final_approval: &#123;Final valuation approval process&#125;
    ```

### Documentation Standards
- **Valuation Report:** {Comprehensive valuation report documentation}
- **Supporting Analysis:** {Detailed supporting calculations and analysis}
- **Data Sources:** {Documentation of data sources and assumptions}
- **Professional Standards:** {Compliance with valuation standards}

## Validation
*Evidence that valuation analysis provides accurate value estimates, follows professional standards, and supports informed decision making*

## Quality Standards

### Bronze Level (Minimum Viable)
- [ ] Basic valuation using single approach
- [ ] Simple comparable analysis with market multiples
- [ ] Basic sensitivity analysis and documentation
- [ ] All required metadata fields completed

### Silver Level (Investment Ready)
- [ ] Multiple valuation approaches with appropriate weighting
- [ ] Comprehensive market analysis and comparable selection
- [ ] Detailed sensitivity analysis and scenario modeling
- [ ] Professional quality documentation and review

### Gold Level (Operational Excellence)
- [ ] Sophisticated valuation modeling with advanced analytics
- [ ] Comprehensive risk analysis and Monte Carlo simulation
- [ ] Complex security valuation and special situation analysis
- [ ] Expert-level documentation meeting professional standards

## Common Pitfalls

1. **Inappropriate methodology**: Using valuation approaches not suitable for purpose or industry
2. **Poor comparable selection**: Using non-comparable organizations or transactions
3. **Unrealistic assumptions**: Overly optimistic or pessimistic assumptions
4. **Inadequate support**: Insufficient documentation and analysis supporting conclusions

## Success Patterns

1. **Multiple methodology validation**: Using multiple approaches to validate value conclusions with appropriate weighting
2. **Thorough market analysis**: Comprehensive analysis of market conditions and comparable transactions
3. **Robust sensitivity testing**: Extensive sensitivity and scenario analysis to test assumption impact
4. **Professional standards compliance**: Adherence to professional valuation standards and quality control

## Relationship Guidelines

### Typical Dependencies
- **FIN (Financial Model)**: Financial projections drive DCF and earnings-based valuations
- **FOR (Forecasts)**: Market forecasts drive growth assumptions and terminal value calculations
- **MKT (Market Definition)**: Market analysis drives comparable selection and market approach
- **Industry Analysis**: Industry analysis drives risk assessment and multiple selection

### Typical Enablements
- **FND (Funding)**: Valuation analysis drives funding pricing and negotiation
- **INV (Investment)**: Valuation targets drive investment decision criteria
- **STR (Strategy)**: Valuation insights drive strategic planning and resource allocation
- **Negotiation**: Valuation analysis drives deal negotiation and structuring

## Document Relationships

This document type commonly relates to:

- **Depends on**: FIN (Financial Model), FOR (Forecasts), MKT (Market Definition), IND (Industry Analysis)
- **Enables**: FND (Funding), INV (Investment), STR (Strategy), NEG (Negotiation)
- **Informs**: Investment decisions, transaction pricing, strategic planning
- **Guides**: Deal structure, funding terms, acquisition strategy

## Validation Checklist

- [ ] Executive summary with clear purpose, valuation date, standard, and value conclusion
- [ ] Valuation framework with purpose, scope, standards, and business overview
- [ ] Valuation approaches with income, market, and asset-based methodologies
- [ ] Market analysis with industry assessment, comparable analysis, and market multiples
- [ ] Discount and premium analysis with discount factors, premiums, and size analysis
- [ ] Sensitivity analysis with key variable testing, scenarios, and Monte Carlo simulation
- [ ] Valuation reconciliation with approach weighting and value conclusion
- [ ] Special considerations for complex securities and special situations
- [ ] Quality control with review process and documentation standards
- [ ] Validation evidence of valuation accuracy, professional standards compliance, and decision support


---

## Related Documents

- [Back to Document Types](/docs/document-types)
- [Domain: Financial & Investment](/docs/domains/financial-investment)
- [Full Specification](/spec/v1-0-0)
