// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { MemoryRouter as Router } from 'react-router-dom'
import { fireEvent, within, waitFor } from '@testing-library/react'
import { renderWithProviders } from '../../../../test-utils'
import ServiceDetailView from './index'

jest.mock('./InwaysSection', () => () => <div />)
jest.mock('./AccessRequestSection', () => () => <div />)
jest.mock('./AccessGrantSection', () => () => <div />)

const service = {
  name: 'name',
  internal: false,
  inways: [],
}

describe('ServiceDetails', () => {
  it('should display', () => {
    const { queryByText } = renderWithProviders(
      <Router>
        <ServiceDetailView service={service} removeHandler={jest.fn()} />
      </Router>,
    )

    const heading = queryByText('Published in central directory')
    expect(heading).toBeInTheDocument()
    expect(heading).toHaveTextContent('visible.svg')
  })

  it('should show hidden icon', () => {
    const { queryByText } = renderWithProviders(
      <Router>
        <ServiceDetailView
          service={{
            ...service,
            internal: true,
          }}
          removeHandler={jest.fn()}
        />
      </Router>,
    )

    const heading = queryByText('Not visible in central directory')
    expect(heading).toBeInTheDocument()
    expect(heading).toHaveTextContent('hidden.svg')
  })

  it('should call the removeHandler on remove', async () => {
    const handleRemove = jest.fn()
    const { getByTitle, getByRole } = renderWithProviders(
      <Router>
        <ServiceDetailView service={service} removeHandler={handleRemove} />
      </Router>,
    )

    fireEvent.click(getByTitle('Remove service'))

    const confirmModal = getByRole('dialog')
    const okButton = within(confirmModal).getByText('Remove')

    fireEvent.click(okButton)
    await waitFor(() => expect(handleRemove).toHaveBeenCalled())
  })
})
