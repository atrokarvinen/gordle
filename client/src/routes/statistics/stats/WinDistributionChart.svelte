<script lang="ts">
	type EChartsOption = echarts.EChartsOption;

	import type { GameDto } from '$lib/models';

	import { i18n } from '$lib/translations/i18n';
	import { echarts } from './echarts';
	export let filteredGames: GameDto[];

	const maxGuesses = 8;
	const guessArray = Array.from({ length: maxGuesses }, (_, i) => i + 1);
	const yLabels = [...guessArray, $i18n.t('lose')].reverse();
	let winData: number[][];
	let lossData: number;
	$: winData = guessArray.map((guess) => {
		const gamesWithGuess = filteredGames.filter((game) => {
			const guessCountMatches = game.guesses?.length === guess;
			const isWin = game.state === 2;
			return guessCountMatches && isWin;
		});
		const count = gamesWithGuess.length;
		return [guess, (count / filteredGames.length) * 100];
	});
	$: {
		const totalPlayed = filteredGames.length;
		const wins = filteredGames.filter((g) => g.state === 2).length;
		const totalWinRate = filteredGames.length === 0 ? 0 : (wins / filteredGames.length) * 100;
		if (totalPlayed === 0) {
			lossData = 0;
		} else {
			lossData = 100 - totalWinRate;
		}
	}

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
				valueFormatter: (value) => (value ? `${(+value).toFixed(1)}%` : '0%')
			},
			xAxis: { type: 'category', axisLabel: { color: textColor } },
			yAxis: [{ type: 'value', axisLabel: { color: textColor } }],
			series: [
				{
					name: $i18n.t('wins'),
					type: 'bar',
					stack: 'total',
					emphasis: { focus: 'series' },
					color: '#11ba81',
					data: winData
				},
				{
					name: $i18n.t('losses'),
					type: 'bar',
					stack: 'total',
					emphasis: { focus: 'series' },
					color: '#d31b76',
					data: [[$i18n.t('lose'), lossData]]
				}
			]
		};
	}
</script>

<div id="chartWrapper" class="h-96" use:echarts={option}></div>
