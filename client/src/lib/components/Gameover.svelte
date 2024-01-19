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
		<p>Game over! You {isGameWon ? 'win' : 'lose'}!</p>
		<p>
			Answer was: '{answer}'
		</p>
		<p>
			{description}
		</p>
	</div>
{/if}
