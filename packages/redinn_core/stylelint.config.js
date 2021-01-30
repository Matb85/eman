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
    // "font-family-no-missing-generic-family-keyword": [
    //   true,
    //   {
    //     ignoreFontFamilies: ["Lato Light"],
    //   },
    // ],
    "selector-pseudo-element-no-unknown": [
      true,
      {
        ignorePseudoElements: ["v-deep"],
      },
    ],
    "function-name-case": null,
    "declaration-block-single-line-max-declarations": 1,
  },
};
