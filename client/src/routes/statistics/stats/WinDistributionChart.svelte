<script lang="ts">
	type EChartsOption = echarts.EChartsOption;

	import type { GameDto } from '$lib/models';

	import { i18n } from '$lib/translations/i18n';
	import { echarts } from './echarts';
	export let filteredGames: GameDto[];

	const maxGuesses = 8;
	const guessArray = Array.from({ length: maxGuesses }, (_, i) => i + 1);
	const yLabels = [...guessArray, $i18n.t('lose')].reverse();
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

	const getLost = () => {
		const totalPlayed = filteredGames.length;
		const wins = filteredGames.filter((g) => g.state === 2).length;
		const totalWinRate = filteredGames.length === 0 ? 0 : (wins / filteredGames.length) * 100;
		if (totalPlayed === 0) return 0;
		return 100 - totalWinRate;
	};

	const getWon = (guess: number) => {
		const gamesWithGuess = filteredGames.filter((game) => {
			const guessCountMatches = game.guesses?.length === guess;
			const isWin = game.state === 2;
			return guessCountMatches && isWin;
		});
		const count = gamesWithGuess.length;
		return (count / filteredGames.length) * 100;
	};

	$: allData = yLabels.map((label, i) => {
		const isLost = label === $i18n.t('lose');
		const guess = +label;
		return isLost ? getLost() : getWon(guess);
	});

	const getColor = (cssVar: string) => {
		const color = getComputedStyle(document.body).getPropertyValue(cssVar);
		return `rgba(${color})`;
	};

	$: {
		console.log('ylabels:', yLabels);
		console.log('winData:', winData);
		console.log('lossData:', lossData);
		console.log('allData:', allData);
	}

	const textColor = '#e4e5ec';
	let option: EChartsOption;
	$: {
		option = {
			title: {
				text: $i18n.t('games_by_guess_count'),
				left: 'center',
				textStyle: { color: textColor }
			},
			textStyle: { color: textColor },
			tooltip: {
				trigger: 'axis',
				axisPointer: { type: 'shadow' },
				renderMode: 'richText',
				// formatter: '{a}, {b}: {c}%',
				valueFormatter: (value) => (value ? `${(+value).toFixed(2)}%` : '0%')
			},
			legend: { textStyle: { color: textColor } },
			xAxis: {
				type: 'value',
				axisLabel: { color: textColor }
			},
			yAxis: [
				{
					type: 'category',
					data: yLabels,
					axisLabel: { color: textColor }
				}
			],
			series: [
				{
					name: $i18n.t('wins'),
					type: 'bar',
					stack: 'total',
					emphasis: { focus: 'series' },
					color: '#11ba81',
					data: winData.slice().reverse().slice(0, 8)
				},
				{
					name: $i18n.t('losses'),
					type: 'bar',
					stack: 'total',
					emphasis: { focus: 'series' },
					color: '#d31b76',
					data: [lossData[8]]
				}
			]
		};
	}
</script>

<div id="chartWrapper" class="h-96" use:echarts={option}>
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
