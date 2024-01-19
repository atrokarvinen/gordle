export enum LetterState {
	CORRECT,
	CONTAINED,
	INCORRECT,
	UNKNOWN
}

export type GuessedLetter = {
	letter: string;
	state: LetterState;
};

export type Guess = {
	word: string;
	letters: GuessedLetter[];
};

export type GuessDto = {
	id?: number;
	gameId: number;
	word: string;
};

export type GuessResultDto = string[];

export type GameDto = {
	id: number;
	name: string;
	maxAttempts: number;
	wordLength: number;
};
