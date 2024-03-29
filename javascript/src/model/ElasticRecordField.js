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
 * The ElasticRecordField model module.
 * @module model/ElasticRecordField
 * @version 1.0.0
 */
class ElasticRecordField {
    /**
     * Constructs a new <code>ElasticRecordField</code>.
     * @alias module:model/ElasticRecordField
     */
    constructor() { 
        
        ElasticRecordField.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>ElasticRecordField</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ElasticRecordField} obj Optional instance to populate.
     * @return {module:model/ElasticRecordField} The populated <code>ElasticRecordField</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ElasticRecordField();

            if (data.hasOwnProperty('dynamic')) {
                obj['dynamic'] = ApiClient.convertToType(data['dynamic'], 'Boolean');
            }
            if (data.hasOwnProperty('fields')) {
                obj['fields'] = ApiClient.convertToType(data['fields'], {'String': ElasticRecordField});
            }
            if (data.hasOwnProperty('index')) {
                obj['index'] = ApiClient.convertToType(data['index'], 'Boolean');
            }
            if (data.hasOwnProperty('properties')) {
                obj['properties'] = ApiClient.convertToType(data['properties'], {'String': ElasticRecordField});
            }
            if (data.hasOwnProperty('store')) {
                obj['store'] = ApiClient.convertToType(data['store'], 'Boolean');
            }
            if (data.hasOwnProperty('type')) {
                obj['type'] = ApiClient.convertToType(data['type'], 'String');
            }
        }
        return obj;
    }


}

/**
 * @member {Boolean} dynamic
 */
ElasticRecordField.prototype['dynamic'] = undefined;

/**
 * @member {Object.<String, module:model/ElasticRecordField>} fields
 */
ElasticRecordField.prototype['fields'] = undefined;

/**
 * @member {Boolean} index
 */
ElasticRecordField.prototype['index'] = undefined;

/**
 * The properties associated with this field
 * @member {Object.<String, module:model/ElasticRecordField>} properties
 */
ElasticRecordField.prototype['properties'] = undefined;

/**
 * @member {Boolean} store
 */
ElasticRecordField.prototype['store'] = undefined;

/**
 * The type used for indexing, keyword=matches entire seaerch term, text=partial match, ip=ip address, float=number
 * @member {String} type
 */
ElasticRecordField.prototype['type'] = undefined;






export default ElasticRecordField;

