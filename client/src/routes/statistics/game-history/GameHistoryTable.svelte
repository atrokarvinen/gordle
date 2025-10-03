<script lang="ts">
	import { goto } from '$app/navigation';
	import { base, resolve } from '$app/paths';
	import { page } from '$app/state';
	import type { GameDto } from '$lib/models';
	import { i18n } from '$lib/translations/i18n';
	import { uiLanguageStore } from '$lib/translations/uiLanguageStore';
	import { Pagination } from '@skeletonlabs/skeleton-svelte';
	import { DEFAULT_LIMIT, DEFAULT_PAGE } from './defaults';

	interface Props {
		games: GameDto[];
		totalCount: number;
	}

	const { games, totalCount }: Props = $props();

	const rowClass = (game: GameDto) => {
		return game.gameover.win ? 'preset-filled-success-500' : 'preset-filled-error-500';
	};
	const rowClassEven = (game: GameDto) => {
		return game.gameover.win ? 'even:preset-filled-success-500' : 'even:preset-filled-error-500';
	};

	let paginationSettings = {
		page: +page.url.searchParams.get('page')! || DEFAULT_PAGE,
		limit: +page.url.searchParams.get('limit')! || DEFAULT_LIMIT,
		size: totalCount,
		amounts: [5, 10, 20, 50]
	};

	$effect(() => {
		paginationSettings.size = totalCount;
	});

	let pageNumber = $state(paginationSettings.page);
	let pageSize = $state(paginationSettings.limit);

	let sourceData: any[] = Array.from({ length: 20 }, (_, i) => ({
		position: i + 1,
		name: `Element ${i + 1}`,
		weight: Math.round(Math.random() * 1000) / 100,
		symbol: String.fromCharCode(65 + (i % 26))
	}));

	const onPageChange = (pageNumber: number) => {
		let query = new URLSearchParams(page.url.searchParams.toString());
		query.set('page', pageNumber.toString());
		goto(`?${query.toString()}`);
	};

	const onAmountChange = (amount: number) => {
		let query = new URLSearchParams(page.url.searchParams.toString());
		const maxPage = Math.ceil(totalCount / amount);
		const pageNumber = Math.min(paginationSettings.page, maxPage);

		console.log('amount changed to:', amount, 'maxPage:', maxPage, 'new page:', pageNumber);

		query.set('page', pageNumber.toString());
		query.set('limit', amount.toString());
		goto(`?${query.toString()}`);
	};
</script>

<div class="space-y-2">
	<div class="flex flex-col items-center justify-between gap-y-2 md:flex-row">
		<select
			class="select max-w-[150px]"
			bind:value={pageSize}
			onchange={(e) => {
				const value = Number(e.currentTarget.value);
				onAmountChange(value);
			}}
		>
			{#each paginationSettings.amounts as v}
				<option value={v}>{v} {$i18n.t('games')}</option>
			{/each}
		</select>
		<Pagination
			data={sourceData}
			page={pageNumber}
			onPageChange={(e) => onPageChange(e.page)}
			{pageSize}
			count={totalCount}
			siblingCount={0}
		></Pagination>
	</div>
	<div class="table-container">
		<table class="table-hover table">
			<tbody>
				{#each games as game}
					<tr
						id="table-row"
						class={`${rowClass(game)} ${rowClassEven(game)} ${game.gameover.win ? 'success' : 'error'}`}
					>
						<td class="w-16"><img class="w-8" src="{base}/{game.language}.png" alt="en" /></td>
						<td class="capitalize">{game.gameover.answer}</td>
						<td class="w-16">{new Date(game.createdAt).toLocaleDateString($uiLanguageStore)}</td>
						<td class="w-8"
							><a href={resolve(`/game/${game.id}`)} class="btn btn-sm preset-filled-primary-500"
								>{$i18n.t('view')}</a
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
	:global(.paginator-controls.preset-filled-surface-500) {
		background-color: rgb(var(--color-surface-700));
	}
	#table-row.success:hover {
		background-color: rgb(var(--color-success-700));
	}
	#table-row.error:hover {
		background-color: rgb(var(--color-error-700));
	}
</style>
