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
 * The OrgTestNotificationTargetInput model module.
 * @module model/OrgTestNotificationTargetInput
 * @version 1.0.0
 */
class OrgTestNotificationTargetInput {
    /**
     * Constructs a new <code>OrgTestNotificationTargetInput</code>.
     * @alias module:model/OrgTestNotificationTargetInput
     */
    constructor() { 
        
        OrgTestNotificationTargetInput.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>OrgTestNotificationTargetInput</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/OrgTestNotificationTargetInput} obj Optional instance to populate.
     * @return {module:model/OrgTestNotificationTargetInput} The populated <code>OrgTestNotificationTargetInput</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new OrgTestNotificationTargetInput();

            if (data.hasOwnProperty('target')) {
                obj['target'] = ApiClient.convertToType(data['target'], 'String');
            }
        }
        return obj;
    }


}

/**
 * @member {String} target
 */
OrgTestNotificationTargetInput.prototype['target'] = undefined;






export default OrgTestNotificationTargetInput;

