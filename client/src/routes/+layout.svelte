<script lang="ts">
	import '../app.css';
	import Navigation from './Navigation.svelte';
</script>

<!-- <Toast position="t" /> -->

<Navigation />
<div class="align-center flex justify-center">
	<slot />
</div>
