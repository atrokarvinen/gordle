import { browser } from '$app/environment';
import { axios } from '$lib/axios';
import type { GameDto } from '$lib/models';

export const load = async () => {
	if (!browser) {
		return { game: null, loading: true };
	}

	console.log('[load] latest game');

	await login();

	try {
		const response = await axios.get<GameDto>(`/games/latest`);
		console.log(response.data);
		const game = response.data;
		return { gameId: game.id, loading: false };
	} catch (error) {
		return { gameId: '-1', loading: false };
	}
};

const login = async () => {
	console.log('logging in...');
	const userId = localStorage.getItem('userId') ?? 0;
	console.log('userId in localStorage:', userId);
	const response = await axios.post('/users/login', { userId: Number(userId) });
	const user = response.data;
	const loginId = user.ID;

	localStorage.setItem('userId', loginId);
	console.log('received loginId:', loginId);
};
