import { browser } from '$app/environment';
import { writable } from 'svelte/store';
import type { Language } from './language';

export const detectLanguage = () => {
	if (!browser) return 'en';
	const localStorageLang = localStorage.getItem('uiLanguage');
	const browserLang = navigator.language;
	const lang = localStorageLang ?? browserLang ?? 'en';
	return lang === 'fi' ? 'fi' : 'en';
};

export const uiLanguageStore = writable<Language>(detectLanguage());
