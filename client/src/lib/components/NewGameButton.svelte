<script lang="ts">
	import type { Difficulty } from '$lib/models';
	import { i18n } from '$lib/translations/i18n';
	import Modal from './Modal.svelte';
	import NewGameModal from './NewGameModal.svelte';

	export let wordLength: number;
	export let maxAttempts: number;
	export let difficulty: Difficulty;

	let show = false;
</script>

<Modal {show} onClose={() => (show = false)} title={$i18n.t('new_game')}>
	{#snippet body()}
		<p>{$i18n.t('new_game_info')}</p>
		<NewGameModal {wordLength} {maxAttempts} {difficulty} onClose={() => (show = false)} />
	{/snippet}
</Modal>
<button class="btn preset-filled-primary-500" on:click={() => (show = true)}
	>{$i18n.t('new_game')}</button
>
