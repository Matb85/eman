import Page from "./main.page";

class AddEnterprisePage extends Page {
  get name() {
    return $("[data-testid='add-e-name']");
  }
  get city() {
    return $("[data-testid='add-e-city']");
  }
  get zipcode() {
    return $("[data-testid='add-e-zipcode']");
  }
  get country() {
    return $("[data-testid='add-e-country']");
  }
  get street() {
    return $("[data-testid='add-e-street']");
  }
  get upload() {
    return $("[data-testid='add-e-upload']");
  }
  get submitBtn() {
    return $("[data-testid='add-e-submit']");
  }
  open() {
    return super.open("/addenterprise");
  }
  async submit() {
    await (await this.submitBtn).click();
  }
}

export default new AddEnterprisePage();
