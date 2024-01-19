<script lang="ts">
	import { goto } from '$app/navigation';
	import { axios } from '$lib/axios';
	import type { GameDto } from '$lib/models';

	const createGame = async () => {
		const response = await axios.post<GameDto>(`/games`);
		const game = response.data;
		console.log('created game:', game);
		goto(`/game/${game.id}`);
	};

	const loadGame = async () => {
		const response = await axios.get<GameDto>(`/games/latest`);
		console.log(response.data);
		const game = response.data;
		goto(`/game/${game.id}`);
	};
</script>

<div>
	<button on:click={createGame}>Start new game</button>
	<button on:click={loadGame}>Continue</button>
</div>
