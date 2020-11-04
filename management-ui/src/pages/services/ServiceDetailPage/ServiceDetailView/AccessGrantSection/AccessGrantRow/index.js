// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { object, func } from 'prop-types'
import { Table } from '@commonground/design-system'
import { useTranslation } from 'react-i18next'

import ButtonWithIcon from '../../../../../../components/ButtonWithIcon'
import { IconRevoke } from '../../../../../../icons'
import { TdActions } from './index.styles'

const AccessGrantRow = ({ accessGrant, revokeHandler }) => {
  const { t } = useTranslation()

  const handleRevoke = (evt) => {
    evt.stopPropagation()

    const confirmed = window.confirm(
      t(
        'Access will be revoked for the serviceName service from organizationName',
        {
          organizationName: accessGrant.organizationName,
          serviceName: accessGrant.serviceName,
        },
      ),
    )

    if (confirmed) {
      revokeHandler(accessGrant)
    }
  }

  return (
    <Table.Tr data-testid="service-accessgrants" key={accessGrant.id}>
      <Table.Td>{accessGrant.organizationName}</Table.Td>
      <TdActions>
        <ButtonWithIcon size="small" variant="link" onClick={handleRevoke}>
          <IconRevoke />
          {t('Revoke')}
        </ButtonWithIcon>
      </TdActions>
    </Table.Tr>
  )
}

AccessGrantRow.propTypes = {
  accessGrant: object,
  revokeHandler: func,
}
AccessGrantRow.defaultProps = {}

export default AccessGrantRow
