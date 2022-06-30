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
import OrcApiRuntimeDetails from './OrcApiRuntimeDetails';

/**
 * The SrcCreateInput model module.
 * @module model/SrcCreateInput
 * @version 0.1.0
 */
class SrcCreateInput {
    /**
     * Constructs a new <code>SrcCreateInput</code>.
     * @alias module:model/SrcCreateInput
     */
    constructor() { 
        
        SrcCreateInput.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>SrcCreateInput</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/SrcCreateInput} obj Optional instance to populate.
     * @return {module:model/SrcCreateInput} The populated <code>SrcCreateInput</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new SrcCreateInput();

            if (data.hasOwnProperty('description')) {
                obj['description'] = ApiClient.convertToType(data['description'], 'String');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('runtime_description')) {
                obj['runtime_description'] = ApiClient.convertToType(data['runtime_description'], 'String');
            }
            if (data.hasOwnProperty('runtime_details')) {
                obj['runtime_details'] = OrcApiRuntimeDetails.constructFromObject(data['runtime_details']);
            }
            if (data.hasOwnProperty('tags')) {
                obj['tags'] = ApiClient.convertToType(data['tags'], ['String']);
            }
            if (data.hasOwnProperty('type')) {
                obj['type'] = ApiClient.convertToType(data['type'], 'String');
            }
            if (data.hasOwnProperty('uid')) {
                obj['uid'] = ApiClient.convertToType(data['uid'], 'String');
            }
        }
        return obj;
    }


}

/**
 * User supplied description of the source
 * @member {String} description
 */
SrcCreateInput.prototype['description'] = undefined;

/**
 * User supplied name of the source
 * @member {String} name
 */
SrcCreateInput.prototype['name'] = undefined;

/**
 * Description of the runtime of the source
 * @member {String} runtime_description
 */
SrcCreateInput.prototype['runtime_description'] = undefined;

/**
 * @member {module:model/OrcApiRuntimeDetails} runtime_details
 */
SrcCreateInput.prototype['runtime_details'] = undefined;

/**
 * User supplied tags
 * @member {Array.<String>} tags
 */
SrcCreateInput.prototype['tags'] = undefined;

/**
 * Type of source
 * @member {String} type
 */
SrcCreateInput.prototype['type'] = undefined;

/**
 * The UID of the source
 * @member {String} uid
 */
SrcCreateInput.prototype['uid'] = undefined;






export default SrcCreateInput;

