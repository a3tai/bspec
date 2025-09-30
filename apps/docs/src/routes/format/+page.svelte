<script lang="ts">
	import Sidebar from '$lib/components/Sidebar.svelte';
	import type { PageData } from './$types';

	let { data }: { data: PageData } = $props();
</script>

<svelte:head>
	<title>{data.title}</title>
	<meta name="description" content={data.description} />
</svelte:head>

<!-- Sidebar -->
<Sidebar specification={data.specification} />

<!-- Main Content -->
<main class="flex-1 min-w-0">
	<div class="max-w-4xl mx-auto px-6 py-8">
		<!-- Header -->
		<header class="mb-8">
			<nav class="text-sm text-[rgb(var(--color-text-muted))] mb-4">
				<a href="/" class="hover:text-[rgb(var(--color-link-hover))]">Home</a>
				<span class="mx-2">/</span>
				<span>Format</span>
			</nav>
			<h1 class="text-3xl font-bold text-[rgb(var(--color-text-primary))] mb-4">{data.title || 'BSpec Format Guide'}</h1>
			{#if data.content}
				<div class="flex items-center gap-4 text-sm text-[rgb(var(--color-text-tertiary))]">
					{#if data.content.wordCount}
						<span>{data.content.wordCount} words</span>
						<span>•</span>
					{/if}
					{#if data.content.readingTime}
						<span>{data.content.readingTime} min read</span>
					{/if}
					{#if data.specification?.version}
						<span>•</span>
						<span>Version {data.specification.version}</span>
					{/if}
				</div>
			{/if}
		</header>

		{#if data.content}
			<!-- Table of Contents -->
			{#if data.content.headings?.length > 0}
				<aside class="bg-[rgb(var(--color-bg-primary)/.8)] backdrop-blur-lg rounded-lg border border-[rgb(var(--color-border-light))] p-6 mb-8">
					<h2 class="text-lg font-semibold text-[rgb(var(--color-text-primary))] mb-4">Table of Contents</h2>
					<nav class="space-y-1">
						{#each data.content.headings as heading}
							<a
								href="#{heading.id}"
								class="block py-1 text-[rgb(var(--color-text-secondary))] hover:text-[rgb(var(--color-link-hover))] transition-colors"
								style="padding-left: {(heading.level - 1) * 16}px"
							>
								{heading.text}
							</a>
						{/each}
					</nav>
				</aside>
			{/if}

			<!-- Content -->
			<article class="prose prose-lg max-w-none">
				{@html data.content.content}
			</article>
		{:else}
			<div class="bg-[rgb(var(--color-bg-secondary))] border border-[rgb(var(--color-border-light))] rounded-lg p-8 text-center">
				<p class="text-[rgb(var(--color-text-secondary))]">Format guide content is being prepared.</p>
			</div>
		{/if}

		<!-- Navigation -->
		<nav class="mt-8 flex justify-between">
			<a
				href="/spec"
				class="inline-flex items-center px-4 py-2 text-[rgb(var(--color-text-tertiary))] hover:text-[rgb(var(--color-link-hover))] transition-colors"
			>
				<svg class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
					<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7" />
				</svg>
				Specification
			</a>
			<a
				href="/domains"
				class="inline-flex items-center px-4 py-2 bg-[rgb(var(--color-primary-600))] text-white rounded-lg hover:bg-[rgb(var(--color-primary-700))] transition-colors"
			>
				Explore Domains
				<svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
					<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
				</svg>
			</a>
		</nav>
	</div>
</main>