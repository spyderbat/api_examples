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
 * The DaoAgentClass model module.
 * @module model/DaoAgentClass
 * @version 1.0.0
 */
class DaoAgentClass {
    /**
     * Constructs a new <code>DaoAgentClass</code>.
     * @alias module:model/DaoAgentClass
     * @param name {String} 
     */
    constructor(name) { 
        
        DaoAgentClass.initialize(this, name);
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
     * Constructs a <code>DaoAgentClass</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/DaoAgentClass} obj Optional instance to populate.
     * @return {module:model/DaoAgentClass} The populated <code>DaoAgentClass</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new DaoAgentClass();

            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('rbac_roles')) {
                obj['rbac_roles'] = ApiClient.convertToType(data['rbac_roles'], ['String']);
            }
        }
        return obj;
    }


}

/**
 * @member {String} name
 */
DaoAgentClass.prototype['name'] = undefined;

/**
 * @member {Array.<String>} rbac_roles
 */
DaoAgentClass.prototype['rbac_roles'] = undefined;






export default DaoAgentClass;

