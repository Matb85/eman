module.exports = {
  root: true,
  env: {
    browser: true,
    node: true,
  },
  extends: [
    "@nuxtjs/eslint-config-typescript",
    "prettier",
    "prettier/vue",
    "plugin:prettier/recommended",
    "plugin:wdio/recommended",
    "plugin:nuxt/recommended",
  ],
  plugins: ["prettier", "wdio"],
  // add your custom rules here
  rules: {
    noImplicitThis: true,
    strict: true,
    "prettier/prettier": ["warn", { trailingComma: "es5", printWidth: 120 }],
    "max-len": [1, 120],
  },
};
