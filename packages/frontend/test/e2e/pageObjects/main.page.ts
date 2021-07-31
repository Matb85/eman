export default class Page {
  async open(path = "/") {
    await browser.url("/app" + path);
  }
}
