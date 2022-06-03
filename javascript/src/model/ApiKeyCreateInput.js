/**
 * Spyderbat API UI & Public APIs
 * Restful APIs for use by UI & customers.
 *
 * The version of the OpenAPI document: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

import ApiClient from '../ApiClient';

/**
 * The ApiKeyCreateInput model module.
 * @module model/ApiKeyCreateInput
 * @version 1.0.0
 */
class ApiKeyCreateInput {
    /**
     * Constructs a new <code>ApiKeyCreateInput</code>.
     * @alias module:model/ApiKeyCreateInput
     * @param validTo {Date} When the key should expire, in a RFC3339 formated string
     */
    constructor(validTo) { 
        
        ApiKeyCreateInput.initialize(this, validTo);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, validTo) { 
        obj['valid_to'] = validTo;
    }

    /**
     * Constructs a <code>ApiKeyCreateInput</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ApiKeyCreateInput} obj Optional instance to populate.
     * @return {module:model/ApiKeyCreateInput} The populated <code>ApiKeyCreateInput</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ApiKeyCreateInput();

            if (data.hasOwnProperty('description')) {
                obj['description'] = ApiClient.convertToType(data['description'], 'String');
            }
            if (data.hasOwnProperty('valid_to')) {
                obj['valid_to'] = ApiClient.convertToType(data['valid_to'], 'Date');
            }
        }
        return obj;
    }


}

/**
 * Description of the API key
 * @member {String} description
 */
ApiKeyCreateInput.prototype['description'] = undefined;

/**
 * When the key should expire, in a RFC3339 formated string
 * @member {Date} valid_to
 */
ApiKeyCreateInput.prototype['valid_to'] = undefined;






export default ApiKeyCreateInput;

