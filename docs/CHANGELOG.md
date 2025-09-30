# BSpec Changelog

All notable changes to the BSpec project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- Semantic versioning specification and automation
- Version synchronization script (`scripts/version-sync.py`)
- Comprehensive documentation professionalization across all 82 document types

### Changed
- All specification documents now meet MBA/IETF-level professional standards
- Enhanced Purpose and Scope sections with normative language
- Standardized metadata schemas across all document types

### Fixed
- Version constant added to Go SDK for proper version tracking

## [1.0.0] - 2025-01-30

### Added
- Complete BSpec 1.0 Universal Business Specification Standard
- 82 document types across 11 business domains
- Multi-language SDK support (TypeScript, Python, Go, Rust, JSON Schema)
- BSpec CLI tool for document management and validation
- Web application for specification browsing (bspec.dev)
- MCP server for AI agent integration
- Comprehensive specification cleanup and professionalization
- Automated code generation from specification files
- Professional Abstract sections for all specifications
- Standardized validation checklists for all document types

### Core Domains
- **Strategic Foundation**: Mission, Vision, Strategy, Values, Objectives, Theory, Motivation, Purpose
- **Market Environment**: Market Analysis, Competitive Analysis, Positioning, Opportunities, Threats, Trends, Macroeconomic Analysis, Ecosystem Analysis, Regulatory Analysis, Segmentation Analysis
- **Customer Value**: Personas, Customer Journey Maps, Use Cases, User Stories, Jobs-to-be-Done, Pain Points, Gains Analysis, Behavior Analysis, Empathy Maps, Surveys, Customer Insights, Feedback Analysis
- **Product Service**: Product Definition, Features, Requirements, Service Definition, User Experience Design, Quality Standards, Integration Specifications, Support Definition, Product Strategy, Roadmap
- **Business Model**: Business Model Canvas, Value Streams, Key Partners, Key Resources, Key Activities, Cost Structure, Revenue Model, Customer Segments, Channels
- **Operations Execution**: Processes, Workflows, Organizational Structure, Roles, Teams, Skills, Capabilities, Facilities, Vendors, Policies, Service Level Agreements
- **Technology Data**: Architecture, Systems, Infrastructure, Development, Data Management, APIs, Security, Analytics
- **Financial Investment**: Financial Model, Budgets, Forecasts, Funding, Investment, Metrics, Reports, Tax, Audit, Valuation
- **Risk Governance**: Risk Analysis, Governance, Compliance, Legal, Ethics, Insurance, Controls, Incident Management
- **Growth Innovation**: Innovation Strategy, Research & Development, Experiments, Future Planning, Learning, Adoption, Ignition
- **Learning Decisions**: Hypothesis, Learning Plans, Decision Making, Knowledge Management, Retrospectives, Wisdom
- **Brand Marketing**: Brand Strategy, Messaging, Campaigns, Content, Social Media, SEO, Performance Marketing, Influencer Marketing, Video Marketing, Channel Marketing, Brand Positioning, Tone of Voice

### SDKs and Tools
- **TypeScript SDK**: Full type definitions and interfaces for all document types
- **Python SDK**: Pydantic models and validation for all document types  
- **Go SDK**: Struct definitions and methods for all document types
- **Rust SDK**: Type-safe structs and implementations for all document types
- **JSON Schema**: Complete JSON Schema definitions for validation
- **CLI Tool**: Command-line interface for document management, validation, and AI chat

### Applications
- **Web Application**: Interactive specification browser and conformance calculator
- **MCP Server**: Model Context Protocol server for AI agent integration
- **Documentation Site**: Comprehensive documentation and examples

### Technical Features
- Specification-first architecture with automated SDK generation
- Machine-readable YAML frontmatter for all document types
- Rich relationship modeling (dependencies, enablements, conflicts)
- Conformance levels (Bronze, Silver, Gold) for organizational maturity
- Industry-specific profiles (Software/SaaS, Physical Products, Services, Nonprofits)
- Comprehensive validation and quality assurance frameworks

---

## Version History Format

Each version entry should include:

### [Version] - YYYY-MM-DD

#### Added
- New features, document types, or capabilities

#### Changed  
- Changes to existing features or breaking changes

#### Deprecated
- Features marked for removal (with timeline)

#### Removed
- Features removed from the specification or SDKs

#### Fixed
- Bug fixes and corrections

#### Security
- Security-related changes or fixes

---

## Component Versioning

This changelog tracks changes across all BSpec components:

- **Core Specification** (`spec/v1/`): The master specification files
- **TypeScript SDK** (`sdk/v1/typescript/`): TypeScript/JavaScript interfaces and types
- **Python SDK** (`sdk/v1/python/`): Python models and validation
- **Go SDK** (`sdk/v1/go/`): Go structs and methods
- **Rust SDK** (`sdk/v1/rust/`): Rust types and implementations
- **JSON Schema** (`sdk/v1/json/`): JSON Schema definitions
- **CLI Tool** (`sdk/cli/`): Command-line interface tool
- **Web Application** (`apps/web/`): bspec.dev website
- **MCP Server** (`apps/mcp/`): AI integration server
- **Documentation Site** (`apps/docs/`): Documentation website

---

[Unreleased]: https://github.com/a3tai/bspec-foundations/compare/v1.0.0...HEAD
[1.0.0]: https://github.com/a3tai/bspec-foundations/releases/tag/v1.0.0