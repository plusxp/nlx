// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import { Selector, t } from 'testcafe'

class Page {
  constructor() {
    // Wait until visible. Drawer animation can take some time and click interaction requires visibility
    this.serviceName = Selector('[data-testid="service-name"]', {
      visibilityCheck: true,
    })
    this.published = Selector('[data-testid="service-published"]')
    this.editButton = Selector('a[title="Service aanpassen"]')
    this.removeButton = Selector('button[title="Service verwijderen"]')
    this.closeButton = Selector('[data-testid="close-button"]')
    this.inways = Selector('[data-testid="service-inways"]')
    this.alert = Selector('div[role="alert"]')
    this.alertTitle = this.alert.find('[data-testid="title"]')
  }

  async removeService() {
    await t.click(this.removeButton)
    await t.wait(3000)
  }
}

export default new Page()
