# A3T Business Workbench

A native Wails v3 workbench for the canonical A3T business service.
The application manages business identity, structured records, immutable history, BSpec imports, and bounded graph exploration through the service REST API.

PostgreSQL-backed service reads are authoritative.
Neo4j supplies focused current-state graph neighborhoods and paths.
BSpec is an import and interchange format rather than a parallel local database.

## Start the business service

From the A3T Core repository:

```bash
cd /Users/steverude/prj/a3tai/core
docker compose -f infra/docker-compose.dev.yml --profile business up -d --build business-api business-projector
curl -fsS http://localhost:7290/ready
```

The local Compose profile supplies a fixed development tenant and actor.
The desktop app therefore needs no token or development identity headers for the default local workflow.

## Run the desktop app

```bash
cd /Users/steverude/prj/a3tai/bspec/apps/graph
PACKAGE_MANAGER=bun wails3 task dev
```

If port `9255` is already in use:

```bash
PACKAGE_MANAGER=bun WAILS_VITE_PORT=9265 wails3 task dev
```

## Connection configuration

The native Go client defaults to `http://localhost:7290`.
The Connection panel can switch the service URL and set an in-memory bearer token or development identity headers.
After configuration, the native client owns the active credentials and never returns them through status or read bindings.
Hidden credentials can be preserved only when reconnecting to the same normalized service URL.
Bearer credentials require HTTPS except for loopback development services.
An unavailable candidate connection never replaces the last working client.

The initial configuration can also come from environment variables:

- `A3T_BUSINESS_API_URL` or `BUSINESS_API_URL`
- `A3T_BUSINESS_ACCESS_TOKEN` or `A3T_ACCESS_TOKEN`
- `A3T_BUSINESS_TENANT_ID` or `A3T_TENANT_ID`
- `A3T_BUSINESS_ACTOR_ID` or `A3T_ACTOR_ID`

## Workbench lenses

- Landscape shows a bounded projected neighborhood, explicit relationship direction, unresolved terminal references, truncation state, and shortest path explanations.
- Record reads canonical narrative and metadata, then creates a complete next revision with optimistic concurrency.
- History reconstructs immutable record snapshots and separates content changes from graph changes.
- Identity distinguishes the managed business from its owning tenant, optional organization association, and external identifiers.

The persistent record index is the keyboard-accessible alternative to spatial graph navigation.
The UI distinguishes canonical loading, empty, stale, forbidden, pending projection, graph unavailable, unresolved, and truncated states.

## BSpec import review

The import workflow always runs a service dry run first.
It reports created, revised, unchanged, resolved, unresolved, and diagnostic counts before an explicit apply action.
The native client freezes the exact reviewed bytes in a bounded, expiring in-memory review and reuses one durable operation ID across ambiguous apply retries.

`.bspec` archives are staged and sent as bounded multipart uploads.
Directories are converted into bounded JSON document imports without recreating service importer semantics in the client.

## Build and verify

```bash
cd /Users/steverude/prj/a3tai/bspec/apps/graph
go test ./...
cd frontend
bun run check
bun test
bun run build
cd ..
PACKAGE_MANAGER=bun wails3 build
PACKAGE_MANAGER=bun wails3 package
```

The native binary is written to `bin/a3t-business-workbench`.
On macOS, `wails3 package` creates `bin/a3t-business-workbench.app`.

## Service contract

The normative REST contract is `/Users/steverude/prj/a3tai/core/services/business/openapi.yaml`.
The product interaction contract is `/Users/steverude/prj/a3tai/core/services/business/UI.md`.

The Wails Go boundary owns service URLs, bearer credentials, pagination cursors, `ETag` and `If-Match`, idempotency keys, structured errors, and multipart streaming.
The webview consumes typed Wails bindings and does not make cross-origin REST requests.
