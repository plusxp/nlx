// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React, { useContext } from 'react'
import { observer } from 'mobx-react'
import { array } from 'prop-types'
import { Table, ToasterContext } from '@commonground/design-system'
import { useTranslation } from 'react-i18next'

import Collapsible from '../../../../../components/Collapsible'
import {
  DetailHeading,
  StyledCollapsibleBody,
  StyledCollapsibleEmptyBody,
} from '../../../../../components/DetailView'
import Amount from '../../../../../components/Amount'
import { IconKey } from '../../../../../icons'
import IncomingAccessRequestRow from './IncomingAccessRequestRow'

const AccessRequestsSection = ({ accessRequests }) => {
  const { t } = useTranslation()
  const { showToast } = useContext(ToasterContext)

  const approveHandler = async (accessRequest) => {
    await accessRequest.approve()

    if (!accessRequest.error) {
      showToast({
        title: t('Access request approved'),
        body: t('Organization has access to service', {
          organizationName: accessRequest.organizationName,
          serviceName: accessRequest.serviceName,
        }),
        variant: 'success',
      })
    } else {
      showToast({
        title: t('Failed to approve access request'),
        body: t('Please try again'),
        variant: 'error',
      })
    }
  }

  return (
    <Collapsible
      title={
        <DetailHeading data-testid="service-incoming-accessrequests">
          <IconKey />
          {t('Access requests')}
          <Amount value={accessRequests.length} isAccented />
        </DetailHeading>
      }
      ariaLabel={t('Access requests')}
    >
      <StyledCollapsibleBody>
        <Table data-testid="service-incoming-accessrequests-list">
          <tbody>
            {accessRequests.length ? (
              accessRequests.map((accessRequest, i) => (
                <IncomingAccessRequestRow
                  key={accessRequest.id}
                  accessRequest={accessRequest}
                  approveHandler={approveHandler}
                />
              ))
            ) : (
              <Table.Tr data-testid="service-no-incoming-accessrequests">
                <Table.Td>
                  <StyledCollapsibleEmptyBody>
                    {t('There are no access requests')}
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

AccessRequestsSection.propTypes = {
  accessRequests: array,
}

AccessRequestsSection.defaultProps = {
  accessRequests: [],
}

export default observer(AccessRequestsSection)