<script lang="ts">
	import { generateBreadcrumbs, generatePageMetadata } from '$lib/content-processor';
	import type { PageData } from './$types';
	import GlassCard from '$lib/components/GlassCard.svelte';
	import DocumentCard from '$lib/components/DocumentCard.svelte';
	import GlassButton from '$lib/components/GlassButton.svelte';

	let { data }: { data: PageData } = $props();

	const breadcrumbs = $derived(data.specification ? generateBreadcrumbs(data.specification, data.domain.slug) : []);
	const metadata = $derived(generatePageMetadata(undefined, data.domain.slug, data.domain.name));
</script>

<svelte:head>
	<title>{metadata.title}</title>
	<meta name="description" content={metadata.description} />
	<meta name="keywords" content={metadata.keywords.join(', ')} />
</svelte:head>

<main class="flex-1 min-w-0 overflow-y-auto">
	<div class="max-w-6xl mx-auto px-6 py-8">
		{#if data.domain}
			<!-- Header -->
			<header class="mb-12">
				<!-- Breadcrumbs -->
				<nav class="text-sm text-[rgb(var(--color-text-muted))] mb-6 animate-fade-in">
					{#each breadcrumbs as crumb, index}
						{#if index === breadcrumbs.length - 1}
							<span class="text-[rgb(var(--color-text-primary))]">{crumb.name}</span>
						{:else}
							<a href={crumb.path} class="hover:text-[rgb(var(--color-link-hover))] transition-colors">{crumb.name}</a>
							<span class="mx-2">/</span>
						{/if}
					{/each}
				</nav>

				<div class="animate-slide-up">
					<h1 class="text-4xl md:text-5xl font-bold text-[rgb(var(--color-text-primary))] mb-4">
						{data.domain.name}
					</h1>
					{#if data.domain.description}
						<p class="text-xl text-[rgb(var(--color-text-secondary))] max-w-3xl mb-6">
							{data.domain.description}
						</p>
					{:else}
						<p class="text-xl text-[rgb(var(--color-text-secondary))] max-w-3xl mb-6">
							Comprehensive documentation for the {data.domain.name.toLowerCase()} domain,
							covering {data.domain.documents.length} document types that define this critical business area.
						</p>
					{/if}
				</div>
			</header>

			<!-- Domain Stats -->
			<div class="grid md:grid-cols-3 gap-6 mb-12">
				<GlassCard variant="medium" elevation={2} highlight className="text-center animate-scale-in animate-stagger-1">
					<div class="mb-3">
						<div class="text-4xl font-bold text-[rgb(var(--color-primary-600))] dark:text-[rgb(var(--color-primary-400))]">
							{data.domain.documents.length}
						</div>
					</div>
					<div class="text-sm text-[rgb(var(--color-text-secondary))]">Document Types</div>
				</GlassCard>

				<GlassCard variant="medium" elevation={2} highlight className="text-center animate-scale-in animate-stagger-2">
					<div class="mb-3">
						<svg class="w-12 h-12 mx-auto text-[rgb(var(--color-success))]" fill="none" stroke="currentColor" viewBox="0 0 24 24">
							<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
						</svg>
					</div>
					<div class="text-sm text-[rgb(var(--color-text-secondary))]">Complete Coverage</div>
				</GlassCard>

				<GlassCard variant="medium" elevation={2} highlight className="text-center animate-scale-in animate-stagger-3">
					<div class="mb-3">
						<svg class="w-12 h-12 mx-auto text-[rgb(var(--color-info))]" fill="none" stroke="currentColor" viewBox="0 0 24 24">
							<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
						</svg>
					</div>
					<div class="text-sm text-[rgb(var(--color-text-secondary))]">Structured Documentation</div>
				</GlassCard>
			</div>

			<!-- Document Types -->
			<section class="mb-12">
				<h2 class="text-3xl font-bold text-[rgb(var(--color-text-primary))] mb-8 animate-slide-up">
					Document Types
				</h2>
				<div class="grid md:grid-cols-2 lg:grid-cols-3 gap-6">
					{#each data.domain.documents as document, i}
						<div class="animate-scale-in animate-stagger-{(i % 5) + 1}">
							<DocumentCard {document} />
						</div>
					{/each}
				</div>
			</section>

			<!-- Navigation -->
			<nav class="flex flex-col sm:flex-row justify-between items-center gap-4 pt-8 border-t border-[rgb(var(--color-border-light))] animate-fade-in">
				<GlassButton href="/domains" variant="ghost" size="md">
					<svg class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
						<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7" />
					</svg>
					All Domains
				</GlassButton>

				<GlassButton href="/spec" variant="primary" size="md">
					Read Specification
					<svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
						<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
					</svg>
				</GlassButton>
			</nav>
		{:else}
			<!-- Domain Not Found -->
			<div class="text-center py-16 animate-fade-in">
				<GlassCard variant="medium" elevation={2} className="max-w-md mx-auto">
					<div class="w-16 h-16 mx-auto mb-6 bg-gradient-to-br from-[rgb(var(--color-error)/.2)] to-[rgb(var(--color-error)/.1)] rounded-full flex items-center justify-center">
						<svg class="w-8 h-8 text-[rgb(var(--color-error))]" fill="none" stroke="currentColor" viewBox="0 0 24 24">
							<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
						</svg>
					</div>
					<h1 class="text-2xl font-bold text-[rgb(var(--color-text-primary))] mb-4">
						Domain Not Found
					</h1>
					<p class="text-[rgb(var(--color-text-secondary))] mb-8">
						The requested domain could not be found in the loaded specification.
					</p>
					<div class="flex flex-col sm:flex-row gap-4 justify-center">
						<GlassButton href="/domains" variant="primary" size="md">
							View All Domains
						</GlassButton>
						<GlassButton href="/" variant="secondary" size="md">
							Go Home
						</GlassButton>
					</div>
				</GlassCard>
			</div>
		{/if}
	</div>
</main>