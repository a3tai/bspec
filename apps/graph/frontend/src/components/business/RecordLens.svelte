<script lang="ts">
  import type { RecordView } from '../../../bindings/github.com/a3tai/bspec/apps/graph';
  import type { IncomingRelationshipView } from '../../lib/business';
  import { dataAsJSON, formatDate, shortID, titleCase } from '../../lib/business';
  import { renderMarkdown } from '../../utils/markdown';
  import Badge from '../ui/badge/badge.svelte';
  import Button from '../ui/button/button.svelte';
  import Empty from '../ui/empty/empty.svelte';

  type Props = {
    record: RecordView | null;
    incomingRelationships: IncomingRelationshipView[];
    projectionPending: boolean;
    graphContextUnavailable: boolean;
    graphContextPartial: boolean;
    onCreate: () => void;
    onEdit: () => void;
    onSelectRecord: (recordID: string) => void;
  };

  let {
    record,
    incomingRelationships,
    projectionPending,
    graphContextUnavailable,
    graphContextPartial,
    onCreate,
    onEdit,
    onSelectRecord,
  }: Props = $props();

  const narrativeHTML = $derived(
    renderMarkdown(record?.revision.narrative_markdown || 'No narrative has been written for this record.'),
  );
  const structuredData = $derived(dataAsJSON(record?.revision.data_json));
</script>

{#if !record}
  <div class="lens-empty">
    <Empty
      title="No record selected"
      description="Create the first business record or select one from the record index."
    />
    <Button onclick={onCreate}>Create record</Button>
  </div>
{:else}
  <article class="record-lens">
    <header class="lens-heading record-heading">
      <div>
        <div class="eyebrow-row">
          <span class="eyebrow">{record.record.type_code}</span>
          <Badge variant={record.revision.status === 'active' || record.revision.status === 'accepted' ? 'success' : 'neutral'}>
            {record.revision.status || 'No status'}
          </Badge>
          <Badge variant="outline">Revision {record.revision.revision}</Badge>
        </div>
        <h2>{record.revision.title}</h2>
        <p class="stable-key">{record.record.key}</p>
      </div>
      <Button variant="outline" onclick={onEdit}>Edit next revision</Button>
    </header>

    {#if projectionPending}
      <div class="state-banner state-banner-warn">
        Canonical changes are saved. Graph-derived relationship context is still catching up.
      </div>
    {/if}

    <div class="record-layout">
      <section class="record-reading" aria-label="Record narrative">
        <div class="markdown-reading">{@html narrativeHTML}</div>

        <section class="relationship-section">
          <div class="section-heading">
            <div>
              <span class="eyebrow">Relationships</span>
              <h3>Outgoing</h3>
            </div>
            <Badge variant="neutral">{record.relationships.length}</Badge>
          </div>

          <div class="relationship-list">
            {#each record.relationships as relationship (relationship.id)}
              <button
                type="button"
                class="relationship-row"
                class:relationship-unresolved={relationship.resolution === 'unresolved'}
                onclick={() => relationship.target_record_id && onSelectRecord(relationship.target_record_id)}
                disabled={!relationship.target_record_id}
              >
                <span class={`relationship-glyph relationship-${relationship.kind}`} aria-hidden="true"></span>
                <span>
                  <strong>{titleCase(relationship.kind)}</strong>
                  <small>{relationship.target_key}</small>
                </span>
                <Badge variant={relationship.resolution === 'resolved' ? 'success' : 'warn'} size="sm">
                  {relationship.resolution}
                </Badge>
              </button>
            {:else}
              <p class="quiet-empty">This revision has no outgoing relationships.</p>
            {/each}
          </div>
        </section>

        <section class="relationship-section">
          <div class="section-heading">
            <div>
              <span class="eyebrow">Projected context</span>
              <h3>Incoming</h3>
            </div>
            <Badge variant="neutral">{incomingRelationships.length}</Badge>
          </div>

          {#if graphContextUnavailable}
            <div class="state-banner state-banner-danger">
              Incoming relationship context is unavailable. Canonical record content remains available.
            </div>
          {:else}
            {#if graphContextPartial}
              <div class="state-banner state-banner-neutral">
                Incoming context is a bounded projected view and may be partial because filters, truncation, or projection lag are active.
              </div>
            {/if}
            <div class="relationship-list">
              {#each incomingRelationships as relationship (relationship.id)}
                <button type="button" class="relationship-row" onclick={() => onSelectRecord(relationship.sourceID)}>
                  <span class={`relationship-glyph relationship-${relationship.kind}`} aria-hidden="true"></span>
                  <span>
                    <strong>{relationship.sourceTitle}</strong>
                    <small>{relationship.sourceKey} · {titleCase(relationship.kind)}</small>
                  </span>
                </button>
              {:else}
                <p class="quiet-empty">No incoming relationships in the current projected neighborhood.</p>
              {/each}
            </div>
          {/if}
        </section>

        <details class="structured-data">
          <summary>Structured record data</summary>
          <pre>{structuredData}</pre>
        </details>
      </section>

      <aside class="context-rail" aria-label="Record context">
        <section>
          <span class="eyebrow">Stable identity</span>
          <dl class="context-list">
            <div><dt>Key</dt><dd>{record.record.key}</dd></div>
            <div><dt>Record ID</dt><dd title={record.record.id}>{shortID(record.record.id)}</dd></div>
            <div><dt>Type</dt><dd>{record.record.type_code}</dd></div>
            <div><dt>Version</dt><dd>v{record.record.version}</dd></div>
          </dl>
        </section>

        <section>
          <span class="eyebrow">Current revision</span>
          <dl class="context-list">
            <div><dt>Status</dt><dd>{record.revision.status || 'Not set'}</dd></div>
            <div><dt>Domain</dt><dd>{record.revision.domain || 'Not set'}</dd></div>
            <div><dt>Visibility</dt><dd>{record.revision.visibility || 'Not set'}</dd></div>
            <div><dt>Owner</dt><dd>{record.revision.owner_ref || 'Not set'}</dd></div>
            <div><dt>Updated</dt><dd>{formatDate(record.record.updated_at, true)}</dd></div>
          </dl>
        </section>

        <section>
          <span class="eyebrow">Source</span>
          <dl class="context-list">
            <div><dt>URI</dt><dd>{record.revision.source_uri || 'Native record'}</dd></div>
            <div><dt>Media</dt><dd>{record.revision.source_media_type || 'Not set'}</dd></div>
            <div><dt>Schema</dt><dd>{record.revision.schema_version || 'Not set'}</dd></div>
            <div><dt>Source version</dt><dd>{record.revision.source_version || 'Not set'}</dd></div>
          </dl>
        </section>
      </aside>
    </div>
  </article>
{/if}
