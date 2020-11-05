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
    NlxmanagementService,
    NlxmanagementServiceFromJSON,
    NlxmanagementServiceFromJSONTyped,
    NlxmanagementServiceToJSON,
} from './';

/**
 * 
 * @export
 * @interface ManagementListServicesResponse
 */
export interface ManagementListServicesResponse {
    /**
     * 
     * @type {Array<NlxmanagementService>}
     * @memberof ManagementListServicesResponse
     */
    services?: Array<NlxmanagementService>;
}

export function ManagementListServicesResponseFromJSON(json: any): ManagementListServicesResponse {
    return ManagementListServicesResponseFromJSONTyped(json, false);
}

export function ManagementListServicesResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): ManagementListServicesResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'services': !exists(json, 'services') ? undefined : ((json['services'] as Array<any>).map(NlxmanagementServiceFromJSON)),
    };
}

export function ManagementListServicesResponseToJSON(value?: ManagementListServicesResponse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'services': value.services === undefined ? undefined : ((value.services as Array<any>).map(NlxmanagementServiceToJSON)),
    };
}

