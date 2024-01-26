<script lang="ts">
	import { goto } from '$app/navigation';
	import { base } from '$app/paths';
	import { ProgressRadial } from '@skeletonlabs/skeleton';

	export let data;

	$: {
		if (data.loading === false) {
			goto(`${base}/game/${data.gameId}`);
		}
	}
</script>

{#if data.loading}
	<ProgressRadial />
{/if}
