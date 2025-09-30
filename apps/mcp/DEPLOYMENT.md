# BSpec MCP Server - Deployment Guide

## 🚀 Production Status: ✅ LIVE

**Current Version**: 1.0.0  
**Deployment Date**: 2025-09-30  
**Status**: Healthy and operational

## Production Endpoints

### Primary Service
- **MCP Endpoint**: https://mcp.bspec.dev/mcp
- **Health Check**: https://mcp.bspec.dev/health
- **Web Interface**: https://mcp.bspec.dev/

### Backup (Workers Domain)  
- https://bspec-prod.rude-company-llc.workers.dev

## Quick Start

### Testing the Live Server
```bash
# Health check
curl https://mcp.bspec.dev/health

# Test with MCP Inspector
npx @modelcontextprotocol/inspector --transport http --server-url https://mcp.bspec.dev/mcp
```

### Integration with AI Agents
Add to your MCP client configuration:
```json
{
  "mcpServers": {
    "bspec": {
      "transport": "http",
      "url": "https://mcp.bspec.dev/mcp"
    }
  }
}
```

## Available Capabilities

### 7 MCP Tools
1. **get_bspec_overview** - Comprehensive BSpec specification overview
2. **list_domains** - List all 12 business domains
3. **get_domain_info** - Get detailed domain information
4. **list_document_types** - List all 82+ document types with filtering
5. **get_document_type** - Get complete specification for a document type
6. **search_document_types** - Search document types by keywords
7. **get_document_relationships** - Understand document dependencies and relationships

### MCP Resources
Access via standard MCP resource URIs:
- `https://mcp.bspec.dev/spec/v1-0-0/overview`
- `https://mcp.bspec.dev/spec/v1-0-0/domains/{domain-key}`
- `https://mcp.bspec.dev/spec/v1-0-0/document-types/{TYPE-CODE}`

## Development

### Local Development
```bash
# Install dependencies
cd apps/mcp
bun install

# Start development server
bun run dev
# or
./start-dev.sh

# Test endpoints
./test-http.sh
```

### Deployment Commands
```bash
# Deploy to production
bun run deploy:prod

# Deploy to development
bun run deploy:dev

# View logs
bun run logs
```

## Infrastructure

### Platform Details
- **Platform**: Cloudflare Workers
- **R2 Bucket**: bspec-assets
- **Custom Domain**: mcp.bspec.dev
- **Protocol Support**: MCP 2024-11-05, 2025-03-26, 2025-06-18
- **Transport**: HTTP (primary), SSE (optional)

### Architecture
```
┌─────────────────────────────────────────┐
│         Cloudflare Workers              │
│                                         │
│  ┌─────────────────────────────────┐   │
│  │   BSpec MCP Server (index.ts)   │   │
│  │   - HTTP/SSE Transport          │   │
│  │   - MCP Protocol Handler        │   │
│  │   - CORS & Error Handling       │   │
│  └──────────────┬──────────────────┘   │
│                 │                       │
│  ┌──────────────▼──────────────────┐   │
│  │   Server Implementation         │   │
│  │   (server.ts)                   │   │
│  │   - TGZ Parser                  │   │
│  │   - Document Indexer            │   │
│  │   - Tool Implementations        │   │
│  │   - Resource Handler            │   │
│  └──────────────┬──────────────────┘   │
│                 │                       │
│  ┌──────────────▼──────────────────┐   │
│  │   R2 Bucket (bspec-assets)      │   │
│  │   - bspec-v1-0-0.tgz           │   │
│  │   - bspec-v1.0.0.tgz           │   │
│  │   - web/index.html              │   │
│  └─────────────────────────────────┘   │
└─────────────────────────────────────────┘
            │
            ▼
    https://mcp.bspec.dev
```

## Troubleshooting

### Common Issues

**Build Errors**
```bash
bun run clean
bun install
bun run build
```

**Authentication Issues**  
```bash
wrangler logout
wrangler login
```

**Route Conflicts**
- Verify domain is added to Cloudflare
- Check `wrangler.toml` route patterns
- Ensure DNS is pointing to Cloudflare

### Debug Mode
```bash
# Local development with debugging
bun run preview

# Test local endpoints
curl http://localhost:8787/mcp/health
```

### Health Monitoring
Expected health response:
```json
{
  "status": "healthy",
  "name": "bspec-specification-server",
  "version": "1.0.0",
  "timestamp": "2025-09-30T..."
}
```

## Security

1. **CORS Configuration**: Enabled for all origins in development
2. **Input Validation**: All inputs validated via Zod schemas  
3. **Error Handling**: Errors don't expose internal details
4. **Rate Limiting**: Consider implementing for production usage

## Updates and Maintenance

### Updating the Server
```bash
# Pull latest changes
git pull origin main

# Build and deploy
bun run deploy:prod
```

### Rollback Procedure
```bash
# View deployments
wrangler deployments list

# Rollback to specific version
wrangler rollback [version-id]
```

### Monitoring
```bash
# View production logs
bun run logs

# Check deployments
wrangler deployments list

# Monitor health
curl https://mcp.bspec.dev/health
```

## Example Usage

### Basic MCP Tool Call
```bash
curl -X POST https://mcp.bspec.dev/mcp \
  -H "Content-Type: application/json" \
  -d '{
    "jsonrpc": "2.0",
    "id": 1,
    "method": "tools/call",
    "params": {
      "name": "get_document_type",
      "arguments": {"type_code": "MSN"}
    }
  }'
```

### Resource Access
```bash
curl -X POST https://mcp.bspec.dev/mcp \
  -H "Content-Type: application/json" \
  -d '{
    "jsonrpc": "2.0",
    "id": 2,
    "method": "resources/read",
    "params": {
      "uri": "https://mcp.bspec.dev/spec/v1.0.0/overview"
    }
  }'
```

---

**The BSpec MCP Server is live and ready for AI agent integration!** 🎉

For support: https://github.com/a3tai/bspec/issues

## Support

For issues or questions:
- Repository: https://github.com/a3tai/bspec-foundations
- Issues: https://github.com/a3tai/bspec-foundations/issues