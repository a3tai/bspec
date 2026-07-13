import type {
  GraphEdge,
  GraphNode,
  Neighborhood,
  RecordView,
} from '../../bindings/github.com/a3tai/bspec/apps/graph';

export type BusinessLens = 'landscape' | 'record' | 'history' | 'identity';

export type ServiceError = {
  code: string;
  message: string;
};

export type VisualGraphNode = {
  id: string;
  key: string;
  title: string;
  typeCode: string;
  status: string;
  domain: string;
  x: number;
  y: number;
  tone: number;
  unresolved: boolean;
};

export type VisualGraphEdge = {
  id: string;
  sourceID: string;
  targetID: string;
  kind: string;
  strength: string;
  unresolved: boolean;
  source: VisualGraphNode;
  target: VisualGraphNode;
  path: string;
  labelX: number;
  labelY: number;
};

export type GraphScene = {
  nodes: VisualGraphNode[];
  edges: VisualGraphEdge[];
  domains: string[];
  width: number;
  height: number;
};

export type IncomingRelationshipView = {
  id: string;
  sourceID: string;
  sourceKey: string;
  sourceTitle: string;
  kind: string;
  strength: string;
};

export type RecordMutation = {
  key?: string;
  type_code: string;
  revision: {
    title: string;
    status: string;
    domain: string;
    visibility: string;
    owner_ref: string;
    schema_version: string;
    source_version: string;
    data_json: string;
    narrative_markdown: string;
    raw_markdown: string;
    source_uri: string;
    source_media_type: string;
    effective_at: string | null;
    relationships: Array<{
      kind: string;
      target_key: string;
      strength: string;
      provenance_json?: string;
    }>;
  };
};

const serviceCodePattern = /\b(BUSINESS_[A-Z_]+)\s*:\s*/;

export function serviceError(error: unknown): ServiceError {
  const message = error instanceof Error ? error.message : String(error);
  const match = serviceCodePattern.exec(message);
  return {
    code: match?.[1] ?? '',
    message: match ? message.replace(serviceCodePattern, '').trim() : message,
  };
}

export function formatDate(value: string | null | undefined, includeTime = false): string {
  if (!value) return 'Not set';
  const date = new Date(value);
  if (Number.isNaN(date.getTime())) return value;
  return new Intl.DateTimeFormat(undefined, {
    dateStyle: 'medium',
    ...(includeTime ? { timeStyle: 'short' as const } : {}),
  }).format(date);
}

export function titleCase(value: string): string {
  return value
    .replaceAll('_', ' ')
    .replaceAll('-', ' ')
    .replace(/\b\w/g, (letter) => letter.toUpperCase());
}

export function shortID(value: string | null | undefined): string {
  if (!value) return 'None';
  return value.length > 16 ? `${value.slice(0, 8)}…${value.slice(-4)}` : value;
}

export function dataAsJSON(value: string | null | undefined): string {
  return value?.trim() || '{}';
}

