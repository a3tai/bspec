<script lang="ts">
  import { onMount } from 'svelte';
  import {
    BusinessService,
    type Business,
    type BusinessIdentifier,
    type BusinessServiceStatus,
    type CreateBusinessInput,
    type CreateRecordInput,
    type ImportReport,
    type ImportReviewResult,
    type Neighborhood,
    type PathResult,
    type ProjectionStatus,
    type RecordPage,
    type RecordRevision,
    type RecordView,
    type ReviseRecordInput,
    type RevisionPage,
  } from '../bindings/github.com/a3tai/bspec/apps/graph';
  import HistoryLens from './components/business/HistoryLens.svelte';
  import IdentityLens from './components/business/IdentityLens.svelte';
  import ImportReview from './components/business/ImportReview.svelte';
  import LandscapeLens from './components/business/LandscapeLens.svelte';
  import RecordEditor from './components/business/RecordEditor.svelte';
  import RecordLens from './components/business/RecordLens.svelte';
  import Badge from './components/ui/badge/badge.svelte';
  import Button from './components/ui/button/button.svelte';
  import Empty from './components/ui/empty/empty.svelte';
  import Lamp from './components/ui/lamp/lamp.svelte';
  import NativeSelect from './components/ui/native-select/native-select.svelte';
  import SearchInput from './components/ui/search-input/search-input.svelte';
  import type {
    BusinessLens,
    IncomingRelationshipView,
    RecordMutation,
  } from './lib/business';
  import { formatDate, serviceError, shortID, titleCase } from './lib/business';
  import { focusTrap } from './lib/focusTrap';

  const allRelationshipKinds = ['depends_on', 'enables', 'conflicts_with', 'related', 'supersedes'];

  let serviceStatus = $state<BusinessServiceStatus | null>(null);
  let businesses = $state<Business[]>([]);
  let business = $state<Business | null>(null);
  let identifiers = $state<BusinessIdentifier[]>([]);
  let records = $state<RecordView[]>([]);
  let recordCache = $state<Record<string, RecordView>>({});
  let recordsNextCursor = $state('');
  let activeRecord = $state<RecordView | null>(null);
  let selectedRecordID = $state('');
  let projection = $state<ProjectionStatus | null>(null);
  let projectionLoading = $state(false);
  let projectionError = $state('');
  let neighborhood = $state<Neighborhood | null>(null);
  let pathResult = $state<PathResult | null>(null);
  let revisions = $state<RecordRevision[]>([]);
  let revisionsNextCursor = $state('');
  let historySnapshot = $state<RecordView | null>(null);
  let previousHistorySnapshot = $state<RecordView | null>(null);
  let selectedRevisionID = $state('');

  let lens = $state<BusinessLens>('landscape');
  let recordQuery = $state('');
  let recordTypeFilter = $state('');
  let recordStatusFilter = $state('');
  let activeRelationshipKinds = $state([...allRelationshipKinds]);
  let pathTargetID = $state('');
  let pathTargetQuery = $state('');
  let pathTargetRecords = $state<RecordView[]>([]);
  let pathTargetsLoading = $state(false);

  let booting = $state(true);
  let portfolioLoading = $state(false);
  let businessLoading = $state(false);
  let recordsLoading = $state(false);
  let graphLoading = $state(false);
  let pathLoading = $state(false);
  let historyLoading = $state(false);
  let globalError = $state('');
  let graphErrorCode = $state('');
  let graphErrorMessage = $state('');
  let pathErrorCode = $state('');
  let pathErrorMessage = $state('');

  let settingsOpen = $state(false);
  let settingsSaving = $state(false);
  let settingsError = $state('');
  let configBaseURL = $state('http://localhost:7290');
  let configAccessToken = $state('');
  let configTenantID = $state('');
  let configActorID = $state('');
  let configPreserveCredentials = $state(false);

  let createBusinessOpen = $state(false);
  let createBusinessSaving = $state(false);
  let createBusinessError = $state('');
  let newBusinessSlug = $state('');
  let newBusinessName = $state('');
  let newBusinessLegalName = $state('');
  let newBusinessDescription = $state('');
  let newBusinessOntology = $state('');
  let createBusinessOperationID = $state('');

  let recordEditorOpen = $state(false);
  let recordEditorRecord = $state<RecordView | null>(null);
  let createRecordOperationID = $state('');

  let importOpen = $state(false);
  let importLoading = $state(false);
  let importPath = $state('');
  let importReport = $state<ImportReport | null>(null);
  let importReviewID = $state('');
  let importApplied = $state(false);

  let searchTimer: number | undefined;
  let pathSearchTimer: number | undefined;
  let portfolioRequest = 0;
  let businessRequest = 0;
  let recordsRequest = 0;
  let recordRequest = 0;
  let neighborhoodRequest = 0;
  let pathRequest = 0;
  let pathTargetsRequest = 0;
  let projectionRequest = 0;
  let historyPageRequest = 0;
  let historySelectionRequest = 0;

  const selectedRecord = $derived(
    recordCache[selectedRecordID] ??
      (activeRecord?.record.id === selectedRecordID ? activeRecord : null),
  );
  const cachedRecords = $derived(Object.values(recordCache));
  const projectionPending = $derived(projection?.status === 'pending');
  const primaryIdentifiers = $derived(identifiers.filter((identifier) => identifier.primary).slice(0, 3));
  const recordTypeOptions = $derived([
    { value: '', label: 'All types' },
    ...[...new Set([
      ...records.map((record) => record.record.type_code),
      ...(recordTypeFilter ? [recordTypeFilter] : []),
    ])]
      .sort()
      .map((value) => ({ value, label: value })),
  ]);
  const recordStatusOptions = $derived([
    { value: '', label: 'All statuses' },
    ...[...new Set([
      ...records.map((record) => record.revision.status),
      ...(recordStatusFilter ? [recordStatusFilter] : []),
    ]
      .filter((value): value is string => typeof value === 'string' && value.length > 0))]
      .sort()
      .map((value) => ({ value, label: titleCase(value) })),
  ]);
  const incomingRelationships = $derived.by<IncomingRelationshipView[]>(() => {
    if (!neighborhood || !selectedRecordID) return [];
    const nodes = new Map(neighborhood.nodes.map((node) => [node.id, node]));
    return neighborhood.edges
      .filter((edge) => edge.target_id === selectedRecordID)
      .map((edge) => {
        const source = nodes.get(edge.source_id);
        return {
          id: edge.id,
          sourceID: edge.source_id,
          sourceKey: source?.key ?? edge.source_id,
          sourceTitle: source?.title ?? 'Related record',
          kind: edge.kind,
          strength: edge.strength || '',
        };
      });
  });

  onMount(() => {
    const savedLens = localStorage.getItem('a3t-business-lens');
    if (savedLens === 'landscape' || savedLens === 'record' || savedLens === 'history' || savedLens === 'identity') {
      lens = savedLens;
    }
    void bootstrap();
    const projectionTimer = window.setInterval(() => {
      if (business && projection?.status === 'pending') void refreshProjection();
    }, 2500);

    return () => {
      window.clearInterval(projectionTimer);
      if (searchTimer) window.clearTimeout(searchTimer);
      if (pathSearchTimer) window.clearTimeout(pathSearchTimer);
    };
  });

  async function bootstrap() {
    booting = true;
    globalError = '';
    try {
      serviceStatus = await BusinessService.Status();
      if (!serviceStatus?.ready) {
        globalError = serviceStatus?.message || 'The business service is not ready.';
        return;
      }
      configBaseURL = serviceStatus.base_url;
      await loadPortfolio(true);
    } catch (error) {
      globalError = serviceError(error).message;
    } finally {
      booting = false;
    }
  }

  async function loadPortfolio(restoreSelection = false) {
    const request = ++portfolioRequest;
    portfolioLoading = true;
    globalError = '';
    try {
      const loadedBusinesses = await BusinessService.ListBusinesses();
      if (request !== portfolioRequest) return;
      businesses = loadedBusinesses;
      if (restoreSelection) {
        const savedID = localStorage.getItem('a3t-business-id');
        const candidate = businesses.find((item) => item.id === savedID) ?? (businesses.length === 1 ? businesses[0] : null);
        if (candidate) await selectBusiness(candidate.id);
      }
    } catch (error) {
      if (request !== portfolioRequest) return;
      const failure = serviceError(error);
      globalError = failure.code === 'BUSINESS_FORBIDDEN'
        ? 'Your current identity does not have businesses.read permission.'
        : failure.message;
    } finally {
      if (request === portfolioRequest) portfolioLoading = false;
    }
  }

  async function selectBusiness(businessID: string) {
    businessLoading = true;
    globalError = '';
    clearBusinessState();
    const request = ++businessRequest;
    try {
      const [loadedBusiness, loadedIdentifiers] = await Promise.all([
        BusinessService.GetBusiness(businessID),
        BusinessService.ListIdentifiers(businessID),
      ]);
      if (request !== businessRequest) return;
      if (!loadedBusiness) throw new Error('Business service returned no business.');
      business = loadedBusiness;
      identifiers = loadedIdentifiers;
      localStorage.setItem('a3t-business-id', loadedBusiness.id);
      selectedRecordID = localStorage.getItem(`a3t-business-record:${loadedBusiness.id}`) ?? '';
      await loadRecordPage(false);
      if (request !== businessRequest) return;
      void loadPathTargetRecords('');
      await refreshProjection();
      if (request !== businessRequest) return;
      if (selectedRecordID) {
        await selectRecord(selectedRecordID, false);
      }
    } catch (error) {
      if (request !== businessRequest) return;
      globalError = serviceError(error).message;
      business = null;
    } finally {
      if (request === businessRequest) businessLoading = false;
    }
  }

  function clearBusinessState() {
    businessRequest++;
    recordsRequest++;
    recordRequest++;
    neighborhoodRequest++;
    pathRequest++;
    pathTargetsRequest++;
    projectionRequest++;
    historyPageRequest++;
    historySelectionRequest++;
    business = null;
    identifiers = [];
    records = [];
    recordCache = {};
    recordsNextCursor = '';
    activeRecord = null;
    selectedRecordID = '';
    projection = null;
    projectionLoading = false;
    projectionError = '';
    neighborhood = null;
    pathResult = null;
    pathTargetQuery = '';
    pathTargetRecords = [];
    pathTargetsLoading = false;
    revisions = [];
    historySnapshot = null;
    previousHistorySnapshot = null;
    graphErrorCode = '';
    graphErrorMessage = '';
    pathErrorCode = '';
    pathErrorMessage = '';
  }

  function rememberRecords(views: RecordView[]) {
    const next = { ...recordCache };
    for (const view of views) next[view.record.id] = view;
    recordCache = next;
  }

  function showPortfolio() {
    clearBusinessState();
    localStorage.removeItem('a3t-business-id');
  }

  async function loadRecordPage(append: boolean) {
    if (!business) return;
    const request = ++recordsRequest;
    const businessID = business.id;
    recordsLoading = true;
    try {
      const cursor = append ? recordsNextCursor : '';
      const page = await BusinessService.ListRecordPage(
        businessID,
        recordQuery.trim(),
        recordTypeFilter,
        recordStatusFilter,
        cursor,
        100,
      ) as RecordPage | null;
      if (request !== recordsRequest || business?.id !== businessID) return;
      if (!page) throw new Error('Business service returned no record page.');
      rememberRecords(page.items);
      if (append) {
        const byID = new Map(records.map((view) => [view.record.id, view]));
        for (const view of page.items) byID.set(view.record.id, view);
        records = [...byID.values()];
      } else {
        records = page.items;
      }
      recordsNextCursor = page.next_cursor || '';

      if (!selectedRecordID && records.length > 0) {
        selectedRecordID = records[0].record.id;
        activeRecord = records[0];
      } else if (selectedRecordID) {
        activeRecord = records.find((view) => view.record.id === selectedRecordID) ?? activeRecord;
      }
    } catch (error) {
      if (request !== recordsRequest || business?.id !== businessID) return;
      globalError = serviceError(error).message;
    } finally {
      if (request === recordsRequest) recordsLoading = false;
    }
  }

  function scheduleRecordSearch() {
    if (searchTimer) window.clearTimeout(searchTimer);
    recordsRequest++;
    searchTimer = window.setTimeout(() => void loadRecordPage(false), 280);
  }

  function changeRecordFilter() {
    void loadRecordPage(false);
  }

  function schedulePathTargetSearch(event: Event) {
    pathTargetQuery = (event.currentTarget as HTMLInputElement).value;
    if (pathSearchTimer) window.clearTimeout(pathSearchTimer);
    pathTargetsRequest++;
    pathSearchTimer = window.setTimeout(() => void loadPathTargetRecords(pathTargetQuery), 240);
  }

  async function loadPathTargetRecords(query: string) {
    if (!business) return;
    const request = ++pathTargetsRequest;
    const businessID = business.id;
    pathTargetsLoading = true;
    try {
      const page = await BusinessService.ListRecordPage(businessID, query.trim(), '', '', '', 50) as RecordPage | null;
      if (request !== pathTargetsRequest || business?.id !== businessID) return;
      if (!page) throw new Error('Business service returned no path destination page.');
      pathTargetRecords = page.items;
      rememberRecords(page.items);
    } catch (error) {
      if (request !== pathTargetsRequest || business?.id !== businessID) return;
      pathErrorCode = serviceError(error).code;
      pathErrorMessage = serviceError(error).message;
    } finally {
      if (request === pathTargetsRequest) pathTargetsLoading = false;
    }
  }

  async function selectRecord(recordID: string, persist = true) {
    if (!business || !recordID) return;
    const request = ++recordRequest;
    const businessID = business.id;
    selectedRecordID = recordID;
    pathTargetID = '';
    pathResult = null;
    if (persist) localStorage.setItem(`a3t-business-record:${business.id}`, recordID);

    activeRecord = recordCache[recordID] ?? records.find((view) => view.record.id === recordID) ?? null;
    if (!activeRecord) {
      try {
        const loadedRecord = await BusinessService.GetRecord(businessID, recordID);
        if (request !== recordRequest || business?.id !== businessID || selectedRecordID !== recordID) return;
        activeRecord = loadedRecord;
        if (loadedRecord) rememberRecords([loadedRecord]);
      } catch (error) {
        if (request !== recordRequest || business?.id !== businessID || selectedRecordID !== recordID) return;
        globalError = serviceError(error).message;
        return;
      }
    }

    if (request !== recordRequest || business?.id !== businessID || selectedRecordID !== recordID) return;
    await loadNeighborhood();
    if (request !== recordRequest || selectedRecordID !== recordID) return;
    if (lens === 'history') await loadHistory(true);
  }

  async function loadNeighborhood() {
    if (!business || !selectedRecordID) {
      neighborhood = null;
      return;
    }
    const request = ++neighborhoodRequest;
    const businessID = business.id;
    const recordID = selectedRecordID;
    const kinds = [...activeRelationshipKinds];
    graphLoading = true;
    graphErrorCode = '';
    graphErrorMessage = '';
    try {
      const loadedNeighborhood = await BusinessService.Neighborhood(
        businessID,
        recordID,
        2,
        100,
        kinds,
      );
      if (request !== neighborhoodRequest || business?.id !== businessID || selectedRecordID !== recordID) return;
      neighborhood = loadedNeighborhood;
      if (loadedNeighborhood) void enrichNeighborhoodRecords(loadedNeighborhood, request, businessID, recordID);
    } catch (error) {
      if (request !== neighborhoodRequest || business?.id !== businessID || selectedRecordID !== recordID) return;
      const failure = serviceError(error);
      graphErrorCode = failure.code;
      graphErrorMessage = failure.message;
      neighborhood = null;
    } finally {
      if (request === neighborhoodRequest) graphLoading = false;
    }
  }

  async function enrichNeighborhoodRecords(loaded: Neighborhood, request: number, businessID: string, recordID: string) {
    const missingIDs = loaded.nodes
      .map((node) => node.id)
      .filter((nodeID) => !recordCache[nodeID]);
    if (missingIDs.length === 0) return;
    try {
      const views = await BusinessService.GetRecordsByIDs(businessID, missingIDs);
      if (request !== neighborhoodRequest || business?.id !== businessID || selectedRecordID !== recordID) return;
      rememberRecords(views);
    } catch {
      // Graph nodes still contain readable canonical labels when optional metadata enrichment fails.
    }
  }

  function toggleRelationshipKind(kind: string) {
    if (activeRelationshipKinds.includes(kind)) {
      const next = activeRelationshipKinds.filter((candidate) => candidate !== kind);
      activeRelationshipKinds = next.length ? next : [...allRelationshipKinds];
    } else {
      activeRelationshipKinds = [...activeRelationshipKinds, kind];
    }
    pathRequest++;
    pathResult = null;
    pathErrorCode = '';
    pathErrorMessage = '';
    void loadNeighborhood();
  }

  async function findPath() {
    if (!business || !selectedRecordID || !pathTargetID) return;
    const request = ++pathRequest;
    const businessID = business.id;
    const recordID = selectedRecordID;
    const targetID = pathTargetID;
    const kinds = [...activeRelationshipKinds];
    pathLoading = true;
    pathErrorCode = '';
    pathErrorMessage = '';
    pathResult = null;
    try {
      const loadedPath = await BusinessService.FindPath(
        businessID,
        recordID,
        targetID,
        6,
        kinds,
      );
      if (request !== pathRequest || business?.id !== businessID || selectedRecordID !== recordID || pathTargetID !== targetID) return;
      pathResult = loadedPath;
    } catch (error) {
      if (request !== pathRequest || business?.id !== businessID || selectedRecordID !== recordID || pathTargetID !== targetID) return;
      const failure = serviceError(error);
      pathErrorCode = failure.code;
      pathErrorMessage = failure.message;
    } finally {
      if (request === pathRequest) pathLoading = false;
    }
  }

  async function refreshProjection() {
    if (!business) return;
    const request = ++projectionRequest;
    const businessID = business.id;
    const wasPending = projection?.status === 'pending';
    projectionLoading = true;
    projectionError = '';
    try {
      const loadedProjection = await BusinessService.ProjectionStatus(businessID);
      if (request !== projectionRequest || business?.id !== businessID) return;
      projection = loadedProjection;
      if (wasPending && projection?.status === 'current' && selectedRecordID) {
        await loadNeighborhood();
      }
    } catch (error) {
      if (request !== projectionRequest || business?.id !== businessID) return;
      projection = null;
      projectionError = serviceError(error).message;
    } finally {
      if (request === projectionRequest) projectionLoading = false;
    }
  }

  async function setLens(nextLens: BusinessLens) {
    lens = nextLens;
    localStorage.setItem('a3t-business-lens', nextLens);
    if (nextLens === 'history') await loadHistory(true);
    if (nextLens === 'landscape' && selectedRecordID && !neighborhood) await loadNeighborhood();
  }

  async function loadHistory(reset: boolean) {
    if (!business || !selectedRecordID) {
      revisions = [];
      return;
    }
    const request = ++historyPageRequest;
    const businessID = business.id;
    const recordID = selectedRecordID;
    if (reset) {
      revisions = [];
      revisionsNextCursor = '';
      selectedRevisionID = '';
      historySnapshot = null;
      previousHistorySnapshot = null;
      historySelectionRequest++;
    }
    historyLoading = true;
    try {
      const page = await BusinessService.ListRevisionPage(
        businessID,
        recordID,
        reset ? '' : revisionsNextCursor,
        100,
      ) as RevisionPage | null;
      if (request !== historyPageRequest || business?.id !== businessID || selectedRecordID !== recordID) return;
      if (!page) throw new Error('Business service returned no revision page.');
      revisions = reset ? page.items : [...revisions, ...page.items];
      revisionsNextCursor = page.next_cursor || '';
      if (reset) {
        selectedRevisionID = revisions[0]?.id ?? '';
        if (selectedRevisionID) await selectRevision(selectedRevisionID);
      }
    } catch (error) {
      if (request !== historyPageRequest || business?.id !== businessID || selectedRecordID !== recordID) return;
      globalError = serviceError(error).message;
    } finally {
      if (request === historyPageRequest) historyLoading = false;
    }
  }

  async function selectRevision(revisionID: string) {
    if (!business || !selectedRecordID) return;
    const request = ++historySelectionRequest;
    const businessID = business.id;
    const recordID = selectedRecordID;
    selectedRevisionID = revisionID;
    historyLoading = true;
    try {
      const index = revisions.findIndex((revision) => revision.id === revisionID);
      let previousRevision = index >= 0 ? revisions[index + 1] : undefined;
      const selected = index >= 0 ? revisions[index] : undefined;
      if (!previousRevision && selected && selected.revision > 1 && revisionsNextCursor) {
        const page = await BusinessService.ListRevisionPage(
          businessID,
          recordID,
          revisionsNextCursor,
          100,
        ) as RevisionPage | null;
        if (request !== historySelectionRequest || business?.id !== businessID || selectedRecordID !== recordID || selectedRevisionID !== revisionID) return;
        if (page) {
          revisions = [...revisions, ...page.items];
          revisionsNextCursor = page.next_cursor || '';
          previousRevision = page.items[0];
        }
      }
      const [snapshot, previous] = await Promise.all([
        BusinessService.GetRevision(businessID, recordID, revisionID),
        previousRevision
          ? BusinessService.GetRevision(businessID, recordID, previousRevision.id)
          : Promise.resolve(null),
      ]);
      if (request !== historySelectionRequest || business?.id !== businessID || selectedRecordID !== recordID || selectedRevisionID !== revisionID) return;
      historySnapshot = snapshot;
      previousHistorySnapshot = previous;
    } catch (error) {
      if (request !== historySelectionRequest || business?.id !== businessID || selectedRecordID !== recordID || selectedRevisionID !== revisionID) return;
      globalError = serviceError(error).message;
    } finally {
      if (request === historySelectionRequest) historyLoading = false;
    }
  }

  function openCreateRecord() {
    recordEditorRecord = null;
    createRecordOperationID = crypto.randomUUID();
    recordEditorOpen = true;
  }

  function openEditRecord() {
    if (!selectedRecord) return;
    recordEditorRecord = selectedRecord;
    recordEditorOpen = true;
  }

  async function saveRecord(mutation: RecordMutation) {
    if (!business) return;
    try {
      let saved: RecordView | null;
      if (recordEditorRecord) {
        saved = await BusinessService.ReviseRecord(
          business.id,
          recordEditorRecord.record.id,
          recordEditorRecord.record.version,
          mutation as ReviseRecordInput,
        );
      } else {
        saved = await BusinessService.CreateRecord(
          business.id,
          { ...mutation, key: mutation.key || '' } as CreateRecordInput,
          createRecordOperationID,
        );
      }
      if (!saved) throw new Error('Business service returned no saved record.');
      selectedRecordID = saved.record.id;
      activeRecord = saved;
      rememberRecords([saved]);
      recordEditorOpen = false;
      await loadRecordPage(false);
      await refreshProjection();
      await loadNeighborhood();
      await setLens('record');
    } catch (error) {
      const failure = serviceError(error);
      if (failure.code === 'BUSINESS_PRECONDITION_FAILED' && recordEditorRecord) {
        const latest = await BusinessService.GetRecord(business.id, recordEditorRecord.record.id);
        if (latest) {
          activeRecord = latest;
          recordEditorRecord = latest;
          rememberRecords([latest]);
        }
        await loadRecordPage(false);
      }
      throw error;
    }
  }

  function openImport() {
    importOpen = true;
    importPath = '';
    importReport = null;
    importReviewID = '';
    importApplied = false;
  }

  async function reviewImport(path: string) {
    if (!business) return;
    importPath = path;
    importReport = null;
    importReviewID = '';
    importApplied = false;
    importLoading = true;
    try {
      const review = await BusinessService.ReviewImportSource(business.id, path) as ImportReviewResult | null;
      if (!review) throw new Error('Business service returned no import review.');
      importReport = review.report;
      importReviewID = review.review_id;
    } finally {
      importLoading = false;
    }
  }

  async function applyImport() {
    if (!business || !importReviewID) return;
    importLoading = true;
    try {
      importReport = await BusinessService.ApplyReviewedImport(business.id, importReviewID);
      importApplied = true;
      await loadRecordPage(false);
      await refreshProjection();
      if (selectedRecordID) await loadNeighborhood();
    } finally {
      importLoading = false;
    }
  }

  function openSettings() {
    configBaseURL = serviceStatus?.base_url || 'http://localhost:7290';
    configAccessToken = '';
    configTenantID = '';
    configActorID = '';
    configPreserveCredentials = Boolean(
      serviceStatus?.has_access_token || serviceStatus?.has_development_identity,
    );
    settingsError = '';
    settingsOpen = true;
  }

  function updateConfigBaseURL(event: Event) {
    configBaseURL = (event.currentTarget as HTMLInputElement).value;
    const comparable = (value: string) => value.trim().replace(/\/+$/, '').replace(/\/api\/v1$/, '');
    if (comparable(configBaseURL) !== comparable(serviceStatus?.base_url || '')) {
      configPreserveCredentials = false;
    }
  }

  async function saveSettings(event: SubmitEvent) {
    event.preventDefault();
    settingsSaving = true;
    settingsError = '';
    try {
      serviceStatus = await BusinessService.Configure({
        base_url: configBaseURL,
        access_token: configAccessToken,
        tenant_id: configTenantID,
        actor_id: configActorID,
        preserve_credentials: configPreserveCredentials,
      });
      if (!serviceStatus?.ready) {
        settingsError = serviceStatus?.message || 'The configured service is not ready.';
        return;
      }
      configAccessToken = '';
      configTenantID = '';
      configActorID = '';
      settingsOpen = false;
      clearBusinessState();
      await loadPortfolio(true);
    } catch (error) {
      settingsError = serviceError(error).message;
    } finally {
      settingsSaving = false;
    }
  }

  function openCreateBusiness() {
    newBusinessSlug = '';
    newBusinessName = '';
    newBusinessLegalName = '';
    newBusinessDescription = '';
    newBusinessOntology = '';
    createBusinessOperationID = crypto.randomUUID();
    createBusinessError = '';
    createBusinessOpen = true;
  }

  async function createBusiness(event: SubmitEvent) {
    event.preventDefault();
    createBusinessSaving = true;
    createBusinessError = '';
    try {
      const created = await BusinessService.CreateBusiness(
        {
          slug: newBusinessSlug,
          display_name: newBusinessName,
          legal_name: newBusinessLegalName,
          description: newBusinessDescription,
          status: 'draft',
          visibility: 'internal',
          ontology_version: newBusinessOntology,
        } as CreateBusinessInput,
        createBusinessOperationID,
      );
      if (!created) throw new Error('Business service returned no created business.');
      createBusinessOpen = false;
      await loadPortfolio(false);
      await selectBusiness(created.id);
    } catch (error) {
      createBusinessError = serviceError(error).message;
    } finally {
      createBusinessSaving = false;
    }
  }

  function statusBadgeVariant(status: string): 'success' | 'warn' | 'neutral' {
    if (status === 'active' || status === 'accepted') return 'success';
    if (status === 'draft' || status === 'review') return 'warn';
    return 'neutral';
  }
