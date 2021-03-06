// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { fireEvent, act } from '@testing-library/react'
import { MemoryRouter } from 'react-router-dom'

import UserContext from '../../../user-context'
import { renderWithProviders } from '../../../test-utils'
import InsightSettings from './index'

// eslint-disable-next-line react/prop-types
jest.mock('./Form', () => ({ onSubmitHandler }) => (
  <form onSubmit={() => onSubmitHandler({ foo: 'bar' })} data-testid="form">
    <button type="submit" />
  </form>
))

describe('the Insight settings section', () => {
  afterEach(() => {
    jest.resetModules()
  })

  it('on initialization', async () => {
    const getSettingsHandler = jest.fn().mockResolvedValue({})
    const userContext = { user: { id: '42' } }
    const { findByTestId, queryByTestId } = renderWithProviders(
      <MemoryRouter>
        <UserContext.Provider value={userContext}>
          <InsightSettings
            getSettings={getSettingsHandler}
            updateHandler={() => {}}
          />
        </UserContext.Provider>
      </MemoryRouter>,
    )

    const formElement = await findByTestId('form')

    expect(formElement).toBeTruthy()
    expect(queryByTestId('error-message')).toBeNull()
  })

  it('successfully submitting the form', async () => {
    const updateHandler = jest.fn().mockResolvedValue(null)
    const getSettingsHandler = jest.fn().mockResolvedValue({})
    const userContext = { user: { id: '42' } }
    const { findByTestId } = renderWithProviders(
      <MemoryRouter>
        <UserContext.Provider value={userContext}>
          <InsightSettings
            updateHandler={updateHandler}
            getSettings={getSettingsHandler}
          />
        </UserContext.Provider>
      </MemoryRouter>,
    )

    const settingsForm = await findByTestId('form')
    await act(async () => {
      fireEvent.submit(settingsForm)
    })

    expect(updateHandler).toHaveBeenCalledWith({
      foo: 'bar',
    })
  })

  it('re-submitting the form when the previous submission went wrong', async () => {
    const updateHandler = jest
      .fn()
      .mockResolvedValue(null)
      .mockRejectedValueOnce(new Error('arbitrary error'))

    const getSettingsHandler = jest.fn().mockResolvedValue({})
    const userContext = { user: { id: '42' } }

    const { findByTestId, getByRole } = renderWithProviders(
      <MemoryRouter>
        <UserContext.Provider value={userContext}>
          <InsightSettings
            updateHandler={updateHandler}
            getSettings={getSettingsHandler}
          />
        </UserContext.Provider>
      </MemoryRouter>,
    )

    const settingsForm = await findByTestId('form')

    await act(async () => {
      await fireEvent.submit(settingsForm)
    })

    expect(updateHandler).toHaveBeenCalledTimes(1)
    expect(getByRole('alert')).toBeTruthy()
    expect(getByRole('alert').textContent).toBe('Failed to update the settings')

    await act(async () => {
      await fireEvent.submit(settingsForm)
    })

    expect(updateHandler).toHaveBeenCalledTimes(2)
  })

  it('submitting when the HTTP response is not ok', async () => {
    const getSettingsHandler = jest.fn().mockResolvedValue({})
    const updateHandler = jest
      .fn()
      .mockRejectedValue(new Error('arbitrary error'))

    const userContext = { user: { id: '42' } }

    const { findByTestId, getByRole } = renderWithProviders(
      <MemoryRouter>
        <UserContext.Provider value={userContext}>
          <InsightSettings
            updateHandler={updateHandler}
            getSettings={getSettingsHandler}
          />
        </UserContext.Provider>
      </MemoryRouter>,
    )

    const settingsForm = await findByTestId('form')

    await act(async () => {
      await fireEvent.submit(settingsForm)
    })

    expect(getByRole('alert')).toBeTruthy()
    expect(getByRole('alert').textContent).toBe('Failed to update the settings')
  })
})
