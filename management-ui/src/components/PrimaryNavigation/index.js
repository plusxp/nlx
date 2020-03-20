import { Link } from 'react-router-dom'
import React from 'react'
import { useTranslation } from 'react-i18next'
import {
  StyledIcon,
  StyledLink,
  StyledNLXManagementLogo,
  StyledNav,
} from './index.styles'

const IconServices = ({ ...props }) => (
  <svg viewBox="0 0 24 24" {...props}>
    <path d="M13 18v2h6v2h-6a2 2 0 0 1-2-2v-2H8a4 4 0 0 1-4-4V7a1 1 0 0 1 1-1h3V2h2v4h4V2h2v4h3a1 1 0 0 1 1 1v7a4 4 0 0 1-4 4h-3zm-5-2h8a2 2 0 0 0 2-2v-3H6v3a2 2 0 0 0 2 2zm10-8H6v1h12V8zm-6 6.5a1 1 0 1 1 0-2 1 1 0 0 1 0 2z" />
  </svg>
)

const PrimaryNavigation = () => {
  const { t } = useTranslation()
  return (
    <StyledNav>
      <Link to="/" title={t('NLX Dashboard homepage')} aria-label="Homepage">
        <StyledNLXManagementLogo />
      </Link>

      <StyledLink to="/services" aria-label="Services page">
        <StyledIcon as={IconServices} />
        {t('Services')}
      </StyledLink>
    </StyledNav>
  )
}

export default PrimaryNavigation