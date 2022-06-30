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
 * The DaoOrgRoles model module.
 * @module model/DaoOrgRoles
 * @version 0.1.0
 */
class DaoOrgRoles {
    /**
     * Constructs a new <code>DaoOrgRoles</code>.
     * @alias module:model/DaoOrgRoles
     */
    constructor() { 
        
        DaoOrgRoles.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>DaoOrgRoles</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/DaoOrgRoles} obj Optional instance to populate.
     * @return {module:model/DaoOrgRoles} The populated <code>DaoOrgRoles</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new DaoOrgRoles();

            if (data.hasOwnProperty('default_roles')) {
                obj['default_roles'] = ApiClient.convertToType(data['default_roles'], ['String']);
            }
            if (data.hasOwnProperty('org_uid')) {
                obj['org_uid'] = ApiClient.convertToType(data['org_uid'], 'String');
            }
        }
        return obj;
    }


}

/**
 * Default roles for the user
 * @member {Array.<String>} default_roles
 */
DaoOrgRoles.prototype['default_roles'] = undefined;

/**
 * Org UID
 * @member {String} org_uid
 */
DaoOrgRoles.prototype['org_uid'] = undefined;






export default DaoOrgRoles;

