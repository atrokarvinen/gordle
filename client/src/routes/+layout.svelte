<script lang="ts">
	import '../app.css';
	import Navigation from './Navigation.svelte';
</script>

<!-- <Drawer>
	{#if $drawerStore.id === 'mobile-nav'}
		<MobileNavigation />
	{:else}
		<p>Unknown drawer id '{$drawerStore.id}'</p>
	{/if}
</Drawer> -->
<!-- <Toast position="t" /> -->

<Navigation />
<div class="align-center flex justify-center">
	<slot />
</div>
