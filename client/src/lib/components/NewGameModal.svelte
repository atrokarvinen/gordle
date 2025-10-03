<script lang="ts">
	import { goto } from '$app/navigation';
	import { resolve } from '$app/paths';
	import { axios, getApiErrorMessage } from '$lib/axios';
	import { languageStore } from '$lib/languageStore';
	import type { GameDto } from '$lib/models';
	import { toaster } from '$lib/toaster';
	import { i18n } from '$lib/translations/i18n';
	import LanguageSelect from './LanguageSelect.svelte';

	export let maxAttempts = 6;
	export let wordLength = 6;
	export let onClose: () => void;
	let language = $languageStore;

	const attemptsOptions = [4, 5, 6, 7, 8];
	const wordLengthOptions = [5, 6, 7, 8];

	const createGame = async () => {
		try {
			const payload = { maxAttempts, wordLength, language };
			console.log('payload:', payload);
			const response = await axios.post<GameDto>(`/games`, payload);
			const game = response.data;
			console.log('created game:', game);

			toaster.success({
				title: $i18n.t('new_game_started')
			});
			goto(resolve(`/game/[gameId]`, { gameId: game.id.toString() }));
		} catch (error) {
			const err = getApiErrorMessage(error);
			toaster.error({
				title: $i18n.t(err.message, { data: err.data })
			});
		}
	};
</script>

<div class="card w-modal space-y-4 p-4">
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
	<footer class="flex flex-row justify-end gap-x-2">
		<button class="btn preset-filled-secondary-500" on:click={onClose}>{$i18n.t('close')}</button>
		<button class="btn preset-filled-primary-500" on:click={createGame}>{$i18n.t('create')}</button>
	</footer>
</div>
