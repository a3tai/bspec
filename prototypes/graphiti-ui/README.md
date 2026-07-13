# BSpec Graphiti Workbench

Quick Mittsu/Svelte prototype for exploring BSpec graph data before wiring it to
a live Graphiti service.

## Catalog Mode

Catalog mode reads the BSpec standard from `spec/v1` and shows document type
nodes plus relationship-rule facts.

```bash
bun install
bun run build:data
bun run dev
```

## Existing BSpec Project Mode

Project mode reads an existing BSpec project directory, a `documents/` folder, or
a packed `.bspec` archive and emits the same UI data shape.

```bash
bun scripts/build-graph-data.js --input ./my-project --mode project
bun scripts/build-graph-data.js --input ./my-project.bspec
bun run dev
```

The generated `src/data/graph.json` is intentionally static for now. That keeps
the UI fast to iterate on while the Graphiti service boundary is still being
prototyped.

## License

MIT, matching BSpec and Mittsu.
