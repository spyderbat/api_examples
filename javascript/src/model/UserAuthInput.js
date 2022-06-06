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
 * The UserAuthInput model module.
 * @module model/UserAuthInput
 * @version 1.0.0
 */
class UserAuthInput {
    /**
     * Constructs a new <code>UserAuthInput</code>.
     * @alias module:model/UserAuthInput
     * @param email {String} User email
     * @param password {String} 
     */
    constructor(email, password) { 
        
        UserAuthInput.initialize(this, email, password);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, email, password) { 
        obj['email'] = email;
        obj['password'] = password;
    }

    /**
     * Constructs a <code>UserAuthInput</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/UserAuthInput} obj Optional instance to populate.
     * @return {module:model/UserAuthInput} The populated <code>UserAuthInput</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new UserAuthInput();

            if (data.hasOwnProperty('email')) {
                obj['email'] = ApiClient.convertToType(data['email'], 'String');
            }
            if (data.hasOwnProperty('password')) {
                obj['password'] = ApiClient.convertToType(data['password'], 'String');
            }
        }
        return obj;
    }


}

/**
 * User email
 * @member {String} email
 */
UserAuthInput.prototype['email'] = undefined;

/**
 * @member {String} password
 */
UserAuthInput.prototype['password'] = undefined;






export default UserAuthInput;
