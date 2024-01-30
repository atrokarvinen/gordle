<script lang="ts">
	import { LetterState } from '$lib/models';
	import LetterInfo from './LetterInfo.svelte';

	const textCorrect = 'Letter is in correct place';
	const textContained = 'Letter is in the word but in wrong place';
	const textIncorrect = 'Letter is not in the word';
	const textUnknown = 'Letter has not been guessed yet';
</script>

<div class="flex flex-col items-start gap-y-2 px-2 md:w-96">
	<h1 class="h3 self-center">How to play</h1>
	<p>
		Guess the word by typing in letters and submit by pressing Enter. Inserted words must be real
		words. Player has limited amount of guesses before the game is over.
	</p>
	<LetterInfo letter={'A'} letterState={LetterState.CORRECT} text={textCorrect} />
	<LetterInfo letter={'A'} letterState={LetterState.CONTAINED} text={textContained} />
	<LetterInfo letter={'A'} letterState={LetterState.INCORRECT} text={textIncorrect} />
	<LetterInfo letter={'A'} letterState={LetterState.UNKNOWN} text={textUnknown} />
</div>
