<script lang="ts">
	import { page } from '$app/stores';
	import { bspecStore } from '$lib/bspec-store.svelte';
	import { generateBreadcrumbs, generatePageMetadata } from '$lib/content-processor.svelte';
	import Sidebar from '$lib/components/Sidebar.svelte';

	const domainSlug = $derived($page.params.domain);
	const domain = $derived(bspecStore.specification?.domains.find(d => d.slug === domainSlug));
	const breadcrumbs = $derived(bspecStore.specification ? generateBreadcrumbs(bspecStore.specification, domainSlug) : []);
	const metadata = $derived(generatePageMetadata(undefined, domainSlug, domain?.name));
</script>

<svelte:head>
	<title>{metadata.title}</title>
	<meta name="description" content={metadata.description} />
	<meta name="keywords" content={metadata.keywords.join(', ')} />
</svelte:head>

<div class="min-h-screen bg-gray-50">
	<div class="container mx-auto px-4 py-8">
		{#if bspecStore.loading}
			<div class="text-center py-16">
				<div class="inline-block animate-spin rounded-full h-8 w-8 border-b-2 border-blue-600"></div>
				<p class="mt-2 text-gray-600">Loading domain...</p>
			</div>
		{:else if bspecStore.error}
			<div class="bg-red-50 border border-red-200 rounded-lg p-4">
				<p class="text-red-800">Error: {bspecStore.error}</p>
			</div>
		{:else if domain}
			<!-- Header -->
			<header class="mb-8">
				<!-- Breadcrumbs -->
				<nav class="text-sm text-gray-500 mb-4">
					{#each breadcrumbs as crumb, index}
						{#if index === breadcrumbs.length - 1}
							<span>{crumb.name}</span>
						{:else}
							<a href={crumb.path} class="hover:text-blue-600">{crumb.name}</a>
							<span class="mx-2">/</span>
						{/if}
					{/each}
				</nav>

				<h1 class="text-4xl font-bold text-gray-900 mb-4">{domain.name}</h1>
				<p class="text-xl text-gray-600 max-w-3xl">
					Comprehensive documentation for the {domain.name.toLowerCase()} domain,
					covering {domain.documents.length} document types that define this critical business area.
				</p>
			</header>

			<!-- Domain Stats -->
			<div class="grid md:grid-cols-3 gap-6 mb-8">
				<div class="bg-white rounded-lg shadow-sm border p-6">
					<h3 class="text-lg font-semibold text-gray-900 mb-2">Document Types</h3>
					<div class="text-3xl font-bold text-blue-600">{domain.documents.length}</div>
				</div>
				<div class="bg-white rounded-lg shadow-sm border p-6">
					<h3 class="text-lg font-semibold text-gray-900 mb-2">Domain</h3>
					<div class="text-lg text-gray-700">{domain.name}</div>
				</div>
				<div class="bg-white rounded-lg shadow-sm border p-6">
					<h3 class="text-lg font-semibold text-gray-900 mb-2">Coverage</h3>
					<div class="text-lg text-gray-700">Complete</div>
				</div>
			</div>

			<!-- Document Types -->
			<section class="mb-8">
				<h2 class="text-2xl font-bold text-gray-900 mb-6">Document Types</h2>
				<div class="grid md:grid-cols-2 lg:grid-cols-3 gap-6">
					{#each domain.documents as document}
						<a
							href="/domains/{domainSlug}/{document.slug}"
							class="bg-white rounded-lg shadow-sm border p-6 hover:shadow-lg hover:border-blue-300 transition-all duration-200 group"
						>
							<!-- Document Type Badge -->
							<div class="flex items-center mb-4">
								<span class="bg-blue-100 text-blue-800 text-sm font-mono px-3 py-1 rounded-full">
									{document.type}
								</span>
							</div>

							<h3 class="text-lg font-semibold text-gray-900 mb-2 group-hover:text-blue-700">
								{document.name}
							</h3>

							<!-- Frontmatter Info -->
							{#if document.frontmatter.description}
								<p class="text-gray-600 text-sm mb-4 line-clamp-2">
									{document.frontmatter.description}
								</p>
							{/if}

							<!-- Meta Info -->
							<div class="text-xs text-gray-500 space-y-1">
								{#if document.frontmatter.status}
									<div class="flex items-center">
										<span class="w-2 h-2 rounded-full mr-2 {
											document.frontmatter.status === 'Draft' ? 'bg-yellow-400' :
											document.frontmatter.status === 'Review' ? 'bg-orange-400' :
											document.frontmatter.status === 'Accepted' ? 'bg-green-400' :
											'bg-gray-400'
										}"></span>
										Status: {document.frontmatter.status}
									</div>
								{/if}
								{#if document.frontmatter.priority}
									<div>Priority: {document.frontmatter.priority}</div>
								{/if}
							</div>

							<div class="mt-4 flex items-center text-blue-600 text-sm font-medium">
								View Specification
								<svg class="w-4 h-4 ml-1 group-hover:translate-x-1 transition-transform" fill="none" stroke="currentColor" viewBox="0 0 24 24">
									<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
								</svg>
							</div>
						</a>
					{/each}
				</div>
			</section>

			<!-- Navigation -->
			<nav class="flex justify-between">
				<a
					href="/domains"
					class="inline-flex items-center px-4 py-2 text-gray-600 hover:text-blue-600 transition-colors"
				>
					<svg class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
						<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7" />
					</svg>
					All Domains
				</a>
				<a
					href="/spec"
					class="inline-flex items-center px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors"
				>
					Read Specification
					<svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
						<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
					</svg>
				</a>
			</nav>
		{:else}
			<!-- Domain Not Found -->
			<div class="text-center py-16">
				<h1 class="text-2xl font-bold text-gray-900 mb-4">Domain Not Found</h1>
				<p class="text-gray-600 mb-6">
					The domain "{domainSlug}" could not be found in the loaded specification.
				</p>
				<div class="space-x-4">
					<a
						href="/domains"
						class="inline-flex items-center px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors"
					>
						View All Domains
					</a>
					<a
						href="/"
						class="inline-flex items-center px-4 py-2 text-gray-600 hover:text-blue-600 transition-colors"
					>
						Go Home
					</a>
				</div>
			</div>
		{/if}
	</div>
</div>