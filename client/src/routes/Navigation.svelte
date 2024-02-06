<script lang="ts">
	import { base } from '$app/paths';
	import { AppBar, getDrawerStore } from '@skeletonlabs/skeleton';
	import MediaQuery from 'svelte-media-queries';
	import { links } from './links';

	const drawerStore = getDrawerStore();
	const openDrawer = () => {
		drawerStore.open({
			regionDrawer: 'w-64',
			position: 'right',
			id: 'mobile-nav'
		});
	};

	let isSmallDevice: boolean;
</script>

<MediaQuery query="(max-width: 480px)" bind:matches={isSmallDevice} />
<AppBar>
	<svelte:fragment slot="lead">
		<a data-testid="home-link" class="btn-icon variant-filled-surface" href="{base}/"
			><i class="fa-solid fa-house" /></a
		>
	</svelte:fragment>
	<svelte:fragment slot="trail">
		{#if isSmallDevice}
			<button data-testid="hamburger" class="btn-icon variant-filled-surface" on:click={openDrawer}>
				<i class="fas fa-bars" />
			</button>
		{:else}
			{#each links as link}
				<a data-testid={link.datatestid} class="btn-icon variant-filled-surface" href={link.href}
					><i class={link.iconClass} /></a
				>
			{/each}
		{/if}
	</svelte:fragment>
</AppBar>