</script>

<svelte:head>
  <title>{business ? `${business.display_name} · A3T Business Workbench` : 'A3T Business Workbench'}</title>
  <meta name="description" content="Canonical business identity, records, history, and graph exploration." />
</svelte:head>

<main class="business-app">
  <header class="app-topbar">
    <div class="window-spacer" aria-hidden="true"></div>
    <button type="button" class="a3t-brand" onclick={showPortfolio} aria-label="Open business portfolio">
      <span class="a3t-mark">A3</span>
      <span>
        <strong>A3T</strong>
        <small>Business Workbench</small>
      </span>
    </button>

    <div class="topbar-context">
      {#if business}
        <button type="button" class="breadcrumb" onclick={showPortfolio}>Portfolio</button>
        <span>/</span>
        <strong>{business.display_name}</strong>
      {:else}
        <strong>Business portfolio</strong>
      {/if}
    </div>

    <div class="topbar-actions">
      <div class="service-indicator" title={serviceStatus?.message || 'Checking service'}>
        <Lamp state={serviceStatus?.ready ? 'running' : serviceStatus?.connected ? 'warn' : 'fail'} size={9} />
        <span>{serviceStatus?.ready ? 'Service ready' : serviceStatus?.connected ? 'Service degraded' : 'Service offline'}</span>
      </div>
      {#if business}
        <Button variant="outline" size="sm" onclick={openImport}>Import BSpec</Button>
      {/if}
      <Button variant="ghost" size="sm" onclick={openSettings}>Connection</Button>
    </div>
  </header>

  {#if globalError}
    <div class="global-banner">
      <span>{globalError}</span>
      <button type="button" onclick={() => globalError = ''} aria-label="Dismiss message">×</button>
    </div>
  {/if}

  {#if booting || businessLoading}
    <section class="app-loading" aria-live="polite">
      <span class="a3t-mark large">A3</span>
      <span class="spinner"></span>
      <strong>{businessLoading ? 'Opening business workbench' : 'Connecting to canonical business service'}</strong>
      <small>{serviceStatus?.base_url || 'http://localhost:7290'}</small>
    </section>
  {:else if !serviceStatus?.ready}
    <section class="service-offline">
      <div class="offline-orbit" aria-hidden="true"><span>A3</span></div>
      <span class="eyebrow">Connection required</span>
      <h1>Business service is unavailable</h1>
      <p>{serviceStatus?.message || globalError || 'Start the A3T business service or configure another endpoint.'}</p>
      <div>
        <Button onclick={bootstrap}>Retry</Button>
        <Button variant="outline" onclick={openSettings}>Configure connection</Button>
      </div>
      <code>{serviceStatus?.base_url || 'http://localhost:7290'}</code>
    </section>
  {:else if !business}
    <section class="portfolio">
      <header class="portfolio-hero">
        <div>
          <span class="eyebrow">Tenant portfolio</span>
          <h1>Managed businesses</h1>
          <p>Choose a canonical business root, or create one for a new living business model.</p>
        </div>
        <Button onclick={openCreateBusiness}>Create business</Button>
      </header>

      <div class="portfolio-summary">
        <div><span>Businesses</span><strong>{businesses.length}</strong></div>
        <div><span>Active</span><strong>{businesses.filter((item) => item.status === 'active').length}</strong></div>
        <div><span>Draft</span><strong>{businesses.filter((item) => item.status === 'draft').length}</strong></div>
        <div><span>Service</span><strong>{serviceStatus.auth_mode}</strong></div>
      </div>

      {#if portfolioLoading}
        <div class="graph-loading"><span class="spinner"></span>Loading tenant businesses</div>
      {:else if businesses.length === 0}
        <div class="portfolio-empty">
          <Empty
            title="No managed businesses"
            description="Create a business root before importing BSpec records or building its knowledge graph."
          />
          <Button onclick={openCreateBusiness}>Create first business</Button>
        </div>
      {:else}
        <div class="business-grid">
          {#each businesses as item (item.id)}
            <button type="button" class="business-card" onclick={() => void selectBusiness(item.id)}>
              <header>
                <span class="business-monogram">{item.display_name.slice(0, 2).toUpperCase()}</span>
                <Badge variant={statusBadgeVariant(item.status)}>{item.status}</Badge>
              </header>
              <h2>{item.display_name}</h2>
              <p>{item.description || item.legal_name || 'Canonical business identity and structured knowledge.'}</p>
              <dl>
                <div><dt>Slug</dt><dd>{item.slug}</dd></div>
                <div><dt>Ontology</dt><dd>{item.ontology_version || 'Not set'}</dd></div>
                <div><dt>Updated</dt><dd>{formatDate(item.updated_at)}</dd></div>
              </dl>
              <span class="card-open">Open workbench <b>→</b></span>
            </button>
          {/each}
        </div>
      {/if}
    </section>
  {:else}
    <section class="business-shell">
      <header class="business-contextbar">
        <div class="business-title">
          <span class="business-monogram compact">{business.display_name.slice(0, 2).toUpperCase()}</span>
          <div>
            <span class="eyebrow">{business.slug}</span>
            <h1>{business.display_name}</h1>
            {#if primaryIdentifiers.length > 0}
              <div class="header-identifiers" aria-label="Primary business identifiers">
                {#each primaryIdentifiers as identifier (identifier.id)}
                  <span title={`${identifier.scheme}: ${identifier.value}`}>{identifier.scheme}: {identifier.value}</span>
                {/each}
              </div>
            {/if}
          </div>
        </div>
        <div class="business-context-items">
          <div><span>Lifecycle</span><Badge variant={statusBadgeVariant(business.status)}>{business.status}</Badge></div>
          <div><span>Visibility</span><strong>{business.visibility}</strong></div>
          <div><span>Ontology</span><strong>{business.ontology_version || 'Not set'}</strong></div>
          <div>
            <span>Graph projection</span>
            {#if projectionLoading}
              <Badge variant="neutral">Checking</Badge>
            {:else if projectionPending}
              <Badge variant="warn">{projection?.pending_events ?? 0} pending</Badge>
            {:else if projection?.status === 'current'}
              <Badge variant="success">Outbox current</Badge>
            {:else}
              <span title={projectionError || 'Projection status has not been loaded'}><Badge variant="danger">Unavailable</Badge></span>
            {/if}
          </div>
        </div>
      </header>

      <div class="workbench-shell">
        <aside class="record-index">
          <header>
            <div>
              <span class="eyebrow">Canonical index</span>
              <h2>Records</h2>
            </div>
            <Badge variant="neutral">{records.length}{recordsNextCursor ? '+' : ''}</Badge>
          </header>

          <div class="record-index-controls">
            <SearchInput
              bind:value={recordQuery}
              aria-label="Search business records"
              placeholder="Search key or title"
              showShortcut={false}
              oninput={scheduleRecordSearch}
              onclear={() => void loadRecordPage(false)}
            />
            <div class="record-filter-row">
              <NativeSelect bind:value={recordTypeFilter} options={recordTypeOptions} size="sm" aria-label="Filter records by type" onchange={changeRecordFilter} />
              <NativeSelect bind:value={recordStatusFilter} options={recordStatusOptions} size="sm" aria-label="Filter records by status" onchange={changeRecordFilter} />
            </div>
          </div>

          <nav class="record-list" aria-label="Business records">
            {#each records as view (view.record.id)}
              <button
                type="button"
                class:active={selectedRecordID === view.record.id}
                aria-current={selectedRecordID === view.record.id ? 'true' : undefined}
                onclick={() => void selectRecord(view.record.id)}
              >
                <span class="record-type">{view.record.type_code}</span>
                <span>
                  <strong>{view.revision.title}</strong>
                  <small>{view.record.key}</small>
                </span>
                <i class={`status-dot status-${view.revision.status || 'none'}`} title={view.revision.status || 'No status'}></i>
              </button>
            {:else}
              {#if !recordsLoading}
                <div class="record-list-empty">
                  <span>No records match this index view.</span>
                  {#if !recordQuery && !recordTypeFilter && !recordStatusFilter}
                    <Button size="sm" onclick={openCreateRecord}>Create record</Button>
                  {/if}
                </div>
              {/if}
            {/each}
          </nav>

          {#if recordsLoading}
            <div class="index-loading"><span class="spinner"></span>Loading records</div>
          {:else if recordsNextCursor}
            <Button variant="ghost" size="sm" onclick={() => void loadRecordPage(true)}>Load more</Button>
          {/if}

          <footer>
            <Button onclick={openCreateRecord}>New record</Button>
            <small>PostgreSQL canonical state</small>
          </footer>
        </aside>

        <section class="workbench-main">
          <nav class="lens-navigation" aria-label="Business workbench lenses">
            <button type="button" class:active={lens === 'landscape'} aria-current={lens === 'landscape' ? 'page' : undefined} onclick={() => void setLens('landscape')}>
              <span>Landscape</span><small>Orient and trace impact</small>
            </button>
            <button type="button" class:active={lens === 'record'} aria-current={lens === 'record' ? 'page' : undefined} onclick={() => void setLens('record')}>
              <span>Record</span><small>Read and revise meaning</small>
            </button>
            <button type="button" class:active={lens === 'history'} aria-current={lens === 'history' ? 'page' : undefined} onclick={() => void setLens('history')}>
              <span>History</span><small>Audit immutable change</small>
            </button>
            <button type="button" class:active={lens === 'identity'} aria-current={lens === 'identity' ? 'page' : undefined} onclick={() => void setLens('identity')}>
              <span>Identity</span><small>Understand the root</small>
            </button>
          </nav>

          <div class="lens-content">
            {#if lens === 'landscape'}
              <LandscapeLens
                {neighborhood}
                records={cachedRecords}
                {selectedRecordID}
                loading={graphLoading}
                errorCode={graphErrorCode}
                errorMessage={graphErrorMessage}
                {projectionPending}
                {pathResult}
                {pathTargetID}
                {pathTargetQuery}
                pathCandidates={pathTargetRecords}
                {pathTargetsLoading}
                {pathLoading}
                pathErrorCode={pathErrorCode}
                pathErrorMessage={pathErrorMessage}
                activeKinds={activeRelationshipKinds}
                onPathTargetChange={(recordID) => {
                  pathRequest++;
                  pathTargetID = recordID;
                  pathResult = null;
                  pathErrorCode = '';
                  pathErrorMessage = '';
                }}
                onPathTargetSearch={schedulePathTargetSearch}
                onToggleKind={toggleRelationshipKind}
                onFindPath={() => void findPath()}
                onRetry={() => void loadNeighborhood()}
                onSelectRecord={(recordID) => void selectRecord(recordID)}
              />
            {:else if lens === 'record'}
              <RecordLens
                record={selectedRecord}
                {incomingRelationships}
                {projectionPending}
                graphContextUnavailable={Boolean(graphErrorMessage)}
                graphContextPartial={projectionPending || Boolean(neighborhood?.truncated) || activeRelationshipKinds.length !== allRelationshipKinds.length}
                onCreate={openCreateRecord}
                onEdit={openEditRecord}
                onSelectRecord={(recordID) => void selectRecord(recordID)}
              />
            {:else if lens === 'history'}
              <HistoryLens
                {revisions}
                {selectedRevisionID}
                snapshot={historySnapshot}
                previousSnapshot={previousHistorySnapshot}
                loading={historyLoading}
                nextCursor={revisionsNextCursor}
                onSelectRevision={(revisionID) => void selectRevision(revisionID)}
                onLoadMore={() => void loadHistory(false)}
              />
            {:else}
              <IdentityLens {business} {identifiers} />
            {/if}
          </div>
        </section>
      </div>
    </section>
  {/if}

  {#if settingsOpen}
    <div class="modal-backdrop" role="presentation">
      <div class="modal connection-modal" role="dialog" aria-modal="true" aria-labelledby="connection-title" tabindex="-1" use:focusTrap={{ onEscape: () => { if (!settingsSaving) settingsOpen = false; } }}>
        <header class="modal-header">
          <div>
            <span class="eyebrow">Go-side service boundary</span>
            <h2 id="connection-title">Business service connection</h2>
            <p>Credentials are handed to the native client and are never returned through status or read bindings.</p>
          </div>
          <button class="modal-close" type="button" onclick={() => settingsOpen = false} aria-label="Close connection settings">×</button>
        </header>
        <form class="connection-form" onsubmit={saveSettings}>
          {#if settingsError}<div class="state-banner state-banner-danger">{settingsError}</div>{/if}
          <label>
            <span>Service URL</span>
            <input value={configBaseURL} oninput={updateConfigBaseURL} type="url" required placeholder="http://localhost:7290" />
            <small>Use the service origin. A trailing /api/v1 is normalized automatically.</small>
          </label>
          <label>
            <span>Bearer token or API key</span>
            <input bind:value={configAccessToken} type="password" autocomplete="off" placeholder={serviceStatus?.has_access_token ? 'Stored credential is hidden' : 'Optional for local Compose'} />
          </label>
          {#if serviceStatus?.has_access_token || serviceStatus?.has_development_identity}
            <label class="credential-preserve">
              <input bind:checked={configPreserveCredentials} type="checkbox" />
              <span>Keep the currently configured hidden credential and development identity when their fields are blank.</span>
            </label>
            <small class="credential-preserve-help">Turn this off and leave the fields blank to clear them when connecting.</small>
          {/if}
          <div class="connection-dev-fields">
            <label>
              <span>Development tenant ID</span>
              <input bind:value={configTenantID} placeholder="Optional local override" />
            </label>
            <label>
              <span>Development actor ID</span>
              <input bind:value={configActorID} placeholder="Optional local override" />
            </label>
          </div>
          <div class="state-banner state-banner-neutral">
            Tenant and actor identifiers are request headers only in development mode. They are never placed in REST bodies, queries, or MCP inputs.
          </div>
          <footer class="modal-actions">
            <Button variant="ghost" onclick={() => settingsOpen = false} disabled={settingsSaving}>Cancel</Button>
            <Button type="submit" disabled={settingsSaving}>{settingsSaving ? 'Connecting…' : 'Connect'}</Button>
          </footer>
        </form>
      </div>
    </div>
  {/if}

  {#if createBusinessOpen}
    <div class="modal-backdrop" role="presentation">
      <div class="modal create-business-modal" role="dialog" aria-modal="true" aria-labelledby="create-business-title" tabindex="-1" use:focusTrap={{ onEscape: () => { if (!createBusinessSaving) createBusinessOpen = false; } }}>
        <header class="modal-header">
          <div>
            <span class="eyebrow">New managed root</span>
            <h2 id="create-business-title">Create business</h2>
            <p>A tenant may own multiple businesses. This does not create a tenant or organization.</p>
          </div>
          <button class="modal-close" type="button" onclick={() => createBusinessOpen = false} aria-label="Close create business dialog">×</button>
        </header>
        <form class="create-business-form" onsubmit={createBusiness}>
          {#if createBusinessError}<div class="state-banner state-banner-danger">{createBusinessError}</div>{/if}
          <div class="editor-grid">
            <label>
              <span>Display name</span>
              <input bind:value={newBusinessName} required maxlength="200" placeholder="Agent Three" />
            </label>
            <label>
              <span>Slug</span>
              <input bind:value={newBusinessSlug} required maxlength="63" placeholder="agent-three" />
            </label>
            <label class="editor-span-2">
              <span>Legal name</span>
              <input bind:value={newBusinessLegalName} placeholder="Optional legal identity" />
            </label>
            <label class="editor-span-2">
              <span>Description</span>
              <textarea bind:value={newBusinessDescription} placeholder="What this managed business is and does."></textarea>
            </label>
            <label>
              <span>Ontology version</span>
              <input bind:value={newBusinessOntology} placeholder="1.0.0" />
            </label>
          </div>
          <footer class="modal-actions">
            <Button variant="ghost" onclick={() => createBusinessOpen = false} disabled={createBusinessSaving}>Cancel</Button>
            <Button type="submit" disabled={createBusinessSaving}>{createBusinessSaving ? 'Creating…' : 'Create draft business'}</Button>
          </footer>
        </form>
      </div>
    </div>
  {/if}

  {#if recordEditorOpen}
    <RecordEditor record={recordEditorRecord} onSave={saveRecord} onClose={() => recordEditorOpen = false} />
  {/if}

  {#if importOpen && business}
    <ImportReview
      {business}
      report={importReport}
      inputPath={importPath}
      loading={importLoading}
      applied={importApplied}
      onReview={reviewImport}
      onApply={applyImport}
      onClose={() => importOpen = false}
    />
  {/if}
</main>
