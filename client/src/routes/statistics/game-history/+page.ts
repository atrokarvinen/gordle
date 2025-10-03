import { browser } from '$app/environment';
import { axios } from '$lib/axios';
import type { GameHistoryDto } from '$lib/models';

export const load = async ({ url }) => {
	if (!browser) return { games: [] };
	const pageStr = url.searchParams.get('page');
	const page = pageStr ? parseInt(pageStr) - 1 : 0;
	const limit = url.searchParams.get('limit');

	const loadPromise = axios.get<GameHistoryDto>(`/statistics/history?page=${page}&limit=${limit}`);

	return { loadPromise };
};
