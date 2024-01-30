<script lang="ts">
	import { page } from '$app/stores';
	import Gameover from '$lib/components/Gameover.svelte';
	import Keyboard from '$lib/components/Keyboard.svelte';
	import KeyboardObserver from '$lib/components/KeyboardObserver.svelte';
	import NewGameButton from '$lib/components/NewGameButton.svelte';
	import QuitGameButton from '$lib/components/QuitGameButton.svelte';
	import WordBoard from '$lib/components/WordBoard.svelte';
	import type { GameoverDto } from '$lib/models';
	import { getToastStore } from '@skeletonlabs/skeleton';
	import { submitGuess as requestCreateGuess } from './api.js';

	export let data;

	const toastStore = getToastStore();

	let currentIndex = 0;
	let isGameStopped = false;
	let submitting = false;
	let currentGuess: string[] = [''];

	$: gameId = Number($page.params.gameId);
	$: guesses = data.game?.guesses ?? [];
	$: gameover = data.game?.gameover ?? undefined;
	$: currentGuessIndex = guesses.length;
	$: word = currentGuess.join('');
	$: wordLength = data.game?.wordLength ?? 0;
	$: maxAttempts = data.game?.maxAttempts ?? 0;
	$: emptyGuess = Array.from(Array(wordLength).keys()).map(() => '');
	$: console.log('loaded game:', data.game);
	$: console.log('currentGuess:', currentGuess);

	$: {
		const isGameover = gameover?.isGameover ?? false;
		const noGameFound = gameId === -1;
		isGameStopped = isGameover || noGameFound;
		resetGuess();
	}

	const submitGuess = async () => {
		if (submitting) return;
		if (isGameStopped) return;
		submitting = true;
		try {
			const result = await requestCreateGuess(gameId, word, wordLength);
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
		} finally {
			submitting = false;
		}
	};

	const letterClicked = (i: number) => (currentIndex = i);
	const onGameover = (go: GameoverDto) => (gameover = go);
	const resetGuess = () => {
		currentGuess = emptyGuess;
		currentIndex = 0;
	};
</script>

<KeyboardObserver bind:currentGuess bind:currentIndex {submitGuess} />

<div class="flex w-full flex-col items-center gap-y-3">
	{#if gameover}
		<Gameover {gameover} />
	{/if}
	<WordBoard
		{currentGuess}
		words={guesses}
		currentLetterIndex={currentIndex}
		{maxAttempts}
		{wordLength}
		{currentGuessIndex}
		{isGameStopped}
		{letterClicked}
	/>
	<Keyboard {guesses} {submitting} />

	<div class="flex flex-col gap-y-3">
		{#if isGameStopped}
			<NewGameButton />
		{:else}
			<QuitGameButton {gameId} {onGameover} />
		{/if}
	</div>
</div>
