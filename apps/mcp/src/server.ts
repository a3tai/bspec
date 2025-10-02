/**
 * Ultra-Simple BSpec Specification File Server
 *
 * Loads TGZ from R2 and serves raw specification files via MCP
 */

// Use Cloudflare Workers built-in compression APIs

// BSpec specification index with intelligence
export interface BSpecIndex {
  version: string;
  files: Map<string, string>; // filename -> content
  documentTypes: Map<string, DocumentTypeSpec>; // type code -> spec
  domains: Map<string, DomainInfo>; // domain -> info
  readme: string; // main README content
  spec: string; // main specification content
}

export interface DocumentTypeSpec {
  code: string; // 3-letter code (MSN, VSN, etc.)
  name: string;
  domain: string;
  description: string;
  purpose: string;
  content: string; // full specification content
}

export interface DomainInfo {
  name: string;
  description: string;
  documentTypes: string[]; // array of type codes
  totalTypes: number;
}

let bspecIndex: BSpecIndex | null = null;

/**
 * Load TGZ from R2 and index all files
 */
export async function loadSpecificationFromR2(env: any): Promise<BSpecIndex> {
  const startTime = Date.now();
  console.log('[TIMING] Loading BSpec specification from R2...');

  try {
    // Load the standard TGZ file
    const fetchStart = Date.now();
    const object = await env.ASSETS.get('bspec-v1-0-0.tgz');
    console.log(`[TIMING] R2 fetch completed in ${Date.now() - fetchStart}ms`);
    
    if (!object) {
      console.error('ERROR: BSpec specification file not found in R2 bucket');
      console.log('Expected file: bspec-v1-0-0.tgz');
      throw new Error('BSpec specification file not found in R2 bucket. Please ensure bspec-v1-0-0.tgz is uploaded to the ASSETS bucket.');
    }

    const bufferStart = Date.now();
    console.log('[TIMING] Found TGZ file in R2, reading buffer...');
    const arrayBuffer = await object.arrayBuffer();
    console.log(`[TIMING] Buffer read completed in ${Date.now() - bufferStart}ms`);

    if (!arrayBuffer || arrayBuffer.byteLength === 0) {
      console.error('ERROR: TGZ file is empty or corrupted');
      throw new Error('BSpec specification file is empty or corrupted');
    }

    const uint8Array = new Uint8Array(arrayBuffer);
    console.log(`[TIMING] TGZ file loaded, size: ${uint8Array.length} bytes`);

    // Decompress the gzip file using Cloudflare Workers built-in streams
    const readable = new ReadableStream({
      start(controller) {
        controller.enqueue(uint8Array);
        controller.close();
      }
    });

    const decompressStart = Date.now();
    console.log('[TIMING] Decompressing TGZ file...');
    const decompressed = await new Response(
      readable.pipeThrough(new DecompressionStream('gzip'))
    ).arrayBuffer();
    console.log(`[TIMING] Decompression completed in ${Date.now() - decompressStart}ms`);

    if (!decompressed || decompressed.byteLength === 0) {
      console.error('ERROR: Decompression failed or resulted in empty data');
      throw new Error('Failed to decompress BSpec specification file');
    }

    const decompressedBytes = new Uint8Array(decompressed);
    console.log(`Decompressed size: ${decompressedBytes.length} bytes`);

    // Parse TAR file (simple implementation)
    const parseStart = Date.now();
    console.log('[TIMING] Parsing TAR contents...');
    const files = new Map<string, string>();
    let version = '1.0.0'; // default

    let offset = 0;
    let fileCount = 0;
    while (offset < decompressedBytes.length) {
    // TAR header is 512 bytes
    if (offset + 512 > decompressedBytes.length) break;

    const header = decompressedBytes.slice(offset, offset + 512);

    // Extract filename (first 100 bytes, null-terminated)
    let nameEnd = 0;
    while (nameEnd < 100 && header[nameEnd] !== 0) nameEnd++;
    const filename = new TextDecoder().decode(header.slice(0, nameEnd));

    if (!filename) break;

    // Extract file size (bytes 124-135, octal)
    const sizeStr = new TextDecoder().decode(header.slice(124, 135)).replace(/\0/g, '').trim();
    const size = parseInt(sizeStr, 8) || 0;

    offset += 512; // Skip header

      if (size > 0) {
        const content = new TextDecoder().decode(decompressedBytes.slice(offset, offset + size));
        files.set(filename, content);
        fileCount++;

        // Check for version.txt
        if (filename === 'version.txt' || filename.endsWith('/version.txt')) {
          version = content.trim();
          console.log(`Found version: ${version}`);
        }

        // Round up to next 512-byte boundary
        offset += Math.ceil(size / 512) * 512;
      }
    }

    console.log(`[TIMING] TAR parsing completed in ${Date.now() - parseStart}ms`);
    
    if (files.size === 0) {
      console.error('ERROR: No files found in TGZ archive');
      throw new Error('BSpec specification archive is empty or corrupted');
    }

    console.log(`[TIMING] Successfully loaded ${files.size} files from BSpec v${version}`);

    // Build intelligent BSpec index
    const indexStart = Date.now();
    console.log('[TIMING] Building BSpec intelligence index...');
    const bspecIndex = await buildBSpecIndex(files, version);
    console.log(`[TIMING] Index built in ${Date.now() - indexStart}ms`);

    const totalTime = Date.now() - startTime;
    console.log(`[TIMING] ✅ Total specification load time: ${totalTime}ms`);
    
    return bspecIndex;

  } catch (error) {
    console.error('ERROR: Failed to load BSpec specification from R2:', error);
    throw error instanceof Error ? error : new Error('Unknown error loading BSpec specification');
  }
}

