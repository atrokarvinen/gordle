<script lang="ts">
	import { page } from '$app/stores';
	import { axios } from '$lib/axios';
	import Keyboard from '$lib/components/Keyboard.svelte';
	import WordBoard from '$lib/components/WordBoard.svelte';
	import WordGuess from '$lib/components/WordGuess.svelte';
	import type { Guess, GuessDto, GuessResultDto, GuessedLetter } from '$lib/models';
	import { convertLetterState } from '$lib/utils';

	export let data;

	const gameId = Number($page.params.gameId);
	console.log('gameId:', gameId);

	$: console.log('loaded game:', data.game);

	let guesses: Guess[] = data.guesses ?? [];

	const submitGuess = async () => {
		const payload: GuessDto = { gameId, word };
		const response = await axios.post(`/games/${gameId}/guesses`, payload);
		console.log(response.data);

		const results: GuessResultDto = response.data;
		const letters: GuessedLetter[] = results.map((x, i) => {
			const letter = word[i];
			const state = convertLetterState(x);
			return { letter, state };
		});
		const g: Guess = { letters, word };
		guesses = [...guesses, g];
	};

	const quit = async () => {
		const response = await axios.delete(`/games/${gameId}`);
		console.log(response.data);
	};

	let currentGuess = 'koalas'.split('');
	// let currentGuess = Array.from(Array(LETTERS_COUNT).keys()).map(() => '');
	$: word = currentGuess.join('');
	$: console.log('word: "' + word + '"');
	$: console.log('guesses:', guesses);
</script>

<WordBoard words={guesses} />
<WordGuess bind:inputLetters={currentGuess} />
<Keyboard {guesses} onKeyDown={(e) => console.log(`pressed '${e}'`)} />

<button on:click={quit}>Quit</button>
<button on:click={submitGuess}>Guess</button>

<style>
	button {
		color: 'white';
		background-color: 'green';
	}
</style>
