#!/usr/bin/env node

// Simple Node.js test script for MCP server
import { createBSpecMCPServer } from './dist/server.js';

async function testMCPServer() {
  console.log('Testing BSpec MCP Server...');

  try {
    // Create server instance
    const server = createBSpecMCPServer();
    console.log('Server created successfully');

    // Check if server has expected properties
    console.log('Server properties:');
    console.log('- name:', server.name);
    console.log('- version:', server.version);
    console.log('- available methods:', Object.getOwnPropertyNames(server));

    // Check what methods are available on the server prototype
    console.log('- prototype methods:', Object.getOwnPropertyNames(Object.getPrototypeOf(server)));

    // Check registered tools
    console.log('\nRegistered tools:');
    console.log('- count:', server._registeredTools?.size || 0);
    if (server._registeredTools) {
      console.log('- tool names:', Array.from(server._registeredTools.keys()));
    }

  } catch (error) {
    console.error('Test failed:', error);
  }
}

testMCPServer().catch(console.error);