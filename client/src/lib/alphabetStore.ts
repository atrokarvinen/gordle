import { writable } from 'svelte/store';
import { languageStore } from './languageStore';

export type Alphabets = {
	alphabets: string[];
	qwerty1: string[];
	qwerty2: string[];
	qwerty3: string[];
};

export const alphabetStore = writable<Alphabets>({
	alphabets: 'abcdefghijklmnopqrstuvwxyz'.split(''),
	qwerty1: 'qwertyuiop'.split(''),
	qwerty2: 'asdfghjkl'.split(''),
	qwerty3: 'zxcvbnm'.split('')
});

languageStore.subscribe((lang) => {
	if (lang === 'en') {
		alphabetStore.set({
			alphabets: 'abcdefghijklmnopqrstuvwxyz'.split(''),
			qwerty1: 'qwertyuiop'.split(''),
			qwerty2: 'asdfghjkl'.split(''),
			qwerty3: 'zxcvbnm'.split('')
		});
	} else if (lang === 'fi') {
		alphabetStore.set({
			alphabets: 'abcdefghijklmnopqrstuvwxyzäöå'.split(''),
			qwerty1: 'qwertyuiopå'.split(''),
			qwerty2: 'asdfghjkläö'.split(''),
			qwerty3: 'zxcvbnm'.split('')
		});
	}
});
