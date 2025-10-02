<script>
	import StatisticsTabs from './StatisticsTabs.svelte';
</script>

<div class="flex w-full flex-col items-center px-2 md:w-96">
	<StatisticsTabs />
	<div class="mt-4">
		<slot />
	</div>
</div>
