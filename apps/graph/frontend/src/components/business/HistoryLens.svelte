<script lang="ts">
  import type {
    RecordRevision,
    RecordView,
  } from '../../../bindings/github.com/a3tai/bspec/apps/graph';
  import { dataAsJSON, formatDate, titleCase } from '../../lib/business';
  import { renderMarkdown } from '../../utils/markdown';
  import Badge from '../ui/badge/badge.svelte';
  import Button from '../ui/button/button.svelte';
  import Empty from '../ui/empty/empty.svelte';

  type Props = {
    revisions: RecordRevision[];
    selectedRevisionID: string;
    snapshot: RecordView | null;
    previousSnapshot: RecordView | null;
    loading: boolean;
    nextCursor: string;
    onSelectRevision: (revisionID: string) => void;
    onLoadMore: () => void;
  };

  let {
    revisions,
    selectedRevisionID,
    snapshot,
    previousSnapshot,
    loading,
    nextCursor,
    onSelectRevision,
    onLoadMore,
  }: Props = $props();

  const selectedHTML = $derived(renderMarkdown(snapshot?.revision.narrative_markdown || 'No narrative in this revision.'));
  const previousHTML = $derived(renderMarkdown(previousSnapshot?.revision.narrative_markdown || 'No previous narrative.'));
  const selectedData = $derived(dataAsJSON(snapshot?.revision.data_json));
  const previousData = $derived(dataAsJSON(previousSnapshot?.revision.data_json));

  const metadataComparison = $derived.by(() => {
    if (!snapshot) return [];
    const previous = previousSnapshot;
    return [
      { label: 'Type', before: previous?.record.type_code ?? '', after: snapshot.record.type_code },
      { label: 'Status', before: previous?.revision.status ?? '', after: snapshot.revision.status ?? '' },
      { label: 'Domain', before: previous?.revision.domain ?? '', after: snapshot.revision.domain ?? '' },
      { label: 'Visibility', before: previous?.revision.visibility ?? '', after: snapshot.revision.visibility ?? '' },
      { label: 'Owner', before: previous?.revision.owner_ref ?? '', after: snapshot.revision.owner_ref ?? '' },
      { label: 'Schema', before: previous?.revision.schema_version ?? '', after: snapshot.revision.schema_version ?? '' },
      { label: 'Source version', before: previous?.revision.source_version ?? '', after: snapshot.revision.source_version ?? '' },
      { label: 'Source URI', before: previous?.revision.source_uri ?? '', after: snapshot.revision.source_uri ?? '' },
      { label: 'Source media', before: previous?.revision.source_media_type ?? '', after: snapshot.revision.source_media_type ?? '' },
      { label: 'Effective at', before: previous?.revision.effective_at ?? '', after: snapshot.revision.effective_at ?? '' },
    ];
  });

  const relationshipChanges = $derived.by(() => {
    if (!snapshot || !previousSnapshot) return [];
    const identity = (relationship: RecordView['relationships'][number]) => `${relationship.kind}:${relationship.target_key}`;
    const value = (relationship: RecordView['relationships'][number]) => JSON.stringify({
      targetRecordID: relationship.target_record_id ?? '',
      resolution: relationship.resolution,
      strength: relationship.strength ?? '',
      provenanceJSON: relationship.provenance_json ?? '',
    });
    const before = new Map(previousSnapshot.relationships.map((relationship) => [identity(relationship), relationship]));
    const after = new Map(snapshot.relationships.map((relationship) => [identity(relationship), relationship]));
    const changes: Array<{ key: string; state: 'added' | 'removed' | 'changed'; label: string }> = [];
    for (const [key, relationship] of after) {
      const prior = before.get(key);
      if (!prior) changes.push({ key, state: 'added', label: `${titleCase(relationship.kind)} ${relationship.target_key}` });
      else if (value(prior) !== value(relationship)) changes.push({ key, state: 'changed', label: `${titleCase(relationship.kind)} ${relationship.target_key}` });
    }
    for (const [key, relationship] of before) {
      if (!after.has(key)) changes.push({ key, state: 'removed', label: `${titleCase(relationship.kind)} ${relationship.target_key}` });
    }
    return changes.sort((left, right) => left.key.localeCompare(right.key));
  });

  function relationshipSnapshot(view: RecordView): string {
    return JSON.stringify(
      view.relationships
        .map((relationship) => ({
          kind: relationship.kind,
          targetKey: relationship.target_key,
          targetRecordID: relationship.target_record_id ?? '',
          resolution: relationship.resolution,
          strength: relationship.strength ?? '',
          provenanceJSON: relationship.provenance_json ?? '',
        }))
        .sort((left, right) => JSON.stringify(left).localeCompare(JSON.stringify(right))),
    );
  }

  const changeSummary = $derived.by(() => {
    if (!snapshot) return [];
    if (!previousSnapshot) {
      return snapshot.revision.revision === 1 ? ['Initial revision'] : ['Previous revision not loaded'];
    }
    const changes: string[] = [];
    if (snapshot.revision.content_hash !== previousSnapshot.revision.content_hash) changes.push('Content changed');
    if (relationshipSnapshot(snapshot) !== relationshipSnapshot(previousSnapshot)) changes.push('Relationships changed');
    if (snapshot.record.type_code !== previousSnapshot.record.type_code) changes.push('Type changed');
    if (snapshot.revision.status !== previousSnapshot.revision.status) changes.push('Status changed');
    if (snapshot.revision.domain !== previousSnapshot.revision.domain) changes.push('Domain changed');
    return changes.length ? changes : ['No semantic change'];
  });
