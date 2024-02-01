import { axios } from '$lib/axios';
import type { GameoverDto, Guess, GuessDto, GuessResultDto, GuessedLetter } from '$lib/models';
import { convertLetterState } from '$lib/utils';

export type SubmitGuessResponse = { gameover: GameoverDto; guess: Guess };

export const submitGuess = async (gameId: number, word: string): Promise<SubmitGuessResponse> => {
	const payload: GuessDto = { gameId, word };
	const response = await axios.post(`/games/${gameId}/guesses`, payload);
	console.log(response.data);
	const results: GuessResultDto = response.data;
	const letters: GuessedLetter[] = results.results.map((x, i) => {
		const letter = word[i];
		const state = convertLetterState(x);
		return { letter, state };
	});
	const guess: Guess = { letters, word };
	const gameover = results.gameover;
	return { gameover, guess };
};
