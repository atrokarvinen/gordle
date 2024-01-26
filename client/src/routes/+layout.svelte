<script lang="ts">
	import { base } from '$app/paths';
	import { AppBar, AppShell } from '@skeletonlabs/skeleton';
	import '../app.pcss';
</script>

<AppShell regionPage="" slotPageContent="p-4 items-center justify-center flex flex-col">
	<svelte:fragment slot="header">
		<AppBar>
			<nav class="navi">
				<a href="{base}/">Home</a>
				<a href="{base}/game/2">Game</a>
			</nav>
		</AppBar>
	</svelte:fragment>
	<slot />
</AppShell>

<style global>
	:global(*) {
		box-sizing: border-box;
	}
</style>
