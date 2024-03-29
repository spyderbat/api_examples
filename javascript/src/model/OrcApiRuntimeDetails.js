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
 * The OrcApiRuntimeDetails model module.
 * @module model/OrcApiRuntimeDetails
 * @version 1.0.0
 */
class OrcApiRuntimeDetails {
    /**
     * Constructs a new <code>OrcApiRuntimeDetails</code>.
     * Runtime details
     * @alias module:model/OrcApiRuntimeDetails
     */
    constructor() { 
        
        OrcApiRuntimeDetails.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>OrcApiRuntimeDetails</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/OrcApiRuntimeDetails} obj Optional instance to populate.
     * @return {module:model/OrcApiRuntimeDetails} The populated <code>OrcApiRuntimeDetails</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new OrcApiRuntimeDetails();

            if (data.hasOwnProperty('agent_arch')) {
                obj['agent_arch'] = ApiClient.convertToType(data['agent_arch'], 'String');
            }
            if (data.hasOwnProperty('agent_version')) {
                obj['agent_version'] = ApiClient.convertToType(data['agent_version'], 'String');
            }
            if (data.hasOwnProperty('boot_time')) {
                obj['boot_time'] = ApiClient.convertToType(data['boot_time'], 'Number');
            }
            if (data.hasOwnProperty('cloud_image_id')) {
                obj['cloud_image_id'] = ApiClient.convertToType(data['cloud_image_id'], 'String');
            }
            if (data.hasOwnProperty('cloud_instance_id')) {
                obj['cloud_instance_id'] = ApiClient.convertToType(data['cloud_instance_id'], 'String');
            }
            if (data.hasOwnProperty('cloud_region')) {
                obj['cloud_region'] = ApiClient.convertToType(data['cloud_region'], 'String');
            }
            if (data.hasOwnProperty('cloud_tags')) {
                obj['cloud_tags'] = ApiClient.convertToType(data['cloud_tags'], ['String']);
            }
            if (data.hasOwnProperty('cloud_type')) {
                obj['cloud_type'] = ApiClient.convertToType(data['cloud_type'], 'String');
            }
            if (data.hasOwnProperty('cpu_cores')) {
                obj['cpu_cores'] = ApiClient.convertToType(data['cpu_cores'], 'Number');
            }
            if (data.hasOwnProperty('cpu_make')) {
                obj['cpu_make'] = ApiClient.convertToType(data['cpu_make'], 'String');
            }
            if (data.hasOwnProperty('cpu_model')) {
                obj['cpu_model'] = ApiClient.convertToType(data['cpu_model'], 'String');
            }
            if (data.hasOwnProperty('hostname')) {
                obj['hostname'] = ApiClient.convertToType(data['hostname'], 'String');
            }
            if (data.hasOwnProperty('ip_addresses')) {
                obj['ip_addresses'] = ApiClient.convertToType(data['ip_addresses'], ['String']);
            }
            if (data.hasOwnProperty('mac_addresses')) {
                obj['mac_addresses'] = ApiClient.convertToType(data['mac_addresses'], ['String']);
            }
            if (data.hasOwnProperty('os_name')) {
                obj['os_name'] = ApiClient.convertToType(data['os_name'], 'String');
            }
            if (data.hasOwnProperty('os_pretty_name')) {
                obj['os_pretty_name'] = ApiClient.convertToType(data['os_pretty_name'], 'String');
            }
            if (data.hasOwnProperty('request_ip')) {
                obj['request_ip'] = ApiClient.convertToType(data['request_ip'], 'String');
            }
            if (data.hasOwnProperty('src_uid')) {
                obj['src_uid'] = ApiClient.convertToType(data['src_uid'], 'String');
            }
            if (data.hasOwnProperty('uname')) {
                obj['uname'] = ApiClient.convertToType(data['uname'], 'String');
            }
        }
        return obj;
    }


}

/**
 * @member {String} agent_arch
 */
OrcApiRuntimeDetails.prototype['agent_arch'] = undefined;

/**
 * @member {String} agent_version
 */
OrcApiRuntimeDetails.prototype['agent_version'] = undefined;

/**
 * @member {Number} boot_time
 */
OrcApiRuntimeDetails.prototype['boot_time'] = undefined;

/**
 * @member {String} cloud_image_id
 */
OrcApiRuntimeDetails.prototype['cloud_image_id'] = undefined;

/**
 * @member {String} cloud_instance_id
 */
OrcApiRuntimeDetails.prototype['cloud_instance_id'] = undefined;

/**
 * @member {String} cloud_region
 */
OrcApiRuntimeDetails.prototype['cloud_region'] = undefined;

/**
 * @member {Array.<String>} cloud_tags
 */
OrcApiRuntimeDetails.prototype['cloud_tags'] = undefined;

/**
 * @member {String} cloud_type
 */
OrcApiRuntimeDetails.prototype['cloud_type'] = undefined;

/**
 * @member {Number} cpu_cores
 */
OrcApiRuntimeDetails.prototype['cpu_cores'] = undefined;

/**
 * @member {String} cpu_make
 */
OrcApiRuntimeDetails.prototype['cpu_make'] = undefined;

/**
 * @member {String} cpu_model
 */
OrcApiRuntimeDetails.prototype['cpu_model'] = undefined;

/**
 * @member {String} hostname
 */
OrcApiRuntimeDetails.prototype['hostname'] = undefined;

/**
 * @member {Array.<String>} ip_addresses
 */
OrcApiRuntimeDetails.prototype['ip_addresses'] = undefined;

/**
 * @member {Array.<String>} mac_addresses
 */
OrcApiRuntimeDetails.prototype['mac_addresses'] = undefined;

/**
 * @member {String} os_name
 */
OrcApiRuntimeDetails.prototype['os_name'] = undefined;

/**
 * @member {String} os_pretty_name
 */
OrcApiRuntimeDetails.prototype['os_pretty_name'] = undefined;

/**
 * @member {String} request_ip
 */
OrcApiRuntimeDetails.prototype['request_ip'] = undefined;

/**
 * @member {String} src_uid
 */
OrcApiRuntimeDetails.prototype['src_uid'] = undefined;

/**
 * @member {String} uname
 */
OrcApiRuntimeDetails.prototype['uname'] = undefined;






export default OrcApiRuntimeDetails;

