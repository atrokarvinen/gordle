<script lang="ts">
	import LanguageSelect from '$lib/components/LanguageSelect.svelte';
	import type { GameDto, StatisticsDto } from '$lib/models';
	import type { Language } from '$lib/translations/language';
	import { Bar } from 'svelte-chartjs';

	import { BarElement, CategoryScale, Chart, Legend, LinearScale, Title, Tooltip } from 'chart.js';

	Chart.register(Title, Tooltip, Legend, BarElement, CategoryScale, LinearScale);

	export let data: StatisticsDto;

	$: allGames = data.allGames;
	let totalWinRate = 0;
	let totalPlayed = 0;

	let filteredGames: GameDto[] = [];
	$: {
		filteredGames = allGames.filter((game) => {
			const languageMatches = game.language === selectedLanguage || selectedLanguage === undefined;
			const wordLengthMatches = game.wordLength === selectedWordLength || selectedWordLength === -1;
			return languageMatches && wordLengthMatches;
		});
	}

	let selectedLanguage: Language | undefined;
	let selectedWordLength: number = -1;
	const languageChanged = (language: Language) => {
		if (selectedLanguage === language) {
			selectedLanguage = undefined;
			return;
		}
		selectedLanguage = language;
	};

	const wordLengthOptions = [
		{ value: -1, label: 'All' },
		{ value: 5, label: '5' },
		{ value: 6, label: '6' },
		{ value: 7, label: '7' },
		{ value: 8, label: '8' }
	];
	const maxGuesses = 8;
	const guessArray = Array.from({ length: maxGuesses }, (_, i) => i + 1);
	const yLabels = [...guessArray, 'Lose'];
	let winData = Array.from({ length: maxGuesses + 1 }, () => 0);
	let lossData = Array.from({ length: maxGuesses + 1 }, () => 0);
	console.log('guessArray', guessArray);
	$: {
		console.log('filteredGames: ', filteredGames);
		winData = guessArray.map((guess) => {
			const gamesWithGuess = filteredGames.filter((game) => game.guesses?.length === guess);
			console.log('gamesWithGuess: ', gamesWithGuess);
			const count = gamesWithGuess.length;
			return (count / filteredGames.length) * 100;
		});
		const wins = filteredGames.filter((g) => g.state === 2).length;
		totalPlayed = filteredGames.length;
		totalWinRate = filteredGames.length === 0 ? 0 : (wins / filteredGames.length) * 100;
		lossData = lossData.map((_, i) => {
			if (i < lossData.length - 1) return 0;
			if (totalPlayed === 0) return 0;
			return 100 - totalWinRate;
		});
	}
	$: console.log('chartData: ', winData);

	const getColor = (cssVar: string) => {
		const color = getComputedStyle(document.body).getPropertyValue(cssVar);
		return `rgba(${color})`;
	};
</script>

<div class="space-y-3">
	<div>
		<p>Filter by language:</p>
		<div class="flex items-center">
			<LanguageSelect
				languageOptions={['en', 'fi']}
				onChange={languageChanged}
				value={selectedLanguage}
			/>
			{#if selectedLanguage}
				<div class="ml-5">
					<button
						class="btn variant-filled-secondary"
						on:click={() => (selectedLanguage = undefined)}
					>
						<span>
							<i class="fas fa-times" />
						</span>
						<span>Clear</span>
					</button>
				</div>
			{/if}
		</div>
	</div>

	<label>
		Filter by word length:
		<select class="select" bind:value={selectedWordLength}>
			{#each wordLengthOptions as wordLengthOption}
				<option value={wordLengthOption.value}>{wordLengthOption.label}</option>
			{/each}
		</select>
	</label>

	<p>Games played: {totalPlayed}</p>
	<p>Win: {totalWinRate.toFixed(1)}%</p>

	<div class="h-96">
		<Bar
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
						border: { color: 'white' },
						grid: { color: 'white' },
						ticks: { color: 'white' }
					}
				}
			}}
		/>
	</div>
</div>
