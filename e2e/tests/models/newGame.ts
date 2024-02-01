export type NewGame = {
  maxAttempts: string;
  wordLength: string;
  language: "en" | "fi";
};

export const defaultGame: NewGame = {
  maxAttempts: "6",
  wordLength: "6",
  language: "en",
};
