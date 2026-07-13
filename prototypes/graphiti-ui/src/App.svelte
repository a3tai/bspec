<script lang="ts">
  import graphData from './data/graph.json';
  import Badge from './components/ui/badge/badge.svelte';
  import Button from './components/ui/button/button.svelte';
  import Card from './components/ui/card/card.svelte';
  import Chip from './components/ui/chip/chip.svelte';
  import Empty from './components/ui/empty/empty.svelte';
  import Inspector from './components/ui/inspector/inspector.svelte';
  import Lamp from './components/ui/lamp/lamp.svelte';
  import NativeSelect from './components/ui/native-select/native-select.svelte';
  import SearchInput from './components/ui/search-input/search-input.svelte';
  import SegmentedControl from './components/ui/segmented-control/segmented-control.svelte';
  import SourceList from './components/ui/source-list/source-list.svelte';
  import Sparkline from './components/ui/sparkline/sparkline.svelte';
  import Stat from './components/ui/stat/stat.svelte';
  import Table from './components/ui/table/table.svelte';
  import Toolbar from './components/ui/toolbar/toolbar.svelte';

  type Domain = {
    id: string;
    slug: string;
    name: string;
    description: string;
    documentCount: number;
    outgoingEdgeCount: number;
    incomingEdgeCount: number;
  };

  type RelationshipRef = {
    code: string;
    pattern: string;
  };

  type DocumentNode = {
    id: string;
    kind: string;
    documentId: string;
    code: string;
    name: string;
    domainSlug: string;
    domainName: string;
    status: string;
    version: string;
    lastUpdated: string;
    owner: string;
    sourcePath: string;
    abstract: string;
    purpose: string;
    relationshipGuidance: string;
    relationships: Record<string, RelationshipRef[]>;
    outgoingCount: number;
  };

  type Edge = {
    id: string;
    source: string;
    sourceCode: string;
    sourceDocumentId: string;
    sourceName: string;
    sourceDomainSlug: string;
    target: string;
    targetCode: string;
    targetDocumentId: string;
    targetName: string;
    targetDomainSlug: string;
    targetPattern: string;
    field: string;
    label: string;
    fact: string;
  };

  type Mode = 'graph' | 'facts' | 'episodes';

  const data = graphData as {
    title: string;
    description: string;
    sourceKind?: 'catalog' | 'project';
    sourceRoot: string;
    version: string;
    groupId: string;
    nodeLabelSingular?: string;
    nodeLabelPlural?: string;
    domainHeading?: string;
    graphitiCoreEpisodes: number;
    graphitiRelationshipRuleEpisodes: number;
    domains: Domain[];
    documents: DocumentNode[];
    edges: Edge[];
    relationCounts: { field: string; label: string; count: number }[];
    domainSeries: number[];
  };

  const sourceKind = data.sourceKind ?? 'catalog';
  const nodeLabelSingular = data.nodeLabelSingular ?? (sourceKind === 'project' ? 'Document' : 'Document type');
  const nodeLabelPlural = data.nodeLabelPlural ?? (sourceKind === 'project' ? 'Documents' : 'Document types');
  const domainHeading = data.domainHeading ?? (sourceKind === 'project' ? 'Project Domains' : 'BSpec Domains');
  const workbenchTitle = data.title ?? 'BSpec Graphiti Workbench';
  const workbenchDescription =
    data.description ?? 'Interactive Mittsu prototype for exploring BSpec as Graphiti-ready graph data.';
  const sourceRoot = data.sourceRoot ?? '../../spec/v1';

  const relationOptions = [
    { value: 'all', label: 'All relationships' },
    ...data.relationCounts.map((item) => ({
      value: item.field,
      label: `${item.label} (${item.count})`,
    })),
  ];

  const modeOptions: { value: Mode; label: string }[] = [
    { value: 'graph', label: 'Graph' },
    { value: 'facts', label: 'Facts' },
    { value: 'episodes', label: 'Episodes' },
  ];

  const domainItems = [
    { type: 'heading' as const, label: domainHeading },
    { value: 'all', label: `All ${nodeLabelPlural.toLowerCase()}`, count: data.documents.length },
    ...data.domains.map((domain) => ({
      value: domain.slug,
      label: domain.name,
      count: domain.documentCount,
    })),
  ];

  let selectedDomain = $state('all');
  let relationFilter = $state('all');
  let mode = $state<Mode>('graph');
  let query = $state('');
  let selectedDocumentId = $state(data.documents[0]?.id ?? '');

  const normalizedQuery = $derived(query.trim().toLowerCase());

  const filteredDocuments = $derived.by(() => {
    return data.documents.filter((document) => {
      const inDomain = selectedDomain === 'all' || document.domainSlug === selectedDomain;
      const matchesQuery =
        normalizedQuery.length === 0 ||
        [
          document.code,
          document.documentId,
          document.name,
          document.domainName,
          document.abstract,
          document.purpose,
          document.sourcePath,
        ]
          .join(' ')
          .toLowerCase()
          .includes(normalizedQuery);

      return inDomain && matchesQuery;
    });
  });

  const visibleDocumentIds = $derived(new Set(filteredDocuments.map((document) => document.id)));

  const filteredEdges = $derived.by(() => {
    return data.edges.filter((edge) => {
      const sourceVisible = visibleDocumentIds.has(edge.source);
      const targetVisible = visibleDocumentIds.has(edge.target);
      const relationMatches = relationFilter === 'all' || edge.field === relationFilter;
      const queryMatches =
        normalizedQuery.length === 0 ||
        [edge.fact, edge.sourceName, edge.sourceCode, edge.targetName, edge.targetCode, edge.label]
          .join(' ')
          .toLowerCase()
          .includes(normalizedQuery);

      if (mode === 'facts') {
        return relationMatches && queryMatches && (selectedDomain === 'all' || sourceVisible || targetVisible);
      }

      return relationMatches && sourceVisible && targetVisible && queryMatches;
    });
  });

  const selectedDocument = $derived(
    data.documents.find((document) => document.id === selectedDocumentId) ?? filteredDocuments[0] ?? data.documents[0],
  );

  const selectedDocumentEdges = $derived(
    data.edges.filter((edge) => edge.source === selectedDocument?.id || edge.target === selectedDocument?.id),
  );

  const graphNodes = $derived(filteredDocuments.slice(0, 42));
  const graphNodeIds = $derived(new Set(graphNodes.map((node) => node.id)));
  const graphEdges = $derived(
    filteredEdges.filter((edge) => graphNodeIds.has(edge.source) && graphNodeIds.has(edge.target)).slice(0, 120),
  );

  const graphPositions = $derived.by(() => {
    const centerX = 420;
    const centerY = 260;
    const radiusX = 310;
    const radiusY = 185;
    const byId = new Map<string, { x: number; y: number }>();

    graphNodes.forEach((node, index) => {
      const angle = (index / Math.max(graphNodes.length, 1)) * Math.PI * 2 - Math.PI / 2;
      byId.set(node.id, {
        x: centerX + Math.cos(angle) * radiusX,
        y: centerY + Math.sin(angle) * radiusY,
      });
    });

    return byId;
  });

  const relationRows = $derived(
    filteredEdges.slice(0, 250).map((edge) => ({
      source: edge.sourceCode,
      relation: edge.label,
      target: edge.targetPattern,
      domain: edge.sourceDomainSlug,
      fact: edge.fact,
    })),
  );

  const episodeRows = $derived(
    filteredDocuments.slice(0, 250).map((document) => ({
      episode: `${document.documentId || document.code}: ${document.name}`,
      domain: document.domainName,
      relationships: document.outgoingCount,
      source: document.sourcePath,
      updated: document.lastUpdated,
    })),
  );

  const searchResults = $derived.by(() => {
    if (!normalizedQuery) {
      return selectedDocumentEdges.slice(0, 12).map((edge) => ({
        id: edge.id,
        label: edge.label,
        title: edge.fact,
        detail: `${edge.sourceName} -> ${edge.targetName}`,
      }));
    }

    const documentHits = data.documents
      .filter((document) =>
        [document.code, document.name, document.domainName, document.purpose, document.abstract]
          .join(' ')
          .toLowerCase()
          .includes(normalizedQuery),
      )
      .slice(0, 8)
      .map((document) => ({
        id: document.id,
        label: sourceKind === 'catalog' ? 'DOCUMENT_TYPE' : 'DOCUMENT',
        title: `${document.documentId || document.code}: ${document.name}`,
        detail: document.purpose || document.abstract || document.sourcePath,
      }));

    const edgeHits = data.edges
      .filter((edge) => edge.fact.toLowerCase().includes(normalizedQuery))
      .slice(0, 8)
      .map((edge) => ({
        id: edge.id,
        label: edge.label,
        title: edge.fact,
        detail: `${edge.sourceName} -> ${edge.targetName}`,
      }));

    return [...documentHits, ...edgeHits].slice(0, 12);
  });

  const selectedEpisode = $derived.by(() => {
    if (!selectedDocument) return '{}';

    const isCatalog = selectedDocument.kind === 'bspec.document_type';
    const episodeName = isCatalog
      ? `BSpec ${data.version} document type ${selectedDocument.code}: ${selectedDocument.name}`
      : `BSpec document ${selectedDocument.documentId}: ${selectedDocument.name}`;

    return JSON.stringify(
      {
        group_id: data.groupId,
        name: episodeName,
        source: 'json',
        source_description: `${nodeLabelSingular}: ${selectedDocument.sourcePath}`,
        episode_body: {
          kind: selectedDocument.kind,
          standard: 'BSpec',
          standard_version: isCatalog ? data.version : undefined,
          id: isCatalog ? undefined : selectedDocument.documentId,
          code: selectedDocument.code,
          name: selectedDocument.name,
          title: selectedDocument.name,
          domain: {
            slug: selectedDocument.domainSlug,
            name: selectedDocument.domainName,
          },
          owner: selectedDocument.owner,
          status: selectedDocument.status,
          relationships: selectedDocument.relationships,
          purpose: selectedDocument.purpose,
        },
      },
      null,
      2,
    );
  });

  const inspectorSections = $derived([
    {
      title: nodeLabelSingular,
      rows: [
        { label: sourceKind === 'project' ? 'ID' : 'Code', value: selectedDocument?.documentId ?? selectedDocument?.code ?? '' },
        { label: 'Type', value: selectedDocument?.code ?? '' },
        { label: 'Name', value: selectedDocument?.name ?? '' },
        { label: 'Domain', value: selectedDocument?.domainName ?? '' },
        { label: 'Status', value: selectedDocument?.status ?? '' },
        { label: 'Owner', value: selectedDocument?.owner ?? '' },
        { label: 'Updated', value: selectedDocument?.lastUpdated ?? '' },
      ],
    },
    {
      title: 'Graph',
      rows: [
        { label: 'Edges', value: String(selectedDocumentEdges.length) },
        { label: 'Source', value: selectedDocument?.sourcePath ?? '' },
      ],
    },
  ]);

  function selectDocument(documentId: string) {
    selectedDocumentId = documentId;
  }

  function selectFromResult(resultId: string) {
    const document = data.documents.find((item) => item.id === resultId);
    const edge = data.edges.find((item) => item.id === resultId);
    selectedDocumentId = document?.id ?? edge?.source ?? selectedDocumentId;
  }

  function resetFilters() {
    selectedDomain = 'all';
    relationFilter = 'all';
    query = '';
  }
