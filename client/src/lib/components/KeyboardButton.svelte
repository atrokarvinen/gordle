<script lang="ts">
	import type { LetterState } from '$lib/models';
	import { ProgressRadial } from '@skeletonlabs/skeleton';

	import LetterBox from './LetterBox.svelte';

	export let letter: string;
	export let state: LetterState;
	export let icon: string | undefined = undefined;
	export let submitting = false;

	const dispatchKeyboardEvent = () => {
		const event = new KeyboardEvent('keydown', { key: letter });
		window.dispatchEvent(event);
	};
</script>

<button data-testid="keyboard-button" on:click={dispatchKeyboardEvent}>
	{#if submitting}
		<div class="bg-surface-400 flex h-7 w-12 items-center justify-center md:h-12 md:w-20">
			<ProgressRadial width="md:w-10 w-6" />
		</div>
	{:else}
		<LetterBox {letter} letterState={state} cursor="cursor-pointer" as="button" {icon} />
	{/if}
</button>
