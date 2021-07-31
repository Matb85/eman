import Page from "./main.page";

class LoginPage extends Page {
  get email() {
    return $("[data-testid='register-email']");
  }
  get password() {
    return $("[data-testid='register-password']");
  }
  get firstname() {
    return $("[data-testid='register-firstName']");
  }
  get lastname() {
    return $("[data-testid='register-lastName']");
  }
  get submitBtn() {
    return $("[data-testid='register-submit']");
  }
  open() {
    return super.open("/register");
  }
  async submit() {
    return (await this.submitBtn).click();
  }
}

export default new LoginPage();
