import { browser } from '$app/environment';
import { axios } from '$lib/axios';
import type { StatisticsDto } from '$lib/models';

export const load = async ({ url }) => {
	if (!browser) return { games: [], dataLoading: true };
	const page = url.searchParams.get('page');
	const limit = url.searchParams.get('limit');

	const loadPromise = axios.get<StatisticsDto>(`/statistics?page=${page}&limit=${limit}`);

	return { loadPromise, dataLoading: true };
};
