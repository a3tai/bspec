# WARP.md

This file provides guidance to WARP (warp.dev) when working with code in this repository.

## Project Overview

BSpec-Foundations is the home of the Business Specification Standard (BSpec) v1.0 - a universal language for describing any business as a structured, machine-readable knowledge graph. This repository contains the complete specification, reference implementations, and the foundation for a comprehensive ecosystem of tools and services.

**Core Innovation**: Every business is treated as a system of atomic, interconnected documents that together form a complete picture of strategy, operations, and execution.

## Essential Commands

### Development Setup
```bash
# Python-based development (primary)
python3 -m pip install -r sdk/v1/python/requirements.txt

# Web development (SvelteKit + Tailwind CSS 4.0)
cd apps/web && bun install && bun run dev

# MCP server development (Cloudflare Workers)
cd apps/mcp && bun install && bun run dev

# CLI development (Go)
cd sdk/cli && make deps && make build

# TypeScript SDK
cd sdk/v1/typescript && bun install && bun run build
```

### Core Development Workflows

#### Generate All SDKs from Specification
```bash
# Generate all language SDKs (TypeScript, Python, Go, JSON)
python3 scripts/generate.py

# Generate specific SDK only
python3 scripts/generate.py --lang typescript

# Clean and regenerate all
python3 scripts/generate.py --clean

# Verify specification before generation
python3 scripts/generate.py --verify
```

#### Testing
```bash
# Test specification parser
python3 scripts/test_generator.py

# Test Go CLI
cd sdk/cli && make test

# Test web app
cd apps/web && bun test

# Test with coverage
cd sdk/cli && make test-coverage
```

#### Building and Packaging
```bash
# Build CLI for all platforms
cd sdk/cli && make build-all

# Build web application
cd apps/web && bun run build

# Deploy MCP server
cd apps/mcp && bun run deploy

# Create release package
python3 scripts/create_tgz.py
```

#### Linting and Formatting
```bash
# Go formatting and linting
cd sdk/cli && make fmt && make lint

# Web app linting
cd apps/web && bun run lint && bun run format

# TypeScript checks
cd apps/mcp && bun run type-check
```

## High-Level Architecture

### Repository Structure
The codebase follows a modular, domain-driven architecture:

```
BSpec-Foundations/
├── spec/v1/                    # Complete BSpec 1.0 specification
│   ├── spec.md                 # Master specification document
│   ├── strategic-foundation/   # Strategic document types
│   ├── market-environment/     # Market analysis documents
│   ├── customer-value/         # Customer-focused documents
│   ├── product-service/        # Product definition documents
│   ├── business-model/         # Business model documents
│   ├── operations-execution/   # Operations documents
│   ├── technology-data/        # Technology documents
│   ├── financial-investment/   # Financial documents
│   ├── risk-governance/        # Risk management documents
│   ├── growth-innovation/      # Growth strategy documents
│   └── learning-decisions/     # Learning and decision documents
├── scripts/                    # Code generation and automation
│   ├── parsers/               # Specification parsers
│   ├── generators/            # SDK generators for all languages
│   └── generate.py            # Master orchestrator
├── sdk/v1/                    # Generated SDKs
│   ├── typescript/            # TypeScript SDK with full types
│   ├── python/                # Python SDK with Pydantic models
│   ├── go/                    # Go SDK with struct definitions
│   └── json/                  # JSON schemas for validation
├── apps/                      # BSpec ecosystem applications
│   ├── web/                   # bspec.dev website (SvelteKit)
│   ├── mcp/                   # MCP server (Cloudflare Workers)
│   └── docs/                  # Documentation site
└── sdk/cli/                   # Go-based CLI tool
```

### Core System Components

#### 1. Specification-First Architecture
- **Single Source of Truth**: All SDKs and tools are generated from `spec/v1/spec.md`
- **Document Types**: 82 standardized business document types across 11 domains
- **Relationship Model**: Documents declare dependencies, enablements, and conflicts
- **Machine + Human Readable**: YAML frontmatter + Markdown content

#### 2. Multi-Language SDK Generation
- **Code Generators**: Python-based generators create SDKs in TypeScript, Python, Go, and JSON
- **Type Safety**: Full type definitions in all target languages
- **Validation**: Schema-based validation for all document types
- **Consistent APIs**: Same interface patterns across all languages

