---
title: "Python SDK"
description: "Pythonic SDK for BSpec document management with Pydantic models"
---

# Python SDK

Work with BSpec documents using Python with full Pydantic validation and type hints.

## Installation

```bash
pip install bspec-sdk
```

## Quick Start

```python
from bspec import BSpec, DocumentTypes

# Create a Vision document
vision = DocumentTypes.VSN(
    id="VSN-product-2025",
    title="Product Vision 2025",
    type="VSN",
    status="Draft",
    version="1.0.0",
    owner="Product Team",
    domain="strategic-foundation"
)

# Validate
is_valid = vision.validate()

# Convert to dict/JSON
data = vision.model_dump()
json_str = vision.model_dump_json()

# Parse from JSON
vision = DocumentTypes.VSN.model_validate_json(json_str)
```

## Features

- Full Pydantic v2 models for all 112 document types
- Type hints and IDE autocomplete
- Automatic validation
- JSON/YAML serialization
- Markdown parsing and generation
- Relationship analysis
- Conformance checking

## Documentation

See the [GitHub repository](https://github.com/a3tai/bspec) for full documentation.
