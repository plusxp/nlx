// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { checkPropTypes } from 'prop-types'

import deferredPromise from '../test-utils/deferred-promise'
import OutgoingAccessRequestModel, {
  createAccessRequestInstance,
  outgoingAccessRequestPropTypes,
} from './OutgoingAccessRequestModel'

jest.mock('../domain/access-request-repository', (obj) => obj)

let serviceData
let accessRequestJson
let domain

beforeEach(() => {
  serviceData = {
    organizationName: 'Organization',
    serviceName: 'Service',
  }

  accessRequestJson = {
    ...serviceData,
    id: 'abcd',
    createdAt: 'datetime1',
    updatedAt: 'datetime2',
  }

  domain = {
    requestAccess: jest.fn(),
  }
})

test('model implements proptypes', () => {
  const errorSpy = jest.spyOn(console, 'error').mockImplementation(() => {})
  const accessRequest = new OutgoingAccessRequestModel({
    json: accessRequestJson,
    domain,
  })

  checkPropTypes(
    outgoingAccessRequestPropTypes,
    accessRequest,
    'prop',
    'OutgoingAccessRequestModel',
  )

  expect(errorSpy).not.toHaveBeenCalled()
  errorSpy.mockRestore()
})

test('createAccessRequestInstance creates an instance', () => {
  expect(createAccessRequestInstance(serviceData)).toBeInstanceOf(
    OutgoingAccessRequestModel,
  )
})

test('sending a request', async () => {
  const request = deferredPromise()
  domain = {
    requestAccess: jest.fn(() => request),
  }

  const accessRequest = new OutgoingAccessRequestModel({
    json: serviceData,
    domain,
  })

  expect(accessRequest.state).toBe('')

  accessRequest.send()

  expect(accessRequest.isLoading).toBe(true)
  expect(accessRequest.state).toBe('CREATED')
  expect(domain.requestAccess).toHaveBeenCalled()

  await request.resolve(accessRequestJson)

  expect(accessRequest.isLoading).toBe(false)
  expect(accessRequest.id).toBe('abcd')
})

test('update should ignore properties that do not belong on object', () => {
  const accessRequest = new OutgoingAccessRequestModel({
    json: accessRequestJson,
    domain,
  })

  accessRequest.update({ yada: 'blada' })

  expect('yada' in accessRequest).toBe(false)
})