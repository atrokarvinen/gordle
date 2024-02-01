import { test as base } from "@playwright/test";
import { GamePageModel } from "./gamePageModel";

type GameFixture = {
  game: GamePageModel;
};

export const test = base.extend<GameFixture>({
  game: async ({ page }, use) => {
    const gamePage = new GamePageModel(page);
    await page.goto("/");
    await use(gamePage);
  },
});
