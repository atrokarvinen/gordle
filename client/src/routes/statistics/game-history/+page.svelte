<script lang="ts">
	import { i18n } from '$lib/translations/i18n';
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
		<p>{$i18n.t('no_games_found')}</p>
	{:else}
		<GameHistoryTable games={dto?.data.games ?? []} totalCount={dto?.data.totalCount ?? 0} />
	{/if}
{:catch}
	<p>{$i18n.t('failed_to_load_history')}</p>
{/await}
