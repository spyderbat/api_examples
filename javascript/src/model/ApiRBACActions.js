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
import RBACAction from './RBACAction';

/**
 * The ApiRBACActions model module.
 * @module model/ApiRBACActions
 * @version 1.0.0
 */
class ApiRBACActions {
    /**
     * Constructs a new <code>ApiRBACActions</code>.
     * @alias module:model/ApiRBACActions
     */
    constructor() { 
        
        ApiRBACActions.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>ApiRBACActions</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ApiRBACActions} obj Optional instance to populate.
     * @return {module:model/ApiRBACActions} The populated <code>ApiRBACActions</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ApiRBACActions();

            if (data.hasOwnProperty('actions')) {
                obj['actions'] = ApiClient.convertToType(data['actions'], [RBACAction]);
            }
        }
        return obj;
    }


}

/**
 * @member {Array.<module:model/RBACAction>} actions
 */
ApiRBACActions.prototype['actions'] = undefined;






export default ApiRBACActions;

