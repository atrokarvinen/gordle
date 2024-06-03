<script lang="ts">
	import type { GameDto } from '$lib/models';
	import { i18n } from '$lib/translations/i18n';

	export let filteredGames: GameDto[];

	$: totalPlayed = filteredGames.length;
	$: totalWinRate = calculateWinRate(filteredGames);
	let currentWinStreak = 0;
	let longestWinStreak = 0;
	$: {
		const { currentStreak, maxStreak } = calculateWinStreaks(filteredGames);
		currentWinStreak = currentStreak;
		longestWinStreak = maxStreak;
	}

	const calculateWinRate = (games: GameDto[]) => {
		if (games.length === 0) return 0;
		return (games.filter((g) => g.state === 2).length / totalPlayed) * 100;
	};

	const calculateWinStreaks = (games: GameDto[]) => {
		let streak = 0;
		let currentStreakSet = false;
		let currentStreak = 0;
		let maxStreak = 0;
		let streaks = [];
		for (let i = 0; i < games.length; i++) {
			if (games[i].state === 2) {
				streak++;
			} else {
				if (!currentStreakSet) {
					currentStreak = streak;
					currentStreakSet = true;
				}
				if (streak > maxStreak) {
					maxStreak = streak;
				}
				streaks.push({ streak, first: games[i].gameover.answer });
				streak = 0;
			}
		}
		console.log(streaks);

		return { currentStreak, maxStreak };
	};
</script>

<p>{$i18n.t('games_played')}: {totalPlayed}</p>
<p>{$i18n.t('win')}: {totalWinRate.toFixed(1)}%</p>
<p>{$i18n.t('streak_current')}: {currentWinStreak}</p>
<p>{$i18n.t('streak_longest')}: {longestWinStreak}</p>
