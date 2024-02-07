import { browser } from '$app/environment';
import { axios } from '$lib/axios';
import type { GameHistoryDto } from '$lib/models';

export const load = async ({ url }) => {
	if (!browser) return { games: [] };
	const page = url.searchParams.get('page');
	const limit = url.searchParams.get('limit');

	const loadPromise = axios.get<GameHistoryDto>(`/statistics/history?page=${page}&limit=${limit}`);

	return { loadPromise };
};
