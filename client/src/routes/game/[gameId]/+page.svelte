<script lang="ts">
	import { page } from '$app/stores';
	import Gameover from '$lib/components/Gameover.svelte';
	import Keyboard from '$lib/components/Keyboard.svelte';
	import KeyboardObserver from '$lib/components/KeyboardObserver.svelte';
	import NewGameButton from '$lib/components/NewGameButton.svelte';
	import QuitGameButton from '$lib/components/QuitGameButton.svelte';
	import WordBoard from '$lib/components/WordBoard.svelte';
	import { LETTERS_COUNT } from '$lib/constants.js';
	import type { GameoverDto } from '$lib/models';
	import { getToastStore } from '@skeletonlabs/skeleton';
	import { submitGuess as createGuessRequest } from './api.js';

	export let data;

	const toastStore = getToastStore();
	const emptyGuess = Array.from(Array(LETTERS_COUNT).keys()).map(() => '');

	let currentIndex = 0;
	let currentGuess = emptyGuess;
	let isGameStopped = false;

	$: gameId = Number($page.params.gameId);
	$: guesses = data.game?.guesses ?? [];
	$: gameover = data.game?.gameover ?? undefined;
	$: currentGuessIndex = guesses.length;
	$: word = currentGuess.join('');

	$: console.log('loaded game:', data.game);
	$: console.log('currentGuess:', currentGuess);

	$: {
		const isGameover = gameover?.isGameover ?? false;
		const noGameFound = gameId === -1;
		isGameStopped = isGameover || noGameFound;
		resetGuess();
	}

	const submitGuess = async () => {
		const result = await createGuessRequest(gameId, word);
		console.log('error:', result.errorMessage);
		if (typeof result.errorMessage === 'string') {
			toastStore.trigger({
				background: 'variant-filled-error',
				message: result.errorMessage,
				autohide: true
			});
			return;
		}
		const { gameover: go, guess } = result;
		gameover = go;
		guesses = [...guesses, guess];
	};

	const letterClicked = (i: number) => (currentIndex = i);
	const onGameover = (go: GameoverDto) => (gameover = go);
	const resetGuess = () => {
		currentGuess = emptyGuess;
		currentIndex = 0;
	};
</script>

<KeyboardObserver bind:currentGuess bind:currentIndex {submitGuess} />

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
		<QuitGameButton {gameId} {onGameover} />
	</div>
</div>
