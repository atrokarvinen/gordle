<script lang="ts">
	import { axios, getApiErrorMessage } from '$lib/axios';
	import type { GameDto, GameoverDto } from '$lib/models';
	import { toaster } from '$lib/toaster';
	import { i18n } from '$lib/translations/i18n';

	export let gameId: number;
	export let onGameover: (gameover: GameoverDto) => void;

	const quit = async () => {
		try {
			const response = await axios.delete<GameDto>(`/games/${gameId}`);
			const gameover = response.data.gameover;
			onGameover(gameover);
		} catch (error) {
			toaster.error({
				title: $i18n.t(getApiErrorMessage(error).data, { data: getApiErrorMessage(error).data })
			});
		}
	};

	const confirmQuit = () => {
		// modalStore.trigger({
		// 	type: 'confirm',
		// 	title: $i18n.t('quit_game'),
		// 	body: $i18n.t('give_up_confirm') + '?',
		// 	buttonTextCancel: $i18n.t('cancel'),
		// 	buttonTextConfirm: $i18n.t('confirm'),
		// 	response: (response) => {
		// 		if (response) {
		// 			quit();
		// 		}
		// 	}
		// });
	};
</script>

<button class="btn variant-filled-secondary" on:click={confirmQuit}>{$i18n.t('give_up')}</button>