#### 3. AI-First Integration
- **MCP Protocol**: Standard Model Context Protocol server for AI agent integration
- **JSON Conversion**: All documents can be converted to structured JSON for AI consumption
- **Prompt Templates**: Standard prompts for document generation and business analysis
- **Knowledge Graphs**: Rich relationship modeling for intelligent analysis

#### 4. Web Ecosystem
- **SvelteKit Application**: Modern web framework with TypeScript
- **Apple-Inspired Design**: Premium UI with Liquid Glass effects using Tailwind CSS 4.0
- **Interactive Features**: Document type explorer, conformance calculator
- **Static Site Generation**: Optimized for performance and SEO

### Data Flow Architecture

1. **Specification**: Human-authored business document type definitions in `spec/v1/`
2. **Parsing**: Python parsers extract structured data from specification files
3. **Generation**: Language-specific generators create type-safe SDKs
4. **Validation**: Generated SDKs include comprehensive validation logic
5. **Applications**: Web apps, CLI tools, and MCP servers consume the SDKs
6. **AI Integration**: MCP server provides structured business data to AI agents

### Critical Design Patterns

#### 1. Document Structure Standard
Every BSpec document follows this exact pattern:
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

# Human-readable content in Markdown
```

#### 2. Atomic Document Principle
- Each document covers exactly one business concern
- Documents are reusable and combinable
- Clear boundaries prevent scope creep
- Enables granular analysis and updates

#### 3. Rich Relationship Modeling
- **Dependencies**: What this document requires to be meaningful
- **Enablements**: What this document makes possible  
- **Conflicts**: Mutually exclusive documents
- **Related**: Contextual connections

This creates a business knowledge graph that AI systems can traverse and reason about.

## Development Guidelines

### Package Management
- **Node.js Projects**: Use `bun` instead of npm or node (as specified in CLAUDE.md)
- **Python**: Use standard pip with requirements.txt
- **Go**: Use Go modules with vendor directories

### Code Generation Workflow
1. Modify specification files in `spec/v1/` if needed
2. Run `python3 scripts/generate.py --verify` to validate changes
3. Generate updated SDKs with `python3 scripts/generate.py`
4. Test generated code in target languages
5. Update applications that consume the SDKs

### Testing Strategy
- **Specification**: Verify all document types parse correctly
- **SDKs**: Type checking, validation logic, and API consistency
- **Applications**: Unit tests for web components and CLI commands
- **Integration**: End-to-end workflows including generation and deployment

### AI Integration Patterns
The MCP server at `apps/mcp/` provides these tools for AI agents:
- `parse_bspec_document`: Parse and validate BSpec documents
- `analyze_business_spec`: Analyze conformance and identify gaps
- `convert_to_json`: Convert documents to JSON for AI consumption
- `generate_recommendations`: Generate improvement suggestions

### Quality Standards
- **Type Safety**: Full TypeScript/type annotations where applicable
- **Validation**: Comprehensive input validation and error handling  
- **Documentation**: Inline docs and comprehensive README files
- **Performance**: Efficient parsing and generation algorithms
- **Consistency**: Aligned patterns across all languages and tools

## Important Notes

### Technology Stack
- **Primary Language**: Python (for generators and core tooling)
- **Web Frontend**: SvelteKit with TypeScript and Tailwind CSS 4.0
- **CLI Tool**: Go with Cobra framework
- **MCP Server**: TypeScript on Cloudflare Workers
- **Package Management**: Bun for Node.js projects (critical requirement)

### Deployment Architecture
- **Website**: Static site deployment with edge functions at bspec.dev
- **MCP Server**: Containerized deployment at mcp.bspec.dev/mcp  
- **CLI**: Cross-platform binaries distributed via GitHub releases
- **SDKs**: Published to npm, PyPI, and Go modules registry

### Business Context
This is a standardization project similar to how SAFE agreements standardized early-stage investing. The goal is to create a universal language for describing any business as a structured knowledge graph that both humans and AI systems can understand and work with effectively.

The specification defines 82 document types across 11 comprehensive business domains, with conformance levels from Bronze (12+ docs) to Gold (45+ docs) and industry-specific profiles for Software/SaaS, Physical Products, Service Businesses, and Nonprofits.