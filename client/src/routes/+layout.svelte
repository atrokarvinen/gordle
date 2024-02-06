<script lang="ts">
	import NewGameModal from '$lib/components/NewGameModal.svelte';
	import {
		AppShell,
		Drawer,
		Modal,
		Toast,
		getDrawerStore,
		initializeStores,
		type ModalComponent
	} from '@skeletonlabs/skeleton';
	import '../app.pcss';
	import MobileNavigation from './MobileNavigation.svelte';
	import Navigation from './Navigation.svelte';

	initializeStores();

	const drawerStore = getDrawerStore();
	const modalRegistry: Record<string, ModalComponent> = {
		NewGameModal: { ref: NewGameModal }
	};
</script>

<Drawer>
	{#if $drawerStore.id === 'mobile-nav'}
		<MobileNavigation />
	{:else}
		<p>Unknown drawer id '{$drawerStore.id}'</p>
	{/if}
</Drawer>
<Modal components={modalRegistry} />
<Toast position="t" />
<AppShell regionPage="" slotPageContent="md:p-4 p-1 items-center flex flex-col">
	<svelte:fragment slot="header">
		<Navigation />
	</svelte:fragment>
	<slot />
</AppShell>
