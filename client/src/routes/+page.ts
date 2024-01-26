import { browser } from '$app/environment';
import { axios } from '$lib/axios';
import type { GameDto } from '$lib/models';

export const load = async () => {
	if (!browser) {
		return { game: null, loading: true };
	}
	console.log('[load] latest game');
	try {
		const response = await axios.get<GameDto>(`/games/latest`);
		console.log(response.data);
		const game = response.data;
		return { gameId: game.id, loading: false };
	} catch (error) {
		return { gameId: '-1', loading: false };
	}
};
