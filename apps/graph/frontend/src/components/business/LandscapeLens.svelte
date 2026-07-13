<script lang="ts">
  import type {
    Neighborhood,
    PathResult,
    RecordView,
  } from '../../../bindings/github.com/a3tai/bspec/apps/graph';
  import { buildGraphScene, titleCase } from '../../lib/business';
  import Badge from '../ui/badge/badge.svelte';
  import Button from '../ui/button/button.svelte';
  import Empty from '../ui/empty/empty.svelte';

  type Props = {
    neighborhood: Neighborhood | null;
    records: RecordView[];
    selectedRecordID: string;
    loading: boolean;
    errorCode: string;
    errorMessage: string;
    projectionPending: boolean;
    pathResult: PathResult | null;
    pathTargetID: string;
    pathTargetQuery: string;
    pathCandidates: RecordView[];
    pathTargetsLoading: boolean;
    pathLoading: boolean;
    pathErrorCode: string;
    pathErrorMessage: string;
    activeKinds: string[];
    onPathTargetChange: (recordID: string) => void;
    onPathTargetSearch: (event: Event) => void;
    onToggleKind: (kind: string) => void;
    onFindPath: () => void;
    onRetry: () => void;
    onSelectRecord: (recordID: string) => void;
  };

  let {
    neighborhood,
    records,
    selectedRecordID,
    loading,
    errorCode,
    errorMessage,
    projectionPending,
    pathResult,
    pathTargetID,
    pathTargetQuery,
    pathCandidates,
    pathTargetsLoading,
    pathLoading,
    pathErrorCode,
    pathErrorMessage,
    activeKinds,
    onPathTargetChange,
    onPathTargetSearch,
    onToggleKind,
    onFindPath,
    onRetry,
    onSelectRecord,
  }: Props = $props();

  const relationshipKinds = ['depends_on', 'enables', 'conflicts_with', 'related', 'supersedes'];
  const scene = $derived(buildGraphScene(neighborhood, records));
  const selectedRecord = $derived(records.find((record) => record.record.id === selectedRecordID) ?? null);
  const pathTargets = $derived.by(() => {
    const byID = new Map(pathCandidates.map((record) => [record.record.id, record]));
    const selectedTarget = records.find((record) => record.record.id === pathTargetID);
    if (selectedTarget) byID.set(selectedTarget.record.id, selectedTarget);
    byID.delete(selectedRecordID);
    return [...byID.values()];
  });

  function pathEdge(index: number) {
    return pathResult?.edges[index];
  }

  function pathDirection(index: number): string {
    const edge = pathEdge(index);
    const node = pathResult?.nodes[index];
    if (!edge || !node) return '→';
    return edge.source_id === node.id ? '→' : '←';
  }
</script>

