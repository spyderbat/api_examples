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


import ApiClient from "../ApiClient";
import SrcDataQueryInput from '../model/SrcDataQueryInput';
import ValidationError from '../model/ValidationError';

/**
* SourceData service.
* @module api/SourceDataApi
* @version 1.0.0
*/
export default class SourceDataApi {

    /**
    * Constructs a new SourceDataApi. 
    * @alias module:api/SourceDataApi
    * @class
    * @param {module:ApiClient} [apiClient] Optional API client implementation to use,
    * default to {@link module:ApiClient#instance} if unspecified.
    */
    constructor(apiClient) {
        this.apiClient = apiClient || ApiClient.instance;
    }


    /**
     * Callback function to receive the result of the srcDataQuery operation.
     * @callback module:api/SourceDataApi~srcDataQueryCallback
     * @param {String} error Error message, if any.
     * @param {String} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Query source data
     *  Allows querying of the source data, data is stored as 'records' which are returned as json objects, in nd-json (see ndjson.org) format.   * Data is returned as it is matched, and no ordering guarentees exist.  * The call completes after it has finished searching for matching records.  * The query expression is limited to seaching a 24 hour period of time, it is the callers responsibility to construct an appropriate 24 hour range. * Documentation for the returned spydergraph datatype can be found at https://app.spyderbat.com/schema/spydergraph/index.html  * The user must have both the *org.ListSourceData* action on the org and *source_data.Query* action on the source * To get a count of results (up to 10K) but no data, use querySize: 0  
     * @param {Object} opts Optional parameters
     * @param {module:model/SrcDataQueryInput} opts.srcDataQueryInput 
     * @param {module:api/SourceDataApi~srcDataQueryCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link String}
     */
    srcDataQuery(opts, callback) {
      opts = opts || {};
      let postBody = opts['srcDataQueryInput'];

      let pathParams = {
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = ['apiToken'];
      let contentTypes = ['application/json'];
      let accepts = ['application/x-ndjson', 'application/json'];
      let returnType = 'String';
      return this.apiClient.callApi(
        '/api/v1/source/query/', 'POST',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the srcDataQueryV2 operation.
     * @callback module:api/SourceDataApi~srcDataQueryV2Callback
     * @param {String} error Error message, if any.
     * @param {String} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Query source data
     * Same as the post query above except results are cached
     * @param {String} orgUID Organization UID to query
     * @param {String} dt DataType to query
     * @param {Object} opts Optional parameters
     * @param {String} opts.e Data which matches this expression are returned, as a json object
     * @param {Number} opts.et Time in unix epoch time, records < time are returned
     * @param {Array.<String>} opts.id List of IDs to resolve
     * @param {Array.<String>} opts.pj Array of top level object properties which will be emitted, if none are specified all will be emitted; ex pj=id&pj=version
     * @param {String} opts.q Lucene based search query
     * @param {Number} opts.qf Where to start the query in the result set from
     * @param {Number} opts.qs Size of the query result set
     * @param {Boolean} opts.rr Return rptrs mixed with the data
     * @param {String} opts.src Source UID to query
     * @param {Number} opts.st Time in unix epoch time, records >= time are returned
     * @param {module:api/SourceDataApi~srcDataQueryV2Callback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link String}
     */
    srcDataQueryV2(orgUID, dt, opts, callback) {
      opts = opts || {};
      let postBody = null;
      // verify the required parameter 'orgUID' is set
      if (orgUID === undefined || orgUID === null) {
        throw new Error("Missing the required parameter 'orgUID' when calling srcDataQueryV2");
      }
      // verify the required parameter 'dt' is set
      if (dt === undefined || dt === null) {
        throw new Error("Missing the required parameter 'dt' when calling srcDataQueryV2");
      }

      let pathParams = {
        'orgUID': orgUID
      };
      let queryParams = {
        'dt': dt,
        'e': opts['e'],
        'et': opts['et'],
        'id': this.apiClient.buildCollectionParam(opts['id'], 'multi'),
        'pj': this.apiClient.buildCollectionParam(opts['pj'], 'multi'),
        'q': opts['q'],
        'qf': opts['qf'],
        'qs': opts['qs'],
        'rr': opts['rr'],
        'src': opts['src'],
        'st': opts['st']
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = ['apiToken'];
      let contentTypes = [];
      let accepts = ['application/x-ndjson', 'application/json'];
      let returnType = 'String';
      return this.apiClient.callApi(
        '/api/v1/org/{orgUID}/data/', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the srcSendData operation.
     * @callback module:api/SourceDataApi~srcSendDataCallback
     * @param {String} error Error message, if any.
     * @param data This operation does not return a value.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Send data to a source, this is expected to be gzip compressed nd-json. The 'Content-Encoding' header should be specified with a value of 'gzip'. Alternatively, a multi-part form upload may be used with gzipped data up to a maximum size of 1MB.
     * Sends data to a source
     * @param {String} dataType 
     * @param {String} orgUID 
     * @param {String} sourceUID 
     * @param {String} encoding must be gzip
     * @param {File} file The file to upload. The file must be a valid gzip-ed JSON file.
     * @param {module:api/SourceDataApi~srcSendDataCallback} callback The callback function, accepting three arguments: error, data, response
     */
    srcSendData(dataType, orgUID, sourceUID, encoding, file, callback) {
      let postBody = null;
      // verify the required parameter 'dataType' is set
      if (dataType === undefined || dataType === null) {
        throw new Error("Missing the required parameter 'dataType' when calling srcSendData");
      }
      // verify the required parameter 'orgUID' is set
      if (orgUID === undefined || orgUID === null) {
        throw new Error("Missing the required parameter 'orgUID' when calling srcSendData");
      }
      // verify the required parameter 'sourceUID' is set
      if (sourceUID === undefined || sourceUID === null) {
        throw new Error("Missing the required parameter 'sourceUID' when calling srcSendData");
      }
      // verify the required parameter 'encoding' is set
      if (encoding === undefined || encoding === null) {
        throw new Error("Missing the required parameter 'encoding' when calling srcSendData");
      }
      // verify the required parameter 'file' is set
      if (file === undefined || file === null) {
        throw new Error("Missing the required parameter 'file' when calling srcSendData");
      }

      let pathParams = {
        'dataType': dataType,
        'orgUID': orgUID,
        'sourceUID': sourceUID
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
        'encoding': encoding,
        'file': file
      };

      let authNames = ['apiToken'];
      let contentTypes = ['multipart/form-data'];
      let accepts = ['application/json'];
      let returnType = null;
      return this.apiClient.callApi(
        '/api/v1/org/{orgUID}/source/{sourceUID}/data/{dataType}', 'POST',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }


}
