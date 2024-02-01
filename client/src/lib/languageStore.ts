import { writable } from 'svelte/store';
import type { Language } from './translations/language';

export const languageStore = writable<Language>('en');
