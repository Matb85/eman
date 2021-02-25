import pl from "./pl.js";
import en from "./en.js";
export default {
  locales: [
    { code: "en", iso: "en-US", file: "en.js" },
    { code: "pl", iso: "pl-PL", file: "pl.js" },
  ],
  defaultLocale: "pl",
  vueI18n: {
    fallbackLocale: "en",
    messages: {
      en,
      pl,
    },
  },
  detectBrowserLanguage: {
    useCookie: true,
    cookieKey: "i18n_redirected",
    // onlyOnRoot: true,
  },
};
