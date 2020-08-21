// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React, { useContext, useState } from 'react'
import { useTranslation } from 'react-i18next'
import { Alert, ToasterContext } from '@commonground/design-system'
import { func } from 'prop-types'
import PageTemplate from '../../components/PageTemplate'
import SettingsForm from '../../components/SettingsForm'
import usePromise from '../../hooks/use-promise'
import LoadingMessage from '../../components/LoadingMessage'
import { StyledUpdatedError } from '../services/EditServicePage/index.styles'
import SettingsRepository from '../../domain/settings-repository'

const SettingsPage = ({ getSettings, updateHandler }) => {
  const { t } = useTranslation()
  const { showToast } = useContext(ToasterContext)
  const [updateError, setUpdatedError] = useState(null)
  const { isReady, error, result: settings } = usePromise(getSettings)

  const updateSettings = async (values) => {
    try {
      await updateHandler(values)

      showToast({
        body: t('Successfully updated the settings.'),
        variant: 'success',
      })

      setUpdatedError(false)
    } catch (err) {
      setUpdatedError(err.message)
    }
  }

  return (
    <PageTemplate>
      <PageTemplate.Header title={t('Settings')} />

      {!isReady || (!error && !settings) ? (
        <LoadingMessage />
      ) : error ? (
        <Alert variant="error" data-testid="error-message">
          {t('Failed to load the settings.')}
        </Alert>
      ) : settings ? (
        <>
          {updateError ? (
            <StyledUpdatedError
              title={t('Failed to update the settings.')}
              variant="error"
              data-testid="error-message"
            >
              {t(`${updateError}`)}
            </StyledUpdatedError>
          ) : null}

          <SettingsForm
            initialValues={settings}
            onSubmitHandler={(values) => updateSettings(values)}
          />
        </>
      ) : null}
    </PageTemplate>
  )
}

SettingsPage.propTypes = {
  updateHandler: func,
  getSettings: func,
}

SettingsPage.defaultProps = {
  updateHandler: SettingsRepository.update,
  getSettings: SettingsRepository.get,
}

export default SettingsPage