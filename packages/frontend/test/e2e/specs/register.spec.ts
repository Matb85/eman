import Page from "../pageObjects/register.page";
import { mockUser } from "../utils/index";

describe("/register tab - a11y", () => {
  it("should have a proper title", async () => {
    await Page.open();
    expect(browser).toHaveTitle("Redinn");
  });
});

describe("/register tab - input form", () => {
  it("register - OK", async () => {
    await Page.open();
    const user = mockUser();
    // set credentials
    await (await Page.email).setValue(user.email);
    await (await Page.password).setValue(user.password);
    await (await Page.firstname).setValue(user.firstname);
    await (await Page.lastname).setValue(user.lastname);
    // submit
    await Page.submit();
    // wait for a redirect
    await browser.pause(1000);
    expect(browser).toHaveUrlContaining("/app/addenterprise");
    // if all correct clear the session cookies
    await browser.deleteCookies(["auth._token.local", "auth._token_expiration.local"]);
  });
  it("register - wrong credentials", async () => {
    await Page.open();
    const user = mockUser();
    // skip registering the user
    // set credentials
    await (await Page.email).setValue(user.email);
    await (await Page.password).setValue(user.password);
    // submit
    await Page.submit();
    await browser.pause(500);
    // look for the snack bar
    const snackbar = await $(".snackbar .text");
    expect(snackbar).toHaveText("Coś poszło nie tak... Spróbuj ponownie później");
    // wait for a redirect
    expect(browser).toHaveUrlContaining("/app/login");
    // if all correct clear the session cookies
    await browser.deleteCookies(["auth._token.local", "auth._token_expiration.local"]);
  });
});
