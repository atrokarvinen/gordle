<script lang="ts">
	import type { LetterState } from '$lib/models';

	import LetterBox from './LetterBox.svelte';

	export let letter: string;
	export let state: LetterState;
	export let onKeyDown: (key: string) => void;
</script>

<button class="button" on:click={() => onKeyDown(letter)}>
	<LetterBox {letter} letterState={state} />
</button>

<style>
	.button {
		border: none;
		background-color: transparent;
		padding: 0;
		margin: 0;
	}
	.button:hover {
		cursor: pointer;
	}
</style>
