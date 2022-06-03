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
import OrcApiRuntimeDetails from './OrcApiRuntimeDetails';
import ResourcePolicy from './ResourcePolicy';

/**
 * The SrcUpdateInput model module.
 * @module model/SrcUpdateInput
 * @version 1.0.0
 */
class SrcUpdateInput {
    /**
     * Constructs a new <code>SrcUpdateInput</code>.
     * @alias module:model/SrcUpdateInput
     */
    constructor() { 
        
        SrcUpdateInput.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>SrcUpdateInput</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/SrcUpdateInput} obj Optional instance to populate.
     * @return {module:model/SrcUpdateInput} The populated <code>SrcUpdateInput</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new SrcUpdateInput();

            if (data.hasOwnProperty('description')) {
                obj['description'] = ApiClient.convertToType(data['description'], 'String');
            }
            if (data.hasOwnProperty('last_data')) {
                obj['last_data'] = ApiClient.convertToType(data['last_data'], 'Date');
            }
            if (data.hasOwnProperty('last_ingest_chunk_end_time')) {
                obj['last_ingest_chunk_end_time'] = ApiClient.convertToType(data['last_ingest_chunk_end_time'], 'Date');
            }
            if (data.hasOwnProperty('last_stored_chunk_end_time')) {
                obj['last_stored_chunk_end_time'] = ApiClient.convertToType(data['last_stored_chunk_end_time'], 'Date');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('resource_name')) {
                obj['resource_name'] = ApiClient.convertToType(data['resource_name'], 'String');
            }
            if (data.hasOwnProperty('resource_policy')) {
                obj['resource_policy'] = ResourcePolicy.constructFromObject(data['resource_policy']);
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
 * User supplied description of the source
 * @member {String} description
 */
SrcUpdateInput.prototype['description'] = undefined;

/**
 * @member {Date} last_data
 */
SrcUpdateInput.prototype['last_data'] = undefined;

/**
 * The end of the last chunk ingested from the agent
 * @member {Date} last_ingest_chunk_end_time
 */
SrcUpdateInput.prototype['last_ingest_chunk_end_time'] = undefined;

/**
 * The end of the last chunk stored from the agent
 * @member {Date} last_stored_chunk_end_time
 */
SrcUpdateInput.prototype['last_stored_chunk_end_time'] = undefined;

/**
 * User supplied name of the source
 * @member {String} name
 */
SrcUpdateInput.prototype['name'] = undefined;

/**
 * Resource name used for RBAC
 * @member {String} resource_name
 */
SrcUpdateInput.prototype['resource_name'] = undefined;

/**
 * @member {module:model/ResourcePolicy} resource_policy
 */
SrcUpdateInput.prototype['resource_policy'] = undefined;

/**
 * Description of the runtime of the source
 * @member {String} runtime_description
 */
SrcUpdateInput.prototype['runtime_description'] = undefined;

/**
 * @member {module:model/OrcApiRuntimeDetails} runtime_details
 */
SrcUpdateInput.prototype['runtime_details'] = undefined;

/**
 * User supplied tags
 * @member {Array.<String>} tags
 */
SrcUpdateInput.prototype['tags'] = undefined;

/**
 * Type of source
 * @member {String} type
 */
SrcUpdateInput.prototype['type'] = undefined;

/**
 * Valid from date, the first date this object was valid
 * @member {Date} valid_from
 */
SrcUpdateInput.prototype['valid_from'] = undefined;

/**
 * Valid to date, the date this object is valid to
 * @member {Date} valid_to
 */
SrcUpdateInput.prototype['valid_to'] = undefined;






export default SrcUpdateInput;

