<script lang="ts">
	import { goto } from '$app/navigation';
	import { resolve } from '$app/paths';
	import { page } from '$app/state';
	import { i18n } from '$lib/translations/i18n';
	import { ProgressRing } from '@skeletonlabs/skeleton-svelte';

	export let data;

	$: {
		if (data.loader) {
			data.loader().then((r) => {
				const gameId = r.gameId;

				// Check if user was redirected already once and game was not found.
				// This blocks endless redirect loop.
				const redirected = page.state.redirected;
				const gameNotFound = gameId === -1;
				if (redirected && gameNotFound) {
					retryFailed = true;
					return;
				}

				goto(resolve(`/game/[gameId]`, { gameId: gameId.toString() }));
			});
		}
	}

	let retryFailed = false;
</script>

{#if retryFailed}
	<aside class="alert mt-4 flex items-center preset-filled-error-500">
		<h1 class="h3">Error</h1>
		<div class=" flex flex-row items-center gap-x-2">
			<i class="fas fa-exclamation-triangle fa-xl"></i>
			<p>{$i18n.t('auth_error')}</p>
		</div>
	</aside>
{:else}
	<ProgressRing size="size-16" />
{/if}
