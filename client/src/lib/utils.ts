import { LetterState } from './models';

export const convertLetterState = (state: string): LetterState => {
	if (state === 'v') return LetterState.CORRECT;
	if (state === '?') return LetterState.CONTAINED;
	return LetterState.INCORRECT;
};
