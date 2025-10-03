<script lang="ts">
	import type { Snippet } from 'svelte';

	interface Props {
		show: boolean;
		onClose: () => void;
		title?: string;
		buttonTextCancel?: string;
		buttonTextConfirm?: string;
		response?: (response: boolean) => void;
		body?: Snippet;
	}

	let {
		show,
		onClose,
		title = 'Title',
		buttonTextCancel,
		buttonTextConfirm,
		response = () => {},
		body
	}: Props = $props();

	let htmlDialog: HTMLDialogElement | null = null;
	$effect(() => {
		if (show) {
			htmlDialog?.showModal();
		} else {
			htmlDialog?.close();
		}
	});
</script>

<dialog
	data-dialog
	class="rounded-container bg-surface-100-900
		-translate-1/2 backdrop:bg-surface-50/75
		dark:backdrop:bg-surface-950/75 left-1/2
		top-1/2 z-10 w-72 max-w-[640px] space-y-4 p-4 text-inherit"
	bind:this={htmlDialog}
>
	<h2 class="h3">{title}</h2>
	{#if body}
		{@render body()}
	{/if}
	{#if buttonTextCancel || buttonTextConfirm}
		<form method="dialog" class="flex justify-end gap-x-2">
			<button
				type="button"
				class="btn preset-filled-secondary-500"
				data-dialog-close
				onclick={() => {
					response(false);
					onClose();
				}}>{buttonTextCancel}</button
			>
			<button
				type="button"
				class="btn preset-filled-primary-500"
				onclick={() => {
					response(true);
					onClose();
				}}>{buttonTextConfirm}</button
			>
		</form>
	{/if}
</dialog>

<style>
	/* NOTE: add the following styles to your global stylesheet. */
	dialog,
	dialog::backdrop {
		--anim-duration: 250ms;
		transition:
			display var(--anim-duration) allow-discrete,
			overlay var(--anim-duration) allow-discrete,
			opacity var(--anim-duration);
		opacity: 0;
	}
	/* Animate In */
	dialog[open],
	dialog[open]::backdrop {
		opacity: 1;
	}
	/* Animate Out */
	@starting-style {
		dialog[open],
		dialog[open]::backdrop {
			opacity: 0;
		}
	}
</style>