</script>

<svelte:head>
  <title>{workbenchTitle}</title>
  <meta
    name="description"
    content={workbenchDescription}
  />
</svelte:head>

<main class="workbench">
  <header class="topbar">
    <div class="brand">
      <div class="mark">B</div>
      <div>
        <h1>{workbenchTitle}</h1>
        <p>v{data.version} · {data.groupId} · {sourceRoot}</p>
      </div>
    </div>
    <div class="runtime">
      <Lamp state="running" size={10} />
      <span>{sourceKind === 'project' ? 'BSpec project loaded' : 'BSpec catalog loaded'}</span>
      <Badge variant="outline">MIT</Badge>
      <Badge variant="accent">Mittsu</Badge>
    </div>
  </header>

  <section class="stats">
    <div class="uin-stat-row">
      <Stat label={nodeLabelPlural} value={data.documents.length} tone="accent" size="lg" />
      <Stat label="Domains" value={data.domains.length} size="lg" />
      <Stat label="Relationship facts" value={data.edges.length} size="lg" />
      <Stat label="Core episodes" value={data.graphitiCoreEpisodes} size="lg" />
      <Stat label="Rule episodes" value={data.graphitiRelationshipRuleEpisodes} size="lg" />
    </div>
    <div class="spark-panel" aria-label="Domain graph density">
      <span>Domain density</span>
      <Sparkline values={data.domainSeries} width={180} height={34} area tone="success" />
    </div>
  </section>

  <section class="layout">
    <aside class="sidebar">
      <SourceList
        items={domainItems}
        bind:value={selectedDomain}
        ariaLabel="BSpec domains"
      />
      <div class="sidebar-actions">
        <Button variant="outline" size="sm" onclick={resetFilters}>Reset</Button>
      </div>
    </aside>

    <section class="main-panel">
      <Toolbar density="roomy" class="workbench-toolbar">
        <SearchInput
          bind:value={query}
          placeholder={`Search ${nodeLabelPlural.toLowerCase()}, relationships, or Graphiti facts`}
          shortcutLabel="/"
        />
        <SegmentedControl
          bind:value={mode}
          options={modeOptions}
          ariaLabel="Workbench mode"
          size="sm"
        />
        <NativeSelect
          bind:value={relationFilter}
          options={relationOptions}
          size="sm"
          aria-label="Relationship filter"
        />
      </Toolbar>

      {#if mode === 'graph'}
        <Card padding={false} class="graph-card">
          <div class="graph-head">
            <div>
              <h2>{nodeLabelSingular} Relationship Graph</h2>
              <p>{graphNodes.length} visible {nodeLabelPlural.toLowerCase()} · {graphEdges.length} visible facts</p>
            </div>
            <div class="chips">
              {#each data.relationCounts.filter((item) => item.count > 0) as item}
                <Chip
                  active={relationFilter === item.field}
                  count={item.count}
                  onClick={() => relationFilter = relationFilter === item.field ? 'all' : item.field}
                >
                  {item.label}
                </Chip>
              {/each}
            </div>
          </div>

          {#if graphNodes.length === 0}
            <Empty title="No graph nodes" description="Try a broader query or reset the filters." />
          {:else}
            <svg class="graph" viewBox="0 0 840 520" role="img" aria-label="BSpec document relationship graph">
              {#each graphEdges as edge (edge.id)}
                {@const source = graphPositions.get(edge.source)}
                {@const target = graphPositions.get(edge.target)}
                {#if source && target}
                  <line
                    x1={source.x}
                    y1={source.y}
                    x2={target.x}
                    y2={target.y}
                    class:edge-enable={edge.field === 'enables'}
                    class:edge-depends={edge.field === 'depends_on'}
                  />
                {/if}
              {/each}

              {#each graphNodes as node (node.id)}
                {@const point = graphPositions.get(node.id)}
                {#if point}
                  <g
                    role="button"
                    tabindex="0"
                    aria-label={`Select ${node.documentId || node.code}: ${node.name}`}
                    class:selected-node={selectedDocument?.id === node.id}
                    onclick={() => selectDocument(node.id)}
                    onkeydown={(event) => {
                      if (event.key === 'Enter' || event.key === ' ') selectDocument(node.id);
                    }}
                  >
                    <circle cx={point.x} cy={point.y} r={node.outgoingCount > 8 ? 25 : 20} />
                    <text x={point.x} y={point.y + 4}>{node.code}</text>
                  </g>
                {/if}
              {/each}
            </svg>
          {/if}
        </Card>
      {:else if mode === 'facts'}
        <Card padding={false} class="table-card">
          <Table
            columns={[
              { key: 'source', label: 'Source', sortable: true, width: '90px' },
              { key: 'relation', label: 'Relation', sortable: true, width: '150px' },
              { key: 'target', label: 'Target', sortable: true, width: '120px' },
              { key: 'fact', label: 'Graphiti fact' },
            ]}
            rows={relationRows}
            striped
            density="compact"
            emptyLabel="No relationship facts match the current filters"
          />
        </Card>
      {:else}
        <Card padding={false} class="table-card">
          <Table
            columns={[
              { key: 'episode', label: 'Episode', sortable: true },
              { key: 'domain', label: 'Domain', sortable: true },
              { key: 'relationships', label: 'Edges', sortable: true, align: 'right', width: '80px' },
              { key: 'updated', label: 'Updated', sortable: true, width: '120px' },
              { key: 'source', label: 'Source path' },
            ]}
            rows={episodeRows}
            striped
            density="compact"
            emptyLabel="No episodes match the current filters"
          />
        </Card>
      {/if}

      <section class="search-results">
        <div class="section-title">
          <h2>Local Graphiti Search Preview</h2>
          <Badge variant="neutral">{searchResults.length}</Badge>
        </div>
        <div class="result-list">
          {#each searchResults as result (result.id)}
            <button class="result-row" type="button" onclick={() => selectFromResult(result.id)}>
              <span class="result-label">{result.label}</span>
              <span class="result-title">{result.title}</span>
              <span class="result-detail">{result.detail}</span>
            </button>
          {:else}
            <Empty title="No local results" description="Try searching for PRD, strategy, risk, or enables." />
          {/each}
        </div>
      </section>
    </section>

    <Inspector title="Selection" sections={inspectorSections} width="360px" class="details">
      <section class="inspector-block">
        <p class="inspector-block-title">Purpose</p>
        <p>{selectedDocument?.purpose || selectedDocument?.abstract || 'No document body extracted.'}</p>
      </section>

      <section class="inspector-block">
        <p class="inspector-block-title">Graphiti Episode</p>
        <pre>{selectedEpisode}</pre>
      </section>
    </Inspector>
  </section>
</main>
