// Copyright © VNG Realisatie 2018
// Licensed under the EUPL
//

import React from 'react'
import ReactDOM from 'react-dom'
import Card from './Card'

test('renders without crashing', () => {
  expect(() => {
    const div = document.createElement('div')
    ReactDOM.render(<Card />, div)
  }).not.toThrow()
})
