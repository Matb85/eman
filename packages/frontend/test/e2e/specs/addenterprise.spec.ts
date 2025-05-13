import Page from "../pageObjects/addenterprise.page";
import { RegisterUser, Login, Logout } from "../utils/auth";
import { mockEnterpise } from "../utils/";
import path from "path";

describe("/addenterprise tab - a11y", () => {
  it("should have a proper title", async () => {
    await Page.open();
    expect(browser).toHaveTitle("Eman");
  });
});

describe("/addenterprise tab - input form", () => {
  it("add an enterprise - OK", async () => {
    const user = await RegisterUser();
    await Login(user);

    await Page.open();
    const enterprise = mockEnterpise();
    // set necessary data
    await (await Page.name).setValue(enterprise.name);
    await (await Page.country).setValue(enterprise.address.country);
    await (await Page.zipcode).setValue(enterprise.address.zipcode);
    await (await Page.street).setValue(enterprise.address.street);
    await (await Page.city).setValue(enterprise.address.city);
    // upload a logo
    const filePath = path.join(process.env.BASE_DIR!, "/assets/image.jpg");
    await (await Page.upload).setValue(filePath);
    // submit
    await Page.submit();
    // wait for a redirect
    await browser.pause(1000);
    expect(browser).toHaveUrlContaining("/app/0/media");
    // if all correct clear the session cookies
    await Logout();
  });
});
