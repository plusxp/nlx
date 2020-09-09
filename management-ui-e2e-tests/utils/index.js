// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import fs from 'fs'
import path from 'path'

import saveBrowserConsoleMessages from './save-browser-console-messages'
import saveRequests from './save-requests'

export { default as getBaseUrl } from './getBaseUrl'
export { default as getLocation } from './getLocation'
export { default as doAccessibilityTest } from './doAccessibilityTest'
export { default as dismissAlertWithText } from './dismissAlertWithText'

export const saveBrowserConsoleAndRequests = async (
  testController,
  requests,
) => {
  const userAgent =
    testController.testRun.browserConnection.browserInfo.parsedUserAgent
      .prettyUserAgent
  const fixtureName = testController.testRun.test.testFile.currentFixture.name

  const directory = path.join(
    process.cwd(),
    'test-results',
    'extra',
    userAgent.replace('/', '-'),
    fixtureName,
  )
  const fileNameBase = `${testController.testRun.test.name}`

  await fs.promises.mkdir(directory, { recursive: true })

  await saveBrowserConsoleMessages({
    testController,
    directory,
    fileName: `${fileNameBase} - browser console.json`,
  })

  await saveRequests({
    testController,
    requests,
    directory,
    fileName: `${fileNameBase} - requests.json`,
  })
}
