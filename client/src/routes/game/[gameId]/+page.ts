import { browser } from '$app/environment';
import { axios } from '$lib/axios';
import type { GameDto } from '$lib/models.js';
import { isAxiosError, type AxiosError } from 'axios';

export const load = async ({ params }) => {
	if (!browser) {
		return { game: null };
	}
	const gameId = params.gameId;
	if (gameId === '-1') {
		return { game: null };
	}
	console.log('[load] gameId', gameId);

	try {
		const gameResponse = await axios.get<GameDto>(`/games/${gameId}`);
		const game = gameResponse.data;
		return { game };
	} catch (error: unknown) {
		if (isAxiosError(error)) {
			const axiosError: AxiosError = error;
			if (axiosError?.response?.status === 403) {
				return { unauthorized: true };
			}
		}
		return { game: null };
	}
};
