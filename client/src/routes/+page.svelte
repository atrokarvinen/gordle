<script lang="ts">
	import { goto } from '$app/navigation';
	import { base } from '$app/paths';
	import { page } from '$app/stores';
	import { i18n } from '$lib/translations/i18n';
	import { ProgressRadial } from '@skeletonlabs/skeleton';

	export let data;

	$: {
		if (data.loader) {
			data.loader().then((r) => {
				const gameId = r.gameId;

				// Check if user was redirected already once and game was not found.
				// This blocks endless redirect loop.
				const redirected = $page.state.redirected;
				const gameNotFound = gameId === -1;
				if (redirected && gameNotFound) {
					retryFailed = true;
					return;
				}

				goto(`${base}/game/${gameId}`);
			});
		}
	}

	let retryFailed = false;
</script>

{#if retryFailed}
	<aside class="alert variant-filled-error mt-4 flex items-center">
		<h1 class="h3">Error</h1>
		<div class=" flex flex-row items-center gap-x-2">
			<i class="fas fa-exclamation-triangle fa-xl" />
			<p>{$i18n.t('auth_error')}</p>
		</div>
	</aside>
{:else}
	<ProgressRadial width="w-16" />
{/if}
