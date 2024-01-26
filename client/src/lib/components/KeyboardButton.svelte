<script lang="ts">
	import type { LetterState } from '$lib/models';

	import LetterBox from './LetterBox.svelte';

	export let letter: string;
	export let state: LetterState;
	export let onKeyDown: (key: string) => void;
</script>

<button on:click={() => onKeyDown(letter)}>
	<LetterBox {letter} letterState={state} cursor="cursor-pointer" />
</button>
