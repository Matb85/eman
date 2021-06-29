module.exports = {
  testEnvironment: "node",
  transform: {
    "^.+\\.tsx?$": "ts-jest",
  },
  testRegex: "(/__tests__/.*|(\\.|/)(test|spec))\\.(jsx?|tsx?)$",
  testPathIgnorePatterns: ["/node_modules/", "/__tests__/setup", "/dist/"],
  moduleFileExtensions: ["ts", "tsx", "js", "jsx", "json"],
  preset: "ts-jest",
  moduleNameMapper: { "^@/(.*)$": "<rootDir>/src/$1" },
  verbose: false,

  globalSetup: "<rootDir>/__tests__/setup/globalSetup.ts",
  globalTeardown: "<rootDir>/__tests__/setup/globalTeardown.ts",
  setupFilesAfterEnv: ["<rootDir>/__tests__/setup/setupFilesAfterEnv.ts"],
};
