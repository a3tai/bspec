# BSpec TypeScript SDK

TypeScript interfaces and types for the BSpec v1.0.0 Universal Business Specification Standard.

## Installation

```bash
npm install @bspec/typescript-sdk
```

## Usage

```typescript
import { BaseBSpecDocument, MSNDocument, DOCUMENT_TYPES } from '@bspec/typescript-sdk';

// Use type-safe document interfaces
const mission: MSNDocument = {
  id: 'MSN-company-mission',
  title: 'Company Mission Statement',
  type: DOCUMENT_TYPES.MSN,
  status: 'Approved',
  version: '1.0.0',
  owner: 'executive-team',
  domain: 'Strategic Foundation',
  // ... other fields
};
```

## Contents

- `base.ts` - Base interfaces and types
- `interfaces/` - Individual document type interfaces
- `constants.ts` - Type-safe constants and enums
- `bspec-data.json` - Complete specification data
- `version.txt` - BSpec version number

## Document Types

This SDK includes TypeScript interfaces for all 112 BSpec document types across 15 business domains.

## Generated

- **From**: BSpec v1.0.0 specification
- **At**: 2025-09-28T21:08:19.116300
- **Generator**: typescript-generator-v1.0.0

## License

CC BY 4.0 - See [BSpec Foundation](https://bspec.dev) for details.
