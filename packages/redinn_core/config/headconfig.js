import { ContentMetaTag } from "../helpers/helperclasses.ts";

const TITLE = "Redinn";
const DESC = "Redinn - The ultimate platform for managing your business";
const IMG = "/icon.png";
const ALT = "Redinn logo";
export default {
  title: TITLE,
  meta: [
    { charset: "utf-8" },
    { name: "viewport", content: "width=device-width, initial-scale=1" },
    new ContentMetaTag("description", DESC),
    new ContentMetaTag("twitter:title", TITLE),
    new ContentMetaTag("twitter:description", DESC),
    new ContentMetaTag("twitter:image", IMG),
    new ContentMetaTag("twitter:image:alt", ALT),
    new ContentMetaTag("og:title", TITLE),
    new ContentMetaTag("og:description", DESC),
    new ContentMetaTag("og:image", IMG),
    new ContentMetaTag("og:image:alt", ALT),
    new ContentMetaTag("og:image:secure_url", IMG),
    new ContentMetaTag("og:type", "website"),
    new ContentMetaTag("og:site_name", "Redinn"),
    new ContentMetaTag(
      "og:url",
      process.env.BASE_URL || "http://localhost:3000"
    ),
  ],
  link: [{ rel: "icon", type: "image/x-icon", href: "/favicon.ico" }],
};
