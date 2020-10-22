// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { observable } from 'mobx'

export const mockServicesStore = ({
  services = [],
  isInitiallyFetched = true,
  error = '',
  fetchAll = jest.fn(),
  selectService = jest.fn(),
  create = jest.fn(),
  removeService = jest.fn(),
}) =>
  observable({
    servicesStore: {
      services,
      isInitiallyFetched,
      error,
      fetchAll,
      selectService,
      removeService,
      create,
    },
  })

export const mockServiceModel = (service, fetch = jest.fn()) => ({
  ...service,
  fetch,
})
