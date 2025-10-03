<script lang="ts">
	import '../app.css';
	import Navigation from './Navigation.svelte';
</script>

<div class="flex min-h-screen flex-col">
	<Navigation />
	<div class="align-center flex grow justify-center">
		<slot />
	</div>
</div>