export function buildGraphScene(neighborhood: Neighborhood | null, records: RecordView[]): GraphScene {
  const sceneWidth = 1000;
  if (!neighborhood) return { nodes: [], edges: [], domains: [], width: sceneWidth, height: 540 };

  const recordByID = new Map(records.map((view) => [view.record.id, view]));
  const grouped = new Map<string, GraphNode[]>();
  for (const node of neighborhood.nodes) {
    const domain = recordByID.get(node.id)?.revision.domain?.trim() || 'Unspecified';
    const group = grouped.get(domain) ?? [];
    group.push(node);
    grouped.set(domain, group);
  }

  const domains = [...grouped.keys()].sort((left, right) => left.localeCompare(right));
  const nodeByID = new Map<string, VisualGraphNode>();
  const columns = 6;
  const columnWidth = 145;
  const rowHeight = 78;
  let nextDomainY = 30;

  domains.forEach((domain) => {
    const domainNodes = [...(grouped.get(domain) ?? [])].sort((left, right) => left.key.localeCompare(right.key));
    const rows = Math.max(1, Math.ceil(domainNodes.length / columns));

    domainNodes.forEach((node, index) => {
      const column = index % columns;
      const row = Math.floor(index / columns);
      nodeByID.set(node.id, {
        id: node.id,
        key: node.key,
        title: node.title,
        typeCode: node.type_code,
        status: node.status || 'unspecified',
        domain,
        x: 135 + column * columnWidth,
        y: nextDomainY + 68 + row * rowHeight,
        tone: hashString(domain) % 6,
        unresolved: false,
      });
    });
    nextDomainY += Math.max(145, rows * rowHeight + 65);
  });

  const unresolvedEdges = neighborhood.edges
    .filter((edge) => (!edge.target_id || !nodeByID.has(edge.target_id)) && edge.target_ref && nodeByID.has(edge.source_id))
    .sort((left, right) => `${left.target_ref}:${left.id}`.localeCompare(`${right.target_ref}:${right.id}`));
  if (unresolvedEdges.length > 0) {
    domains.push('Unresolved');
    const rows = Math.ceil(unresolvedEdges.length / columns);
    unresolvedEdges.forEach((edge, index) => {
      nodeByID.set(`unresolved:${edge.id}`, {
        id: `unresolved:${edge.id}`,
        key: edge.target_ref || 'Unknown reference',
        title: 'Unresolved reference',
        typeCode: 'REF',
        status: 'unresolved',
        domain: 'Unresolved',
        x: 135 + (index % columns) * columnWidth,
        y: nextDomainY + 68 + Math.floor(index / columns) * rowHeight,
        tone: 5,
        unresolved: true,
      });
    });
    nextDomainY += Math.max(145, rows * rowHeight + 65);
  }

  const edges: VisualGraphEdge[] = [];
  for (const edge of neighborhood.edges) {
    const source = nodeByID.get(edge.source_id);
    if (!source) continue;
    let target = edge.target_id ? nodeByID.get(edge.target_id) : undefined;
    let unresolved = false;

    if (!target && edge.target_ref) {
      unresolved = true;
      target = nodeByID.get(`unresolved:${edge.id}`);
    }

    if (!target) continue;
    edges.push({
      id: edge.id,
      sourceID: source.id,
      targetID: target.id,
      kind: edge.kind,
      strength: edge.strength || '',
      unresolved,
      source,
      target,
      path: '',
      labelX: 0,
      labelY: 0,
    });
  }

  const edgeGroups = new Map<string, VisualGraphEdge[]>();
  for (const edge of edges) {
    const pair = [edge.sourceID, edge.targetID].sort();
    const key = `${pair[0]}:${pair[1]}`;
    const group = edgeGroups.get(key) ?? [];
    group.push(edge);
    edgeGroups.set(key, group);
  }
  for (const group of edgeGroups.values()) {
    group.sort((left, right) => `${left.kind}:${left.id}`.localeCompare(`${right.kind}:${right.id}`));
    const canonical = [group[0].source, group[0].target].sort((left, right) => left.id.localeCompare(right.id));
    const deltaX = canonical[1].x - canonical[0].x;
    const deltaY = canonical[1].y - canonical[0].y;
    const length = Math.max(1, Math.hypot(deltaX, deltaY));
    const perpendicularX = -deltaY / length;
    const perpendicularY = deltaX / length;
    group.forEach((edge, index) => {
      const offset = (index - (group.length - 1) / 2) * 24;
      const midpointX = (edge.source.x + edge.target.x) / 2;
      const midpointY = (edge.source.y + edge.target.y) / 2;
      const controlX = midpointX + perpendicularX * offset * 2;
      const controlY = midpointY + perpendicularY * offset * 2;
      edge.path = offset === 0
        ? `M ${edge.source.x} ${edge.source.y} L ${edge.target.x} ${edge.target.y}`
        : `M ${edge.source.x} ${edge.source.y} Q ${controlX} ${controlY} ${edge.target.x} ${edge.target.y}`;
      edge.labelX = (edge.source.x + 2 * controlX + edge.target.x) / 4;
      edge.labelY = (edge.source.y + 2 * controlY + edge.target.y) / 4 - 7;
    });
  }

  return {
    nodes: [...nodeByID.values()],
    edges,
    domains,
    width: sceneWidth,
    height: Math.max(540, nextDomainY + 25),
  };
}

export function relationshipTarget(edge: GraphEdge): string {
  return edge.target_id ?? edge.target_ref ?? '';
}

function hashString(value: string): number {
  let hash = 2166136261;
  for (let index = 0; index < value.length; index += 1) {
    hash ^= value.charCodeAt(index);
    hash = Math.imul(hash, 16777619);
  }
  return hash >>> 0;
}
