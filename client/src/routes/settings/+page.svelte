<script lang="ts">
	import LanguageSelect from '$lib/components/LanguageSelect.svelte';
	import { i18n } from '$lib/translations/i18n';
	import { uiLanguageStore } from '$lib/translations/uiLanguageStore';
</script>

<div class="space-y-2 p-4 md:w-96">
	<p>{$i18n.t('language_select_info')}</p>
	<LanguageSelect
		onChange={(lang) => {
			$uiLanguageStore = lang;
			localStorage.setItem('uiLanguage', lang);
		}}
		value={$uiLanguageStore}
	/>
</div>
