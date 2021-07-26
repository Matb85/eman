import Page from "../pageObjects/main.page";

describe("Example test", () => {
  it("should open correct app", () => {
    Page.open("/login");
    expect(browser).toHaveTitle("Redinn");
  });
});
