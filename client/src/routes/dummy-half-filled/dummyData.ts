import { LetterState, type Guess } from '../../lib/models';

export const guesses: Guess[] = [
	{
		word: 'grail',
		letters: [
			{ letter: 'g', state: LetterState.INCORRECT },
			{ letter: 'r', state: LetterState.CORRECT },
			{ letter: 'a', state: LetterState.CORRECT },
			{ letter: 'i', state: LetterState.INCORRECT },
			{ letter: 'l', state: LetterState.INCORRECT }
		]
	},
	{
		word: 'track',
		letters: [
			{ letter: 't', state: LetterState.INCORRECT },
			{ letter: 'r', state: LetterState.CORRECT },
			{ letter: 'a', state: LetterState.CORRECT },
			{ letter: 'c', state: LetterState.CONTAINED },
			{ letter: 'k', state: LetterState.INCORRECT }
		]
	},
	{
		word: 'cramp',
		letters: [
			{ letter: 'c', state: LetterState.CORRECT },
			{ letter: 'r', state: LetterState.CORRECT },
			{ letter: 'a', state: LetterState.CORRECT },
			{ letter: 'm', state: LetterState.INCORRECT },
			{ letter: 'p', state: LetterState.INCORRECT }
		]
	}
];
