import { browser } from '$app/environment';
import { axios } from '$lib/axios';
import type { GameDto, Guess } from '$lib/models.js';

export const load = async ({ params }) => {
	if (!browser) {
		return { game: null };
	}
	const gameId = params.gameId;
	if (gameId === '-1') {
		return { game: null };
	}
	console.log('[load] gameId', gameId);

	const gameResponse = await axios.get<GameDto>(`/games/${gameId}`);
	const game = gameResponse.data;

	const guessResponse = await axios.get<Guess[]>(`/games/${gameId}/guesses`);
	const guesses = guessResponse.data;

	return { game, guesses };
};
