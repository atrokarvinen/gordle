<script lang="ts">
	import { LetterState } from '$lib/models';
	import { i18n } from '$lib/translations/i18n';
	import LetterInfo from './LetterInfo.svelte';
</script>

<div class="flex flex-col items-start gap-y-2 px-2 md:w-96">
	<h1 class="self-center h3">{$i18n.t('how_to_play')}</h1>
	<p>{$i18n.t('info_paragraph')}</p>
	<LetterInfo letter="A" letterState={LetterState.CORRECT} text={$i18n.t('letter_correct')} />
	<LetterInfo letter="A" letterState={LetterState.CONTAINED} text={$i18n.t('letter_contained')} />
	<LetterInfo letter="A" letterState={LetterState.INCORRECT} text={$i18n.t('letter_incorrect')} />
	<LetterInfo letter="A" letterState={LetterState.UNKNOWN} text={$i18n.t('letter_unknown')} />
</div>
