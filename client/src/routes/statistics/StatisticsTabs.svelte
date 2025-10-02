<script lang="ts">
	import { resolve } from '$app/paths';
	import { page } from '$app/state';
	import { i18n } from '$lib/translations/i18n';
	import { Tabs } from '@skeletonlabs/skeleton-svelte';
	import { DEFAULT_LIMIT, DEFAULT_PAGE } from './game-history/defaults';

	let value: string | undefined = $state();
	$effect(() => {
		const route = page.url.pathname;
		if (route.endsWith('/game-history')) {
			value = 'game-history';
		} else if (route.endsWith('/stats')) {
			value = 'stats';
		}
	});
</script>

<Tabs {value} fluid>
	{#snippet list()}
		<Tabs.Control value="game-history">
			<a
				href={resolve('/statistics/game-history') + `?page=${DEFAULT_PAGE}&limit=${DEFAULT_LIMIT}`}
			>
				{$i18n.t('game_history')}
			</a>
		</Tabs.Control>
		<Tabs.Control value="stats">
			<a href={resolve('/statistics/stats')} class="block h-full w-full">
				{$i18n.t('statistics')}
			</a>
		</Tabs.Control>
	{/snippet}
</Tabs>
