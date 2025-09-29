# BSpec Package Format (.bspec) Specification

**Version:** 1.0.0
**Status:** Draft
**Last Updated:** 2025-09-28
**Part of:** BSpec 1.0 Universal Business Specification Standard

## Overview

The .bspec file format provides a standardized way to package, distribute, and consume complete business specifications. Like how .zip archives contain multiple files in a single container, .bspec files contain all documents, assets, and metadata needed to represent a complete business specification.

## Design Principles

### 1. Self-Contained
Every .bspec file contains everything needed to understand and analyze the business specification without external dependencies.

### 2. Version Agnostic
The format supports multiple versions of the BSpec standard and provides clear migration paths between versions.

### 3. Asset-Rich
Business specifications often reference diagrams, images, presentations, and other assets. The .bspec format includes dedicated asset storage with organized folder structures.

### 4. Machine-Readable
AI systems and analysis tools can parse, validate, and consume .bspec files programmatically with complete context.

### 5. Human-Friendly
When extracted, the contents maintain a logical folder structure that humans can navigate and understand.

## File Structure

A .bspec file is a compressed archive (using gzip compression) with the following standardized internal structure:

```
business-specification.bspec
├── manifest.json                    # Package metadata and validation
├── documents/                       # All BSpec documents organized by domain
│   ├── 01-strategic/
│   │   ├── MSN-company-mission-v1.0.0.md
│   │   ├── VSN-company-vision-v1.0.0.md
│   │   ├── VAL-core-values-v1.0.0.md
│   │   └── STR-platform-strategy-v1.0.0.md
│   ├── 02-market/
│   │   ├── MKT-ai-platform-market-v1.0.0.md
│   │   ├── CMP-competitive-landscape-v1.0.0.md
│   │   └── POS-market-positioning-v1.0.0.md
│   ├── 03-customer/
│   │   ├── PER-enterprise-cto-v1.0.0.md
│   │   ├── PER-ml-platform-lead-v1.0.0.md
│   │   └── JTB-ai-platform-adoption-v1.0.0.md
│   ├── 04-product/
│   │   ├── PRD-inference-api-v1.0.0.md
│   │   └── SVC-model-hosting-v1.0.0.md
│   ├── 05-business-model/
│   │   ├── BMC-platform-canvas-v1.0.0.md
│   │   ├── REV-api-usage-revenue-v1.0.0.md
│   │   └── UNT-unit-economics-v1.0.0.md
│   ├── 06-operations/
│   │   └── PRC-customer-onboarding-v1.0.0.md
│   ├── 07-technology/
│   │   ├── ARC-system-architecture-v1.0.0.md
│   │   └── API-inference-endpoints-v1.0.0.md
│   ├── 08-financial/
│   │   ├── FIN-financial-model-v1.0.0.md
│   │   └── MET-key-metrics-v1.0.0.md
│   ├── 09-risk/
│   │   ├── RSK-platform-risks-v1.0.0.md
│   │   └── MIT-risk-mitigations-v1.0.0.md
│   ├── 10-growth/
│   │   └── GTM-developer-first-strategy-v1.0.0.md
│   └── 11-learning/
│       └── DEC-platform-technology-choice-v1.0.0.md
├── assets/                          # All referenced assets
│   ├── diagrams/                    # SVG diagrams and flowcharts
│   │   ├── system-architecture.svg
│   │   ├── user-journey-flow.svg
│   │   └── business-model-canvas.svg
│   ├── images/                      # Photos, screenshots, mockups
│   │   ├── product-screenshots/
│   │   ├── team-photos/
│   │   └── brand-assets/
│   ├── documents/                   # PDFs, presentations, external docs
│   │   ├── market-research-report.pdf
│   │   ├── investor-presentation.pdf
│   │   └── technical-specifications.pdf
│   └── data/                        # Spreadsheets, datasets, models
│       ├── financial-projections.xlsx
│       └── market-sizing-data.csv
├── computed/                        # Generated analysis and relationships
│   ├── relationships/
│   │   ├── dependency-graph.json
│   │   ├── impact-analysis.json
│   │   └── conflict-detection.json
│   ├── analysis/
│   │   ├── conformance-report.json
│   │   ├── gap-analysis.json
│   │   └── readiness-assessment.json
│   └── exports/
│       ├── complete-specification.json
│       └── artifact-templates/
├── schemas/                         # Validation schemas used
│   ├── bspec-version.json
│   ├── document-schemas.json
│   └── relationship-schemas.json
└── CHANGELOG.md                     # Version history and changes
```

