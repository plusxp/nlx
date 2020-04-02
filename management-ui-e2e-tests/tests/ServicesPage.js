// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

import { Selector } from 'testcafe'
import { axeCheck, createReport } from 'axe-testcafe'
import { adminUser } from "./roles";

const getBaseUrl = require('../getBaseUrl')
const baseUrl = getBaseUrl()

fixture `Services page`
  .page(`${baseUrl}/services`)
  .beforeEach(async (t) => {
    await t.useRole(adminUser)
  })

test('Automated accessibility testing', async t => {
  const { violations } = await axeCheck(t)
  await t.expect(violations.length === 0).ok(createReport(violations));
})

test('Page title is visible', async t => {
  const pageTitle = Selector('h1')
  const servicesList = Selector('[data-testid="services-list"]');

  await t
    .expect(pageTitle.visible).ok()
    .expect(pageTitle.innerText).eql('Services')
})

test('Service details are displayed', async t => {
  const servicesList = Selector('[data-testid="services-list"]');
  const kentekenService = servicesList.find('tbody tr').nth(0)
  const kentekenServiceColumns = kentekenService.find('td')

  const nameCell = kentekenServiceColumns.nth(0)
  const accessCell = kentekenServiceColumns.nth(1)

  await t
    .expect(servicesList.visible).ok()
    .expect(servicesList.find('tbody tr').count).eql(1)

    .expect(nameCell.textContent).eql('kentekenregister')
    .expect(accessCell.textContent).eql('Open')
})