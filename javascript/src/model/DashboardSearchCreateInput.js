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
 * The DashboardSearchCreateInput model module.
 * @module model/DashboardSearchCreateInput
 * @version 1.0.0
 */
class DashboardSearchCreateInput {
    /**
     * Constructs a new <code>DashboardSearchCreateInput</code>.
     * @alias module:model/DashboardSearchCreateInput
     */
    constructor() { 
        
        DashboardSearchCreateInput.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>DashboardSearchCreateInput</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/DashboardSearchCreateInput} obj Optional instance to populate.
     * @return {module:model/DashboardSearchCreateInput} The populated <code>DashboardSearchCreateInput</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new DashboardSearchCreateInput();

            if (data.hasOwnProperty('')) {
                obj[''] = ApiClient.convertToType(data[''], 'Date');
            }
            if (data.hasOwnProperty('data')) {
                obj['data'] = ApiClient.convertToType(data['data'], {'String': Object});
            }
            if (data.hasOwnProperty('description')) {
                obj['description'] = ApiClient.convertToType(data['description'], 'String');
            }
            if (data.hasOwnProperty('notify')) {
                obj['notify'] = ApiClient.convertToType(data['notify'], 'Boolean');
            }
            if (data.hasOwnProperty('notify_frequency')) {
                obj['notify_frequency'] = ApiClient.convertToType(data['notify_frequency'], 'Number');
            }
            if (data.hasOwnProperty('search')) {
                obj['search'] = ApiClient.convertToType(data['search'], 'String');
            }
            if (data.hasOwnProperty('tags')) {
                obj['tags'] = ApiClient.convertToType(data['tags'], ['String']);
            }
        }
        return obj;
    }


}

/**
 * Time last notification window was completed
 * @member {Date} 
 */
DashboardSearchCreateInput.prototype[''] = undefined;

/**
 * UI supplied JSON object
 * @member {Object.<String, Object>} data
 */
DashboardSearchCreateInput.prototype['data'] = undefined;

/**
 * Description of query (max 64 chars)
 * @member {String} description
 */
DashboardSearchCreateInput.prototype['description'] = undefined;

/**
 * Are notifications generated from this search
 * @member {Boolean} notify
 */
DashboardSearchCreateInput.prototype['notify'] = undefined;

/**
 * Frequency of notification in seconds. One of 300, 3600, or 86400.
 * @member {Number} notify_frequency
 */
DashboardSearchCreateInput.prototype['notify_frequency'] = undefined;

/**
 * Search query to run
 * @member {String} search
 */
DashboardSearchCreateInput.prototype['search'] = undefined;

/**
 * User supplied tags
 * @member {Array.<String>} tags
 */
DashboardSearchCreateInput.prototype['tags'] = undefined;






export default DashboardSearchCreateInput;
