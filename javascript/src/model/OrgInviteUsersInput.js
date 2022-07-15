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
 * The OrgInviteUsersInput model module.
 * @module model/OrgInviteUsersInput
 * @version 1.0.0
 */
class OrgInviteUsersInput {
    /**
     * Constructs a new <code>OrgInviteUsersInput</code>.
     * @alias module:model/OrgInviteUsersInput
     */
    constructor() { 
        
        OrgInviteUsersInput.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>OrgInviteUsersInput</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/OrgInviteUsersInput} obj Optional instance to populate.
     * @return {module:model/OrgInviteUsersInput} The populated <code>OrgInviteUsersInput</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new OrgInviteUsersInput();

            if (data.hasOwnProperty('emails')) {
                obj['emails'] = ApiClient.convertToType(data['emails'], ['String']);
            }
            if (data.hasOwnProperty('roles')) {
                obj['roles'] = ApiClient.convertToType(data['roles'], ['String']);
            }
        }
        return obj;
    }


}

/**
 * @member {Array.<String>} emails
 */
OrgInviteUsersInput.prototype['emails'] = undefined;

/**
 * @member {Array.<String>} roles
 */
OrgInviteUsersInput.prototype['roles'] = undefined;






export default OrgInviteUsersInput;

