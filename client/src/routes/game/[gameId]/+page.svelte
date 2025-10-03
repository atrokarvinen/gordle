<script lang="ts">
	import { goto } from '$app/navigation';
	import { resolve } from '$app/paths';
	import { getApiErrorMessage } from '$lib/axios.js';
	import Gameover from '$lib/components/Gameover.svelte';
	import Keyboard from '$lib/components/Keyboard.svelte';
	import KeyboardObserver from '$lib/components/KeyboardObserver.svelte';
	import NewGameButton from '$lib/components/NewGameButton.svelte';
	import QuitGameButton from '$lib/components/QuitGameButton.svelte';
	import WordBoard from '$lib/components/WordBoard.svelte';
	import { languageStore } from '$lib/languageStore.js';
	import type { GameoverDto } from '$lib/models';
	import { toaster } from '$lib/toaster.js';
	import { i18n } from '$lib/translations/i18n.js';
	import { submitGuess as requestCreateGuess } from './api.js';

	export let data;

	$: {
		if (data.unauthorized) {
			goto(resolve('/'), { state: { redirected: true } });
		}
	}

	let currentIndex = 0;
	let isGameStopped = false;
	let submitting = false;
	let currentGuess: string[] = [''];

	$: gameId = data.game?.id ?? -1;
	$: guesses = data.game?.guesses ?? [];
	$: gameover = data.game?.gameover ?? undefined;
	$: currentGuessIndex = guesses.length;
	$: word = currentGuess.join('');
	$: wordLength = data.game?.wordLength ?? 6;
	$: maxAttempts = data.game?.maxAttempts ?? 6;
	$: emptyGuess = Array.from(Array(wordLength).keys()).map(() => '');
	$: console.log('loaded game:', data.game);
	$: console.log('currentGuess:', currentGuess);

	$: $languageStore = data.game?.language ?? 'en';

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
			if (word.length !== wordLength) throw 'word_length_wrong';
			const result = await requestCreateGuess(gameId, word);
			gameover = result.gameover;
			guesses = [...guesses, result.guess];
			resetGuess();
		} catch (error) {
			const err = getApiErrorMessage(error);
			toaster.error({
				//
				title: $i18n.t(err.message, { data: err.data })
			});
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

<div class="flex h-full w-full flex-col items-center gap-y-3">
	{#if gameover}
		<Gameover {gameover} />
	{/if}
	<div class="flex flex-grow flex-col items-center justify-end gap-y-3 pb-6 sm:justify-start">
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
				<NewGameButton {wordLength} {maxAttempts} />
			{:else}
				<QuitGameButton {gameId} {onGameover} />
			{/if}
		</div>
	</div>
</div>
