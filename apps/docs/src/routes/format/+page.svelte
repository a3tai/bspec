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
			<nav class="text-sm text-zinc-500 dark:text-zinc-400 mb-4">
				<a href="/" class="hover:text-orange-600 dark:hover:text-orange-400">Home</a>
				<span class="mx-2">/</span>
				<span>Format</span>
			</nav>
			<h1 class="text-3xl font-bold text-zinc-900 dark:text-zinc-100 mb-4">BSpec Format Guide</h1>
			<div class="flex items-center gap-4 text-sm text-zinc-600 dark:text-zinc-400">
				<span>{data.content.wordCount} words</span>
				<span>•</span>
				<span>{data.content.readingTime} min read</span>
				{#if data.specification?.version}
					<span>•</span>
					<span>Version {data.specification.version}</span>
				{/if}
			</div>
		</header>

		<!-- Table of Contents -->
		{#if data.content.headings.length > 0}
			<aside class="bg-white/80 dark:bg-zinc-800/80 backdrop-blur-lg rounded-lg border border-zinc-200 dark:border-zinc-700 p-6 mb-8">
				<h2 class="text-lg font-semibold text-zinc-900 dark:text-zinc-100 mb-4">Table of Contents</h2>
				<nav class="space-y-1">
					{#each data.content.headings as heading}
						<a
							href="#{heading.id}"
							class="block py-1 text-zinc-600 dark:text-zinc-400 hover:text-orange-600 dark:hover:text-orange-400 transition-colors"
							style="padding-left: {(heading.level - 1) * 16}px"
						>
							{heading.text}
						</a>
					{/each}
				</nav>
			</aside>
		{/if}

		<!-- Content -->
		<article class="prose prose-lg max-w-none prose-zinc dark:prose-invert">
			{@html data.content.content}
		</article>

		<!-- Navigation -->
		<nav class="mt-8 flex justify-between">
			<a
				href="/spec"
				class="inline-flex items-center px-4 py-2 text-zinc-600 dark:text-zinc-400 hover:text-orange-600 dark:hover:text-orange-400 transition-colors"
			>
				<svg class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
					<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7" />
				</svg>
				Specification
			</a>
			<a
				href="/domains"
				class="inline-flex items-center px-4 py-2 bg-orange-600 text-white rounded-lg hover:bg-orange-700 transition-colors"
			>
				Explore Domains
				<svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
					<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
				</svg>
			</a>
		</nav>
	</div>
</main>