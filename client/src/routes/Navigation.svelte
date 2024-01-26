<script lang="ts">
	import { base } from '$app/paths';
	import { AppBar } from '@skeletonlabs/skeleton';
</script>

<AppBar>
	<svelte:fragment slot="lead">
		<a class="btn-icon variant-filled-surface" href="{base}/"><i class="fa-solid fa-house" /></a>
	</svelte:fragment>
	<svelte:fragment slot="trail">
		<a class="btn-icon variant-filled-surface" href="{base}/info"><i class="fa-solid fa-info" /></a>
	</svelte:fragment>
</AppBar>
