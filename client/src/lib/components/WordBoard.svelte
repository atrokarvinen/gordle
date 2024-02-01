<script lang="ts">
	import { LetterState, type Guess } from '$lib/models';
	import LetterBox from './LetterBox.svelte';
	import WordGuess from './WordGuess.svelte';

	export let currentGuess: string[];
	export let words: Guess[];
	export let currentGuessIndex: number;
	export let currentLetterIndex: number;
	export let isGameStopped: boolean;
	export let letterClicked: (index: number) => void;
	export let maxAttempts: number;
	export let wordLength: number;

	const emptyLetter = { letter: '', state: LetterState.UNKNOWN };
	$: mappedData = Array.from(Array(maxAttempts).keys()).map((i) => {
		if (words.length > i) return words[i];
		else return { word: '', letters: Array.from(Array(wordLength).keys()).map(() => emptyLetter) };
	});
</script>

<div class="flex flex-col gap-y-1">
	{#each mappedData as word, index}
		<div data-testid="guess-row" class="flex gap-x-1">
			{#if index === currentGuessIndex && !isGameStopped}
				<WordGuess inputLetters={currentGuess} currentIndex={currentLetterIndex} {letterClicked} />
			{:else}
				{#each word.letters as letter}
					<LetterBox letter={letter.letter} letterState={letter.state} />
				{/each}
				<span data-testid="guessed-word" hidden>{word.word}</span>
			{/if}
		</div>
	{/each}
</div>
