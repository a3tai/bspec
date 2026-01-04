---
title: "TypeScript SDK"
description: "Type-safe TypeScript/JavaScript SDK for working with BSpec documents"
---

# TypeScript SDK

The BSpec TypeScript SDK provides a fully type-safe interface for creating, validating, and managing Business Specification documents in Node.js and browser environments.

## Installation

```bash
npm install @bspec/sdk

# or with yarn
yarn add @bspec/sdk

# or with pnpm
pnpm add @bspec/sdk

# or with bun
bun add @bspec/sdk
```

## Quick Start

```typescript
import { BSpec, DocumentTypes } from '@bspec/sdk'

// Create a new Vision document
const vision: DocumentTypes.VSN = {
  id: 'VSN-product-2025',
  type: 'VSN',
  title: 'Product Vision 2025',
  status: 'Draft',
  version: '1.0.0',
  owner: 'Product Team',
  domain: 'strategic-foundation',
  
  // Vision-specific fields
  vision_statement: 'Transform how businesses document their operations',
  time_horizon: '3-5 years',
  target_state: {
    market_position: 'Leading business documentation standard',
    key_capabilities: ['Universal language', 'AI-ready', 'Machine-readable']
  }
}

// Validate the document
const result = BSpec.validate(vision)
if (!result.valid) {
  console.error('Validation errors:', result.errors)
}

// Parse from markdown
const doc = await BSpec.parseMarkdown(markdownContent)

// Convert to JSON
const json = BSpec.toJSON(vision)

// Generate markdown
const markdown = BSpec.toMarkdown(vision)
```

## Core Features

### Type Safety

Full TypeScript type definitions for all 112 document types:

```typescript
import type { 
  VSN, // Vision
  STR, // Strategy  
  MSN, // Mission
  BRD, // Business Requirements
  PRD  // Product Requirements
} from '@bspec/sdk'

// Autocomplete and type checking
const strategy: STR = {
  // TypeScript will validate all required fields
}
```

### Validation

```typescript
import { validate, ValidationResult } from '@bspec/sdk'

const result: ValidationResult = validate(document)

if (result.valid) {
  console.log('Document is valid!')
} else {
  result.errors.forEach(error => {
    console.error(`${error.path}: ${error.message}`)
  })
}

// Strict validation (warnings as errors)
const strictResult = validate(document, { strict: true })
```

### Document Parsing

```typescript
import { parseMarkdown, parseJSON } from '@bspec/sdk'

// Parse from markdown
const doc = await parseMarkdown(markdownString)

// Parse from JSON
const doc = parseJSON(jsonString)

// Parse with validation
const doc = await parseMarkdown(markdownString, { validate: true })
```

### Document Generation

```typescript
import { generate, DocumentType } from '@bspec/sdk'

// Generate from template
const template = generate('VSN', {
  title: 'My Vision',
  owner: 'Product Team'
})

// Convert to markdown
const markdown = template.toMarkdown()

// Save to file (Node.js)
await template.save('docs/vision.md')
```

### Relationships

```typescript
import { analyzeRelationships, DependencyGraph } from '@bspec/sdk'

// Analyze document relationships
const graph: DependencyGraph = analyzeRelationships(documents)

// Find dependencies
const deps = graph.getDependencies('VSN-product-2025')

// Find what a document enables
const enabled = graph.getEnabled('STR-growth-2025')

// Check for conflicts
const conflicts = graph.getConflicts('BRD-feature-a')

// Topological sort for build order
const buildOrder = graph.topologicalSort()
```

### Conformance Analysis

```typescript
import { analyzeConformance, ConformanceLevel } from '@bspec/sdk'

const analysis = analyzeConformance(documents)

console.log(`Conformance Level: ${analysis.level}`) // Bronze, Silver, or Gold
console.log(`Documents: ${analysis.documentCount}/82`)
console.log(`Coverage: ${analysis.coverage}%`)

// Get missing documents for next level
const missing = analysis.getMissingForLevel('Silver')
missing.forEach(type => {
  console.log(`Missing: ${type.code} - ${type.name}`)
})
```

## API Reference

### Document Types

All 82 BSpec document types are available:

