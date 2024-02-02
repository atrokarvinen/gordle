<script lang="ts">
	import type { GameoverDto } from '$lib/models';
	import { i18n } from '$lib/translations/i18n';
	import GameoverDescription from './GameoverDescription.svelte';

	export let gameover: GameoverDto;

	$: isGameover = gameover.isGameover;
	$: isGameWon = gameover.win;
	$: answer = gameover.answer;
</script>

{#if isGameover}
	<div class="w-full space-y-2 px-2 sm:w-96">
		<p>{$i18n.t('gameover')}!</p>
		{#if isGameWon}
			<p>{$i18n.t('victory')}!</p>
		{/if}
		<p>
			{$i18n.t('answer_was')}: <span class="font-bold capitalize">{answer}</span>
		</p>
		<GameoverDescription definitions={gameover.definitions} examples={gameover.examples} />
	</div>
{/if}
