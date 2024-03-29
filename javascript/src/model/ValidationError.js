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
 * The ValidationError model module.
 * @module model/ValidationError
 * @version 1.0.0
 */
class ValidationError {
    /**
     * Constructs a new <code>ValidationError</code>.
     * @alias module:model/ValidationError
     */
    constructor() { 
        
        ValidationError.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>ValidationError</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ValidationError} obj Optional instance to populate.
     * @return {module:model/ValidationError} The populated <code>ValidationError</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ValidationError();

            if (data.hasOwnProperty('err_msg')) {
                obj['err_msg'] = ApiClient.convertToType(data['err_msg'], 'String');
            }
            if (data.hasOwnProperty('field')) {
                obj['field'] = ApiClient.convertToType(data['field'], 'String');
            }
            if (data.hasOwnProperty('property')) {
                obj['property'] = ApiClient.convertToType(data['property'], 'String');
            }
            if (data.hasOwnProperty('tags')) {
                obj['tags'] = ApiClient.convertToType(data['tags'], 'String');
            }
        }
        return obj;
    }


}

/**
 * Message regarding the validation failure
 * @member {String} err_msg
 */
ValidationError.prototype['err_msg'] = undefined;

/**
 * Field name which failed validation
 * @member {String} field
 */
ValidationError.prototype['field'] = undefined;

/**
 * JSON property name of the field which failed validation
 * @member {String} property
 */
ValidationError.prototype['property'] = undefined;

/**
 * Validation tag which failed
 * @member {String} tags
 */
ValidationError.prototype['tags'] = undefined;






export default ValidationError;