/**
 * Build intelligent BSpec index from loaded files
 */
async function buildBSpecIndex(files: Map<string, string>, version: string): Promise<BSpecIndex> {
  const buildStart = Date.now();
  const documentTypes = new Map<string, DocumentTypeSpec>();
  const domains = new Map<string, DomainInfo>();
  let parsedFiles = 0;

  // Get main README and spec content
  const readme = files.get('README.md') || files.get('README.json') || '';
  const spec = files.get('spec.json') || '';

  // Domain mapping based on BSpec 1.0 specification
  const domainMapping = {
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
    'brand-marketing': 'Brand & Marketing',
    'learning-decisions': 'Learning & Decisions'
  };

  // Process each domain directory
  for (const [domainDir, domainName] of Object.entries(domainMapping)) {
    const domainTypes: string[] = [];

    // Find all document type specs in this domain
    for (const [filename, content] of files.entries()) {
      if (filename.startsWith(domainDir + '/') && filename.endsWith('-spec.json')) {
        try {
          const jsonContent = JSON.parse(content);
          const specContent = jsonContent.content || '';

          // Extract document type code from filename (e.g., "MSN-spec.json" -> "MSN")
          const typeCode = filename.split('/')[1].split('-spec.json')[0].toUpperCase();

          // Parse the document type info from the specification content
          const docType = parseDocumentTypeSpec(typeCode, domainDir, specContent);
          documentTypes.set(typeCode, docType);
          domainTypes.push(typeCode);
          parsedFiles++;

        } catch (error) {
          console.warn(`[TIMING] Failed to parse document type spec: ${filename}`, error);
        }
      }
    }

    if (domainTypes.length > 0) {
      domains.set(domainDir, {
        name: domainName,
        description: getDomainDescription(domainDir),
        documentTypes: domainTypes.sort(),
        totalTypes: domainTypes.length
      });
    }
  }

  const buildTime = Date.now() - buildStart;
  console.log(`[TIMING] Built BSpec index: ${documentTypes.size} document types across ${domains.size} domains (parsed ${parsedFiles} files) in ${buildTime}ms`);

  return {
    version,
    files,
    documentTypes,
    domains,
    readme,
    spec
  };
}

/**
 * Parse document type specification from content
 */
