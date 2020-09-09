// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { checkPropTypes } from 'prop-types'

import deferredPromise from '../test-utils/deferred-promise'
import InwayModel, { createInway, inwayModelPropTypes } from './InwayModel'

let store
let inway

beforeEach(() => {
  store = {}
  inway = {
    name: 'Service',
  }
})

test('createInway returns an instance', () => {
  const inwayModel = createInway({ store, inway })
  expect(inwayModel).toBeInstanceOf(InwayModel)
})

test('model implements proptypes', () => {
  const errorSpy = jest.spyOn(console, 'error').mockImplementation(() => {})
  const inwayModel = new InwayModel({ store, inway })

  checkPropTypes(inwayModelPropTypes, inwayModel, 'prop', 'InwayModel')

  expect(errorSpy).not.toHaveBeenCalled()
  errorSpy.mockRestore()
})

test('fetches data', async () => {
  const request = deferredPromise()
  store = {
    domain: {
      getByName: jest.fn(() => request),
    },
  }

  const inwayModel = new InwayModel({ store, inway })

  inwayModel.fetch()

  expect(store.domain.getByName).toHaveBeenCalled()

  await request.resolve(inway)

  expect(inwayModel).toBeInstanceOf(InwayModel)
})