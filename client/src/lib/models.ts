import type { Language } from './translations/language';

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
	definitions: string[];
	examples: string[];
};

export type GameDto = {
	id: number;
	maxAttempts: number;
	wordLength: number;
	language: Language;
	gameover: GameoverDto;
	guesses: Guess[];
	state: GameState;
	createdAt: string;
};

export enum GameState {
	UNKNOWN,
	ACTIVE,
	WIN,
	LOSE
}

export type GameHistoryDto = {
	totalCount: number;
	games: GameDto[];
};

export type StatisticsDto = {
	allGames: GameDto[];
	total: Statistics;
	byLanguage: Record<string, Record<number, WordLengthMap>>;
};

export type WordLengthMap = {
	total: Statistics;
	byGuessCount: Record<number, number>;
};

export type Statistics = {
	winCount: number;
	lossCount: number;
	totalCount: number;
	winRate: number;
};
