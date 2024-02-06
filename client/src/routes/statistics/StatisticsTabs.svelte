<script lang="ts">
	import { base } from '$app/paths';
	import { page } from '$app/stores';
	import { i18n } from '$lib/translations/i18n';
	import { TabAnchor, TabGroup } from '@skeletonlabs/skeleton';
	import { DEFAULT_LIMIT, DEFAULT_PAGE } from './game-history/defaults';
</script>

<TabGroup>
	<TabAnchor
		href="{base}/statistics/game-history?page={DEFAULT_PAGE}&limit={DEFAULT_LIMIT}"
		selected={$page.url.pathname.includes('/game-history')}>{$i18n.t('game_history')}</TabAnchor
	>
	<TabAnchor href="{base}/statistics/stats" selected={$page.url.pathname.endsWith('/stats')}
		>{$i18n.t('statistics')}</TabAnchor
	>
</TabGroup>
