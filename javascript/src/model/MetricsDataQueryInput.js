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
import Expr from './Expr';
import RstoreCausalQuery from './RstoreCausalQuery';

/**
 * The MetricsDataQueryInput model module.
 * @module model/MetricsDataQueryInput
 * @version 0.1.0
 */
class MetricsDataQueryInput {
    /**
     * Constructs a new <code>MetricsDataQueryInput</code>.
     * @alias module:model/MetricsDataQueryInput
     * @param dataType {String} DataType to query
     * @param orgUid {String} Organization UID to query
     */
    constructor(dataType, orgUid) { 
        
        MetricsDataQueryInput.initialize(this, dataType, orgUid);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, dataType, orgUid) { 
        obj['data_type'] = dataType;
        obj['org_uid'] = orgUid;
    }

    /**
     * Constructs a <code>MetricsDataQueryInput</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/MetricsDataQueryInput} obj Optional instance to populate.
     * @return {module:model/MetricsDataQueryInput} The populated <code>MetricsDataQueryInput</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new MetricsDataQueryInput();

            if (data.hasOwnProperty('aggs')) {
                obj['aggs'] = ApiClient.convertToType(data['aggs'], {'String': Object});
            }
            if (data.hasOwnProperty('causal')) {
                obj['causal'] = RstoreCausalQuery.constructFromObject(data['causal']);
            }
            if (data.hasOwnProperty('context_uid')) {
                obj['context_uid'] = ApiClient.convertToType(data['context_uid'], 'String');
            }
            if (data.hasOwnProperty('data_type')) {
                obj['data_type'] = ApiClient.convertToType(data['data_type'], 'String');
            }
            if (data.hasOwnProperty('end_time')) {
                obj['end_time'] = ApiClient.convertToType(data['end_time'], 'Number');
            }
            if (data.hasOwnProperty('expr')) {
                obj['expr'] = Expr.constructFromObject(data['expr']);
            }
            if (data.hasOwnProperty('ids')) {
                obj['ids'] = ApiClient.convertToType(data['ids'], ['String']);
            }
            if (data.hasOwnProperty('org_uid')) {
                obj['org_uid'] = ApiClient.convertToType(data['org_uid'], 'String');
            }
            if (data.hasOwnProperty('projection')) {
                obj['projection'] = ApiClient.convertToType(data['projection'], ['String']);
            }
            if (data.hasOwnProperty('query')) {
                obj['query'] = ApiClient.convertToType(data['query'], 'String');
            }
            if (data.hasOwnProperty('query_from')) {
                obj['query_from'] = ApiClient.convertToType(data['query_from'], 'Number');
            }
            if (data.hasOwnProperty('query_size')) {
                obj['query_size'] = ApiClient.convertToType(data['query_size'], 'Number');
            }
            if (data.hasOwnProperty('return_rptrs')) {
                obj['return_rptrs'] = ApiClient.convertToType(data['return_rptrs'], 'Boolean');
            }
            if (data.hasOwnProperty('rptrs')) {
                obj['rptrs'] = ApiClient.convertToType(data['rptrs'], ['String']);
            }
            if (data.hasOwnProperty('sort_col')) {
                obj['sort_col'] = ApiClient.convertToType(data['sort_col'], 'String');
            }
            if (data.hasOwnProperty('sort_order')) {
                obj['sort_order'] = ApiClient.convertToType(data['sort_order'], 'String');
            }
            if (data.hasOwnProperty('src_uid')) {
                obj['src_uid'] = ApiClient.convertToType(data['src_uid'], 'String');
            }
            if (data.hasOwnProperty('start_time')) {
                obj['start_time'] = ApiClient.convertToType(data['start_time'], 'Number');
            }
        }
        return obj;
    }


}

/**
 * Aggregations
 * @member {Object.<String, Object>} aggs
 */
MetricsDataQueryInput.prototype['aggs'] = undefined;

/**
 * @member {module:model/RstoreCausalQuery} causal
 */
MetricsDataQueryInput.prototype['causal'] = undefined;

/**
 * Context UID for this query, it's used to track the query as it flows through the system, and shouldn't be exposed to customers
 * @member {String} context_uid
 */
MetricsDataQueryInput.prototype['context_uid'] = undefined;

/**
 * DataType to query
 * @member {String} data_type
 */
MetricsDataQueryInput.prototype['data_type'] = undefined;

/**
 * Time in unix epoch time, records < time are returned
 * @member {Number} end_time
 */
MetricsDataQueryInput.prototype['end_time'] = undefined;

/**
 * @member {module:model/Expr} expr
 */
MetricsDataQueryInput.prototype['expr'] = undefined;

/**
 * Array of IDs to resolve into records
 * @member {Array.<String>} ids
 */
MetricsDataQueryInput.prototype['ids'] = undefined;

/**
 * Organization UID to query
 * @member {String} org_uid
 */
MetricsDataQueryInput.prototype['org_uid'] = undefined;

/**
 * Array of top level object properties which will be emitted, if none are specified all will be emitted
 * @member {Array.<String>} projection
 */
MetricsDataQueryInput.prototype['projection'] = undefined;

/**
 * Lucene based search query
 * @member {String} query
 */
MetricsDataQueryInput.prototype['query'] = undefined;

/**
 * Where to start the query in the result set from
 * @member {Number} query_from
 */
MetricsDataQueryInput.prototype['query_from'] = undefined;

/**
 * Size of the query result set
 * @member {Number} query_size
 */
MetricsDataQueryInput.prototype['query_size'] = undefined;

/**
 * @member {Boolean} return_rptrs
 */
MetricsDataQueryInput.prototype['return_rptrs'] = undefined;

/**
 * @member {Array.<String>} rptrs
 */
MetricsDataQueryInput.prototype['rptrs'] = undefined;

/**
 * Sort column
 * @member {String} sort_col
 */
MetricsDataQueryInput.prototype['sort_col'] = undefined;

/**
 * Sort order
 * @member {String} sort_order
 */
MetricsDataQueryInput.prototype['sort_order'] = undefined;

/**
 * Source UID to query
 * @member {String} src_uid
 */
MetricsDataQueryInput.prototype['src_uid'] = undefined;

/**
 * Time in unix epoch time, records >= time are returned
 * @member {Number} start_time
 */
MetricsDataQueryInput.prototype['start_time'] = undefined;






export default MetricsDataQueryInput;

