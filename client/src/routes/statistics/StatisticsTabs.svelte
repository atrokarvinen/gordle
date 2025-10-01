<script lang="ts">
	import { resolve } from '$app/paths';
	import { page } from '$app/state';
	import { i18n } from '$lib/translations/i18n';
	import { Navigation } from '@skeletonlabs/skeleton-svelte';
	import { DEFAULT_LIMIT, DEFAULT_PAGE } from './game-history/defaults';
</script>

<Navigation.Bar>
	<Navigation.Tile
		href={resolve('/statistics/game-history') + `?page=${DEFAULT_PAGE}&limit=${DEFAULT_LIMIT}`}
		selected={page.url.pathname.includes('/game-history')}
	>
		{$i18n.t('game_history')}
	</Navigation.Tile>
	<Navigation.Tile
		href={resolve(`/statistics/stats`)}
		selected={page.url.pathname.endsWith('/stats')}>{$i18n.t('statistics')}</Navigation.Tile
	>
</Navigation.Bar>
