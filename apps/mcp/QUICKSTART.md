# Quick Start: Testing BSpec MCP Server

## Step 1: Start the Server

Open a terminal and run:

```bash
cd /Users/steverude/prj/a3tai/bspec/apps/mcp
./start-dev.sh
```

Keep this terminal open. The server will run on `http://localhost:8787`

## Step 2: Test with MCP Inspector

Open a **new terminal** and run:

```bash
cd /Users/steverude/prj/a3tai/bspec/apps/mcp
npx @modelcontextprotocol/inspector --transport http --server-url http://localhost:8787/mcp
```

Or using the config file:

```bash
npx @modelcontextprotocol/inspector --config mcp-config.json --server bspec
```

This will:
1. Start a web server (usually on http://localhost:5173)
2. Open your browser automatically
3. Show you the MCP Inspector interface

## Step 3: Use the Inspector

In the MCP Inspector web interface, you can:

1. **View Tools** - See all 7 available tools
2. **View Resources** - See all BSpec specification resources
3. **Test Tools** - Click any tool to try it out
4. **View Responses** - See the full JSON responses

### Try These First:

1. **get_bspec_overview** - Get a summary of BSpec (no parameters needed)
2. **list_domains** - See all 12 business domains
3. **get_document_type** with `type_code: "MSN"` - Get Mission Statement spec

## Alternative: Quick HTTP Test

If you just want to verify the server is working:

```bash
cd /Users/steverude/prj/a3tai/bspec/apps/mcp
./test-http.sh
```

This runs a series of HTTP tests and shows the results.

## Troubleshooting

**"Connection refused" error?**
- Make sure the server is running in the first terminal
- Check that you're using port 8787

**MCP Inspector won't start?**
- Try: `npx -y @modelcontextprotocol/inspector ...`
- The `-y` flag auto-installs if needed

**Build errors?**
- Run: `bun install` then `bun run build`

**Need to reset everything?**
```bash
bun run clean
rm -rf .wrangler
bun run build
./start-dev.sh
```