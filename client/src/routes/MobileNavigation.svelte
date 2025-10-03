<script lang="ts">
	import { links, type LinkType } from './links';
	import MobileNavigationLink from './MobileNavigationLink.svelte';

	const homeLink: LinkType = {
		datatestid: 'mobile-nav-home-link',
		href: '/',
		iconClass: 'fa-solid fa-house',
		name: 'home',
		ariaLabel: 'Go to home page'
	};
</script>

<nav>
	<ul class="flex flex-col items-center gap-2 p-2">
		<MobileNavigationLink link={homeLink} />
		{#each links as link (link.href)}
			<MobileNavigationLink {link} />
		{/each}
	</ul>
</nav>
