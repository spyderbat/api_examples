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
 * The RstreamTimeRange model module.
 * @module model/RstreamTimeRange
 * @version 1.0.0
 */
class RstreamTimeRange {
    /**
     * Constructs a new <code>RstreamTimeRange</code>.
     * returns true if the property is a time and is start&lt;&#x3D;time&lt;end
     * @alias module:model/RstreamTimeRange
     */
    constructor() { 
        
        RstreamTimeRange.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>RstreamTimeRange</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/RstreamTimeRange} obj Optional instance to populate.
     * @return {module:model/RstreamTimeRange} The populated <code>RstreamTimeRange</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new RstreamTimeRange();

            if (data.hasOwnProperty('end_time')) {
                obj['end_time'] = ApiClient.convertToType(data['end_time'], 'Number');
            }
            if (data.hasOwnProperty('start_time')) {
                obj['start_time'] = ApiClient.convertToType(data['start_time'], 'Number');
            }
        }
        return obj;
    }


}

/**
 * @member {Number} end_time
 */
RstreamTimeRange.prototype['end_time'] = undefined;

/**
 * @member {Number} start_time
 */
RstreamTimeRange.prototype['start_time'] = undefined;






export default RstreamTimeRange;
