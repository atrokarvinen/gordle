<script lang="ts">
	import { goto } from '$app/navigation';
	import { base } from '$app/paths';
	import { axios } from '$lib/axios';
	import type { GameDto } from '$lib/models';
	import { getModalStore } from '@skeletonlabs/skeleton';

	export let isGameover: boolean;

	const createGame = async () => {
		const response = await axios.post<GameDto>(`/games`);
		const game = response.data;
		console.log('created game:', game);
		goto(`${base}/game/${game.id}`);
	};

	const modalStore = getModalStore();
	const confirmNew = () => {
		if (isGameover) {
			createGame();
			return;
		}
		modalStore.trigger({
			type: 'confirm',
			title: 'Confirm',
			body: 'Are you sure you want to start a new game?',
			response: (response) => {
				if (response) {
					createGame();
				}
			}
		});
	};
</script>

<button class="btn variant-filled-primary" on:click={confirmNew}>New game</button>
