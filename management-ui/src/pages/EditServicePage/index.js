// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React, { useState } from 'react'
import { func } from 'prop-types'
import { useTranslation } from 'react-i18next'
import { Alert } from '@commonground/design-system'
import { useParams } from 'react-router-dom'
import ServiceForm from '../../components/ServiceForm'
import ServiceRepository from '../../domain/service-repository'
import PageTemplate from '../../components/PageTemplate'
import usePromise from '../../hooks/use-promise'
import LoadingMessage from '../../components/LoadingMessage'
import { StyledUpdatedError } from './index.styles'

const EditServicePage = ({ updateHandler, getServiceByName }) => {
  const { name } = useParams()
  const { t } = useTranslation()
  const [isUpdated, setisUpdated] = useState(false)
  const [updateError, setUpdatedError] = useState(null)
  const { isReady, error, result } = usePromise(getServiceByName, name)

  const submitService = async (service) => {
    // placeholder until we've implemented adding authorizations in the form
    service.authorizationSettings = service.authorizationSettings || {}
    service.authorizationSettings.authorizations =
      service.authorizationSettings.authorizations || []

    try {
      await updateHandler(name, service)
      setUpdatedError(null)
      setisUpdated(true)
    } catch (err) {
      setUpdatedError(err.message)
      setisUpdated(false)
    }
  }

  return (
    <PageTemplate>
      <PageTemplate.HeaderWithBackNavigation
        backButtonTo={`/services/${name}`}
        title={t('Edit existing service')}
      />

      {!isReady || (!error && !result) ? (
        <LoadingMessage />
      ) : error ? (
        <Alert variant="error" data-testid="error-message">
          {t('Failed to load the service.', { name })}
        </Alert>
      ) : result ? (
        <>
          {updateError ? (
            <StyledUpdatedError
              title={t('Failed to update the service.')}
              variant="error"
              data-testid="error-message"
            >
              {t(updateError)}
            </StyledUpdatedError>
          ) : null}

          {isUpdated && !updateError ? (
            <Alert variant="success" data-testid="error-message">
              {t('The service has been updated.')}
            </Alert>
          ) : null}

          {!isUpdated ? (
            <ServiceForm
              initialValues={result}
              onSubmitHandler={(values) => submitService(values)}
              disableName
              submitButtonText={t('Update service')}
            />
          ) : null}
        </>
      ) : null}
    </PageTemplate>
  )
}

EditServicePage.propTypes = {
  updateHandler: func,
  getServiceByName: func,
}

EditServicePage.defaultProps = {
  updateHandler: ServiceRepository.update,
  getServiceByName: ServiceRepository.getByName,
}

export default EditServicePage
