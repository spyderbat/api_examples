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
import DaoAgentClass from './DaoAgentClass';

/**
 * The DaoAgentConfig model module.
 * @module model/DaoAgentConfig
 * @version 1.0.0
 */
class DaoAgentConfig {
    /**
     * Constructs a new <code>DaoAgentConfig</code>.
     * Initial configuration to use when creating an agent
     * @alias module:model/DaoAgentConfig
     * @param classes {Array.<module:model/DaoAgentClass>} Agent classes
     */
    constructor(classes) { 
        
        DaoAgentConfig.initialize(this, classes);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, classes) { 
        obj['classes'] = classes;
    }

    /**
     * Constructs a <code>DaoAgentConfig</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/DaoAgentConfig} obj Optional instance to populate.
     * @return {module:model/DaoAgentConfig} The populated <code>DaoAgentConfig</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new DaoAgentConfig();

            if (data.hasOwnProperty('classes')) {
                obj['classes'] = ApiClient.convertToType(data['classes'], [DaoAgentClass]);
            }
            if (data.hasOwnProperty('description')) {
                obj['description'] = ApiClient.convertToType(data['description'], 'String');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('source_tags')) {
                obj['source_tags'] = ApiClient.convertToType(data['source_tags'], ['String']);
            }
            if (data.hasOwnProperty('type')) {
                obj['type'] = ApiClient.convertToType(data['type'], 'String');
            }
        }
        return obj;
    }


}

/**
 * Agent classes
 * @member {Array.<module:model/DaoAgentClass>} classes
 */
DaoAgentConfig.prototype['classes'] = undefined;

/**
 * Description of the agent config
 * @member {String} description
 */
DaoAgentConfig.prototype['description'] = undefined;

/**
 * Name of the agent config
 * @member {String} name
 */
DaoAgentConfig.prototype['name'] = undefined;

/**
 * Tags to be applied to a source upon agent registration
 * @member {Array.<String>} source_tags
 */
DaoAgentConfig.prototype['source_tags'] = undefined;

/**
 * Type of agent config
 * @member {String} type
 */
DaoAgentConfig.prototype['type'] = undefined;






export default DaoAgentConfig;

