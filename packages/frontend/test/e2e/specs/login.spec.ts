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
    expect(response.status).toBe(200);
    expect(response.data.message).toBe("registration successful");
    // set credentials
    console.log(Page.email);
    await (await Page.email).setValue(user.email);
    await (await Page.password).setValue(user.password);
    // submit
    await Page.submit();
    await browser.pause(1000);
    // wait for a redirect
    expect(browser).toHaveUrlContaining("/app/addenterprise");
  });
});