function parseDocumentTypeSpec(code: string, domain: string, content: string): DocumentTypeSpec {
  // Extract name and purpose from markdown content
  const lines = content.split('\n');
  let name = code;
  let description = '';
  let purpose = '';

  for (let i = 0; i < lines.length; i++) {
    const line = lines[i].trim();

    // Look for document type name
    if (line.startsWith('**Document Type Name:**')) {
      name = line.replace('**Document Type Name:**', '').trim();
    }

    // Look for purpose section
    if (line.startsWith('## Purpose and Scope')) {
      // Get the next few lines as description
      for (let j = i + 1; j < Math.min(i + 5, lines.length); j++) {
        const nextLine = lines[j].trim();
        if (nextLine && !nextLine.startsWith('#') && !nextLine.startsWith('**')) {
          description += nextLine + ' ';
        }
      }
      break;
    }
  }

  return {
    code,
    name: name || `${code} Document`,
    domain: domain.replace('-', ' ').replace(/\b\w/g, l => l.toUpperCase()),
    description: description.trim() || `${code} document specification`,
    purpose: purpose || description.trim(),
    content
  };
}

/**
 * Get domain description based on BSpec specification
 */
function getDomainDescription(domain: string): string {
  const descriptions: { [key: string]: string } = {
    'strategic-foundation': 'Core organizational purpose and direction',
    'market-environment': 'External market context and competitive landscape',
    'customer-value': 'Customer insights and value propositions',
    'product-service': 'Product and service definitions and specifications',
    'business-model': 'Revenue generation and value delivery mechanisms',
    'operations-execution': 'Operational processes and organizational structure',
    'technology-data': 'Technical architecture and data management',
    'financial-investment': 'Financial planning and investment management',
    'risk-governance': 'Risk management and governance frameworks',
    'growth-innovation': 'Innovation strategies and growth initiatives',
    'brand-marketing': 'Brand strategy and marketing execution',
    'learning-decisions': 'Organizational learning and decision frameworks'
  };

  return descriptions[domain] || 'Business domain specification';
}

/**
 * Create intelligent BSpec MCP server with Resources
 */
