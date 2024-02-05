<script lang="ts">
	import { base } from '$app/paths';

	export let data;

	$: console.log('loaded games:', data.games);
	$: games = data.games;
</script>

History

<div class="flex flex-col gap-y-2">
	{#each games as game}
		<div class="flex gap-x-2">
			<p>ID: {game.id}</p>
			<p>Win: {game.gameover.win}</p>
			<p>Answer: {game.gameover.answer}</p>
			<a href="{base}/game/{game.id}" class="btn variant-filled-primary">View</a>
		</div>
	{/each}
</div>
