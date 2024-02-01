import { expect } from "@playwright/test";
import { test } from "./game-fixture";

test("allows typing in guess with PC keyboard", async ({ game }) => {
  await game.verifyGameActive();
  await game.page.keyboard.type("tester");

  await expect(game.page.getByTestId("word-guess")).toHaveValue("TESTER");
});

test("allows typing in guess with UI keyboard", async ({ game }) => {
  await game.verifyGameActive();
  await game.typeLetters("tester");

  await expect(game.page.getByTestId("word-guess")).toHaveValue("TESTER");
});

test("erases word correctly", async ({ game }) => {
  await game.verifyGameActive();
  await game.page.keyboard.type("tester");
  await game.page.keyboard.press("Backspace");

  await expect(game.page.getByTestId("word-guess")).toHaveValue("TESTE");
  expect(await game.isLetterActive(5)).toBeTruthy();

  await game.page.keyboard.press("Backspace");
  await expect(game.page.getByTestId("word-guess")).toHaveValue("TEST");
  expect(await game.isLetterActive(4)).toBeTruthy();

  await game.page.keyboard.press("Backspace");
  await game.page.keyboard.press("Backspace");
  await game.page.keyboard.press("Backspace");
  await game.page.keyboard.press("Backspace");
  await game.page.keyboard.press("Backspace");
  await expect(game.page.getByTestId("word-guess")).toHaveValue("");
  expect(await game.isLetterActive(0)).toBeTruthy();
});
