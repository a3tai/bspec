# BSpec TypeScript SDK

TypeScript interfaces and types for the BSpec v1.1.2 Universal Business Specification Standard.

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

This SDK includes TypeScript interfaces for all 127 BSpec document types across 15 business domains.

## Generated

- **From**: BSpec v1.1.2 specification
- **At**: 2026-07-12T17:47:25.168020
- **Generator**: typescript-generator-v1.0.0

## License

MIT License - see the repository LICENSE file for details.
