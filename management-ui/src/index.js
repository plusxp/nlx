// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import 'react-app-polyfill/ie11'
import 'core-js/features/array/find'
import 'core-js/features/array/includes'

import React from 'react'
import ReactDOM from 'react-dom'
import './i18n'

import { BrowserRouter as Router } from 'react-router-dom'
import App from './App'
import { UserContextProvider } from './user-context'

ReactDOM.render(
  <Router>
    <UserContextProvider>
      <App />
    </UserContextProvider>
  </Router>,
  document.getElementById('root'),
)
