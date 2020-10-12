// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { array } from 'prop-types'
import { Table } from '@commonground/design-system'
import { useTranslation } from 'react-i18next'

import Collapsible from '../../../../../components/Collapsible'
import {
  DetailHeading,
  StyledCollapsibleBody,
  StyledCollapsibleEmptyBody,
} from '../../../../../components/DetailView'
import Amount from '../../../../../components/Amount'
import { IconCheckboxMultiple } from '../../../../../icons'

const AccessGrantSection = ({ accessGrants }) => {
  const { t } = useTranslation()

  return (
    <Collapsible
      title={
        <DetailHeading data-testid="service-accessgrants">
          <IconCheckboxMultiple />
          {t('Organizations with access')}
          <Amount value={accessGrants.length} />
        </DetailHeading>
      }
      ariaLabel={t('Organizations with access')}
    >
      <StyledCollapsibleBody>
        <Table data-testid="service-accessgrant-list">
          <tbody>
            {accessGrants.length ? (
              accessGrants.map(({ id, organizationName }) => (
                <Table.Tr data-testid="service-accessgrants" key={id}>
                  <Table.Td>{organizationName}</Table.Td>
                </Table.Tr>
              ))
            ) : (
              <Table.Tr data-testid="service-no-accessgrants">
                <Table.Td>
                  <StyledCollapsibleEmptyBody>
                    {t('There are no organizations with access')}
                  </StyledCollapsibleEmptyBody>
                </Table.Td>
              </Table.Tr>
            )}
          </tbody>
        </Table>
      </StyledCollapsibleBody>
    </Collapsible>
  )
}

AccessGrantSection.propTypes = {
  accessGrants: array,
}
AccessGrantSection.defaultProps = {
  accessGrants: [],
}

export default AccessGrantSection