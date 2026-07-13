import { execFileSync } from 'node:child_process';
import {
  existsSync,
  mkdirSync,
  mkdtempSync,
  readdirSync,
  readFileSync,
  rmSync,
  statSync,
  writeFileSync,
} from 'node:fs';
import { tmpdir } from 'node:os';
import { basename, dirname, extname, join, relative, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';

const __dirname = dirname(fileURLToPath(import.meta.url));
const appRoot = resolve(__dirname, '..');
const repoRoot = resolve(appRoot, '..', '..');
const defaultSpecDir = resolve(repoRoot, 'spec', 'v1');
const defaultOutputPath = resolve(appRoot, 'src', 'data', 'graph.json');

const relationshipFields = ['depends_on', 'enables', 'conflicts_with', 'related', 'supersedes'];

const relationshipLabels = {
  depends_on: 'DEPENDS_ON',
  enables: 'ENABLES',
  conflicts_with: 'CONFLICTS_WITH',
  related: 'RELATED_TO',
  supersedes: 'SUPERSEDES',
};

const domainNames = {
  'strategic-foundation': 'Strategic Foundation',
  'market-environment': 'Market & Environment',
  'customer-value': 'Customer & Value',
  'product-service': 'Product & Service',
  'business-model': 'Business Model',
  'operations-execution': 'Operations & Execution',
  'technology-data': 'Technology & Data',
  'financial-investment': 'Financial & Investment',
  'risk-governance': 'Risk & Governance',
  'growth-innovation': 'Growth & Innovation',
  'learning-decisions': 'Learning & Decisions',
  'brand-marketing': 'Brand & Marketing',
  unspecified: 'Unspecified',
};

const domainDescriptions = {
  'strategic-foundation': 'Core purpose, vision, values, and strategic direction',
  'market-environment': 'Market analysis, competitive landscape, and external factors',
  'customer-value': 'Customer understanding, journey mapping, and value delivery',
  'product-service': 'Product specifications, features, and service delivery',
  'business-model': 'Revenue streams, cost structure, and business mechanics',
  'operations-execution': 'Processes, organization, roles, and operational capabilities',
  'technology-data': 'Technical architecture, systems, data, and infrastructure',
  'financial-investment': 'Financial planning, budgets, forecasts, and investment strategy',
  'risk-governance': 'Risk management, compliance, governance, and controls',
  'growth-innovation': 'Growth strategies, innovation, R&D, and expansion',
  'learning-decisions': 'Decision frameworks, learning systems, and knowledge management',
  'brand-marketing': 'Brand positioning, marketing strategy, and communications',
  unspecified: 'Documents without a declared BSpec domain',
};

const domainAliases = {
  strategic: 'strategic-foundation',
  market: 'market-environment',
  customer: 'customer-value',
  product: 'product-service',
  model: 'business-model',
  business: 'business-model',
  operations: 'operations-execution',
  technology: 'technology-data',
  financial: 'financial-investment',
  finance: 'financial-investment',
  risk: 'risk-governance',
  growth: 'growth-innovation',
  learning: 'learning-decisions',
  brand: 'brand-marketing',
  marketing: 'brand-marketing',
};

function parseArgs(argv) {
  const args = {
    input: defaultSpecDir,
    output: defaultOutputPath,
    mode: 'auto',
    title: '',
  };

  for (let index = 0; index < argv.length; index += 1) {
    const arg = argv[index];
    const [key, inlineValue] = arg.split('=', 2);
    const nextValue = inlineValue ?? argv[index + 1];

    if (key === '--input') {
      args.input = nextValue;
      if (inlineValue === undefined) index += 1;
    } else if (key === '--output') {
      args.output = nextValue;
      if (inlineValue === undefined) index += 1;
    } else if (key === '--mode') {
      args.mode = nextValue;
      if (inlineValue === undefined) index += 1;
    } else if (key === '--title') {
      args.title = nextValue;
      if (inlineValue === undefined) index += 1;
    } else if (arg === '--help' || arg === '-h') {
      printHelp();
      process.exit(0);
    }
  }

  if (!['auto', 'catalog', 'project'].includes(args.mode)) {
    throw new Error(`Unsupported --mode ${args.mode}. Use auto, catalog, or project.`);
  }

  return {
    input: resolve(process.cwd(), args.input),
    output: resolve(process.cwd(), args.output),
    mode: args.mode,
    title: args.title,
  };
}

function printHelp() {
  console.log(`Usage: bun scripts/build-graph-data.js [options]

Options:
  --input <path>     spec/v1, extracted BSpec project, documents dir, or .bspec file
  --output <path>    output graph JSON path
  --mode <mode>      auto, catalog, or project
  --title <title>    override workbench title
`);
}

function stableId(...parts) {
  return parts
    .filter(Boolean)
    .join(':')
    .toLowerCase()
    .replace(/[^a-z0-9]+/g, '-')
    .replace(/^-|-$/g, '') || 'item';
}

function compactText(value, limit = 900) {
  const compacted = String(value ?? '').replace(/\s+/g, ' ').trim();
  return compacted.length > limit ? `${compacted.slice(0, limit - 1).trim()}...` : compacted;
}

function displayPath(path) {
  const absolute = resolve(path);
  const relativePath = relative(repoRoot, absolute);
  return relativePath.startsWith('..') ? absolute : relativePath;
}

function slugify(value) {
  return String(value ?? 'unspecified')
    .trim()
    .toLowerCase()
    .replace(/&/g, 'and')
    .replace(/[^a-z0-9]+/g, '-')
    .replace(/^-|-$/g, '') || 'unspecified';
}

function normalizeDomain(value) {
  const slug = slugify(value);
  const canonical = domainAliases[slug] ?? slug;
  return {
    slug: domainNames[canonical] ? canonical : canonical || 'unspecified',
    name: domainNames[canonical] ?? String(value || 'Unspecified'),
  };
}

function readMarkdownFiles(dir, predicate = () => true, relativeDir = '') {
  const files = [];

  for (const item of readdirSync(dir)) {
    if (item === '.DS_Store' || item.startsWith('._')) continue;

    const fullPath = join(dir, item);
    const itemPath = relativeDir ? join(relativeDir, item) : item;

    if (statSync(fullPath).isDirectory()) {
      files.push(...readMarkdownFiles(fullPath, predicate, itemPath));
    } else if (item.endsWith('.md') && predicate(itemPath)) {
      files.push({
        path: itemPath,
        fullPath,
        content: readFileSync(fullPath, 'utf8'),
      });
    }
  }

  return files;
}

function extractField(content, label) {
  const match = content.match(new RegExp(`^\\*\\*${label}:\\*\\*\\s*(.+?)\\s*$`, 'm'));
  return match ? match[1].trim() : '';
}

function extractSection(content, heading) {
  const pattern = new RegExp(`^##\\s+${heading}\\s*$\\n([\\s\\S]*?)(?=^##\\s+|\\Z)`, 'gm');
  const values = [];
  let match;

  while ((match = pattern.exec(content)) !== null) {
    const value = compactText(match[1]);
    if (value && !values.includes(value)) {
      values.push(value);
    }
  }

  return values.join('\n\n');
}

function extractMetadataSchemaBlock(content) {
  const lines = content.split('\n');
  const headingIndex = lines.findIndex((line) => line.trim() === '## Document Metadata Schema');
  if (headingIndex === -1) return '';

  const fenceStart = lines.findIndex(
    (line, index) => index > headingIndex && line.trim().startsWith('```'),
  );
  if (fenceStart === -1) return '';

  const block = [];
  for (let index = fenceStart + 1; index < lines.length; index += 1) {
    if (lines[index].trim() === '```') break;
    block.push(lines[index]);
  }

  return block
    .join('\n')
    .trim()
    .replace(/^---\s*/, '')
    .replace(/\s*---$/, '')
    .trim();
}

function normalizeScalar(value) {
  return String(value ?? '').trim().replace(/^['"]|['"]$/g, '').trim();
}

function parseInlineList(value) {
  const normalized = normalizeScalar(value);
  const inner = normalized.startsWith('[') && normalized.endsWith(']')
    ? normalized.slice(1, -1)
    : normalized;

  return inner
    .split(',')
    .map(normalizeScalar)
    .filter(Boolean);
}

function splitFrontmatter(content) {
  if (!content.startsWith('---\n')) {
    return { metadata: {}, markdown: content };
  }

  const parts = content.slice(4).split('\n---\n');
  if (parts.length < 2) {
    return { metadata: {}, markdown: content };
  }

  return {
    metadata: parseFrontmatter(parts[0]),
    markdown: parts.slice(1).join('\n---\n').trim(),
  };
}

function parseFrontmatter(frontmatter) {
  const metadata = {};
  let currentKey = '';

  for (const rawLine of frontmatter.split('\n')) {
    if (!rawLine.trim() || rawLine.trimStart().startsWith('#')) continue;

    if (!rawLine.startsWith(' ') && !rawLine.startsWith('\t') && rawLine.includes(':')) {
      const separator = rawLine.indexOf(':');
      const key = rawLine.slice(0, separator).trim();
      const value = rawLine.slice(separator + 1).trim();
      currentKey = key;

      if (!value) {
        metadata[key] = [];
      } else if (value.startsWith('[') && value.endsWith(']')) {
        metadata[key] = parseInlineList(value);
      } else {
        metadata[key] = normalizeScalar(value);
      }
      continue;
    }

    if (currentKey && rawLine.trim().startsWith('- ')) {
      const existing = Array.isArray(metadata[currentKey])
        ? metadata[currentKey]
        : metadata[currentKey]
          ? [metadata[currentKey]]
          : [];
      existing.push(normalizeScalar(rawLine.trim().slice(2)));
      metadata[currentKey] = existing;
    }
  }

  return metadata;
}

function parseCatalogRelationshipToken(token) {
  const pattern = normalizeScalar(token);
  if (!pattern || pattern === '[]') return null;

  const code = pattern.match(/^([A-Z]{3})/)?.[1];
  return code ? { code, pattern, documentId: null } : null;
}

function parseProjectRelationshipToken(token) {
  const pattern = normalizeScalar(token);
  if (!pattern || pattern === '[]') return null;

  const upperCode = pattern.match(/^([A-Z]{3})(?:[-_*]|$)/)?.[1];
  const lowerCode = pattern.match(/^([a-z]{3})(?:[-_]|$)/)?.[1]?.toUpperCase();

  return {
    code: upperCode ?? lowerCode ?? '',
    pattern,
    documentId: pattern.includes('*') ? null : pattern,
  };
}

function parseRelationshipValue(value, parser) {
  if (Array.isArray(value)) {
    return value.map((item) => parser(String(item))).filter(Boolean);
  }

  const normalized = normalizeScalar(value).replace(/\s+#.*$/, '');
  if (!normalized || normalized === '[]') return [];
  return parseInlineList(normalized).map(parser).filter(Boolean);
}

function extractCatalogRelationships(content) {
  const schemaBlock = extractMetadataSchemaBlock(content);
  const relationships = Object.fromEntries(relationshipFields.map((field) => [field, []]));

  for (const line of schemaBlock.split('\n')) {
    const match = line.match(/^\s*(depends_on|enables|conflicts_with|related|supersedes):\s*(.*?)\s*(?:#.*)?$/);
    if (!match) continue;

    relationships[match[1]] = parseRelationshipValue(match[2], parseCatalogRelationshipToken);
  }

  return relationships;
}

function extractProjectRelationships(metadata) {
  return Object.fromEntries(
    relationshipFields.map((field) => [field, parseRelationshipValue(metadata[field], parseProjectRelationshipToken)]),
  );
}

function loadVersion(specDir) {
  const versionPath = resolve(specDir, 'version.txt');
  try {
    return readFileSync(versionPath, 'utf8').trim();
  } catch {
    return '1.0.0';
  }
}

function buildCatalogGraph(inputPath, titleOverride) {
  const files = readMarkdownFiles(inputPath, (path) => path.endsWith('-spec.md')).sort((a, b) =>
    a.path.localeCompare(b.path),
  );
  const version = loadVersion(inputPath);

  const documents = files.map((file) => {
    const [domainSlugFromPath, filename] = file.path.split('/');
    const pathDomain = normalizeDomain(domainSlugFromPath);
    const code = extractField(file.content, 'Document Type Code') || filename.replace('-spec.md', '');
    const name = extractField(file.content, 'Document Type Name') || code;
    const domainName = extractField(file.content, 'Domain') || pathDomain.name;
    const normalizedDomain = { slug: pathDomain.slug, name: domainName };
    const relationships = extractCatalogRelationships(file.content);
    const outgoingCount = relationshipFields.reduce((sum, field) => sum + relationships[field].length, 0);

    return {
      id: stableId('document-type', normalizedDomain.slug, code),
      kind: 'bspec.document_type',
      documentId: code,
      code,
      name,
      domainSlug: normalizedDomain.slug,
      domainName: normalizedDomain.name,
      status: extractField(file.content, 'Status') || 'Draft',
      version: extractField(file.content, 'Version') || '1.0.0',
      lastUpdated: extractField(file.content, 'Last Updated') || '2025-09-30',
      owner: 'BSpec Foundation',
      sourcePath: `${displayPath(inputPath)}/${file.path}`,
      abstract: extractSection(file.content, 'Abstract'),
      purpose: extractSection(file.content, 'Purpose and Scope'),
      relationshipGuidance: compactText(
        [extractSection(file.content, 'Relationship Guidelines'), extractSection(file.content, 'Document Relationships')]
          .filter(Boolean)
          .join('\n\n'),
      ),
      relationships,
      outgoingCount,
    };
  });

  const documentByCode = new Map();
  for (const document of documents) {
    if (!documentByCode.has(document.code)) {
      documentByCode.set(document.code, []);
    }
    documentByCode.get(document.code).push(document);
  }

  const edges = [];
  for (const document of documents) {
    for (const field of relationshipFields) {
      for (const target of document.relationships[field]) {
        const targetDocument = documentByCode.get(target.code)?.[0];
        edges.push(buildEdge(document, target, targetDocument, field));
      }
    }
  }

  return finalizeGraph({
    sourceKind: 'catalog',
    title: titleOverride || 'BSpec Catalog Graphiti Workbench',
    description: 'Interactive Mittsu prototype for exploring BSpec document types as Graphiti-ready graph data.',
    version,
    groupId: `bspec-v${version}`,
    sourceRoot: displayPath(inputPath),
    graphitiCoreEpisodes: countUnique(documents, 'domainSlug') + documents.length,
    graphitiRelationshipRuleEpisodes: countUnique(documents, 'domainSlug') + documents.length + edges.length,
    documents,
    edges,
    nodeLabelSingular: 'Document type',
    nodeLabelPlural: 'Document types',
    domainHeading: 'BSpec Domains',
  });
}

function materializeProjectInput(inputPath) {
  if (statSync(inputPath).isDirectory()) {
    return {
      root: inputPath,
      label: displayPath(inputPath),
      cleanup: () => {},
    };
  }

  if (extname(inputPath) !== '.bspec') {
    throw new Error(`Project input must be a directory or .bspec file: ${inputPath}`);
  }

  const tempRoot = mkdtempSync(join(tmpdir(), 'bspec-ui-'));
  execFileSync('tar', ['-xzf', inputPath, '-C', tempRoot], { stdio: 'ignore' });

  return {
    root: tempRoot,
    label: displayPath(inputPath),
    cleanup: () => rmSync(tempRoot, { recursive: true, force: true }),
  };
}

function buildProjectGraph(inputPath, titleOverride) {
  const materialized = materializeProjectInput(inputPath);

  try {
    const manifestPath = join(materialized.root, 'manifest.json');
    const hasManifest = existsSync(manifestPath);
    const manifest = hasManifest ? JSON.parse(readFileSync(manifestPath, 'utf8')) : {};
    const documentsRoot = hasManifest && existsSync(join(materialized.root, 'documents'))
      ? join(materialized.root, 'documents')
      : materialized.root;
    const files = readMarkdownFiles(documentsRoot).sort((a, b) => a.path.localeCompare(b.path));
    const projectName = manifest.name || basename(inputPath, extname(inputPath)) || 'BSpec Project';
    const version = manifest.bspec_version || manifest.bspecVersion || '1.0.0';

    const documents = files.map((file) => {
      const { metadata, markdown } = splitFrontmatter(file.content);
      const docType = String(metadata.type || basename(file.path).split('-')[0] || 'DOC').toUpperCase();
      const documentId = String(metadata.id || basename(file.path, extname(file.path)));
      const domain = normalizeDomain(metadata.domain || 'unspecified');
      const relationships = extractProjectRelationships(metadata);
      const outgoingCount = relationshipFields.reduce((sum, field) => sum + relationships[field].length, 0);
      const sourcePath = extname(inputPath) === '.bspec'
        ? `${materialized.label}!/documents/${file.path}`
        : displayPath(file.fullPath);

      return {
        id: stableId('document', documentId, sourcePath),
        kind: 'bspec.document',
        documentId,
        code: docType,
        name: String(metadata.title || documentId),
        domainSlug: domain.slug,
        domainName: domain.name,
        status: String(metadata.status || 'Draft'),
        version: String(metadata.version || '1.0.0'),
        lastUpdated: String(metadata.updated || metadata.created || ''),
        owner: String(metadata.owner || 'Unknown'),
        sourcePath,
        abstract: '',
        purpose: compactText(markdown, 900),
        relationshipGuidance: '',
        relationships,
        outgoingCount,
      };
    });

    const documentById = new Map();
    const documentByCode = new Map();
    for (const document of documents) {
      documentById.set(document.documentId, document);
      documentById.set(document.documentId.toLowerCase(), document);
      if (!documentByCode.has(document.code)) {
        documentByCode.set(document.code, []);
      }
      documentByCode.get(document.code).push(document);
    }

    const edges = [];
    for (const document of documents) {
      for (const field of relationshipFields) {
        for (const target of document.relationships[field]) {
          const targetDocument = target.documentId
            ? documentById.get(target.documentId) ?? documentById.get(target.documentId.toLowerCase())
            : documentByCode.get(target.code)?.[0];
          edges.push(buildEdge(document, target, targetDocument, field));
        }
      }
    }

    return finalizeGraph({
      sourceKind: 'project',
      title: titleOverride || `${projectName} Graphiti Workbench`,
      description: 'Interactive Mittsu prototype for exploring an existing BSpec project as Graphiti-ready graph data.',
      version,
      groupId: `bspec-${stableId(projectName)}`,
      sourceRoot: materialized.label,
      graphitiCoreEpisodes: 1 + documents.length,
      graphitiRelationshipRuleEpisodes: 1 + documents.length + edges.length,
      documents,
      edges,
      manifest,
      nodeLabelSingular: 'Document',
      nodeLabelPlural: 'Documents',
      domainHeading: 'Project Domains',
    });
  } finally {
    materialized.cleanup();
  }
}

function buildEdge(document, target, targetDocument, field) {
  return {
    id: stableId('edge', document.id, field, target.pattern),
    source: document.id,
    sourceCode: document.code,
    sourceDocumentId: document.documentId,
    sourceName: document.name,
    sourceDomainSlug: document.domainSlug,
    target: targetDocument?.id ?? stableId('external', target.pattern),
    targetCode: targetDocument?.code ?? target.code ?? target.pattern,
    targetDocumentId: targetDocument?.documentId ?? target.documentId ?? '',
    targetName: targetDocument?.name ?? target.pattern,
    targetDomainSlug: targetDocument?.domainSlug ?? 'external',
    targetPattern: target.pattern,
    field,
    label: relationshipLabels[field],
    fact: `${document.documentId || document.code}: ${document.name} ${relationshipLabels[field]} ${target.pattern}`,
  };
}

function countUnique(items, key) {
  return new Set(items.map((item) => item[key]).filter(Boolean)).size;
}

function buildDomains(documents, edges) {
  const domainMap = new Map();

  for (const document of documents) {
    if (!domainMap.has(document.domainSlug)) {
      domainMap.set(document.domainSlug, {
        id: stableId('domain', document.domainSlug),
        slug: document.domainSlug,
        name: document.domainName,
        description: domainDescriptions[document.domainSlug] ?? '',
        documentCount: 0,
        outgoingEdgeCount: 0,
        incomingEdgeCount: 0,
      });
    }
    domainMap.get(document.domainSlug).documentCount += 1;
  }

  for (const edge of edges) {
    if (domainMap.has(edge.sourceDomainSlug)) {
      domainMap.get(edge.sourceDomainSlug).outgoingEdgeCount += 1;
    }
    if (domainMap.has(edge.targetDomainSlug)) {
      domainMap.get(edge.targetDomainSlug).incomingEdgeCount += 1;
    }
  }

  return [...domainMap.values()].sort((a, b) => a.name.localeCompare(b.name));
}

function finalizeGraph({
  sourceKind,
  title,
  description,
  version,
  groupId,
  sourceRoot,
  graphitiCoreEpisodes,
  graphitiRelationshipRuleEpisodes,
  documents,
  edges,
  manifest = null,
  nodeLabelSingular,
  nodeLabelPlural,
  domainHeading,
}) {
  const domains = buildDomains(documents, edges);
  const relationCounts = relationshipFields.map((field) => ({
    field,
    label: relationshipLabels[field],
    count: edges.filter((edge) => edge.field === field).length,
  }));

  return {
    schemaVersion: 1,
    sourceKind,
    title,
    description,
    generatedAt: new Date().toISOString(),
    version,
    groupId,
    sourceRoot,
    manifest,
    nodeLabelSingular,
    nodeLabelPlural,
    domainHeading,
    graphitiCoreEpisodes,
    graphitiRelationshipRuleEpisodes,
    domains,
    documents,
    edges,
    relationCounts,
    domainSeries: domains.map((domain) => domain.documentCount + domain.outgoingEdgeCount),
  };
}

function detectMode(inputPath, requestedMode) {
  if (requestedMode !== 'auto') return requestedMode;
  if (!statSync(inputPath).isDirectory()) return 'project';
  if (existsSync(join(inputPath, 'manifest.json')) || existsSync(join(inputPath, 'documents'))) return 'project';
  if (existsSync(join(inputPath, 'version.txt')) || existsSync(join(inputPath, 'spec.md'))) return 'catalog';
  return 'project';
}

const args = parseArgs(process.argv.slice(2));
if (!existsSync(args.input)) {
  throw new Error(`Input does not exist: ${args.input}`);
}

const mode = detectMode(args.input, args.mode);
const graphData = mode === 'catalog'
  ? buildCatalogGraph(args.input, args.title)
  : buildProjectGraph(args.input, args.title);

mkdirSync(dirname(args.output), { recursive: true });
writeFileSync(args.output, `${JSON.stringify(graphData, null, 2)}\n`);

console.log(`Wrote ${args.output}`);
console.log(`${graphData.documents.length} ${graphData.nodeLabelPlural.toLowerCase()}, ${graphData.edges.length} relationships`);
console.log(`Source: ${graphData.sourceRoot} (${graphData.sourceKind})`);
