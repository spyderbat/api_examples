/**
 * Spyderbat API UI & Public APIs
 * Restful APIs for use by UI & customers.
 *
 * The version of the OpenAPI document: 1.0.0
 * Contact: apisupport@spyderbat.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

import ApiClient from '../ApiClient';

/**
 * The UiDataSetUserOrgDataInput model module.
 * @module model/UiDataSetUserOrgDataInput
 * @version 1.0.0
 */
class UiDataSetUserOrgDataInput {
    /**
     * Constructs a new <code>UiDataSetUserOrgDataInput</code>.
     * @alias module:model/UiDataSetUserOrgDataInput
     */
    constructor() { 
        
        UiDataSetUserOrgDataInput.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>UiDataSetUserOrgDataInput</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/UiDataSetUserOrgDataInput} obj Optional instance to populate.
     * @return {module:model/UiDataSetUserOrgDataInput} The populated <code>UiDataSetUserOrgDataInput</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new UiDataSetUserOrgDataInput();

            if (data.hasOwnProperty('')) {
                obj[''] = ApiClient.convertToType(data[''], 'String');
            }
            if (data.hasOwnProperty('data')) {
                obj['data'] = ApiClient.convertToType(data['data'], {'String': Object});
            }
            if (data.hasOwnProperty('last_used')) {
                obj['last_used'] = ApiClient.convertToType(data['last_used'], 'Date');
            }
            if (data.hasOwnProperty('tags')) {
                obj['tags'] = ApiClient.convertToType(data['tags'], ['String']);
            }
            if (data.hasOwnProperty('uid')) {
                obj['uid'] = ApiClient.convertToType(data['uid'], 'String');
            }
            if (data.hasOwnProperty('valid_from')) {
                obj['valid_from'] = ApiClient.convertToType(data['valid_from'], 'Date');
            }
            if (data.hasOwnProperty('valid_to')) {
                obj['valid_to'] = ApiClient.convertToType(data['valid_to'], 'Date');
            }
        }
        return obj;
    }


}

/**
 * @member {String} 
 */
UiDataSetUserOrgDataInput.prototype[''] = undefined;

/**
 * UI supplied JSON object
 * @member {Object.<String, Object>} data
 */
UiDataSetUserOrgDataInput.prototype['data'] = undefined;

/**
 * The time of las use of the key; only updated every 5 minutes
 * @member {Date} last_used
 */
UiDataSetUserOrgDataInput.prototype['last_used'] = undefined;

/**
 * User supplied tags
 * @member {Array.<String>} tags
 */
UiDataSetUserOrgDataInput.prototype['tags'] = undefined;

/**
 * UID for the UIData
 * @member {String} uid
 */
UiDataSetUserOrgDataInput.prototype['uid'] = undefined;

/**
 * Valid from is the creation date or first date when the API key became valid
 * @member {Date} valid_from
 */
UiDataSetUserOrgDataInput.prototype['valid_from'] = undefined;

/**
 * Valid to is the expiration date or the last date this API key will be valid
 * @member {Date} valid_to
 */
UiDataSetUserOrgDataInput.prototype['valid_to'] = undefined;






export default UiDataSetUserOrgDataInput;

