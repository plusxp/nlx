// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import deferredPromise from '../test-utils/deferred-promise'
import DirectoryStore from './DirectoryStore'
import { mockDirectoryServiceModel } from './DirectoryStore.mock'

jest.mock('../models/DirectoryServiceModel', () => ({
  createDirectoryService: (...args) => mockDirectoryServiceModel(...args),
}))

let rootStore
let directoryRepository

beforeEach(() => {
  rootStore = {}
  directoryRepository = {}
})

test('fetching directory services', async () => {
  const request = deferredPromise()
  directoryRepository = {
    getAll: jest.fn().mockReturnValue(request),
  }

  const serviceList = [
    { organizationName: 'Org A', serviceName: 'Service A' },
    { organizationName: 'Org B', serviceName: 'Service B' },
  ]

  const directoryStore = new DirectoryStore({
    rootStore,
    directoryRepository,
  })

  expect(directoryStore.services).toEqual([])
  expect(directoryStore.isInitiallyFetched).toBe(false)

  directoryStore.fetchServices()

  expect(directoryStore.isInitiallyFetched).toBe(false)
  expect(directoryRepository.getAll).toHaveBeenCalled()

  await request.resolve(serviceList)

  await expect(directoryStore.isInitiallyFetched).toBe(true)
  expect(directoryStore.services).toHaveLength(2)
  expect(directoryStore.services).not.toBe([])
})

test('handle error while fetching directory services', async () => {
  const errorSpy = jest.spyOn(console, 'error').mockImplementation(() => {})

  const request = deferredPromise()
  directoryRepository = {
    getAll: jest.fn(() => request),
  }

  const directoryStore = new DirectoryStore({
    rootStore,
    directoryRepository,
  })

  expect(directoryStore.services).toEqual([])

  directoryStore.fetchServices()

  expect(directoryStore.isInitiallyFetched).toBe(false)
  expect(directoryRepository.getAll).toHaveBeenCalled()

  await request.reject('some error')

  expect(directoryStore.error).toEqual('some error')
  expect(directoryStore.services).toEqual([])
  expect(directoryStore.isInitiallyFetched).toBe(true)

  errorSpy.mockRestore()
})

test('selecting a directory service', () => {
  const mockDirectoryServiceModelA = mockDirectoryServiceModel({
    organizationName: 'Org A',
    serviceName: 'Service A',
    state: 'state-a',
  })
  const mockDirectoryServiceModelB = mockDirectoryServiceModel({
    organizationName: 'Org B',
    serviceName: 'Service B',
    state: 'state-b',
  })
  const serviceList = [mockDirectoryServiceModelA, mockDirectoryServiceModelB]

  const directoryStore = new DirectoryStore({
    rootStore,
    directoryRepository,
  })
  directoryStore.services = serviceList

  const selectedService = directoryStore.selectService({
    organizationName: 'Org A',
    serviceName: 'Service A',
  })

  expect(mockDirectoryServiceModelA.fetch).toHaveBeenCalled()
  expect(mockDirectoryServiceModelB.fetch).not.toHaveBeenCalled()
  expect(selectedService.state).toBe('state-a')
})