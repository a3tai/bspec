/**
 * Ultra-Simple BSpec Specification Server
 *
 * Cloudflare Worker that loads TGZ from R2 and serves specification files via MCP
 * Implements Streamable HTTP transport with full MCP 2025-03-26 specification compliance
 */

import { initializeServer } from './server.js';

// Cloudflare Workers types
interface Env {
  ASSETS: R2Bucket;
}

// Global server instance and connection state
let serverInstance: any = null;
let isInitialized = false;
const SUPPORTED_PROTOCOL_VERSIONS = ['2024-11-05', '2025-03-26', '2025-06-18'];
const DEFAULT_PROTOCOL_VERSION = '2024-11-05';

// Session management for SSE connections
const activeSessions = new Map<string, any>();

// Event counter for resumability
let eventCounter = 0;

/**
 * Initialize server if not already done
 */
async function ensureServerLoaded(env: Env) {
  if (!serverInstance) {
    console.log('Initializing BSpec specification server...');
    serverInstance = await initializeServer(env);
    console.log(`Server ready: ${serverInstance.name} v${serverInstance.version}`);
  }
  return serverInstance;
}

/**
 * Generate unique session ID for SSE connections
 */
function generateSessionId(): string {
  return crypto.randomUUID();
}

/**
 * Generate unique event ID for SSE resumability
 */
function generateEventId(): string {
  return `event-${++eventCounter}-${Date.now()}`;
}

/**
 * Validate Accept header contains required values
 */
function validateAcceptHeader(acceptHeader: string | null, requiredTypes: string[]): boolean {
  if (!acceptHeader) return false;

  const acceptedTypes = acceptHeader.toLowerCase().split(',').map(t => t.trim());
  return requiredTypes.every(required =>
    acceptedTypes.some(accepted => accepted.includes(required.toLowerCase()))
  );
}

/**
 * Validate Origin header for security
 */
function validateOrigin(origin: string | null): boolean {
  // For public deployment, we allow any origin
  // In production, you might want to restrict this
  return true;
}

/**
 * Process MCP request and return appropriate response
 */
