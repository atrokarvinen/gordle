import { browser } from '$app/environment';
import { axios } from '$lib/axios';
import type { GameDto } from '$lib/models';
import { detectLanguage } from '$lib/translations/uiLanguageStore';

export const load = async () => {
	if (!browser) {
		return { gameId: null };
	}

	const loader = async () => {
		await login();
		return await getGame();
	};

	return { loader };
};

const login = async () => {
	console.log('logging in...');
	const userId = localStorage.getItem('userId') ?? 0;
	console.log('userId in localStorage:', userId);
	const response = await axios.post('/users/login', { userId: Number(userId) });
	const user = response.data;
	const loginId = user.ID;
	console.log('received loginId:', loginId);
	if (loginId > 0) {
		localStorage.setItem('userId', loginId);
		axios.defaults.headers.common['Authorization'] = `Bearer ${loginId}`;
		return;
	}
};

const getGame = async () => {
	try {
		const response = await axios.get<GameDto>(`/games/latest`);
		console.log(response.data);
		const game = response.data;
		return { gameId: game.id };
	} catch (error) {
		console.log('No game found:', error);
	}
	try {
		console.log('Creating new game...');
		const payload = { maxAttempts: 6, wordLength: 6, language: detectLanguage() };
		const response = await axios.post<GameDto>(`/games`, payload);
		const game = response.data;
		return { gameId: game.id };
	} catch (error) {
		return { gameId: '-1' };
	}
};
