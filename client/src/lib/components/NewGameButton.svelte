<script lang="ts">
	import { goto } from '$app/navigation';
	import { base } from '$app/paths';
	import { axios } from '$lib/axios';
	import type { GameDto } from '$lib/models';

	const createGame = async () => {
		const response = await axios.post<GameDto>(`/games`);
		const game = response.data;
		console.log('created game:', game);
		goto(`${base}/game/${game.id}`);
	};
</script>

<button on:click={createGame}>New game</button>
