<script lang="ts">
	import { goto } from '$app/navigation';
	import { base } from '$app/paths';
	import { axios, getApiErrorMessage } from '$lib/axios';
	import type { GameDto } from '$lib/models';
	import { getToastStore } from '@skeletonlabs/skeleton';

	const toastStore = getToastStore();
	const createGame = async () => {
		try {
			const response = await axios.post<GameDto>(`/games`);
			const game = response.data;
			console.log('created game:', game);

			toastStore.trigger({
				background: 'variant-filled-success',
				message: 'New game started',
				autohide: true
			});
			goto(`${base}/game/${game.id}`);
		} catch (error) {
			toastStore.trigger({
				background: 'variant-filled-error',
				message: getApiErrorMessage(error),
				autohide: true
			});
		}
	};
</script>

<button class="btn variant-filled-primary" on:click={createGame}>New game</button>
