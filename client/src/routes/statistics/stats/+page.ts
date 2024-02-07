import { browser } from '$app/environment';
import { axios } from '$lib/axios';
import type { GameDto } from '$lib/models';

export const load = async () => {
	if (!browser) return { games: [] };

	const queryUrl = `/statistics/stats`;
	const loadPromise = axios.get<GameDto[]>(queryUrl);

	return { loadPromise };
};
