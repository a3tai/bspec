<script lang="ts">
	import { getNavigationStructure } from '$lib/navigation';
	import type { PageData } from './$types';
	import GlassCard from '$lib/components/GlassCard.svelte';
	import DomainCard from '$lib/components/DomainCard.svelte';
	import GlassButton from '$lib/components/GlassButton.svelte';

	let { data }: { data: PageData } = $props();
	const navigation = $derived(getNavigationStructure(data.specification));

	const totalDocuments = $derived(
		navigation.domains.reduce((sum, domain) => sum + domain.documents.length, 0)
	);
</script>

<svelte:head>
	<title>{data.title}</title>
	<meta name="description" content={data.description} />
</svelte:head>

<!-- Main Content -->
<main class="flex-1 min-w-0 overflow-y-auto">
	<div class="max-w-6xl mx-auto px-6 py-8">
		<!-- Breadcrumb -->
		<nav class="text-sm text-[rgb(var(--color-text-muted))] mb-6 animate-fade-in">
			<a href="/" class="hover:text-[rgb(var(--color-link-hover))] transition-colors">Home</a>
			<span class="mx-2">/</span>
			<span class="text-[rgb(var(--color-text-primary))]">Domains</span>
		</nav>

		<!-- Header -->
		<header class="mb-12 animate-slide-up">
			<h1 class="text-4xl md:text-5xl font-bold text-[rgb(var(--color-text-primary))] mb-4">
				Business Domains
			</h1>
			<p class="text-xl text-[rgb(var(--color-text-secondary))] max-w-3xl">
				BSpec organizes business documentation into {navigation.domains.length} comprehensive domains,
				covering every aspect of business specification from strategy to execution.
			</p>
		</header>

		<!-- Stats Overview -->
		<div class="grid grid-cols-1 md:grid-cols-3 gap-6 mb-12">
			<GlassCard variant="medium" elevation={2} highlight className="text-center animate-scale-in animate-stagger-1">
				<div class="text-4xl font-bold text-[rgb(var(--color-primary-600))] dark:text-[rgb(var(--color-primary-400))] mb-2">
					{navigation.domains.length}
				</div>
				<div class="text-[rgb(var(--color-text-secondary))]">Business Domains</div>
			</GlassCard>

			<GlassCard variant="medium" elevation={2} highlight className="text-center animate-scale-in animate-stagger-2">
				<div class="text-4xl font-bold text-[rgb(var(--color-success))] mb-2">
					{totalDocuments}
				</div>
				<div class="text-[rgb(var(--color-text-secondary))]">Document Types</div>
			</GlassCard>

			<GlassCard variant="medium" elevation={2} highlight className="text-center animate-scale-in animate-stagger-3">
				<div class="text-4xl font-bold text-[rgb(var(--color-info))] mb-2">
					v{data.specification?.version}
				</div>
				<div class="text-[rgb(var(--color-text-secondary))]">Specification Version</div>
			</GlassCard>
		</div>

		<!-- Domains Grid -->
		<div class="grid md:grid-cols-2 lg:grid-cols-3 gap-6 mb-12">
			{#each navigation.domains as domain, i}
				<div class="animate-scale-in animate-stagger-{(i % 5) + 1}">
					<DomainCard {domain} previewCount={3} />
				</div>
			{/each}
		</div>

		<!-- Navigation -->
		<nav class="flex flex-col sm:flex-row justify-between items-center gap-4 pt-8 border-t border-[rgb(var(--color-border-light))] animate-fade-in">
			<GlassButton href="/" variant="ghost" size="md">
				<svg class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
					<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7" />
				</svg>
				Back to Home
			</GlassButton>

			<GlassButton href="/spec" variant="primary" size="md">
				Read Specification
				<svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
					<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
				</svg>
			</GlassButton>
		</nav>
	</div>
</main>