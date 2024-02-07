import { expect, test } from "@playwright/test";

test.beforeEach(async ({ page }) => {
  await page.goto("/");
});

test("opens game page", async ({ page }) => {
  await expect(page.getByText("Give up")).toBeVisible();
});

test("navigation works", async ({ page }) => {
  await page.getByTestId("info").click();
  await expect(page.getByText("How to play")).toBeVisible();

  await page.getByTestId("home-link").click();
  await expect(page.getByText("Give up")).toBeVisible();
});
