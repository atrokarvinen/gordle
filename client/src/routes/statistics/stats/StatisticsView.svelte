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
	const languageChanged = (language: Language) => {
		selectedLanguage = language;
	};
	$: languageData = data.byLanguage[selectedLanguage];

	const wordLengths = [5, 6, 7, 8];
	$: {
		const byGuesses = wordLengths.map((x) => {
			// console.log(`languageData[${x}]`, languageData[x]);

			if (!languageData[x]) return;
			const totals = languageData[x].total;
			const wordLenInfo = languageData[x].byGuessCount;
			const guesses = [4, 5, 6, 7, 8];
			const winRates = guesses.map((guess) => {
				const winCount = wordLenInfo[guess] ?? 0;
				const totalCount = totals.totalCount;
				return totalCount === 0 ? 0 : +((winCount / totalCount) * 100).toFixed(2);
			});
			return winRates;
		});
		// console.log(byGuesses);
	}
	// $: console.log('languageData: ', languageData);

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
		<select class="select">
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
						data: [52, 11, 32, 36, 0, 51, 87, 99, 0],
						backgroundColor: getColor('--color-primary-500')
					},
					{
						data: [0, 0, 0, 0, 0, 0, 0, 0, 12],
						backgroundColor: getColor('--color-error-500')
					}
				],
				yLabels: [0, 1, 2, 3, 4, 5, 6, 7, 'Lose']
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
