<script lang="ts">
	import { page } from '$app/stores';
	import { axios, getApiErrorMessage } from '$lib/axios';
	import Gameover from '$lib/components/Gameover.svelte';
	import Keyboard from '$lib/components/Keyboard.svelte';
	import NewGameButton from '$lib/components/NewGameButton.svelte';
	import WordBoard from '$lib/components/WordBoard.svelte';
	import { LETTERS_COUNT } from '$lib/constants.js';
	import type { Guess, GuessDto, GuessResultDto, GuessedLetter } from '$lib/models';
	import { convertLetterState } from '$lib/utils';
	import { getModalStore, getToastStore } from '@skeletonlabs/skeleton';

	export let data;

	$: gameId = Number($page.params.gameId);
	$: console.log('gameId:', gameId);
	$: console.log('loaded game:', data.game);
	$: guesses = data.game?.guesses ?? [];
	$: gameover = data.game?.gameover ?? undefined;

	const submitGuess = async () => {
		if (word.length !== LETTERS_COUNT) {
			console.log('word "%s" length is not correct', word);
			return;
		}
		const payload: GuessDto = { gameId, word };
		try {
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
			currentIndex = 0;
		} catch (error) {
			toastStore.trigger({
				background: 'variant-filled-error',
				message: getApiErrorMessage(error),
				autohide: true
			});
		}
	};

	const quit = async () => {
		const response = await axios.delete(`/games/${gameId}`);
		console.log(response.data);
	};

	const confirmQuit = () => {
		modalStore.trigger({
			type: 'confirm',
			title: 'Confirm',
			body: 'Are you sure you want to quit?',
			response: (response) => {
				if (response) {
					quit();
				}
			}
		});
	};

	let currentIndex = 0;
	let currentGuess = Array.from(Array(LETTERS_COUNT).keys()).map(() => '');
	$: currentGuessIndex = guesses.length;
	$: word = currentGuess.join('');
	$: console.log('word: "' + word + '"');
	$: console.log('guesses:', guesses);
	$: console.log('currentGuess:', currentGuess);

	const onKeyDown = (e: KeyboardEvent) => {
		const index = currentIndex;
		const key = e.key;
		console.log('keydown:', key);
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
		const alphabets = 'abcdefghijklmnopqrstuvwxyz'.split('');
		if (alphabets.includes(key.toLocaleLowerCase())) {
			currentGuess = currentGuess.map((l, i) => (i === index ? key.toUpperCase() : l));
			currentIndex = Math.min(index + 1, currentGuess.length - 1);
		}
	};

	const modalStore = getModalStore();
	const toastStore = getToastStore();

	const letterClicked = (i: number) => {
		currentIndex = i;
	};

	let isGameStopped = false;
	$: {
		const isGameover = gameover?.isGameover ?? false;
		const isNewGame = gameId === -1;
		isGameStopped = isGameover || isNewGame;
		console.log('isGameStopped:', isGameStopped);
	}
</script>

<svelte:window on:keydown={onKeyDown} />

<div class="flex flex-col items-center gap-y-3">
	{#if gameover}
		<Gameover {gameover} />
	{/if}
	<WordBoard
		bind:currentGuess
		words={guesses}
		{currentGuessIndex}
		currentLetterIndex={currentIndex}
		isGameover={gameover?.isGameover ?? false}
		{letterClicked}
	/>
	<Keyboard {guesses} />

	<div class="flex flex-col gap-y-3">
		<NewGameButton {isGameStopped} />
		<button class="btn variant-filled-secondary" on:click={confirmQuit}>Quit</button>
	</div>
</div>
