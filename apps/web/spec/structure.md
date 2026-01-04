---
title: "Document Structure"
description: "Standard structure for BSpec documents"
---

# Document Structure

Every BSpec document follows a consistent structure that makes it both human-readable and machine-parseable.

## File Format

BSpec documents are written in **Markdown** with **YAML frontmatter**:

```markdown
---
# YAML metadata here
---

# Markdown content here
```

## Complete Document Template

```yaml
---
# ============================================================================
# CORE IDENTITY
# ============================================================================
id: {TYPE}-{kebab-case-identifier}
type: {TYPE}
title: "{Document Title}"
status: Draft|Review|Approved|Active|Deprecated
version: 1.0.0

# ============================================================================
# OWNERSHIP & STAKEHOLDERS
# ============================================================================
owner: {owner-name}
stakeholders: [stakeholder-1, stakeholder-2]

# ============================================================================
# BUSINESS CONTEXT
# ============================================================================
domain: {domain-name}
priority: Critical|High|Medium|Low
scope: {scope-description}
horizon: tactical|strategic|operational
visibility: public|internal|confidential

# ============================================================================
# RELATIONSHIPS (Creates Knowledge Graph)
# ============================================================================
depends_on: [document-ids-this-depends-on]
enables: [document-ids-this-enables]
conflicts_with: [document-ids-that-conflict]
related_to: [related-document-ids]

# ============================================================================
# SUCCESS & VALIDATION
# ============================================================================
success_criteria:
  - "Measurable success criterion 1"
  - "Measurable success criterion 2"

assumptions:
  - "Key assumption 1"
  - "Key assumption 2"

constraints: [constraint-1, constraint-2]

# ============================================================================
# METADATA
# ============================================================================
standards: [relevant-standards]
review_cycle: weekly|monthly|quarterly|annually
tags: [tag1, tag2]

# ============================================================================
# TYPE-SPECIFIC FIELDS
# ============================================================================
# Each document type has its own specific fields here
# See individual document type specifications
---

# {Document Title}

## Executive Summary

Brief overview of the document purpose and key points.

## Main Content

Document content in Markdown format with:
- Headings
- Lists
- Tables
- Code blocks
- Links

## Appendices

Additional supporting information.
```

## Required Fields

All documents **must** have:

- `id` - Unique identifier in format `{TYPE}-{identifier}`
- `type` - Three-letter document type code (e.g., VSN, STR, MSN)
- `title` - Human-readable title
- `status` - Current document status
- `version` - Semantic version number
- `owner` - Person or team responsible
- `domain` - Business domain the document belongs to

## Optional Fields

Documents **may** include:

- `stakeholders` - People or teams who are stakeholders
- `priority` - Business priority level
- `scope` - What the document covers
- `horizon` - Time horizon (tactical/strategic/operational)
- `visibility` - Who can access this document
- `success_criteria` - How to measure success
- `assumptions` - Key assumptions made
- `constraints` - Limiting factors
- `standards` - Standards this conforms to
- `review_cycle` - How often to review
- `tags` - Classification tags

## Relationship Fields

These fields create the business knowledge graph:

### `depends_on`
Documents this document requires to be meaningful.

```yaml
depends_on:
  - VSN-product-2025  # Needs vision
  - STR-growth-2025   # Needs strategy
```

### `enables`
Documents that this document makes possible.

```yaml
enables:
  - PRD-feature-a     # Enables product requirements
  - BRD-integration-b # Enables business requirements
```

### `conflicts_with`
Mutually exclusive documents.

```yaml
conflicts_with:
  - STR-alternative-approach  # Can't coexist with alternative
```

### `related_to`
Related but not dependent documents.

```yaml
related_to:
  - MKT-analysis-q4  # Related market analysis
  - CMP-landscape    # Related competitive info
```

## Status Values

Documents progress through these statuses:

- **Draft** - Initial creation, work in progress
- **Review** - Ready for stakeholder review
- **Approved** - Approved but not yet active
- **Active** - Currently in effect
- **Deprecated** - No longer valid, kept for reference

## Version Format

Use semantic versioning: `MAJOR.MINOR.PATCH`

- **MAJOR** - Breaking changes (1.0.0 → 2.0.0)
- **MINOR** - New features, backward compatible (1.0.0 → 1.1.0)
- **PATCH** - Bug fixes, minor updates (1.0.0 → 1.0.1)

## Document Type Codes

All 82 BSpec document types use three-letter codes:

### Strategic Foundation (8)
VSN, STR, MSN, VAL, OBJ, PUR, THY, MOT

### Market Environment (10)
MKT, SEG, CMP, ECO, TRN, REG, OPP, POS, THR, MAC

### Customer Value (13)
CUS, PER, USE, STO, PAI, FEE, CJM, SUR, JTB, CIN, EMP, GAI, BEH

### Product & Service (10)
PRD, FEA, SVC, REQ, QUA, UXD, ROD, INT, SUP, PSP

### Business Model (9)
REV, PRI, CST, VST, CHN, KPT, KRS, KAC, CUS

### Operations & Execution (12)
OPS, PRO, WFL, ORG, TEA, ROL, POL, VND, SKI, CAP, SLA, FAC

### Technology & Data (8)
SYS, API, DAT, ARC, DEV, INF, SEC, ANA

### Financial & Investment (10)
FIN, BUD, FOR, FND, INV, VLU, MET, REP, TAX, AUD

### Risk & Governance (8)
RSK, GOV, COM, INS, INC, CTL, LEG, ETH

### Growth & Innovation (7)
INN, EXP, RND, FUT, LEA, ADT, IGN

### Learning & Decisions (6)
DEC, KNO, LRN, HYP, RET, WIS

### Brand & Marketing (12)
BRD, MSG, CNT, CAM, SOC, SEO, MCH, BPO, IFL, TON, VID, PRF

## File Naming Convention

Files should be named: `{TYPE}-{identifier}.md`

**Examples:**
- `VSN-product-2025.md`
- `STR-growth-strategy.md`
- `MSN-company-mission.md`
- `BRD-feature-x-requirements.md`

## Best Practices

1. **One Concern Per Document**: Each document should cover exactly one topic
2. **Clear IDs**: Use descriptive, human-readable identifiers
3. **Explicit Relationships**: Always declare dependencies and enablements
4. **Meaningful Status**: Keep status current and accurate
5. **Version Everything**: Increment version with every meaningful change
6. **Measure Success**: Include concrete success criteria
7. **Regular Reviews**: Set appropriate review cycles
8. **Rich Context**: Provide enough context for the document to stand alone

## Validation

Documents can be validated using:

- **CLI**: `bspec validate docs/`
- **TypeScript SDK**: `validate(document)`
- **Python SDK**: `document.validate()`
- **Go SDK**: `doc.Validate()`

## Examples

See the [document types](/docs/document-types) for specific examples of each type.

## Next Steps

- View [All Document Types](/docs/document-types)
- Learn about [Conformance Levels](/spec/conformance)
- Read the [Full Specification](/spec/v1-0-0)
