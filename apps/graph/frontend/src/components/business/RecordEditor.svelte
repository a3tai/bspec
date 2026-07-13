<script lang="ts">
  import type { RecordView } from '../../../bindings/github.com/a3tai/bspec/apps/graph';
  import type { RecordMutation } from '../../lib/business';
  import { dataAsJSON, serviceError, titleCase } from '../../lib/business';
  import { focusTrap } from '../../lib/focusTrap';
  import Badge from '../ui/badge/badge.svelte';
  import Button from '../ui/button/button.svelte';

  type EditableRelationship = {
    kind: string;
    targetKey: string;
    strength: string;
    provenanceJSON?: string;
  };

  type ReconciliationField =
    | 'type_code'
    | 'title'
    | 'status'
    | 'domain'
    | 'visibility'
    | 'owner_ref'
    | 'schema_version'
    | 'source_version'
    | 'data_json'
    | 'narrative_markdown'
    | 'raw_markdown'
    | 'source_uri'
    | 'source_media_type'
    | 'effective_at'
    | 'relationships';

  type ReconciliationConflict = {
    key: ReconciliationField;
    label: string;
    baseValue: string;
    latestValue: string;
    draftValue: string;
  };

  type Props = {
    record: RecordView | null;
    onSave: (input: RecordMutation) => Promise<void>;
    onClose: () => void;
  };

  let { record, onSave, onClose }: Props = $props();

  const relationshipKinds = ['depends_on', 'enables', 'conflicts_with', 'related', 'supersedes'];
  const isCreating = $derived(record === null);

  function initial<T>(read: (value: RecordView | null) => T): T {
    return read(record);
  }

  function editableRelationships(value: RecordView | null): EditableRelationship[] {
    return (value?.relationships ?? []).map((relationship) => ({
      kind: relationship.kind,
      targetKey: relationship.target_key,
      strength: relationship.strength ?? '',
      provenanceJSON: relationship.provenance_json,
    }));
  }

  let key = $state(initial((value) => value?.record.key ?? ''));
  let typeCode = $state(initial((value) => value?.record.type_code ?? ''));
  let title = $state(initial((value) => value?.revision.title ?? ''));
  let status = $state(initial((value) => value?.revision.status ?? 'draft'));
  let domain = $state(initial((value) => value?.revision.domain ?? ''));
  let visibility = $state(initial((value) => value ? (value.revision.visibility ?? '') : 'internal'));
  let ownerRef = $state(initial((value) => value?.revision.owner_ref ?? ''));
  let schemaVersion = $state(initial((value) => value?.revision.schema_version ?? ''));
  let sourceVersion = $state(initial((value) => value?.revision.source_version ?? ''));
  let narrative = $state(initial((value) => value?.revision.narrative_markdown ?? ''));
  let rawMarkdown = $state(initial((value) => value?.revision.raw_markdown ?? ''));
  let sourceURI = $state(initial((value) => value?.revision.source_uri ?? ''));
  let sourceMediaType = $state(initial((value) => value?.revision.source_media_type ?? ''));
  let effectiveAt = $state(initial((value) => value?.revision.effective_at ?? ''));
  let dataText = $state(initial((value) => dataAsJSON(value?.revision.data_json)));
  let relationships = $state<EditableRelationship[]>(initial(editableRelationships));
  let saving = $state(false);
  let errorMessage = $state('');
  let staleConflict = $state(false);
  let baseVersion = $state(initial((value) => value?.record.version ?? 0));
  let baseRecord = $state<RecordView | null>(initial((value) => value));
  let reconciliationPrepared = $state(false);
  let reconciliationConflicts = $state<ReconciliationConflict[]>([]);
  const latestAvailable = $derived(staleConflict && record !== null && record.record.version > baseVersion);

  function addRelationship() {
    relationships = [
      ...relationships,
      { kind: 'related', targetKey: '', strength: '' },
    ];
  }

  function removeRelationship(index: number) {
    relationships = relationships.filter((_, relationshipIndex) => relationshipIndex !== index);
  }

  function applyLatestValue(key: ReconciliationField, value: string) {
    switch (key) {
      case 'type_code': typeCode = value; break;
      case 'title': title = value; break;
      case 'status': status = value; break;
      case 'domain': domain = value; break;
      case 'visibility': visibility = value; break;
      case 'owner_ref': ownerRef = value; break;
      case 'schema_version': schemaVersion = value; break;
      case 'source_version': sourceVersion = value; break;
      case 'data_json': dataText = value; break;
      case 'narrative_markdown': narrative = value; break;
      case 'raw_markdown': rawMarkdown = value; break;
      case 'source_uri': sourceURI = value; break;
      case 'source_media_type': sourceMediaType = value; break;
      case 'effective_at': effectiveAt = value; break;
      case 'relationships': relationships = JSON.parse(value) as EditableRelationship[]; break;
    }
  }

  function prepareReconciliation() {
    if (!baseRecord || !record || record.record.version <= baseVersion) return;
    const latestRelationships = JSON.stringify(editableRelationships(record));
    const fields: ReconciliationConflict[] = [
      { key: 'type_code', label: 'Type code', baseValue: baseRecord.record.type_code, latestValue: record.record.type_code, draftValue: typeCode },
      { key: 'title', label: 'Title', baseValue: baseRecord.revision.title, latestValue: record.revision.title, draftValue: title },
      { key: 'status', label: 'Status', baseValue: baseRecord.revision.status ?? '', latestValue: record.revision.status ?? '', draftValue: status },
      { key: 'domain', label: 'Domain', baseValue: baseRecord.revision.domain ?? '', latestValue: record.revision.domain ?? '', draftValue: domain },
      { key: 'visibility', label: 'Visibility', baseValue: baseRecord.revision.visibility ?? '', latestValue: record.revision.visibility ?? '', draftValue: visibility },
      { key: 'owner_ref', label: 'Owner', baseValue: baseRecord.revision.owner_ref ?? '', latestValue: record.revision.owner_ref ?? '', draftValue: ownerRef },
      { key: 'schema_version', label: 'Schema version', baseValue: baseRecord.revision.schema_version ?? '', latestValue: record.revision.schema_version ?? '', draftValue: schemaVersion },
      { key: 'source_version', label: 'Source version', baseValue: baseRecord.revision.source_version ?? '', latestValue: record.revision.source_version ?? '', draftValue: sourceVersion },
      { key: 'data_json', label: 'Structured data', baseValue: dataAsJSON(baseRecord.revision.data_json), latestValue: dataAsJSON(record.revision.data_json), draftValue: dataText },
      { key: 'narrative_markdown', label: 'Narrative', baseValue: baseRecord.revision.narrative_markdown ?? '', latestValue: record.revision.narrative_markdown ?? '', draftValue: narrative },
      { key: 'raw_markdown', label: 'Raw Markdown', baseValue: baseRecord.revision.raw_markdown ?? '', latestValue: record.revision.raw_markdown ?? '', draftValue: rawMarkdown },
      { key: 'source_uri', label: 'Source URI', baseValue: baseRecord.revision.source_uri ?? '', latestValue: record.revision.source_uri ?? '', draftValue: sourceURI },
      { key: 'source_media_type', label: 'Source media type', baseValue: baseRecord.revision.source_media_type ?? '', latestValue: record.revision.source_media_type ?? '', draftValue: sourceMediaType },
      { key: 'effective_at', label: 'Effective at', baseValue: baseRecord.revision.effective_at ?? '', latestValue: record.revision.effective_at ?? '', draftValue: effectiveAt },
      { key: 'relationships', label: 'Outgoing relationships', baseValue: JSON.stringify(editableRelationships(baseRecord)), latestValue: latestRelationships, draftValue: JSON.stringify(relationships) },
    ];

    const conflicts: ReconciliationConflict[] = [];
    for (const field of fields) {
      if (field.latestValue === field.baseValue || field.latestValue === field.draftValue) continue;
      if (field.draftValue === field.baseValue) {
        applyLatestValue(field.key, field.latestValue);
      } else {
        conflicts.push(field);
      }
    }
    reconciliationConflicts = conflicts;
    reconciliationPrepared = true;
  }

  function resolveConflict(conflict: ReconciliationConflict, useLatest: boolean) {
    if (useLatest) applyLatestValue(conflict.key, conflict.latestValue);
    reconciliationConflicts = reconciliationConflicts.filter((candidate) => candidate.key !== conflict.key);
  }

  function completeReconciliation() {
    if (!record || !reconciliationPrepared || reconciliationConflicts.length > 0) return;
    baseRecord = record;
    baseVersion = record.record.version;
    staleConflict = false;
    reconciliationPrepared = false;
    errorMessage = '';
  }

  function reconciliationPreview(value: string): string {
    if (value === '') return 'Not set';
    return value.length > 180 ? `${value.slice(0, 177)}…` : value;
  }

  async function submit(event: SubmitEvent) {
    event.preventDefault();
    if (staleConflict) return;
    errorMessage = '';
    staleConflict = false;

    const normalizedKey = key.trim();
    const normalizedType = typeCode.trim().toUpperCase();
    if (isCreating && !/^[A-Za-z0-9][A-Za-z0-9._:-]{0,191}$/.test(normalizedKey)) {
      errorMessage = 'The stable key must be one URL-safe segment using letters, numbers, dot, underscore, colon, or hyphen.';
      return;
    }
    if (!/^[A-Z][A-Z0-9._-]{1,31}$/.test(normalizedType)) {
      errorMessage = 'Type code must be 2 to 32 uppercase letters, numbers, dots, underscores, or hyphens.';
      return;
    }
    if (!title.trim()) {
      errorMessage = 'Title is required.';
      return;
    }

    const dataJSON = (dataText || '{}').trim();
    try {
      JSON.parse(dataJSON);
    } catch {
      errorMessage = 'Structured data must be valid JSON.';
      return;
    }

    const normalizedRelationships = relationships
      .map((relationship) => ({
        kind: relationship.kind,
        target_key: relationship.targetKey.trim(),
        strength: relationship.strength.trim(),
        ...(relationship.provenanceJSON === undefined ? {} : { provenance_json: relationship.provenanceJSON }),
      }))
      .filter((relationship) => relationship.target_key);

    const duplicateCheck = new Set<string>();
    for (const relationship of normalizedRelationships) {
      const identity = `${relationship.kind}:${relationship.target_key}`;
      if (duplicateCheck.has(identity)) {
        errorMessage = `Duplicate ${titleCase(relationship.kind)} relationship to ${relationship.target_key}.`;
        return;
      }
      duplicateCheck.add(identity);
      if (relationship.target_key === normalizedKey) {
        errorMessage = 'A record cannot relate to itself.';
        return;
      }
    }

    const mutation: RecordMutation = {
      ...(isCreating ? { key: normalizedKey } : {}),
      type_code: normalizedType,
      revision: {
        title: title.trim(),
        status: status.trim(),
        domain: domain.trim(),
        visibility,
        owner_ref: ownerRef.trim(),
        schema_version: schemaVersion.trim(),
        source_version: sourceVersion.trim(),
        data_json: dataJSON,
        narrative_markdown: narrative,
        raw_markdown: rawMarkdown,
        source_uri: sourceURI.trim(),
        source_media_type: sourceMediaType.trim(),
        effective_at: effectiveAt || null,
        relationships: normalizedRelationships,
      },
    };

    saving = true;
    try {
      await onSave(mutation);
    } catch (error) {
      const failure = serviceError(error);
      staleConflict = failure.code === 'BUSINESS_PRECONDITION_FAILED';
      reconciliationPrepared = false;
      reconciliationConflicts = [];
      errorMessage = staleConflict
        ? 'A newer revision was saved while this editor was open. Review the latest canonical summary, then explicitly rebase this preserved draft before saving.'
        : failure.message;
    } finally {
      saving = false;
    }
  }
