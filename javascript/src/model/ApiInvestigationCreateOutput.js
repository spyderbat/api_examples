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
 * The ApiInvestigationCreateOutput model module.
 * @module model/ApiInvestigationCreateOutput
 * @version 0.1.0
 */
class ApiInvestigationCreateOutput {
    /**
     * Constructs a new <code>ApiInvestigationCreateOutput</code>.
     * @alias module:model/ApiInvestigationCreateOutput
     */
    constructor() { 
        
        ApiInvestigationCreateOutput.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>ApiInvestigationCreateOutput</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ApiInvestigationCreateOutput} obj Optional instance to populate.
     * @return {module:model/ApiInvestigationCreateOutput} The populated <code>ApiInvestigationCreateOutput</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ApiInvestigationCreateOutput();

            if (data.hasOwnProperty('uid')) {
                obj['uid'] = ApiClient.convertToType(data['uid'], 'String');
            }
        }
        return obj;
    }


}

/**
 * @member {String} uid
 */
ApiInvestigationCreateOutput.prototype['uid'] = undefined;






export default ApiInvestigationCreateOutput;

