/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API Patient
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntPatient,
    EntPatientFromJSON,
    EntPatientFromJSONTyped,
    EntPatientToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntJobEdges
 */
export interface EntJobEdges {
    /**
     * Patients holds the value of the patients edge.
     * @type {Array<EntPatient>}
     * @memberof EntJobEdges
     */
    patients?: Array<EntPatient>;
}

export function EntJobEdgesFromJSON(json: any): EntJobEdges {
    return EntJobEdgesFromJSONTyped(json, false);
}

export function EntJobEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntJobEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'patients': !exists(json, 'patients') ? undefined : ((json['patients'] as Array<any>).map(EntPatientFromJSON)),
    };
}

export function EntJobEdgesToJSON(value?: EntJobEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'patients': value.patients === undefined ? undefined : ((value.patients as Array<any>).map(EntPatientToJSON)),
    };
}

