<script lang="ts">
	import { page } from '$app/stores';
	import { axios } from '$lib/axios';
	import Gameover from '$lib/components/Gameover.svelte';
	import Keyboard from '$lib/components/Keyboard.svelte';
	import NewGameButton from '$lib/components/NewGameButton.svelte';
	import WordBoard from '$lib/components/WordBoard.svelte';
	import { LETTERS_COUNT } from '$lib/constants.js';
	import type { Guess, GuessDto, GuessResultDto, GuessedLetter } from '$lib/models';
	import { convertLetterState } from '$lib/utils';

	export let data;

	$: gameId = Number($page.params.gameId);
	$: console.log('gameId:', gameId);
	$: console.log('loaded game:', data.game);
	$: guesses = data.guesses ?? [];
	$: gameover = data.game?.gameover ?? undefined;

	const submitGuess = async () => {
		const payload: GuessDto = { gameId, word };
		const response = await axios.post(`/games/${gameId}/guesses`, payload);
		console.log(response.data);

		const results: GuessResultDto = response.data;
		const letters: GuessedLetter[] = results.results.map((x, i) => {
			const letter = word[i];
			const state = convertLetterState(x);
			return { letter, state };
		});
		const g: Guess = { letters, word };
		guesses = [...guesses, g];
		gameover = results.gameover;
		currentGuess = Array.from(Array(LETTERS_COUNT).keys()).map(() => '');
	};

	const quit = async () => {
		const response = await axios.delete(`/games/${gameId}`);
		console.log(response.data);
	};

	let currentGuess = Array.from(Array(LETTERS_COUNT).keys()).map(() => '');
	$: currentGuessIndex = guesses.length;
	$: word = currentGuess.join('');
	$: console.log('word: "' + word + '"');
	$: console.log('guesses:', guesses);
	$: console.log('currentGuess:', currentGuess);

	let currentIndex = 0;

	const onKeyDown = (e: KeyboardEvent) => {
		const index = currentIndex;
		const key = e.key;
		console.log('keydown:', key);
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
			console.log('submitting word...');
		}

		const alphabets = 'abcdefghijklmnopqrstuvwxyz'.split('');
		if (alphabets.includes(key.toLocaleLowerCase())) {
			currentGuess = currentGuess.map((l, i) => (i === index ? key.toUpperCase() : l));
			currentIndex = Math.min(index + 1, currentGuess.length - 1);
		}
	};
</script>

<svelte:window on:keydown={onKeyDown} />

<div class="flex flex-col items-center gap-y-3">
	{#if gameover}
		<Gameover {gameover} />
	{/if}
	<WordBoard
		words={guesses}
		{currentGuessIndex}
		currentLetterIndex={currentIndex}
		bind:currentGuess
		isGameover={gameover?.isGameover ?? false}
	/>
	<Keyboard {guesses} />

	<div class="flex flex-row gap-x-3">
		<NewGameButton />
		<button class="btn variant-filled-secondary" on:click={quit}>Quit</button>
		<button class="btn variant-filled-secondary" on:click={submitGuess}>Guess</button>
	</div>
</div>
