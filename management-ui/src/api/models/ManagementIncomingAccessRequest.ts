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
    ManagementAccessRequestState,
    ManagementAccessRequestStateFromJSON,
    ManagementAccessRequestStateFromJSONTyped,
    ManagementAccessRequestStateToJSON,
} from './';

/**
 * 
 * @export
 * @interface ManagementIncomingAccessRequest
 */
export interface ManagementIncomingAccessRequest {
    /**
     * 
     * @type {string}
     * @memberof ManagementIncomingAccessRequest
     */
    id?: string;
    /**
     * 
     * @type {string}
     * @memberof ManagementIncomingAccessRequest
     */
    organizationName?: string;
    /**
     * 
     * @type {string}
     * @memberof ManagementIncomingAccessRequest
     */
    serviceName?: string;
    /**
     * 
     * @type {ManagementAccessRequestState}
     * @memberof ManagementIncomingAccessRequest
     */
    state?: ManagementAccessRequestState;
    /**
     * 
     * @type {Date}
     * @memberof ManagementIncomingAccessRequest
     */
    createdAt?: Date;
    /**
     * 
     * @type {Date}
     * @memberof ManagementIncomingAccessRequest
     */
    updatedAt?: Date;
}

export function ManagementIncomingAccessRequestFromJSON(json: any): ManagementIncomingAccessRequest {
    return ManagementIncomingAccessRequestFromJSONTyped(json, false);
}

export function ManagementIncomingAccessRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): ManagementIncomingAccessRequest {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'id': !exists(json, 'id') ? undefined : json['id'],
        'organizationName': !exists(json, 'organizationName') ? undefined : json['organizationName'],
        'serviceName': !exists(json, 'serviceName') ? undefined : json['serviceName'],
        'state': !exists(json, 'state') ? undefined : ManagementAccessRequestStateFromJSON(json['state']),
        'createdAt': !exists(json, 'createdAt') ? undefined : (new Date(json['createdAt'])),
        'updatedAt': !exists(json, 'updatedAt') ? undefined : (new Date(json['updatedAt'])),
    };
}

export function ManagementIncomingAccessRequestToJSON(value?: ManagementIncomingAccessRequest | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'id': value.id,
        'organizationName': value.organizationName,
        'serviceName': value.serviceName,
        'state': ManagementAccessRequestStateToJSON(value.state),
        'createdAt': value.createdAt === undefined ? undefined : (value.createdAt.toISOString()),
        'updatedAt': value.updatedAt === undefined ? undefined : (value.updatedAt.toISOString()),
    };
}


