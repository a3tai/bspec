<script lang="ts">
	import { getNavigationStructure } from '$lib/bspec-parser.svelte.js';
	import Sidebar from '$lib/components/Sidebar.svelte';
	import type { PageData } from './$types';

	let { data }: { data: PageData } = $props();
	const navigation = $derived(getNavigationStructure(data.specification));
</script>

<svelte:head>
	<title>{data.title}</title>
	<meta name="description" content={data.description} />
</svelte:head>

<!-- Sidebar -->
<Sidebar specification={data.specification} />

<!-- Main Content -->
<main class="flex-1 min-w-0">
	<div class="max-w-6xl mx-auto px-6 py-8">
		<!-- Header -->
		<header class="mb-8">
			<nav class="text-sm text-zinc-500 dark:text-zinc-400 mb-4">
				<a href="/" class="hover:text-orange-600 dark:hover:text-orange-400">Home</a>
				<span class="mx-2">/</span>
				<span>Domains</span>
			</nav>
			<h1 class="text-3xl font-bold text-zinc-900 dark:text-zinc-100 mb-4">Business Domains</h1>
			<p class="text-xl text-zinc-600 dark:text-zinc-400 max-w-3xl">
				BSpec organizes business documentation into {navigation.domains.length} comprehensive domains,
				covering every aspect of business specification from strategy to execution.
			</p>
		</header>

		<!-- Stats Overview -->
		<div class="grid md:grid-cols-3 gap-6 mb-12">
			<div class="bg-white/80 dark:bg-zinc-800/80 backdrop-blur-lg rounded-lg shadow-sm border border-zinc-200 dark:border-zinc-700 p-6 text-center">
				<div class="text-3xl font-bold text-orange-600 dark:text-orange-400 mb-2">{navigation.domains.length}</div>
				<div class="text-zinc-600 dark:text-zinc-400">Business Domains</div>
			</div>
			<div class="bg-white/80 dark:bg-zinc-800/80 backdrop-blur-lg rounded-lg shadow-sm border border-zinc-200 dark:border-zinc-700 p-6 text-center">
				<div class="text-3xl font-bold text-green-600 dark:text-green-400 mb-2">
					{navigation.domains.reduce((sum, domain) => sum + domain.documents.length, 0)}
				</div>
				<div class="text-zinc-600 dark:text-zinc-400">Document Types</div>
			</div>
			<div class="bg-white/80 dark:bg-zinc-800/80 backdrop-blur-lg rounded-lg shadow-sm border border-zinc-200 dark:border-zinc-700 p-6 text-center">
				<div class="text-3xl font-bold text-purple-600 dark:text-purple-400 mb-2">{data.specification?.version}</div>
				<div class="text-zinc-600 dark:text-zinc-400">Specification Version</div>
			</div>
		</div>

		<!-- Domains Grid -->
		<div class="grid md:grid-cols-2 lg:grid-cols-3 gap-6 mb-8">
			{#each navigation.domains as domain}
				<a
					href={domain.path}
					class="bg-white/80 dark:bg-zinc-800/80 backdrop-blur-lg rounded-lg shadow-sm border border-zinc-200 dark:border-zinc-700 p-6 hover:shadow-lg hover:border-orange-300 dark:hover:border-orange-600 transition-all duration-200 group"
				>
					<h2 class="text-xl font-semibold text-zinc-900 dark:text-zinc-100 mb-3 group-hover:text-orange-600 dark:group-hover:text-orange-400">
						{domain.name}
					</h2>
					<p class="text-zinc-600 dark:text-zinc-400 mb-4">
						{domain.documents.length} document types covering all aspects of {domain.name.toLowerCase()}.
					</p>

					<!-- Document Types Preview -->
					<div class="space-y-2">
						{#each domain.documents.slice(0, 3) as document}
							<div class="text-sm text-zinc-500 dark:text-zinc-400 flex items-center">
								<span class="w-8 h-5 bg-zinc-100 dark:bg-zinc-700 rounded text-xs font-mono flex items-center justify-center mr-2">
									{document.type}
								</span>
								{document.name}
							</div>
						{/each}
						{#if domain.documents.length > 3}
							<div class="text-sm text-orange-600 dark:text-orange-400 font-medium">
								+{domain.documents.length - 3} more...
							</div>
						{/if}
					</div>

					<div class="mt-4 flex items-center text-orange-600 dark:text-orange-400 text-sm font-medium">
						Explore Domain
						<svg class="w-4 h-4 ml-1 group-hover:translate-x-1 transition-transform" fill="none" stroke="currentColor" viewBox="0 0 24 24">
							<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
						</svg>
					</div>
				</a>
			{/each}
		</div>

		<!-- Navigation -->
		<nav class="flex justify-between">
			<a
				href="/"
				class="inline-flex items-center px-4 py-2 text-zinc-600 dark:text-zinc-400 hover:text-orange-600 dark:hover:text-orange-400 transition-colors"
			>
				<svg class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
					<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7" />
				</svg>
				Back to Home
			</a>
			<a
				href="/spec"
				class="inline-flex items-center px-4 py-2 bg-orange-600 text-white rounded-lg hover:bg-orange-700 transition-colors"
			>
				Read Specification
				<svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
					<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
				</svg>
			</a>
		</nav>
	</div>
</main>