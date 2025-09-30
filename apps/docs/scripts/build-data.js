import { readFileSync, writeFileSync, mkdirSync, readdirSync, statSync } from 'fs';
import { resolve, dirname, join } from 'path';
import matter from 'gray-matter';
import { marked } from 'marked';
import { markedHighlight } from 'marked-highlight';
import hljs from 'highlight.js';

// Domain name mapping: folder name -> friendly display name
const DOMAIN_NAMES = {
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
	'brand-marketing': 'Brand & Marketing'
};

// Domain descriptions for richer context
const DOMAIN_DESCRIPTIONS = {
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
	'brand-marketing': 'Brand positioning, marketing strategy, and communications'
};

async function loadBSpecFromDirectory(specPath) {
	const files = [];

	// Process markdown files directly from spec directory
	function processDirectory(dirPath, relativePath = '') {
		const items = readdirSync(dirPath);

		for (const item of items) {
			const fullPath = join(dirPath, item);
			const stat = statSync(fullPath);
			const itemRelativePath = relativePath ? join(relativePath, item) : item;

			if (stat.isDirectory()) {
				processDirectory(fullPath, itemRelativePath);
			} else if (item.endsWith('.md')) {
				try {
					const content = readFileSync(fullPath, 'utf8');
					files.push({
						path: itemRelativePath,
						content
					});
				} catch (error) {
					console.warn(`Warning: Could not read ${fullPath}:`, error.message);
				}
			}
		}
	}

	processDirectory(specPath);

	// Parse the specification structure
	const specification = {
		version: '1.0.0',
		mainFiles: {},
		domains: [],
		allFiles: files
	};

	// Process files to extract structured data
	for (const file of files) {
		try {
			// Handle main files
			if (file.path === 'spec.md') {
				const { data: frontmatter, content: markdownContent } = matter(file.content);
				specification.mainFiles.spec = {
					path: file.path,
					content: markdownContent,
					metadata: frontmatter
				};
			} else if (file.path === 'index.md') {
				const { data: frontmatter, content: markdownContent } = matter(file.content);
				specification.mainFiles.index = {
					path: file.path,
					content: markdownContent,
					metadata: frontmatter
				};
			} else if (file.path === 'README.md') {
				const { data: frontmatter, content: markdownContent } = matter(file.content);
				specification.mainFiles.readme = {
					path: file.path,
					content: markdownContent,
					metadata: frontmatter
				};
			}
			// Handle domain files
			else if (file.path.includes('/') && file.path.endsWith('-spec.md')) {
				const { data: frontmatter, content: markdownContent } = matter(file.content);
				const pathParts = file.path.split('/');

				if (pathParts.length === 2) { // domain/TYPE-spec.md
					const domainSlug = pathParts[0];
					const fileName = pathParts[1];
					const documentType = fileName.replace('-spec.md', '');

					// Get friendly domain name from mapping
					const domainName = DOMAIN_NAMES[domainSlug] || domainSlug;
					const domainDescription = DOMAIN_DESCRIPTIONS[domainSlug] || '';

					// Find or create domain
					let domain = specification.domains.find(d => d.slug === domainSlug);
					if (!domain) {
						domain = {
							name: domainName,
							slug: domainSlug,
							description: domainDescription,
							documents: []
						};
						specification.domains.push(domain);
					}

					// Extract document name from markdown content
					// Look for "**Document Type Name:**" in the content
					let documentName = documentType.toUpperCase(); // Default fallback
					const docTypeNameMatch = markdownContent.match(/\*\*Document Type Name:\*\*\s*(.+)/);
					if (docTypeNameMatch && docTypeNameMatch[1]) {
						documentName = docTypeNameMatch[1].trim();
					} else if (frontmatter['Document Type Name']) {
						documentName = frontmatter['Document Type Name'];
					} else if (frontmatter.title) {
						documentName = frontmatter.title;
					}

					// Add document to domain
					domain.documents.push({
						path: file.path,
						type: documentType,
						name: documentName,
						slug: documentType.toLowerCase(),
						content: markdownContent,
						frontmatter,
						domainName,
						domainSlug,
						metadata: frontmatter
					});
				}
			}
		} catch (error) {
			console.warn(`Warning: Could not parse ${file.path}:`, error.message);
		}
	}

	return specification;
}


// Configure marked with syntax highlighting
marked.use(markedHighlight({
	langPrefix: 'hljs language-',
	highlight(code, lang) {
		const language = hljs.getLanguage(lang) ? lang : 'plaintext';
		return hljs.highlight(code, { language }).value;
	}
}));

// Configure marked options
marked.setOptions({
	breaks: false,
	gfm: true,
	headerIds: false,
	silent: true
});

async function processMarkdownContent(content) {
	if (!content) return '';

	try {
		// Process links to make them work in the docs site
		let processedContent = content;

		// Process links in order from most specific to most general
		// 1. First handle specific main files before domain patterns
		processedContent = processedContent.replace(
			/\[([^\]]+)\]\(\.\/spec\.md\)/g,
			'[$1](/spec)'
		);

		processedContent = processedContent.replace(
			/\[([^\]]+)\]\(\.\/README\.md\)/g,
			'[$1](/)'
		);

		processedContent = processedContent.replace(
			/\[([^\]]+)\]\(\.\/index\.md\)/g,
			'[$1](/index)'
		);

		// 2. Convert ./domain/TYPE-spec.md links to /domains/domain/type format
		processedContent = processedContent.replace(
			/\[([^\]]+)\]\(\.\/([^/]+)\/([^)]+)-spec\.md\)/g,
			'[$1](/domains/$2/$3)'
		);

		return marked.parse(processedContent);
	} catch (error) {
		console.warn('Error processing markdown:', error.message);
		return content; // Return raw content if parsing fails
	}
}

