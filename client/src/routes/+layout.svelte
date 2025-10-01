<script lang="ts">
	import Navigation from './Navigation.svelte';
</script>

<!-- <Drawer>
	{#if $drawerStore.id === 'mobile-nav'}
		<MobileNavigation />
	{:else}
		<p>Unknown drawer id '{$drawerStore.id}'</p>
	{/if}
</Drawer> -->
<!-- <Modal components={modalRegistry} /> -->
<!-- <Toast position="t" /> -->

<Navigation />
<slot />
