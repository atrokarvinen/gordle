<script lang="ts">
	import { base } from '$app/paths';
	import type { Language } from '$lib/translations/language';

	export let value: Language | undefined;
	export let onChange: (language: Language) => void;
	export let languageOptions: Language[] = ['en', 'fi', 'se', 'de', 'pl'];

	const selectedLangBorder = 'border-success-500';
	const unselectedLangBorder = 'border-transparent';
</script>

<div class="flex flex-wrap gap-x-3">
	{#each languageOptions as language (language)}
		<button
			aria-label="Select {language}"
			on:click={() => onChange(language)}
			class={`rounded border-2 border-solid p-2 ${value === language ? selectedLangBorder : unselectedLangBorder}`}
			style="outline: none;"
		>
			<img src="{base}/{language}.png" alt={language} class="w-10" />
		</button>
	{/each}
</div>
