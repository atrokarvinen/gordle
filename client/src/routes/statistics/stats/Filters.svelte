<script lang="ts">
	import LanguageSelect from '$lib/components/LanguageSelect.svelte';
	import { i18n } from '$lib/translations/i18n';
	import type { Language } from '$lib/translations/language';

	export let selectedLanguage: Language | undefined;
	export let selectedWordLength: number = -1;
	const languageChanged = (language: Language) => {
		if (selectedLanguage === language) {
			selectedLanguage = undefined;
			return;
		}
		selectedLanguage = language;
	};

	const wordLengthOptions = [
		{ value: -1, label: $i18n.t('all') },
		{ value: 5, label: '5' },
		{ value: 6, label: '6' },
		{ value: 7, label: '7' },
		{ value: 8, label: '8' }
	];
</script>

<p>{$i18n.t('filter_by_language')}:</p>
<div class="flex items-center">
	<LanguageSelect
		languageOptions={['en', 'fi']}
		onChange={languageChanged}
		value={selectedLanguage}
	/>
	{#if selectedLanguage}
		<div class="ml-5">
			<button
				aria-label="Clear language filter"
				class="btn variant-filled-secondary"
				on:click={() => (selectedLanguage = undefined)}
			>
				<span>
					<i class="fas fa-times"></i>
				</span>
				<span>{$i18n.t('clear')}</span>
			</button>
		</div>
	{/if}
</div>

<label>
	{$i18n.t('filter_by_word_length')}:
	<select class="select w-32" bind:value={selectedWordLength}>
		{#each wordLengthOptions as wordLengthOption}
			<option value={wordLengthOption.value}>{wordLengthOption.label}</option>
		{/each}
	</select>
</label>
