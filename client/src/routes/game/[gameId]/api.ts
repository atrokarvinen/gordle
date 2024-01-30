import { axios, getApiErrorMessage } from '$lib/axios';
import type { GameoverDto, Guess, GuessDto, GuessResultDto, GuessedLetter } from '$lib/models';
import { convertLetterState } from '$lib/utils';

export type SubmitGuessResponse =
	| {
			errorMessage: string;
			gameover?: undefined;
			guess?: undefined;
	  }
	| { errorMessage?: undefined; gameover: GameoverDto; guess: Guess };

export const submitGuess = async (
	gameId: number,
	word: string,
	wordLength: number
): Promise<SubmitGuessResponse> => {
	if (word.length !== wordLength) {
		return { errorMessage: 'Word length is not correct' };
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
		const guess: Guess = { letters, word };
		const gameover = results.gameover;
		return { gameover, guess };
	} catch (error) {
		const errorMessage = getApiErrorMessage(error);
		return { errorMessage };
	}
};
