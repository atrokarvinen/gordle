<script lang="ts">
	import { base } from '$app/paths';
	import type { Language } from '$lib/translations/language';

	export let value: Language | undefined;
	export let onChange: (language: Language) => void;
	export let languageOptions: Language[] = ['en', 'fi'];

	const selectedLangBorder = 'border-success-500';
	const unselectedLangBorder = 'border-transparent';
</script>

<div class="flex gap-x-3">
	{#each languageOptions as language}
		<button
			on:click={() => onChange(language)}
			class={`rounded border-2 border-solid p-2 ${value === language ? selectedLangBorder : unselectedLangBorder}`}
			style="outline: none;"
		>
			<img src="{base}/{language}.png" alt={language} class="w-16" />
		</button>
	{/each}
</div>
