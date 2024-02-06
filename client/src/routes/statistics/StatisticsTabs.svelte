<script lang="ts">
	import { base } from '$app/paths';
	import { page } from '$app/stores';
	import { TabAnchor, TabGroup } from '@skeletonlabs/skeleton';
	import { DEFAULT_LIMIT, DEFAULT_PAGE } from './game-history/defaults';
</script>

<TabGroup>
	<TabAnchor
		href="{base}/statistics/game-history?page={DEFAULT_PAGE}&limit={DEFAULT_LIMIT}"
		selected={$page.url.pathname.includes('/game-history')}>Game history</TabAnchor
	>
	<TabAnchor href="{base}/statistics/stats" selected={$page.url.pathname.endsWith('/stats')}
		>Statistics</TabAnchor
	>
</TabGroup>
