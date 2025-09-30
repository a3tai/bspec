<script lang="ts">
	import '../app.css';
	import { onMount } from 'svelte';
	import { page } from '$app/state';
	import { initializeTheme } from '$lib/theme-store.svelte';
	import ThemeToggle from '$lib/components/ThemeToggle.svelte';
	import Sidebar from '$lib/components/Sidebar.svelte';
	import type { LayoutData } from './$types';

	let { children, data }: { children: any; data: LayoutData } = $props();

	const currentPath = $derived(page.url.pathname);
	let mobileMenuOpen = $state(false);
	let sidebarOpen = $state(false);

	onMount(() => {
		initializeTheme();
	});

	function isActiveRoute(path: string): boolean {
		if (path === '/' && currentPath === '/') return true;
		if (path !== '/' && currentPath.startsWith(path)) return true;
		return false;
	}
</script>

<svelte:head>
	<!-- Favicons -->
	<link rel="icon" href="/favicon.ico" />
	<link rel="icon" type="image/png" sizes="32x32" href="/favicon-32x32.png" />
	<link rel="icon" type="image/png" sizes="16x16" href="/favicon-16x16.png" />
	<link rel="apple-touch-icon" sizes="180x180" href="/apple-touch-icon.png" />

	<!-- Web App Manifest -->
	<link rel="manifest" href="/site.webmanifest" />

	<!-- Theme Color -->
	<meta name="theme-color" content="#ea580c" />
	<meta name="msapplication-TileColor" content="#ea580c" />
</svelte:head>

<div class="h-screen bg-[rgb(var(--color-bg-primary))] transition-colors flex flex-col">
	<!-- Header -->
	<header class="border-b border-[rgb(var(--color-border-light))] z-50 bg-[rgb(var(--color-bg-primary)/.8)] backdrop-blur-lg flex-shrink-0">
		<div class="px-6">
			<div class="flex items-center justify-between h-16">
				<!-- Logo -->
				<div class="flex items-center">
					<a href="/" class="flex items-center space-x-3">
						<div class="relative w-9 h-9 rounded-lg bg-gradient-to-br from-[#bd5e21] via-[#d16a26] to-[#e8762c] dark:from-[#bd5e21] dark:via-[#d16a26] dark:to-[#f2842f] shadow-lg overflow-hidden">
							<img src="/bspec-icon.png" alt="BSpec" class="w-full h-full object-cover drop-shadow-sm" />
						</div>
						<div class="flex items-center space-x-2">
							<span class="text-xl font-headers font-bold text-[rgb(var(--color-text-primary))]">BSpec</span>
						</div>
					</a>
				</div>

				<!-- Right side utilities -->
				<div class="flex items-center space-x-4">
					<!-- Sidebar toggle (desktop) -->
					<button
						class="hidden lg:flex p-2 rounded-md text-[rgb(var(--color-text-tertiary))] hover:text-[rgb(var(--color-text-primary))] hover:bg-[rgb(var(--color-bg-hover))] transition-colors"
						onclick={() => sidebarOpen = !sidebarOpen}
						aria-label="Toggle sidebar"
					>
						<svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
							<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h8M4 18h16" />
						</svg>
					</button>

					<!-- Theme Toggle -->
					<ThemeToggle />

					<!-- Mobile menu button -->
					<button
						class="lg:hidden p-2 rounded-md text-[rgb(var(--color-text-tertiary))] hover:text-[rgb(var(--color-text-primary))] hover:bg-[rgb(var(--color-bg-hover))] transition-colors"
						onclick={() => mobileMenuOpen = !mobileMenuOpen}
						aria-label="Toggle mobile menu"
					>
						<svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
							<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16" />
						</svg>
					</button>
				</div>
			</div>
		</div>
	</header>

	<!-- Main Content with Sidebar -->
	<div class="flex flex-1 overflow-hidden">
		<!-- Sidebar - Always visible on desktop, toggleable on mobile -->
		{#if data.specification}
			<aside class="hidden lg:block lg:w-80 lg:flex-shrink-0 lg:border-r lg:border-[rgb(var(--color-border-light))] lg:overflow-y-auto">
				<Sidebar specification={data.specification} />
			</aside>

			<!-- Mobile sidebar overlay -->
			<aside class="fixed inset-y-0 left-0 z-40 w-80 transform transition-transform duration-300 ease-in-out lg:hidden {mobileMenuOpen ? 'translate-x-0' : '-translate-x-full'} mt-16">
				<Sidebar specification={data.specification} />
			</aside>
		{/if}

		<!-- Mobile overlay backdrop -->
		{#if mobileMenuOpen}
			<div
				class="fixed inset-0 z-30 bg-[rgb(var(--color-neutral-700)/.5)] lg:hidden mt-16"
				role="button"
				tabindex="0"
				onclick={() => {
					mobileMenuOpen = false;
				}}
				onkeydown={(e) => {
					if (e.key === 'Enter' || e.key === ' ') {
						mobileMenuOpen = false;
					}
				}}
				aria-label="Close sidebar"
			></div>
		{/if}

		<!-- Main content area -->
		<main class="flex-1 overflow-y-auto">
			{@render children()}
		</main>
	</div>
</div>
