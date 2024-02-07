<script lang="ts">
	import { i18n } from '$lib/translations/i18n';
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
	<p>{$i18n.t('failed_to_load_statistics')}</p>
{/await}