</script>

<div class="modal-backdrop" role="presentation">
  <div class="modal record-editor" role="dialog" aria-modal="true" aria-labelledby="record-editor-title" tabindex="-1" use:focusTrap={{ onEscape: () => { if (!saving) onClose(); } }}>
    <header class="modal-header">
      <div>
        <span class="eyebrow">{isCreating ? 'New stable record' : `Complete revision after v${baseVersion}`}</span>
        <h2 id="record-editor-title">{isCreating ? 'Create business record' : `Revise ${record?.revision.title}`}</h2>
        <p>Every save is a complete immutable snapshot of content, metadata, and outgoing relationships.</p>
      </div>
      <button class="modal-close" type="button" onclick={onClose} aria-label="Close record editor" disabled={saving}>×</button>
    </header>

    <form class="record-editor-form" onsubmit={submit}>
      {#if errorMessage}
        <div class="state-banner" class:state-banner-danger={!staleConflict} class:state-banner-warn={staleConflict}>
          <strong>{staleConflict ? 'Reconciliation required' : 'Could not save'}</strong>
          <span>{errorMessage}</span>
          {#if staleConflict && record}
            <div class="reconciliation-card">
              <div class="reconciliation-latest">
                <span>Latest canonical v{record.record.version}</span>
                <strong>{record.revision.title}</strong>
                <small>{record.revision.status || 'No status'} · {record.revision.domain || 'No domain'} · {record.relationships.length} outgoing relationships</small>
              </div>
              {#if !reconciliationPrepared}
                <Button variant="outline" size="sm" onclick={prepareReconciliation} disabled={!latestAvailable}>
                  {latestAvailable ? 'Compare base, latest, and draft' : 'Loading latest canonical revision'}
                </Button>
              {:else if reconciliationConflicts.length > 0}
                <div class="reconciliation-conflicts">
                  <strong>{reconciliationConflicts.length} conflicting {reconciliationConflicts.length === 1 ? 'field' : 'fields'} need a choice</strong>
                  {#each reconciliationConflicts as conflict (conflict.key)}
                    <section>
                      <h4>{conflict.label}</h4>
                      <div><span>Base</span><code>{reconciliationPreview(conflict.baseValue)}</code></div>
                      <div><span>Latest</span><code>{reconciliationPreview(conflict.latestValue)}</code></div>
                      <div><span>Your draft</span><code>{reconciliationPreview(conflict.draftValue)}</code></div>
                      <footer>
                        <Button variant="ghost" size="sm" onclick={() => resolveConflict(conflict, false)}>Keep draft</Button>
                        <Button variant="outline" size="sm" onclick={() => resolveConflict(conflict, true)}>Use latest</Button>
                      </footer>
                    </section>
                  {/each}
                </div>
              {:else}
                <div class="reconciliation-ready">
                  <span>Non-conflicting concurrent changes were merged into the draft.</span>
                  <Button variant="outline" size="sm" onclick={completeReconciliation}>Complete rebase onto v{record.record.version}</Button>
                </div>
              {/if}
            </div>
          {/if}
        </div>
      {/if}

      <fieldset class="record-editor-fields" disabled={staleConflict}>
      <div class="editor-grid">
        <label>
          <span>Stable key</span>
          <input bind:value={key} required disabled={!isCreating} placeholder="strategy-2026" />
          <small>{isCreating ? 'Permanent business-scoped identity' : 'Stable identity cannot be changed'}</small>
        </label>
        <label>
          <span>Type code</span>
          <input bind:value={typeCode} required maxlength="32" placeholder="STR" />
          <small>Uppercase ontology type</small>
        </label>
        <label class="editor-span-2">
          <span>Title</span>
          <input bind:value={title} required maxlength="500" placeholder="Business strategy" />
        </label>
        <label>
          <span>Status</span>
          <input bind:value={status} placeholder="draft" />
        </label>
        <label>
          <span>Domain</span>
          <input bind:value={domain} placeholder="Strategic Foundation" />
        </label>
        <label>
          <span>Visibility classification</span>
          <select bind:value={visibility}>
            <option value="">Not set</option>
            <option value="public">Public</option>
            <option value="internal">Internal</option>
            <option value="confidential">Confidential</option>
          </select>
          <small>Classification only, not record-level authorization</small>
        </label>
        <label>
          <span>Owner reference</span>
          <input bind:value={ownerRef} placeholder="team:strategy" />
        </label>
      </div>

      <label class="editor-field">
        <span>Narrative Markdown</span>
        <textarea class="narrative-editor" bind:value={narrative} placeholder="Write the business meaning in Markdown..."></textarea>
      </label>

      <section class="relationship-editor">
        <header class="section-heading">
          <div>
            <span class="eyebrow">Complete outgoing snapshot</span>
            <h3>Relationships</h3>
          </div>
          <Button variant="outline" size="sm" onclick={addRelationship}>Add relationship</Button>
        </header>

        <div class="relationship-edit-list">
          {#each relationships as relationship, index (index)}
            <div class="relationship-edit-row">
              <select bind:value={relationship.kind} aria-label="Relationship kind">
                {#each relationshipKinds as kind}
                  <option value={kind}>{titleCase(kind)}</option>
                {/each}
              </select>
              <input bind:value={relationship.targetKey} placeholder="Target stable key" aria-label="Target stable key" />
              <input bind:value={relationship.strength} placeholder="Strength (optional)" aria-label="Relationship strength" />
              <button type="button" onclick={() => removeRelationship(index)} aria-label="Remove relationship">×</button>
            </div>
          {:else}
            <p class="quiet-empty">No outgoing relationships. Add one only when the direction and target are explicit.</p>
          {/each}
        </div>
      </section>

      <details class="advanced-editor">
        <summary>
          <span>Structured data and source provenance</span>
          <Badge variant="outline" size="sm">Preserved on revision</Badge>
        </summary>
        <div class="editor-grid">
          <label>
            <span>Schema version</span>
            <input bind:value={schemaVersion} />
          </label>
          <label>
            <span>Source version</span>
            <input bind:value={sourceVersion} />
          </label>
          <label>
            <span>Source URI</span>
            <input bind:value={sourceURI} />
          </label>
          <label>
            <span>Source media type</span>
            <input bind:value={sourceMediaType} />
          </label>
          <label>
            <span>Effective at</span>
            <input bind:value={effectiveAt} placeholder="RFC 3339 timestamp" />
          </label>
          <label class="editor-span-2">
            <span>Structured JSON</span>
            <textarea class="data-editor" bind:value={dataText}></textarea>
          </label>
          <label class="editor-span-2">
            <span>Original raw Markdown</span>
            <textarea class="raw-editor" bind:value={rawMarkdown}></textarea>
            <small>Preserved as source provenance. Narrative above remains the canonical reading surface.</small>
          </label>
        </div>
      </details>
      </fieldset>

      <footer class="modal-actions">
        <Button variant="ghost" onclick={onClose} disabled={saving}>Cancel</Button>
        <Button type="submit" disabled={saving || staleConflict}>{saving ? 'Saving…' : isCreating ? 'Create record' : 'Save next revision'}</Button>
      </footer>
    </form>
  </div>
</div>
