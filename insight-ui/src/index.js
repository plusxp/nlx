import React from 'react'
import ReactDOM from 'react-dom'
import App from './App'

import { MuiThemeProvider } from '@material-ui/core/styles'
import muiTheme from './styles/muiTheme'

import { Provider } from 'react-redux'
import configureStore from './configureStore'

const store = configureStore()

ReactDOM.render(
    <Provider store={store}>
        <MuiThemeProvider theme={muiTheme}>
            <App />
        </MuiThemeProvider>
    </Provider>,
    document.getElementById('root'),
)