export function createBSpecMCPServer(index: BSpecIndex) {
  return {
    name: 'bspec-specification-server',
    version: index.version,

    listResources() {
      const resources = [];

      // BSpec Overview Resource
      resources.push({
        uri: `https://mcp.bspec.dev/spec/v${index.version}/overview`,
        name: `BSpec v${index.version} Overview`,
        description: `Complete overview of BSpec ${index.version} specification with ${index.documentTypes.size} document types across ${index.domains.size} domains`,
        mimeType: 'text/markdown'
      });

      // Domain Resources
      for (const [domainKey, domainInfo] of index.domains.entries()) {
        resources.push({
          uri: `https://mcp.bspec.dev/spec/v${index.version}/domains/${domainKey}`,
          name: `Domain: ${domainInfo.name}`,
          description: `${domainInfo.description} - Contains ${domainInfo.totalTypes} document types`,
          mimeType: 'text/markdown'
        });
      }

      // Document Type Resources
      for (const [typeCode, docType] of index.documentTypes.entries()) {
        resources.push({
          uri: `https://mcp.bspec.dev/spec/v${index.version}/document-types/${typeCode}`,
          name: `${typeCode}: ${docType.name}`,
          description: `${docType.description} (${docType.domain} domain)`,
          mimeType: 'text/markdown'
        });
      }

      return resources;
    },

    async readResource(uri: string) {
      const url = new URL(uri);

      // Validate that this is a mcp.bspec.dev URI
      if (url.hostname !== 'mcp.bspec.dev' || url.protocol !== 'https:') {
        return {
          contents: [{
            uri,
            text: `Error: Only https://mcp.bspec.dev/* URIs are supported. Received: ${uri}`,
            mimeType: 'text/plain'
          }]
        };
      }

      const path = url.pathname;

      try {
        // Handle overview
        if (path === `/spec/v${index.version}/overview`) {
          const content = await this.callTool('get_bspec_overview');
          return {
            contents: [{
              uri,
              text: content.content[0].text,
              mimeType: 'text/markdown'
            }]
          };
        }

        // Handle domain resources
        if (path.startsWith(`/spec/v${index.version}/domains/`)) {
          const domainKey = path.substring(`/spec/v${index.version}/domains/`.length);
          const content = await this.callTool('get_domain_info', { domain: domainKey });

          // Check if the tool call was successful
          if (content.content[0].text.startsWith('Error:')) {
            return {
              contents: [{
                uri,
                text: content.content[0].text,
                mimeType: 'text/plain'
              }]
            };
          }

          return {
            contents: [{
              uri,
              text: content.content[0].text,
              mimeType: 'text/markdown'
            }]
          };
        }

        // Handle document type resources
        if (path.startsWith(`/spec/v${index.version}/document-types/`)) {
          const typeCode = path.substring(`/spec/v${index.version}/document-types/`.length);
          const content = await this.callTool('get_document_type', { type_code: typeCode });

          // Check if the tool call was successful
          if (content.content[0].text.startsWith('Error:')) {
            return {
              contents: [{
                uri,
                text: content.content[0].text,
                mimeType: 'text/plain'
              }]
            };
          }

          return {
            contents: [{
              uri,
              text: content.content[0].text,
              mimeType: 'text/markdown'
            }]
          };
        }

        return {
          contents: [{
            uri,
            text: `Error: Resource not found: ${uri}. Available paths: /spec/v${index.version}/overview, /spec/v${index.version}/domains/{key}, /spec/v${index.version}/document-types/{code}`,
            mimeType: 'text/plain'
          }]
        };

      } catch (error) {
        return {
          contents: [{
            uri,
            text: `Error reading resource: ${error instanceof Error ? error.message : 'Unknown error'}`,
            mimeType: 'text/plain'
          }]
        };
      }
    },

    listTools() {
      return [
        // BSpec Intelligence Tools
        {
          name: 'get_document_type',
          description: 'Get detailed information about a specific BSpec document type',
          inputSchema: {
            type: 'object',
            properties: {
              type_code: { type: 'string', description: 'Three-letter document type code (e.g., "MSN", "VSN", "STR")' }
            },
            required: ['type_code']
          }
        },
        {
          name: 'list_document_types',
          description: 'List all BSpec document types with filtering options',
          inputSchema: {
            type: 'object',
            properties: {
              domain: { type: 'string', description: 'Filter by domain (e.g., "strategic-foundation", "market-environment")' },
              include_content: { type: 'boolean', description: 'Include full specification content (default: false)' }
            }
          }
        },
        {
          name: 'get_domain_info',
          description: 'Get information about a specific BSpec business domain',
          inputSchema: {
            type: 'object',
            properties: {
              domain: { type: 'string', description: 'Domain name (e.g., "strategic-foundation", "customer-value")' }
            },
            required: ['domain']
          }
        },
        {
          name: 'list_domains',
          description: 'List all BSpec business domains with their document types',
          inputSchema: {
            type: 'object',
            properties: {
              include_types: { type: 'boolean', description: 'Include document types in each domain (default: true)' }
            }
          }
        },
        {
          name: 'get_bspec_overview',
          description: 'Get a comprehensive overview of the BSpec specification',
          inputSchema: { type: 'object' }
        },
        {
          name: 'search_document_types',
          description: 'Search for document types by name, description, or purpose',
          inputSchema: {
            type: 'object',
            properties: {
              query: { type: 'string', description: 'Search query to match against document type names, descriptions, or purposes' }
            },
            required: ['query']
          }
        },
        {
          name: 'get_document_relationships',
          description: 'Get information about how document types typically relate to each other',
          inputSchema: {
            type: 'object',
            properties: {
              type_code: { type: 'string', description: 'Document type code to get relationships for' }
            },
            required: ['type_code']
          }
        },
      ];
    },

    async callTool(name: string, args: any = {}) {
      switch (name) {
        // BSpec Intelligence Tools
        case 'get_document_type':
          const typeCode = args.type_code?.toUpperCase();
          if (!typeCode) {
            return { content: [{ type: 'text', text: 'Error: type_code required' }] };
          }

          const docType = index.documentTypes.get(typeCode);
          if (!docType) {
            return { content: [{ type: 'text', text: `Error: Document type "${typeCode}" not found. Use list_document_types to see all available types.` }] };
          }

          return {
            content: [{
              type: 'text',
              text: `# ${docType.code}: ${docType.name}\n\n**Domain:** ${docType.domain}\n**Description:** ${docType.description}\n\n**Full Specification:**\n\n${docType.content}`
            }]
          };

        case 'list_document_types':
          const domainFilter = args.domain;
          const includeContent = args.include_content || false;

          let filteredTypes = Array.from(index.documentTypes.values());
          if (domainFilter) {
            filteredTypes = filteredTypes.filter(type =>
              type.domain.toLowerCase().includes(domainFilter.toLowerCase()) ||
              type.code.toLowerCase().includes(domainFilter.toLowerCase())
            );
          }

          let response = `# BSpec Document Types (${filteredTypes.length} types)\n\n`;

          for (const type of filteredTypes.sort((a, b) => a.code.localeCompare(b.code))) {
            response += `## ${type.code}: ${type.name}\n`;
            response += `**Domain:** ${type.domain}\n`;
            response += `**Description:** ${type.description}\n`;

            if (includeContent) {
              response += `\n**Full Specification:**\n${type.content}\n`;
            }
            response += '\n---\n\n';
          }

          return { content: [{ type: 'text', text: response }] };

        case 'get_domain_info':
          const domain = args.domain;
          if (!domain) {
            return { content: [{ type: 'text', text: 'Error: domain required' }] };
          }

          const domainInfo = index.domains.get(domain);
          if (!domainInfo) {
            const availableDomains = Array.from(index.domains.keys()).join(', ');
            return { content: [{ type: 'text', text: `Error: Domain "${domain}" not found. Available domains: ${availableDomains}` }] };
          }

          let domainResponse = `# ${domainInfo.name}\n\n`;
          domainResponse += `**Description:** ${domainInfo.description}\n`;
          domainResponse += `**Total Document Types:** ${domainInfo.totalTypes}\n\n`;
          domainResponse += `## Document Types in this Domain:\n\n`;

          for (const typeCode of domainInfo.documentTypes) {
            const type = index.documentTypes.get(typeCode);
            if (type) {
              domainResponse += `- **${typeCode}**: ${type.name} - ${type.description}\n`;
            }
          }

          return { content: [{ type: 'text', text: domainResponse }] };

        case 'list_domains':
          const includeTypes = args.include_types !== false; // default true

          let domainsResponse = `# BSpec Business Domains (${index.domains.size} domains)\n\n`;

          for (const [domainKey, domainInfo] of index.domains.entries()) {
            domainsResponse += `## ${domainInfo.name}\n`;
            domainsResponse += `**Key:** ${domainKey}\n`;
            domainsResponse += `**Description:** ${domainInfo.description}\n`;
            domainsResponse += `**Document Types:** ${domainInfo.totalTypes}\n`;

            if (includeTypes) {
              domainsResponse += `\n**Types:** ${domainInfo.documentTypes.join(', ')}\n`;
            }
            domainsResponse += '\n---\n\n';
          }

          return { content: [{ type: 'text', text: domainsResponse }] };

        case 'get_bspec_overview':
          const totalTypes = index.documentTypes.size;
          const totalDomains = index.domains.size;

          let overview = `# BSpec ${index.version} Specification Overview\n\n`;
          overview += `**Universal Business Specification Standard**\n\n`;
          overview += `- **Version:** ${index.version}\n`;
          overview += `- **Document Types:** ${totalTypes}\n`;
          overview += `- **Business Domains:** ${totalDomains}\n\n`;

          overview += `## Business Domains:\n\n`;
          for (const [domainKey, domainInfo] of index.domains.entries()) {
            overview += `- **${domainInfo.name}** (${domainInfo.totalTypes} types): ${domainInfo.description}\n`;
          }

          overview += `\n## Key Features:\n\n`;
          overview += `- **Atomic Documents**: Each document covers exactly one business concern\n`;
          overview += `- **Rich Relationships**: Documents declare dependencies and enablements\n`;
          overview += `- **Machine + Human Readable**: YAML frontmatter + Markdown content\n`;
          overview += `- **Quality Levels**: Bronze (minimum), Silver (investment ready), Gold (operational excellence)\n`;
          overview += `- **Business Knowledge Graph**: Connected documents form comprehensive business picture\n\n`;

          overview += `Use the other tools to explore specific document types, domains, or search for specific capabilities.`;

          return { content: [{ type: 'text', text: overview }] };

        case 'search_document_types':
          const query = args.query?.toLowerCase();
          if (!query) {
            return { content: [{ type: 'text', text: 'Error: query required' }] };
          }

          const matchingTypes = Array.from(index.documentTypes.values()).filter(type =>
            type.name.toLowerCase().includes(query) ||
            type.description.toLowerCase().includes(query) ||
            type.purpose.toLowerCase().includes(query) ||
            type.code.toLowerCase().includes(query)
          );

          if (matchingTypes.length === 0) {
            return { content: [{ type: 'text', text: `No document types found matching "${query}". Try broader search terms or use list_document_types to see all types.` }] };
          }

          let searchResponse = `# Search Results for "${query}" (${matchingTypes.length} matches)\n\n`;

          for (const type of matchingTypes.sort((a, b) => a.code.localeCompare(b.code))) {
            searchResponse += `## ${type.code}: ${type.name}\n`;
            searchResponse += `**Domain:** ${type.domain}\n`;
            searchResponse += `**Description:** ${type.description}\n\n`;
          }

          return { content: [{ type: 'text', text: searchResponse }] };

        case 'get_document_relationships':
          const relTypeCode = args.type_code?.toUpperCase();
          if (!relTypeCode) {
            return { content: [{ type: 'text', text: 'Error: type_code required' }] };
          }

          const relDocType = index.documentTypes.get(relTypeCode);
          if (!relDocType) {
            return { content: [{ type: 'text', text: `Error: Document type "${relTypeCode}" not found.` }] };
          }

          // Extract relationship information from the content
          const relationshipInfo = extractRelationshipInfo(relDocType.content);

          let relResponse = `# ${relTypeCode}: ${relDocType.name} - Relationships\n\n`;
          relResponse += `**Domain:** ${relDocType.domain}\n\n`;
          relResponse += relationshipInfo;

          return { content: [{ type: 'text', text: relResponse }] };


        default:
          return { content: [{ type: 'text', text: `Unknown tool: ${name}. Use get_bspec_overview to see available capabilities.` }] };
      }
    }
  };
}