async function processMcpRequest(body: any, server: any, sessionId?: string): Promise<any> {
  // MCP protocol handling
  if (body.method === 'initialize') {
    // Validate protocol version
    const clientVersion = body.params?.protocolVersion;
    let protocolVersion = DEFAULT_PROTOCOL_VERSION;

    if (clientVersion && SUPPORTED_PROTOCOL_VERSIONS.includes(clientVersion)) {
      protocolVersion = clientVersion;
    } else if (clientVersion) {
      // Client requested unsupported version
      return {
        jsonrpc: '2.0',
        id: body.id,
        error: {
          code: -32602,
          message: 'Unsupported protocol version',
          data: {
            supported: SUPPORTED_PROTOCOL_VERSIONS,
            requested: clientVersion
          }
        }
      };
    }

    return {
      jsonrpc: '2.0',
      id: body.id,
      result: {
        protocolVersion,
        capabilities: {
          tools: {
            listChanged: false
          },
          resources: {
            listChanged: false,
            subscribe: false
          }
        },
        serverInfo: {
          name: server.name,
          version: server.version
        },
        instructions: 'BSpec MCP Server provides access to the Business Specification Standard (BSpec) v1.0 through resources and tools. Use resources to explore the specification structure, and tools for intelligent queries.'
      }
    };
  }

  if (body.method === 'tools/list') {
    try {
      const tools = server.listTools();
      const progressToken = body.params?._meta?.progressToken;
      const result: any = { tools };
      if (progressToken) {
        result._meta = { progressToken };
      }

      return {
        jsonrpc: '2.0',
        id: body.id,
        result
      };
    } catch (error) {
      return {
        jsonrpc: '2.0',
        id: body.id,
        error: {
          code: -32603,
          message: error instanceof Error ? error.message : 'Internal error listing tools'
        }
      };
    }
  }

  if (body.method === 'tools/call') {
    try {
      const { name, arguments: args } = body.params;

      // Validate tool name
      if (!name || typeof name !== 'string') {
        return {
          jsonrpc: '2.0',
          id: body.id,
          error: {
            code: -32602,
            message: 'Invalid params: tool name is required and must be a string'
          }
        };
      }

      const result = await server.callTool(name, args);
      const progressToken = body.params?._meta?.progressToken;

      if (progressToken && result._meta) {
        result._meta.progressToken = progressToken;
      } else if (progressToken) {
        result._meta = { progressToken };
      }

      return {
        jsonrpc: '2.0',
        id: body.id,
        result
      };
    } catch (error) {
      const errorMessage = error instanceof Error ? error.message : 'Internal error';
      let errorCode = -32603; // Internal error

      // More specific error codes
      if (errorMessage.includes('not found') || errorMessage.includes('Unknown tool')) {
        errorCode = -32601; // Method not found
      } else if (errorMessage.includes('required') || errorMessage.includes('Invalid')) {
        errorCode = -32602; // Invalid params
      }

      return {
        jsonrpc: '2.0',
        id: body.id,
        error: {
          code: errorCode,
          message: errorMessage
        }
      };
    }
  }

  if (body.method === 'resources/list') {
    try {
      const resources = server.listResources();
      const progressToken = body.params?._meta?.progressToken;
      const result: any = { resources };
      if (progressToken) {
        result._meta = { progressToken };
      }

      return {
        jsonrpc: '2.0',
        id: body.id,
        result
      };
    } catch (error) {
      return {
        jsonrpc: '2.0',
        id: body.id,
        error: {
          code: -32603,
          message: error instanceof Error ? error.message : 'Internal error listing resources'
        }
      };
    }
  }

  if (body.method === 'resources/read') {
    try {
      const { uri } = body.params;

      // Validate URI parameter
      if (!uri || typeof uri !== 'string') {
        return {
          jsonrpc: '2.0',
          id: body.id,
          error: {
            code: -32602,
            message: 'Invalid params: uri is required and must be a string'
          }
        };
      }

      // Basic URI validation (must be https and mcp.bspec.dev)
      try {
        const parsedUri = new URL(uri);
        if (parsedUri.protocol !== 'https:' || parsedUri.hostname !== 'mcp.bspec.dev') {
          return {
            jsonrpc: '2.0',
            id: body.id,
            error: {
              code: -32602,
              message: 'Invalid URI: Only https://mcp.bspec.dev/* URIs are supported'
            }
          };
        }
      } catch (urlError) {
        return {
          jsonrpc: '2.0',
          id: body.id,
          error: {
            code: -32602,
            message: 'Invalid URI format'
          }
        };
      }

      const result = await server.readResource(uri);
      const progressToken = body.params?._meta?.progressToken;

      if (progressToken && result._meta) {
        result._meta.progressToken = progressToken;
      } else if (progressToken) {
        result._meta = { progressToken };
      }

      return {
        jsonrpc: '2.0',
        id: body.id,
        result
      };
    } catch (error) {
      const errorMessage = error instanceof Error ? error.message : 'Internal error';
      let errorCode = -32603; // Internal error

      // More specific error codes for resource errors
      if (errorMessage.includes('not found') || errorMessage.includes('Resource not found')) {
        errorCode = -32602; // Invalid params (resource doesn't exist)
      }

      return {
        jsonrpc: '2.0',
        id: body.id,
        error: {
          code: errorCode,
          message: errorMessage
        }
      };
    }
  }

  // Handle initialized notification
  if (body.method === 'notifications/initialized') {
    isInitialized = true;
    console.log('Client completed initialization');
    // Notifications don't return responses - return null to indicate no response
    return null;
  }

  // Check if we're properly initialized for non-ping requests
  if (!isInitialized && body.method !== 'ping' && body.method !== 'initialize') {
    return {
      jsonrpc: '2.0',
      id: body.id,
      error: {
        code: -32002,
        message: 'Server not initialized. Send initialize request followed by initialized notification first.'
      }
    };
  }

  if (body.method === 'ping') {
    // Extract progress token if present
    const progressToken = body.params?._meta?.progressToken;
    const result: any = {};
    if (progressToken) {
      result._meta = { progressToken };
    }

    return {
      jsonrpc: '2.0',
      id: body.id,
      result
    };
  }

  return {
    jsonrpc: '2.0',
    id: body.id,
    error: {
      code: -32601,
      message: 'Method not found'
    }
  };
}

/**
 * Main Cloudflare Workers fetch handler - Implements Streamable HTTP Transport
 */
