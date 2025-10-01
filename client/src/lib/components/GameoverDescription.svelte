<script lang="ts">
	import { i18n } from '$lib/translations/i18n';
	import { Accordion } from '@skeletonlabs/skeleton-svelte';
	import GameoverDescriptionBrowser from './GameoverDescriptionBrowser.svelte';

	export let definitions: string[];
	export let examples: string[];

	let definitionsIndex = 0;
	$: displayedDefinition = getElement(definitions, definitionsIndex);

	let examplesIndex = 0;
	$: displayedExample = getElement(examples, examplesIndex);

	const getElement = (items: string[], index: number) => {
		if (items.length === 0 || index >= items.length) return '';
		const element = items[index];
		return element.trim();
	};

	const padSentence = (sentence: string) => {
		if (sentence.endsWith('.')) return sentence;
		if (sentence.endsWith('?')) return sentence;
		if (sentence.endsWith('!')) return sentence;
		return `${sentence}.`;
	};
</script>

<Accordion padding="p-2">
	{#if definitions.length > 0}
		<Accordion.Item value="definitions">
			{#snippet lead()}<i class="fas fa-book w-2"></i>{/snippet}
			{#snippet control()}{$i18n.t('definition')}{/snippet}
			{#snippet panel()}
				<div class="my-2 flex flex-col justify-between">
					<div class="h-20 overflow-y-auto">
						<p class="italic first-letter:capitalize">
							"{padSentence(displayedDefinition)}"
						</p>
					</div>
					{#if definitions.length > 1}
						<GameoverDescriptionBrowser bind:browseIndex={definitionsIndex} items={definitions} />
					{/if}
				</div>
			{/snippet}
		</Accordion.Item>
	{/if}
	{#if examples.length > 0}
		<Accordion.Item value="examples">
			{#snippet lead()}<i class="fas fa-lightbulb w-2"></i>{/snippet}
			{#snippet control()}{$i18n.t('example')}{/snippet}
			{#snippet panel()}
				<div class="my-2 flex flex-col justify-between">
					<div class="h-16 overflow-y-auto">
						<p class="italic first-letter:capitalize">
							"{padSentence(displayedExample)}"
						</p>
					</div>
					{#if examples.length > 1}
						<GameoverDescriptionBrowser bind:browseIndex={examplesIndex} items={examples} />
					{/if}
				</div>
			{/snippet}
		</Accordion.Item>
	{/if}
</Accordion>
