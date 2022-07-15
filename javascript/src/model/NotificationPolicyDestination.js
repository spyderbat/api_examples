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
import NotificationPolicyDestinationSlack from './NotificationPolicyDestinationSlack';
import NotificationPolicyDestinationWebhook from './NotificationPolicyDestinationWebhook';

/**
 * The NotificationPolicyDestination model module.
 * @module model/NotificationPolicyDestination
 * @version 1.0.0
 */
class NotificationPolicyDestination {
    /**
     * Constructs a new <code>NotificationPolicyDestination</code>.
     * A notification policy destination, containing one and only one of the available types
     * @alias module:model/NotificationPolicyDestination
     */
    constructor() { 
        
        NotificationPolicyDestination.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>NotificationPolicyDestination</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/NotificationPolicyDestination} obj Optional instance to populate.
     * @return {module:model/NotificationPolicyDestination} The populated <code>NotificationPolicyDestination</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new NotificationPolicyDestination();

            if (data.hasOwnProperty('data')) {
                obj['data'] = ApiClient.convertToType(data['data'], Object);
            }
            if (data.hasOwnProperty('description')) {
                obj['description'] = ApiClient.convertToType(data['description'], 'String');
            }
            if (data.hasOwnProperty('email')) {
                obj['email'] = ApiClient.convertToType(data['email'], ['String']);
            }
            if (data.hasOwnProperty('org_uid')) {
                obj['org_uid'] = ApiClient.convertToType(data['org_uid'], 'String');
            }
            if (data.hasOwnProperty('slack')) {
                obj['slack'] = NotificationPolicyDestinationSlack.constructFromObject(data['slack']);
            }
            if (data.hasOwnProperty('users')) {
                obj['users'] = ApiClient.convertToType(data['users'], ['String']);
            }
            if (data.hasOwnProperty('webhook')) {
                obj['webhook'] = NotificationPolicyDestinationWebhook.constructFromObject(data['webhook']);
            }
        }
        return obj;
    }


}

/**
 * UI-supplied data
 * @member {Object} data
 */
NotificationPolicyDestination.prototype['data'] = undefined;

/**
 * @member {String} description
 */
NotificationPolicyDestination.prototype['description'] = undefined;

/**
 * @member {Array.<String>} email
 */
NotificationPolicyDestination.prototype['email'] = undefined;

/**
 * @member {String} org_uid
 */
NotificationPolicyDestination.prototype['org_uid'] = undefined;

/**
 * @member {module:model/NotificationPolicyDestinationSlack} slack
 */
NotificationPolicyDestination.prototype['slack'] = undefined;

/**
 * @member {Array.<String>} users
 */
NotificationPolicyDestination.prototype['users'] = undefined;

/**
 * @member {module:model/NotificationPolicyDestinationWebhook} webhook
 */
NotificationPolicyDestination.prototype['webhook'] = undefined;






export default NotificationPolicyDestination;

