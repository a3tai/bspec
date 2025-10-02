/**
 * BSpec MCP Agent using Cloudflare Agents Framework
 * 
 * Provides access to the Business Specification Standard (BSpec) v1.0 through 
 * tools and resources using the official Cloudflare Agents SDK.
 */

import { McpServer } from "@modelcontextprotocol/sdk/server/mcp.js";
import { McpAgent } from "agents/mcp";
import { z } from "zod";
import { 
  BSpecIndex, 
  loadSpecificationFromR2, 
  DocumentTypeSpec, 
  DomainInfo 
} from './server';

// Environment interface for Cloudflare Workers
interface Env {
  ASSETS: R2Bucket;
  BSpecMcpAgent: DurableObjectNamespace<BSpecMcpAgent>;
  ENVIRONMENT?: string;
  MCP_SERVER_VERSION?: string;
  CORS_ORIGIN?: string;
}

// State interface for the agent
interface BSpecState {
  initialized: boolean;
  version: string;
  lastLoadTime: number;
}

/**
 * BSpec MCP Agent implementing the Cloudflare Agents framework
 */
export class BSpecMcpAgent extends McpAgent<Env, BSpecState, {}> {
  server = new McpServer({
    name: "bspec-specification-server",
    version: "1.0.0"
  });

  // Initial state
  initialState: BSpecState = {
    initialized: false,
    version: "1.0.0",
    lastLoadTime: 0
  };

  // Cache for the loaded BSpec specification
  private bspecIndex: BSpecIndex | null = null;

  async init() {
    console.log('Initializing BSpec MCP Agent...');
    
    // DON'T load specification during init - do it lazily on first request
    // This prevents timeout during initialization
    
    // Set up tools immediately (they're lightweight)
    this.setupTools();

    // Update state to indicate initialization
    this.setState({
      ...this.state,
      initialized: true,
      lastLoadTime: 0 // Will be updated on first load
    });

    console.log('BSpec MCP Agent initialized (lazy loading enabled)');
  }

  onStateUpdate(state: BSpecState) {
    console.log('BSpec Agent state updated:', { 
      initialized: state.initialized, 
      version: state.version,
      lastLoad: new Date(state.lastLoadTime).toISOString()
    });
  }

  /**
   * Ensure the BSpec specification is loaded from R2 with timeout protection
   */
  private async ensureSpecificationLoaded(): Promise<void> {
    if (this.bspecIndex) {
      // Already loaded
      return;
    }

    // Check if we're already loading (prevent duplicate loads)
    if (this._loadingPromise) {
      return this._loadingPromise;
    }

    console.log('Loading BSpec specification from R2...');
    
    // Store the loading promise to prevent concurrent loads
    this._loadingPromise = this._loadSpecWithTimeout();
    
    try {
      await this._loadingPromise;
    } finally {
      this._loadingPromise = null;
    }
  }

  private _loadingPromise: Promise<void> | null = null;

  /**
   * Load specification with timeout protection
   */
  private async _loadSpecWithTimeout(): Promise<void> {
    const startTime = Date.now();
    const TIMEOUT_MS = 8000; // 8 seconds (Cloudflare Workers have 10s CPU limit)
    
    console.log('[TIMING] Starting specification load...');
    
    const timeoutPromise = new Promise<never>((_, reject) => {
      setTimeout(() => reject(new Error('Specification loading timed out')), TIMEOUT_MS);
    });

    try {
      this.bspecIndex = await Promise.race([
        loadSpecificationFromR2(this.env),
        timeoutPromise
      ]);

      const loadTime = Date.now() - startTime;
      console.log(`[TIMING] ✅ Loaded BSpec ${this.bspecIndex.version} with ${this.bspecIndex.documentTypes.size} document types in ${loadTime}ms`);
      
      // Update version in state
      this.setState({
        ...this.state,
        version: this.bspecIndex.version,
        lastLoadTime: Date.now()
      });
    } catch (error) {
      const failTime = Date.now() - startTime;
      console.error(`[TIMING] ❌ Failed to load BSpec specification after ${failTime}ms:`, error);
      throw new Error(`Failed to load specification: ${error instanceof Error ? error.message : 'Unknown error'}`);
    }
  }

