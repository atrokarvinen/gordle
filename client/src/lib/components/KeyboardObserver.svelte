<script lang="ts">
	import { alphabetStore } from '$lib/alphabetStore';
	import { languageStore } from '$lib/languageStore';

	export let currentIndex: number;
	export let currentGuess: string[];
	export let submitGuess: () => void;
	$: lang = $languageStore;

	const onKeyDown = (e: KeyboardEvent) => {
		const index = currentIndex;
		const key = e.key;
		if (e.altKey || e.ctrlKey || e.metaKey || e.shiftKey) {
			return;
		}
		const currentValue = currentGuess[index];
		if (key === 'Backspace' && currentValue === '') {
			const previous = Math.max(index - 1, 0);
			currentGuess = currentGuess.map((l, i) => (i === previous ? '' : l));
			currentIndex = Math.max(index - 1, 0);
		} else if (key === 'Backspace') {
			currentGuess = currentGuess.map((l, i) => (i === index ? '' : l));
		}
		if (key === 'ArrowLeft') {
			currentIndex = Math.max(index - 1, 0);
		}
		if (key === 'ArrowRight') {
			currentIndex = Math.min(index + 1, currentGuess.length - 1);
		}
		if (key === 'ArrowUp') {
			currentIndex = 0;
		}
		if (key === 'ArrowDown') {
			currentIndex = currentGuess.length - 1;
		}
		if (key === 'Enter') {
			submitGuess();
		}
		if ($alphabetStore.alphabets.includes(key.toLocaleLowerCase())) {
			currentGuess = currentGuess.map((l, i) => (i === index ? key.toUpperCase() : l));
			currentIndex = Math.min(index + 1, currentGuess.length - 1);
		}
	};
</script>

<svelte:window on:keydown={onKeyDown} />
