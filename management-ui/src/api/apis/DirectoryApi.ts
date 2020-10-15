/* tslint:disable */
/* eslint-disable */
/**
 * management.proto
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: version not set
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import * as runtime from '../runtime';
import {
    ManagementDirectoryListServicesResponse,
    ManagementDirectoryListServicesResponseFromJSON,
    ManagementDirectoryListServicesResponseToJSON,
    ManagementDirectoryService,
    ManagementDirectoryServiceFromJSON,
    ManagementDirectoryServiceToJSON,
    ManagementOutgoingAccessRequest,
    ManagementOutgoingAccessRequestFromJSON,
    ManagementOutgoingAccessRequestToJSON,
    RuntimeError,
    RuntimeErrorFromJSON,
    RuntimeErrorToJSON,
} from '../models';

export interface DirectoryGetOrganizationServiceRequest {
    organizationName: string;
    serviceName: string;
}

export interface DirectoryRequestAccessToServiceRequest {
    organizationName: string;
    serviceName: string;
}

/**
 * 
 */
export class DirectoryApi extends runtime.BaseAPI {

    /**
     */
    async directoryGetOrganizationServiceRaw(requestParameters: DirectoryGetOrganizationServiceRequest): Promise<runtime.ApiResponse<ManagementDirectoryService>> {
        if (requestParameters.organizationName === null || requestParameters.organizationName === undefined) {
            throw new runtime.RequiredError('organizationName','Required parameter requestParameters.organizationName was null or undefined when calling directoryGetOrganizationService.');
        }

        if (requestParameters.serviceName === null || requestParameters.serviceName === undefined) {
            throw new runtime.RequiredError('serviceName','Required parameter requestParameters.serviceName was null or undefined when calling directoryGetOrganizationService.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/api/v1/directory/organizations/{organizationName}/services/{serviceName}`.replace(`{${"organizationName"}}`, encodeURIComponent(String(requestParameters.organizationName))).replace(`{${"serviceName"}}`, encodeURIComponent(String(requestParameters.serviceName))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => ManagementDirectoryServiceFromJSON(jsonValue));
    }

    /**
     */
    async directoryGetOrganizationService(requestParameters: DirectoryGetOrganizationServiceRequest): Promise<ManagementDirectoryService> {
        const response = await this.directoryGetOrganizationServiceRaw(requestParameters);
        return await response.value();
    }

    /**
     */
    async directoryListServicesRaw(): Promise<runtime.ApiResponse<ManagementDirectoryListServicesResponse>> {
        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/api/v1/directory/services`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => ManagementDirectoryListServicesResponseFromJSON(jsonValue));
    }

    /**
     */
    async directoryListServices(): Promise<ManagementDirectoryListServicesResponse> {
        const response = await this.directoryListServicesRaw();
        return await response.value();
    }

    /**
     */
    async directoryRequestAccessToServiceRaw(requestParameters: DirectoryRequestAccessToServiceRequest): Promise<runtime.ApiResponse<ManagementOutgoingAccessRequest>> {
        if (requestParameters.organizationName === null || requestParameters.organizationName === undefined) {
            throw new runtime.RequiredError('organizationName','Required parameter requestParameters.organizationName was null or undefined when calling directoryRequestAccessToService.');
        }

        if (requestParameters.serviceName === null || requestParameters.serviceName === undefined) {
            throw new runtime.RequiredError('serviceName','Required parameter requestParameters.serviceName was null or undefined when calling directoryRequestAccessToService.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/api/v1/directory/organizations/{organizationName}/services/{serviceName}/access-requests`.replace(`{${"organizationName"}}`, encodeURIComponent(String(requestParameters.organizationName))).replace(`{${"serviceName"}}`, encodeURIComponent(String(requestParameters.serviceName))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => ManagementOutgoingAccessRequestFromJSON(jsonValue));
    }

    /**
     */
    async directoryRequestAccessToService(requestParameters: DirectoryRequestAccessToServiceRequest): Promise<ManagementOutgoingAccessRequest> {
        const response = await this.directoryRequestAccessToServiceRaw(requestParameters);
        return await response.value();
    }

}
