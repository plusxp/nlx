// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import AccessRequestRepository from './access-request-repository'

describe('the AccessRequestRepository', () => {
  describe('requesting access to a service', () => {
    beforeEach(() => {
      jest.spyOn(global, 'fetch').mockImplementation(() =>
        Promise.resolve({
          ok: true,
          status: 201,
          json: () =>
            Promise.resolve({
              id: 'string',
              organizationName: 'organization',
              serviceName: 'service',
              status: 'CREATED',
              createdAt: '2020-06-30T10:36:57.100Z',
              updatedAt: '2020-06-30T10:36:57.100Z',
            }),
        }),
      )
    })

    afterEach(() => global.fetch.mockRestore())

    it('should return the services', async () => {
      const result = await AccessRequestRepository.requestAccess({
        organizationName: 'organization',
        serviceName: 'service',
      })

      expect(result).toEqual({
        id: 'string',
        organizationName: 'organization',
        serviceName: 'service',
        status: 'CREATED',
        createdAt: '2020-06-30T10:36:57.100Z',
        updatedAt: '2020-06-30T10:36:57.100Z',
      })
    })
  })
})