<script lang="ts">
	import type { StatisticsDto } from '$lib/models';
	import type { Language } from '$lib/translations/language';

	import DataVisualization from './DataVisualization.svelte';
	import Filters from './Filters.svelte';
	import WinDistributionChart from './WinDistributionChart.svelte';

	export let data: StatisticsDto;

	let selectedLanguage: Language | undefined;
	let selectedWordLength: number = -1;

	$: allGames = data.allGames;
	$: filteredGames = allGames.filter((game) => {
		const languageMatches = game.language === selectedLanguage || selectedLanguage === undefined;
		const wordLengthMatches = game.wordLength === selectedWordLength || selectedWordLength === -1;
		return languageMatches && wordLengthMatches;
	});
</script>

<div class="space-y-3">
	<Filters bind:selectedLanguage bind:selectedWordLength />
	<DataVisualization {filteredGames} />
	<WinDistributionChart {filteredGames} />
</div>
