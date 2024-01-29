<script lang="ts">
	import type { GameoverDto } from '$lib/models';

	export let gameover: GameoverDto;

	$: isGameover = gameover.isGameover;
	$: isGameWon = gameover.win;
	$: answer = gameover.answer;
	$: description = gameover.answerDescription;
</script>

{#if isGameover}
	<div>
		<p>Game over!</p>
		{#if isGameWon}
			<p>Victory!</p>
		{/if}
		<p>
			Answer was: <span class="font-bold capitalize">{answer}</span>
		</p>
		<p class="italic first-letter:capitalize">
			"{description}"
		</p>
	</div>
{/if}
