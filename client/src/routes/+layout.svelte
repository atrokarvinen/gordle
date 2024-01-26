<script lang="ts">
	import { AppBar, AppShell } from '@skeletonlabs/skeleton';
	import '../app.pcss';
</script>

<AppShell regionPage="" slotPageContent="p-4 items-center justify-center flex flex-col">
	<svelte:fragment slot="header">
		<AppBar></AppBar>
	</svelte:fragment>
	<slot />
</AppShell>

<style global>
	:global(*) {
		box-sizing: border-box;
	}
</style>
