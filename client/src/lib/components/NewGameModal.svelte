<script lang="ts">
	import { goto } from '$app/navigation';
	import { base } from '$app/paths';
	import { axios, getApiErrorMessage } from '$lib/axios';
	import { languageStore } from '$lib/languageStore';
	import type { GameDto } from '$lib/models';
	import { i18n } from '$lib/translations/i18n';
	import { getModalStore, getToastStore } from '@skeletonlabs/skeleton';
	import type { SvelteComponent } from 'svelte';
	import LanguageSelect from './LanguageSelect.svelte';

	export let parent: SvelteComponent;

	const modalStore = getModalStore();

	let maxAttempts = $modalStore[0]?.meta.maxAttempts ?? 6;
	let wordLength = $modalStore[0]?.meta.wordLength ?? 6;
	let language = $languageStore;

	const attemptsOptions = [4, 5, 6, 7, 8];
	const wordLengthOptions = [5, 6, 7, 8];

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
				message: $i18n.t('new_game_started'),
				autohide: true
			});
			goto(`${base}/game/${game.id}`);
			modalStore.close();
		} catch (error) {
			const err = getApiErrorMessage(error);
			toastStore.trigger({
				background: 'variant-filled-error',
				message: $i18n.t(err.message, { data: err.data }),
				autohide: true
			});
		}
	};
</script>

{#if $modalStore[0]}
	<div class="card w-modal space-y-4 p-4">
		<header class="h2">{$modalStore[0].title}</header>
		<article>{$modalStore[0].body}</article>
		<form class="form space-y-2">
			<LanguageSelect value={language} onChange={(lang) => (language = lang)} />
			<label class="label">
				{$i18n.t('word_length')}
				<select class="select" bind:value={wordLength}>
					{#each wordLengthOptions as option}
						<option value={option}>{option}</option>
					{/each}
				</select>
			</label>
			<label class="label">
				{$i18n.t('max_guesses')}
				<select class="select" bind:value={maxAttempts}>
					{#each attemptsOptions as option}
						<option value={option}>{option}</option>
					{/each}
				</select>
			</label>
		</form>
		<footer class="modal-footer {parent.regionFooter}">
			<button class="btn variant-filled-secondary" on:click={modalStore.close}
				>{$i18n.t('close')}</button
			>
			<button class="btn variant-filled-primary" on:click={createGame}>{$i18n.t('create')}</button>
		</footer>
	</div>
{/if}
