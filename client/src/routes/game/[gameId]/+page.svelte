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
</script>

<div class="flex flex-col items-center gap-y-3">
	{#if gameover}
		<Gameover {gameover} />
	{/if}
	<WordBoard
		words={guesses}
		{currentGuessIndex}
		bind:currentGuess
		isGameover={gameover?.isGameover ?? false}
	/>
	<Keyboard {guesses} onKeyDown={(e) => console.log(`pressed '${e}'`)} />

	<div class="flex flex-row gap-x-3">
		<NewGameButton />
		<button class="btn variant-filled-secondary" on:click={quit}>Quit</button>
		<button class="btn variant-filled-secondary" on:click={submitGuess}>Guess</button>
	</div>
</div>
