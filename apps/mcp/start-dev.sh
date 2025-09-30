#!/bin/bash

# Start the BSpec MCP server in development mode

echo "🚀 Starting BSpec MCP Server in development mode..."
echo ""

# Build first
echo "📦 Building..."
bun run build

# Start wrangler dev
echo ""
echo "🔄 Starting Wrangler dev server on http://localhost:8787"
echo "   Press Ctrl+C to stop"
echo ""
echo "   Available endpoints:"
echo "   - http://localhost:8787/health (health check)"
echo "   - http://localhost:8787/ (web interface)"
echo "   - http://localhost:8787/mcp (MCP endpoint)"
echo ""

wrangler dev --port 8787 --local