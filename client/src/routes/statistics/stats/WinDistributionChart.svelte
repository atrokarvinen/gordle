<script lang="ts">
	import type { GameDto } from '$lib/models';
	import { i18n } from '$lib/translations/i18n';
	import type { EChartsOption } from 'echarts';
	import { echarts } from './echarts';

	interface Props {
		filteredGames: GameDto[];
	}

	const { filteredGames }: Props = $props();

	const maxGuesses = 8;

	const getWinData = (filteredGames: GameDto[]) => {
		const guessArray = Array.from({ length: maxGuesses }, (_, i) => i + 1);
		const winsByGuess = guessArray.map((guess) => {
			const gamesWithGuess = filteredGames.filter((game) => {
				const guessCountMatches = game.guesses?.length === guess;
				const isWin = game.state === 2;
				return guessCountMatches && isWin;
			});
			const count = gamesWithGuess.length;
			return [guess, (count / filteredGames.length) * 100];
		});
		return winsByGuess;
	};

	const getLossData = (filteredGames: GameDto[]) => {
		const losses = filteredGames.filter((game) => game.state === 3).length;
		const totalPlayed = filteredGames.length;
		if (totalPlayed === 0) return 0;
		return (losses / totalPlayed) * 100;
	};

	const getChartOptions = (winData: number[][], lossData: number) => {
		const textColor = '#e4e5ec';
		const option: EChartsOption = {
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
		return option;
	};

	let winData = $derived(getWinData(filteredGames));
	let lossData = $derived(getLossData(filteredGames));
	let option = $derived(getChartOptions(winData, lossData));
</script>

<div id="chartWrapper" class="h-96" use:echarts={option}></div>
