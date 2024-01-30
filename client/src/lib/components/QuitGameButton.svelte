<script lang="ts">
	import { axios, getApiErrorMessage } from '$lib/axios';
	import type { GameDto, GameoverDto } from '$lib/models';
	import { getModalStore, getToastStore } from '@skeletonlabs/skeleton';

	export let gameId: number;
	export let onGameover: (gameover: GameoverDto) => void;

	const toastStore = getToastStore();
	const quit = async () => {
		try {
			const response = await axios.delete<GameDto>(`/games/${gameId}`);
			const gameover = response.data.gameover;
			onGameover(gameover);
		} catch (error) {
			toastStore.trigger({
				background: 'variant-filled-error',
				message: getApiErrorMessage(error),
				autohide: true
			});
		}
	};

	const modalStore = getModalStore();
	const confirmQuit = () => {
		modalStore.trigger({
			type: 'confirm',
			title: 'Quit game',
			body: 'Are you sure you want to give up current game?',
			response: (response) => {
				if (response) {
					quit();
				}
			}
		});
	};
</script>

<button class="btn variant-filled-secondary" on:click={confirmQuit}>Give up</button>
