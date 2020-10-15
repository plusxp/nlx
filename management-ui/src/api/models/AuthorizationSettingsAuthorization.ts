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
/**
 * 
 * @export
 * @interface AuthorizationSettingsAuthorization
 */
export interface AuthorizationSettingsAuthorization {
    /**
     * 
     * @type {string}
     * @memberof AuthorizationSettingsAuthorization
     */
    organizationName?: string;
    /**
     * 
     * @type {string}
     * @memberof AuthorizationSettingsAuthorization
     */
    publicKeyHash?: string;
}

export function AuthorizationSettingsAuthorizationFromJSON(json: any): AuthorizationSettingsAuthorization {
    return AuthorizationSettingsAuthorizationFromJSONTyped(json, false);
}

export function AuthorizationSettingsAuthorizationFromJSONTyped(json: any, ignoreDiscriminator: boolean): AuthorizationSettingsAuthorization {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'organizationName': !exists(json, 'organizationName') ? undefined : json['organizationName'],
        'publicKeyHash': !exists(json, 'publicKeyHash') ? undefined : json['publicKeyHash'],
    };
}

export function AuthorizationSettingsAuthorizationToJSON(value?: AuthorizationSettingsAuthorization | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'organizationName': value.organizationName,
        'publicKeyHash': value.publicKeyHash,
    };
}


