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
 * The NotificationPolicyDestinationSlack model module.
 * @module model/NotificationPolicyDestinationSlack
 * @version 1.0.0
 */
class NotificationPolicyDestinationSlack {
    /**
     * Constructs a new <code>NotificationPolicyDestinationSlack</code>.
     * @alias module:model/NotificationPolicyDestinationSlack
     * @param url {String} 
     */
    constructor(url) { 
        
        NotificationPolicyDestinationSlack.initialize(this, url);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, url) { 
        obj['url'] = url;
    }

    /**
     * Constructs a <code>NotificationPolicyDestinationSlack</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/NotificationPolicyDestinationSlack} obj Optional instance to populate.
     * @return {module:model/NotificationPolicyDestinationSlack} The populated <code>NotificationPolicyDestinationSlack</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new NotificationPolicyDestinationSlack();

            if (data.hasOwnProperty('url')) {
                obj['url'] = ApiClient.convertToType(data['url'], 'String');
            }
        }
        return obj;
    }


}

/**
 * @member {String} url
 */
NotificationPolicyDestinationSlack.prototype['url'] = undefined;






export default NotificationPolicyDestinationSlack;
