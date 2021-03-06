// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
module.exports = {
  extends: [
    'standard',
    'plugin:prettier/recommended',
    '@commonground/eslint-config/rules/reactAppGenerics',
    '@commonground/eslint-config/rules/generic',
    '@commonground/eslint-config/rules/prettier',
    '@commonground/eslint-config/rules/header',
    '@commonground/eslint-config/rules/import',
  ],

  globals: {
    cy: 'readonly',
    Cypress: 'readonly',
    it: 'readonly',
    describe: 'readonly',
    beforeEach: 'readonly',
  },
}
