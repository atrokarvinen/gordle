import { expect } from "@playwright/test";
import { BACKEND_URL } from "./config";
import { test } from "./game-fixture";

test.beforeEach(async ({ game }) => {
  await game.verifyGameActive();
  const request = game.page.request;
  const url = BACKEND_URL + "/test/answer";
  await request.put(url, { data: { answer: "please" } });
});

test("loads unfinished game", async ({ game }) => {
  await game.guessWord("sudden");
  await game.verifyRowHasGuess("SUDDEN", 0);

  await game.page.reload();

  await game.verifyRowHasGuess("SUDDEN", 0);
});

test("loads ended game and answer", async ({ game }) => {
  await game.giveUp();
  await game.verifyGameover();

  await game.page.reload();

  await game.verifyGameover();
});

test("loses game by guessing wrong too many times", async ({ game }) => {
  await game.guessWord("sudden");
  await game.guessWord("sudden");
  await game.guessWord("sudden");
  await game.guessWord("sudden");
  await game.guessWord("sudden");

  await game.verifyGameActive();

  await game.guessWord("sudden", true);

  await game.verifyGameover();
});

test("wins game and loads won game correctly", async ({ game }) => {
  await game.guessWord("please", true);

  await game.verifyVictory("please");

  await game.page.reload();

  await game.verifyVictory("please");
});

test("does not allow guessing non-existing word", async ({ game }) => {
  await game.guessWord("asdasd", true);

  const err = 'word "ASDASD" not found in language "en"';
  await expect(game.page.getByText(err)).toBeVisible();
  await game.verifyLetterActive(5);
});
