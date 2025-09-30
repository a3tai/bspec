# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

BSpec-Foundations is the home of the Business Specification Standard (BSpec) v1.0 - a universal language for describing any business as a structured, machine-readable knowledge graph. This repository contains the complete specification, reference implementations, and the foundation for a comprehensive ecosystem of tools and services hosted at bspec.dev.

## Repository Structure

```
BSpec-Foundations/
├── spec/v1/                     # BSpec 1.0 Specification (COMPLETE)
│   └── spec.md                  # Complete 
└── README.md                    # Project overview and getting started
```

## BSpec 1.0 Core Concepts

### Universal Business Vocabulary (82 Document Types)

BSpec defines a complete vocabulary across 11 domains:

1. **Strategic Foundation (8)**: MSN, VSN, VAL, STR, OBJ, MOT, PUR, THY
2. **Market & Environment (10)**: MKT, SEG, CMP, POS, TRN, ECO, OPP, THR, REG, MAC
3. **Customer & Value (12)**: PER, JTB, CJM, USE, STO, PAI, GAI, EMP, FEE, INT, SUR, BEH
4. **Product & Service (10)**: PRD, SVC, FEA, ROD, REQ, QUA, UXD, PER, INT, SUP
5. **Business Model (12)**: BMC, REV, PRC, CST, CHN, REL, RES, ACT, PRT, UNT, LTV, CAC
6. **Operations & Execution (12)**: PRC, WFL, ORG, ROL, TEA, SKI, POL, SLA, VND, FAC, TOO, CAP
7. **Technology & Data (8)**: ARC, SYS, DAT, API, INF, SEC, DEV, ANA
8. **Financial & Investment (10)**: FIN, BUD, FOR, FND, INV, VAL, MET, REP, AUD, TAX
9. **Risk & Governance (8)**: RSK, MIT, CMP, GVN, CTL, CRI, ETH, STA
10. **Growth & Innovation (8)**: GTM, GRW, SCL, EXP, INN, RND, ACQ, EXP
11. **Learning & Decisions (6)**: DEC, LRN, RET, HYP, KNO, WIS

### Document Structure Standard

Every BSpec document follows this exact YAML frontmatter + Markdown structure:

```yaml
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
```

### Conformance Levels

- **Bronze (12+ docs)**: Minimum viable business spec
- **Silver (25+ docs)**: Investment ready
- **Gold (45+ docs)**: Operational excellence

### Industry Profiles

Specialized requirements for:
- **Software/SaaS**: ARC, API, SEC, SUP + SaaS metrics
- **Physical Product**: INF, QUA, VND, REG + supply chain
- **Service Business**: PRC, SLA, SKI, QUA + delivery processes
- **Nonprofit**: PUR, STA, MET, GVN + impact measurement

## Planned Ecosystem Architecture (bspec.dev)

### Domain Structure
- **bspec.dev**: Main marketing website and community hub
- **bspec.dev/docs**: Technical documentation and specification
- **mcp.bspec.dev/mcp**: Hosted MCP server for AI agent integration

### Technology Stack

#### bspec.dev Website (apps/web/)
- **Framework**: SvelteKit with TypeScript
- **Design**: Apple-inspired aesthetic with Liquid Glass styling
- **Components**: Interactive document type explorer, conformance calculator
- **Deployment**: Static site with edge functions
- **Features**:
  - Getting started guides and tutorials
  - Live BSpec document validator
  - Community showcase and case studies
  - SDK and tool downloads

#### SDKs (sdk/v1/)

**TypeScript SDK**: Core parsing, validation, and JSON conversion
```typescript
import { BSpecParser, BSpecValidator, BSpecConverter } from '@bspec/typescript-sdk';

// Parse BSpec document
const document = BSpecParser.parse(markdownContent);

// Validate against schema
const validation = BSpecValidator.validate(document);

// Convert to JSON for AI consumption
const json = BSpecConverter.toJSON(document);

// Analyze business relationships
const dependencies = BSpecAnalyzer.getDependencies(document.id);
```

