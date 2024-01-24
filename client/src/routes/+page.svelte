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

	const echo = async () => {
		const response = await axios.get<string>(`/echo`);
		console.log(response.data);
	};
</script>

<div>
	<NewGameButton />
	<button on:click={loadGame}>Continue</button>
	<button on:click={echo}>Echo</button>
</div>
