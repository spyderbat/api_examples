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
import ResourcePolicy from './ResourcePolicy';

/**
 * The Org model module.
 * @module model/Org
 * @version 1.0.0
 */
class Org {
    /**
     * Constructs a new <code>Org</code>.
     * @alias module:model/Org
     * @param name {String} Name of the organization
     */
    constructor(name) { 
        
        Org.initialize(this, name);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, name) { 
        obj['name'] = name;
    }

    /**
     * Constructs a <code>Org</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/Org} obj Optional instance to populate.
     * @return {module:model/Org} The populated <code>Org</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new Org();

            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('org_type_uid')) {
                obj['org_type_uid'] = ApiClient.convertToType(data['org_type_uid'], 'String');
            }
            if (data.hasOwnProperty('owner_email')) {
                obj['owner_email'] = ApiClient.convertToType(data['owner_email'], 'String');
            }
            if (data.hasOwnProperty('owner_uid')) {
                obj['owner_uid'] = ApiClient.convertToType(data['owner_uid'], 'String');
            }
            if (data.hasOwnProperty('resource_name')) {
                obj['resource_name'] = ApiClient.convertToType(data['resource_name'], 'String');
            }
            if (data.hasOwnProperty('resource_policy')) {
                obj['resource_policy'] = ResourcePolicy.constructFromObject(data['resource_policy']);
            }
            if (data.hasOwnProperty('tags')) {
                obj['tags'] = ApiClient.convertToType(data['tags'], ['String']);
            }
            if (data.hasOwnProperty('uid')) {
                obj['uid'] = ApiClient.convertToType(data['uid'], 'String');
            }
            if (data.hasOwnProperty('valid_from')) {
                obj['valid_from'] = ApiClient.convertToType(data['valid_from'], 'Date');
            }
            if (data.hasOwnProperty('valid_to')) {
                obj['valid_to'] = ApiClient.convertToType(data['valid_to'], 'Date');
            }
        }
        return obj;
    }


}

/**
 * Name of the organization
 * @member {String} name
 */
Org.prototype['name'] = undefined;

/**
 * Organization Type
 * @member {String} org_type_uid
 */
Org.prototype['org_type_uid'] = undefined;

/**
 * The email address of the user who owns this org
 * @member {String} owner_email
 */
Org.prototype['owner_email'] = undefined;

/**
 * The user UID who owns this organization
 * @member {String} owner_uid
 */
Org.prototype['owner_uid'] = undefined;

/**
 * Resource name utilized by RBAC
 * @member {String} resource_name
 */
Org.prototype['resource_name'] = undefined;

/**
 * @member {module:model/ResourcePolicy} resource_policy
 */
Org.prototype['resource_policy'] = undefined;

/**
 * User supplied tags
 * @member {Array.<String>} tags
 */
Org.prototype['tags'] = undefined;

/**
 * Org UID
 * @member {String} uid
 */
Org.prototype['uid'] = undefined;

/**
 * Valid from date, the first date this object was valid
 * @member {Date} valid_from
 */
Org.prototype['valid_from'] = undefined;

/**
 * Valid to date, the date this object is valid to
 * @member {Date} valid_to
 */
Org.prototype['valid_to'] = undefined;






export default Org;