## Manifest File Specification

The `manifest.json` file contains essential metadata about the .bspec package:

```json
{
  "format_version": "1.0.0",
  "bspec_version": "1.0.0",
  "package": {
    "id": "com.example.platform-business-spec",
    "name": "AI Platform Business Specification",
    "description": "Complete business specification for an AI inference platform",
    "version": "2.1.0",
    "created": "2025-09-28T10:00:00Z",
    "updated": "2025-09-28T15:30:00Z",
    "creator": {
      "name": "Platform Team",
      "email": "team@example.com",
      "organization": "Example AI Inc."
    }
  },
  "business": {
    "industry": ["software", "artificial-intelligence", "b2b"],
    "stage": "growth",
    "geography": ["us", "europe"],
    "revenue_model": ["subscription", "usage-based"],
    "target_conformance": "silver"
  },
  "contents": {
    "documents": {
      "count": 23,
      "by_domain": {
        "strategic": 4,
        "market": 3,
        "customer": 3,
        "product": 2,
        "business-model": 3,
        "operations": 1,
        "technology": 2,
        "financial": 2,
        "risk": 2,
        "growth": 1,
        "learning": 1
      },
      "by_status": {
        "accepted": 20,
        "review": 2,
        "draft": 1
      },
      "conformance_level": "silver",
      "missing_for_gold": ["ORG", "TEA", "CMP", "GVN", "INN"]
    },
    "assets": {
      "count": 15,
      "types": ["svg", "png", "pdf", "xlsx"],
      "total_size_mb": 12.5
    },
    "relationships": {
      "dependencies": 45,
      "enablements": 38,
      "conflicts": 2,
      "related": 67
    }
  },
  "validation": {
    "schema_compliance": true,
    "relationship_integrity": true,
    "asset_references": true,
    "circular_dependencies": false,
    "warnings": [],
    "errors": []
  },
  "tools": {
    "created_with": "bspec-cli v1.2.0",
    "compatible_with": ["@bspec/typescript-sdk >=1.0.0", "bspec-python >=1.0.0"],
    "recommended_tools": ["bspec-analyzer", "bspec-generator"]
  },
  "export_options": {
    "formats": ["json", "yaml", "html"],
    "artifacts": ["pitch-deck", "business-plan", "executive-summary"],
    "analysis": ["conformance", "readiness", "gap-analysis"]
  }
}
```

## Directory Organization Standards

### Document Organization (documents/)
Documents are organized by domain using numbered prefixes for consistent ordering:

```
01-strategic/     # Strategic Foundation (MSN, VSN, VAL, STR, OBJ, MOT, PUR, THY)
02-market/        # Market & Environment (MKT, SEG, CMP, POS, TRN, ECO, OPP, THR, REG, MAC)
03-customer/      # Customer & Value (PER, JTB, CJM, USE, STO, PAI, GAI, EMP, FEE, INT, SUR, BEH)
04-product/       # Product & Service (PRD, SVC, FEA, ROD, REQ, QUA, UXD, PER, INT, SUP)
05-business-model/ # Business Model (BMC, REV, PRC, CST, CHN, REL, RES, ACT, PRT, UNT, LTV, CAC)
06-operations/    # Operations & Execution (PRC, WFL, ORG, ROL, TEA, SKI, POL, SLA, VND, FAC, TOO, CAP)
07-technology/    # Technology & Data (ARC, SYS, DAT, API, INF, SEC, DEV, ANA)
08-financial/     # Financial & Investment (FIN, BUD, FOR, FND, INV, VAL, MET, REP, AUD, TAX)
09-risk/          # Risk & Governance (RSK, MIT, CMP, GVN, CTL, CRI, ETH, STA)
10-growth/        # Growth & Innovation (GTM, GRW, SCL, EXP, INN, RND, ACQ, EXP)
11-learning/      # Learning & Decisions (DEC, LRN, RET, HYP, KNO, WIS)
```

### Asset Organization (assets/)
Assets are organized by type and purpose:

