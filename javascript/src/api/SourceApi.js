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
import ApiSOARListHandlerOutput from '../model/ApiSOARListHandlerOutput';
import ApiSourceCreateHandlerOutput from '../model/ApiSourceCreateHandlerOutput';
import Source from '../model/Source';
import SrcCreateInput from '../model/SrcCreateInput';
import SrcUpdateInput from '../model/SrcUpdateInput';
import ValidationError from '../model/ValidationError';

/**
* Source service.
* @module api/SourceApi
* @version 0.1.0
*/
export default class SourceApi {

    /**
    * Constructs a new SourceApi. 
    * @alias module:api/SourceApi
    * @class
    * @param {module:ApiClient} [apiClient] Optional API client implementation to use,
    * default to {@link module:ApiClient#instance} if unspecified.
    */
    constructor(apiClient) {
        this.apiClient = apiClient || ApiClient.instance;
    }


    /**
     * Callback function to receive the result of the integrationSoarSrcList operation.
     * @callback module:api/SourceApi~integrationSoarSrcListCallback
     * @param {String} error Error message, if any.
     * @param {Array.<module:model/ApiSOARListHandlerOutput>} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * List sources for integration with SOARs
     *  Lists the sources of data that match the specified query parameters, and return  URL entry points into the UI for matching sources.   * Requires the action  *org:ListSources* on the organization 
     * @param {String} orgUID 
     * @param {Object} opts Optional parameters
     * @param {Number} opts.et optional end time of the query
     * @param {String} opts.hostname A single hostname to match
     * @param {String} opts.ipAddress A single IP address to match
     * @param {String} opts.macAddress A single mac address to match
     * @param {Number} opts.page 
     * @param {Number} opts.pageSize 
     * @param {Number} opts.st optional start time of the query, if only a start time is provided, end time will be start+10m
     * @param {module:api/SourceApi~integrationSoarSrcListCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link Array.<module:model/ApiSOARListHandlerOutput>}
     */
    integrationSoarSrcList(orgUID, opts, callback) {
      opts = opts || {};
      let postBody = null;
      // verify the required parameter 'orgUID' is set
      if (orgUID === undefined || orgUID === null) {
        throw new Error("Missing the required parameter 'orgUID' when calling integrationSoarSrcList");
      }

      let pathParams = {
        'orgUID': orgUID
      };
      let queryParams = {
        'et': opts['et'],
        'hostname': opts['hostname'],
        'ip_address': opts['ipAddress'],
        'mac_address': opts['macAddress'],
        'page': opts['page'],
        'page_size': opts['pageSize'],
        'st': opts['st']
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = ['apiToken'];
      let contentTypes = [];
      let accepts = ['application/json'];
      let returnType = [ApiSOARListHandlerOutput];
      return this.apiClient.callApi(
        '/api/v1/integration/soar/org/{orgUID}/source/', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the srcCreate operation.
     * @callback module:api/SourceApi~srcCreateCallback
     * @param {String} error Error message, if any.
     * @param {module:model/ApiSourceCreateHandlerOutput} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Create a source
     *  Creates a new source of data  * Requires the action  *org:CreateSource* on the organization 
     * @param {String} orgUID The org this source is associated with
     * @param {Object} opts Optional parameters
     * @param {module:model/SrcCreateInput} opts.srcCreateInput 
     * @param {module:api/SourceApi~srcCreateCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/ApiSourceCreateHandlerOutput}
     */
    srcCreate(orgUID, opts, callback) {
      opts = opts || {};
      let postBody = opts['srcCreateInput'];
      // verify the required parameter 'orgUID' is set
      if (orgUID === undefined || orgUID === null) {
        throw new Error("Missing the required parameter 'orgUID' when calling srcCreate");
      }

      let pathParams = {
        'orgUID': orgUID
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = ['apiToken'];
      let contentTypes = ['application/json'];
      let accepts = ['application/json'];
      let returnType = ApiSourceCreateHandlerOutput;
      return this.apiClient.callApi(
        '/api/v1/org/{orgUID}/source/', 'POST',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the srcDelete operation.
     * @callback module:api/SourceApi~srcDeleteCallback
     * @param {String} error Error message, if any.
     * @param data This operation does not return a value.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Delete a source
     *  Delete a source  * Requires the action  *org:DeleteSource* on the organization 
     * @param {String} orgUID 
     * @param {String} sourceUID 
     * @param {module:api/SourceApi~srcDeleteCallback} callback The callback function, accepting three arguments: error, data, response
     */
    srcDelete(orgUID, sourceUID, callback) {
      let postBody = null;
      // verify the required parameter 'orgUID' is set
      if (orgUID === undefined || orgUID === null) {
        throw new Error("Missing the required parameter 'orgUID' when calling srcDelete");
      }
      // verify the required parameter 'sourceUID' is set
      if (sourceUID === undefined || sourceUID === null) {
        throw new Error("Missing the required parameter 'sourceUID' when calling srcDelete");
      }

      let pathParams = {
        'orgUID': orgUID,
        'sourceUID': sourceUID
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = ['apiToken'];
      let contentTypes = [];
      let accepts = [];
      let returnType = null;
      return this.apiClient.callApi(
        '/api/v1/org/{orgUID}/source/{sourceUID}', 'DELETE',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the srcList operation.
     * @callback module:api/SourceApi~srcListCallback
     * @param {String} error Error message, if any.
     * @param {Array.<module:model/Source>} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * List sources
     *  Lists the sources of data for an organization  * Requires the action  *org:ListSources* on the organization 
     * @param {String} orgUID 
     * @param {Object} opts Optional parameters
     * @param {String} opts.agentUidEquals 
     * @param {String} opts.descriptionContains 
     * @param {Array.<String>} opts.hasTags 
     * @param {Boolean} opts.originalAssociation 
     * @param {Number} opts.page 
     * @param {Number} opts.pageSize 
     * @param {module:api/SourceApi~srcListCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link Array.<module:model/Source>}
     */
    srcList(orgUID, opts, callback) {
      opts = opts || {};
      let postBody = null;
      // verify the required parameter 'orgUID' is set
      if (orgUID === undefined || orgUID === null) {
        throw new Error("Missing the required parameter 'orgUID' when calling srcList");
      }

      let pathParams = {
        'orgUID': orgUID
      };
      let queryParams = {
        'agent_uid_equals': opts['agentUidEquals'],
        'description_contains': opts['descriptionContains'],
        'has_tags': this.apiClient.buildCollectionParam(opts['hasTags'], 'multi'),
        'original_association': opts['originalAssociation'],
        'page': opts['page'],
        'page_size': opts['pageSize']
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = ['apiToken'];
      let contentTypes = [];
      let accepts = ['application/json'];
      let returnType = [Source];
      return this.apiClient.callApi(
        '/api/v1/org/{orgUID}/source/', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the srcLoad operation.
     * @callback module:api/SourceApi~srcLoadCallback
     * @param {String} error Error message, if any.
     * @param {module:model/Source} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Load a source
     *  Loads the meta data for a source of data  * Requires the action  *org:LoadSource* on the organization 
     * @param {String} orgUID 
     * @param {String} sourceUID 
     * @param {module:api/SourceApi~srcLoadCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/Source}
     */
    srcLoad(orgUID, sourceUID, callback) {
      let postBody = null;
      // verify the required parameter 'orgUID' is set
      if (orgUID === undefined || orgUID === null) {
        throw new Error("Missing the required parameter 'orgUID' when calling srcLoad");
      }
      // verify the required parameter 'sourceUID' is set
      if (sourceUID === undefined || sourceUID === null) {
        throw new Error("Missing the required parameter 'sourceUID' when calling srcLoad");
      }

      let pathParams = {
        'orgUID': orgUID,
        'sourceUID': sourceUID
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
      let returnType = Source;
      return this.apiClient.callApi(
        '/api/v1/org/{orgUID}/source/{sourceUID}', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the srcUpdate operation.
     * @callback module:api/SourceApi~srcUpdateCallback
     * @param {String} error Error message, if any.
     * @param data This operation does not return a value.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Update a source
     *  Updates the meta data for a source of data  * Requires the action  *org:UpdateSource* on the organization 
     * @param {String} orgUID The org this source is associated with
     * @param {String} sourceUID The UID of the source
     * @param {Object} opts Optional parameters
     * @param {module:model/SrcUpdateInput} opts.srcUpdateInput 
     * @param {module:api/SourceApi~srcUpdateCallback} callback The callback function, accepting three arguments: error, data, response
     */
    srcUpdate(orgUID, sourceUID, opts, callback) {
      opts = opts || {};
      let postBody = opts['srcUpdateInput'];
      // verify the required parameter 'orgUID' is set
      if (orgUID === undefined || orgUID === null) {
        throw new Error("Missing the required parameter 'orgUID' when calling srcUpdate");
      }
      // verify the required parameter 'sourceUID' is set
      if (sourceUID === undefined || sourceUID === null) {
        throw new Error("Missing the required parameter 'sourceUID' when calling srcUpdate");
      }

      let pathParams = {
        'orgUID': orgUID,
        'sourceUID': sourceUID
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = ['apiToken'];
      let contentTypes = ['application/json'];
      let accepts = ['application/json'];
      let returnType = null;
      return this.apiClient.callApi(
        '/api/v1/org/{orgUID}/source/{sourceUID}', 'PUT',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }


}
