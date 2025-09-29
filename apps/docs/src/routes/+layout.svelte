<script lang="ts">
	import '../app.css';
	import { onMount } from 'svelte';
	import { page } from '$app/stores';
	import { initializeTheme } from '$lib/theme-store.svelte.js';
	import ThemeToggle from '$lib/components/ThemeToggle.svelte';
	import Footer from '$lib/components/Footer.svelte';
	import Navigation from '$lib/components/Navigation.svelte';
	import type { BSpecSpecification } from '$lib/bspec-parser.svelte.js';

	let { children } = $props();

	const currentPath = $derived($page.url.pathname);
	let mobileMenuOpen = $state(false);
	let sidebarOpen = $state(false);
	let specification: BSpecSpecification | null = $state(null);

	onMount(async () => {
		initializeTheme();

		// Load specification data for navigation (dev mode uses JSON, production uses TGZ)
		try {
			const response = await fetch('/src/lib/data/home.json');
			if (response.ok) {
				const data = await response.json();
				specification = data.specification;
			}
		} catch (error) {
			console.warn('Failed to load specification for navigation:', error);
		}
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

<div class="h-screen bg-white dark:bg-zinc-900 transition-colors flex flex-col">
	<!-- Header -->
	<header class="border-b border-zinc-200 dark:border-zinc-800 z-50 bg-white/80 dark:bg-zinc-900/80 backdrop-blur-lg flex-shrink-0">
		<div class="px-6">
			<div class="flex items-center justify-between h-16">
				<!-- Logo -->
				<div class="flex items-center">
					<a href="/" class="flex items-center space-x-3">
						<div class="relative w-9 h-9 rounded-lg bg-gradient-to-br from-[#bd5e21] via-[#d16a26] to-[#e8762c] dark:from-[#bd5e21] dark:via-[#d16a26] dark:to-[#f2842f] shadow-lg overflow-hidden">
							<img src="/bspec-icon.png" alt="BSpec" class="w-full h-full object-cover drop-shadow-sm" />
						</div>
						<div class="flex items-center space-x-2">
							<span class="text-xl font-headers font-bold text-zinc-900 dark:text-zinc-100">BSpec</span>
						</div>
					</a>
				</div>

				<!-- Right side utilities -->
				<div class="flex items-center space-x-4">
					<!-- Sidebar toggle (desktop) -->
					<button
						class="hidden lg:flex p-2 rounded-md text-zinc-600 dark:text-zinc-400 hover:text-zinc-900 dark:hover:text-zinc-100 hover:bg-zinc-100 dark:hover:bg-zinc-800 transition-colors"
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
						class="lg:hidden p-2 rounded-md text-zinc-600 dark:text-zinc-400 hover:text-zinc-900 dark:hover:text-zinc-100 hover:bg-zinc-100 dark:hover:bg-zinc-800 transition-colors"
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
		<!-- Sidebar -->
		<aside class="fixed inset-y-0 left-0 z-40 w-64 transform transition-transform duration-300 ease-in-out lg:translate-x-0 lg:static lg:inset-0
			{sidebarOpen || mobileMenuOpen ? 'translate-x-0' : '-translate-x-full lg:translate-x-0'}">
			{#if specification}
				<Navigation {specification} />
			{:else}
				<div class="h-full bg-zinc-50 dark:bg-zinc-900 border-r border-zinc-200 dark:border-zinc-800 animate-pulse">
					<div class="p-4 space-y-4">
						<div class="h-4 bg-zinc-300 dark:bg-zinc-700 rounded w-3/4"></div>
						<div class="space-y-2">
							<div class="h-3 bg-zinc-200 dark:bg-zinc-800 rounded w-full"></div>
							<div class="h-3 bg-zinc-200 dark:bg-zinc-800 rounded w-5/6"></div>
							<div class="h-3 bg-zinc-200 dark:bg-zinc-800 rounded w-4/5"></div>
						</div>
					</div>
				</div>
			{/if}
		</aside>

		<!-- Mobile sidebar overlay -->
		{#if sidebarOpen || mobileMenuOpen}
			<div
				class="fixed inset-0 z-30 bg-zinc-600 bg-opacity-50 lg:hidden"
				role="button"
				tabindex="0"
				onclick={() => {
					sidebarOpen = false;
					mobileMenuOpen = false;
				}}
				onkeydown={(e) => {
					if (e.key === 'Enter' || e.key === ' ') {
						sidebarOpen = false;
						mobileMenuOpen = false;
					}
				}}
				aria-label="Close sidebar"
			></div>
		{/if}

		<!-- Main content area -->
		<main class="flex-1 lg:ml-0 overflow-y-auto">
			<div class="relative min-h-full">
				{@render children?.()}
				<!-- Footer Component -->
				<Footer />
			</div>
		</main>
	</div>
</div>
