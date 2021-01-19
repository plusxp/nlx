// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import React from 'react'
import { renderWithProviders } from '../../../../../../test-utils'
import CollapsibleHeader from './index'

test('the CollapsibleHeader component', async () => {
  const { rerender, getByText, getByTestId } = renderWithProviders(
    <CollapsibleHeader counter={0} />,
  )

  const toggler = getByText(/Access requests/i)

  expect(toggler).toHaveTextContent(
    'key.svg' + 'Access requests' + '0', // eslint-disable-line no-useless-concat
  )
  expect(getByTestId('amount')).toBeInTheDocument()

  rerender(<CollapsibleHeader counter={1} />)

  expect(toggler).toHaveTextContent(
    'key.svg' + 'Access requests' + '1', // eslint-disable-line no-useless-concat
  )
  expect(getByTestId('amount-accented')).toBeInTheDocument()
})
