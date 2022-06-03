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
import ResourcePolicy from './ResourcePolicy';

/**
 * The InvestigationCreateInput model module.
 * @module model/InvestigationCreateInput
 * @version 1.0.0
 */
class InvestigationCreateInput {
    /**
     * Constructs a new <code>InvestigationCreateInput</code>.
     * @alias module:model/InvestigationCreateInput
     */
    constructor() { 
        
        InvestigationCreateInput.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>InvestigationCreateInput</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/InvestigationCreateInput} obj Optional instance to populate.
     * @return {module:model/InvestigationCreateInput} The populated <code>InvestigationCreateInput</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new InvestigationCreateInput();

            if (data.hasOwnProperty('created_by')) {
                obj['created_by'] = ApiClient.convertToType(data['created_by'], 'String');
            }
            if (data.hasOwnProperty('data')) {
                obj['data'] = ApiClient.convertToType(data['data'], {'String': Object});
            }
            if (data.hasOwnProperty('modified_by')) {
                obj['modified_by'] = ApiClient.convertToType(data['modified_by'], 'String');
            }
            if (data.hasOwnProperty('modified_on')) {
                obj['modified_on'] = ApiClient.convertToType(data['modified_on'], 'Date');
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
            if (data.hasOwnProperty('tags')) {
                obj['tags'] = ApiClient.convertToType(data['tags'], ['String']);
            }
            if (data.hasOwnProperty('valid_from')) {
                obj['valid_from'] = ApiClient.convertToType(data['valid_from'], 'Date');
            }
            if (data.hasOwnProperty('valid_to')) {
                obj['valid_to'] = ApiClient.convertToType(data['valid_to'], 'Date');
            }
            if (data.hasOwnProperty('version')) {
                obj['version'] = ApiClient.convertToType(data['version'], 'Number');
            }
        }
        return obj;
    }


}

/**
 * UID of user who created the investigation
 * @member {String} created_by
 */
InvestigationCreateInput.prototype['created_by'] = undefined;

/**
 * JSON Object associated with the investigation
 * @member {Object.<String, Object>} data
 */
InvestigationCreateInput.prototype['data'] = undefined;

/**
 * UID of the user who last modified the investigation
 * @member {String} modified_by
 */
InvestigationCreateInput.prototype['modified_by'] = undefined;

/**
 * Date the investigation was last modified
 * @member {Date} modified_on
 */
InvestigationCreateInput.prototype['modified_on'] = undefined;

/**
 * Name of the investigation
 * @member {String} name
 */
InvestigationCreateInput.prototype['name'] = undefined;

/**
 * Resource name used for RBAC
 * @member {String} resource_name
 */
InvestigationCreateInput.prototype['resource_name'] = undefined;

/**
 * @member {module:model/ResourcePolicy} resource_policy
 */
InvestigationCreateInput.prototype['resource_policy'] = undefined;

/**
 * User supplied tags
 * @member {Array.<String>} tags
 */
InvestigationCreateInput.prototype['tags'] = undefined;

/**
 * Valid from date, the first date this object was valid
 * @member {Date} valid_from
 */
InvestigationCreateInput.prototype['valid_from'] = undefined;

/**
 * Valid to date, the date this object is valid to
 * @member {Date} valid_to
 */
InvestigationCreateInput.prototype['valid_to'] = undefined;

/**
 * Version of the investigation
 * @member {Number} version
 */
InvestigationCreateInput.prototype['version'] = undefined;






export default InvestigationCreateInput;

