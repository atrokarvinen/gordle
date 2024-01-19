<script lang="ts">
	import { page } from '$app/stores';
	import { axios } from '$lib/axios';
	import Gameover from '$lib/components/Gameover.svelte';
	import Keyboard from '$lib/components/Keyboard.svelte';
	import NewGameButton from '$lib/components/NewGameButton.svelte';
	import WordBoard from '$lib/components/WordBoard.svelte';
	import type { GameoverDto, Guess, GuessDto, GuessResultDto, GuessedLetter } from '$lib/models';
	import { convertLetterState } from '$lib/utils';

	export let data;

	$: gameId = Number($page.params.gameId);
	$: console.log('gameId:', gameId);
	$: console.log('loaded game:', data.game);
	$: guesses = data.guesses ?? [];
	let gameover: GameoverDto | undefined;

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
		// currentGuess = Array.from(Array(LETTERS_COUNT).keys()).map(() => '');
	};

	const quit = async () => {
		const response = await axios.delete(`/games/${gameId}`);
		console.log(response.data);
	};

	let currentGuess = 'koalas'.split('');
	$: currentGuessIndex = guesses.length;
	$: word = currentGuess.join('');
	$: console.log('word: "' + word + '"');
	$: console.log('guesses:', guesses);
	$: console.log('currentGuess:', currentGuess);
</script>

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

<NewGameButton />
<button on:click={quit}>Quit</button>
<button on:click={submitGuess}>Guess</button>

<style>
	button {
		color: 'white';
		background-color: 'green';
	}
</style>
