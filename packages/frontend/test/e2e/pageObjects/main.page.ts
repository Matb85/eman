export default class Page {
  static open(path = "/") {
    browser.url("/app" + path);
  }
}