</script>

<section class="history-lens">
  <header class="lens-heading">
    <div>
      <span class="eyebrow">Immutable trust surface</span>
      <h2>Revision history</h2>
      <p>Each entry reconstructs the complete record envelope and relationship snapshot as it existed then.</p>
    </div>
    <Badge variant="outline">Read only</Badge>
  </header>

  {#if revisions.length === 0 && !loading}
    <div class="lens-empty">
      <Empty title="No revision history" description="Select a record to inspect its immutable revision timeline." />
    </div>
  {:else}
    <div class="history-layout">
      <aside class="revision-timeline" aria-label="Record revisions">
        {#each revisions as revision (revision.id)}
          <button
            type="button"
            class:active={selectedRevisionID === revision.id}
            aria-current={selectedRevisionID === revision.id ? 'true' : undefined}
            onclick={() => onSelectRevision(revision.id)}
          >
            <span class="timeline-marker" aria-hidden="true"></span>
            <span>
              <strong>Revision {revision.revision}</strong>
              <small>{formatDate(revision.created_at, true)}</small>
              <em>{revision.created_by}</em>
            </span>
            <Badge variant="neutral" size="sm">v{revision.record_version}</Badge>
          </button>
        {/each}
        {#if nextCursor}
          <Button variant="ghost" size="sm" onclick={onLoadMore} disabled={loading}>Load older revisions</Button>
        {/if}
      </aside>

      <section class="history-detail">
        {#if loading && !snapshot}
          <div class="graph-loading"><span class="spinner"></span>Loading historical snapshot</div>
        {:else if snapshot}
          <header class="historical-banner">
            <div>
              <span class="eyebrow">Historical snapshot · read only</span>
              <h3>{snapshot.revision.title}</h3>
              <p>{snapshot.record.key} · {snapshot.record.type_code} · revision {snapshot.revision.revision}</p>
            </div>
            <div class="change-chips">
              {#each changeSummary as change}
                <Badge variant={change === 'No semantic change' ? 'neutral' : 'accent'}>{change}</Badge>
              {/each}
            </div>
          </header>

          <div class="history-metadata history-metadata-comparison">
            {#each metadataComparison as item (item.label)}
              <div class:changed={previousSnapshot !== null && item.before !== item.after}>
                <span>{item.label}</span>
                {#if previousSnapshot && item.before !== item.after}<del>{item.before || 'Not set'}</del>{/if}
                <strong>{item.after || 'Not set'}</strong>
              </div>
            {/each}
          </div>

          <section class="history-change-section">
            <div class="section-heading">
              <div>
                <span class="eyebrow">Content comparison</span>
                <h3>Narrative</h3>
              </div>
            </div>
            <div class="narrative-comparison" class:single={previousSnapshot === null}>
              {#if previousSnapshot}
                <article>
                  <span class="comparison-label">Revision {previousSnapshot.revision.revision}</span>
                  <div class="markdown-reading compact-reading">{@html previousHTML}</div>
                </article>
              {/if}
              <article>
                <span class="comparison-label comparison-current">Revision {snapshot.revision.revision}</span>
                <div class="markdown-reading compact-reading">{@html selectedHTML}</div>
              </article>
            </div>
          </section>

          <section class="history-change-section">
            <div class="section-heading">
              <div>
                <span class="eyebrow">Graph comparison</span>
                <h3>Relationship snapshot</h3>
              </div>
              <Badge variant="neutral">{snapshot.relationships.length}</Badge>
            </div>
            {#if previousSnapshot}
              <div class="relationship-change-list">
                {#each relationshipChanges as change (change.key)}
                  <div>
                    <Badge variant={change.state === 'removed' ? 'danger' : change.state === 'added' ? 'success' : 'warn'} size="sm">{titleCase(change.state)}</Badge>
                    <span>{change.label}</span>
                  </div>
                {:else}
                  <p class="quiet-empty">No relationship changes from the previous revision.</p>
                {/each}
              </div>
            {/if}
            <div class="history-relationships">
              {#each snapshot.relationships as relationship (relationship.id)}
                <div>
                  <span class={`relationship-glyph relationship-${relationship.kind}`} aria-hidden="true"></span>
                  <strong>{titleCase(relationship.kind)}</strong>
                  <span>{relationship.target_key}</span>
                  <Badge variant={relationship.resolution === 'resolved' ? 'success' : 'warn'} size="sm">
                    {relationship.resolution}
                  </Badge>
                </div>
              {:else}
                <p class="quiet-empty">No outgoing relationships in this historical revision.</p>
              {/each}
            </div>
          </section>

          <details class="structured-data">
            <summary>Structured data comparison</summary>
            <div class="data-comparison" class:single={previousSnapshot === null}>
              {#if previousSnapshot}
                <div><span>Revision {previousSnapshot.revision.revision}</span><pre>{previousData}</pre></div>
              {/if}
              <div><span>Revision {snapshot.revision.revision}</span><pre>{selectedData}</pre></div>
            </div>
          </details>
        {/if}
      </section>
    </div>
  {/if}
</section>
