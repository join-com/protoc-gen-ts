module.exports = {
  preset: "ts-jest",
  testEnvironment: "node",
  testPathIgnorePatterns: ["generated"],
  setupFilesAfterEnv: ["jest-extended"],
  watchPlugins: ["jest-watch-typeahead/filename"],
  globals: {
    "ts-jest": {
      diagnostics: {
        warnOnly: true
      }
    }
  }
};
