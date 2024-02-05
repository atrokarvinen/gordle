import { browser } from '$app/environment';
import { axios } from '$lib/axios';
import type { GameDto } from '$lib/models';

export const load = async () => {
	if (!browser) return { games: [] };

	const response = await axios.get<GameDto[]>('/statistics');
	const games = response.data;

	return { games };
};
