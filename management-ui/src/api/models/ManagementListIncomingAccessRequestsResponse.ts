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

import { exists, mapValues } from '../runtime';
import {
    ManagementIncomingAccessRequest,
    ManagementIncomingAccessRequestFromJSON,
    ManagementIncomingAccessRequestFromJSONTyped,
    ManagementIncomingAccessRequestToJSON,
} from './';

/**
 * 
 * @export
 * @interface ManagementListIncomingAccessRequestsResponse
 */
export interface ManagementListIncomingAccessRequestsResponse {
    /**
     * 
     * @type {Array<ManagementIncomingAccessRequest>}
     * @memberof ManagementListIncomingAccessRequestsResponse
     */
    accessRequests?: Array<ManagementIncomingAccessRequest>;
}

export function ManagementListIncomingAccessRequestsResponseFromJSON(json: any): ManagementListIncomingAccessRequestsResponse {
    return ManagementListIncomingAccessRequestsResponseFromJSONTyped(json, false);
}

export function ManagementListIncomingAccessRequestsResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): ManagementListIncomingAccessRequestsResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'accessRequests': !exists(json, 'accessRequests') ? undefined : ((json['accessRequests'] as Array<any>).map(ManagementIncomingAccessRequestFromJSON)),
    };
}

export function ManagementListIncomingAccessRequestsResponseToJSON(value?: ManagementListIncomingAccessRequestsResponse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'accessRequests': value.accessRequests === undefined ? undefined : ((value.accessRequests as Array<any>).map(ManagementIncomingAccessRequestToJSON)),
    };
}

