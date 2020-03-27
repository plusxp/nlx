// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React from 'react'
import { renderWithProviders } from '../../test-utils'
import Amount from './index'

test('renders without crashing', () => {
  expect(() => renderWithProviders(<Amount value={1} />)).not.toThrow()
})
