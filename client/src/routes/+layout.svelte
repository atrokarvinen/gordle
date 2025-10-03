<script lang="ts">
	import { toaster } from '$lib/toaster';
	import { Toaster } from '@skeletonlabs/skeleton-svelte';
	import '../app.css';
	import Navigation from './Navigation.svelte';
</script>

<Toaster {toaster} />
<div class="flex min-h-screen flex-col">
	<Navigation />
	<div class="align-center flex grow justify-center">
		<slot />
	</div>
</div>
