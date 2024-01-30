<script lang="ts">
	import { goto } from '$app/navigation';
	import { base } from '$app/paths';
	import { axios, getApiErrorMessage } from '$lib/axios';
	import { languageStore } from '$lib/languageStore';
	import type { GameDto } from '$lib/models';
	import { getModalStore, getToastStore } from '@skeletonlabs/skeleton';
	import type { SvelteComponent } from 'svelte';

	export let parent: SvelteComponent;

	const modalStore = getModalStore();

	let maxAttempts = 6;
	let wordLength = 6;
	let language = $languageStore;

	const attemptsOptions = [4, 5, 6, 7, 8];
	const wordLengthOptions = [5, 6, 7, 8];
	const languageOptions = ['en', 'fi'];

	const toastStore = getToastStore();
	const createGame = async () => {
		try {
			const payload = { maxAttempts, wordLength, language };
			console.log('payload:', payload);
			const response = await axios.post<GameDto>(`/games`, payload);
			const game = response.data;
			console.log('created game:', game);

			toastStore.trigger({
				background: 'variant-filled-success',
				message: 'New game started',
				autohide: true
			});
			goto(`${base}/game/${game.id}`);
			modalStore.close();
		} catch (error) {
			toastStore.trigger({
				background: 'variant-filled-error',
				message: getApiErrorMessage(error),
				autohide: true
			});
		}
	};

	const selectedLangBorder = 'border-success-500 ';
	const unselectedLangBorder = 'border-transparent ';
</script>

{#if $modalStore[0]}
	<div class="card w-modal space-y-4 p-4">
		<header class="h2">{$modalStore[0].title}</header>
		<article>{$modalStore[0].body}</article>
		<form class="form space-y-2">
			<div class="flex gap-x-3">
				{#each languageOptions as lang}
					<button
						on:click={() => (language = lang)}
						class={`rounded border-2 border-solid p-2 ${lang === language ? selectedLangBorder : unselectedLangBorder}`}
						style="outline: none;"
					>
						<img alt={lang} src="{base}/{lang}.png" class="w-16" />
					</button>
				{/each}
			</div>
			<label class="label">
				Max guesses
				<select class="select" bind:value={maxAttempts}>
					{#each attemptsOptions as option}
						<option value={option}>{option}</option>
					{/each}
				</select>
			</label>
			<label class="label">
				Word length
				<select class="select" bind:value={wordLength}>
					{#each wordLengthOptions as option}
						<option value={option}>{option}</option>
					{/each}
				</select>
			</label>
		</form>
		<footer class="modal-footer {parent.regionFooter}">
			<button class="btn variant-filled-secondary" on:click={modalStore.close}>Close</button>
			<button class="btn variant-filled-primary" on:click={createGame}>Create</button>
		</footer>
	</div>
{/if}
