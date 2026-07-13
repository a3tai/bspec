<script lang="ts">
  import { Dialogs } from '@wailsio/runtime';
  import type {
    Business,
    ImportReport,
  } from '../../../bindings/github.com/a3tai/bspec/apps/graph';
  import { serviceError, titleCase } from '../../lib/business';
  import { focusTrap } from '../../lib/focusTrap';
  import Badge from '../ui/badge/badge.svelte';
  import Button from '../ui/button/button.svelte';
  import Empty from '../ui/empty/empty.svelte';

  type Props = {
    business: Business;
    report: ImportReport | null;
    inputPath: string;
    loading: boolean;
    applied: boolean;
    onReview: (path: string) => Promise<void>;
    onApply: () => Promise<void>;
    onClose: () => void;
  };

  let { business, report, inputPath, loading, applied, onReview, onApply, onClose }: Props = $props();
  let errorMessage = $state('');

  const errorDiagnostics = $derived(report?.diagnostics?.filter((item) => item.severity === 'error').length ?? 0);

  async function chooseSource() {
    errorMessage = '';
    try {
      const path = await Dialogs.OpenFile({
        Title: 'Review BSpec import',
        Message: 'Choose a .bspec archive, extracted implementation, or Markdown documents folder.',
        ButtonText: 'Review',
        CanChooseFiles: true,
        CanChooseDirectories: true,
        AllowsMultipleSelection: false,
        AllowsOtherFiletypes: true,
        Filters: [{ DisplayName: 'BSpec archives', Pattern: '*.bspec' }],
      });
      if (!path) return;
      await onReview(path);
    } catch (error) {
      errorMessage = serviceError(error).message;
    }
  }

  async function applyImport() {
    errorMessage = '';
    try {
      await onApply();
    } catch (error) {
      errorMessage = serviceError(error).message;
    }
  }
</script>

<div class="modal-backdrop" role="presentation">
  <div class="modal import-review" role="dialog" aria-modal="true" aria-labelledby="import-review-title" tabindex="-1" use:focusTrap={{ onEscape: () => { if (!loading) onClose(); } }}>
    <header class="modal-header">
      <div>
        <span class="eyebrow">BSpec interchange boundary</span>
        <h2 id="import-review-title">Review import into {business.display_name}</h2>
        <p>Dry run first. The service parses, resolves, hashes, and cycle-checks the complete source without writing canonical state.</p>
      </div>
      <button class="modal-close" type="button" onclick={onClose} aria-label="Close import review" disabled={loading}>×</button>
    </header>

    <div class="import-review-body">
      {#if errorMessage}
        <div class="state-banner state-banner-danger"><strong>Import failed</strong><span>{errorMessage}</span></div>
      {/if}

      <section class="import-source-card">
        <div>
          <span class="eyebrow">Source</span>
          <strong title={inputPath}>{inputPath || 'No BSpec source selected'}</strong>
          <small>.bspec archives are streamed. Directories are sent as bounded Markdown document imports.</small>
        </div>
        <Button variant="outline" onclick={chooseSource} disabled={loading}>{inputPath ? 'Choose another' : 'Choose source'}</Button>
      </section>

      {#if loading}
        <div class="graph-loading"><span class="spinner"></span>{report ? 'Applying canonical import' : 'Running service dry run'}</div>
      {:else if !report}
        <Empty
          title="Choose a source to review"
          description="Nothing is written until a dry-run report has been inspected and explicitly applied."
        />
      {:else}
        <section class="import-summary">
          <header class="section-heading">
            <div>
              <span class="eyebrow">{applied ? 'Applied report' : 'Dry-run report'}</span>
              <h3>{report.source_name}</h3>
            </div>
            <Badge variant={applied ? 'success' : 'accent'}>{applied ? 'Canonical' : 'No writes'}</Badge>
          </header>

          <div class="import-metrics">
            <div><span>Created</span><strong>{report.created_records}</strong></div>
            <div><span>Revised</span><strong>{report.revised_records}</strong></div>
            <div><span>Unchanged</span><strong>{report.unchanged_records}</strong></div>
            <div><span>Resolved edges</span><strong>{report.resolved_relationships}</strong></div>
            <div class:metric-warn={report.unresolved_references > 0}>
              <span>Unresolved</span><strong>{report.unresolved_references}</strong>
            </div>
          </div>

          <div class="import-diagnostics">
            <div class="section-heading">
              <div>
                <span class="eyebrow">Service diagnostics</span>
                <h3>Review findings</h3>
              </div>
              <Badge variant={errorDiagnostics > 0 ? 'danger' : 'neutral'}>{report.diagnostics?.length ?? 0}</Badge>
            </div>
            <div class="diagnostic-list">
              {#each report.diagnostics ?? [] as diagnostic, index (`${diagnostic.code}-${diagnostic.path}-${index}`)}
                <div class={`diagnostic diagnostic-${diagnostic.severity}`}>
                  <Badge variant={diagnostic.severity === 'error' ? 'danger' : diagnostic.severity === 'warning' ? 'warn' : 'neutral'} size="sm">
                    {titleCase(diagnostic.severity)}
                  </Badge>
                  <div>
                    <strong>{diagnostic.message}</strong>
                    <small>{diagnostic.code}{diagnostic.path ? ` · ${diagnostic.path}` : ''}</small>
                  </div>
                </div>
              {:else}
                <p class="quiet-empty">The service returned no diagnostics.</p>
              {/each}
            </div>
          </div>

          <div class="state-banner state-banner-neutral">
            The current API reports aggregate changes and diagnostics. It does not yet expose a per-record preview plan, so this UI does not recreate importer semantics locally.
          </div>
        </section>
      {/if}
    </div>

    <footer class="modal-actions">
      <Button variant="ghost" onclick={onClose} disabled={loading}>{applied ? 'Done' : 'Cancel'}</Button>
      {#if report && !applied}
        <Button onclick={applyImport} disabled={loading || errorDiagnostics > 0}>
          Apply with fresh idempotency key
        </Button>
      {/if}
    </footer>
  </div>
</div>
