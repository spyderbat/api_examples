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
 * The DashboardSearchUpdateInput model module.
 * @module model/DashboardSearchUpdateInput
 * @version 1.0.0
 */
class DashboardSearchUpdateInput {
    /**
     * Constructs a new <code>DashboardSearchUpdateInput</code>.
     * @alias module:model/DashboardSearchUpdateInput
     */
    constructor() { 
        
        DashboardSearchUpdateInput.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>DashboardSearchUpdateInput</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/DashboardSearchUpdateInput} obj Optional instance to populate.
     * @return {module:model/DashboardSearchUpdateInput} The populated <code>DashboardSearchUpdateInput</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new DashboardSearchUpdateInput();

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
DashboardSearchUpdateInput.prototype[''] = undefined;

/**
 * UI supplied JSON object
 * @member {Object.<String, Object>} data
 */
DashboardSearchUpdateInput.prototype['data'] = undefined;

/**
 * Description of query (max 64 chars)
 * @member {String} description
 */
DashboardSearchUpdateInput.prototype['description'] = undefined;

/**
 * Are notifications generated from this search
 * @member {Boolean} notify
 */
DashboardSearchUpdateInput.prototype['notify'] = undefined;

/**
 * Frequency of notification in seconds. One of 300, 3600, or 86400.
 * @member {Number} notify_frequency
 */
DashboardSearchUpdateInput.prototype['notify_frequency'] = undefined;

/**
 * Search query to run
 * @member {String} search
 */
DashboardSearchUpdateInput.prototype['search'] = undefined;

/**
 * User supplied tags
 * @member {Array.<String>} tags
 */
DashboardSearchUpdateInput.prototype['tags'] = undefined;






export default DashboardSearchUpdateInput;

