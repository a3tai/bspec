/**
 * BSpec Document Parser
 *
 * Parses BSpec documents from Markdown with YAML frontmatter
 * Supports all 82 document types defined in BSpec 1.0
 */

import { parse as parseYAML } from 'yaml';
import type { BaseBSpecDocument, DocumentType } from '../types.js';

export interface ParsedBSpecDocument extends BaseBSpecDocument {
  /** Raw markdown content without frontmatter */
  content: string;
  /** Original file path or source identifier */
  source?: string;
}

export interface ParseResult {
  success: boolean;
  document?: ParsedBSpecDocument;
  errors?: string[];
}

/**
 * Parse a BSpec document from raw markdown content
 * Extracts YAML frontmatter and validates basic structure
 */
export function parseBSpecDocument(content: string, source?: string): ParseResult {
  try {
    const { frontmatter, markdown } = extractFrontmatter(content);

    if (!frontmatter) {
      return {
        success: false,
        errors: ['No YAML frontmatter found in document']
      };
    }

    const parsedFrontmatter = parseYAML(frontmatter);

    // Validate required fields
    const errors = validateRequiredFields(parsedFrontmatter);
    if (errors.length > 0) {
      return {
        success: false,
        errors
      };
    }

    const document: ParsedBSpecDocument = {
      ...parsedFrontmatter,
      content: markdown.trim(),
      source
    };

    return {
      success: true,
      document
    };

  } catch (error) {
    return {
      success: false,
      errors: [`Parse error: ${error instanceof Error ? error.message : 'Unknown error'}`]
    };
  }
}

/**
 * Extract YAML frontmatter from markdown content
 */
function extractFrontmatter(content: string): { frontmatter: string | null, markdown: string } {
  const frontmatterRegex = /^---\s*\n([\s\S]*?)\n---\s*\n([\s\S]*)$/;
  const match = content.match(frontmatterRegex);

  if (!match) {
    return { frontmatter: null, markdown: content };
  }

  return {
    frontmatter: match[1],
    markdown: match[2]
  };
}

/**
 * Validate that required BSpec fields are present
 */
function validateRequiredFields(frontmatter: any): string[] {
  const errors: string[] = [];

  // Required fields for all BSpec documents
  const requiredFields = ['id', 'title', 'type', 'status', 'version', 'owner', 'created', 'updated'];

  for (const field of requiredFields) {
    if (!frontmatter[field]) {
      errors.push(`Missing required field: ${field}`);
    }
  }

  // Validate document type is one of the 82 valid types
  if (frontmatter.type && !isValidDocumentType(frontmatter.type)) {
    errors.push(`Invalid document type: ${frontmatter.type}. Must be one of the 82 BSpec document types.`);
  }

  // Validate status
  const validStatuses = ['Draft', 'Review', 'Accepted', 'Deprecated'];
  if (frontmatter.status && !validStatuses.includes(frontmatter.status)) {
    errors.push(`Invalid status: ${frontmatter.status}. Must be one of: ${validStatuses.join(', ')}`);
  }

  // Validate date formats (basic check)
  if (frontmatter.created && !isValidDate(frontmatter.created)) {
    errors.push(`Invalid created date format: ${frontmatter.created}. Expected YYYY-MM-DD`);
  }

  if (frontmatter.updated && !isValidDate(frontmatter.updated)) {
    errors.push(`Invalid updated date format: ${frontmatter.updated}. Expected YYYY-MM-DD`);
  }

  return errors;
}

/**
 * Check if a string is a valid BSpec document type
 */
function isValidDocumentType(type: string): type is DocumentType {
  const validTypes: DocumentType[] = [
    // Strategic Foundation (8)
    'MSN', 'VSN', 'VAL', 'STR', 'OBJ', 'MOT', 'PUR', 'THY',
    // Market & Environment (10)
    'MKT', 'SEG', 'CMP', 'POS', 'TRN', 'ECO', 'OPP', 'THR', 'REG', 'MAC',
    // Customer & Value (12)
    'PER', 'JTB', 'CJM', 'USE', 'STO', 'PAI', 'GAI', 'EMP', 'FEE', 'INT', 'SUR', 'BEH',
    // Product & Service (10)
    'PRD', 'SVC', 'FEA', 'ROD', 'REQ', 'QUA', 'UXD', 'PER', 'INT', 'SUP',
    // Business Model (12)
    'BMC', 'REV', 'PRC', 'CST', 'CHN', 'REL', 'RES', 'ACT', 'PRT', 'UNT', 'LTV', 'CAC',
    // Operations & Execution (12)
    'PRC', 'WFL', 'ORG', 'ROL', 'TEA', 'SKI', 'POL', 'SLA', 'VND', 'FAC', 'TOO', 'CAP',
    // Technology & Data (8)
    'ARC', 'SYS', 'DAT', 'API', 'INF', 'SEC', 'DEV', 'ANA',
    // Financial & Investment (10)
    'FIN', 'BUD', 'FOR', 'FND', 'INV', 'VAL', 'MET', 'REP', 'AUD', 'TAX',
    // Risk & Governance (8)
    'RSK', 'MIT', 'CMP', 'GVN', 'CTL', 'CRI', 'ETH', 'STA',
    // Growth & Innovation (8)
    'GTM', 'GRW', 'SCL', 'EXP', 'INN', 'RND', 'ACQ', 'EXP',
    // Learning & Decisions (6)
    'DEC', 'LRN', 'RET', 'HYP', 'KNO', 'WIS'
  ];

  return validTypes.includes(type as DocumentType);
}

/**
 * Basic date validation for YYYY-MM-DD format
 */
function isValidDate(dateString: string): boolean {
  const dateRegex = /^\d{4}-\d{2}-\d{2}$/;
  if (!dateRegex.test(dateString)) {
    return false;
  }

  const date = new Date(dateString);
  return date.toISOString().slice(0, 10) === dateString;
}

/**
 * Parse multiple BSpec documents from an array of content strings
 */
export function parseBSpecDocuments(documents: Array<{ content: string; source?: string }>): ParseResult[] {
  return documents.map(({ content, source }) => parseBSpecDocument(content, source));
}

/**
 * Extract document metadata without parsing full content
 * Useful for quick queries and listings
 */
export function extractDocumentMetadata(content: string): {
  id?: string;
  type?: DocumentType;
  title?: string;
  status?: string;
  domain?: string;
} {
  const { frontmatter } = extractFrontmatter(content);

  if (!frontmatter) {
    return {};
  }

  try {
    const parsed = parseYAML(frontmatter);
    return {
      id: parsed.id,
      type: parsed.type,
      title: parsed.title,
      status: parsed.status,
      domain: parsed.domain
    };
  } catch {
    return {};
  }
}