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
import Source from './Source';

/**
 * The ApiSOARListHandlerOutput model module.
 * @module model/ApiSOARListHandlerOutput
 * @version 1.0.0
 */
class ApiSOARListHandlerOutput {
    /**
     * Constructs a new <code>ApiSOARListHandlerOutput</code>.
     * @alias module:model/ApiSOARListHandlerOutput
     */
    constructor() { 
        
        ApiSOARListHandlerOutput.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>ApiSOARListHandlerOutput</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ApiSOARListHandlerOutput} obj Optional instance to populate.
     * @return {module:model/ApiSOARListHandlerOutput} The populated <code>ApiSOARListHandlerOutput</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ApiSOARListHandlerOutput();

            if (data.hasOwnProperty('investigate_source_url')) {
                obj['investigate_source_url'] = ApiClient.convertToType(data['investigate_source_url'], 'String');
            }
            if (data.hasOwnProperty('source')) {
                obj['source'] = Source.constructFromObject(data['source']);
            }
        }
        return obj;
    }


}

/**
 * @member {String} investigate_source_url
 */
ApiSOARListHandlerOutput.prototype['investigate_source_url'] = undefined;

/**
 * @member {module:model/Source} source
 */
ApiSOARListHandlerOutput.prototype['source'] = undefined;






export default ApiSOARListHandlerOutput;

