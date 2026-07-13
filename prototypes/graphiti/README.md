# BSpec Graphiti Prototype

This prototype keeps the BSpec specification as the source of truth and projects it into Graphiti as structured episodes. The goal is to test whether BSpec's document taxonomy, domains, and relationship rules become more useful when backed by a temporal graph database and Graphiti's hybrid search.

## What This Proves

- BSpec document type specs can become Graphiti episodes without changing the spec.
- Domains become first-class graph context.
- Document type relationships such as `depends_on`, `enables`, `conflicts_with`, and `related` are preserved as structured episode facts.
- The same exported data can be ingested directly with `graphiti-core` or through Graphiti's FastAPI service.

## Architecture

```text
spec/v1/*/*-spec.md
  -> bspec_graphiti.py export
  -> bspec-episodes.jsonl
  -> Graphiti Core or Graphiti FastAPI service
  -> Neo4j/FalkorDB/Neptune-backed temporal context graph
  -> hybrid graph search for agents and tools
```

The first prototype targets Neo4j because it is Graphiti's default documented local backend. FalkorDB and Neptune are viable follow-ups once the BSpec episode shape settles.

## Offline Commands

These commands do not require Graphiti, Neo4j, or an API key.

```bash
pyenv install -s 3.14.4
pyenv local 3.14.4
python prototypes/graphiti/bspec_graphiti.py summary
python prototypes/graphiti/bspec_graphiti.py export
```

To include one explicit episode for every relationship rule:

```bash
python prototypes/graphiti/bspec_graphiti.py export --include-relationship-rules
```

## Existing BSpec Projects

The same prototype can export a user-authored BSpec project, either as an extracted
project directory or a packed `.bspec` archive. This follows the existing CLI
shape: `manifest.json`, `documents/*.md`, optional `assets/`, and optional
`computed/`.

```bash
python prototypes/graphiti/bspec_graphiti.py summary-bspec ./my-project
python prototypes/graphiti/bspec_graphiti.py export-bspec ./my-project \
  --include-relationship-facts

python prototypes/graphiti/bspec_graphiti.py summary-bspec ./my-project.bspec
python prototypes/graphiti/bspec_graphiti.py export-bspec ./my-project.bspec \
  --output prototypes/graphiti/bspec-project-episodes.jsonl \
  --include-relationship-facts
```

Catalog exports use `bspec.document_type` and `bspec.relationship_rule` episodes.
Project exports use `bspec.archive`, `bspec.document`, and
`bspec.document_relationship` episodes. Keep these in separate Graphiti
`group_id`s when comparing the normative standard against a real project.

## Direct Graphiti Core Path

Use this path when you want the richest structured ingestion.

```bash
pyenv install -s 3.14.4
pyenv local 3.14.4
python -m venv .venv
. .venv/bin/activate
python -m pip install -r prototypes/graphiti/requirements.txt

export OPENAI_API_KEY=...
export NEO4J_URI=bolt://localhost:7687
export NEO4J_USER=neo4j
export NEO4J_PASSWORD=password

python prototypes/graphiti/bspec_graphiti.py export
python prototypes/graphiti/bspec_graphiti.py ingest-core
python prototypes/graphiti/bspec_graphiti.py search-core "What does a PRD depend on?"
```

## Graphiti Service Path

Use this path when evaluating Graphiti as an independently deployed service.

```bash
cp prototypes/graphiti/.env.example prototypes/graphiti/.env
# edit OPENAI_API_KEY if needed
docker compose --env-file prototypes/graphiti/.env -f prototypes/graphiti/docker-compose.yml up

python prototypes/graphiti/bspec_graphiti.py export
python prototypes/graphiti/bspec_graphiti.py ingest-service
python prototypes/graphiti/bspec_graphiti.py search-service "Which BSpec documents enable product requirements?"
```

The service path maps each BSpec JSON episode into Graphiti's `/messages` endpoint because the current service API is message-oriented. The direct Core path uses `EpisodeType.json` and is the better fit for BSpec's structured spec data.

## Prototype Questions

- Should BSpec define a prescribed Graphiti ontology with Pydantic entity types for `BSpecDomain`, `BSpecDocumentType`, `BSpecDocumentInstance`, `BusinessCapability`, `Stakeholder`, and `Metric`?
- Should relationship rules be stored as Graphiti-derived facts only, or also written as native graph edges in a sidecar adapter?
- Should the MCP server query Graphiti for semantic/temporal context while retaining the current static validator/parser for deterministic conformance?
- Should user-authored BSpec documents be ingested as separate episodes from the normative spec, grouped by organization/project?
- Should the CLI write Graphiti-ready JSONL into `computed/graphiti-episodes.jsonl` during `bspec pack`?

## Suggested Next Step

Wire one BSpec document set from a real project into this prototype and compare answers from:

1. static BSpec parser queries,
2. Graphiti hybrid search,
3. graph-distance reranked Graphiti search.

That comparison will show whether Graphiti adds enough value to justify a production service boundary.
