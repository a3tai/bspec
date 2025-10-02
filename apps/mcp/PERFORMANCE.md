# MCP Server Performance Troubleshooting

## Issue: SSE Stream Disconnected

The "SSE stream disconnected: aborted" error was caused by Cloudflare Workers timing out during the MCP agent initialization phase.

## Root Causes Identified

1. **Heavy Synchronous Work During Init**: Loading the entire BSpec specification (TGZ decompression, TAR parsing, indexing 82+ document types) during the `init()` method exceeded Cloudflare Workers' 10-second CPU time limit.

2. **Large Resource Registration**: Registering 82+ document type resources plus domain resources all at once created excessive overhead.

3. **No Timeout Protection**: The specification loading process had no timeout or error recovery mechanism.

## Solutions Implemented

### 1. Lazy Loading Architecture
**Before**: Specification loaded synchronously during `init()`
```typescript
async init() {
  await this.ensureSpecificationLoaded();  // BLOCKS initialization
  this.setupResources();
  this.setupTools();
}
```

**After**: Specification loaded on first tool call
```typescript
async init() {
  // Only setup lightweight tools (no spec loading)
  this.setupTools();
  // Spec loads lazily on first request
}
```

### 2. Timeout Protection
Added 8-second timeout with race condition:
```typescript
private async _loadSpecWithTimeout(): Promise<void> {
  const TIMEOUT_MS = 8000; // 8s (Workers have 10s CPU limit)
  
  this.bspecIndex = await Promise.race([
    loadSpecificationFromR2(this.env),
    timeoutPromise
  ]);
}
```

### 3. Loading Deduplication
Prevents multiple concurrent specification loads:
```typescript
if (this._loadingPromise) {
  return this._loadingPromise;  // Reuse existing load
}
this._loadingPromise = this._loadSpecWithTimeout();
```

### 4. Minimal Resource Registration
Only registers overview resource during initialization. Document types and domains are accessible via tools (not as heavyweight resources).

## Timing Logs

Comprehensive timing logs are now available to diagnose performance bottlenecks:

### Viewing Logs Locally
```bash
cd apps/mcp
bun run dev
```

### Viewing Production Logs
```bash
cd apps/mcp
wrangler tail
# or
wrangler tail --env production
```

### Timing Log Format
All timing logs are prefixed with `[TIMING]` for easy filtering:

```
[TIMING] Loading BSpec specification from R2...
[TIMING] R2 fetch completed in 45ms
[TIMING] Buffer read completed in 12ms
[TIMING] TGZ file loaded, size: 123456 bytes
[TIMING] Decompressing TGZ file...
[TIMING] Decompression completed in 234ms
[TIMING] Parsing TAR contents...
[TIMING] TAR parsing completed in 89ms
[TIMING] Building BSpec intelligence index...
[TIMING] Index built in 156ms
[TIMING] ✅ Total specification load time: 536ms

[TIMING] Tool: get_document_type(MSN)
[TIMING] Spec loaded in 2ms (already cached)
[TIMING] get_document_type completed in 4ms
```

## Expected Performance

### First Request (Cold Start)
- **R2 Fetch**: 20-100ms
- **Buffer Read**: 5-20ms
- **Decompression**: 100-300ms
- **TAR Parsing**: 50-150ms
- **Index Building**: 100-200ms
- **Total**: 300-800ms ✅

### Subsequent Requests (Warm)
- Specification cached in memory
- Tool calls: 1-10ms ✅

## Monitoring Commands

### Check if server is timing out
```bash
wrangler tail --format json | grep -i "timeout\|abort\|error"
```

### View only timing logs
```bash
wrangler tail | grep "\[TIMING\]"
```

### Monitor tool call performance
```bash
wrangler tail | grep "Tool:"
```

### Check cold start vs warm performance
```bash
# First request after deploy (cold start)
curl https://mcp.bspec.dev/mcp

# Check logs for timing breakdown
wrangler tail --since 1m | grep TIMING
```

## Deployment

Deploy with timing logs enabled:
```bash
cd apps/mcp
bun run build
wrangler deploy --env production

# Monitor first requests
wrangler tail --env production | grep TIMING
```

## Further Optimizations (if needed)

If timing issues persist, consider:

1. **Pre-built Index**: Generate the BSpec index at build time and store as JSON
2. **KV Cache**: Cache the loaded specification in Cloudflare KV
3. **Streaming Responses**: Stream large document specifications incrementally
4. **Durable Object Storage**: Store specification in DO SQLite storage

## Testing

### Test Initialization Speed
```bash
# Start dev server and measure time to first response
time curl http://localhost:8787/mcp
```

### Test Tool Call Performance
Use the MCP Inspector:
```bash
npx @modelcontextprotocol/inspector http://localhost:8787/mcp
```

Then call tools and observe timing logs in the terminal.

## Common Issues

### "Specification loading timed out"
- TGZ file too large (check R2 bucket)
- Network latency to R2
- CPU-intensive decompression

**Solution**: Increase `TIMEOUT_MS` or implement KV caching

### "SSE stream disconnected"
- Initialization taking too long
- Large response payload

**Solution**: Already implemented lazy loading

### Memory exhaustion
- Loading too much data into memory

**Solution**: Stream responses or implement pagination

## Contact

For performance issues, collect timing logs and create an issue at:
https://github.com/a3tai/bspec/issues
