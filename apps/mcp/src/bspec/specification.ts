/**
 * BSpec 1.0 Specification Bundle
 *
 * Contains the complete BSpec specification embedded for Cloudflare Workers deployment
 * This avoids external file dependencies and ensures everything is packaged together
 */

export const BSPEC_SPECIFICATION = `# BSpec 1.0: Business Specification Standard

## Overview

The Business Specification Standard (BSpec) 1.0 is a universal language for describing any business as a structured, machine-readable knowledge graph. It provides a comprehensive framework for documenting, analyzing, and evolving business models, strategies, and operations.

## Core Principles

1. **Universal Vocabulary**: 82 document types covering all aspects of business
2. **Machine Readable**: Structured format for AI and automation
3. **Relationship-Driven**: Documents connected through dependency graphs
4. **Conformance Levels**: Bronze (12+), Silver (25+), Gold (45+) documents
5. **Industry Agnostic**: Applicable to any business domain

## Document Structure

Every BSpec document follows this standard format:

\`\`\`yaml
---
# CORE IDENTITY
id: {TYPE}-{kebab-case-identifier}
title: "Human Readable Document Title"
type: {TYPE}
status: Draft|Review|Accepted|Deprecated
version: 1.0.0

# OWNERSHIP & RESPONSIBILITY
owner: Team/Person
stakeholders: [Person1, Team2]
reviewers: [Person1, Person2]

# TEMPORAL METADATA
created: YYYY-MM-DD
updated: YYYY-MM-DD
review_cycle: monthly|quarterly|annually

# RELATIONSHIP GRAPH
depends_on: [{document-ids}]
enables: [{document-ids}]
conflicts_with: [{document-ids}]
related: [{document-ids}]

# BUSINESS CONTEXT
domain: strategic|market|customer|product|model|operations|technology|financial|risk|growth|learning
priority: critical|high|medium|low
success_criteria:
  - "Measurable success criterion 1"
  - "Measurable success criterion 2"

# VALIDATION & MEASUREMENT
assumptions: [list of key assumptions]
constraints: [list of key constraints]
risks: [{risk-document-ids}]
metrics: [{metric-document-ids}]
---

# {Document Title}

## Overview
*Brief description of purpose and business context*

## {Type-Specific Content}
*Content structure varies by document type*

## Dependencies
*What this document requires to be valid*

## Enables
*What this document makes possible*

## Implementation
*How this translates to concrete actions*

## Validation
*How to verify correctness and effectiveness*
\`\`\`

## Document Types by Domain

### 1. Strategic Foundation (8 types)
- **MSN** (Mission): Core purpose and reason for existence
- **VSN** (Vision): Future state and aspirational goals
- **VAL** (Values): Fundamental beliefs and principles
- **STR** (Strategy): High-level approach to achieving vision
- **OBJ** (Objectives): Specific, measurable outcomes
- **MOT** (Motivation): Driving forces and rationale
- **PUR** (Purpose): Why the organization exists
- **THY** (Theory): Underlying assumptions and hypotheses

### 2. Market & Environment (10 types)
- **MKT** (Market): Target market definition and analysis
- **SEG** (Segments): Customer segmentation strategy
- **CMP** (Competition): Competitive landscape analysis
- **POS** (Positioning): Market positioning strategy
- **TRN** (Trends): Market trends and patterns
- **ECO** (Economics): Economic factors and conditions
- **OPP** (Opportunities): Market opportunities identified
- **THR** (Threats): External threats and challenges
- **REG** (Regulations): Regulatory environment and compliance
- **MAC** (Macroeconomics): Broader economic context

### 3. Customer & Value (12 types)
- **PER** (Personas): User archetypes and characteristics
- **JTB** (Jobs-to-be-Done): Customer needs and outcomes
- **CJM** (Customer Journey): Customer experience mapping
- **USE** (Use Cases): Specific usage scenarios
- **STO** (Stories): User stories and narratives
- **PAI** (Pain Points): Customer problems and frustrations
- **GAI** (Gains): Customer benefits and value
- **EMP** (Empathy): Customer emotional understanding
- **FEE** (Feedback): Customer input and insights
- **INT** (Interactions): Customer touchpoints
- **SUR** (Surveys): Customer research data
- **BEH** (Behavior): Customer behavior patterns

### 4. Product & Service (10 types)
- **PRD** (Product): Product specifications and features
- **SVC** (Service): Service offerings and delivery
- **FEA** (Features): Specific product capabilities
- **ROD** (Roadmap): Product development timeline
- **REQ** (Requirements): Functional and technical needs
- **QUA** (Quality): Quality standards and measures
- **UXD** (UX Design): User experience design
- **PER** (Performance): Performance metrics and SLAs
- **INT** (Integration): System integration requirements
- **SUP** (Support): Customer support and maintenance

### 5. Business Model (12 types)
- **BMC** (Business Model Canvas): Core business model
- **REV** (Revenue): Revenue streams and models
- **PRC** (Pricing): Pricing strategy and structure
- **CST** (Costs): Cost structure and drivers
- **CHN** (Channels): Distribution and sales channels
- **REL** (Relationships): Customer relationship types
- **RES** (Resources): Key resources required
- **ACT** (Activities): Key business activities
- **PRT** (Partners): Key partnerships and alliances
- **UNT** (Unit Economics): Per-unit financial metrics
- **LTV** (Lifetime Value): Customer lifetime value
- **CAC** (Customer Acquisition Cost): Cost to acquire customers

### 6. Operations & Execution (12 types)
- **PRC** (Processes): Business processes and workflows
- **WFL** (Workflows): Operational workflows
- **ORG** (Organization): Organizational structure
- **ROL** (Roles): Job roles and responsibilities
- **TEA** (Teams): Team structures and composition
- **SKI** (Skills): Required skills and competencies
- **POL** (Policies): Operating policies and guidelines
- **SLA** (Service Level Agreements): Performance commitments
- **VND** (Vendors): Supplier relationships
- **FAC** (Facilities): Physical and virtual facilities
- **TOO** (Tools): Technology tools and systems
- **CAP** (Capabilities): Organizational capabilities

### 7. Technology & Data (8 types)
- **ARC** (Architecture): Technical architecture
- **SYS** (Systems): Information systems
- **DAT** (Data): Data models and management
- **API** (APIs): Application programming interfaces
- **INF** (Infrastructure): Technology infrastructure
- **SEC** (Security): Security requirements and measures
- **DEV** (Development): Development processes
- **ANA** (Analytics): Data analytics and insights

### 8. Financial & Investment (10 types)
- **FIN** (Financial): Financial models and projections
- **BUD** (Budget): Budget allocations and planning
- **FOR** (Forecasts): Financial forecasts
- **FND** (Funding): Funding sources and requirements
- **INV** (Investment): Investment strategy and allocation
- **VAL** (Valuation): Business valuation methods
- **MET** (Metrics): Key performance indicators
- **REP** (Reports): Financial reporting requirements
- **AUD** (Audit): Audit requirements and processes
- **TAX** (Tax): Tax strategy and compliance

### 9. Risk & Governance (8 types)
- **RSK** (Risk): Risk assessment and management
- **MIT** (Mitigation): Risk mitigation strategies
- **CMP** (Compliance): Regulatory compliance
- **GVN** (Governance): Corporate governance
- **CTL** (Controls): Internal controls and procedures
- **CRI** (Crisis): Crisis management planning
- **ETH** (Ethics): Ethical guidelines and standards
- **STA** (Stakeholders): Stakeholder management

### 10. Growth & Innovation (8 types)
- **GTM** (Go-to-Market): Market entry strategy
- **GRW** (Growth): Growth strategy and tactics
- **SCL** (Scale): Scaling operations and systems
- **EXP** (Expansion): Market expansion plans
- **INN** (Innovation): Innovation processes
- **RND** (Research & Development): R&D activities
- **ACQ** (Acquisition): Acquisition strategy
- **EXP** (Experiments): Business experiments and tests

### 11. Learning & Decisions (6 types)
- **DEC** (Decisions): Decision records and rationale
- **LRN** (Learning): Organizational learning
- **RET** (Retrospectives): Process improvement reviews
- **HYP** (Hypotheses): Business hypotheses and assumptions
- **KNO** (Knowledge): Knowledge management
- **WIS** (Wisdom): Institutional wisdom and insights

## Conformance Levels

### Bronze Level (12+ documents)
Minimum viable business specification covering:
- Core identity (MSN, VSN, VAL)
- Key strategy (STR, OBJ)
- Market understanding (MKT, SEG)
- Value proposition (PER, JTB)
- Business model (BMC, REV)
- Basic operations (PRC)

### Silver Level (25+ documents)
Investment-ready specification adding:
- Detailed market analysis
- Customer research and insights
- Product/service specifications
- Financial projections
- Risk management
- Operational processes

### Gold Level (45+ documents)
Operational excellence specification including:
- Comprehensive documentation across all domains
- Detailed relationship mapping
- Advanced analytics and metrics
- Innovation and growth planning
- Full governance framework

## Industry Profiles

### Software/SaaS Profile
Required additional focus on:
- **ARC** (Technical Architecture)
- **API** (API Design and Management)
- **SEC** (Security and Compliance)
- **SUP** (Customer Support Systems)
- **MET** (SaaS Metrics: MRR, Churn, etc.)

### Physical Product Profile
Required additional focus on:
- **INF** (Manufacturing Infrastructure)
- **QUA** (Quality Control and Testing)
- **VND** (Supply Chain Management)
- **REG** (Product Safety Regulations)

### Service Business Profile
Required additional focus on:
- **PRC** (Service Delivery Processes)
- **SLA** (Service Level Agreements)
- **SKI** (Service Skills and Training)
- **QUA** (Service Quality Standards)

### Nonprofit Profile
Required additional focus on:
- **PUR** (Mission and Purpose)
- **STA** (Stakeholder Management)
- **MET** (Impact Measurement)
- **GVN** (Board Governance)

## Validation Rules

### Document Structure Validation
1. **ID Format**: Must follow {TYPE}-{kebab-case-name} pattern
2. **Type Consistency**: Document type must match ID prefix
3. **Version Format**: Must follow semantic versioning (x.y.z)
4. **Date Sequence**: Created date must be before or equal to updated date
5. **Required Fields**: All core fields must be present and valid

### Relationship Validation
1. **Dependency Existence**: All referenced dependencies must exist
2. **Circular Dependencies**: No circular dependency chains allowed
3. **Conflict Resolution**: Conflicting documents cannot both be accepted
4. **Domain Alignment**: Document types should align with expected domains

### Content Validation
1. **Minimum Content**: Documents should have substantive content (>100 words)
2. **No Placeholders**: Remove TODO, TBD, and placeholder content
3. **Success Criteria**: Define measurable success criteria
4. **Implementation Details**: Include actionable implementation guidance

## Machine Readability

### JSON Schema Support
Every document type has corresponding JSON schema for:
- Automated validation
- AI/ML processing
- Tool integration
- Quality assurance

### AI Integration Patterns
Standard approaches for:
- Document generation from context
- Relationship analysis and mapping
- Conformance assessment
- Gap analysis and recommendations

### Query Capabilities
Support for complex queries including:
- Document type filtering
- Domain-based searches
- Relationship traversal
- Status and ownership filtering
- Full-text content search

## Ecosystem Integration

### Model Context Protocol (MCP)
Standard MCP server implementation provides:
- Document parsing and validation
- Relationship analysis
- Query and search capabilities
- Format conversion (Markdown ↔ JSON)
- Business intelligence insights

### SDK Support
Language-specific SDKs available for:
- TypeScript/JavaScript
- Python
- Go
- C#/.NET
- Java

### Tool Ecosystem
Compatible with:
- Business planning software
- Project management tools
- Documentation platforms
- AI/ML frameworks
- Analytics platforms

## Getting Started

1. **Choose Your Level**: Start with Bronze (12 docs) for MVP
2. **Select Industry Profile**: Use specialized requirements if applicable
3. **Begin with Strategy**: Start with MSN, VSN, VAL documents
4. **Build Incrementally**: Add documents based on business priorities
5. **Validate Continuously**: Use validation tools to ensure quality
6. **Iterate and Improve**: Regular reviews and updates

## Resources

- **Specification**: Complete technical specification
- **Templates**: Document templates for all 82 types
- **Examples**: Real-world implementation examples
- **Tools**: SDKs, validators, and analysis tools
- **Community**: Support and best practices sharing

---

*BSpec 1.0 - A universal language for business*
*For more information, visit: https://bspec.dev*
`;

