<script lang="ts">
	import { axios, getApiErrorMessage } from '$lib/axios';
	import type { GameDto, GameoverDto } from '$lib/models';
	import { toaster } from '$lib/toaster';
	import { i18n } from '$lib/translations/i18n';
	import Modal from './Modal.svelte';

	interface Props {
		gameId: number;
		onGameover: (gameover: GameoverDto) => void;
	}

	const { gameId, onGameover }: Props = $props();

	let show = $state(false);

	async function quit() {
		try {
			const response = await axios.delete<GameDto>(`/games/${gameId}`);
			const gameover = response.data.gameover;
			onGameover(gameover);
		} catch (error) {
			toaster.error({
				title: $i18n.t(getApiErrorMessage(error).data, { data: getApiErrorMessage(error).data })
			});
		}
	}

	const confirmQuit = () => {
		show = true;
	};
</script>

<Modal
	{show}
	onClose={() => (show = false)}
	title={$i18n.t('quit_game')}
	buttonTextCancel={$i18n.t('cancel')}
	buttonTextConfirm={$i18n.t('confirm')}
	response={(response) => {
		if (response) {
			quit();
		}
	}}
>
	{#snippet body()}
		<p>{$i18n.t('give_up_confirm') + '?'}</p>
	{/snippet}
</Modal>
<button class="btn preset-filled-secondary-500" onclick={confirmQuit}>{$i18n.t('give_up')}</button>
