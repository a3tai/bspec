<script lang="ts">
	import { getNavigationStructure } from '$lib/bspec-parser.svelte.js';
	import Sidebar from '$lib/components/Sidebar.svelte';
	import type { PageData } from './$types';

	let { data }: { data: PageData } = $props();
	const navigation = $derived(getNavigationStructure(data.specification));
</script>

<svelte:head>
	<title>{metadata.title}</title>
	<meta name="description" content={metadata.description} />
	<meta name="keywords" content={metadata.keywords.join(', ')} />
</svelte:head>

{#if bspecStore.loading}
	<div class="flex-1 flex items-center justify-center min-h-[80vh]">
		<div class="text-center">
			<div class="inline-block animate-spin rounded-full h-8 w-8 border-b-2 border-gray-900"></div>
			<p class="mt-2 text-gray-600">Loading specification...</p>
		</div>
	</div>
{:else if bspecStore.error}
	<div class="flex-1 flex items-center justify-center min-h-[80vh]">
		<div class="bg-red-50 border border-red-200 rounded-lg p-6 max-w-lg">
			<p class="text-red-800">Error: {bspecStore.error}</p>
		</div>
	</div>
{:else if processedContent && navigation}
	<!-- Sidebar -->
	<Sidebar />

	<!-- Main Content -->
	<main class="flex-1 min-w-0">
		<div class="max-w-4xl mx-auto px-6 py-8">
			<!-- Header -->
			<header class="mb-8">
				<!-- <h1 class="text-3xl font-bold text-gray-900 mb-4">
					BSpec Documentation
				</h1>
				<p class="text-lg text-gray-600 mb-6">
					Universal Business Specification Standard - A structured, machine-readable knowledge graph for describing any business.
				</p> -->

				<div class="flex items-center gap-4 text-sm text-gray-600 mb-6">
					<span>{processedContent.wordCount} words</span>
					<span>•</span>
					<span>{processedContent.readingTime} min read</span>
					{#if bspecStore.specification?.version}
						<span>•</span>
						<span>Version {bspecStore.specification.version}</span>
					{/if}
				</div>
			</header>

			<!-- Content -->
			<article class="prose prose-lg max-w-none">
				{@html processedContent.content}
			</article>
		</div>
	</main>
{:else}
	<!-- File Upload Interface -->
	<div class="flex-1 flex items-center justify-center min-h-[80vh]">
		<div class="max-w-2xl mx-auto px-6 text-center">
			<h1 class="text-3xl font-bold text-gray-900 mb-4">Load BSpec Documentation</h1>
			<p class="text-gray-600 mb-8">
				Upload a BSpec .tgz file to generate documentation or use the local specification for development.
			</p>

			<input
				type="file"
				accept=".tgz,.tar.gz"
				on:change={handleFileUpload}
				bind:this={fileInput}
				class="hidden"
			/>

			<button
				on:click={() => fileInput.click()}
				class="inline-flex items-center px-6 py-3 bg-gray-900 text-white rounded-lg hover:bg-gray-800 transition-colors mb-4"
			>
				<svg class="w-5 h-5 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
					<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 011 9.9M15 13l-3-3m0 0l-3 3m3-3v12" />
				</svg>
				Upload BSpec File
			</button>

			<div class="text-sm text-gray-500">
				<p>Supports .tgz and .tar.gz files</p>
			</div>
		</div>
	</div>
{/if}
