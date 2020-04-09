// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

import { Selector, Role, ClientFunction } from 'testcafe'

const getBaseUrl = require('../getBaseUrl')
const baseUrl = getBaseUrl()

const getLocation  = ClientFunction(() => document.location.href.toString());

export const adminUser = Role(`${baseUrl}/login`, async t => {
  const managementLoginButton = Selector('#login');
  const dexLoginText = Selector('#login');
  const dexPasswordText = Selector('#password');
  const dexSubmitLoginButton = Selector('#submit-login');
  const dexGrantAccessButton = Selector('button[type="submit"]');

  await t
    .setTestSpeed(0.5)
    .click(managementLoginButton)
    .typeText(dexLoginText, 'admin@example.com')
    .typeText(dexPasswordText, 'password')
    .click(dexSubmitLoginButton)
    .click(dexGrantAccessButton)
    .expect(getLocation()).notContains('auth/local')
    .navigateTo(getBaseUrl())
});
