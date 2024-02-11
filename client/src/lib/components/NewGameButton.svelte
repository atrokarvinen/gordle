<script lang="ts">
	import { i18n } from '$lib/translations/i18n';
	import { getModalStore } from '@skeletonlabs/skeleton';

	export let wordLength: number;
	export let maxAttempts: number;

	const modalStore = getModalStore();

	const openModal = () => {
		modalStore.trigger({
			title: $i18n.t('new_game'),
			body: $i18n.t('new_game_info'),
			type: 'component',
			component: 'NewGameModal',
			meta: {
				wordLength,
				maxAttempts
			}
		});
	};
</script>

<button class="btn variant-filled-primary" on:click={openModal}>{$i18n.t('new_game')}</button>
