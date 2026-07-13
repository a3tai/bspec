<script lang="ts">
  import type {
    Business,
    BusinessIdentifier,
  } from '../../../bindings/github.com/a3tai/bspec/apps/graph';
  import { formatDate, shortID, titleCase } from '../../lib/business';
  import Badge from '../ui/badge/badge.svelte';
  import Empty from '../ui/empty/empty.svelte';

  type Props = {
    business: Business;
    identifiers: BusinessIdentifier[];
  };

  let { business, identifiers }: Props = $props();

  const groupedIdentifiers = $derived.by(() => {
    const groups = new Map<string, BusinessIdentifier[]>();
    for (const identifier of identifiers) {
      const group = groups.get(identifier.scheme) ?? [];
      group.push(identifier);
      groups.set(identifier.scheme, group);
    }
    return [...groups.entries()].sort(([left], [right]) => left.localeCompare(right));
  });
</script>

<section class="identity-lens">
  <header class="lens-heading">
    <div>
      <span class="eyebrow">Managed business identity</span>
      <h2>{business.display_name}</h2>
      <p>The business is a graph root owned by a tenant. It is not the tenant or its optional A3T organization.</p>
    </div>
    <Badge variant="outline">Read only</Badge>
  </header>

  <div class="state-banner state-banner-neutral">
    Visibility is classification data in the current service. It is not record-level authorization enforcement.
  </div>

  <div class="identity-map" aria-label="Business identity hierarchy">
    <div class="identity-node identity-tenant">
      <span>Tenant owner</span>
      <strong title={business.tenant_id}>{shortID(business.tenant_id)}</strong>
      <small>Authenticated scope</small>
    </div>
    <div class="identity-link"><span>owns</span></div>
    <div class="identity-node identity-business">
      <span>Business root</span>
      <strong>{business.display_name}</strong>
      <small>{business.slug} · {business.status}</small>
    </div>
    {#if business.organization_id}
      <div class="identity-link"><span>associated with</span></div>
      <div class="identity-node identity-organization">
        <span>A3T organization</span>
        <strong title={business.organization_id}>{shortID(business.organization_id)}</strong>
        <small>Same-tenant association</small>
      </div>
    {/if}
  </div>

  <div class="identity-grid">
    <section class="identity-details">
      <div class="section-heading">
        <div>
          <span class="eyebrow">Business profile</span>
          <h3>Canonical root</h3>
        </div>
        <Badge variant={business.status === 'active' ? 'success' : 'neutral'}>{business.status}</Badge>
      </div>
      <dl class="context-list identity-context">
        <div><dt>Business ID</dt><dd title={business.id}>{business.id}</dd></div>
        <div><dt>Slug</dt><dd>{business.slug}</dd></div>
        <div><dt>Legal name</dt><dd>{business.legal_name || 'Not set'}</dd></div>
        <div><dt>Visibility</dt><dd>{business.visibility}</dd></div>
        <div><dt>Ontology</dt><dd>{business.ontology_version || 'Not set'}</dd></div>
        <div><dt>Version</dt><dd>v{business.version}</dd></div>
        <div><dt>Updated</dt><dd>{formatDate(business.updated_at, true)}</dd></div>
      </dl>
      {#if business.description}
        <p class="business-description">{business.description}</p>
      {/if}
    </section>

    <section class="identifier-panel">
      <div class="section-heading">
        <div>
          <span class="eyebrow">External identity</span>
          <h3>Namespaced identifiers</h3>
        </div>
        <Badge variant="neutral">{identifiers.length}</Badge>
      </div>

      {#if groupedIdentifiers.length === 0}
        <Empty
          title="No external identifiers"
          description="Identifiers such as domains, brands, and registrations will appear here when added through the service."
        />
      {:else}
        <div class="identifier-groups">
          {#each groupedIdentifiers as [scheme, items] (scheme)}
            <section>
              <header>
                <h4>{titleCase(scheme)}</h4>
                <Badge variant="outline" size="sm">{items.length}</Badge>
              </header>
              {#each items as identifier (identifier.id)}
                <div class="identifier-row">
                  <div>
                    <strong>{identifier.value}</strong>
                    <small>{identifier.normalized_value}</small>
                  </div>
                  <div class="identifier-state">
                    {#if identifier.primary}<Badge variant="accent" size="sm">Primary</Badge>{/if}
                    <Badge variant={identifier.verified_at ? 'success' : 'neutral'} size="sm">
                      {identifier.verified_at ? 'Verified' : 'Unverified'}
                    </Badge>
                  </div>
                </div>
              {/each}
            </section>
          {/each}
        </div>
      {/if}
    </section>
  </div>

  <p class="read-only-note">
    Identifier management stays read only until update, verification, removal, and primary-selection lifecycle operations exist in the service.
  </p>
</section>
