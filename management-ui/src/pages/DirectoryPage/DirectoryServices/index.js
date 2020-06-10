// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { useTranslation } from 'react-i18next'
import { func } from 'prop-types'
import React from 'react'
import { Table } from '@commonground/design-system'

import EmptyContentMessage from '../../../components/EmptyContentMessage'
import StatusIndicator from '../../../components/StatusIndicator'

const DirectoryServices = ({ directoryServices }) => {
  const { t } = useTranslation()
  const { services = [] } = directoryServices()

  return services.length === 0 ? (
    <EmptyContentMessage data-testid="directory-no-services">
      {t('There are no services yet.')}
    </EmptyContentMessage>
  ) : (
    <Table withLinks role="grid" data-testid="directory-services">
      <thead>
        <Table.TrHead>
          <Table.Th>{t('Organization')}</Table.Th>
          <Table.Th>{t('Service')}</Table.Th>
          <Table.Th>{t('Status')}</Table.Th>
          <Table.Th>{t('API Type')}</Table.Th>
        </Table.TrHead>
      </thead>
      <tbody>
        {services.map((service, i) => (
          <Table.Tr
            key={i}
            to={`/directory/${service.organizationName}/${service.serviceName}`}
            data-testid={`directory-service-row-${i}`}
          >
            <Table.Td>{service.organizationName}</Table.Td>
            <Table.Td>{service.serviceName}</Table.Td>
            <Table.Td>
              <StatusIndicator status={service.status} />
            </Table.Td>
            <Table.Td>{service.apiSpecificationType}</Table.Td>
          </Table.Tr>
        ))}
      </tbody>
    </Table>
  )
}

DirectoryServices.propTypes = { directoryServices: func.isRequired }

export default DirectoryServices
