<script lang="ts">
	import type { GameDto } from '$lib/models';
	// import { Bar } from 'svelte-chartjs';

	import { i18n } from '$lib/translations/i18n';
	// import { BarElement, CategoryScale, Chart, Legend, LinearScale, Title, Tooltip } from 'chart.js';

	// Chart.register(Title, Tooltip, Legend, BarElement, CategoryScale, LinearScale);

	export let filteredGames: GameDto[];

	const maxGuesses = 8;
	const guessArray = Array.from({ length: maxGuesses }, (_, i) => i + 1);
	const yLabels = [...guessArray, $i18n.t('lose')];
	let winData = Array.from({ length: maxGuesses + 1 }, () => 0);
	let lossData = Array.from({ length: maxGuesses + 1 }, () => 0);
	$: winData = guessArray.map((guess) => {
		const gamesWithGuess = filteredGames.filter((game) => {
			const guessCountMatches = game.guesses?.length === guess;
			const isWin = game.state === 2;
			return guessCountMatches && isWin;
		});
		const count = gamesWithGuess.length;
		return (count / filteredGames.length) * 100;
	});
	$: lossData = lossData.map((_, i) => {
		if (i < lossData.length - 1) return 0;
		const totalPlayed = filteredGames.length;
		const wins = filteredGames.filter((g) => g.state === 2).length;
		const totalWinRate = filteredGames.length === 0 ? 0 : (wins / filteredGames.length) * 100;
		if (totalPlayed === 0) return 0;
		return 100 - totalWinRate;
	});

	const getColor = (cssVar: string) => {
		const color = getComputedStyle(document.body).getPropertyValue(cssVar);
		return `rgba(${color})`;
	};
</script>

<div class="h-96">
	<!-- <Bar
		data={{
			datasets: [
				{ data: winData, backgroundColor: getColor('--color-primary-500') },
				{ data: lossData, backgroundColor: getColor('--color-error-500') }
			],
			yLabels: yLabels
		}}
		options={{
			plugins: {
				legend: { display: false },
				title: { display: true, text: $i18n.t('games_by_guess_count'), color: 'white' }
			},
			maintainAspectRatio: false,
			responsive: true,
			indexAxis: 'y',
			backgroundColor: 'yellow',
			color: 'white',
			scales: {
				x: {
					stacked: true,
					title: { display: true, text: '%', color: 'white' },
					border: { color: 'white' },
					grid: { color: 'white' },
					ticks: { color: 'white' }
				},
				y: {
					stacked: true,
					border: { color: 'white' },
					grid: { color: 'white' },
					ticks: { color: 'white' }
				}
			}
		}}
	/> -->
</div>
