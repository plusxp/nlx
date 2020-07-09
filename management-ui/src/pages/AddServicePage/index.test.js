// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { fireEvent, act } from '@testing-library/react'
import { MemoryRouter, Router } from 'react-router-dom'

import { createMemoryHistory } from 'history'
import UserContext from '../../user-context'
import { renderWithProviders } from '../../test-utils'
import AddServicePage from './index'

jest.mock('../../components/ServiceForm', () => ({ onSubmitHandler }) => (
  <form onSubmit={() => onSubmitHandler({ foo: 'bar' })} data-testid="form">
    <button type="submit" />
  </form>
))

describe('the AddServicePage', () => {
  afterEach(() => {
    jest.resetModules()
  })

  it('on initialization', () => {
    const userContext = { user: { id: '42' } }
    const { getByTestId, queryByRole, getByLabelText } = renderWithProviders(
      <MemoryRouter>
        <UserContext.Provider value={userContext}>
          <AddServicePage createHandler={() => {}} />
        </UserContext.Provider>
      </MemoryRouter>,
    )

    const linkBack = getByLabelText(/Back/)
    expect(linkBack.getAttribute('href')).toBe('/services')
    expect(getByTestId('form')).toBeTruthy()
    expect(queryByRole('dialog')).toBeNull()
  })

  it('successfully submitting the form', async () => {
    const history = createMemoryHistory()
    const createHandler = jest.fn().mockResolvedValue({
      name: 'my-service',
    })
    const { findByTestId } = renderWithProviders(
      <Router history={history}>
        <AddServicePage createHandler={createHandler} />
      </Router>,
    )

    const addComponentForm = await findByTestId('form')
    await act(async () => {
      fireEvent.submit(addComponentForm)
    })

    expect(history.location.pathname).toEqual('/services/my-service')
    expect(history.location.search).toEqual('?lastAction=added')
  })

  it('re-submitting the form when the previous submission went wrong', async () => {
    const createHandler = jest
      .fn()
      .mockResolvedValue({ name: 'my-service' })
      .mockRejectedValueOnce(new Error('arbitrary error'))

    const history = createMemoryHistory()
    const { findByTestId, queryByRole } = renderWithProviders(
      <Router history={history}>
        <AddServicePage createHandler={createHandler} />
      </Router>,
    )

    const addComponentForm = await findByTestId('form')

    await act(async () => {
      await fireEvent.submit(addComponentForm)
    })

    expect(createHandler).toHaveBeenCalledTimes(1)
    expect(queryByRole('alert')).toBeTruthy()
    expect(queryByRole('alert').textContent).toBe(
      'Failed adding servicearbitrary error',
    )

    await act(async () => {
      await fireEvent.submit(addComponentForm)
    })

    expect(await queryByRole('alert')).toBeTruthy()

    expect(createHandler).toHaveBeenCalledTimes(2)
    expect(history.location.pathname).toEqual('/services/my-service')
    expect(history.location.search).toEqual('?lastAction=added')
  })

  it('submitting when the HTTP response is not ok', async () => {
    const createHandler = jest
      .fn()
      .mockRejectedValue(new Error('arbitrary error'))

    const { findByTestId, queryByRole } = renderWithProviders(
      <MemoryRouter>
        <AddServicePage createHandler={createHandler} />
      </MemoryRouter>,
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
