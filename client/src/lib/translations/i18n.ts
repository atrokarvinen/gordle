import i18next from 'i18next';
import { createI18nStore } from 'svelte-i18next';
import { englishKeys } from './en';
import { finnishKeys } from './fi';

i18next.init({
	lng: 'en',
	resources: {
		...englishKeys,
		...finnishKeys
	},
	interpolation: {
		escapeValue: false // not needed for svelte as it escapes by default
	}
});

export const i18n = createI18nStore(i18next);
