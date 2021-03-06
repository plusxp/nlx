// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React, { useContext } from 'react'
import { useTranslation } from 'react-i18next'
import { Alert, ToasterContext } from '@commonground/design-system'
import { func } from 'prop-types'

import { useApplicationStore } from '../../../hooks/use-stores'
import usePromise from '../../../hooks/use-promise'
import LoadingMessage from '../../../components/LoadingMessage'
import SettingsRepository from '../../../domain/settings-repository'
import Form from './Form'

const GeneralSettings = ({ getSettings, updateHandler }) => {
  const { t } = useTranslation()
  const { showToast } = useContext(ToasterContext)
  const { isReady, error, result: settings } = usePromise(getSettings)
  const applicationStore = useApplicationStore()

  const updateSettings = async (values) => {
    try {
      await updateHandler(values)
      applicationStore.update({
        isOrganizationInwaySet: !!values.organizationInway,
      })

      showToast({
        body: t('Successfully updated the settings'),
        variant: 'success',
      })
    } catch (err) {
      showToast({
        body: t('Failed to update the settings'),
        variant: 'error',
      })
    }
  }

  return (
    <>
      {!isReady ? (
        <LoadingMessage />
      ) : error ? (
        <Alert variant="error" data-testid="error-message">
          {t('Failed to load the settings')}
        </Alert>
      ) : settings ? (
        <Form initialValues={settings} onSubmitHandler={updateSettings} />
      ) : null}
    </>
  )
}

GeneralSettings.propTypes = {
  updateHandler: func,
  getSettings: func,
}

GeneralSettings.defaultProps = {
  updateHandler: SettingsRepository.updateGeneralSettings,
  getSettings: SettingsRepository.getGeneralSettings,
}

export default GeneralSettings
