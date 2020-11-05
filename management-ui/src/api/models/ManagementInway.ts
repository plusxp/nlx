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
    ManagementInwayService,
    ManagementInwayServiceFromJSON,
    ManagementInwayServiceFromJSONTyped,
    ManagementInwayServiceToJSON,
} from './';

/**
 * 
 * @export
 * @interface ManagementInway
 */
export interface ManagementInway {
    /**
     * 
     * @type {string}
     * @memberof ManagementInway
     */
    name?: string;
    /**
     * 
     * @type {string}
     * @memberof ManagementInway
     */
    version?: string;
    /**
     * 
     * @type {string}
     * @memberof ManagementInway
     */
    hostname?: string;
    /**
     * 
     * @type {string}
     * @memberof ManagementInway
     */
    selfAddress?: string;
    /**
     * 
     * @type {Array<ManagementInwayService>}
     * @memberof ManagementInway
     */
    services?: Array<ManagementInwayService>;
    /**
     * 
     * @type {string}
     * @memberof ManagementInway
     */
    ipAddress?: string;
}

export function ManagementInwayFromJSON(json: any): ManagementInway {
    return ManagementInwayFromJSONTyped(json, false);
}

export function ManagementInwayFromJSONTyped(json: any, ignoreDiscriminator: boolean): ManagementInway {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'name': !exists(json, 'name') ? undefined : json['name'],
        'version': !exists(json, 'version') ? undefined : json['version'],
        'hostname': !exists(json, 'hostname') ? undefined : json['hostname'],
        'selfAddress': !exists(json, 'selfAddress') ? undefined : json['selfAddress'],
        'services': !exists(json, 'services') ? undefined : ((json['services'] as Array<any>).map(ManagementInwayServiceFromJSON)),
        'ipAddress': !exists(json, 'ipAddress') ? undefined : json['ipAddress'],
    };
}

export function ManagementInwayToJSON(value?: ManagementInway | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'name': value.name,
        'version': value.version,
        'hostname': value.hostname,
        'selfAddress': value.selfAddress,
        'services': value.services === undefined ? undefined : ((value.services as Array<any>).map(ManagementInwayServiceToJSON)),
        'ipAddress': value.ipAddress,
    };
}

