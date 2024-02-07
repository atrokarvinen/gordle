<script lang="ts">
	import { ProgressRadial } from '@skeletonlabs/skeleton';
	import StatisticsView from './StatisticsView.svelte';

	export let data;
</script>

{#await data.loadPromise}
	<ProgressRadial width="w-16" />
{:then value}
	{#if value}
		<StatisticsView data={value.data} />
	{/if}
{:catch error}
	<p>Failed to load statistics.</p>
{/await}
