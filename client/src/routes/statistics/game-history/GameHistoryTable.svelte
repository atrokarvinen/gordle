<script lang="ts">
	import { goto } from '$app/navigation';
	import { base } from '$app/paths';
	import { page } from '$app/stores';
	import type { GameDto } from '$lib/models';
	import { uiLanguageStore } from '$lib/translations/uiLanguageStore';
	import { Paginator } from '@skeletonlabs/skeleton';

	export let games: GameDto[];
	export let totalCount: number;

	const rowClass = (game: GameDto) => {
		return game.gameover.win ? 'variant-filled-success' : 'variant-filled-error';
	};
	const rowClassEven = (game: GameDto) => {
		return game.gameover.win ? 'even:variant-filled-success' : 'even:variant-filled-error';
	};

	let paginationSettings = {
		page: +$page.url.searchParams.get('page')! || 0,
		limit: +$page.url.searchParams.get('limit')! || 5,
		size: totalCount,
		amounts: [5, 10, 20, 50]
	};

	$: paginationSettings.size = totalCount;

	const onPageChange = (e: CustomEvent<number>) => {
		const pageNumber = e.detail;
		let query = new URLSearchParams($page.url.searchParams.toString());
		query.set('page', pageNumber.toString());
		goto(`?${query.toString()}`);
	};

	const onAmountChange = (e: CustomEvent<number>) => {
		const amount = e.detail;
		let query = new URLSearchParams($page.url.searchParams.toString());
		const maxPage = Math.ceil(totalCount / amount) - 1;
		const pageNumber = Math.min(paginationSettings.page, maxPage);

		query.set('page', pageNumber.toString());
		query.set('limit', amount.toString());
		goto(`?${query.toString()}`);
	};
</script>

<div class="space-y-2">
	<Paginator
		select="select w-36 flex m-auto"
		amountText="Games"
		showNumerals
		maxNumerals={1}
		bind:settings={paginationSettings}
		on:page={onPageChange}
		on:amount={onAmountChange}
	/>
	<div class="table-container">
		<table class="table-hover table">
			<tbody>
				{#each games as game}
					<tr class={`${rowClass(game)} ${rowClassEven(game)}`}>
						<td class="w-16"><img class="w-8" src="{base}/{game.language}.png" alt="en" /></td>
						<td class="capitalize">{game.gameover.answer}</td>
						<td class="w-16">{new Date(game.createdAt).toLocaleDateString($uiLanguageStore)}</td>
						<td class="w-8"
							><a href="{base}/game/{game.id}" class="btn btn-sm variant-filled-primary">View</a
							></td
						>
					</tr>
				{/each}
			</tbody>
		</table>
	</div>
</div>

<style>
	.table tbody td {
		vertical-align: middle;
	}
</style>
