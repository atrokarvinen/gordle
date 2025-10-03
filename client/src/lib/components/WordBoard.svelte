<script lang="ts">
	import { LetterState, type Guess } from '$lib/models';
	import WordRow from './WordRow.svelte';
	import WordRowInput from './WordRowInput.svelte';

	export let currentGuess: string[];
	export let words: Guess[];
	export let currentGuessIndex: number;
	export let currentLetterIndex: number;
	export let isGameStopped: boolean;
	export let letterClicked: (index: number) => void;
	export let maxAttempts: number;
	export let wordLength: number;

	const emptyLetter = { letter: '', state: LetterState.UNKNOWN };
	$: emptyLetters = Array.from(Array(wordLength).keys()).map(() => emptyLetter);
	$: emptyWords = Array.from(Array(maxAttempts - words.length).keys()).map(() => {
		return { word: '', letters: emptyLetters };
	});
	$: allWords = [...words, ...emptyWords];
</script>

<div class="flex flex-col gap-y-1">
	{#each allWords as word, index (index)}
		<div data-testid="guess-row" class="flex gap-x-1">
			{#if index === currentGuessIndex && !isGameStopped}
				<WordRowInput {currentGuess} currentIndex={currentLetterIndex} {letterClicked} />
			{:else}
				<WordRow {word} />
			{/if}
		</div>
	{/each}
</div>
