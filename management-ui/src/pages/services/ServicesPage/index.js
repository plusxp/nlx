// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React, { useEffect } from 'react'
import { useTranslation } from 'react-i18next'
import { Alert, Button } from '@commonground/design-system'
import { Link, Route } from 'react-router-dom'

import { observer } from 'mobx-react'
import PageTemplate from '../../../components/PageTemplate'
import ServiceDetailPage from '../ServiceDetailPage'

import LoadingMessage from '../../../components/LoadingMessage'

import { useServicesStore } from '../../../hooks/use-stores'
import ServiceToastManager from './ServiceToastManager'
import ServiceCount from './ServiceCount'
import ServicesPageView from './ServicesPageView'
import { StyledActionsBar, StyledIconPlus } from './index.styles'

const ServicesPage = () => {
  const { t } = useTranslation()
  const { isReady, services, error, fetchServices } = useServicesStore()

  useEffect(() => {
    fetchServices()
  }, []) // eslint-disable-line react-hooks/exhaustive-deps

  return (
    <PageTemplate>
      <ServiceToastManager />

      <PageTemplate.Header title={t('Services')} />

      <StyledActionsBar>
        <ServiceCount
          count={services ? services.length : 0}
          data-testid="service-count"
        />
        <Button
          as={Link}
          to="/services/add-service"
          aria-label={t('Add service')}
        >
          <StyledIconPlus />
          {t('Add service')}
        </Button>
      </StyledActionsBar>

      {!isReady ? (
        <LoadingMessage />
      ) : error ? (
        <Alert variant="error" data-testid="error-message">
          {t('Failed to load the services.')}
        </Alert>
      ) : (
        <ServicesPageView services={services} />
      )}

      <Route path="/services/:name">
        <ServiceDetailPage parentUrl="/services" />
      </Route>
    </PageTemplate>
  )
}

export default observer(ServicesPage)