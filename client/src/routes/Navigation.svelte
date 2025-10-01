<script lang="ts">
	import { resolve } from '$app/paths';
	import { AppBar } from '@skeletonlabs/skeleton-svelte';
	// import MediaQuery from 'svelte-media-queries';
	import { MediaQuery } from 'svelte/reactivity';
	import { links } from './links';

	// const drawerStore = getDrawerStore();
	// const openDrawer = () => {
	// 	drawerStore.open({
	// 		width: 'w-48',
	// 		position: 'right',
	// 		id: 'mobile-nav'
	// 	});
	// };

	let isSmallDevice = new MediaQuery('(max-width: 480px)');
</script>

<AppBar>
	{#snippet lead()}
		<a
			aria-label="Go to home page"
			data-testid="home-link"
			class="btn-icon variant-filled-surface"
			href={resolve('/')}><i class="fa-solid fa-house"></i></a
		>
	{/snippet}
	{#snippet trail()}
		{#if isSmallDevice.current}
			<button
				aria-label="Open navigation menu"
				data-testid="hamburger"
				class="btn-icon variant-filled-surface"
				on:click={() => {
					// openDrawer();
				}}
			>
				<i class="fas fa-bars"></i>
			</button>
		{:else}
			{#each links as link}
				<a
					aria-label={link.ariaLabel}
					data-testid={link.datatestid}
					class="btn-icon variant-filled-surface"
					href={link.href}><i class={link.iconClass}></i></a
				>
			{/each}
		{/if}
	{/snippet}
</AppBar>
