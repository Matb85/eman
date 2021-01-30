import { defaultlinks } from "./globals/globals";

const dirname = "/media";
export const state = () => ({
  media: ["google_business", "facebook", "instagram", "pinterest", "twitter"],
  subtreelinks: [
    defaultlinks(dirname),
    {
      name: "Media",
      sublinks: [
        {
          name: "Google Business",
          to: dirname + "/statistics",
          customIcon: "gbusiness",
          type: "is-gbusiness",
          class: "backlit",
        },
        {
          name: "Facebook",
          to: dirname + "/calendar",
          customIcon: "facebook",
          type: "is-facebook",
        },
        {
          name: "Instagram",
          to: dirname,
          customIcon: "instagram",
          type: "is-instagram",
          class: "instagram whiteicon",
        },
        {
          name: "Pinterest",
          to: dirname + "/statistics",
          customIcon: "pinterest",
          type: "is-pinterest",
        },
        {
          name: "Twitter",
          to: dirname + "/statistics",
          customIcon: "twitter",
          type: "is-twitter",
        },
      ],
    },
  ],
});

export const mutations = {};

export const actions = {};