// Document type definitions with detailed information
export const DOCUMENT_TYPES = {
  // Strategic Foundation (8)
  'MSN': {
    name: 'Mission',
    domain: 'strategic',
    description: 'Core purpose and reason for existence',
    template: `---
id: MSN-{identifier}
title: "{Organization} Mission Statement"
type: MSN
status: Draft
version: 1.0.0
owner: Leadership Team
stakeholders: [All Teams]
created: {date}
updated: {date}
domain: strategic
priority: critical
success_criteria:
  - "Mission is understood and embraced by all stakeholders"
  - "Mission guides strategic decisions and priorities"
---

# {Organization} Mission Statement

## Overview
Our mission defines why we exist and what we aim to achieve in the world.

## Mission Statement
[Clear, concise statement of organizational purpose]

## Core Elements
- **Purpose**: Why we exist
- **Impact**: What change we create
- **Beneficiaries**: Who we serve
- **Approach**: How we operate

## Alignment
This mission aligns with our vision and values to create a coherent strategic foundation.

## Implementation
- Communicate mission throughout organization
- Use mission to guide strategic decisions
- Regularly review and refine mission
- Measure alignment with mission in all activities

## Validation
- Stakeholder understanding surveys
- Decision alignment assessments
- Mission-driven outcome tracking`
  },

  'VSN': {
    name: 'Vision',
    domain: 'strategic',
    description: 'Future state and aspirational goals',
    template: `---
id: VSN-{identifier}
title: "{Organization} Vision Statement"
type: VSN
status: Draft
version: 1.0.0
owner: Leadership Team
stakeholders: [All Teams]
depends_on: [MSN-{mission-id}]
created: {date}
updated: {date}
domain: strategic
priority: critical
success_criteria:
  - "Vision inspires and motivates stakeholders"
  - "Vision provides clear direction for strategy"
---

# {Organization} Vision Statement

## Overview
Our vision describes the future state we aspire to create.

## Vision Statement
[Inspiring, aspirational statement of future desired state]

## Key Elements
- **Future State**: What success looks like
- **Timeline**: When we aim to achieve this
- **Impact**: The change we'll create
- **Differentiation**: What makes our vision unique

## Strategic Alignment
This vision builds upon our mission and guides our strategic objectives.

## Implementation
- Communicate vision consistently
- Align all strategies with vision
- Track progress toward vision
- Celebrate vision-aligned achievements

## Validation
- Vision clarity and inspiration metrics
- Strategic alignment assessments
- Progress tracking toward vision elements`
  }

  // Add more document types as needed...
};

