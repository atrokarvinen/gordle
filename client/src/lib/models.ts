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

export type GuessResultDto = {
	results: string[];
	gameover: GameoverDto;
};

export type GameoverDto = {
	isGameover: boolean;
	win: boolean;
	answer: string;
	answerDescription: string;
};

export type GameDto = {
	id: number;
	name: string;
	maxAttempts: number;
	wordLength: number;
	gameover: GameoverDto;
};
