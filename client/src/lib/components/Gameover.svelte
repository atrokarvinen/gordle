<script lang="ts">
	import type { GameoverDto } from '$lib/models';

	export let gameover: GameoverDto;

	let descriptionIndex = 0;
	let descriptions: string[];

	$: isGameover = gameover.isGameover;
	$: isGameWon = gameover.win;
	$: answer = gameover.answer;
	$: displayedDescription = descriptions[descriptionIndex];
	$: {
		descriptions = gameover.answerDescription.split(';;');
		descriptionIndex = 0;
	}
</script>

{#if isGameover}
	<div class="w-full px-2 md:w-96">
		<p>Game over!</p>
		{#if isGameWon}
			<p>Victory!</p>
		{/if}
		<p>
			Answer was: <span class="font-bold capitalize">{answer}</span>
		</p>
		<div class="my-2 flex flex-col justify-between">
			<div class="h-16 overflow-y-auto">
				<p class="italic first-letter:capitalize">"{displayedDescription}."</p>
			</div>
			{#if descriptions.length > 1}
				<div class="flex items-center justify-end gap-x-2">
					<span>{`${descriptionIndex + 1} / ${descriptions.length}`}</span>
					<button
						class="btn-icon variant-ghost-surface"
						disabled={descriptionIndex === 0}
						on:click={() => (descriptionIndex = Math.max(descriptionIndex - 1, 0))}
					>
						<i class="fas fa-arrow-left" />
					</button>
					<button
						class="btn-icon variant-ghost-surface"
						disabled={descriptionIndex === descriptions.length - 1}
						on:click={() =>
							(descriptionIndex = Math.min(descriptionIndex + 1, descriptions.length - 1))}
					>
						<i class="fas fa-arrow-right" />
					</button>
				</div>
			{/if}
		</div>
	</div>
{/if}
