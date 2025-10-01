# Core Concepts

Understanding these fundamental concepts will help you work effectively with BSpec.

## Document Structure

Every BSpec document follows a standardized structure with two parts:

### 1. YAML Frontmatter (Machine-Readable)

```yaml
---
# CORE IDENTITY
id: {TYPE}-{kebab-case-identifier}
type: {TYPE}
status: Draft|Review|Accepted|Deprecated
version: 1.0.0

# RELATIONSHIPS (creates knowledge graph)
depends_on: [document-ids]
enables: [document-ids]
conflicts_with: [document-ids]

# BUSINESS CONTEXT
domain: strategic|market|customer|product|...
priority: critical|high|medium|low
success_criteria: ["Measurable criteria"]
---
```

### 2. Markdown Content (Human-Readable)

Standard Markdown for human-readable documentation.

## The Atomic Document Principle

Each BSpec document is **atomic** - it covers exactly one business concern:

✅ **Good**: One Vision Statement document  
✅ **Good**: One Product Requirements document per product  
❌ **Bad**: All company strategies in one document  
❌ **Bad**: Vision + Mission + Values in one document  

This atomicity enables:
- **Reusability**: Documents can be referenced and reused
- **Clarity**: Clear boundaries prevent scope creep  
- **Granular Updates**: Change one thing without touching everything
- **AI Analysis**: Systems can reason about specific concerns

## Rich Relationship Model

Documents don't exist in isolation. BSpec creates a knowledge graph through four relationship types:

### Dependencies (`depends_on`)
What this document requires to be meaningful.

Example: A Product Roadmap depends on:
- Product Requirements (PRD)
- Market Strategy (STR)
- Resource Plan (CAP)

### Enablements (`enables`)
What this document makes possible.

Example: A Vision Statement enables:
- Mission Statement (MSN)
- Strategic Plan (STR)
- Objectives (OBJ)

### Conflicts (`conflicts_with`)
Mutually exclusive documents.

Example: A "Premium Brand" Positioning conflicts with:
- "Budget Leader" Positioning
- "Mass Market" Strategy

### Related (`related_to`)
Contextual connections.

Example: A Feature Specification relates to:
- Customer Journey Map (CJM)
- User Personas (CUS)
- Technical Architecture (ARC)

## 11 Business Domains

All 82 document types are organized into 11 comprehensive domains:

1. **Strategic Foundation** - Vision, mission, strategy, objectives
2. **Market Environment** - Competition, trends, opportunities, threats
3. **Customer Value** - Segments, personas, journeys, value propositions
4. **Product & Service** - Requirements, features, roadmaps, architecture
5. **Business Model** - Revenue, pricing, partnerships, channels
6. **Operations & Execution** - Processes, operations, team, projects
7. **Technology & Data** - Infrastructure, data, AI, integrations
8. **Financial & Investment** - Budget, forecasts, funding, investor relations
9. **Risk & Governance** - Risks, compliance, security, governance, legal
10. **Growth & Innovation** - Innovation, experiments, R&D, change management
11. **Learning & Decisions** - Hypotheses, insights, retrospectives, decisions

## Conformance Levels

BSpec provides three conformance levels to guide your documentation journey:

### Bronze (12+ documents)
Essential foundation covering all critical business areas. Minimum viable business documentation.

**Required Domains**: All 11 domains must be represented with at least one document.

### Silver (25+ documents)
Comprehensive documentation suitable for scaling businesses, investor due diligence, and serious business operations.

### Gold (45+ documents)
Complete documentation for mature businesses, public companies, or those requiring thorough governance and compliance.

[Learn more about conformance levels →](/spec/conformance)

## Industry Profiles

BSpec provides specialized profiles for different business types:

- **Software/SaaS**: Emphasis on product, technology, and user experience
- **Physical Products**: Focus on manufacturing, supply chain, and distribution
- **Service Businesses**: Highlights on delivery, customer relationships, and operations
- **Nonprofits**: Priorities on mission, impact, and stakeholder management

Each profile suggests which document types are most important for that industry.

## Status Workflow

Documents progress through a standard lifecycle:

```
Draft → Review → Accepted → [Deprecated]
```

- **Draft**: Work in progress, not yet complete
- **Review**: Complete and awaiting approval
- **Accepted**: Approved and in use
- **Deprecated**: No longer current, replaced by newer version

## Versioning

BSpec uses semantic versioning (MAJOR.MINOR.PATCH) for documents:

- **MAJOR**: Fundamental changes to the document's purpose or structure
- **MINOR**: Significant updates that add or remove content sections
- **PATCH**: Small corrections, clarifications, or minor updates

Example: `1.0.0` → `1.1.0` (added new section) → `2.0.0` (restructured entire document)

## Machine-Readable Format

The YAML frontmatter makes BSpec documents:

- **Parseable**: Tools can read and validate documents
- **Queryable**: Search and filter by any metadata field
- **Visualizable**: Generate graphs, reports, and dashboards
- **AI-Compatible**: Enable AI systems to understand business structure

This enables powerful tooling including:
- Automated validation
- Relationship visualization
- Gap analysis
- Conformance checking
- AI-powered recommendations

## Next Steps

- [Quick Start Guide](/docs/quickstart) - Create your first document
- [Document Types](/docs/document-types) - Explore all 82 types
- [CLI Tool](/docs/tools/cli) - Learn to use the command-line tool
- [MCP Server](/docs/tools/mcp) - Integrate with AI agents
