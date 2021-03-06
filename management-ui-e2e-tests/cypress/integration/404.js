// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//

import logViolations from '../axe-utilities/log-violations'

describe('404', () => {
  beforeEach(() => {
    cy.visit('/page-that-does-not-exist')
    cy.injectAxe()
  })

  it('Page is accessible', () => {
    cy.findByText('Pagina niet gevonden').should('exist')
    cy.checkA11y(null, null, logViolations)
  })
})