**Python SDK**: Full business analysis and AI integration
```python
from bspec import BSpecDocument, BSpecAnalyzer, BSpecGenerator

# Load and analyze business spec
spec = BSpecDocument.load_directory('./business-spec/')
analysis = BSpecAnalyzer.analyze_conformance(spec, target='silver')

# Generate artifacts
pitch_deck = BSpecGenerator.generate_pitch_deck(spec, audience='investor')
```

**Go SDK**: High-performance processing for servers
```go
package main

import "github.com/bspec/go-sdk/pkg/parser"

func main() {
    documents, err := parser.ParseDirectory("./business-spec/")
    analysis := analyzer.AnalyzeConformance(documents, "gold")
}
```

#### CLI Tool (sdk/cli/)
```bash
# Validate BSpec documents
bspec validate ./business-spec/

# Analyze conformance level
bspec analyze --target silver ./business-spec/

# Convert to JSON for AI consumption
bspec convert --format json ./business-spec/ > business.json

# Generate artifacts
bspec generate pitch-deck --audience investor ./business-spec/

# Package into .bspec format
bspec package ./business-spec/ business.bspec
```

#### MCP Server (sdk/mcp/)
Hosted at `mcp.bspec.dev/mcp` for AI agent integration:

**Capabilities**:
- Parse and validate BSpec documents
- Convert between Markdown and JSON formats
- Query business knowledge graphs
- Generate documents and artifacts
- Analyze conformance and readiness

**MCP Tools**:
```json
{
  "tools": [
    {
      "name": "parse_bspec_document",
      "description": "Parse and validate a BSpec document",
      "inputSchema": {
        "type": "object",
        "properties": {
          "content": {"type": "string"},
          "validate": {"type": "boolean", "default": true}
        }
      }
    },
    {
      "name": "analyze_business_spec",
      "description": "Analyze conformance and gaps in business specification",
      "inputSchema": {
        "type": "object",
        "properties": {
          "documents": {"type": "array"},
          "target_level": {"type": "string", "enum": ["bronze", "silver", "gold"]},
          "industry_profile": {"type": "string"}
        }
      }
    },
    {
      "name": "convert_to_json",
      "description": "Convert BSpec documents to JSON for AI consumption",
      "inputSchema": {
        "type": "object",
        "properties": {
          "documents": {"type": "array"},
          "include_relationships": {"type": "boolean", "default": true}
        }
      }
    }
  ]
}
```

### .bspec File Format

Packaging format for complete business specifications:

```
business-spec.bspec (compressed archive like .tgz)
├── manifest.json           # Package metadata and validation
├── documents/              # All BSpec documents
│   ├── 01-strategic/       # Organized by domain
│   ├── 02-market/
│   ├── 03-customer/
│   └── ...
├── assets/                 # Images, diagrams, attachments
│   ├── diagrams/           # SVG diagrams
│   ├── images/             # Photos and graphics
│   └── documents/          # PDFs and other files
├── relationships/          # Computed relationship graph
│   ├── dependency-graph.json
│   ├── impact-analysis.json
│   └── conformance-report.json
└── schemas/                # Validation schemas used
    ├── document-schemas.json
    └── relationship-schemas.json
```

## Development Workflow

### For BSpec Specification Work
1. **Read Full Spec**: Start with `spec/v1/spec.md`
2. **Follow Document Structure**: Use exact YAML frontmatter format
3. **Maintain Relationships**: Ensure `depends_on`, `enables`, `conflicts_with` are accurate
4. **Validate Conformance**: Check against Bronze/Silver/Gold requirements

### For SDK Development
1. **Follow Language Conventions**: Each SDK should feel idiomatic to its language
2. **Implement Core Features**: Parser, Validator, Converter, Analyzer, Generator
3. **Include Full Schema Support**: All document types with validation
4. **Enable AI Integration**: JSON conversion and MCP compatibility
5. **Provide Rich APIs**: Easy-to-use interfaces for common operations

