<script lang="ts">
	import { LetterState } from '$lib/models';
	import LetterBox from './LetterBox.svelte';

	export let inputLetters: string[];
	let currentIndex = 0;

	const onKeyDown = (e: KeyboardEvent) => {
		const index = currentIndex;
		const key = e.key;
		console.log('keydown:', key);
		const currentValue = inputLetters[index];
		if (key === 'Backspace' && currentValue === '') {
			const previous = Math.max(index - 1, 0);
			inputLetters = inputLetters.map((l, i) => (i === previous ? '' : l));
			currentIndex = Math.max(index - 1, 0);
		} else if (key === 'Backspace') {
			inputLetters = inputLetters.map((l, i) => (i === index ? '' : l));
		}
		if (key === 'ArrowLeft') {
			currentIndex = Math.max(index - 1, 0);
		}
		if (key === 'ArrowRight') {
			currentIndex = Math.min(index + 1, inputLetters.length - 1);
		}
		if (key === 'ArrowUp') {
			currentIndex = 0;
		}
		if (key === 'ArrowDown') {
			currentIndex = inputLetters.length - 1;
		}
		if (key === 'Enter') {
			console.log('submitting word...');
		}

		const alphabets = 'abcdefghijklmnopqrstuvwxyz'.split('');
		if (alphabets.includes(key.toLocaleLowerCase())) {
			inputLetters = inputLetters.map((l, i) => (i === index ? key.toUpperCase() : l));
			currentIndex = Math.min(index + 1, inputLetters.length - 1);
		}
	};
</script>

<svelte:window on:keydown={onKeyDown} />

<div class="flex gap-x-1">
	{#each inputLetters as letter, i}
		<LetterBox
			focused={currentIndex === i}
			{letter}
			letterState={LetterState.UNKNOWN}
			value={letter}
		/>
	{/each}
</div>
