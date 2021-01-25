// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { string } from 'prop-types'
import { useTranslation } from 'react-i18next'
import {
  SectionHeader,
  StyledIconServices,
  ServiceField,
  ServiceData,
} from './index.styles'

const RequestAccessDetails = ({ organizationName, serviceName }) => {
  const { t } = useTranslation()

  return (
    <>
      <p>{t('You are requesting access to a service')}.</p>

      <section>
        <SectionHeader>{t('Service')}</SectionHeader>
        <ServiceField>
          <StyledIconServices />
          <ServiceData>
            <strong>{serviceName}</strong>
            <span>{organizationName}</span>
          </ServiceData>
        </ServiceField>
      </section>
    </>
  )
}

RequestAccessDetails.propTypes = {
  organizationName: string.isRequired,
  serviceName: string.isRequired,
}

export default RequestAccessDetails