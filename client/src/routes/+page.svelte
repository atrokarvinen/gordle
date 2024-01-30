<script lang="ts">
	import { goto } from '$app/navigation';
	import { base } from '$app/paths';
	import { ProgressRadial } from '@skeletonlabs/skeleton';

	export let data;

	$: {
		if (data.loader) {
			data.loader().then((r) => {
				const gameId = r.gameId;
				goto(`${base}/game/${gameId}`);
			});
		}
	}
</script>

<ProgressRadial width="w-16" />
