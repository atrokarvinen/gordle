<script lang="ts">
	import { ProgressRadial } from '@skeletonlabs/skeleton';
	import GameHistoryTable from './GameHistoryTable.svelte';

	export let data;
</script>

{#await data.loadPromise}
	<div class="flex justify-center">
		<ProgressRadial width="w-16" />
	</div>
{:then dto}
	{#if dto?.data.games.length === 0}
		No previous games found.
	{:else}
		<GameHistoryTable games={dto?.data.games ?? []} totalCount={dto?.data.totalCount ?? 0} />
	{/if}
{:catch}
	<p>Failed to load game history.</p>
{/await}
