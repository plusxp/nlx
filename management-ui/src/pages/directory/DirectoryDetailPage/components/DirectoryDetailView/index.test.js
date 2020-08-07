// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { observable, extendObservable } from 'mobx'
import { fireEvent } from '@testing-library/react'

import { renderWithProviders } from '../../../../../test-utils'
import DirectoryDetailView from './index'

describe('detail view of directory service we do not have access to', () => {
  let service

  global.confirm = jest.fn(() => true)

  beforeEach(() => {
    service = observable({
      organizationName: 'Organization',
      latestAccessRequest: null,
      requestAccess: jest.fn(),
    })
  })

  it('should have a button to request access', () => {
    const { getByText } = renderWithProviders(
      <DirectoryDetailView service={service} />,
    )

    const button = getByText('Request Access')
    expect(button).toBeInTheDocument()

    fireEvent.click(button)
    expect(service.requestAccess).toHaveBeenCalled()
  })

  it('should show a loading message', () => {
    const serviceWithOutstandingRequest = extendObservable(service, {
      latestAccessRequest: {
        id: 'string',
        state: 'CREATED',
        createdAt: '2020-06-30T08:31:41.106Z',
        updatedAt: '2020-06-30T08:31:41.106Z',
      },
    })

    const { getByText } = renderWithProviders(
      <DirectoryDetailView service={serviceWithOutstandingRequest} />,
    )

    expect(getByText('Sending request')).toBeInTheDocument()
  })

  it('should show a failed message', () => {
    const serviceWithFailedRequest = extendObservable(service, {
      latestAccessRequest: {
        id: 'string',
        state: 'FAILED',
        createdAt: '2020-06-30T08:31:41.106Z',
        updatedAt: '2020-06-30T08:31:41.106Z',
      },
    })

    const { getAllByText } = renderWithProviders(
      <DirectoryDetailView service={serviceWithFailedRequest} />,
    )

    const failedMessages = getAllByText('Request could not be sent')

    expect(failedMessages).toHaveLength(2)
  })
})