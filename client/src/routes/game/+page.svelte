<script lang="ts">
	import { axios } from '$lib/axios';
	import Keyboard from '$lib/components/Keyboard.svelte';
	import WordBoard from '$lib/components/WordBoard.svelte';
	import WordGuess from '$lib/components/WordGuess.svelte';
	import type { Guess, GuessDto, GuessResultDto, GuessedLetter } from '$lib/models';
	import { convertLetterState } from '$lib/utils';

	// const data = guesses;
	let data: Guess[] = [];
	const gameId = 1;

	const getGames = async () => {
		const response = await axios.get('/games');
		console.log(response.data);
	};

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
		data = [...data, g];
	};

	const createGame = async () => {
		const response = await axios.post(`/games`);

		console.log(response.data);
	};

	let currentGuess = 'koalas'.split('');
	// let currentGuess = Array.from(Array(LETTERS_COUNT).keys()).map(() => '');
	$: word = currentGuess.join('');
	$: console.log('word: "' + word + '"');
	$: console.log('guesses:', data);
</script>

<WordBoard words={data} />
<WordGuess bind:inputLetters={currentGuess} />
<Keyboard guesses={data} onKeyDown={(e) => console.log(`pressed '${e}'`)} />

<button on:click={getGames}>Get games</button>
<button on:click={submitGuess}>Guess</button>
<button on:click={createGame}>New game</button>

<style>
	button {
		color: 'white';
		background-color: 'green';
	}
</style>
