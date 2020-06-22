// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import DirectoryRepository from './directory-repository'

describe('the DirectoryRepository', () => {
  describe('getting all services', () => {
    afterEach(() => global.fetch.mockRestore())

    it('should return the services', async () => {
      jest.spyOn(global, 'fetch').mockResolvedValue({
        ok: true,
        status: 200,
        json: () =>
          Promise.resolve({
            services: [
              {
                serviceName: 'service',
                organizationName: 'organization',
                status: 'unknown',
                apiSpecificationType: 'plain',
                latestAccessRequest: {
                  status: 'FAILED',
                },
              },
            ],
          }),
      })

      const result = await DirectoryRepository.getAll()

      expect(result).toEqual([
        {
          serviceName: 'service',
          organizationName: 'organization',
          status: 'unknown',
          apiSpecificationType: 'plain',
          latestAccessRequest: {
            status: 'FAILED',
          },
        },
      ])

      expect(global.fetch).toHaveBeenCalledWith(
        '/api/v1/directory/services',
        expect.anything(),
      )
    })

    it('should return the services and add missing latestAccessRequest properties', async () => {
      jest.spyOn(global, 'fetch').mockResolvedValue({
        ok: true,
        status: 200,
        json: () =>
          Promise.resolve({
            services: [
              {
                serviceName: 'service',
                organizationName: 'organization',
                status: 'unknown',
                apiSpecificationType: 'plain',
                latestAccessRequest: {
                  status: 'FAILED',
                },
              },
              {
                serviceName: 'service2',
                organizationName: 'organization',
                status: 'unknown',
                apiSpecificationType: 'plain',
              },
            ],
          }),
      })

      const result = await DirectoryRepository.getAll()

      expect(result).toEqual([
        {
          serviceName: 'service',
          organizationName: 'organization',
          status: 'unknown',
          apiSpecificationType: 'plain',
          latestAccessRequest: {
            status: 'FAILED',
          },
        },
        {
          serviceName: 'service2',
          organizationName: 'organization',
          status: 'unknown',
          apiSpecificationType: 'plain',
          latestAccessRequest: null,
        },
      ])

      expect(global.fetch).toHaveBeenCalledWith(
        '/api/v1/directory/services',
        expect.anything(),
      )
    })

    it('should return an empty list when the response is an empty object', async () => {
      jest.spyOn(global, 'fetch').mockResolvedValue({
        ok: true,
        status: 200,
        json: () => Promise.resolve({}),
      })

      const result = await DirectoryRepository.getAll()

      expect(result).toEqual([])

      expect(global.fetch).toHaveBeenCalledWith(
        '/api/v1/directory/services',
        expect.anything(),
      )
    })
  })

  describe('getting a specific service', () => {
    it('should return the expected service', async () => {
      jest.spyOn(global, 'fetch').mockResolvedValue({
        ok: true,
        status: 200,
        json: () =>
          Promise.resolve({
            serviceName: 'service',
            organizationName: 'organization',
            status: 'unknown',
            apiSpecificationType: 'plain',
            latestAccessRequest: {
              status: 'FAILED',
            },
          }),
      })

      const result = await DirectoryRepository.getByName(
        'organization',
        'service',
      )

      expect(result).toEqual({
        serviceName: 'service',
        organizationName: 'organization',
        status: 'unknown',
        apiSpecificationType: 'plain',
        latestAccessRequest: {
          status: 'FAILED',
        },
      })

      expect(global.fetch).toHaveBeenCalledWith(
        '/api/v1/directory/organizations/organization/services/service',
        expect.anything(),
      )
    })

    it('should add latestAccessRequest prop (null)', async () => {
      jest.spyOn(global, 'fetch').mockResolvedValue({
        ok: true,
        status: 200,
        json: () =>
          Promise.resolve({
            serviceName: 'service',
            organizationName: 'organization',
            status: 'unknown',
            apiSpecificationType: 'plain',
          }),
      })

      const result = await DirectoryRepository.getByName(
        'organization',
        'service',
      )

      expect(result).toEqual({
        serviceName: 'service',
        organizationName: 'organization',
        status: 'unknown',
        apiSpecificationType: 'plain',
        latestAccessRequest: null,
      })
    })
  })
})
