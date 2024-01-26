<script lang="ts">
	import type { LetterState } from '$lib/models';

	import LetterBox from './LetterBox.svelte';

	export let letter: string;
	export let state: LetterState;
	export let icon: string | undefined = undefined;

	const dispatchKey = () => {
		console.log('dispatching key:', letter);
		const event = new KeyboardEvent('keydown', { key: letter });
		window.dispatchEvent(event);
	};
</script>

<button on:click={dispatchKey}>
	<LetterBox {letter} letterState={state} cursor="cursor-pointer" as="button" {icon} />
</button>
