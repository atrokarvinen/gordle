<script lang="ts">
	interface Props {
		visible: boolean;
		onClose: () => void;
	}

	let { visible, onClose }: Props = $props();
</script>

{#if visible}
	<div
		class="backdrop"
		tabindex="0"
		role="button"
		aria-label="Close navigation menu"
		onclick={onClose}
	></div>
{/if}

<style>
	.backdrop {
		position: absolute;
		top: 0;
		left: 0;
		width: 100%;
		height: 100%;
		background-color: rgba(0, 0, 0, 0.5);
		display: flex;
		justify-content: flex-end;
		z-index: 1000;
	}
</style>