// Domain information
export const DOMAINS = {
  strategic: {
    name: 'Strategic Foundation',
    description: 'Core organizational identity, purpose, and strategic direction',
    documentTypes: ['MSN', 'VSN', 'VAL', 'STR', 'OBJ', 'MOT', 'PUR', 'THY']
  },
  market: {
    name: 'Market & Environment',
    description: 'Market analysis, competitive landscape, and external factors',
    documentTypes: ['MKT', 'SEG', 'CMP', 'POS', 'TRN', 'ECO', 'OPP', 'THR', 'REG', 'MAC']
  },
  customer: {
    name: 'Customer & Value',
    description: 'Customer understanding, needs, and value delivery',
    documentTypes: ['PER', 'JTB', 'CJM', 'USE', 'STO', 'PAI', 'GAI', 'EMP', 'FEE', 'INT', 'SUR', 'BEH']
  },
  product: {
    name: 'Product & Service',
    description: 'Product and service specifications, features, and delivery',
    documentTypes: ['PRD', 'SVC', 'FEA', 'ROD', 'REQ', 'QUA', 'UXD', 'PER', 'INT', 'SUP']
  },
  model: {
    name: 'Business Model',
    description: 'Revenue model, pricing, costs, and value creation',
    documentTypes: ['BMC', 'REV', 'PRC', 'CST', 'CHN', 'REL', 'RES', 'ACT', 'PRT', 'UNT', 'LTV', 'CAC']
  },
  operations: {
    name: 'Operations & Execution',
    description: 'Operational processes, organization, and execution capabilities',
    documentTypes: ['PRC', 'WFL', 'ORG', 'ROL', 'TEA', 'SKI', 'POL', 'SLA', 'VND', 'FAC', 'TOO', 'CAP']
  },
  technology: {
    name: 'Technology & Data',
    description: 'Technical architecture, systems, and data management',
    documentTypes: ['ARC', 'SYS', 'DAT', 'API', 'INF', 'SEC', 'DEV', 'ANA']
  },
  financial: {
    name: 'Financial & Investment',
    description: 'Financial planning, budgets, metrics, and investment strategy',
    documentTypes: ['FIN', 'BUD', 'FOR', 'FND', 'INV', 'VAL', 'MET', 'REP', 'AUD', 'TAX']
  },
  risk: {
    name: 'Risk & Governance',
    description: 'Risk management, compliance, and governance frameworks',
    documentTypes: ['RSK', 'MIT', 'CMP', 'GVN', 'CTL', 'CRI', 'ETH', 'STA']
  },
  growth: {
    name: 'Growth & Innovation',
    description: 'Growth strategies, innovation, and market expansion',
    documentTypes: ['GTM', 'GRW', 'SCL', 'EXP', 'INN', 'RND', 'ACQ', 'EXP']
  },
  learning: {
    name: 'Learning & Decisions',
    description: 'Decision making, learning, and knowledge management',
    documentTypes: ['DEC', 'LRN', 'RET', 'HYP', 'KNO', 'WIS']
  }
};

