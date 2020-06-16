// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { useTranslation } from 'react-i18next'

import { IconArrowLeftRight, IconServices, IconDirectory } from '../../icons'
import {
  StyledIcon,
  StyledLink,
  StyledHomeLink,
  StyledNLXManagementLogo,
  StyledNav,
} from './index.styles'

const PrimaryNavigation = () => {
  const { t } = useTranslation()
  return (
    <StyledNav>
      <StyledHomeLink
        to="/"
        title={t('NLX Dashboard homepage')}
        aria-label={t('Homepage')}
      >
        <StyledNLXManagementLogo />
      </StyledHomeLink>

      <StyledLink to="/inways" aria-label={t('Inways page')}>
        <StyledIcon as={IconArrowLeftRight} />
        {t('Inways')}
      </StyledLink>

      <StyledLink to="/services" aria-label={t('Services page')}>
        <StyledIcon as={IconServices} />
        {t('Services')}
      </StyledLink>

      <StyledLink to="/directory" aria-label={t('Directory page')}>
        <StyledIcon as={IconDirectory} />
        {t('Directory')}
      </StyledLink>
    </StyledNav>
  )
}

export default PrimaryNavigation
