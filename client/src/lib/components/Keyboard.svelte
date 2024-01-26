<script lang="ts">
	import { LetterState, type Guess } from '$lib/models';
	import KeyboardButton from './KeyboardButton.svelte';

	export let guesses: Guess[];
	export let onKeyDown: (key: string) => void;

	const alphabets = 'abcdefghijklmnopqrstuvwxyz'.split('');
	const qwertyLine1 = 'qwertyuiop'.split('');
	const qwertyLine2 = 'asdfghjkl'.split('');
	const qwertyLine3 = 'zxcvbnm'.split('');

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

<div class="flex flex-col gap-y-1">
	<div class="ml-0 flex gap-x-1">
		{#each qwertyLine1 as letter}
			<KeyboardButton {letter} state={guessLetterMap[letter]} {onKeyDown} />
		{/each}
	</div>
	<div class="ml-8 flex gap-x-1">
		{#each qwertyLine2 as letter}
			<KeyboardButton {letter} state={guessLetterMap[letter]} {onKeyDown} />
		{/each}
	</div>
	<div class="ml-16 flex gap-x-1">
		{#each qwertyLine3 as letter}
			<KeyboardButton {letter} state={guessLetterMap[letter]} {onKeyDown} />
		{/each}
	</div>
</div>
