---
layout: home

hero:
  name: BSpec
  text: Business Specification Standard
  tagline: A universal language for describing any business as a structured, machine-readable knowledge graph
  actions:
    - theme: brand
      text: Get Started
      link: /docs/overview
    - theme: alt
      text: View Specification
      link: /spec/v1-0-0
    - theme: alt
      text: GitHub
      link: https://github.com/yourusername/bspec

features:
  - icon: 📄
    title: 82 Document Types
    details: Comprehensive coverage across 11 business domains, from strategy to execution
  - icon: 🔗
    title: Knowledge Graph
    details: Rich relationships between documents create an interconnected business model
  - icon: 🤖
    title: AI-Ready
    details: MCP server integration enables AI agents to understand and work with your business
  - icon: 🛠️
    title: Multi-Language SDKs
    details: TypeScript, Python, Go, and JSON Schema support for any tech stack
  - icon: 📊
    title: Conformance Levels
    details: Bronze, Silver, and Gold levels help you grow your documentation over time
  - icon: 🏢
    title: Industry Profiles
    details: Specialized profiles for Software/SaaS, Physical Products, Services, and Nonprofits
---

## What is BSpec?

BSpec is a standardization project that creates a universal language for describing any business as a structured knowledge graph. Similar to how SAFE agreements standardized early-stage investing, BSpec provides a common framework that both humans and AI systems can understand.

### Core Innovation

Every business is treated as a system of **atomic, interconnected documents** that together form a complete picture of strategy, operations, and execution.

```yaml
---
# Example BSpec Document
id: MSN-growth-strategy-2025
type: MSN
status: Accepted
version: 1.0.0
depends_on: [VSN-company-vision, STR-market-strategy]
enables: [OBJ-revenue-targets, INI-expansion-initiatives]
domain: strategic-foundation
priority: critical
---

# 2025 Growth Strategy
...your content here...
```

### Why BSpec?

- **Standardization**: Common language across all businesses
- **Machine-Readable**: Enable AI agents to understand your business
- **Flexible**: Start small with Bronze level (12+ docs) and grow to Gold (45+ docs)
- **Open Source**: Free to use, modify, and extend

### Quick Start

```bash
# Install the CLI
curl -sSL https://bspec.dev/install.sh | bash

# Create a new BSpec project
bspec init my-business

# Validate your documents
bspec validate

# Generate knowledge graph
bspec graph
```

### Integration with AI

BSpec includes an MCP (Model Context Protocol) server that enables AI agents like Claude to:
- Parse and validate BSpec documents
- Analyze business conformance and identify gaps
- Generate recommendations
- Create document templates

[Learn more about MCP integration →](/docs/tools/mcp)
