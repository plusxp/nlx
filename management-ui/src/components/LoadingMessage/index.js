// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { Spinner } from '@commonground/design-system'
import React from 'react'
import { useTranslation } from 'react-i18next'
import { StyledLoadingMessage } from './index.styles'

const LoadingMessage = () => {
  const { t } = useTranslation()
  return (
    <StyledLoadingMessage role="progressbar">
      <Spinner /> {t('Loading…')}
    </StyledLoadingMessage>
  )
}

export default LoadingMessage
