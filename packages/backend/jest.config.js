module.exports = {
  testEnvironment: "node",
  transform: {
    "^.+\\.tsx?$": "ts-jest",
  },
  testRegex: "(/__tests__/.*|(\\.|/)(test|spec))\\.(jsx?|tsx?)$",
  testPathIgnorePatterns: ["/node_modules/"],
  moduleFileExtensions: ["ts", "tsx", "js", "jsx", "json"],
  preset: "ts-jest",
  moduleNameMapper: { "^@/(.*)$": "<rootDir>/src/$1" },
  verbose: false,
  setupFilesAfterEnv: ["<rootDir>/__tests__/setup/setupFilesAfterEnv.ts"],
};
