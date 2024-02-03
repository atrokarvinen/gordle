import i18next from 'i18next';
import { createI18nStore } from 'svelte-i18next';
import { englishKeys } from './en';
import { finnishKeys } from './fi';
import { detectLanguage, uiLanguageStore } from './uiLanguageStore';

i18next.init({
	lng: detectLanguage(),
	resources: { ...englishKeys, ...finnishKeys },
	interpolation: { escapeValue: false }
});

uiLanguageStore.subscribe((value) => i18next.changeLanguage(value));

export const i18n = createI18nStore(i18next);
