// class for craeting meta tag objects with properties Name, Hid, Content

export class ContentMetaTag {
  name: string;
  hid: string;
  content: string;
  constructor(name: string, content: string) {
    this.name = name;
    this.hid = name;
    this.content = content;
  }
}
