<script lang="ts">
	import { i18n } from '$lib/translations/i18n';
	import { Accordion, AccordionItem } from '@skeletonlabs/skeleton';
	import GameoverDescriptionBrowser from './GameoverDescriptionBrowser.svelte';

	export let definitions: string[];
	export let examples: string[];

	let definitionsIndex = 0;
	$: displayedDefinition = definitions[definitionsIndex]?.trim();

	let examplesIndex = 0;
	$: displayedExample = examples[examplesIndex]?.trim();
</script>

<Accordion padding="p-2" regionControl="variant-ringed-surface">
	{#if definitions.length > 0}
		<AccordionItem open>
			<svelte:fragment slot="lead"><i class="fas fa-book w-2" /></svelte:fragment>
			<svelte:fragment slot="summary">{$i18n.t('definition')}</svelte:fragment>
			<svelte:fragment slot="content">
				<div class="my-2 flex flex-col justify-between">
					<div class="h-16 overflow-y-auto">
						<p class="italic first-letter:capitalize">
							"{`${displayedDefinition}${displayedDefinition.endsWith('.') ? '' : '.'}`}"
						</p>
					</div>
					{#if definitions.length > 1}
						<GameoverDescriptionBrowser bind:browseIndex={definitionsIndex} items={definitions} />
					{/if}
				</div>
			</svelte:fragment>
		</AccordionItem>
	{/if}
	{#if examples.length > 0}
		<AccordionItem>
			<svelte:fragment slot="lead"><i class="fas fa-lightbulb w-2" /></svelte:fragment>
			<svelte:fragment slot="summary">{$i18n.t('example')}</svelte:fragment>
			<svelte:fragment slot="content">
				<div class="my-2 flex flex-col justify-between">
					<div class="h-16 overflow-y-auto">
						<p class="italic first-letter:capitalize">
							"{`${displayedExample}${displayedExample.endsWith('.') ? '' : '.'}`}"
						</p>
					</div>
					{#if examples.length > 1}
						<GameoverDescriptionBrowser bind:browseIndex={examplesIndex} items={examples} />
					{/if}
				</div>
			</svelte:fragment>
		</AccordionItem>
	{/if}
</Accordion>
