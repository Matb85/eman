import Page from "../pageObjects/login.page";
import { mockUser, axiosApi } from "../utils/index";

describe("/login tab - a11y", () => {
  it("should have a proper title", async () => {
    await Page.open();
    expect(browser).toHaveTitle("Redinn");
  });
});

describe("/login tab - input form", () => {
  it("login - OK", async () => {
    await Page.open();
    const user = mockUser();
    // register the user
    const response = await axiosApi.post("/auth/register", user).catch((err) => err.response);
    expect(response.data.message).toBe("registration successful");
    expect(response.status).toBe(200);
    // set credentials
    await (await Page.email).setValue(user.email);
    await (await Page.password).setValue(user.password);
    // submit
    await Page.submit();
    // wait for a redirect
    await browser.pause(1000);
    expect(browser).toHaveUrlContaining("/app/addenterprise");
    // if all correct clear the session cookies
    await browser.deleteCookies(["auth._token.local", "auth._token_expiration.local"]);
  });
  it("login - wrong credentials", async () => {
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
