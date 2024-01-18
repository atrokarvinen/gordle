<script lang="ts">
	import { LETTERS_COUNT, MAX_GUESSES } from '$lib/constants';
	import { LetterState, type Guess } from '$lib/models';
	import LetterBox from './LetterBox.svelte';

	export let words: Guess[];

	const emptyLetter = { letter: '', state: LetterState.INCORRECT };
	$: mappedData = Array.from(Array(MAX_GUESSES).keys()).map((i) => {
		if (words.length > i) return words[i];
		else
			return { word: '', letters: Array.from(Array(LETTERS_COUNT).keys()).map(() => emptyLetter) };
	});
</script>

<div class="words">
	{#each mappedData as word}
		<div class="word">
			{#each word.letters as letter}
				<LetterBox letter={letter.letter} letterState={letter.state} />
			{/each}
		</div>
	{/each}
</div>

<style>
	.words {
		display: flex;
		flex-direction: column;
	}
	.word {
		display: flex;
		flex-direction: row;
	}
</style>
