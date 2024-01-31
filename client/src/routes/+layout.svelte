<script lang="ts">
	import NewGameModal from '$lib/components/NewGameModal.svelte';
	import {
		AppShell,
		Modal,
		Toast,
		initializeStores,
		type ModalComponent
	} from '@skeletonlabs/skeleton';
	import '../app.pcss';
	import Navigation from './Navigation.svelte';

	initializeStores();

	const modalRegistry: Record<string, ModalComponent> = {
		NewGameModal: { ref: NewGameModal }
	};
</script>

<Modal components={modalRegistry} />
<Toast position="t" />
<AppShell regionPage="" slotPageContent="md:p-4 p-1 items-center flex flex-col">
	<svelte:fragment slot="header">
		<Navigation />
	</svelte:fragment>
	<slot />
</AppShell>
