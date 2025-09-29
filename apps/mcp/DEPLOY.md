# BSpec MCP Server Deployment Guide

This guide covers deploying the BSpec Model Context Protocol (MCP) server to Cloudflare Workers.

## Prerequisites

1. **Cloudflare Account**: Sign up at [cloudflare.com](https://cloudflare.com)
2. **Domain Setup**: Configure `bspec.dev` domain in Cloudflare (or update routes in `wrangler.toml`)
3. **Wrangler CLI**: Install with `npm install -g wrangler`
4. **Authentication**: Run `wrangler login` to authenticate

## Quick Deployment

```bash
# Install dependencies
npm install

# Build and deploy to production
npm run deploy:prod

# Or deploy to development environment
npm run deploy:dev
```

## Deployment Environments

### Development Environment
- **URL**: `https://mcp-dev.bspec.dev/mcp`
- **Purpose**: Testing and development
- **Command**: `npm run deploy:dev`

### Production Environment
- **URL**: `https://mcp.bspec.dev/mcp`
- **Purpose**: Live MCP server for AI agents
- **Command**: `npm run deploy:prod`

## Configuration

### Environment Variables

Set these in your Cloudflare dashboard or via `wrangler secret put`:

```bash
# Optional: Restrict CORS origins in production
wrangler secret put CORS_ORIGIN

# Optional: API rate limiting
wrangler secret put RATE_LIMIT_RPM
```

### Custom Domain Setup

1. **Add Domain to Cloudflare**: Add `bspec.dev` to your Cloudflare account
2. **Update Routes**: Modify `wrangler.toml` routes if using a different domain
3. **SSL Certificate**: Cloudflare automatically provisions SSL certificates

## MCP Server Endpoints

### Core MCP Protocol
- `POST /mcp` - Main MCP JSON-RPC endpoint
- `GET /mcp/health` - Health check endpoint
- `GET /mcp/docs` - API documentation

### Available Tools

1. **parse_bspec_document** - Parse and validate BSpec documents
2. **validate_bspec_document** - Comprehensive document validation
3. **query_documents** - Search and filter documents
4. **analyze_relationships** - Relationship graph analysis
5. **get_document_types** - List all 82 BSpec document types
6. **convert_to_json** - Convert documents to JSON format

### Available Resources

1. 

## Usage with AI Agents

### Claude Code Integration

Add to your MCP client configuration:

```json
{
  "mcpServers": {
    "bspec": {
      "command": "node",
      "args": ["-e", "/* HTTP MCP client pointing to https://mcp.bspec.dev/mcp */"]
    }
  }
}
```

### Example Usage

```javascript
// Query all strategic documents
const strategicDocs = await mcp.call('query_documents', {
  filters: { domain: 'strategic' },
  includeContent: true
});

// Validate a new document
const validation = await mcp.call('validate_bspec_document', {
  document: myDocument,
  validateRelationships: true
});

// Get mission statement template
const template = await mcp.read('bspec://document-types/MSN');
```

## Monitoring and Debugging

### View Logs

```bash
# Real-time logs for production
npm run logs

# Development environment logs
npm run logs:dev
```

### Health Checks

```bash
# Check server health
curl https://mcp.bspec.dev/mcp/health

# Expected response
{
  "status": "healthy",
  "version": "1.0.0",
  "timestamp": "2024-01-15T10:30:00Z",
  "environment": "production"
}
```

### Performance Monitoring

- **Cloudflare Analytics**: View request metrics in Cloudflare dashboard
- **Worker Metrics**: CPU time, memory usage, and error rates
- **Custom Analytics**: Track MCP tool usage and document operations

## Troubleshooting

### Common Issues

1. **Build Errors**
   ```bash
   npm run clean
   npm install
   npm run build
   ```

2. **Route Conflicts**
   - Verify domain is added to Cloudflare
   - Check `wrangler.toml` route patterns
   - Ensure DNS is pointing to Cloudflare

3. **Authentication Issues**
   ```bash
   wrangler logout
   wrangler login
   ```

4. **Memory/CPU Limits**
   - Check document size limits in validator
   - Monitor Worker metrics in dashboard
   - Consider implementing caching for large documents

### Debug Mode

```bash
# Local development with debugging
npm run preview

# Test MCP endpoints locally
curl http://localhost:8787/mcp/health
```

## Security Considerations

1. **CORS Configuration**: Properly configure origins in production
2. **Rate Limiting**: Implement rate limiting for public endpoints
3. **Input Validation**: All inputs are validated via Zod schemas
4. **Error Handling**: Errors don't expose internal implementation details

## Updating the Server

### Code Updates

```bash
# Pull latest changes
git pull origin main

# Build and deploy
npm run deploy:prod
```

### Specification Updates

When BSpec specification changes:

1. Update document type definitions in `src/bspec/`
2. Run validation tests
3. Deploy updated server
4. Verify all tools and resources work correctly

## Support

- **Issues**: [GitHub Issues](https://github.com/a3tai/bspec-foundations/issues)
- **Documentation**: [bspec.dev/docs](https://bspec.dev/docs)
- **Community**: [BSpec Community Discussions](https://github.com/a3tai/bspec-foundations/discussions)