async function buildStaticData() {
	try {
		console.log('Building static data from spec/v1/ directory...');

		// Load directly from the spec directory
		const specPath = resolve('../../spec/v1');
		console.log('Loading from spec path:', specPath);

		const specification = await loadBSpecFromDirectory(specPath);

		// Ensure static data directory exists
		const dataDir = resolve('./src/lib/data');
		console.log('Current working directory:', process.cwd());
		console.log('Data directory:', dataDir);
		mkdirSync(dataDir, { recursive: true });
		console.log('Created directory, now writing files...');

		// Test write
		try {
			writeFileSync(resolve(dataDir, 'test-write.txt'), 'Test content');
			console.log('Test file written successfully');
		} catch (error) {
			console.error('Test file write failed:', error);
		}

		// Generate lightweight data for homepage - only include metadata, not full specification
		const homeData = {
			domains: specification.domains.map(domain => ({
				name: domain.name,
				slug: domain.slug,
				description: domain.description,
				path: domain.path,
				documentCount: domain.documents.length,
				documentTypes: domain.documents.map(doc => ({ type: doc.type, name: doc.name, slug: doc.slug }))
			})),
			version: specification.version,
			totalFiles: specification.allFiles.length,
			content: specification.mainFiles.index
				? await processMarkdownContent(specification.mainFiles.index.content)
				: specification.mainFiles.spec
					? await processMarkdownContent(specification.mainFiles.spec.content)
					: '',
			title: 'BSpec Documentation',
			description: 'Universal Business Specification Standard - A structured, machine-readable knowledge graph for describing any business'
		};

		// Generate lightweight data for spec page
		const specData = {
			content: specification.mainFiles.spec ? await processMarkdownContent(specification.mainFiles.spec.content) : '',
			title: 'BSpec Specification',
			description: 'Complete Business Specification Standard v1.0',
			version: specification.version
		};

		// Write home.json
		console.log('Writing home.json...');
		console.log('Home file path:', resolve(dataDir, 'home.json'));
		console.log('Home data size:', JSON.stringify(homeData).length);
		writeFileSync(resolve(dataDir, 'home.json'), JSON.stringify(homeData, null, 2));
		console.log('Home.json written successfully');
		console.log('File verification - exists and size:', JSON.stringify(homeData).length);

		console.log('Writing spec.json...');
		writeFileSync(resolve(dataDir, 'spec.json'), JSON.stringify(specData, null, 2));
		console.log('Spec.json written');

		// Generate lightweight data for domains page
		const domainsData = {
			domains: specification.domains.map(domain => ({
				name: domain.name,
				slug: domain.slug,
				description: domain.description,
				path: domain.path,
				documentCount: domain.documents.length,
				documents: domain.documents.map(doc => ({
					type: doc.type,
					name: doc.name,
					slug: doc.slug,
					path: doc.path,
					// Only include first 500 chars of content for preview
					preview: doc.content.substring(0, 500) + (doc.content.length > 500 ? '...' : '')
				}))
			})),
			title: 'Business Domains',
			description: 'Explore the 11 core business domains in the BSpec specification'
		};

		console.log('Writing domains.json...');
		writeFileSync(resolve(dataDir, 'domains.json'), JSON.stringify(domainsData, null, 2));
		console.log('Domains.json written');

		// Generate lightweight data for format page
		const formatData = {
			content: specification.mainFiles.format ? await processMarkdownContent(specification.mainFiles.format.content) : '',
			title: 'BSpec Format Guide',
			description: 'Learn the BSpec document format and structure'
		};

		console.log('Writing format.json...');
		writeFileSync(resolve(dataDir, 'format.json'), JSON.stringify(formatData, null, 2));
		console.log('Format.json written');

		// Generate lightweight data for index page (separate from homepage)
		const indexData = {
			domains: specification.domains.map(domain => ({
				name: domain.name,
				slug: domain.slug,
				description: domain.description,
				documentTypes: domain.documents.map(doc => ({ type: doc.type, name: doc.name, slug: doc.slug }))
			})),
			content: specification.mainFiles.index ? await processMarkdownContent(specification.mainFiles.index.content) : '',
			title: 'BSpec 1.0 Documentation Index',
			description: 'Complete navigation guide to all BSpec document types and domains'
		};

		console.log('Writing index.json...');
		writeFileSync(resolve(dataDir, 'index.json'), JSON.stringify(indexData, null, 2));
		console.log('Index.json written');

		console.log('Static data generation complete!');
		console.log(`Generated data for ${specification.domains.length} domains`);
		console.log(`Processed ${specification.allFiles.length} total files`);

		// Debug: Check if main files exist
		console.log('Main files:', Object.keys(specification.mainFiles));
		console.log('Spec file exists:', !!specification.mainFiles.spec);
		console.log('Version:', specification.version);

	} catch (error) {
		console.error('Error building static data:', error);
		process.exit(1);
	}
}

buildStaticData();