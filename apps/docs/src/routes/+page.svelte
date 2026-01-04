<script lang="ts">
	import type { PageData } from './$types';
	import Hero from '$lib/components/Hero.svelte';
	import GlassCard from '$lib/components/GlassCard.svelte';

	let { data }: { data: PageData } = $props();

	const stats = [
		{ value: data.domains.length.toString(), label: 'Business Domains' },
		{ value: data.totalFiles.toString(), label: 'Document Types' },
		{ value: `v${data.version}`, label: 'Specification Version' }
	];

	const features = [
		{
			title: 'Universal Business Language',
			description: 'Standardized document types covering all aspects of business from strategy to execution',
			icon: 'globe'
		},
		{
			title: 'Machine-Readable Knowledge',
			description: 'Structured YAML frontmatter creates business knowledge graphs for AI analysis',
			icon: 'cpu'
		},
		{
			title: 'Atomic Document Design',
			description: 'Each document covers exactly one concern with explicit relationships to others',
			icon: 'atom'
		}
	];

	const mcpFeatures = [
		{
			title: 'AI Agent Integration',
			description: 'Model Context Protocol server provides structured business data to AI systems',
			icon: 'robot'
		},
		{
			title: 'Real-time Analysis',
			description: 'Instant document validation, conformance checking, and gap analysis',
			icon: 'zap'
		},
		{
			title: 'Business Intelligence',
			description: 'AI-powered insights from your complete business knowledge graph',
			icon: 'lightbulb'
		}
	];

	function getIcon(iconName: string) {
		switch (iconName) {
			case 'globe':
				return 'M21 12a9 9 0 01-9 9m9-9a9 9 0 00-9-9m9 9H3m9 9a9 9 0 01-9-9m9 9c1.657 0 3-4.03 3-9s-1.343-9-3-9m0 18c-1.657 0-3-4.03-3-9s1.343-9 3-9m-9 9a9 9 0 019-9';
			case 'cpu':
				return 'M9 3v2m6-2v2M9 19v2m6-2v2M5 9H3m2 6H3m18-6h-2m2 6h-2M7 19h10a2 2 0 002-2V7a2 2 0 00-2-2H7a2 2 0 00-2 2v10a2 2 0 002 2zM9 9h6v6H9V9z';
			case 'atom':
				return 'M12 12m-3 0a3 3 0 1 0 6 0a3 3 0 1 0 -6 0M12 1v6m0 6v6m11-7h-6m-6 0H1m15.5-6.5l-4.24 4.24m-6.36 0L1.66 5.5m16.68 13l-4.24-4.24m-6.36 0L1.66 18.5';
			case 'robot':
				return 'M10 6V4a2 2 0 114 0v2h1a2 2 0 012 2v8a2 2 0 01-2 2H9a2 2 0 01-2-2V8a2 2 0 012-2h1zM9 12h6v2H9v-2zm2-3h2v1h-2V9z';
			case 'zap':
				return 'M13 2L3 14h9l-1 8 10-12h-9l1-8z';
			case 'lightbulb':
				return 'M9 21c0 .55.45 1 1 1h4c.55 0 1-.45 1-1v-1H9v1zm3-19C8.14 2 5 5.14 5 9c0 2.38 1.19 4.47 3 5.74V17c0 .55.45 1 1 1h6c.55 0 1-.45 1-1v-2.26c1.81-1.27 3-3.36 3-5.74 0-3.86-3.14-7-7-7z';
			default:
				return 'M12 2l3.09 6.26L22 9.27l-5 4.87 1.18 6.88L12 17.77l-6.18 3.25L7 14.14 2 9.27l6.91-1.01L12 2z';
		}
	}
</script>

<svelte:head>
	<title>BSpec - Universal Business Specification Standard</title>
	<meta name="description" content="Transform your business documentation into structured, machine-readable knowledge graphs. BSpec provides AI-ready business specifications with 82 standardized document types across 11 domains." />
	<meta property="og:title" content="BSpec - Universal Business Specification Standard" />
	<meta property="og:description" content="Transform your business documentation into structured, machine-readable knowledge graphs." />
	<meta property="og:type" content="website" />
	<meta property="og:image" content="/bspec-icon.png" />
	<meta name="twitter:card" content="summary_large_image" />
</svelte:head>

