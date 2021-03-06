// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { act, fireEvent } from '@testing-library/react'
import { Router } from 'react-router-dom'
import { createMemoryHistory } from 'history'
import { renderWithProviders } from '../../../test-utils'
import { RootStore, StoreProvider } from '../../../stores'
import { ManagementApi } from '../../../api'
import AddServicePage from './index'

jest.mock('../../../components/PageTemplate/OrganizationInwayCheck', () => () =>
  null,
)

// eslint-disable-next-line react/prop-types
jest.mock('../../../components/ServiceForm', () => ({ onSubmitHandler }) => (
  <form onSubmit={() => onSubmitHandler({ foo: 'bar' })} data-testid="form">
    <button type="submit" />
  </form>
))

describe('the AddServicePage', () => {
  afterEach(() => {
    jest.resetModules()
  })

  it('on initialization', () => {
    const managementApiClient = new ManagementApi()
    const store = new RootStore({ managementApiClient })
    const { getByTestId, queryByRole, getByLabelText } = renderWithProviders(
      <Router history={createMemoryHistory()}>
        <StoreProvider rootStore={store}>
          <AddServicePage />
        </StoreProvider>
      </Router>,
    )

    const linkBack = getByLabelText(/Back/)
    expect(linkBack.getAttribute('href')).toBe('/services')
    expect(getByTestId('form')).toBeTruthy()
    expect(queryByRole('dialog')).toBeNull()
  })

  it('successfully submitting the form', async () => {
    const managementApiClient = new ManagementApi()
    managementApiClient.managementCreateService = jest.fn().mockResolvedValue({
      name: 'my-service',
    })

    const rootStore = new RootStore({
      managementApiClient,
    })

    const history = createMemoryHistory()
    const { findByTestId } = renderWithProviders(
      <Router history={history}>
        <StoreProvider rootStore={rootStore}>
          <AddServicePage />
        </StoreProvider>
      </Router>,
    )

    const addComponentForm = await findByTestId('form')
    await act(async () => {
      fireEvent.submit(addComponentForm)
    })

    expect(managementApiClient.managementCreateService).toHaveBeenCalledTimes(1)
    expect(history.location.pathname).toEqual('/services/my-service')
    expect(history.location.search).toEqual('?lastAction=added')
  })

  it('re-submitting the form when the previous submission went wrong', async () => {
    const managementApiClient = new ManagementApi()
    managementApiClient.managementCreateService = jest
      .fn()
      .mockResolvedValue({ name: 'my-service' })
      .mockRejectedValueOnce(new Error('arbitrary error'))

    const rootStore = new RootStore({
      managementApiClient,
    })

    const history = createMemoryHistory()
    const { findByTestId, queryByRole } = renderWithProviders(
      <Router history={history}>
        <StoreProvider rootStore={rootStore}>
          <AddServicePage />
        </StoreProvider>
      </Router>,
    )

    const addComponentForm = await findByTestId('form')

    await act(async () => {
      await fireEvent.submit(addComponentForm)
    })

    expect(managementApiClient.managementCreateService).toHaveBeenCalledTimes(1)
    expect(queryByRole('alert')).toBeTruthy()
    expect(queryByRole('alert').textContent).toBe(
      'Failed adding servicearbitrary error',
    )

    await act(async () => {
      await fireEvent.submit(addComponentForm)
    })

    expect(await queryByRole('alert')).toBeTruthy()

    expect(managementApiClient.managementCreateService).toHaveBeenCalledTimes(2)
    expect(history.location.pathname).toEqual('/services/my-service')
    expect(history.location.search).toEqual('?lastAction=added')
  })

  it('submitting when the HTTP response is not ok', async () => {
    const managementApiClient = new ManagementApi()
    managementApiClient.managementCreateService = jest
      .fn()
      .mockRejectedValue(new Error('arbitrary error'))

    const rootStore = new RootStore({
      managementApiClient,
    })

    const history = createMemoryHistory()
    const { findByTestId, queryByRole } = renderWithProviders(
      <Router history={history}>
        <StoreProvider rootStore={rootStore}>
          <AddServicePage />
        </StoreProvider>
      </Router>,
    )

    const addComponentForm = await findByTestId('form')

    await act(async () => {
      await fireEvent.submit(addComponentForm)
    })

    expect(queryByRole('alert')).toBeTruthy()
    expect(queryByRole('alert').textContent).toBe(
      'Failed adding servicearbitrary error',
    )
  })
})
