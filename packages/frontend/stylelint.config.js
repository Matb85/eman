module.exports = {
  extends: ["stylelint-config-standard", "stylelint-config-prettier"],
  // add your custom config here
  // https://stylelint.io/user-guide/configuration
  rules: {
    "at-rule-no-unknown": [
      true,
      {
        ignoreAtRules: ["function", "if", "each", "include", "mixin", "for", "extend"],
      },
    ],
    "selector-pseudo-element-no-unknown": [true, { ignorePseudoElements: ["v-deep"] }],
    "function-name-case": null,
    "declaration-block-single-line-max-declarations": 1,
    "no-invalid-position-at-import-rule": null,
  },
};
