<script lang="ts">
	import type { GameDto } from '$lib/models';
	import { i18n } from '$lib/translations/i18n';

	export let filteredGames: GameDto[];

	$: totalPlayed = filteredGames.length;
	$: totalWinRate = calculateWinRate(filteredGames);

	const calculateWinRate = (games: GameDto[]) => {
		if (filteredGames.length === 0) return 0;
		return (filteredGames.filter((g) => g.state === 2).length / totalPlayed) * 100;
	};
</script>

<p>{$i18n.t('games_played')}: {totalPlayed}</p>
<p>{$i18n.t('win')}: {totalWinRate.toFixed(1)}%</p>
