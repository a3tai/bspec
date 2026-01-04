---
title: "MCP Server"
description: "Model Context Protocol server for AI integration"
---

# MCP Server

The BSpec MCP Server enables AI assistants like Claude to work with your business specifications using the Model Context Protocol.

## What is MCP?

The Model Context Protocol (MCP) is an open standard that allows AI assistants to securely access your data and tools. The BSpec MCP server provides AI agents with structured access to your business documentation.

## Features

- **Document Type Information**: AI can query all 112 document types
- **Validation**: AI can validate documents against the BSpec standard
- **Conformance Analysis**: AI can analyze your spec for completeness
- **Relationship Mapping**: AI can understand document dependencies
- **Template Generation**: AI can generate new documents from templates

## Installation

The MCP server is deployed at: `https://mcp.bspec.dev/mcp`

## Usage with Claude Desktop

Add to your Claude Desktop configuration:

```json
{
  "mcpServers": {
    "bspec": {
      "url": "https://mcp.bspec.dev/mcp"
    }
  }
}
```

## Usage with Warp

The MCP server is automatically available in Warp terminal with AI features enabled.

## Available Tools

- `get_document_type` - Get information about a specific document type
- `list_document_types` - List all 112 document types
- `search_document_types` - Search for document types
- `get_document_relationships` - Get relationship information
- `get_domain_info` - Get information about a business domain
- `list_domains` - List all business domains
- `get_bspec_overview` - Get BSpec overview

## Example Queries

Ask Claude or other AI assistants:

- "What document types does BSpec define for strategic planning?"
- "Show me the Vision (VSN) document type specification"
- "What documents should I create to reach Silver conformance?"
- "Generate a template for a Mission Statement document"
- "What's the relationship between Strategy and Vision documents?"

## API Documentation

See the [MCP Server API Documentation](https://mcp.bspec.dev/) for complete details.

## Self-Hosting

Deploy your own MCP server:

```bash
git clone https://github.com/a3tai/bspec.git
cd bspec/apps/mcp
bun install
bun run deploy
```

## Support

- **Documentation**: [https://bspec.dev/docs](https://bspec.dev/docs)
- **GitHub**: [a3tai/bspec](https://github.com/a3tai/bspec)
- **Issues**: [Report bugs](https://github.com/a3tai/bspec/issues)
