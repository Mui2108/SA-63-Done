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
    EntGender,
    EntGenderFromJSON,
    EntGenderFromJSONTyped,
    EntGenderToJSON,
    EntJob,
    EntJobFromJSON,
    EntJobFromJSONTyped,
    EntJobToJSON,
    EntTitle,
    EntTitleFromJSON,
    EntTitleFromJSONTyped,
    EntTitleToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntPatientEdges
 */
export interface EntPatientEdges {
    /**
     * 
     * @type {EntGender}
     * @memberof EntPatientEdges
     */
    gender?: EntGender;
    /**
     * 
     * @type {EntJob}
     * @memberof EntPatientEdges
     */
    job?: EntJob;
    /**
     * 
     * @type {EntTitle}
     * @memberof EntPatientEdges
     */
    title?: EntTitle;
}

export function EntPatientEdgesFromJSON(json: any): EntPatientEdges {
    return EntPatientEdgesFromJSONTyped(json, false);
}

export function EntPatientEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntPatientEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'gender': !exists(json, 'gender') ? undefined : EntGenderFromJSON(json['gender']),
        'job': !exists(json, 'job') ? undefined : EntJobFromJSON(json['job']),
        'title': !exists(json, 'title') ? undefined : EntTitleFromJSON(json['title']),
    };
}

export function EntPatientEdgesToJSON(value?: EntPatientEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'gender': EntGenderToJSON(value.gender),
        'job': EntJobToJSON(value.job),
        'title': EntTitleToJSON(value.title),
    };
}


