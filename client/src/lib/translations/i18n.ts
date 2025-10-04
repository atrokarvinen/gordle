import i18next from 'i18next';
import { createI18nStore } from 'svelte-i18next';
import { germanKeys } from './de';
import { englishKeys } from './en';
import { finnishKeys } from './fi';
import { polishKeys } from './pl';
import { swedishKeys } from './se';
import { detectLanguage, uiLanguageStore } from './uiLanguageStore';

i18next.init({
	lng: detectLanguage(),
	resources: {
		...englishKeys,
		...finnishKeys,
		...swedishKeys,
		...germanKeys,
		...polishKeys
	},
	interpolation: { escapeValue: false }
});

uiLanguageStore.subscribe((value) => i18next.changeLanguage(value));

export const i18n = createI18nStore(i18next);