/**
 * Extract relationship information from document content
 */
function extractRelationshipInfo(content: string): string {
  const lines = content.split('\n');
  let relationshipInfo = '';
  let inRelationshipSection = false;

  for (let i = 0; i < lines.length; i++) {
    const line = lines[i].trim();

    // Look for relationship sections
    if (line.includes('Relationship') || line.includes('Dependencies') || line.includes('Enablements')) {
      inRelationshipSection = true;
      relationshipInfo += `${line}\n`;
      continue;
    }

    // Look for typical dependencies/enablements patterns
    if (line.includes('depends_on:') || line.includes('enables:') || line.includes('conflicts_with:')) {
      relationshipInfo += `${line}\n`;
      continue;
    }

    // Look for common relationship keywords
    if (inRelationshipSection && line) {
      if (line.startsWith('#') && !line.includes('Relationship')) {
        inRelationshipSection = false;
      } else {
        relationshipInfo += `${line}\n`;
      }
    }
  }

  if (!relationshipInfo) {
    relationshipInfo = 'Relationship information not explicitly documented in this specification. ';
    relationshipInfo += 'Relationships are typically defined when creating actual business documents based on business context and dependencies.';
  }

  return relationshipInfo;
}

/**
 * Initialize server with loaded specification
 */
export async function initializeServer(env: any) {
  if (!bspecIndex) {
    bspecIndex = await loadSpecificationFromR2(env);
  }
  return createBSpecMCPServer(bspecIndex);
}