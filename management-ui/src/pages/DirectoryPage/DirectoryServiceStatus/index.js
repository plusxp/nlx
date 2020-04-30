// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { oneOf } from 'prop-types'
import { useTranslation } from 'react-i18next'

import { ReactComponent as IconStatusUp } from './status-up.svg'
import { ReactComponent as IconStatusDown } from './status-down.svg'
import { ReactComponent as IconStatusUnknown } from './status-unknown.svg'

import { StyledWrapper, StyledIconStatusDegraded } from './index.styles'

export const DIRECTORY_SERVICE_STATUS = ['degraded', 'down', 'unknown', 'up']

const DirectoryServiceStatus = ({ status }) => {
  const { t } = useTranslation()

  if (!DIRECTORY_SERVICE_STATUS.includes(status)) {
    console.warn(`invalid status '${status}'`)
    return null
  }
  return (
    <StyledWrapper>
      {
        {
          degraded: <StyledIconStatusDegraded title={t('Degraded')} />,
          down: <IconStatusDown title={t('Down')} />,
          up: <IconStatusUp title={t('Up')} />,
          unknown: <IconStatusUnknown title={t('Unknown')} />,
        }[status]
      }
    </StyledWrapper>
  )
}

DirectoryServiceStatus.propTypes = {
  status: oneOf(DIRECTORY_SERVICE_STATUS),
}

export default DirectoryServiceStatus
