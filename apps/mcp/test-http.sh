#!/bin/bash

# Test the BSpec MCP Server with HTTP requests

echo "Testing BSpec MCP Server..."

# Check if server is running
SERVER_URL="${SERVER_URL:-http://localhost:8787}"

# Test health endpoint
echo ""
echo "1. Testing health endpoint:"
curl -s "$SERVER_URL/health" | jq . || echo "Health check failed"

# Test MCP initialize
echo ""
echo "2. Testing MCP initialize:"
curl -s -X POST "$SERVER_URL/mcp" \
  -H "Content-Type: application/json" \
  -H "Accept: application/json" \
  -d '{
    "jsonrpc": "2.0",
    "id": 1,
    "method": "initialize",
    "params": {
      "protocolVersion": "2024-11-05",
      "capabilities": {},
      "clientInfo": {
        "name": "test-client",
        "version": "1.0.0"
      }
    }
  }' | jq . || echo "Initialize failed"

# Send initialized notification
echo ""
echo "3. Sending initialized notification:"
curl -s -X POST "$SERVER_URL/mcp" \
  -H "Content-Type: application/json" \
  -H "Accept: application/json" \
  -d '{
    "jsonrpc": "2.0",
    "method": "notifications/initialized"
  }' || echo "Initialized notification sent"

# Test tools/list
echo ""
echo "4. Testing tools/list:"
curl -s -X POST "$SERVER_URL/mcp" \
  -H "Content-Type: application/json" \
  -H "Accept: application/json" \
  -d '{
    "jsonrpc": "2.0",
    "id": 2,
    "method": "tools/list"
  }' | jq '.result.tools[] | {name: .name, description: .description}' || echo "Tools list failed"

# Test resources/list
echo ""
echo "5. Testing resources/list:"
curl -s -X POST "$SERVER_URL/mcp" \
  -H "Content-Type: application/json" \
  -H "Accept: application/json" \
  -d '{
    "jsonrpc": "2.0",
    "id": 3,
    "method": "resources/list"
  }' | jq '.result.resources | length' || echo "Resources list failed"

# Test a tool call
echo ""
echo "6. Testing get_bspec_overview tool:"
curl -s -X POST "$SERVER_URL/mcp" \
  -H "Content-Type: application/json" \
  -H "Accept: application/json" \
  -d '{
    "jsonrpc": "2.0",
    "id": 4,
    "method": "tools/call",
    "params": {
      "name": "get_bspec_overview",
      "arguments": {}
    }
  }' | jq '.result.content[0].text' | head -20 || echo "Tool call failed"

echo ""
echo "Tests complete!"