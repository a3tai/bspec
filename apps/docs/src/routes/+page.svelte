<script lang="ts">
	import type { PageData } from './$types';
	import Hero from '$lib/components/Hero.svelte';
	import GlassCard from '$lib/components/GlassCard.svelte';
	import DomainCard from '$lib/components/DomainCard.svelte';

	let { data }: { data: PageData } = $props();

	const stats = [
		{ value: data.domains.length.toString(), label: 'Business Domains' },
		{ value: data.totalFiles.toString(), label: 'Document Types' },
		{ value: `v${data.version}`, label: 'Specification Version' }
	];

	// Show top 6 domains on homepage
	const featuredDomains = data.domains.slice(0, 6);
</script>

<svelte:head>
	<title>{data.title}</title>
	<meta name="description" content={data.description} />
</svelte:head>

<!-- Main Content -->
<main class="flex-1 min-w-0 overflow-y-auto">
	<!-- Hero Section -->
	<Hero
		title="BSpec Documentation"
		subtitle="Business Specification Standard"
		description="A structured, machine-readable knowledge graph for describing any business"
		{stats}
		primaryAction={{ label: 'Explore Domains', href: '/domains' }}
		secondaryAction={{ label: 'Read Specification', href: '/spec' }}
	/>

	<!-- Content Section -->
	{#if data.content}
		<section class="max-w-4xl mx-auto px-6 py-12">
			<GlassCard variant="light" elevation={1} padding="lg" className="animate-fade-in">
				<article class="prose prose-lg max-w-none">
					{@html data.content}
				</article>
			</GlassCard>
		</section>
	{/if}

	<!-- Featured Domains -->
	<section class="max-w-6xl mx-auto px-6 py-12">
		<div class="mb-8 animate-slide-up">
			<h2 class="text-3xl font-bold text-[rgb(var(--color-text-primary))] mb-3">
				Explore Business Domains
			</h2>
			<p class="text-lg text-[rgb(var(--color-text-secondary))]">
				Comprehensive coverage of every aspect of business specification
			</p>
		</div>

		<div class="grid md:grid-cols-2 lg:grid-cols-3 gap-6 mb-8">
			{#each featuredDomains as domain, i}
				<div class="animate-scale-in animate-stagger-{(i % 5) + 1}">
					<DomainCard {domain} previewCount={3} />
				</div>
			{/each}
		</div>

		{#if data.domains.length > 6}
			<div class="text-center animate-fade-in">
				<a
					href="/domains"
					class="inline-flex items-center gap-2 text-[rgb(var(--color-link))] hover:text-[rgb(var(--color-link-hover))] font-medium transition-colors"
				>
					<span>View all {data.domains.length} domains</span>
					<svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
						<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
					</svg>
				</a>
			</div>
		{/if}
	</section>

	<!-- Quick Links -->
	<section class="max-w-6xl mx-auto px-6 py-12">
		<div class="grid md:grid-cols-3 gap-6">
			<GlassCard href="/spec" variant="medium" elevation={2} hover className="animate-slide-up animate-stagger-1">
				<div class="flex items-start gap-4">
					<div class="w-12 h-12 bg-gradient-to-br from-[rgb(var(--color-primary-500))] to-[rgb(var(--color-primary-700))] rounded-lg flex items-center justify-center flex-shrink-0 elevation-2">
						<svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
							<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
						</svg>
					</div>
					<div class="flex-1">
						<h3 class="font-semibold text-[rgb(var(--color-text-primary))] mb-2">Specification</h3>
						<p class="text-sm text-[rgb(var(--color-text-secondary))]">Complete BSpec v{data.version} standard</p>
					</div>
				</div>
			</GlassCard>

			<GlassCard href="/index" variant="medium" elevation={2} hover className="animate-slide-up animate-stagger-2">
				<div class="flex items-start gap-4">
					<div class="w-12 h-12 bg-gradient-to-br from-[rgb(var(--color-primary-500))] to-[rgb(var(--color-primary-700))] rounded-lg flex items-center justify-center flex-shrink-0 elevation-2">
						<svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
							<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 10h16M4 14h16M4 18h16" />
						</svg>
					</div>
					<div class="flex-1">
						<h3 class="font-semibold text-[rgb(var(--color-text-primary))] mb-2">Full Index</h3>
						<p class="text-sm text-[rgb(var(--color-text-secondary))]">Complete navigation guide</p>
					</div>
				</div>
			</GlassCard>

			<GlassCard href="/format" variant="medium" elevation={2} hover className="animate-slide-up animate-stagger-3">
				<div class="flex items-start gap-4">
					<div class="w-12 h-12 bg-gradient-to-br from-[rgb(var(--color-primary-500))] to-[rgb(var(--color-primary-700))] rounded-lg flex items-center justify-center flex-shrink-0 elevation-2">
						<svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
							<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 20l4-16m4 4l4 4-4 4M6 16l-4-4 4-4" />
						</svg>
					</div>
					<div class="flex-1">
						<h3 class="font-semibold text-[rgb(var(--color-text-primary))] mb-2">Format Guide</h3>
						<p class="text-sm text-[rgb(var(--color-text-secondary))]">Learn the document format</p>
					</div>
				</div>
			</GlassCard>
		</div>
	</section>
</main>
