<script lang="ts">
	import { goto } from '$app/navigation';
	import { base } from '$app/paths';
	import { axios } from '$lib/axios';
	import NewGameButton from '$lib/components/NewGameButton.svelte';
	import type { GameDto } from '$lib/models';

	const loadGame = async () => {
		const response = await axios.get<GameDto>(`/games/latest`);
		console.log(response.data);
		const game = response.data;
		goto(`${base}/game/${game.id}`);
	};
</script>

<div>
	<NewGameButton />
	<button class="btn variant-filled-secondary" on:click={loadGame}>Continue</button>
</div>
