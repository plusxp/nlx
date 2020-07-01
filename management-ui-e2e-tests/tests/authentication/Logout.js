// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { Selector } from 'testcafe'
import { waitForReact } from 'testcafe-react-selectors'

import { getBaseUrl } from '../../utils'
import { adminUser } from '../roles'
import loginPage from './page-models/login'

const baseUrl = getBaseUrl()

fixture`Logout`.beforeEach(async (t) => {
  await t.useRole(adminUser).navigateTo(`${baseUrl}/`)
  await waitForReact()
})

test('Logging out should navigate to the login page', async (t) => {
  const userMenuButton = Selector('[aria-label="Account menu"]')
  const userMenuLogoutButton = Selector('button').withText('Uitloggen')

  await t.expect(userMenuButton.visible).ok()
  
  await t.click(userMenuButton)
  await t.expect(userMenuLogoutButton.visible).ok()
  
  await t.click(userMenuLogoutButton)
  await t.wait(500)
  await t.expect(loginPage.loginButton.visible).ok()
})