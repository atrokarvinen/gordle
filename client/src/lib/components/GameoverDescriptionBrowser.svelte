<script lang="ts">
	export let items: string[];
	export let browseIndex = 0;
</script>

<div class="flex items-center justify-end gap-x-2">
	<span>{`${browseIndex + 1} / ${items.length}`}</span>
	<button
		class="btn-icon variant-ghost-surface"
		disabled={browseIndex === 0}
		on:click={() => (browseIndex = Math.max(browseIndex - 1, 0))}
	>
		<i class="fas fa-arrow-left" />
	</button>
	<button
		class="btn-icon variant-ghost-surface"
		disabled={browseIndex === items.length - 1}
		on:click={() => (browseIndex = Math.min(browseIndex + 1, items.length - 1))}
	>
		<i class="fas fa-arrow-right" />
	</button>
</div>