<section class="landscape-lens">
  <header class="lens-heading landscape-heading">
    <div>
      <span class="eyebrow">Focused landscape</span>
      <h2>{selectedRecord?.revision.title || 'Business landscape'}</h2>
      <p>
        {#if neighborhood}
          Depth {neighborhood.depth} · {scene.nodes.length} visible nodes · {scene.edges.length} explicit relationships
        {:else}
          Select a record to open its bounded projected neighborhood.
        {/if}
      </p>
    </div>
    <div class="relationship-filters" aria-label="Relationship filters">
      {#each relationshipKinds as kind}
        <button
          type="button"
          class:active={activeKinds.includes(kind)}
          aria-pressed={activeKinds.includes(kind)}
          onclick={() => onToggleKind(kind)}
        >
          <span class={`relationship-glyph relationship-${kind}`} aria-hidden="true"></span>
          {titleCase(kind)}
        </button>
      {/each}
    </div>
  </header>

  {#if projectionPending}
    <div class="state-banner state-banner-warn">
      Canonical records are current, but the graph projection is catching up. This view may temporarily omit recent changes.
    </div>
  {/if}

  {#if errorCode === 'BUSINESS_PROJECTION_UNAVAILABLE'}
    <div class="graph-state-panel graph-state-danger">
      <div class="state-icon">!</div>
      <div>
        <h3>Graph projection unavailable</h3>
        <p>{errorMessage || 'Neo4j cannot serve this graph read. Canonical records remain available.'}</p>
      </div>
      <Button variant="outline" onclick={onRetry}>Retry</Button>
    </div>
  {:else if errorCode === 'BUSINESS_NOT_FOUND' && projectionPending}
    <div class="graph-state-panel graph-state-warn">
      <div class="state-icon">↻</div>
      <div>
        <h3>Saved and projecting</h3>
        <p>The selected record is canonical, but its graph node has not reached the projection yet.</p>
      </div>
      <Button variant="outline" onclick={onRetry}>Check again</Button>
    </div>
  {:else if errorMessage}
    <div class="graph-state-panel graph-state-danger">
      <div class="state-icon">!</div>
      <div>
        <h3>Landscape could not load</h3>
        <p>{errorMessage}</p>
      </div>
      <Button variant="outline" onclick={onRetry}>Retry</Button>
    </div>
  {:else if loading}
    <div class="graph-loading" aria-live="polite">
      <span class="spinner"></span>
      Loading projected neighborhood
    </div>
  {:else if scene.nodes.length === 0}
    <div class="lens-empty">
      <Empty
        title={selectedRecordID ? 'No projected neighborhood yet' : 'Choose an orientation point'}
        description={selectedRecordID ? 'The record may still be projecting, or it may have no projected relationships.' : 'Select a record from the index to explore its graph context.'}
      />
    </div>
  {:else}
    {#if neighborhood?.truncated}
      <div class="state-banner state-banner-neutral">
        This neighborhood reached its node or edge limit. Narrow the relationship filters or orient from another record.
      </div>
    {/if}

    <div class="graph-canvas-wrap">
      <svg
        class="business-graph"
        viewBox={`0 0 ${scene.width} ${scene.height}`}
        style={`height: ${scene.height}px`}
        role="group"
        aria-label="Interactive business record relationship landscape"
      >
        <defs>
          <marker id="edge-arrow" viewBox="0 0 10 10" refX="8" refY="5" markerWidth="7" markerHeight="7" orient="auto-start-reverse">
            <path d="M 0 0 L 10 5 L 0 10 z"></path>
          </marker>
        </defs>

        {#each scene.domains as domain, index (domain)}
          {@const domainNodes = scene.nodes.filter((node) => node.domain === domain)}
          {#if domainNodes.length > 0}
            {@const x = Math.min(...domainNodes.map((node) => node.x)) - 56}
            {@const y = Math.min(...domainNodes.map((node) => node.y)) - 48}
            <text class="domain-label" x={Math.max(18, x)} y={Math.max(22, y)}>{domain}</text>
          {/if}
        {/each}

        {#each scene.edges as edge (edge.id)}
          <g class={`graph-edge graph-edge-${edge.kind}`} class:graph-edge-unresolved={edge.unresolved}>
            <path
              d={edge.path}
              marker-end="url(#edge-arrow)"
            ></path>
            <text x={edge.labelX} y={edge.labelY}>
              {titleCase(edge.kind)}
            </text>
          </g>
        {/each}

        {#each scene.nodes as node (node.id)}
          {#if node.unresolved}
            <g
              class={`graph-node graph-tone-${node.tone} graph-node-unresolved`}
              aria-label={`Unresolved reference ${node.key}`}
            >
              <path d={`M ${node.x} ${node.y - 25} L ${node.x + 38} ${node.y} L ${node.x} ${node.y + 25} L ${node.x - 38} ${node.y} Z`}></path>
              <text class="node-type" x={node.x} y={node.y - 6}>{node.typeCode}</text>
              <text class="node-key" x={node.x} y={node.y + 12}>{node.key.length > 18 ? `${node.key.slice(0, 16)}…` : node.key}</text>
            </g>
          {:else}
            <a
              href={`#record-${node.id}`}
              class={`graph-node graph-tone-${node.tone}`}
              class:graph-node-selected={node.id === selectedRecordID}
              aria-current={node.id === selectedRecordID ? 'true' : undefined}
              aria-label={`Select ${node.typeCode} ${node.title}`}
              onclick={(event) => {
                event.preventDefault();
                onSelectRecord(node.id);
              }}
            >
              <rect x={node.x - 54} y={node.y - 27} width="108" height="54" rx="12"></rect>
              <text class="node-type" x={node.x} y={node.y - 6}>{node.typeCode}</text>
              <text class="node-key" x={node.x} y={node.y + 12}>{node.key.length > 18 ? `${node.key.slice(0, 16)}…` : node.key}</text>
            </a>
          {/if}
        {/each}
      </svg>

      <div class="graph-legend" aria-label="Graph legend">
        <span><i class="legend-record"></i> Canonical record</span>
        <span><i class="legend-reference"></i> Unresolved reference</span>
        <span><i class="legend-conflict"></i> Conflict</span>
        <span><i class="legend-related"></i> Related</span>
      </div>
    </div>
  {/if}

  <section class="path-explainer">
    <header class="section-heading">
      <div>
        <span class="eyebrow">Impact analysis</span>
        <h3>Explain a path</h3>
      </div>
      <Badge variant="outline">Shortest bounded path</Badge>
    </header>

    <div class="path-controls">
      <div class="path-anchor">
        <span>From</span>
        <strong>{selectedRecord?.revision.title || 'Select a record'}</strong>
      </div>
      <span class="path-control-arrow" aria-hidden="true">→</span>
      <label>
        <span>To</span>
        <input value={pathTargetQuery} oninput={onPathTargetSearch} type="search" aria-label="Search path destinations" placeholder="Search destinations" />
        <select value={pathTargetID} aria-label="Choose path destination" onchange={(event) => onPathTargetChange(event.currentTarget.value)}>
          <option value="">{pathTargetsLoading ? 'Searching destinations…' : 'Choose a destination'}</option>
          {#each pathTargets as target (target.record.id)}
            <option value={target.record.id}>{target.record.type_code} · {target.revision.title}</option>
          {/each}
        </select>
      </label>
      <Button variant="outline" onclick={onFindPath} disabled={!selectedRecordID || !pathTargetID || pathLoading}>
        {pathLoading ? 'Tracing…' : 'Explain'}
      </Button>
    </div>

    {#if pathErrorCode === 'BUSINESS_PROJECTION_UNAVAILABLE'}
      <div class="state-banner state-banner-danger">
        Path projection is unavailable. Canonical records are still readable.
      </div>
    {:else if pathErrorCode === 'BUSINESS_NOT_FOUND'}
      <div class="state-banner state-banner-neutral">
        A path endpoint is not projected yet. Wait for projection state to become current and try again.
      </div>
    {:else if pathErrorMessage}
      <div class="state-banner state-banner-danger">{pathErrorMessage}</div>
    {:else if pathResult}
      {#if pathResult.found}
        <ol class="path-result" aria-label="Relationship path">
          {#each pathResult.nodes as node, index (node.id)}
            <li>
              <button type="button" onclick={() => onSelectRecord(node.id)}>
                <span>{node.type_code}</span>
                <strong>{node.title}</strong>
                <small>{node.key}</small>
              </button>
              {#if index < pathResult.nodes.length - 1}
                <div class="path-connector">
                  <span>{titleCase(pathEdge(index)?.kind || 'related')}</span>
                  <b aria-label={pathDirection(index) === '→' ? 'forward direction' : 'reverse direction'}>{pathDirection(index)}</b>
                </div>
              {/if}
            </li>
          {/each}
        </ol>
      {:else}
        <div class="state-banner state-banner-neutral">
          Both records exist, but no path was found within the current bound and relationship filters.
        </div>
      {/if}
    {/if}
  </section>
</section>
