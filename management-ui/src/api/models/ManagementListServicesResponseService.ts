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
    ManagementListServicesResponseServiceAuthorizationSettings,
    ManagementListServicesResponseServiceAuthorizationSettingsFromJSON,
    ManagementListServicesResponseServiceAuthorizationSettingsFromJSONTyped,
    ManagementListServicesResponseServiceAuthorizationSettingsToJSON,
} from './';

/**
 * 
 * @export
 * @interface ManagementListServicesResponseService
 */
export interface ManagementListServicesResponseService {
    /**
     * 
     * @type {string}
     * @memberof ManagementListServicesResponseService
     */
    name?: string;
    /**
     * 
     * @type {string}
     * @memberof ManagementListServicesResponseService
     */
    endpointURL?: string;
    /**
     * 
     * @type {string}
     * @memberof ManagementListServicesResponseService
     */
    documentationURL?: string;
    /**
     * 
     * @type {string}
     * @memberof ManagementListServicesResponseService
     */
    apiSpecificationURL?: string;
    /**
     * 
     * @type {boolean}
     * @memberof ManagementListServicesResponseService
     */
    internal?: boolean;
    /**
     * 
     * @type {string}
     * @memberof ManagementListServicesResponseService
     */
    techSupportContact?: string;
    /**
     * 
     * @type {string}
     * @memberof ManagementListServicesResponseService
     */
    publicSupportContact?: string;
    /**
     * 
     * @type {ManagementListServicesResponseServiceAuthorizationSettings}
     * @memberof ManagementListServicesResponseService
     */
    authorizationSettings?: ManagementListServicesResponseServiceAuthorizationSettings;
    /**
     * 
     * @type {Array<string>}
     * @memberof ManagementListServicesResponseService
     */
    inways?: Array<string>;
    /**
     * 
     * @type {number}
     * @memberof ManagementListServicesResponseService
     */
    incomingAccessRequestsCount?: number;
}

export function ManagementListServicesResponseServiceFromJSON(json: any): ManagementListServicesResponseService {
    return ManagementListServicesResponseServiceFromJSONTyped(json, false);
}

export function ManagementListServicesResponseServiceFromJSONTyped(json: any, ignoreDiscriminator: boolean): ManagementListServicesResponseService {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'name': !exists(json, 'name') ? undefined : json['name'],
        'endpointURL': !exists(json, 'endpointURL') ? undefined : json['endpointURL'],
        'documentationURL': !exists(json, 'documentationURL') ? undefined : json['documentationURL'],
        'apiSpecificationURL': !exists(json, 'apiSpecificationURL') ? undefined : json['apiSpecificationURL'],
        'internal': !exists(json, 'internal') ? undefined : json['internal'],
        'techSupportContact': !exists(json, 'techSupportContact') ? undefined : json['techSupportContact'],
        'publicSupportContact': !exists(json, 'publicSupportContact') ? undefined : json['publicSupportContact'],
        'authorizationSettings': !exists(json, 'authorizationSettings') ? undefined : ManagementListServicesResponseServiceAuthorizationSettingsFromJSON(json['authorizationSettings']),
        'inways': !exists(json, 'inways') ? undefined : json['inways'],
        'incomingAccessRequestsCount': !exists(json, 'incomingAccessRequestsCount') ? undefined : json['incomingAccessRequestsCount'],
    };
}

export function ManagementListServicesResponseServiceToJSON(value?: ManagementListServicesResponseService | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'name': value.name,
        'endpointURL': value.endpointURL,
        'documentationURL': value.documentationURL,
        'apiSpecificationURL': value.apiSpecificationURL,
        'internal': value.internal,
        'techSupportContact': value.techSupportContact,
        'publicSupportContact': value.publicSupportContact,
        'authorizationSettings': ManagementListServicesResponseServiceAuthorizationSettingsToJSON(value.authorizationSettings),
        'inways': value.inways,
        'incomingAccessRequestsCount': value.incomingAccessRequestsCount,
    };
}

