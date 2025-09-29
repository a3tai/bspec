# BSpec 1.0: Universal Business Specification Standard

**Version:** 1.0.0
**Status:** Public Draft
**Website:** [bspec.dev](https://bspec.dev) *(coming soon)*
**Documentation:** [bspec.dev/docs](https://bspec.dev/docs) *(coming soon)*
**License:** CC BY 4.0

## What is BSpec?

Business Specification Standard (BSpec) is the universal language for describing any business as a structured, machine-readable knowledge graph. Like how SAFE agreements standardized early-stage investing, BSpec standardizes business documentation to enable systematic analysis, automated generation, and intelligent decision-making.

**Key Innovation**: Treat every business as a system of atomic, interconnected documents that together form a complete picture of strategy, operations, and execution.

## The Problem BSpec Solves

Today's business documentation is fragmented, inconsistent, and impossible to analyze systematically:

- **Entrepreneurs** reinvent business planning without structured guidance
- **Investors** compare deals using incompatible formats and incomplete information
- **AI systems** cannot understand business context to provide intelligent assistance
- **Organizations** lose knowledge when people leave or contexts change

## The BSpec Solution

BSpec provides a **standardized vocabulary** and **relationship model** for describing any business completely:

```
Business = Graph of Connected Documents
Each Document = One Business Concern + Rich Metadata
Relationships = How Documents Depend On, Enable, or Conflict With Each Other
```

This enables:
- **Systematic thinking** through templates and dependencies
- **Automated analysis** by AI systems that understand business structure
- **Intelligent generation** of business artifacts (pitch decks, code, marketing)
- **Consistent evaluation** across different businesses and industries

## Quick Start

### Understanding BSpec Documents

Every BSpec document follows this structure:

```yaml
---
# Core Identity
id: STR-platform-strategy
title: "AI Platform Strategy"
type: STR
status: Draft|Review|Accepted|Deprecated
version: 1.0.0

# Relationships (creates business knowledge graph)
depends_on: [MSN-mission, MKT-ai-market]
enables: [PRD-inference-api, GTM-developer-first]
conflicts_with: [STR-consumer-focus]

# Business Context
domain: strategic
priority: critical
success_criteria:
  - "Platform adoption exceeds 1000 developers by Q4"
  - "API revenue reaches $100k MRR within 12 months"
---

# Platform Strategy Content
*Document content in structured Markdown...*
```

### The Complete Business Vocabulary

BSpec defines document types across comprehensive business domains:

#### 🎯 Strategic Foundation (8 types)
**MSN** Mission • **VSN** Vision • **VAL** Values • **STR** Strategy • **OBJ** Objectives • **MOT** Moats • **PUR** Purpose • **THY** Theory of Change

#### 🌍 Market & Environment (10 types)
**MKT** Market Definition • **SEG** Market Segments • **CMP** Competitive Analysis • **POS** Positioning • **TRN** Trends • **ECO** Ecosystem • **OPP** Opportunities • **THR** Threats • **REG** Regulatory Environment • **MAC** Macro Environment

#### 👥 Customer & Value (12 types)
**PER** Personas • **JTB** Jobs-to-be-Done • **CJM** Customer Journey • **USE** Use Cases • **STO** User Stories • **PAI** Pain Points • **GAI** Gains • **EMP** Empathy Maps • **FEE** Feedback • **INT** Interviews • **SUR** Surveys • **BEH** Behaviors

#### 📦 Product & Service (10 types)
**PRD** Products • **SVC** Services • **FEA** Features • **ROD** Roadmap • **REQ** Requirements • **QUA** Quality Standards • **UXD** User Experience • **PER** Performance • **INT** Integrations • **SUP** Support

#### 💰 Business Model (12 types)
**BMC** Business Model Canvas • **REV** Revenue Streams • **PRC** Pricing • **CST** Cost Structure • **CHN** Channels • **REL** Customer Relationships • **RES** Key Resources • **ACT** Key Activities • **PRT** Key Partnerships • **UNT** Unit Economics • **LTV** Lifetime Value • **CAC** Customer Acquisition

#### ⚙️ Operations & Execution (12 types)
**PRC** Processes • **WFL** Workflows • **ORG** Organization • **ROL** Roles • **TEA** Teams • **SKI** Skills • **POL** Policies • **SLA** Service Levels • **VND** Vendors • **FAC** Facilities • **TOO** Tools • **CAP** Capabilities

#### 🔧 Technology & Data (8 types)
**ARC** Architecture • **SYS** Systems • **DAT** Data Models • **API** APIs • **INF** Infrastructure • **SEC** Security • **DEV** Development • **ANA** Analytics

#### 📊 Financial & Investment (10 types)
**FIN** Financial Model • **BUD** Budget • **FOR** Forecasts • **FND** Funding • **INV** Investment • **VAL** Valuation • **MET** Metrics • **REP** Reporting • **AUD** Audit • **TAX** Tax Strategy

#### ⚠️ Risk & Governance (8 types)
**RSK** Risks • **MIT** Mitigations • **CMP** Compliance • **GVN** Governance • **CTL** Controls • **CRI** Crisis Management • **ETH** Ethics • **STA** Stakeholders

#### 📈 Growth & Innovation (8 types)
**GTM** Go-to-Market • **GRW** Growth Model • **SCL** Scaling • **EXP** Experiments • **INN** Innovation • **RND** Research • **ACQ** Acquisitions • **EXP** Expansion

#### 🧠 Learning & Decisions (6 types)
**DEC** Decisions • **LRN** Learnings • **RET** Retrospectives • **HYP** Hypotheses • **KNO** Knowledge • **WIS** Wisdom

### Conformance Levels

**🥉 Bronze: Minimum Viable Business Spec (12+ documents)**
- Strategic core (MSN, VSN, VAL)
- Customer understanding (2+ PER with JTB)
- Solution definition (1+ PRD/SVC)
- Business model (REV, CST)
- Risk awareness (2+ RSK with MIT)

**🥈 Silver: Investment Ready (25+ documents)**
- Bronze foundation + strategic depth + market analysis
- Complete business model and financial foundation
- Growth strategy and key metrics

**🥇 Gold: Operational Excellence (45+ documents)**
- Silver foundation + operational processes
- Organization design and technology architecture
- Compliance framework and innovation pipeline

### Industry Profiles

Specialized requirements for different business types:
- **Software/SaaS**: ARC, API, SEC, SUP + SaaS metrics (MRR, CAC, LTV, Churn, NRR)
- **Physical Product**: INF, QUA, VND, REG + supply chain and manufacturing
- **Service Business**: PRC, SLA, SKI, QUA + service delivery and capacity planning
- **Nonprofit**: PUR, STA, MET, GVN + impact measurement and accountability

## Repository Structure

```
BSpec-Foundations/
├── spec/v1/                     # BSpec 1.0 Specification (COMPLETE)
│   ├── spec.md                  # Complete specification document
│   ├── strategic-foundation/    # Strategic document types (MSN, VSN, VAL, etc.)
│   ├── market-environment/      # Market analysis document types
│   ├── customer-value/          # Customer-focused document types
│   ├── product-service/         # Product and service document types
│   ├── business-model/          # Business model document types
│   ├── operations-execution/    # Operations document types
│   ├── technology-data/         # Technology document types
│   ├── financial-investment/    # Financial document types
│   ├── risk-governance/         # Risk and governance document types
│   ├── growth-innovation/       # Growth document types
│   ├── learning-decisions/      # Learning document types
│   └── brand-marketing/         # Brand and marketing document types
├── apps/                        # BSpec Ecosystem Applications
│   ├── web/                     # bspec.dev website (SvelteKit, Apple/Liquid Glass design)
│   └── mcp/                     # Model Context Provider server (Cloudflare Workers)
├── sdk/                         # Software Development Kits
│   ├── v1/
│   │   ├── typescript/          # TypeScript SDK with complete type definitions
│   │   ├── python/              # Python SDK with Pydantic models
│   │   ├── go/                  # Go SDK with struct definitions
│   │   ├── rust/                # Rust SDK with type definitions
│   │   ├── json/                # JSON schemas for all document types
│   │   └── format/              # Format specifications and validation
│   └── cli/                     # Command-line interface (Go binary)
├── scripts/                     # Code generation and utilities
│   ├── generators/              # SDK generators for all languages
│   ├── parsers/                 # Specification parsers
│   └── templates/               # Code generation templates
├── examples/                    # Reference implementations
└── README.md                    # This file
```

## Planned Ecosystem (bspec.dev)

### 🌐 bspec.dev - Main Website
Modern SaaS marketing site with Apple-inspired design and Liquid Glass styling:
- Getting started guides and tutorials
- Interactive document type explorer
- Community showcases and case studies
- Download links for tools and SDKs

### 📚 bspec.dev/docs - Documentation Hub
Comprehensive technical documentation:
- Complete BSpec 1.0 specification
- SDK documentation and API references
- Implementation guides and best practices
- Industry-specific profiles and examples

### 🤖 mcp.bspec.dev/mcp - AI Integration Server
Hosted MCP (Model Context Provider) server for AI agents:
- Parse and validate BSpec documents
- Convert between Markdown and JSON formats
- Query business knowledge graphs
- Generate artifacts from specifications

### 📦 .bspec File Format
Packaging format for complete business specifications:
```
business-spec.bspec  (like .tgz archive)
├── documents/       # All BSpec documents
├── assets/          # Images, diagrams, attachments
├── manifest.json    # Metadata and validation
└── relationships/   # Computed relationship graph
```

## Machine Readability & AI Integration

BSpec is designed for intelligent consumption:

### JSON Schema
Every document type has a corresponding JSON schema for validation and AI consumption.

### AI Prompt Templates
Standard prompts for document generation, business analysis, and artifact creation.

### API Integration
```typescript
interface BSpecAPI {
  // Document management
  getDocument(id: string): BSpecDocument
  createDocument(doc: BSpecDocument): void
  updateDocument(id: string, updates: Partial<BSpecDocument>): void

  // Relationship queries
  getDependencies(id: string, depth?: number): BSpecDocument[]
  getImpacts(id: string, depth?: number): BSpecDocument[]
  findConflicts(id: string): BSpecDocument[]

  // Analysis
  validateConformance(profile?: IndustryProfile): ConformanceReport
  findGaps(targetLevel: ConformanceLevel): DocumentGap[]
  analyzeReadiness(milestone: string): ReadinessReport

  // Generation
  generateDocument(type: DocumentType, context: BusinessContext): BSpecDocument
  generateArtifact(type: ArtifactType, documents: BSpecDocument[]): Artifact
}
```

## Development Status

| Component | Status | Description |
|-----------|--------|-------------|
| **BSpec 1.0 Spec** | ✅ Complete | All document types across many domains, relationship model, conformance levels |
| **TypeScript SDK** | ✅ Complete | Full TypeScript definitions, JSON schemas, validation |
| **Python SDK** | ✅ Complete | Pydantic models, complete type definitions |
| **Go SDK** | ✅ Complete | Go structs, comprehensive type system |
| **Rust SDK** | ✅ Complete | Rust types with serde support |
| **JSON Schemas** | ✅ Complete | Complete JSON schema definitions for all 82 document types |
| **CLI Tool** | ✅ Complete | Go binary with init, query, pack, extract, chat commands |
| **MCP Server** | ✅ Complete | Cloudflare Workers-based MCP server for AI agents |
| **Web Application** | ✅ Complete | SvelteKit website with Apple/Liquid Glass design |
| **Code Generation** | ✅ Complete | Automated SDK generation from specification |
| **.bspec Format** | ✅ Complete | Packaged archive format with validation |

## Getting Started

### Option 1: Use the CLI Tool

1. **Install the CLI**:
   ```bash
   # From the repository
   cd sdk/cli
   make install

   # Verify installation
   bspec --help
   ```

2. **Initialize a new BSpec project**:
   ```bash
   bspec init my-business-spec
   cd my-business-spec
   ```

3. **Start with strategic documents**:
   ```bash
   # Interactive AI-assisted document creation
   bspec chat
   ```

### Option 2: Use the SDKs

**TypeScript/JavaScript**:
```typescript
import { BSpecDocument, validateDocument } from '@bspec/typescript-sdk';

// Load and validate a BSpec document
const doc: BSpecDocument = loadFromMarkdown(content);
const validation = validateDocument(doc);
```

**Python**:
```python
from bspec_sdk import BSpecDocument, DocumentValidator

# Create and validate documents
doc = BSpecDocument.from_markdown(content)
validator = DocumentValidator()
result = validator.validate(doc)
```

**Go**:
```go
import "github.com/a3tai/bspec/sdk/v1/go"

// Parse and validate BSpec documents
doc, err := bspec.ParseDocument(content)
validation := bspec.ValidateDocument(doc)
```

### Option 3: Manual Approach

1. **Read the Specification**: Start with `spec/v1/spec.md` for the complete framework
2. **Choose Your Conformance Target**: Bronze (12 docs), Silver (25 docs), or Gold (45 docs)
3. **Pick Your Industry Profile**: Software/SaaS, Physical Product, Service Business, or Nonprofit
4. **Start with Strategic Foundation**: Create MSN, VSN, and VAL documents first

## Contributing

BSpec is an open standard designed to evolve with the business world:

- **Specification Improvements**: Submit RFCs for new document types or relationship patterns
- **Industry Profiles**: Contribute specialized requirements for your industry
- **Tool Development**: Build tools that consume or generate BSpec documents
- **Examples**: Share real-world BSpec implementations (anonymized)

## License

This specification is released under **Creative Commons Attribution 4.0 International License**. You are free to share, adapt, and build upon this material for any purpose, including commercial use, as long as you provide appropriate attribution to the BSpec Foundation.

---

**BSpec 1.0 is the beginning, not the end.** It's the foundation for a future where every business decision is informed by complete context, where AI can truly understand and assist with business challenges, and where the complexity of modern business becomes manageable through systematic thinking and intelligent automation.