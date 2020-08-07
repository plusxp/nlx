// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { observable } from 'mobx'
import { Route, StaticRouter as Router } from 'react-router-dom'

import { renderWithProviders } from '../../../test-utils'
import DirectoryDetailPage from './index'

jest.mock('./components/DirectoryDetailView', () => ({ service }) => (
  <div data-testid="directory-service-details" />
))

let service

beforeEach(() => {
  service = observable({
    id: 'Test Organization/Test Service',
    organizationName: 'Test Organization',
    serviceName: 'Test Service',
    state: 'degraded',
    apiSpecificationType: 'API',
    latestAccessRequest: null,
    requestAccess: jest.fn(),
    fetch: jest.fn(),

    isLoading: false,
  })
})

test('display directory service details', () => {
  const { getByTestId, getByText } = renderWithProviders(
    // Router & Route still required for hooks
    // Note not they, but the service data is tested
    <Router location="/directory/organization/service">
      <Route path="/directory/:organizationName/:serviceName">
        <DirectoryDetailPage service={service} />
      </Route>
    </Router>,
  )

  expect(getByText('Test Organization')).toBeInTheDocument()
  expect(getByText('Test Service')).toBeInTheDocument()
  expect(getByText('state-degraded.svg')).toBeInTheDocument()
  expect(getByTestId('directory-service-details')).toBeInTheDocument()
})

test('fetch latest state on load and display changes', () => {
  service.fetch = jest.fn(() => (service.state = 'up'))

  const { queryByText } = renderWithProviders(
    <Router location="/directory/organization/service">
      <Route path="/directory/:organizationName/:serviceName">
        <DirectoryDetailPage service={service} />
      </Route>
    </Router>,
  )

  expect(service.fetch).toHaveBeenCalled()
  expect(queryByText('state-degraded.svg')).not.toBeInTheDocument()
  expect(queryByText('state-up.svg')).toBeInTheDocument()
})

test('service does not exist', () => {
  const { getByTestId, getByText, queryByText } = renderWithProviders(
    <Router location="/directory/organization/service">
      <Route path="/directory/:organizationName/:serviceName">
        <DirectoryDetailPage service={undefined} />
      </Route>
    </Router>,
  )

  const message = getByTestId('error-message')
  expect(message).toBeInTheDocument()
  expect(message.textContent).toBe('Failed to load the service.')

  expect(getByText('service')).toBeInTheDocument()
  expect(queryByText('organization')).toBeNull()

  const closeButton = getByTestId('close-button')
  expect(closeButton).toBeInTheDocument()
})