### For Website Development (apps/web/)
1. **Apple-Inspired Design**: Clean, minimal, premium aesthetic
2. **Liquid Glass Styling**: Frosted glass effects, subtle animations
3. **Interactive Features**: Document type explorer, conformance calculator
4. **Performance First**: Fast loading, progressive enhancement
5. **Mobile Responsive**: Excellent experience on all devices

### For MCP Server Development (sdk/mcp/)
1. **Standard Compliance**: Full MCP protocol implementation
2. **Robust Error Handling**: Clear error messages and validation feedback
3. **Performance Optimization**: Efficient parsing and analysis
4. **Security Focused**: Input validation and safe processing
5. **Comprehensive Logging**: Detailed logs for debugging and monitoring

## File Naming Conventions

### BSpec Documents
`{TYPE}-{kebab-case-name}-v{version}.md`

Examples:
- `MSN-company-mission-v1.0.0.md`
- `CAP-model-serving-multi-region-v1.0.0.md`
- `CTX-orchestration-serving-v1.0.0.md`

### Code Files
- **TypeScript**: PascalCase for classes, camelCase for functions
- **Python**: snake_case for all identifiers
- **Go**: PascalCase for exported, camelCase for unexported

## Machine Readability & AI Integration

### JSON Schema Support
Every document type has corresponding JSON schema for validation and AI consumption.

### AI Prompt Templates
Standard prompts for:
- Document generation from business context
- Conformance analysis and gap identification
- Artifact generation (pitch decks, business plans, code)
- Relationship analysis and impact assessment

### MCP Integration Patterns
```typescript
// AI agents can use BSpec through MCP
const analysis = await mcp.call('analyze_business_spec', {
  documents: businessDocuments,
  target_level: 'silver',
  industry_profile: 'saas'
});

const recommendations = await mcp.call('generate_recommendations', {
  gaps: analysis.gaps,
  priority: 'critical'
});
```

## Quality Standards

### Documentation Quality
- **Clarity**: Understandable by target audience without training
- **Completeness**: All required sections meaningfully filled
- **Consistency**: Aligned assumptions and terminology
- **Currency**: Current and accurate information
- **Connectivity**: Accurate and valuable relationships

### Code Quality
- **Type Safety**: Full TypeScript/type hints where applicable
- **Error Handling**: Comprehensive error handling and validation
- **Testing**: Unit tests for all core functionality
- **Documentation**: Inline docs and README files
- **Performance**: Efficient algorithms and memory usage

## Deployment Architecture

### Development
- **Local Development**: All SDKs can be developed and tested locally
- **Testing**: Comprehensive test suites for all components
- **Validation**: Automated validation of BSpec documents

### Production (bspec.dev)
- **Website**: Static site deployment with edge functions
- **MCP Server**: Containerized deployment at mcp.bspec.dev/mcp
- **Documentation**: Generated documentation site at bspec.dev/docs
- **SDKs**: Published to respective package managers (npm, PyPI, Go modules)

## Contributing Guidelines

### For Specification Changes
1. **RFC Process**: Submit RFC for specification changes
2. **Community Review**: 30-day review period
3. **Backward Compatibility**: Maintain compatibility within major versions
4. **Migration Paths**: Provide clear upgrade paths

### For Tool Development
1. **Standard Compliance**: Follow BSpec 1.0 specification exactly
2. **Cross-Platform**: Support major operating systems
3. **Documentation**: Comprehensive documentation and examples
4. **Testing**: Automated tests and validation
5. **Community Feedback**: Incorporate user feedback and requests

## Important Development Notes

- **Never create files unless explicitly needed** for the specific task
- **Prefer editing existing files** over creating new ones
- **Follow established patterns** in the codebase
- **Maintain consistency** with BSpec specification
- **Document all design decisions** in appropriate ADR format
- **Validate all BSpec documents** against schemas before committing
- **Test SDK functionality** thoroughly before releasing
- **Keep ecosystem coherent** - all tools should work together seamlessly
- IMPORANT: Always use bun instead of node or npm