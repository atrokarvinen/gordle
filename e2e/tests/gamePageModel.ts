import { Page, expect } from "@playwright/test";
import { NewGame } from "./models/newGame";

export class GamePageModel {
  page: Page;

  constructor(page: Page) {
    this.page = page;
  }

  async giveUp() {
    await this.page.getByRole("button", { name: "Give up" }).click();
    await this.page.getByRole("button", { name: "Confirm" }).click();
  }

  async createNewGame(values: NewGame | undefined = undefined) {
    await this.page.getByRole("button", { name: "New game" }).click();
    if (values) {
      const { language, maxAttempts, wordLength } = values;
      await this.page.getByLabel("Word length").selectOption(wordLength);
      await this.page.getByLabel("Max guesses").selectOption(maxAttempts);
      await this.page.getByAltText(language).click();
    }
    await this.page.getByRole("button", { name: "Create" }).click();
  }

  async typeLetter(letter: string) {
    await this.page
      .getByTestId("keyboard-button")
      .filter({ hasText: letter })
      .click();
  }

  async typeLetters(word: string) {
    const letters = word.split("");
    for (const letter of letters) {
      await this.typeLetter(letter);
    }
  }

  async guessWord(word: string, isLast = false) {
    await this.typeLetters(word);
    await this.page.keyboard.press("Enter");
    if (isLast) return;
    await this.verifyLetterActive(0);
  }

  async verifyLetterActive(letterIndex: number) {
    const letter = await this.page
      .getByTestId("word-guess-row")
      .getByTestId("guess-letter")
      .nth(letterIndex);
    await expect(letter).toHaveClass(/border-primary-500/);
  }

  async isLetterActive(letterIndex: number) {
    const letter = await this.page
      .getByTestId("word-guess-row")
      .getByTestId("guess-letter")
      .nth(letterIndex);
    const classes = (await letter.getAttribute("class"))?.split(" ") ?? [];
    return classes.includes("border-primary-500");
  }

  async verifyRowHasGuess(word: string, index: number) {
    const row = await this.page.getByTestId("guess-row").nth(index);
    const guessedWord = row.getByTestId("guessed-word");
    await expect(guessedWord).toHaveText(word);
  }

  async verifyGameStarted() {
    await this.page.waitForSelector("text=New game started");
  }

  async verifyGameActive() {
    await expect(this.page.getByText("Give up")).toBeVisible();
  }

  async verifyGameover() {
    await expect(this.page.getByText("Game over!")).toBeVisible();
  }

  async verifyVictory(answer: string) {
    await this.verifyGameover();
    await expect(this.page.getByText("Victory!")).toBeVisible();
    await expect(this.page.getByText(`Answer was: ${answer}`)).toBeVisible();
  }
}
