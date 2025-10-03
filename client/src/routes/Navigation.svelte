<script lang="ts">
	import { resolve } from '$app/paths';
	import Drawer from '$lib/components/Drawer.svelte';
	import { AppBar } from '@skeletonlabs/skeleton-svelte';
	import { MediaQuery } from 'svelte/reactivity';
	import { links } from './links';

	let isDrawerOpen = false;
	let isSmallDevice = new MediaQuery('(max-width: 480px)');
</script>

<Drawer isOpen={isDrawerOpen} onClose={() => (isDrawerOpen = false)} />
<AppBar>
	{#snippet lead()}
		<a
			aria-label="Go to home page"
			data-testid="home-link"
			class="btn-icon preset-filled-surface-500"
			href={resolve('/')}><i class="fa-solid fa-house"></i></a
		>
	{/snippet}
	{#snippet trail()}
		{#if isSmallDevice.current}
			<button
				aria-label="Open navigation menu"
				data-testid="hamburger"
				class="btn-icon preset-filled-surface-500"
				on:click={() => (isDrawerOpen = !isDrawerOpen)}
			>
				<i class="fas fa-bars"></i>
			</button>
		{:else}
			{#each links as link}
				<a
					aria-label={link.ariaLabel}
					data-testid={link.datatestid}
					class="btn-icon preset-filled-surface-500"
					href={link.href}><i class={link.iconClass}></i></a
				>
			{/each}
		{/if}
	{/snippet}
</AppBar>
