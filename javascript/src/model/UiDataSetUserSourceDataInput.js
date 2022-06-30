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
 * The UiDataSetUserSourceDataInput model module.
 * @module model/UiDataSetUserSourceDataInput
 * @version 0.1.0
 */
class UiDataSetUserSourceDataInput {
    /**
     * Constructs a new <code>UiDataSetUserSourceDataInput</code>.
     * @alias module:model/UiDataSetUserSourceDataInput
     */
    constructor() { 
        
        UiDataSetUserSourceDataInput.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>UiDataSetUserSourceDataInput</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/UiDataSetUserSourceDataInput} obj Optional instance to populate.
     * @return {module:model/UiDataSetUserSourceDataInput} The populated <code>UiDataSetUserSourceDataInput</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new UiDataSetUserSourceDataInput();

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
UiDataSetUserSourceDataInput.prototype[''] = undefined;

/**
 * UI supplied JSON object
 * @member {Object.<String, Object>} data
 */
UiDataSetUserSourceDataInput.prototype['data'] = undefined;

/**
 * The time of las use of the key; only updated every 5 minutes
 * @member {Date} last_used
 */
UiDataSetUserSourceDataInput.prototype['last_used'] = undefined;

/**
 * User supplied tags
 * @member {Array.<String>} tags
 */
UiDataSetUserSourceDataInput.prototype['tags'] = undefined;

/**
 * UID for the UIData
 * @member {String} uid
 */
UiDataSetUserSourceDataInput.prototype['uid'] = undefined;

/**
 * Valid from is the creation date or first date when the API key became valid
 * @member {Date} valid_from
 */
UiDataSetUserSourceDataInput.prototype['valid_from'] = undefined;

/**
 * Valid to is the expiration date or the last date this API key will be valid
 * @member {Date} valid_to
 */
UiDataSetUserSourceDataInput.prototype['valid_to'] = undefined;






export default UiDataSetUserSourceDataInput;

