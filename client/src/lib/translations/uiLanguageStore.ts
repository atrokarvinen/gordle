import { writable } from 'svelte/store';
import type { Language } from './language';

export const uiLanguageStore = writable<Language>('en');
