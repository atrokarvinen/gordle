<script lang="ts">
	import { LetterState } from '$lib/models';
	import LetterBox from './LetterBox.svelte';

	export let currentIndex: number;
	export let inputLetters: string[];
	export let letterClicked: (index: number) => void;
</script>

<div class="flex gap-x-1">
	{#each inputLetters as letter, i}
		<LetterBox
			focused={currentIndex === i}
			{letter}
			letterState={LetterState.UNKNOWN}
			value={letter}
			on:click={() => letterClicked(i)}
		/>
	{/each}
</div>
