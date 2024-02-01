import { expect } from "@playwright/test";
import { test } from "./game-fixture";
import { NewGame, defaultGame } from "./models/newGame";

test("creates new game with default values", async ({ game }) => {
  await game.giveUp();
  await game.createNewGame();

  await expect(game.page.getByText("New game started")).toBeVisible();
});

test("creates new game non-default values", async ({ game }) => {
  const form: NewGame = { ...defaultGame, maxAttempts: "5", wordLength: "7" };
  await game.giveUp();
  await game.createNewGame(form);

  await expect(game.page.getByText("New game started")).toBeVisible();
  await expect(game.page.getByTestId("guess-row")).toHaveCount(5);
  const firstRow = await game.page.getByTestId("guess-row").first();
  await expect(firstRow.getByTestId("guess-letter")).toHaveCount(7);
});

test("creates game with different language", async ({ game }) => {
  const form: NewGame = { ...defaultGame, language: "fi" };

  await expect(game.page.getByText("Ä")).toBeHidden();
  await expect(game.page.getByText("Ö")).toBeHidden();
  await expect(game.page.getByText("Å")).toBeHidden();

  await game.giveUp();
  await game.createNewGame(form);

  await expect(game.page.getByText("New game started")).toBeVisible();
  await expect(game.page.getByText("Ä")).toBeVisible();
  await expect(game.page.getByText("Ö")).toBeVisible();
  await expect(game.page.getByText("Å")).toBeVisible();
});
