<script lang="ts">
	import { page } from '$app/stores';
	import { bspecStore } from '$lib/bspec-store.svelte';
	import { findDocument } from '$lib/bspec-parser.svelte';
	import {
		processMarkdownContent,
		generateBreadcrumbs,
		generatePageMetadata,
		findRelatedDocuments
	} from '$lib/content-processor.svelte';

	const domainSlug = $derived($page.params.domain);
	const docTypeSlug = $derived($page.params.doctype);

	const document = $derived(bspecStore.specification && domainSlug && docTypeSlug
		? findDocument(bspecStore.specification, domainSlug, docTypeSlug)
		: null);

	const domain = $derived(bspecStore.specification?.domains.find(d => d.slug === domainSlug));

	const processedContent = $derived(document ? processMarkdownContent(document.content) : null);

	const breadcrumbs = $derived(bspecStore.specification && domainSlug && docTypeSlug
		? generateBreadcrumbs(bspecStore.specification, domainSlug, docTypeSlug)
		: []);

	const metadata = $derived(generatePageMetadata(document || undefined, domainSlug, document?.name));

	const relatedDocs = $derived(document && bspecStore.specification
		? findRelatedDocuments(document, bspecStore.specification)
		: null);
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
				<p class="mt-2 text-gray-600">Loading document...</p>
			</div>
		{:else if bspecStore.error}
			<div class="bg-red-50 border border-red-200 rounded-lg p-4">
				<p class="text-red-800">Error: {bspecStore.error}</p>
			</div>
		{:else if document && processedContent}
			<div class="max-w-6xl mx-auto">
				<div class="grid lg:grid-cols-4 gap-8">
					<!-- Main Content -->
					<main class="lg:col-span-3">
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

							<!-- Document Badge -->
							<div class="flex items-center gap-4 mb-4">
								<span class="bg-blue-100 text-blue-800 text-sm font-mono px-3 py-1 rounded-full">
									{document.type}
								</span>
								{#if document.frontmatter.status}
									<span class="text-xs px-2 py-1 rounded-full {
										document.frontmatter.status === 'Draft' ? 'bg-yellow-100 text-yellow-800' :
										document.frontmatter.status === 'Review' ? 'bg-orange-100 text-orange-800' :
										document.frontmatter.status === 'Accepted' ? 'bg-green-100 text-green-800' :
										'bg-gray-100 text-gray-800'
									}">
										{document.frontmatter.status}
									</span>
								{/if}
							</div>

							<h1 class="text-4xl font-bold text-gray-900 mb-4">{document.name}</h1>

							<!-- Meta Info -->
							<div class="flex flex-wrap items-center gap-4 text-sm text-gray-600 mb-6">
								<span>{processedContent.wordCount} words</span>
								<span>•</span>
								<span>{processedContent.readingTime} min read</span>
								{#if domain}
									<span>•</span>
									<span>Domain: {domain.name}</span>
								{/if}
								{#if document.frontmatter.priority}
									<span>•</span>
									<span>Priority: {document.frontmatter.priority}</span>
								{/if}
							</div>

							<!-- Description -->
							{#if document.frontmatter.description}
								<p class="text-lg text-gray-600 mb-6">
									{document.frontmatter.description}
								</p>
							{/if}
						</header>

						<!-- Content -->
						<article class="bg-white rounded-lg shadow-sm border p-8 mb-8">
							<div class="prose prose-lg max-w-none prose-blue prose-headings:scroll-mt-16">
								{@html processedContent.content}
							</div>
						</article>

						<!-- Related Documents -->
						{#if relatedDocs && (relatedDocs.dependsOn.length || relatedDocs.enables.length || relatedDocs.related.length)}
							<section class="bg-white rounded-lg shadow-sm border p-6 mb-8">
								<h2 class="text-xl font-semibold mb-4">Related Documents</h2>

								{#if relatedDocs.dependsOn.length > 0}
									<div class="mb-4">
										<h3 class="text-lg font-medium text-gray-900 mb-2">Depends On</h3>
										<div class="grid md:grid-cols-2 gap-2">
											{#each relatedDocs.dependsOn as dep}
												<a
													href="/domains/{dep.slug}/{dep.slug}"
													class="text-blue-600 hover:text-blue-800 text-sm"
												>
													{dep.type}: {dep.name}
												</a>
											{/each}
										</div>
									</div>
								{/if}

								{#if relatedDocs.enables.length > 0}
									<div class="mb-4">
										<h3 class="text-lg font-medium text-gray-900 mb-2">Enables</h3>
										<div class="grid md:grid-cols-2 gap-2">
											{#each relatedDocs.enables as enables}
												<a
													href="/domains/{enables.slug}/{enables.slug}"
													class="text-green-600 hover:text-green-800 text-sm"
												>
													{enables.type}: {enables.name}
												</a>
											{/each}
										</div>
									</div>
								{/if}

								{#if relatedDocs.related.length > 0}
									<div>
										<h3 class="text-lg font-medium text-gray-900 mb-2">Related</h3>
										<div class="grid md:grid-cols-2 gap-2">
											{#each relatedDocs.related as related}
												<a
													href="/domains/{related.slug}/{related.slug}"
													class="text-purple-600 hover:text-purple-800 text-sm"
												>
													{related.type}: {related.name}
												</a>
											{/each}
										</div>
									</div>
								{/if}
							</section>
						{/if}

						<!-- Navigation -->
						<nav class="flex justify-between">
							<a
								href="/domains/{domainSlug}"
								class="inline-flex items-center px-4 py-2 text-gray-600 hover:text-blue-600 transition-colors"
							>
								<svg class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
									<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7" />
								</svg>
								Back to {domain?.name}
							</a>
							<a
								href="/domains"
								class="inline-flex items-center px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors"
							>
								All Domains
								<svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
									<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
								</svg>
							</a>
						</nav>
					</main>

					<!-- Sidebar -->
					<aside class="lg:col-span-1">
						<!-- Table of Contents -->
						{#if processedContent.headings.length > 0}
							<div class="bg-white rounded-lg shadow-sm border p-6 mb-6 sticky top-8">
								<h3 class="text-lg font-semibold mb-4">Table of Contents</h3>
								<nav class="space-y-1">
									{#each processedContent.headings as heading}
										<a
											href="#{heading.id}"
											class="block py-1 text-gray-600 hover:text-blue-600 transition-colors text-sm"
											style="padding-left: {(heading.level - 1) * 12}px"
										>
											{heading.text}
										</a>
										{#each heading.children as child}
											<a
												href="#{child.id}"
												class="block py-1 text-gray-500 hover:text-blue-600 transition-colors text-xs"
												style="padding-left: {(child.level - 1) * 12}px"
											>
												{child.text}
											</a>
										{/each}
									{/each}
								</nav>
							</div>
						{/if}

						<!-- Document Info -->
						<div class="bg-white rounded-lg shadow-sm border p-6">
							<h3 class="text-lg font-semibold mb-4">Document Info</h3>
							<dl class="space-y-2 text-sm">
								<div>
									<dt class="font-medium text-gray-900">Type</dt>
									<dd class="text-gray-600">{document.type}</dd>
								</div>
								<div>
									<dt class="font-medium text-gray-900">Domain</dt>
									<dd class="text-gray-600">{domain?.name}</dd>
								</div>
								{#if document.frontmatter.version}
									<div>
										<dt class="font-medium text-gray-900">Version</dt>
										<dd class="text-gray-600">{document.frontmatter.version}</dd>
									</div>
								{/if}
								{#if document.frontmatter.owner}
									<div>
										<dt class="font-medium text-gray-900">Owner</dt>
										<dd class="text-gray-600">{document.frontmatter.owner}</dd>
									</div>
								{/if}
								{#if document.frontmatter.updated}
									<div>
										<dt class="font-medium text-gray-900">Last Updated</dt>
										<dd class="text-gray-600">{document.frontmatter.updated}</dd>
									</div>
								{/if}
							</dl>
						</div>
					</aside>
				</div>
			</div>
		{:else}
			<!-- Document Not Found -->
			<div class="text-center py-16">
				<h1 class="text-2xl font-bold text-gray-900 mb-4">Document Not Found</h1>
				<p class="text-gray-600 mb-6">
					The document "{docTypeSlug}" could not be found in the "{domainSlug}" domain.
				</p>
				<div class="space-x-4">
					{#if domainSlug}
						<a
							href="/domains/{domainSlug}"
							class="inline-flex items-center px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors"
						>
							Back to {domain?.name || 'Domain'}
						</a>
					{/if}
					<a
						href="/domains"
						class="inline-flex items-center px-4 py-2 text-gray-600 hover:text-blue-600 transition-colors"
					>
						All Domains
					</a>
				</div>
			</div>
		{/if}
	</div>
</div>