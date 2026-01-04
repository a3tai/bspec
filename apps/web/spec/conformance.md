---
title: "Conformance Levels"
description: "BSpec conformance levels from Bronze to Gold"
---

# Conformance Levels

BSpec defines three conformance levels that help organizations progressively build comprehensive business specifications.

## Overview

```
Bronze (Starter)     → 12+ documents  → Core foundation
Silver (Professional) → 25+ documents  → Full coverage
Gold (Comprehensive)  → 45+ documents  → Deep detail
```

## Bronze Level (Starter)

**Minimum**: 12 core documents

Bronze level establishes your fundamental business foundation.

### Required Documents

1. **VSN** - Vision: Where are you going?
2. **MSN** - Mission: Why do you exist?
3. **STR** - Strategy: How will you get there?
4. **VAL** - Values: What principles guide you?
5. **OBJ** - Objectives: What are your goals?
6. **CUS** - Customer: Who do you serve?
7. **PRD** - Product: What do you offer?
8. **REV** - Revenue Model: How do you make money?
9. **OPS** - Operations: How do you deliver?
10. **ORG** - Organization: How are you structured?
11. **FIN** - Financial Plan: What are your numbers?
12. **RSK** - Risk Assessment: What could go wrong?

### Benefits

- Clear strategic direction
- Basic business model understanding
- Foundation for growth
- Investor/stakeholder clarity

### Typical For

- Startups
- Small businesses
- New initiatives
- MVP documentation

## Silver Level (Professional)

**Minimum**: 25 documents across all domains

Silver level provides comprehensive coverage of your business.

### Required Coverage

Must include documents from **all 12 domains**:

1. **Strategic Foundation** (3+)
   - VSN, MSN, STR minimum

2. **Market Environment** (2+)
   - Example: MKT, CMP

3. **Customer Value** (2+)
   - Example: PER, USE

4. **Product & Service** (2+)
   - Example: PRD, FEA

5. **Business Model** (2+)
   - Example: REV, PRI

6. **Operations & Execution** (3+)
   - Example: OPS, PRO, TEA

7. **Technology & Data** (2+)
   - Example: ARC, API

8. **Financial & Investment** (2+)
   - Example: FIN, BUD

9. **Risk & Governance** (2+)
   - Example: RSK, GOV

10. **Growth & Innovation** (2+)
    - Example: INN, EXP

11. **Learning & Decisions** (1+)
    - Example: DEC

12. **Brand & Marketing** (2+)
    - Example: BRD, MSG

### Benefits

- Complete business view
- All critical areas covered
- Ready for scale
- Professional documentation

### Typical For

- Growing companies
- Series A+ startups
- SMB with structure
- Organizations seeking investment

## Gold Level (Comprehensive)

**Minimum**: 45+ documents with deep coverage

Gold level provides detailed documentation across all business areas.

### Required Coverage

**Expanded Coverage Across All Domains**:

- **Strategic Foundation**: 5+ documents
- **Market Environment**: 5+ documents
- **Customer Value**: 5+ documents
- **Product & Service**: 5+ documents
- **Business Model**: 4+ documents
- **Operations & Execution**: 5+ documents
- **Technology & Data**: 3+ documents
- **Financial & Investment**: 4+ documents
- **Risk & Governance**: 3+ documents
- **Growth & Innovation**: 3+ documents
- **Learning & Decisions**: 2+ documents
- **Brand & Marketing**: 4+ documents

### Benefits

- Deep operational detail
- Complete knowledge graph
- Full audit trail
- Enterprise-grade documentation
- AI-ready knowledge base

### Typical For

- Enterprise organizations
- Public companies
- Regulated industries
- Complex operations
- Organizations with compliance requirements

## Industry Profiles

BSpec provides industry-specific profiles with recommended documents:

### Software/SaaS Profile

**Priority Documents**:
- VSN, MSN, STR, VAL, OBJ (Strategic)
- PRD, FEA, API, UXD (Product)
- REV, PRI, CST (Business Model)
- ARC, SYS, SEC (Technology)
- BRD, MSG, SEO (Marketing)

### Physical Products Profile

**Priority Documents**:
- VSN, MSN, STR, PUR (Strategic)
- PRD, QUA, SUP (Product)
- CHN, VND, FAC (Operations)
- INV, FOR, CST (Financial)
- MKT, SEG, CMP (Market)

### Service Business Profile

**Priority Documents**:
- VSN, MSN, VAL (Strategic)
- SVC, SLA, SUP (Service)
- CUS, PER, CJM (Customer)
- PRO, WFL, TEA (Operations)
- BRD, MSG, CAM (Marketing)

### Nonprofit Profile

**Priority Documents**:
- MSN, VAL, PUR, THY (Purpose)
- STO, PAI, CIN (Impact)
- FND, BUD, REP (Financial)
- GOV, ETH, COM (Governance)
- CAM, SOC, CNT (Communications)

## Conformance Calculator

Check your conformance level:

```bash
# Using CLI
bspec analyze docs/

# Using TypeScript SDK
import { analyzeConformance } from '@bspec/sdk'
const analysis = analyzeConformance(documents)
console.log(`Level: ${analysis.level}`)
console.log(`Coverage: ${analysis.coverage}%`)
```

## Progression Strategy

### From Zero to Bronze (Week 1-2)

1. Start with Strategic Foundation (VSN, MSN, STR, VAL)
2. Add Customer and Product (CUS, PRD)
3. Define Business Model (REV)
4. Document Operations (OPS, ORG)
5. Add Financial and Risk (FIN, RSK)

### From Bronze to Silver (Month 1-3)

1. Expand each domain to 2+ documents
2. Add missing domains (all 12 must be covered)
3. Define relationships between documents
4. Validate and refine existing documents
5. Add stakeholder reviews

### From Silver to Gold (Month 3-6)

1. Deep dive into each domain
2. Add detailed operational documents
3. Create comprehensive relationship graphs
4. Implement regular review cycles
5. Build living documentation culture

## Best Practices

1. **Start Small**: Begin with Bronze, expand progressively
2. **Quality Over Quantity**: 12 great documents beat 50 mediocre ones
3. **Relationships Matter**: Connect documents with depends_on/enables
4. **Keep Current**: Set review cycles, update regularly
5. **Use Templates**: Generate from templates for consistency
6. **Validate Often**: Run validation after every change
7. **Build Incrementally**: Add 1-2 documents per week
8. **Involve Stakeholders**: Get input from document owners

## Conformance Report Example

```json
{
  "level": "Silver",
  "documentCount": 28,
  "coverage": 34.1,
  "domainCoverage": {
    "strategic-foundation": 4,
    "market-environment": 3,
    "customer-value": 3,
    "product-service": 4,
    "business-model": 3,
    "operations-execution": 4,
    "technology-data": 2,
    "financial-investment": 2,
    "risk-governance": 1,
    "growth-innovation": 1,
    "learning-decisions": 1,
    "brand-marketing": 0
  },
  "missingForNextLevel": [
    "brand-marketing domain needs 2+ documents"
  ],
  "recommendations": [
    "Add BRD and MSG to reach Gold level",
    "Expand risk-governance to 2+ documents"
  ]
}
```

## Next Steps

- View [All Document Types](/docs/document-types)
- Learn about [Document Structure](/spec/structure)
- Read the [Full Specification](/spec/v1-0-0)
- Get started with the [CLI Tool](/docs/tools/cli)
