<script lang="ts">
	import { goto } from '$app/navigation';
	import { resolve } from '$app/paths';
	import { axios, getApiErrorMessage } from '$lib/axios';
	import { languageStore } from '$lib/languageStore';
	import type { Difficulty, GameDto } from '$lib/models';
	import { toaster } from '$lib/toaster';
	import { i18n } from '$lib/translations/i18n';
	import LanguageSelect from './LanguageSelect.svelte';

	interface Props {
		difficulty?: Difficulty;
		maxAttempts?: number;
		wordLength?: number;
		onClose: () => void;
	}

	let { maxAttempts = 6, wordLength = 5, difficulty = 'all_words', onClose }: Props = $props();
	let language = $state($languageStore);

	const difficultyOptions: Difficulty[] = ['easy_words', 'all_words'];
	const attemptsOptions = [4, 5, 6, 7, 8];
	const wordLengthOptions = [5, 6, 7, 8];

	const createGame = async () => {
		try {
			const payload = { difficulty, maxAttempts, wordLength, language };
			console.log('payload:', payload);
			const response = await axios.post<GameDto>(`/games`, payload);
			const game = response.data;
			console.log('created game:', game);

			toaster.success({ title: $i18n.t('new_game_started') });
			goto(resolve(`/game/[gameId]`, { gameId: game.id.toString() }));
		} catch (error) {
			const err = getApiErrorMessage(error);
			toaster.error({
				title: $i18n.t(err.message, { data: err.data })
			});
		}
	};
</script>

<div class="w-modal card space-y-4 p-4">
	<form class="form space-y-2">
		<LanguageSelect value={language} onChange={(lang) => (language = lang)} />
		<label class="label">
			{$i18n.t('difficulty')}
			<select class="select" bind:value={difficulty}>
				{#each difficultyOptions as option (option)}
					<option value={option}>{$i18n.t(option)}</option>
				{/each}
			</select>
		</label>
		<label class="label">
			{$i18n.t('word_length')}
			<select class="select" bind:value={wordLength}>
				{#each wordLengthOptions as option (option)}
					<option value={option}>{option}</option>
				{/each}
			</select>
		</label>
		<label class="label">
			{$i18n.t('max_guesses')}
			<select class="select" bind:value={maxAttempts}>
				{#each attemptsOptions as option (option)}
					<option value={option}>{option}</option>
				{/each}
			</select>
		</label>
	</form>
	<footer class="flex flex-row justify-end gap-x-2">
		<button class="btn preset-filled-secondary-500" onclick={onClose}>{$i18n.t('close')}</button>
		<button class="btn preset-filled-primary-500" onclick={createGame}>{$i18n.t('create')}</button>
	</footer>
</div>
