const { analyzeAccessibility } = require('../accessibility')
const getBaseUrl = require('../getBaseUrl')
const baseUrl = getBaseUrl();

describe('NotFoundPage', () => {
    beforeAll(async () => {
        await page.setBypassCSP(true)

        // page should be loaded
        await page.goto(`${baseUrl}/page-that-does-not-exist`, { waitUntil: 'load' })

        // page content should be rendered
        await page.waitForSelector('h1')
    })

    it('should not have accessibility issues', async () => {
        const accessibilityReport = await analyzeAccessibility(page, `services.accessibility.png`)
        expect(accessibilityReport).toHaveNoAccessibilityIssues();
    })
})