<script lang="ts">
	import { base } from '$app/paths';
	import { AppBar } from '@skeletonlabs/skeleton';
	import { DEFAULT_LIMIT, DEFAULT_PAGE } from './statistics/game-history/defaults';
</script>

<AppBar>
	<svelte:fragment slot="lead">
		<a data-testid="home-link" class="btn-icon variant-filled-surface" href="{base}/"
			><i class="fa-solid fa-house" /></a
		>
	</svelte:fragment>
	<svelte:fragment slot="trail">
		<a
			data-testid="statistics-link"
			class="btn-icon variant-filled-surface"
			href="{base}/statistics/game-history?page={DEFAULT_PAGE}&limit={DEFAULT_LIMIT}"
			><i class="fa-solid fa-history" /></a
		>
		<a data-testid="settings-link" class="btn-icon variant-filled-surface" href="{base}/settings"
			><i class="fa-solid fa-cog" /></a
		>
		<a data-testid="info-link" class="btn-icon variant-filled-surface" href="{base}/info"
			><i class="fa-solid fa-info" /></a
		>
	</svelte:fragment>
</AppBar>
