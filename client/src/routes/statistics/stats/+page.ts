import { browser } from '$app/environment';
import { axios } from '$lib/axios';
import type { StatisticsDto } from '$lib/models';

export const load = async ({ url }) => {
	if (!browser) return { games: [] };
	const lang = url.searchParams.get('lang');
	const wordLength = url.searchParams.get('word-len');

	const queryUrl = `/statistics/stats?lang=${lang}&word-len=${wordLength}`;
	const loadPromise = axios.get<StatisticsDto>(queryUrl);

	return { loadPromise };
};
