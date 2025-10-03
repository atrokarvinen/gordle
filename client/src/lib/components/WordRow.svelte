<script lang="ts">
	import { type Guess } from '$lib/models';
	import LetterBox from './LetterBox.svelte';

	export let word: Guess;
</script>

{#each word.letters as letter (letter.letter + letter.state)}
	<LetterBox letter={letter.letter} letterState={letter.state} />
{/each}
<span data-testid="guessed-word" hidden>{word.word}</span>
