module.exports = {
  root: true,
  parser: '@typescript-eslint/parser',
  parserOptions: {
    ecmaVersion: 2018,
    tsconfigRootDir: __dirname,
    project: ['./tsconfig.json'],
  },
  plugins: ['@typescript-eslint'],
  extends: [
    'eslint:recommended',
    'plugin:@typescript-eslint/recommended',
    'plugin:@typescript-eslint/recommended-requiring-type-checking',
    'plugin:node/recommended-module',
    'prettier',
  ],
  rules: {
    'node/no-missing-import': [
      'error',
      {
        resolvePaths: [
          './__tests__',
          './node_modules',
          './node_modules/@types',
          './node_modules/@types/node',
        ],
        tryExtensions: ['.ts', '.d.ts'],
      },
    ],
    quotes: ['error', 'single', { avoidEscape: true }],
    'sort-imports': 'error',
    'max-len': ['error', { code: 100 }],
  },
  overrides: [
    {
      files: ['**/*.test.ts'],
      plugins: ['jest'],
      extends: ['plugin:jest/recommended', 'plugin:jest/style'],
      env: {
        'jest/globals': true,
      },
      rules: {
        'jest/no-disabled-tests': 'error',
        'jest/no-focused-tests': 'error',
        'jest/no-identical-title': 'error',
        'jest/consistent-test-it': 'error',
        'jest/valid-expect': [
          'error',
          {
            alwaysAwait: true,
          },
        ],
      },
      settings: {
        jest: { version: 26 },
      },
    },
  ],
  ignorePatterns: ['*.js'],
}
