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


import ApiClient from "../ApiClient";
import ElasticRecordSchema from '../model/ElasticRecordSchema';

/**
* MetadataAPI service.
* @module api/MetadataAPIApi
* @version 1.0.0
*/
export default class MetadataAPIApi {

    /**
    * Constructs a new MetadataAPIApi. 
    * @alias module:api/MetadataAPIApi
    * @class
    * @param {module:ApiClient} [apiClient] Optional API client implementation to use,
    * default to {@link module:ApiClient#instance} if unspecified.
    */
    constructor(apiClient) {
        this.apiClient = apiClient || ApiClient.instance;
    }


    /**
     * Callback function to receive the result of the metadataSearchSchema operation.
     * @callback module:api/MetadataAPIApi~metadataSearchSchemaCallback
     * @param {String} error Error message, if any.
     * @param {Object.<String, module:model/{String: ElasticRecordSchema}>} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Returns the schema used for search
     *  Returns meta data around which model and event properties are indexed for search, i.e. the search schema 
     * @param {module:api/MetadataAPIApi~metadataSearchSchemaCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link Object.<String, module:model/{String: ElasticRecordSchema}>}
     */
    metadataSearchSchema(callback) {
      let postBody = null;

      let pathParams = {
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = ['apiToken'];
      let contentTypes = [];
      let accepts = ['application/json'];
      let returnType = {'String': ElasticRecordSchema};
      return this.apiClient.callApi(
        '/api/v1/meta/search/schema', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }


}
