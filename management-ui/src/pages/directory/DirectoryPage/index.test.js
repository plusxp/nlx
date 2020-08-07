// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React from 'react'
import { MemoryRouter as Router } from 'react-router-dom'
import { observable } from 'mobx'

import { act, renderWithProviders } from '../../../test-utils'
import { StoreProvider } from '../../../stores'
import { UserContextProvider } from '../../../user-context'
import DirectoryPage from './index'

// Ignore this deeply nested component which has a separate request flow
jest.mock('../../../components/OrganizationName', () => () => null)

// Simplify showing of the services. We'll only require the serviceName.
jest.mock('./components/DirectoryPageView', () => ({ services }) => {
  return (
    <div data-testid="mock-directory-services">
      {services.map((o, i) => (
        <span key={i} data-testid={`mock-directory-service-${i}`}>
          {o.serviceName}
        </span>
      ))}
    </div>
  )
})

const renderDirectory = (store) =>
  renderWithProviders(
    <StoreProvider store={store}>
      <UserContextProvider user={{}}>
        <Router>
          <DirectoryPage />
        </Router>
      </UserContextProvider>
    </StoreProvider>,
  )

test('listing all services', () => {
  const store = observable({
    directoryStore: {
      services: [],
      isLoading: true,
      error: '',
      fetchServices: jest.fn(),
      selectService: jest.fn(),
    },
  })

  const { getByRole, getByTestId } = renderDirectory(store)

  expect(store.directoryStore.fetchServices).toHaveBeenCalled()
  expect(getByRole('progressbar')).toBeInTheDocument()
  expect(() => getByTestId('mock-directory-services')).toThrow()

  act(() => {
    store.directoryStore.services = [{ serviceName: 'Test Service' }]
    store.directoryStore.isLoading = false
  })

  expect(getByTestId('mock-directory-services')).toBeInTheDocument()
  expect(() => getByRole('progressbar')).toThrow()
  expect(getByTestId('mock-directory-service-0')).toHaveTextContent(
    'Test Service',
  )
})

test('no services', () => {
  const store = observable({
    directoryStore: {
      services: [],
      isLoading: false,
      error: '',
      fetchServices: jest.fn(),
      selectService: jest.fn(),
    },
  })

  const { getByTestId } = renderDirectory(store)

  expect(getByTestId('mock-directory-services')).toBeInTheDocument()
  expect(() => getByTestId('mock-directory-service-0')).toThrow()
})

test('failed to load services', () => {
  const store = observable({
    directoryStore: {
      services: [],
      isLoading: false,
      error: 'There is an error',
      fetchServices: jest.fn(),
      selectService: jest.fn(),
    },
  })

  const { getByTestId } = renderDirectory(store)

  expect(getByTestId('error-message')).toHaveTextContent(
    /^Failed to load the directory\.$/,
  )
  expect(() => getByTestId('mock-directory-services')).toThrow()
})