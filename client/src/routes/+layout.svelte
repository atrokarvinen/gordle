<script lang="ts">
	import { AppShell, Modal, Toast, initializeStores } from '@skeletonlabs/skeleton';
	import '../app.pcss';
	import Navigation from './Navigation.svelte';

	initializeStores();
</script>

<Modal />
<Toast position="t" />
<AppShell regionPage="" slotPageContent="md:p-4 p-1 items-center justify-center flex flex-col">
	<svelte:fragment slot="header">
		<Navigation />
	</svelte:fragment>
	<slot />
</AppShell>