<!-- Main Content -->
<main class="flex-1 min-w-0 overflow-y-auto">
	<!-- Hero Section -->
	<section class="relative min-h-screen flex items-center justify-center overflow-hidden">
		<!-- Background Gradient -->
		<div class="absolute inset-0 bg-gradient-to-br from-[rgb(var(--color-primary-50))] via-[rgb(var(--color-bg-primary))] to-[rgb(var(--color-primary-100))] dark:from-[rgb(var(--color-neutral-950))] dark:via-[rgb(var(--color-neutral-900))] dark:to-[rgb(var(--color-neutral-800))]"></div>
		
		<!-- Animated Background Elements -->
		<div class="absolute inset-0 overflow-hidden">
			<div class="absolute top-1/4 -left-4 w-72 h-72 bg-gradient-to-br from-[rgb(var(--color-primary-500)/.1)] to-transparent rounded-full blur-3xl animate-pulse"></div>
			<div class="absolute bottom-1/4 -right-4 w-96 h-96 bg-gradient-to-tl from-[rgb(var(--color-primary-400)/.08)] to-transparent rounded-full blur-3xl animate-pulse" style="animation-delay: 1s;"></div>
		</div>

		<div class="relative z-10 max-w-6xl mx-auto px-6 py-24 text-center">
			<!-- Logo and Badge -->
			<div class="mb-8 animate-fade-in">
				<div class="inline-flex items-center justify-center w-20 h-20 bg-gradient-to-br from-[rgb(var(--color-primary-500))] to-[rgb(var(--color-primary-700))] rounded-2xl mb-6 elevation-3">
					<img src="/bspec-icon.png" alt="BSpec" class="w-12 h-12 object-cover" />
				</div>
				<div class="inline-flex items-center px-4 py-2 bg-[rgb(var(--color-bg-secondary))] rounded-full border border-[rgb(var(--color-border-light))] mb-6 glass-light">
				<span class="text-sm font-medium text-[rgb(var(--color-text-secondary))]">Now with MCP Integration</span>
				<span class="ml-2 w-2 h-2 bg-[rgb(var(--color-success))] rounded-full animate-pulse"></span>
				</div>
			</div>

			<!-- Main Headline -->
			<h1 class="text-5xl md:text-7xl font-bold text-[rgb(var(--color-text-primary))] mb-6 animate-slide-up">
				<span class="bg-gradient-to-r from-[rgb(var(--color-text-primary))] via-[rgb(var(--color-primary-600))] to-[rgb(var(--color-text-primary))] bg-clip-text text-transparent">
					BSpec
				</span>
			</h1>
			
			<h2 class="text-2xl md:text-3xl font-light text-[rgb(var(--color-text-secondary))] mb-6 animate-slide-up" style="animation-delay: 0.1s;">
				Universal Business Specification Standard
			</h2>

			<p class="text-xl text-[rgb(var(--color-text-tertiary))] max-w-3xl mx-auto mb-12 leading-relaxed animate-slide-up" style="animation-delay: 0.2s;">
				Transform your business documentation into structured, machine-readable knowledge graphs. 
				Designed for the AI era with <strong class="text-[rgb(var(--color-text-primary))]">82 standardized document types</strong> 
				across <strong class="text-[rgb(var(--color-text-primary))]">11 comprehensive domains</strong>.
			</p>

			<!-- CTA Buttons -->
			<div class="flex flex-col sm:flex-row gap-4 justify-center items-center mb-16 animate-slide-up" style="animation-delay: 0.3s;">
				<a
					href="/domains"
					class="inline-flex items-center px-8 py-4 bg-gradient-to-r from-[rgb(var(--color-primary-500))] to-[rgb(var(--color-primary-600))] text-white font-semibold rounded-2xl shadow-lg hover:shadow-xl hover:scale-105 transition-all duration-300 elevation-2"
				>
					<span>Explore Domains</span>
					<svg class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
						<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7l5 5m0 0l-5 5m5-5H6" />
					</svg>
				</a>
				<a
					href="#mcp"
					class="inline-flex items-center px-8 py-4 bg-[rgb(var(--color-bg-secondary))] text-[rgb(var(--color-text-primary))] font-semibold rounded-2xl border border-[rgb(var(--color-border-medium))] hover:bg-[rgb(var(--color-bg-hover))] hover:scale-105 transition-all duration-300 glass-light"
				>
					<span>MCP Integration</span>
					<svg class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
						<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.663 17h4.673M12 3v1m6.364 1.636l-.707.707M21 12h-1M4 12H3m3.343-5.657l-.707-.707m2.828 9.9a5 5 0 117.072 0l-.548.547A3.374 3.374 0 0014 18.469V19a2 2 0 11-4 0v-.531c0-.895-.356-1.754-.988-2.386l-.548-.547z" />
					</svg>
				</a>
			</div>

			<!-- Stats -->
			<div class="grid grid-cols-3 gap-8 max-w-2xl mx-auto animate-fade-in" style="animation-delay: 0.4s;">
				{#each stats as stat}
					<div class="text-center">
						<div class="text-3xl md:text-4xl font-bold text-[rgb(var(--color-primary-600))] mb-2">{stat.value}</div>
						<div class="text-sm text-[rgb(var(--color-text-tertiary))]">{stat.label}</div>
					</div>
				{/each}
			</div>
		</div>
	</section>

	<!-- Features Section -->
	<section class="py-24 bg-[rgb(var(--color-bg-secondary)/.3)]">
		<div class="max-w-6xl mx-auto px-6">
			<div class="text-center mb-16 animate-slide-up">
				<h2 class="text-4xl md:text-5xl font-bold text-[rgb(var(--color-text-primary))] mb-6">
					Business Knowledge,
					<span class="bg-gradient-to-r from-[rgb(var(--color-primary-500))] to-[rgb(var(--color-primary-700))] bg-clip-text text-transparent">
						Reinvented
					</span>
				</h2>
				<p class="text-xl text-[rgb(var(--color-text-secondary))] max-w-3xl mx-auto">
					Move beyond scattered documents to structured business knowledge that both humans and AI can understand.
				</p>
			</div>

			<div class="grid md:grid-cols-3 gap-8">
				{#each features as feature, i}
					<GlassCard variant="medium" elevation={2} className="h-full animate-scale-in animate-stagger-{i+1}">
						<div class="text-center p-8">
							<div class="w-16 h-16 bg-gradient-to-br from-[rgb(var(--color-primary-500))] to-[rgb(var(--color-primary-600))] rounded-2xl flex items-center justify-center mb-6 mx-auto elevation-2">
								<svg class="w-8 h-8 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
									<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="{getIcon(feature.icon)}" />
								</svg>
							</div>
							<h3 class="text-xl font-semibold text-[rgb(var(--color-text-primary))] mb-4">{feature.title}</h3>
							<p class="text-[rgb(var(--color-text-secondary))] leading-relaxed">{feature.description}</p>
						</div>
					</GlassCard>
				{/each}
			</div>
		</div>
	</section>

	<!-- MCP Integration Section -->
	<section id="mcp" class="py-24 relative overflow-hidden">
		<!-- Background Pattern -->
		<div class="absolute inset-0 opacity-5">
			<div class="absolute inset-0 bg-[radial-gradient(circle_at_center,_rgba(var(--color-primary-500),0.1)_1px,_transparent_1px)] bg-[size:60px_60px]"></div>
		</div>

		<div class="relative z-10 max-w-6xl mx-auto px-6">
			<div class="text-center mb-16 animate-slide-up">
				<div class="inline-flex items-center px-4 py-2 bg-gradient-to-r from-[rgb(var(--color-primary-500)/.1)] to-[rgb(var(--color-primary-600)/.1)] rounded-full border border-[rgb(var(--color-primary-500)/.2)] mb-6">
					<span class="text-sm font-medium text-[rgb(var(--color-primary-600))] dark:text-[rgb(var(--color-primary-400))]">
						✨ AI-Native Integration
					</span>
				</div>
				<h2 class="text-4xl md:text-5xl font-bold text-[rgb(var(--color-text-primary))] mb-6">
					MCP Server
					<span class="bg-gradient-to-r from-[rgb(var(--color-primary-500))] to-[rgb(var(--color-primary-700))] bg-clip-text text-transparent">
						Powered
					</span>
				</h2>
				<p class="text-xl text-[rgb(var(--color-text-secondary))] max-w-3xl mx-auto">
					Integrate BSpec directly into your AI workflows with our Model Context Protocol server. 
					Provides real-time access to your business knowledge graph.
				</p>
			</div>

			<!-- MCP Features Grid -->
			<div class="grid md:grid-cols-3 gap-8 mb-12">
				{#each mcpFeatures as feature, i}
					<div class="animate-scale-in animate-stagger-{i+1}">
						<GlassCard variant="medium" elevation={2} className="h-full glass-highlight">
							<div class="p-8">
								<div class="w-12 h-12 bg-gradient-to-br from-[rgb(var(--color-primary-500))] to-[rgb(var(--color-primary-600))] rounded-lg flex items-center justify-center mb-6 elevation-1">
									<svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
										<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="{getIcon(feature.icon)}" />
									</svg>
								</div>
								<h3 class="text-lg font-semibold text-[rgb(var(--color-text-primary))] mb-3">{feature.title}</h3>
								<p class="text-[rgb(var(--color-text-secondary))] leading-relaxed">{feature.description}</p>
							</div>
						</GlassCard>
					</div>
				{/each}
			</div>

			<!-- Code Example -->
			<div class="max-w-4xl mx-auto animate-slide-up">
				<GlassCard variant="heavy" elevation={3} className="glass-highlight">
					<div class="p-8">
						<div class="flex items-center gap-3 mb-6">
							<div class="w-3 h-3 bg-red-500 rounded-full"></div>
							<div class="w-3 h-3 bg-yellow-500 rounded-full"></div>
							<div class="w-3 h-3 bg-green-500 rounded-full"></div>
							<span class="ml-3 text-sm text-[rgb(var(--color-text-tertiary))] font-mono">mcp.bspec.dev</span>
						</div>
						<pre class="text-sm text-[rgb(var(--color-text-secondary))] overflow-x-auto font-mono leading-relaxed"><code>{`{
  "mcpServers": {
    "bspec": {
      "command": "npx",
      "args": ["-y", "@bspec/mcp-server"]
    }
  }
}`}</code></pre>
						<div class="mt-6 flex items-center gap-4">
							<span class="text-sm text-[rgb(var(--color-text-tertiary))]">Ready for Claude Desktop, Cline, and other MCP clients</span>
							<div class="flex items-center gap-2">
								<div class="w-2 h-2 bg-[rgb(var(--color-success))] rounded-full animate-pulse"></div>
								<span class="text-xs text-[rgb(var(--color-success))] font-medium">LIVE</span>
							</div>
						</div>
					</div>
				</GlassCard>
			</div>
		</div>
	</section>

	<!-- Call to Action Section -->
	<section class="py-24 bg-gradient-to-br from-[rgb(var(--color-primary-500)/.05)] to-[rgb(var(--color-primary-600)/.1)]">
		<div class="max-w-4xl mx-auto px-6 text-center">
			<div class="animate-slide-up">
				<h2 class="text-4xl md:text-5xl font-bold text-[rgb(var(--color-text-primary))] mb-6">
					Ready to Transform Your
					<span class="bg-gradient-to-r from-[rgb(var(--color-primary-500))] to-[rgb(var(--color-primary-700))] bg-clip-text text-transparent">
						Business Documentation?
					</span>
				</h2>
				<p class="text-xl text-[rgb(var(--color-text-secondary))] mb-12 max-w-2xl mx-auto">
					Join the movement toward structured, AI-ready business specifications.
				</p>

				<div class="flex flex-col sm:flex-row gap-4 justify-center items-center">
					<a
						href="/spec"
						class="inline-flex items-center px-8 py-4 bg-gradient-to-r from-[rgb(var(--color-primary-500))] to-[rgb(var(--color-primary-600))] text-white font-semibold rounded-2xl shadow-lg hover:shadow-xl hover:scale-105 transition-all duration-300 elevation-2"
					>
						<span>Read the Specification</span>
						<svg class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
							<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6.253v13m0-13C10.832 5.477 9.246 5 7.5 5S4.168 5.477 3 6.253v13C4.168 18.477 5.754 18 7.5 18s3.332.477 4.5 1.253m0-13C13.168 5.477 14.754 5 16.5 5c1.747 0 3.332.477 4.5 1.253v13C19.832 18.477 18.247 18 16.5 18c-1.746 0-3.332.477-4.5 1.253" />
						</svg>
					</a>
					<a
						href="/domains"
						class="inline-flex items-center px-8 py-4 bg-[rgb(var(--color-bg-primary))] text-[rgb(var(--color-text-primary))] font-semibold rounded-2xl border border-[rgb(var(--color-border-medium))] hover:bg-[rgb(var(--color-bg-hover))] hover:scale-105 transition-all duration-300 glass-medium"
					>
						<span>Explore Document Types</span>
						<svg class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
							<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 11H5m14 0a2 2 0 012 2v6a2 2 0 01-2 2H5a2 2 0 01-2-2v-6a2 2 0 012-2m14 0V9a2 2 0 00-2-2M5 11V9a2 2 0 012-2m0 0V5a2 2 0 012-2h6a2 2 0 012 2v2M7 7h10" />
						</svg>
					</a>
				</div>
			</div>
		</div>
	</section>

	<!-- Footer -->
	<footer class="py-12 border-t border-[rgb(var(--color-border-light))] bg-[rgb(var(--color-bg-secondary)/.5)]">
		<div class="max-w-6xl mx-auto px-6">
			<div class="flex flex-col md:flex-row justify-between items-center">
				<div class="flex items-center mb-4 md:mb-0">
					<div class="w-8 h-8 bg-gradient-to-br from-[rgb(var(--color-primary-500))] to-[rgb(var(--color-primary-700))] rounded-lg flex items-center justify-center mr-3">
						<img src="/bspec-icon.png" alt="BSpec" class="w-5 h-5 object-cover" />
					</div>
					<span class="font-semibold text-[rgb(var(--color-text-primary))]">BSpec Foundation</span>
				</div>
				<div class="text-sm text-[rgb(var(--color-text-tertiary))]">
					Enabling systematic business design through structured knowledge representation
				</div>
			</div>
		</div>
	</footer>
</main>
