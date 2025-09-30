/**
 * BSpec MCP Server using Cloudflare Agents Framework
 *
 * Provides access to the Business Specification Standard (BSpec) v1.0 through 
 * tools and resources using the official Cloudflare Agents SDK.
 */

import { BSpecMcpAgent } from './bspec-agent';

// Environment interface for Cloudflare Workers
interface Env {
  ASSETS: R2Bucket;
  BSpecMcpAgent: DurableObjectNamespace<BSpecMcpAgent>;
  ENVIRONMENT?: string;
  MCP_SERVER_VERSION?: string;
  CORS_ORIGIN?: string;
}

// Export the BSpec MCP Agent class for Durable Objects
export { BSpecMcpAgent };

// HTTP Streamable transport (recommended)
export default BSpecMcpAgent.serve('/mcp', {
  binding: 'BSpecMcpAgent'
});