```typescript
import {
  // Strategic Foundation
  VSN, STR, MSN, VAL, OBJ, PUR, THY, MOT,
  
  // Market Environment
  MKT, SEG, CMP, ECO, TRN, REG, OPP, POS, THR, MAC,
  
  // Customer Value  
  CUS, PER, USE, STO, PAI, FEE, CJM, SUR, JTB, CIN,
  EMP, GAI, BEH,
  
  // And 69 more...
} from '@bspec/sdk'
```

### Core Functions

```typescript
// Validation
validate(document: Document, options?: ValidationOptions): ValidationResult
validateAsync(document: Document): Promise<ValidationResult>

// Parsing
parseMarkdown(content: string, options?: ParseOptions): Promise<Document>
parseJSON(content: string): Document
parseYAML(content: string): Document

// Generation
generate(type: DocumentType, metadata: Partial<Document>): Document
toMarkdown(document: Document): string
toJSON(document: Document): string
toYAML(document: Document): string

// Analysis
analyzeRelationships(documents: Document[]): DependencyGraph
analyzeConformance(documents: Document[]): ConformanceAnalysis
findBrokenReferences(documents: Document[]): BrokenReference[]
```

## Framework Integration

### React

```typescript
import { BSpecProvider, useDocument, useValidation } from '@bspec/react'

function App() {
  return (
    <BSpecProvider documents={documents}>
      <DocumentEditor />
    </BSpecProvider>
  )
}

function DocumentEditor() {
  const { document, update } = useDocument('VSN-product-2025')
  const { isValid, errors } = useValidation(document)
  
  return (
    <div>
      <input 
        value={document.title}
        onChange={e => update({ title: e.target.value })}
      />
      {!isValid && <Errors errors={errors} />}
    </div>
  )
}
```

### Next.js

```typescript
// app/api/validate/route.ts
import { validate } from '@bspec/sdk'
import { NextResponse } from 'next/server'

export async function POST(request: Request) {
  const document = await request.json()
  const result = validate(document)
  
  return NextResponse.json(result)
}
```

### Node.js / Express

```typescript
import express from 'express'
import { validate, parseMarkdown } from '@bspec/sdk'

const app = express()

app.post('/validate', async (req, res) => {
  const document = req.body
  const result = validate(document)
  res.json(result)
})

app.post('/parse', async (req, res) => {
  const { markdown } = req.body
  const document = await parseMarkdown(markdown)
  res.json(document)
})
```

## Examples

### Complete Workflow

```typescript
import { 
  generate, 
  validate, 
  toMarkdown, 
  analyzeConformance 
} from '@bspec/sdk'
import { writeFile } from 'fs/promises'

async function createBusinessSpec() {
  // 1. Generate core documents
  const vision = generate('VSN', {
    title: 'Company Vision 2025',
    owner: 'CEO'
  })
  
  const strategy = generate('STR', {
    title: 'Growth Strategy',
    owner: 'Strategy Team',
    depends_on: [vision.id]
  })
  
  // 2. Validate documents
  const documents = [vision, strategy]
  for (const doc of documents) {
    const result = validate(doc)
    if (!result.valid) {
      throw new Error(`Invalid document: ${doc.id}`)
    }
  }
  
  // 3. Analyze conformance
  const analysis = analyzeConformance(documents)
  console.log(`Conformance: ${analysis.level}`)
  console.log(`Coverage: ${analysis.coverage}%`)
  
  // 4. Save as markdown
  for (const doc of documents) {
    const markdown = toMarkdown(doc)
    await writeFile(`docs/${doc.id}.md`, markdown)
  }
}
```

## Best Practices

1. **Use TypeScript**: Get full type safety and autocomplete
2. **Validate Early**: Validate documents before saving or committing
3. **Type Guards**: Use type guards to narrow document types
4. **Async Operations**: Use async parsing for better performance
5. **Error Handling**: Always check validation results

## Support

- **Documentation**: [https://bspec.dev/docs](https://bspec.dev/docs)
- **NPM Package**: [@bspec/sdk](https://www.npmjs.com/package/@bspec/sdk)
- **GitHub**: [a3tai/bspec](https://github.com/a3tai/bspec)
- **Issues**: [Report bugs](https://github.com/a3tai/bspec/issues)

## Next Steps

- Explore the [CLI Tool](/docs/tools/cli)
- Learn about [Python SDK](/docs/tools/python)
- Read the [Full Specification](/spec/v1-0-0)
