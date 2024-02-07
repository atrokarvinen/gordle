import { browser } from '$app/environment';
import { axios } from '$lib/axios';
import type { StatisticsDto } from '$lib/models';

export const load = async ({ url }) => {
	if (!browser) return { games: [] };

	const queryUrl = `/statistics/stats`;
	const loadPromise = axios.get<StatisticsDto>(queryUrl);

	return { loadPromise };
};
