// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { configure } from 'mobx'
import { RootStore } from '../index'
import OutgoingAccessRequestModel, {
  ACCESS_REQUEST_STATES,
} from './OutgoingAccessRequestModel'
import DirectoryServiceModel from './DirectoryServiceModel'
import AccessProofModel from './AccessProofModel'

test('initializing the model', () => {
  const errorSpy = jest.spyOn(console, 'error').mockImplementation(() => {})
  const directoryService = new DirectoryServiceModel({
    directoryServicesStore: {},
    serviceData: {
      organizationName: 'Organization',
      serviceName: 'Service',
      state: 'up',
      apiSpecificationType: 'API',
    },
  })

  expect(errorSpy).not.toHaveBeenCalled()
  errorSpy.mockRestore()

  expect(directoryService.latestAccessRequest).toBeNull()
  expect(directoryService.latestAccessProof).toBeNull()

  const latestAccessRequest = new OutgoingAccessRequestModel({
    outgoingAccessRequestStore: {},
    accessRequestData: {
      state: ACCESS_REQUEST_STATES.CREATED,
    },
  })

  const latestAccessProof = new AccessProofModel({
    accessProofData: {
      id: 'abc',
    },
  })

  directoryService.update({
    serviceData: { state: 'down' },
    latestAccessRequest,
    latestAccessProof,
  })

  expect(directoryService.state).toBe('down')
  expect(directoryService.latestAccessRequest).toEqual(latestAccessRequest)
  expect(directoryService.latestAccessProof).toEqual(latestAccessProof)
})

test('updating the model with an invalid latest access request and access proof', () => {
  const directoryService = new DirectoryServiceModel({
    directoryServicesStore: {},
    serviceData: {
      organizationName: 'Organization',
      serviceName: 'Service',
      state: 'up',
      apiSpecificationType: 'API',
    },
  })

  expect(() =>
    directoryService.update({
      latestAccessRequest: 'invalid',
    }),
  ).toThrow()

  expect(() =>
    directoryService.update({
      latestAccessProof: 'invalid',
    }),
  ).toThrow()
})

test('(re-)fetching the model', async () => {
  configure({ safeDescriptors: false })
  const rootStore = new RootStore({})

  jest.spyOn(rootStore.directoryServicesStore, 'fetch').mockResolvedValue({})

  const directoryService = new DirectoryServiceModel({
    directoryServicesStore: rootStore.directoryServicesStore,
    serviceData: {
      organizationName: 'Organization',
      serviceName: 'Service',
      state: 'up',
      apiSpecificationType: 'API',
    },
  })

  await directoryService.fetch()

  expect(rootStore.directoryServicesStore.fetch).toHaveBeenCalledWith(
    directoryService,
  )
})

describe('requesting access to a service', () => {
  configure({ safeDescriptors: false })
  it('should request access via the directory service store', async () => {
    const rootStore = new RootStore({})

    const directoryService = new DirectoryServiceModel({
      directoryServicesStore: rootStore.directoryServicesStore,
      serviceData: {},
    })

    jest
      .spyOn(rootStore.directoryServicesStore, 'requestAccess')
      .mockResolvedValue(null)

    await directoryService.requestAccess()

    expect(rootStore.directoryServicesStore.requestAccess).toHaveBeenCalledWith(
      directoryService,
    )
  })
})
