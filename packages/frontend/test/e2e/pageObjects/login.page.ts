import Page from "./main.page";

class LoginPage extends Page {
  get email() {
    return $("[data-testid='login-email']");
  }
  get password() {
    return $("[data-testid='login-password']");
  }
  get submitBtn() {
    return $("[data-testid='login-submit']");
  }
  open() {
    return super.open("/login");
  }
  async submit() {
    await (await this.submitBtn).click();
  }
}

export default new LoginPage();
