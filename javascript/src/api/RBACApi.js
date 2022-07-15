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
import ApiRBACActions from '../model/ApiRBACActions';
import CanUserPerformInput from '../model/CanUserPerformInput';
import ValidationError from '../model/ValidationError';

/**
* RBAC service.
* @module api/RBACApi
* @version 1.0.0
*/
export default class RBACApi {

    /**
    * Constructs a new RBACApi. 
    * @alias module:api/RBACApi
    * @class
    * @param {module:ApiClient} [apiClient] Optional API client implementation to use,
    * default to {@link module:ApiClient#instance} if unspecified.
    */
    constructor(apiClient) {
        this.apiClient = apiClient || ApiClient.instance;
    }


    /**
     * Callback function to receive the result of the canUserPerform operation.
     * @callback module:api/RBACApi~canUserPerformCallback
     * @param {String} error Error message, if any.
     * @param {module:model/ApiRBACActions} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Query allows actions on objects
     * Allows for querying of what actions a user can perform; results may be cached for a short period of time.
     * @param {Object} opts Optional parameters
     * @param {module:model/CanUserPerformInput} opts.canUserPerformInput 
     * @param {module:api/RBACApi~canUserPerformCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/ApiRBACActions}
     */
    canUserPerform(opts, callback) {
      opts = opts || {};
      let postBody = opts['canUserPerformInput'];

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
      let accepts = ['application/json'];
      let returnType = ApiRBACActions;
      return this.apiClient.callApi(
        '/api/v1/rbac/capabilities/', 'POST',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }


}
