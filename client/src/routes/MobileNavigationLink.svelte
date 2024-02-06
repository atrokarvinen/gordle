<script lang="ts">
	import { page } from '$app/stores';
	import { i18n } from '$lib/translations/i18n';
	import { getDrawerStore } from '@skeletonlabs/skeleton';
	import { type LinkType } from './links';

	const drawerStore = getDrawerStore();

	export let link: LinkType;

	$: isLinkActive = isActive(link.href, $page.url.pathname);

	const isActive = (href: string, currentUrl: string) => {
		const lastPart = href.split('/').pop();
		if (!lastPart) return false;
		const withoutQuery = lastPart.split('?')[0];
		return currentUrl.includes(withoutQuery);
	};
</script>

<li>
	<a
		href={link.href}
		data-testid={link.datatestid}
		class={isLinkActive ? 'bg-primary-active-token' : ''}
		on:click={drawerStore.close}
	>
		<span class="text-center">
			<i class={link.iconClass + ' w-4'} />
		</span>
		<span class="flex-auto">{$i18n.t(link.name)}</span>
	</a>
</li>
