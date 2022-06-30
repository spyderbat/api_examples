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
import DashboardSearch from '../model/DashboardSearch';
import DashboardSearchCreateInput from '../model/DashboardSearchCreateInput';
import DashboardSearchUpdateInput from '../model/DashboardSearchUpdateInput';

/**
* DashboardSearch service.
* @module api/DashboardSearchApi
* @version 0.1.0
*/
export default class DashboardSearchApi {

    /**
    * Constructs a new DashboardSearchApi. 
    * @alias module:api/DashboardSearchApi
    * @class
    * @param {module:ApiClient} [apiClient] Optional API client implementation to use,
    * default to {@link module:ApiClient#instance} if unspecified.
    */
    constructor(apiClient) {
        this.apiClient = apiClient || ApiClient.instance;
    }


    /**
     * Callback function to receive the result of the dashboardSearchCreate operation.
     * @callback module:api/DashboardSearchApi~dashboardSearchCreateCallback
     * @param {String} error Error message, if any.
     * @param {String} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Create a dashboard search
     *  Create a dashboard search in an org.   * Requires action dashboard_search:Create
     * @param {String} dashboardSearchUID UID for the DashboardSearch
     * @param {String} orgUID Org UID
     * @param {Object} opts Optional parameters
     * @param {module:model/DashboardSearchCreateInput} opts.dashboardSearchCreateInput 
     * @param {module:api/DashboardSearchApi~dashboardSearchCreateCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link String}
     */
    dashboardSearchCreate(dashboardSearchUID, orgUID, opts, callback) {
      opts = opts || {};
      let postBody = opts['dashboardSearchCreateInput'];
      // verify the required parameter 'dashboardSearchUID' is set
      if (dashboardSearchUID === undefined || dashboardSearchUID === null) {
        throw new Error("Missing the required parameter 'dashboardSearchUID' when calling dashboardSearchCreate");
      }
      // verify the required parameter 'orgUID' is set
      if (orgUID === undefined || orgUID === null) {
        throw new Error("Missing the required parameter 'orgUID' when calling dashboardSearchCreate");
      }

      let pathParams = {
        'dashboardSearchUID': dashboardSearchUID,
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
      let returnType = 'String';
      return this.apiClient.callApi(
        '/api/v1/org/{orgUID}/dashboard_search/', 'POST',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the dashboardSearchDelete operation.
     * @callback module:api/DashboardSearchApi~dashboardSearchDeleteCallback
     * @param {String} error Error message, if any.
     * @param data This operation does not return a value.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Get a dashboard search
     *  Get a dashboard search in an org.   * Requires action dashboard_search:Delete
     * @param {String} dashboardSearchUID 
     * @param {String} orgUID 
     * @param {module:api/DashboardSearchApi~dashboardSearchDeleteCallback} callback The callback function, accepting three arguments: error, data, response
     */
    dashboardSearchDelete(dashboardSearchUID, orgUID, callback) {
      let postBody = null;
      // verify the required parameter 'dashboardSearchUID' is set
      if (dashboardSearchUID === undefined || dashboardSearchUID === null) {
        throw new Error("Missing the required parameter 'dashboardSearchUID' when calling dashboardSearchDelete");
      }
      // verify the required parameter 'orgUID' is set
      if (orgUID === undefined || orgUID === null) {
        throw new Error("Missing the required parameter 'orgUID' when calling dashboardSearchDelete");
      }

      let pathParams = {
        'dashboardSearchUID': dashboardSearchUID,
        'orgUID': orgUID
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
        '/api/v1/org/{orgUID}/dashboard_search/{dashboardSearchUID}', 'DELETE',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the dashboardSearchGet operation.
     * @callback module:api/DashboardSearchApi~dashboardSearchGetCallback
     * @param {String} error Error message, if any.
     * @param {module:model/DashboardSearch} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Get a dashboard search
     *  Get a dashboard search in an org.   * Requires action dashboard_search:Get
     * @param {String} dashboardSearchUID 
     * @param {String} orgUID 
     * @param {module:api/DashboardSearchApi~dashboardSearchGetCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/DashboardSearch}
     */
    dashboardSearchGet(dashboardSearchUID, orgUID, callback) {
      let postBody = null;
      // verify the required parameter 'dashboardSearchUID' is set
      if (dashboardSearchUID === undefined || dashboardSearchUID === null) {
        throw new Error("Missing the required parameter 'dashboardSearchUID' when calling dashboardSearchGet");
      }
      // verify the required parameter 'orgUID' is set
      if (orgUID === undefined || orgUID === null) {
        throw new Error("Missing the required parameter 'orgUID' when calling dashboardSearchGet");
      }

      let pathParams = {
        'dashboardSearchUID': dashboardSearchUID,
        'orgUID': orgUID
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
      let returnType = DashboardSearch;
      return this.apiClient.callApi(
        '/api/v1/org/{orgUID}/dashboard_search/{dashboardSearchUID}', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the dashboardSearchList operation.
     * @callback module:api/DashboardSearchApi~dashboardSearchListCallback
     * @param {String} error Error message, if any.
     * @param {Array.<module:model/DashboardSearch>} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * List dashboard searches
     *  Lists dashboard searches by org.   * Requires action dashboard_search:Query
     * @param {String} orgUID Org UID
     * @param {module:api/DashboardSearchApi~dashboardSearchListCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link Array.<module:model/DashboardSearch>}
     */
    dashboardSearchList(orgUID, callback) {
      let postBody = null;
      // verify the required parameter 'orgUID' is set
      if (orgUID === undefined || orgUID === null) {
        throw new Error("Missing the required parameter 'orgUID' when calling dashboardSearchList");
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
      let contentTypes = [];
      let accepts = ['application/json'];
      let returnType = [DashboardSearch];
      return this.apiClient.callApi(
        '/api/v1/org/{orgUID}/dashboard_search/', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the dashboardSearchUpdate operation.
     * @callback module:api/DashboardSearchApi~dashboardSearchUpdateCallback
     * @param {String} error Error message, if any.
     * @param data This operation does not return a value.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Update a dashboard search
     *  Update a dashboard search in an org.   * Requires action dashboard_search:Update
     * @param {String} dashboardSearchUID UID for the DashboardSearch
     * @param {String} orgUID Org UID
     * @param {Object} opts Optional parameters
     * @param {module:model/DashboardSearchUpdateInput} opts.dashboardSearchUpdateInput 
     * @param {module:api/DashboardSearchApi~dashboardSearchUpdateCallback} callback The callback function, accepting three arguments: error, data, response
     */
    dashboardSearchUpdate(dashboardSearchUID, orgUID, opts, callback) {
      opts = opts || {};
      let postBody = opts['dashboardSearchUpdateInput'];
      // verify the required parameter 'dashboardSearchUID' is set
      if (dashboardSearchUID === undefined || dashboardSearchUID === null) {
        throw new Error("Missing the required parameter 'dashboardSearchUID' when calling dashboardSearchUpdate");
      }
      // verify the required parameter 'orgUID' is set
      if (orgUID === undefined || orgUID === null) {
        throw new Error("Missing the required parameter 'orgUID' when calling dashboardSearchUpdate");
      }

      let pathParams = {
        'dashboardSearchUID': dashboardSearchUID,
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
      let accepts = [];
      let returnType = null;
      return this.apiClient.callApi(
        '/api/v1/org/{orgUID}/dashboard_search/{dashboardSearchUID}', 'PUT',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }


}