- **diagrams/**: SVG files for business process flows, system architecture, user journeys
- **images/**: Photos, screenshots, mockups, brand assets
- **documents/**: PDFs, Word docs, presentations, external research
- **data/**: Spreadsheets, CSV files, datasets, financial models

### Computed Analysis (computed/)
Machine-generated analysis and derived data:

- **relationships/**: Dependency graphs, impact analysis, conflict detection
- **analysis/**: Conformance reports, gap analysis, readiness assessments
- **exports/**: JSON representations, artifact templates

## File Naming Conventions

### Document Files
Follow BSpec standard: `{TYPE}-{kebab-case-name}-v{version}.md`

Examples:
- `MSN-company-mission-v1.0.0.md`
- `PER-enterprise-decision-maker-v2.1.0.md`
- `ARC-microservices-architecture-v1.0.0.md`

### Asset Files
Use descriptive names with appropriate extensions:
- `system-architecture-overview.svg`
- `customer-journey-enterprise.svg`
- `q3-financial-projections.xlsx`
- `competitive-analysis-report.pdf`

## Asset Reference System

### Relative Path References
Documents reference assets using relative paths from the document location:

```markdown
## System Architecture

![System Architecture](../../assets/diagrams/system-architecture.svg)

For detailed technical specifications, see the [Technical Architecture Document](../../assets/documents/technical-specs.pdf).
```

### Asset Manifest
The `assets/manifest.json` file tracks all assets and their relationships:

```json
{
  "assets": [
    {
      "path": "diagrams/system-architecture.svg",
      "type": "diagram",
      "format": "svg",
      "size_bytes": 15420,
      "checksum": "sha256:a1b2c3d4...",
      "created": "2025-09-28T14:00:00Z",
      "referenced_by": [
        "documents/07-technology/ARC-system-architecture-v1.0.0.md",
        "documents/04-product/PRD-platform-api-v1.0.0.md"
      ],
      "description": "High-level system architecture showing microservices and data flow"
    }
  ]
}
```

## Computed Relationships

### Dependency Graph (computed/relationships/dependency-graph.json)
```json
{
  "graph_version": "1.0.0",
  "generated": "2025-09-28T15:30:00Z",
  "nodes": [
    {
      "id": "MSN-company-mission-v1.0.0",
      "type": "MSN",
      "title": "Company Mission",
      "status": "accepted",
      "domain": "strategic"
    }
  ],
  "edges": [
    {
      "from": "STR-platform-strategy-v1.0.0",
      "to": "MSN-company-mission-v1.0.0",
      "type": "depends_on",
      "strength": "critical"
    }
  ],
  "analysis": {
    "total_nodes": 23,
    "total_edges": 45,
    "cycles_detected": 0,
    "orphaned_documents": [],
    "critical_path": ["MSN-company-mission-v1.0.0", "STR-platform-strategy-v1.0.0", "PRD-platform-api-v1.0.0"]
  }
}
```

### Conformance Report (computed/analysis/conformance-report.json)
```json
{
  "report_version": "1.0.0",
  "generated": "2025-09-28T15:30:00Z",
  "target_level": "silver",
  "current_level": "silver",
  "industry_profile": "software-saas",
  "assessment": {
    "bronze": {
      "status": "achieved",
      "score": 1.0,
      "requirements_met": 12,
      "requirements_total": 12
    },
    "silver": {
      "status": "achieved",
      "score": 1.0,
      "requirements_met": 25,
      "requirements_total": 25
    },
    "gold": {
      "status": "in_progress",
      "score": 0.78,
      "requirements_met": 35,
      "requirements_total": 45,
      "missing": ["ORG", "TEA", "CMP", "GVN", "INN"]
    }
  },
  "domain_coverage": {
    "strategic": {"documents": 4, "coverage": "excellent"},
    "market": {"documents": 3, "coverage": "good"},
    "customer": {"documents": 3, "coverage": "good"},
    "operations": {"documents": 1, "coverage": "minimal"}
  },
  "recommendations": [
    {
      "priority": "high",
      "type": "missing_document",
      "suggestion": "Create ORG (Organization) document to define team structure",
      "impact": "Required for Gold conformance"
    }
  ]
}
```

## Package Creation Process

### 1. Validation Phase
Before packaging, all documents must pass validation:
- Schema compliance for each document type
- Relationship integrity (no circular dependencies)
- Asset reference verification (all referenced assets exist)
- Conformance level assessment

### 2. Asset Collection
- Copy all referenced assets to appropriate asset directories
- Generate asset manifest with checksums and metadata
- Validate asset formats and accessibility

### 3. Analysis Generation
- Compute dependency graphs and relationship analysis
- Generate conformance reports and gap analysis
- Create export formats (JSON, YAML) for programmatic consumption

### 4. Compression
- Create compressed archive using gzip compression
- Include all directories and files in standardized structure
- Generate final .bspec file with appropriate naming

## Package Consumption

### Extraction
```bash
# Extract .bspec package
tar -xzf business-specification.bspec -C ./extracted/

# Verify package integrity
bspec validate ./extracted/manifest.json
```

### Programmatic Access
```typescript
import { BSpecPackage } from '@bspec/typescript-sdk';

// Load package
const package = await BSpecPackage.load('business-specification.bspec');

// Access documents
const mission = package.getDocument('MSN-company-mission-v1.0.0');
const allStrategic = package.getDocumentsByDomain('strategic');

// Analyze relationships
const dependencies = package.getDependencies('STR-platform-strategy-v1.0.0');
const conformance = package.getConformanceReport();

// Access assets
const diagram = package.getAsset('diagrams/system-architecture.svg');
```

### Export Capabilities
```bash
# Export to JSON for AI consumption
bspec export --format json business-specification.bspec > business.json

# Generate artifacts
bspec generate pitch-deck business-specification.bspec
bspec generate business-plan --audience investor business-specification.bspec

# Extract specific domains
bspec extract --domain strategic,financial business-specification.bspec
```

## Versioning and Migration

### Package Versioning
Packages use semantic versioning (MAJOR.MINOR.PATCH):
- **MAJOR**: Breaking changes to business model or strategy
- **MINOR**: New documents or significant updates to existing documents
- **PATCH**: Minor corrections, typo fixes, asset updates

### Migration Support
When BSpec specification versions change, packages include migration metadata:

```json
{
  "migration": {
    "from_version": "0.9.0",
    "to_version": "1.0.0",
    "compatibility": "backwards_compatible",
    "changes": [
      {
        "type": "document_type_renamed",
        "from": "STRATEGY",
        "to": "STR",
        "documents_affected": ["STRATEGY-platform-approach-v1.0.0.md"]
      }
    ],
    "migration_guide": "See CHANGELOG.md for detailed migration instructions"
  }
}
```

## Security and Integrity

### Checksums
All files include SHA-256 checksums for integrity verification:

```json
{
  "checksums": {
    "manifest.json": "sha256:a1b2c3d4e5f6...",
    "documents/01-strategic/MSN-company-mission-v1.0.0.md": "sha256:1a2b3c4d5e...",
    "assets/diagrams/system-architecture.svg": "sha256:9z8y7x6w5v..."
  }
}
```

### Sensitive Data Handling
The .bspec format includes mechanisms for handling sensitive information:

- **Redaction markers**: `[REDACTED:financial-projections]` for public distribution
- **Sanitization flags**: Mark documents as `public`, `internal`, or `confidential`
- **Export filtering**: Generate sanitized versions for different audiences

## Compliance and Standards

### Industry Compliance
Packages can include compliance metadata for regulatory requirements:

```json
{
  "compliance": {
    "frameworks": ["SOC2", "GDPR", "CCPA"],
    "certifications": ["ISO27001"],
    "privacy_handling": "documented",
    "data_residency": ["us", "eu"],
    "retention_policy": "7_years"
  }
}
```

### Export Control
For international distribution, packages can include export control information:

```json
{
  "export_control": {
    "classification": "public",
    "restrictions": [],
    "approved_countries": ["*"],
    "technology_classification": "dual_use_exempt"
  }
}
```

## Implementation Guidelines

### Creating .bspec Packages
1. **Start with validation**: Ensure all documents pass BSpec schema validation
2. **Organize systematically**: Use standard directory structure and naming
3. **Include comprehensive assets**: Don't forget diagrams, images, and supporting documents
4. **Generate analysis**: Include computed relationship graphs and conformance reports
5. **Test extraction**: Verify package can be extracted and consumed correctly

### Consuming .bspec Packages
1. **Validate integrity**: Check checksums and schema compliance
2. **Understand context**: Read manifest.json for business context and conformance level
3. **Follow relationships**: Use dependency graphs to understand document relationships
4. **Leverage analysis**: Use pre-computed conformance and gap analysis
5. **Access programmatically**: Use SDKs for systematic access and analysis

## Future Extensions

The .bspec format is designed to evolve with the BSpec standard:

### Planned Enhancements
- **Digital signatures**: Cryptographic signing for authenticity verification
- **Collaborative metadata**: Track contributors, reviewers, and approval workflows
- **Temporal versioning**: Support for time-based document snapshots
- **Modular packaging**: Support for partial packages and cross-references
- **AI annotations**: Machine-generated insights and recommendations

### Backward Compatibility
All future versions will maintain backward compatibility with v1.0.0 packages, with clear migration paths for new features and capabilities.

---

The .bspec format enables the complete packaging, distribution, and consumption of business specifications, making BSpec a truly portable and interoperable standard for business documentation.