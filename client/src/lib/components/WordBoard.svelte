<script lang="ts">
	import { LETTERS_COUNT, MAX_GUESSES } from '$lib/constants';
	import { LetterState, type Guess } from '$lib/models';
	import LetterBox from './LetterBox.svelte';
	import WordGuess from './WordGuess.svelte';

	export let currentGuess: string[];
	export let words: Guess[];
	export let currentGuessIndex: number;
	export let isGameover: boolean;

	const emptyLetter = { letter: '', state: LetterState.UNKNOWN };
	$: mappedData = Array.from(Array(MAX_GUESSES).keys()).map((i) => {
		if (words.length > i) return words[i];
		else
			return { word: '', letters: Array.from(Array(LETTERS_COUNT).keys()).map(() => emptyLetter) };
	});
	$: console.log('[WordBoard] currentGuess:', currentGuess);
</script>

<div class="words">
	{#each mappedData as word, index}
		{#if index === currentGuessIndex && !isGameover}
			<WordGuess bind:inputLetters={currentGuess} />
		{:else}
			<div class="word">
				{#each word.letters as letter}
					<LetterBox letter={letter.letter} letterState={letter.state} />
				{/each}
			</div>
		{/if}
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