// Example documents for reference
export const EXAMPLE_DOCUMENTS = {
  'MSN': `---
id: MSN-stratus-ai-cloud
title: "Stratus AI Cloud Mission Statement"
type: MSN
status: Accepted
version: 1.0.0
owner: CEO
stakeholders: [Leadership Team, All Employees, Investors]
created: 2024-01-15
updated: 2024-01-15
domain: strategic
priority: critical
success_criteria:
  - "Mission drives all strategic decisions and product development"
  - "100% of employees can articulate the mission clearly"
  - "Mission differentiates us in the market and guides brand positioning"
---

# Stratus AI Cloud Mission Statement

## Overview
Our mission defines the fundamental purpose of Stratus AI Cloud and guides every aspect of our business strategy, product development, and organizational culture.

## Mission Statement
**To democratize artificial intelligence by providing the world's most accessible, powerful, and secure AI infrastructure that enables any organization to harness the transformative power of machine learning.**

## Core Elements

### Purpose
We exist to break down the barriers that prevent organizations from accessing and implementing AI solutions effectively.

### Impact
We enable digital transformation and innovation across industries by making advanced AI capabilities accessible to organizations of all sizes.

### Beneficiaries
- Enterprises seeking AI transformation
- Developers building AI-powered applications
- Researchers advancing machine learning
- Startups innovating with AI technology

### Approach
Through cloud-native infrastructure, intuitive tools, and comprehensive support that removes technical complexity while maintaining enterprise-grade reliability and security.

## Strategic Alignment
This mission directly supports our vision of becoming the leading AI infrastructure platform and aligns with our core values of accessibility, innovation, and security.

## Implementation Guidelines
- Every product decision must advance AI accessibility
- All features should reduce complexity while maintaining power
- Security and compliance are non-negotiable requirements
- Customer success is measured by AI implementation success

## Success Metrics
- Customer AI project success rates
- Time-to-deployment for AI solutions
- Platform adoption across different organization sizes
- Customer satisfaction with accessibility and ease of use`,

  'PER': `---
id: PER-enterprise-ai-director
title: "Enterprise AI Director Persona"
type: PER
status: Accepted
version: 1.0.0
owner: Product Team
stakeholders: [Engineering, Sales, Marketing]
depends_on: [JTB-ai-transformation]
enables: [CJM-enterprise-onboarding, USE-model-deployment]
created: 2024-01-15
updated: 2024-01-15
domain: customer
priority: high
success_criteria:
  - "Product features align with persona needs and pain points"
  - "Marketing messages resonate with persona motivations"
  - "Sales team can effectively qualify and engage this persona"
---

# Enterprise AI Director Persona

## Overview
The Enterprise AI Director is a key decision-maker responsible for driving AI transformation initiatives within large organizations.

## Demographics
- **Role**: Director/VP of AI, Chief Data Officer, Head of Digital Transformation
- **Company Size**: 1,000+ employees, $100M+ revenue
- **Industry**: Technology, Financial Services, Healthcare, Manufacturing
- **Experience**: 8-15 years in technology leadership, 3-5 years in AI/ML
- **Education**: Advanced degree in Computer Science, Engineering, or Business

## Psychographics
- **Motivations**: Drive competitive advantage through AI, modernize legacy systems, demonstrate ROI on AI investments
- **Concerns**: Integration complexity, security risks, vendor lock-in, skills gaps
- **Values**: Innovation, reliability, strategic thinking, team development
- **Communication Style**: Data-driven, strategic, collaborative

## Goals & Objectives
### Primary Goals
- Successfully implement AI across multiple business units
- Demonstrate measurable business impact from AI initiatives
- Build internal AI capabilities and expertise
- Ensure security and compliance in AI deployments

### Success Metrics
- AI project deployment success rate >80%
- Time-to-value for AI initiatives <6 months
- ROI demonstration within first year
- Zero security incidents in AI implementations

## Pain Points
### Technical Challenges
- Complex integration with existing enterprise systems
- Difficulty scaling AI models across the organization
- Managing multiple AI tools and platforms
- Ensuring data quality and governance

### Organizational Challenges
- Limited AI expertise within the team
- Resistance to change from business units
- Budget constraints and ROI pressure
- Vendor evaluation and selection complexity

## Jobs-to-be-Done
1. **Evaluate AI Infrastructure**: Assess platforms that can support enterprise-scale AI deployment
2. **Plan AI Strategy**: Develop comprehensive AI transformation roadmap
3. **Manage AI Projects**: Oversee multiple concurrent AI implementation projects
4. **Build AI Team**: Recruit, train, and develop AI talent within the organization
5. **Ensure Governance**: Implement AI ethics, security, and compliance frameworks

## Technology Usage
- **Current Tools**: Existing enterprise software (SAP, Salesforce, etc.), cloud platforms (AWS, Azure, GCP)
- **AI Experience**: Familiar with major AI/ML frameworks, cloud AI services
- **Evaluation Criteria**: Scalability, security, integration capabilities, vendor support
- **Decision Process**: Technical evaluation → stakeholder buy-in → pilot project → full deployment

## Preferred Channels
- Industry conferences and events
- Technical webinars and whitepapers
- Peer recommendations and case studies
- Analyst reports (Gartner, Forrester)
- LinkedIn and professional networks

## Buying Journey
1. **Awareness**: Identify need for AI infrastructure platform
2. **Research**: Evaluate multiple vendors and solutions
3. **Consideration**: Conduct pilots and technical assessments
4. **Decision**: Negotiate terms and select platform
5. **Implementation**: Deploy and scale across organization

## Engagement Strategy
- Provide technical deep-dives and architecture guidance
- Share relevant case studies from similar enterprises
- Offer pilot programs and proof-of-concept opportunities
- Connect with technical support and solution architects
- Facilitate peer connections and reference customers`
};