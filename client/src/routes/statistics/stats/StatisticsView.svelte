<script lang="ts">
	import LanguageSelect from '$lib/components/LanguageSelect.svelte';
	import type { StatisticsDto } from '$lib/models';
	import type { Language } from '$lib/translations/language';
	import { Bar } from 'svelte-chartjs';

	import { BarElement, CategoryScale, Chart, Legend, LinearScale, Title, Tooltip } from 'chart.js';

	Chart.register(Title, Tooltip, Legend, BarElement, CategoryScale, LinearScale);

	export let data: StatisticsDto;

	$: totalWinRate = +data.total.winRate.toFixed(2) * 100;

	let selectedLanguage: Language = 'en';
	let selectedWordLength = 6;
	const languageChanged = (language: Language) => {
		selectedLanguage = language;
	};
	$: languageData = data.byLanguage[selectedLanguage];
	$: wordMap = languageData[selectedWordLength];

	const wordLengths = [5, 6, 7, 8];
	const maxGuesses = 8;
	const guessArray = Array.from({ length: maxGuesses }, (_, i) => i + 1);
	const yLabels = [...guessArray, 'Lose'];
	let winData = Array.from({ length: maxGuesses + 1 }, () => 0);
	let lossData = Array.from({ length: maxGuesses + 1 }, () => 0);
	console.log('guessArray', guessArray);
	$: {
		if (wordMap) {
			winData = guessArray.map((guess) => {
				const count = wordMap.byGuessCount[guess] || 0;
				return (count / wordMap.total.totalCount) * 100;
			});
			lossData = lossData.map((_, i) => {
				if (i < lossData.length - 1) return 0;
				return (1 - wordMap.total.winRate) * 100;
			});
		} else {
			winData = Array.from({ length: maxGuesses + 1 }, () => 0);
			lossData = Array.from({ length: maxGuesses + 1 }, () => 0);
		}
	}
	$: console.log('languageData: ', languageData);
	$: console.log('wordMap: ', wordMap);
	$: console.log('chartData: ', winData);

	const getColor = (cssVar: string) => {
		const color = getComputedStyle(document.body).getPropertyValue(cssVar);
		console.log(`color: ${cssVar} = ${color}`);
		return `rgba(${color})`;
	};
</script>

<div class="space-y-3">
	<p>
		All games winrate: {data.total.winCount} / {data.total.totalCount} = {totalWinRate}%
	</p>

	<div>
		<p>Select by language:</p>
		<LanguageSelect
			languageOptions={['en', 'fi']}
			onChange={languageChanged}
			value={selectedLanguage}
		/>
	</div>

	<label>
		Select word length:
		<select class="select" bind:value={selectedWordLength}>
			{#each wordLengths as wordLength}
				<option value={wordLength}>{wordLength}</option>
			{/each}
		</select>
	</label>

	<div class="h-96">
		<Bar
			data={{
				datasets: [
					{
						// data: [52, 11, 32, 36, 0, 51, 87, 99, 0],
						data: winData,
						backgroundColor: getColor('--color-primary-500')
					},
					{
						// data: [0, 0, 0, 0, 0, 0, 0, 0, 12],
						data: lossData,
						backgroundColor: getColor('--color-error-500')
					}
				],
				yLabels: yLabels
			}}
			options={{
				plugins: {
					legend: { display: false },
					title: { display: true, text: 'Games by guess count', color: 'white' }
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
						// title: { display: true, text: 'Guess', color: 'white' },
						border: { color: 'white' },
						grid: { color: 'white' },
						ticks: { color: 'white' }
					}
				}
			}}
		/>
	</div>
</div>
