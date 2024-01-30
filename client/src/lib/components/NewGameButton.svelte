<script lang="ts">
	import { getModalStore } from '@skeletonlabs/skeleton';

	const modalStore = getModalStore();

	const openModal = () => {
		modalStore.trigger({
			title: 'New game',
			body: 'Select the options for the new game.',
			type: 'component',
			component: 'NewGameModal'
		});
	};
</script>

<button class="btn variant-filled-primary" on:click={openModal}>New game</button>
