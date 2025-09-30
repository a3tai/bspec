# Testing the BSpec MCP Server

This guide explains how to test the BSpec MCP server locally.

## Quick Start

### 1. Start the Development Server

```bash
./start-dev.sh
```

Or manually:
```bash
bun run build
bun run dev
```

The server will be available at `http://localhost:8787`

### 2. Test with HTTP Requests

Run the test script:
```bash
./test-http.sh
```

Or test individual endpoints manually:

#### Health Check
```bash
curl http://localhost:8787/health | jq
```

#### MCP Initialize
```bash
curl -X POST http://localhost:8787/mcp \
  -H "Content-Type: application/json" \
  -H "Accept: application/json" \
  -d '{
    "jsonrpc": "2.0",
    "id": 1,
    "method": "initialize",
    "params": {
      "protocolVersion": "2024-11-05",
      "capabilities": {},
      "clientInfo": {"name": "test-client", "version": "1.0.0"}
    }
  }' | jq
```

#### List Tools
```bash
curl -X POST http://localhost:8787/mcp \
  -H "Content-Type: application/json" \
  -H "Accept: application/json" \
  -d '{
    "jsonrpc": "2.0",
    "id": 2,
    "method": "tools/list"
  }' | jq
```

#### List Resources
```bash
curl -X POST http://localhost:8787/mcp \
  -H "Content-Type: application/json" \
  -H "Accept: application/json" \
  -d '{
    "jsonrpc": "2.0",
    "id": 3,
    "method": "resources/list"
  }' | jq
```

### 3. Test with MCP Inspector

The MCP Inspector provides a web UI for testing MCP servers.

#### Using Config File
```bash
npx @modelcontextprotocol/inspector --config mcp-config.json --server bspec
```

#### Direct URL
```bash
npx @modelcontextprotocol/inspector --transport http --server-url http://localhost:8787/mcp
```

This will open a web interface where you can:
- Browse available tools and resources
- Test tool calls interactively
- View responses in real-time

## Available Tools

The BSpec MCP server provides the following tools:

1. **get_document_type** - Get detailed information about a specific BSpec document type
2. **list_document_types** - List all BSpec document types with filtering options
3. **get_domain_info** - Get information about a specific BSpec business domain
4. **list_domains** - List all BSpec business domains
5. **get_bspec_overview** - Get a comprehensive overview of the BSpec specification
6. **search_document_types** - Search for document types by name or description
7. **get_document_relationships** - Get relationship information for document types

## Available Resources

Resources are accessed via URIs like:
- `https://mcp.bspec.dev/spec/v1.0.0/overview` - BSpec overview
- `https://mcp.bspec.dev/spec/v1.0.0/domains/{domain-key}` - Domain info
- `https://mcp.bspec.dev/spec/v1.0.0/document-types/{TYPE-CODE}` - Document type specs

## Troubleshooting

### Server Won't Start

1. **Check R2 bucket setup**: Make sure the TGZ files are in `.wrangler/state/v3/r2/bspec-assets/`
   ```bash
   ls -la .wrangler/state/v3/r2/bspec-assets/
   ```

2. **Rebuild**: 
   ```bash
   bun run clean
   bun run build
   ```

3. **Check logs**: Look for errors in the wrangler output

### No Response from Endpoints

1. Verify server is running: `curl http://localhost:8787/health`
2. Check the wrangler logs for errors
3. Ensure you're using the correct headers (Content-Type and Accept)

### MCP Inspector Can't Connect

1. Make sure the server is running on port 8787
2. Verify the URL in mcp-config.json is correct
3. Try the direct URL method with `--server-url`

## Development Notes

- The server loads the BSpec specification from a TGZ file stored in R2 (or local emulation)
- Changes to TypeScript files require rebuilding with `bun run build`
- The dev server auto-rebuilds on file changes
- Use `--local` flag with wrangler dev for local R2 emulation