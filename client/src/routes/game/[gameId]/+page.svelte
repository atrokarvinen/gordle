<script lang="ts">
	import { page } from '$app/stores';
	import { axios } from '$lib/axios';
	import Keyboard from '$lib/components/Keyboard.svelte';
	import WordBoard from '$lib/components/WordBoard.svelte';
	import type { Guess, GuessDto, GuessResultDto, GuessedLetter } from '$lib/models';
	import { convertLetterState } from '$lib/utils';

	export let data;

	const gameId = Number($page.params.gameId);
	console.log('gameId:', gameId);

	$: console.log('loaded game:', data.game);

	let guesses: Guess[] = data.guesses ?? [];
	let isGameover = false;
	let isGameWon = false;
	let answer: string | undefined;

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
		// currentGuess = Array.from(Array(LETTERS_COUNT).keys()).map(() => '');

		if (results.gameover.isGameover) {
			isGameover = true;
			isGameWon = results.gameover.win;
			answer = results.gameover.answer;
		}
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

{#if isGameover}
	<div>
		<p>Game over! You {isGameWon ? 'win' : 'lose'}!</p>
		<p>
			Answer was: '{answer}'
		</p>
	</div>
{/if}
<WordBoard words={guesses} {currentGuessIndex} bind:currentGuess />
<Keyboard {guesses} onKeyDown={(e) => console.log(`pressed '${e}'`)} />

<button on:click={quit}>Quit</button>
<button on:click={submitGuess}>Guess</button>

<style>
	button {
		color: 'white';
		background-color: 'green';
	}
</style>
