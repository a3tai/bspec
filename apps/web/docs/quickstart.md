# Quick Start

Get started with BSpec in minutes. This guide will walk you through creating your first BSpec document.

## Installation

### Using the CLI (Recommended)

```bash
# Install via curl
curl -sSL https://bspec.dev/install.sh | bash

# Or download from GitHub releases
# Visit: https://github.com/a3tai/bspec/releases
```

### Using SDKs

::: code-group

```bash [TypeScript]
bun add @bspec/typescript
```

```bash [Python]
pip install bspec-python
```

```bash [Go]
go get github.com/a3tai/bspec-go
```

:::

## Your First Document

Let's create a simple Vision Statement (VSN) document.

### 1. Create the Document

Create a file named `VSN-company-vision.md`:

```yaml
---
# CORE IDENTITY
id: VSN-company-vision
type: VSN
status: Draft
version: 1.0.0

# RELATIONSHIPS
depends_on: []
enables: [MSN-mission, STR-strategy]
conflicts_with: []

# BUSINESS CONTEXT
domain: strategic-foundation
priority: critical
success_criteria:
  - "Vision is clear and inspiring"
  - "Team alignment on long-term direction"
  - "Vision guides strategic decisions"

# METADATA
created: 2025-01-15
updated: 2025-01-15
owner: leadership-team
stakeholders:
  - board
  - executive-team
  - all-employees
tags:
  - strategy
  - foundation
  - culture
---

# Company Vision

## Our Vision

[Describe your company's long-term vision - where you want to be in 5-10 years]

## Why This Vision Matters

[Explain why this vision is important and how it guides your company]

## How We'll Achieve It

[Outline the key pillars or strategies that will help realize this vision]
```

### 2. Validate the Document

```bash
bspec validate VSN-company-vision.md
```

You should see:
```
✓ VSN-company-vision.md is valid
- Type: VSN (Vision Statement)
- Status: Draft
- Domain: strategic-foundation
- Priority: critical
```

### 3. View Document Relationships

```bash
bspec graph VSN-company-vision.md
```

This will show you which documents this vision enables and what it depends on.

## Creating a Complete Set

A Bronze-level conformance requires 12+ core documents. Let's create the essential foundation:

### 1. Strategic Foundation (3 documents)
- **VSN** - Vision Statement (what you just created!)
- **MSN** - Mission Statement
- **STR** - Strategy Document

### 2. Market Environment (2 documents)
- **CMP** - Competitive Analysis
- **SEG** - Customer Segments

### 3. Customer Value (1 document)
- **VLU** - Value Proposition

### 4. Product/Service (2 documents)
- **PRD** - Product Requirements
- **FEA** - Feature Specifications

### 5. Operations (2 documents)
- **OPS** - Operations Plan
- **TEA** - Team Structure

### 6. Financial (2 documents)
- **BUD** - Budget Document
- **FOR** - Financial Forecast

## Using with AI (MCP)

If you're using Claude or another AI assistant with MCP support:

1. Enable the BSpec MCP server
2. Ask Claude to analyze your documents:

```
Can you analyze my BSpec documents and identify any gaps 
in my business documentation?
```

Claude will be able to:
- Parse your BSpec documents
- Identify missing document types
- Suggest improvements
- Generate new document templates

[Learn more about MCP integration →](/docs/tools/mcp)

## Next Steps

- [Core Concepts](/docs/concepts) - Understand BSpec fundamentals
- [Document Types](/docs/document-types) - Explore all 82 document types
- [CLI Reference](/docs/tools/cli) - Learn all CLI commands
- [Conformance Levels](/spec/conformance) - Understand Bronze, Silver, and Gold levels
