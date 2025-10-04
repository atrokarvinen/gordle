<script lang="ts">
	import { alphabetStore } from '$lib/alphabetStore';
	import { LetterState, type Guess } from '$lib/models';
	import KeyboardButton from './KeyboardButton.svelte';

	export let guesses: Guess[];
	export let submitting: boolean;
	$: alphabets = $alphabetStore.alphabets;
	$: qwertyLine1 = $alphabetStore.qwerty1;
	$: qwertyLine2 = $alphabetStore.qwerty2;
	$: qwertyLine3 = $alphabetStore.qwerty3;
	$: qwertyLineExtra = $alphabetStore.qwertyExtra;

	let guessLetterMap: Record<string, LetterState> = {};
	$: {
		const newMap: Record<string, LetterState> = {};
		alphabets.forEach((letter) => {
			newMap[letter] = LetterState.UNKNOWN;
		});
		guesses.forEach((guess) => {
			guess.letters.forEach((letter) => {
				const char = letter.letter.toLowerCase();
				const currentValue = newMap[char];
				const newValue = letter.state;
				if (currentValue === LetterState.CORRECT) return;
				if (currentValue === LetterState.CONTAINED && newValue !== LetterState.CORRECT) return;
				newMap[char] = letter.state;
			});
		});
		guessLetterMap = newMap;
	}
</script>

<div class="mt-4 flex flex-col items-center gap-y-2">
	{#if qwertyLineExtra}
		<div class="ml-0 flex gap-x-1">
			{#each qwertyLineExtra as letter (letter)}
				<KeyboardButton {letter} state={guessLetterMap[letter]} />
			{/each}
		</div>
	{/if}
	<div class="ml-0 flex gap-x-1">
		{#each qwertyLine1 as letter (letter)}
			<KeyboardButton {letter} state={guessLetterMap[letter]} />
		{/each}
	</div>
	<div class="flex gap-x-1">
		{#each qwertyLine2 as letter (letter)}
			<KeyboardButton {letter} state={guessLetterMap[letter]} />
		{/each}
	</div>
	<div class="ml-0 flex gap-x-1">
		<div class="mr-1 h-7 md:h-12">
			<KeyboardButton
				letter="Enter"
				state={LetterState.UNKNOWN}
				icon="fa-solid fa-right-to-bracket"
				{submitting}
			/>
		</div>
		{#each qwertyLine3 as letter (letter)}
			<KeyboardButton {letter} state={guessLetterMap[letter]} />
		{/each}
		<div class="ml-1">
			<KeyboardButton
				letter="Backspace"
				state={LetterState.UNKNOWN}
				icon="fa-solid fa-delete-left"
			/>
		</div>
	</div>
</div>
