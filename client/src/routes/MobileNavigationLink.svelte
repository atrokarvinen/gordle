<script lang="ts">
	import { page } from '$app/state';
	import { i18n } from '$lib/translations/i18n';
	import { type LinkType } from './links';

	interface Props {
		link: LinkType;
	}

	const { link }: Props = $props();

	const isActive = (href: string, currentUrl: string) => {
		if (href === '/' && currentUrl.startsWith('/game')) return true;
		const lastPart = href.split('/').pop();
		if (!lastPart) return false;
		const withoutQuery = lastPart.split('?')[0];
		return currentUrl.includes(withoutQuery);
	};

	let isLinkActive = $derived(isActive(link.href, page.url.pathname));
</script>

<li class="w-full">
	<a
		class="btn w-full {isLinkActive ? 'preset-filled-primary-500' : ''}"
		href={link.href}
		data-testid={link.datatestid}
	>
		<span class="text-center">
			<i class={link.iconClass + ' w-4'}></i>
		</span>
		<span class="flex-auto">{$i18n.t(link.name)}</span>
	</a>
</li>
