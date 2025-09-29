<script lang="ts">
	import { onMount } from 'svelte';
	import type { BSpecSpecification } from '$lib/bspec-parser.svelte.js';

	let indexData: {
		specification: BSpecSpecification;
		content: string;
		title: string;
		description: string;
	} | null = null;
	let loading = true;
	let error: string | null = null;

	onMount(async () => {
		try {
			const response = await fetch('/src/lib/data/index.json');
			if (!response.ok) throw new Error('Failed to load index data');
			indexData = await response.json();
		} catch (e) {
			error = e instanceof Error ? e.message : 'Failed to load index data';
		} finally {
			loading = false;
		}
	});
</script>

<svelte:head>
	{#if indexData}
		<title>{indexData.title}</title>
		<meta name="description" content={indexData.description} />
	{/if}
</svelte:head>

<div class="container mx-auto px-4 py-8">
	{#if loading}
		<div class="text-center">
			<div class="animate-spin rounded-full h-32 w-32 border-b-2 border-blue-600 mx-auto"></div>
			<p class="mt-4 text-gray-600 dark:text-gray-400">Loading BSpec documentation index...</p>
		</div>
	{:else if error}
		<div class="bg-red-50 dark:bg-red-900/20 border border-red-200 dark:border-red-800 rounded-lg p-6">
			<h2 class="text-xl font-semibold text-red-800 dark:text-red-200 mb-2">Error Loading Index</h2>
			<p class="text-red-600 dark:text-red-400">{error}</p>
		</div>
	{:else if indexData}
		<div class="prose prose-lg dark:prose-invert max-w-none">
			{@html indexData.content}
		</div>

		{#if indexData.specification?.domains}
			<div class="mt-12 border-t pt-8">
				<h2 class="text-2xl font-bold mb-6">Domain Summary</h2>
				<div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
					{#each indexData.specification.domains as domain}
						<div class="bg-white dark:bg-gray-800 rounded-lg shadow-sm border border-gray-200 dark:border-gray-700 p-6">
							<h3 class="font-semibold text-lg mb-2 capitalize">
								{domain.name.replace('-', ' ')}
							</h3>
							<p class="text-sm text-gray-600 dark:text-gray-400 mb-3">
								{domain.documents?.length || 0} document types
							</p>
							<a
								href="/domains/{domain.slug || domain.name}"
								class="text-blue-600 dark:text-blue-400 hover:text-blue-800 dark:hover:text-blue-200 text-sm font-medium"
							>
								View Domain →
							</a>
						</div>
					{/each}
				</div>
			</div>
		{/if}
	{/if}
</div>

<style>
	.prose :global(h1) {
		@apply text-3xl font-bold text-gray-900 dark:text-white mb-4;
	}

	.prose :global(h2) {
		@apply text-2xl font-semibold text-gray-800 dark:text-gray-200 mt-8 mb-4;
	}

	.prose :global(h3) {
		@apply text-xl font-semibold text-gray-800 dark:text-gray-200 mt-6 mb-3;
	}

	.prose :global(a) {
		@apply text-blue-600 dark:text-blue-400 hover:text-blue-800 dark:hover:text-blue-200;
	}

	.prose :global(ul) {
		@apply list-none pl-0;
	}

	.prose :global(li) {
		@apply mb-1;
	}

	.prose :global(blockquote) {
		@apply border-l-4 border-blue-500 pl-4 italic text-gray-600 dark:text-gray-400;
	}
</style>