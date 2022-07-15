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
 * The SessionOrgTypeMaxLimit model module.
 * @module model/SessionOrgTypeMaxLimit
 * @version 1.0.0
 */
class SessionOrgTypeMaxLimit {
    /**
     * Constructs a new <code>SessionOrgTypeMaxLimit</code>.
     * @alias module:model/SessionOrgTypeMaxLimit
     */
    constructor() { 
        
        SessionOrgTypeMaxLimit.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>SessionOrgTypeMaxLimit</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/SessionOrgTypeMaxLimit} obj Optional instance to populate.
     * @return {module:model/SessionOrgTypeMaxLimit} The populated <code>SessionOrgTypeMaxLimit</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new SessionOrgTypeMaxLimit();

            if (data.hasOwnProperty('current_value')) {
                obj['current_value'] = ApiClient.convertToType(data['current_value'], 'Number');
            }
            if (data.hasOwnProperty('description')) {
                obj['description'] = ApiClient.convertToType(data['description'], 'String');
            }
            if (data.hasOwnProperty('limit_exceeded')) {
                obj['limit_exceeded'] = ApiClient.convertToType(data['limit_exceeded'], 'Boolean');
            }
            if (data.hasOwnProperty('limit_value')) {
                obj['limit_value'] = ApiClient.convertToType(data['limit_value'], 'Number');
            }
            if (data.hasOwnProperty('remaining_capacity')) {
                obj['remaining_capacity'] = ApiClient.convertToType(data['remaining_capacity'], 'Number');
            }
            if (data.hasOwnProperty('time_window_in_seconds')) {
                obj['time_window_in_seconds'] = ApiClient.convertToType(data['time_window_in_seconds'], 'Number');
            }
        }
        return obj;
    }


}

/**
 * What is the current value
 * @member {Number} current_value
 */
SessionOrgTypeMaxLimit.prototype['current_value'] = undefined;

/**
 * Description of the limit
 * @member {String} description
 */
SessionOrgTypeMaxLimit.prototype['description'] = undefined;

/**
 * Has the limit been met or exceeded
 * @member {Boolean} limit_exceeded
 */
SessionOrgTypeMaxLimit.prototype['limit_exceeded'] = undefined;

/**
 * What is the max limit value
 * @member {Number} limit_value
 */
SessionOrgTypeMaxLimit.prototype['limit_value'] = undefined;

/**
 * How many items can be added to the current value before it is exceeded
 * @member {Number} remaining_capacity
 */
SessionOrgTypeMaxLimit.prototype['remaining_capacity'] = undefined;

/**
 * The time window in seconds that the limit is calcuated on
 * @member {Number} time_window_in_seconds
 */
SessionOrgTypeMaxLimit.prototype['time_window_in_seconds'] = undefined;






export default SessionOrgTypeMaxLimit;

