<script lang="ts">
	import { axios } from '$lib/axios';
	import type { GameDto, GameoverDto } from '$lib/models';
	import { getModalStore } from '@skeletonlabs/skeleton';

	export let gameId: number;
	export let onGameover: (gameover: GameoverDto) => void;

	const quit = async () => {
		const response = await axios.delete<GameDto>(`/games/${gameId}`);
		const gameover = response.data.gameover;
		onGameover(gameover);
	};

	const modalStore = getModalStore();
	const confirmQuit = () => {
		modalStore.trigger({
			type: 'confirm',
			title: 'Confirm',
			body: 'Are you sure you want to quit?',
			response: (response) => {
				if (response) {
					quit();
				}
			}
		});
	};
</script>

<button class="btn variant-filled-secondary" on:click={confirmQuit}>Quit</button>