export default {
  async fetch(request: Request, env: Env): Promise<Response> {
    // Handle CORS preflight
    if (request.method === 'OPTIONS') {
      return new Response(null, {
        status: 204,
        headers: {
          'Access-Control-Allow-Origin': '*',
          'Access-Control-Allow-Methods': 'GET, POST, OPTIONS',
          'Access-Control-Allow-Headers': 'Content-Type, Mcp-Session-Id',
        }
      });
    }

    const url = new URL(request.url);

    try {
      // Initialize server
      const server = await ensureServerLoaded(env);

      // Streamable HTTP Transport Implementation
      // Based on MCP Specification 2025-03-26

      if (url.pathname === '/health') {
        return new Response(JSON.stringify({
          status: 'healthy',
          name: server.name,
          version: server.version,
          timestamp: new Date().toISOString()
        }), {
          headers: {
            'Content-Type': 'application/json',
            'Access-Control-Allow-Origin': '*'
          }
        });
      }

      // Serve documentation from R2 under /docs path
      if (request.method === 'GET' && url.pathname.startsWith('/docs')) {
        try {
          // Handle docs routing
          let docPath = url.pathname;

          // Default to index.html for directory paths
          if (docPath === '/docs' || docPath === '/docs/') {
            docPath = '/docs/index.html';
          } else if (docPath.endsWith('/')) {
            docPath = `${docPath}index.html`;
          } else if (!docPath.includes('.')) {
            // Add .html extension if no extension provided
            docPath = `${docPath}.html`;
          }

          // Try to get documentation file from R2
          const docsObject = await env.ASSETS.get(`docs${docPath.replace('/docs', '')}`);
          if (docsObject) {
            // Determine content type
            const ext = docPath.split('.').pop()?.toLowerCase();
            const contentType = ext === 'html' ? 'text/html' :
                              ext === 'css' ? 'text/css' :
                              ext === 'js' ? 'application/javascript' :
                              ext === 'png' ? 'image/png' :
                              ext === 'jpg' || ext === 'jpeg' ? 'image/jpeg' :
                              ext === 'svg' ? 'image/svg+xml' :
                              ext === 'ico' ? 'image/x-icon' :
                              ext === 'xml' ? 'application/xml' :
                              ext === 'txt' ? 'text/plain' :
                              ext === 'json' ? 'application/json' :
                              'application/octet-stream';

            return new Response(docsObject.body, {
              headers: {
                'Content-Type': contentType,
                'Cache-Control': ext === 'html' ? 'no-cache, must-revalidate' : 'public, max-age=86400',
                'Access-Control-Allow-Origin': '*',
                'X-Docs-Version': server.version
              }
            });
          }
        } catch (error) {
          console.error('Error serving documentation:', error);
        }
      }

      // Serve web files from R2 - try to serve any file that exists in web folder
      if (request.method === 'GET' && url.pathname !== '/mcp' && !url.pathname.startsWith('/docs')) {
        try {
          // Default to index.html for root path
          const filePath = url.pathname === '/' ? '/index.html' : url.pathname;
          const object = await env.ASSETS.get(`web${filePath}`);

          if (object) {
            // Determine content type based on file extension
            const ext = filePath.split('.').pop()?.toLowerCase();
            const contentType = ext === 'html' ? 'text/html' :
                              ext === 'css' ? 'text/css' :
                              ext === 'js' ? 'application/javascript' :
                              ext === 'png' ? 'image/png' :
                              ext === 'jpg' || ext === 'jpeg' ? 'image/jpeg' :
                              ext === 'svg' ? 'image/svg+xml' :
                              ext === 'ico' ? 'image/x-icon' :
                              ext === 'webmanifest' ? 'application/manifest+json' :
                              'application/octet-stream';

            return new Response(object.body, {
              headers: {
                'Content-Type': contentType,
                'Cache-Control': ext === 'html' ? 'no-cache' : 'public, max-age=31536000',
                'Access-Control-Allow-Origin': '*'
              }
            });
          }
        } catch (error) {
          console.error('Error serving web file:', error);
        }
      }

      // Streamable HTTP Transport - MCP endpoint only
      if (url.pathname === '/mcp') {

        // Validate Origin header for security
        const origin = request.headers.get('Origin');
        if (!validateOrigin(origin)) {
          return new Response('Invalid Origin', {
            status: 403,
            headers: { 'Access-Control-Allow-Origin': '*' }
          });
        }

        // GET request - Start SSE stream (optional part of Streamable HTTP)
        if (request.method === 'GET') {
          // Validate Accept header includes text/event-stream
          const acceptHeader = request.headers.get('Accept');
          if (!validateAcceptHeader(acceptHeader, ['text/event-stream'])) {
            return new Response('Method not allowed', {
              status: 405,
              headers: {
                'Allow': 'POST',
                'Access-Control-Allow-Origin': '*'
              }
            });
          }
          const sessionId = generateSessionId();
          console.log(`Starting SSE session: ${sessionId}`);

          // Support resumability via Last-Event-ID
          const lastEventId = request.headers.get('Last-Event-ID');
          let eventId = lastEventId ? parseInt(lastEventId.split('-')[1]) || eventCounter : eventCounter;
          eventCounter = Math.max(eventCounter, eventId);

          // Create SSE response stream
          const stream = new ReadableStream({
            start(controller) {
              // Store session
              activeSessions.set(sessionId, { controller, active: true });

              // Send session info as first event with unique ID
              const sessionEventId = generateEventId();
              const sessionEvent = `id: ${sessionEventId}\nevent: session\ndata: ${JSON.stringify({
                sessionId,
                transport: 'sse',
                timestamp: new Date().toISOString()
              })}\n\n`;
              controller.enqueue(new TextEncoder().encode(sessionEvent));

              // Keep connection alive with periodic pings
              const pingInterval = setInterval(() => {
                const session = activeSessions.get(sessionId);
                if (!session?.active) {
                  clearInterval(pingInterval);
                  return;
                }

                try {
                  const pingEventId = generateEventId();
                  const pingEvent = `id: ${pingEventId}\nevent: ping\ndata: ${JSON.stringify({
                    timestamp: new Date().toISOString()
                  })}\n\n`;
                  controller.enqueue(new TextEncoder().encode(pingEvent));
                } catch (error) {
                  console.log(`SSE ping failed for session ${sessionId}:`, error);
                  clearInterval(pingInterval);
                  activeSessions.delete(sessionId);
                  controller.close();
                }
              }, 30000);

              // Cleanup function
              (controller as any).cleanup = () => {
                clearInterval(pingInterval);
                activeSessions.delete(sessionId);
              };
            },
            cancel() {
              // Cleanup on connection close
              console.log(`SSE session closed: ${sessionId}`);
              if ((this as any).cleanup) {
                (this as any).cleanup();
              }
            }
          });

          return new Response(stream, {
            headers: {
              'Content-Type': 'text/event-stream',
              'Cache-Control': 'no-cache',
              'Connection': 'keep-alive',
              'Access-Control-Allow-Origin': '*',
              'Access-Control-Allow-Headers': 'Content-Type, Mcp-Session-Id',
              'Mcp-Session-Id': sessionId
            }
          });
        }

        // POST request - Handle MCP messages (core part of Streamable HTTP)
        if (request.method === 'POST') {
          // Validate Accept header includes application/json
          const acceptHeader = request.headers.get('Accept');
          if (!validateAcceptHeader(acceptHeader, ['application/json'])) {
            return new Response('Invalid Accept header. Expected application/json', {
              status: 400,
              headers: {
                'Access-Control-Allow-Origin': '*',
                'Access-Control-Allow-Headers': 'Content-Type, Mcp-Session-Id'
              }
            });
          }

          const sessionId = request.headers.get('Mcp-Session-Id');
          const body = await request.json() as any;

          // Process the MCP request
          const result = await processMcpRequest(body, server, sessionId || undefined);

          // Handle notifications (no response needed)
          if (result === null) {
            return new Response(null, {
              status: 204,
              headers: {
                'Access-Control-Allow-Origin': '*',
                'Access-Control-Allow-Headers': 'Content-Type, Mcp-Session-Id'
              }
            });
          }

          // For SSE sessions, send response via SSE stream
          if (sessionId && activeSessions.has(sessionId)) {
            const session = activeSessions.get(sessionId);
            if (session?.active) {
              try {
                const messageEventId = generateEventId();
                const messageEvent = `id: ${messageEventId}\nevent: message\ndata: ${JSON.stringify(result)}\n\n`;
                session.controller.enqueue(new TextEncoder().encode(messageEvent));
              } catch (error) {
                console.log(`Failed to send SSE message for session ${sessionId}:`, error);
                activeSessions.delete(sessionId);
              }
            }

            // Return acknowledgment for SSE
            return new Response(JSON.stringify({
              acknowledged: true,
              sessionId,
              timestamp: new Date().toISOString()
            }), {
              headers: {
                'Content-Type': 'application/json',
                'Access-Control-Allow-Origin': '*',
                'Access-Control-Allow-Headers': 'Content-Type, Mcp-Session-Id'
              }
            });
          }

          // Regular HTTP response (non-SSE)
          return new Response(JSON.stringify(result), {
            headers: {
              'Content-Type': 'application/json',
              'Access-Control-Allow-Origin': '*',
              'Access-Control-Allow-Headers': 'Content-Type, Mcp-Session-Id'
            }
          });
        }

        return new Response('Method not allowed', { status: 405 });
      }

      return new Response('Not Found', { status: 404 });

    } catch (error) {
      console.error('Server error:', error);
      return new Response('Internal Server Error', {
        status: 500,
        headers: { 'Access-Control-Allow-Origin': '*' }
      });
    }
  }
};