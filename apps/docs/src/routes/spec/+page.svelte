<script lang="ts">
	import specData from '$lib/data/spec.json';

	let data = specData;
	let content = data.content;
</script>

<svelte:head>
	<title>{data.title}</title>
	<meta name="description" content={data.description} />
</svelte:head>

<div class="max-w-4xl mx-auto px-6 py-8">
	<div class="mb-8">
		<h1 class="text-3xl font-headers font-bold text-[rgb(var(--color-text-primary))] mb-4">
			{data.title}
		</h1>
		<p class="text-lg text-[rgb(var(--color-text-secondary))]">
			{data.description}
		</p>

		{#if data.specification?.version}
			<div class="mt-4">
				<span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-[rgb(var(--color-primary-100))] text-[rgb(var(--color-primary-700))] dark:bg-[rgb(var(--color-primary-900))] dark:text-[rgb(var(--color-primary-300))]">
					Version {data.specification.version}
				</span>
			</div>
		{/if}
	</div>

	<div class="prose prose-base max-w-none">
		{#if content}
			{@html content}
		{:else}
			<p>No specification content available</p>
		{/if}
	</div>
</div>