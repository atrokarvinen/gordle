<script lang="ts">
	import { links } from './links';
	import MovileNavigationLink from './MobileNavigationLink.svelte';
</script>

<div class="p-4">
	<nav class="list-nav">
		<ul>
			{#each links as link}
				<MovileNavigationLink {link} />
			{/each}
		</ul>
	</nav>
</div>
