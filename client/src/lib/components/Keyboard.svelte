<script lang="ts">
	import { LetterState, type Guess } from '$lib/models';
	import KeyboardButton from './KeyboardButton.svelte';

	export let guesses: Guess[];
	export let submitting: boolean;
	export let lang: string;

	$: alphabets = ('abcdefghijklmnopqrstuvwxyz' + (lang === 'fi' ? 'åäö' : '')).split('');
	$: qwertyLine1 = ('qwertyuiop' + (lang === 'fi' ? 'å' : '')).split('');
	$: qwertyLine2 = ('asdfghjkl' + (lang === 'fi' ? 'äö' : '')).split('');
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

<div class="flex flex-col items-center gap-y-1">
	<div class="ml-0 flex gap-x-1">
		{#each qwertyLine1 as letter}
			<KeyboardButton {letter} state={guessLetterMap[letter]} />
		{/each}
	</div>
	<div class={`flex gap-x-1 ${lang === 'en' ? 'ml-4' : ''}`}>
		{#each qwertyLine2 as letter}
			<KeyboardButton {letter} state={guessLetterMap[letter]} />
		{/each}
	</div>
	<div class="ml-0 flex gap-x-1">
		<KeyboardButton
			letter="Enter"
			state={LetterState.UNKNOWN}
			icon="fa-solid fa-right-to-bracket"
			{submitting}
		/>
		{#each qwertyLine3 as letter}
			<KeyboardButton {letter} state={guessLetterMap[letter]} />
		{/each}
		<KeyboardButton letter="Backspace" state={LetterState.UNKNOWN} icon="fa-solid fa-delete-left" />
	</div>
</div>
