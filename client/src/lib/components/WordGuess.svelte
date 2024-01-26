<script lang="ts">
	import { LetterState } from '$lib/models';
	import LetterBox from './LetterBox.svelte';

	export let currentIndex: number;
	export let inputLetters: string[];
</script>

<div class="flex gap-x-1">
	{#each inputLetters as letter, i}
		<LetterBox
			focused={currentIndex === i}
			{letter}
			letterState={LetterState.UNKNOWN}
			value={letter}
		/>
	{/each}
</div>