  /**
   * Set up MCP resources (lazy - only registers resources after spec is loaded)
   */
  private async setupResources() {
    // Ensure spec is loaded first
    await this.ensureSpecificationLoaded();
    
    if (!this.bspecIndex) {
      console.error('Cannot setup resources: BSpec index not loaded');
      return;
    }

    console.log('Setting up MCP resources...');

    // BSpec Overview Resource
    this.server.resource(
      `bspec-overview`,
      `https://mcp.bspec.dev/spec/v${this.bspecIndex.version}/overview`,
      async (uri) => {
        await this.ensureSpecificationLoaded();
        const index = this.bspecIndex!;
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
        overview += `Use the available tools to explore specific document types, domains, or search for specific capabilities.`;

        return {
          contents: [{
            uri: uri.href,
            text: overview,
            mimeType: 'text/markdown'
          }]
        };
      }
    );

    // Register domain and document type resources lazily
    // Only register a few at init, rest on-demand via tools
    console.log(`Registered overview resource. Document types and domains accessible via tools.`);
  }

  /**
   * Set up MCP tools
   */
  private setupTools() {
    // Get Document Type tool
    this.server.tool(
      "get_document_type",
      "Get detailed information about a specific BSpec document type",
      {
        type_code: z.string().describe("Three-letter document type code (e.g., 'MSN', 'VSN', 'STR')")
      },
      async ({ type_code }) => {
        const startTime = Date.now();
        console.log(`[TIMING] Tool: get_document_type(${type_code})`);
        await this.ensureSpecificationLoaded();
        console.log(`[TIMING] Spec loaded in ${Date.now() - startTime}ms`);
        
        const typeCode = type_code.toUpperCase();
        const docType = this.bspecIndex!.documentTypes.get(typeCode);
        
        if (!docType) {
          console.log(`[TIMING] get_document_type completed in ${Date.now() - startTime}ms (not found)`);
          return {
            content: [{
              type: "text",
              text: `Error: Document type "${typeCode}" not found. Use list_document_types to see all available types.`
            }]
          };
        }

        console.log(`[TIMING] get_document_type completed in ${Date.now() - startTime}ms`);
        return {
          content: [{
            type: "text",
            text: `# ${docType.code}: ${docType.name}\n\n**Domain:** ${docType.domain}\n**Description:** ${docType.description}\n\n**Full Specification:**\n\n${docType.content}`
          }]
        };
      }
    );

    // List Document Types tool
    this.server.tool(
      "list_document_types",
      "List all BSpec document types with filtering options",
      {
        domain: z.string().optional().describe("Filter by domain (e.g., 'strategic-foundation', 'market-environment')"),
        include_content: z.boolean().optional().default(false).describe("Include full specification content")
      },
      async ({ domain, include_content }) => {
        const startTime = Date.now();
        console.log(`[TIMING] Tool: list_document_types(domain=${domain}, include_content=${include_content})`);
        await this.ensureSpecificationLoaded();
        
        let filteredTypes = Array.from(this.bspecIndex!.documentTypes.values());
        if (domain) {
          filteredTypes = filteredTypes.filter(type =>
            type.domain.toLowerCase().includes(domain.toLowerCase()) ||
            type.code.toLowerCase().includes(domain.toLowerCase())
          );
        }

        let response = `# BSpec Document Types (${filteredTypes.length} types)\n\n`;

        for (const type of filteredTypes.sort((a, b) => a.code.localeCompare(b.code))) {
          response += `## ${type.code}: ${type.name}\n`;
          response += `**Domain:** ${type.domain}\n`;
          response += `**Description:** ${type.description}\n`;

          if (include_content) {
            response += `\n**Full Specification:**\n${type.content}\n`;
          }
          response += '\n---\n\n';
        }

        console.log(`[TIMING] list_document_types completed in ${Date.now() - startTime}ms (${filteredTypes.length} types)`);
        return {
          content: [{
            type: "text",
            text: response
          }]
        };
      }
    );

    // Get Domain Info tool
    this.server.tool(
      "get_domain_info",
      "Get information about a specific BSpec business domain",
      {
        domain: z.string().describe("Domain name (e.g., 'strategic-foundation', 'customer-value')")
      },
      async ({ domain }) => {
        await this.ensureSpecificationLoaded();
        
        const domainInfo = this.bspecIndex!.domains.get(domain);
        if (!domainInfo) {
          const availableDomains = Array.from(this.bspecIndex!.domains.keys()).join(', ');
          return {
            content: [{
              type: "text",
              text: `Error: Domain "${domain}" not found. Available domains: ${availableDomains}`
            }]
          };
        }

        let domainResponse = `# ${domainInfo.name}\n\n`;
        domainResponse += `**Description:** ${domainInfo.description}\n`;
        domainResponse += `**Total Document Types:** ${domainInfo.totalTypes}\n\n`;
        domainResponse += `## Document Types in this Domain:\n\n`;

        for (const typeCode of domainInfo.documentTypes) {
          const type = this.bspecIndex!.documentTypes.get(typeCode);
          if (type) {
            domainResponse += `- **${typeCode}**: ${type.name} - ${type.description}\n`;
          }
        }

        return {
          content: [{
            type: "text",
            text: domainResponse
          }]
        };
      }
    );

    // List Domains tool
    this.server.tool(
      "list_domains",
      "List all BSpec business domains with their document types",
      {
        include_types: z.boolean().optional().default(true).describe("Include document types in each domain")
      },
      async ({ include_types }) => {
        await this.ensureSpecificationLoaded();
        
        let domainsResponse = `# BSpec Business Domains (${this.bspecIndex!.domains.size} domains)\n\n`;

        for (const [domainKey, domainInfo] of this.bspecIndex!.domains.entries()) {
          domainsResponse += `## ${domainInfo.name}\n`;
          domainsResponse += `**Key:** ${domainKey}\n`;
          domainsResponse += `**Description:** ${domainInfo.description}\n`;
          domainsResponse += `**Document Types:** ${domainInfo.totalTypes}\n`;

          if (include_types) {
            domainsResponse += `\n**Types:** ${domainInfo.documentTypes.join(', ')}\n`;
          }
          domainsResponse += '\n---\n\n';
        }

        return {
          content: [{
            type: "text",
            text: domainsResponse
          }]
        };
      }
    );

    // Get BSpec Overview tool
    this.server.tool(
      "get_bspec_overview",
      "Get a comprehensive overview of the BSpec specification",
      {},
      async () => {
        await this.ensureSpecificationLoaded();
        
        const index = this.bspecIndex!;
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

        return {
          content: [{
            type: "text",
            text: overview
          }]
        };
      }
    );

    // Search Document Types tool
    this.server.tool(
      "search_document_types",
      "Search for document types by name, description, or purpose",
      {
        query: z.string().describe("Search query to match against document type names, descriptions, or purposes")
      },
      async ({ query }) => {
        await this.ensureSpecificationLoaded();
        
        const searchQuery = query.toLowerCase();
        const matchingTypes = Array.from(this.bspecIndex!.documentTypes.values()).filter(type =>
          type.name.toLowerCase().includes(searchQuery) ||
          type.description.toLowerCase().includes(searchQuery) ||
          type.purpose.toLowerCase().includes(searchQuery) ||
          type.code.toLowerCase().includes(searchQuery)
        );

        if (matchingTypes.length === 0) {
          return {
            content: [{
              type: "text",
              text: `No document types found matching "${query}". Try broader search terms or use list_document_types to see all types.`
            }]
          };
        }

        let searchResponse = `# Search Results for "${query}" (${matchingTypes.length} matches)\n\n`;

        for (const type of matchingTypes.sort((a, b) => a.code.localeCompare(b.code))) {
          searchResponse += `## ${type.code}: ${type.name}\n`;
          searchResponse += `**Domain:** ${type.domain}\n`;
          searchResponse += `**Description:** ${type.description}\n\n`;
        }

        return {
          content: [{
            type: "text",
            text: searchResponse
          }]
        };
      }
    );

    // Get Document Relationships tool
    this.server.tool(
      "get_document_relationships",
      "Get information about how document types typically relate to each other",
      {
        type_code: z.string().describe("Document type code to get relationships for")
      },
      async ({ type_code }) => {
        await this.ensureSpecificationLoaded();
        
        const typeCode = type_code.toUpperCase();
        const docType = this.bspecIndex!.documentTypes.get(typeCode);
        
        if (!docType) {
          return {
            content: [{
              type: "text",
              text: `Error: Document type "${typeCode}" not found.`
            }]
          };
        }

        // Extract relationship information from the content
        const relationshipInfo = this.extractRelationshipInfo(docType.content);
        
        let relResponse = `# ${typeCode}: ${docType.name} - Relationships\n\n`;
        relResponse += `**Domain:** ${docType.domain}\n\n`;
        relResponse += relationshipInfo;

        return {
          content: [{
            type: "text",
            text: relResponse
          }]
        };
      }
    );
  }

  /**
   * Extract relationship information from document content
   */
  private extractRelationshipInfo(content: string): string {
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
}