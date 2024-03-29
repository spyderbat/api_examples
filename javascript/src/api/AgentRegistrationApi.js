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
import AgentRegistration from '../model/AgentRegistration';
import AgentRegistrationCreateInput from '../model/AgentRegistrationCreateInput';
import AgentRegistrationUpdateInput from '../model/AgentRegistrationUpdateInput';
import ApiAgentCreateHandlerOutput from '../model/ApiAgentCreateHandlerOutput';
import ApiAgentRegistrationDownloadLinkHandlerOutput from '../model/ApiAgentRegistrationDownloadLinkHandlerOutput';
import DaoAgentLog from '../model/DaoAgentLog';
import ValidationError from '../model/ValidationError';

/**
* AgentRegistration service.
* @module api/AgentRegistrationApi
* @version 1.0.0
*/
export default class AgentRegistrationApi {

    /**
    * Constructs a new AgentRegistrationApi. 
    * @alias module:api/AgentRegistrationApi
    * @class
    * @param {module:ApiClient} [apiClient] Optional API client implementation to use,
    * default to {@link module:ApiClient#instance} if unspecified.
    */
    constructor(apiClient) {
        this.apiClient = apiClient || ApiClient.instance;
    }


    /**
     * Callback function to receive the result of the agentRegistrationCreate operation.
     * @callback module:api/AgentRegistrationApi~agentRegistrationCreateCallback
     * @param {String} error Error message, if any.
     * @param {module:model/ApiAgentCreateHandlerOutput} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Create an agent registration
     *  Creates a new agent registration  * Requires the action  *agent_registration:Create* on the organization 
     * @param {String} orgUID The OrgUID the registration is associated with
     * @param {String} uid Agent Registration UID
     * @param {Object} opts Optional parameters
     * @param {module:model/AgentRegistrationCreateInput} opts.agentRegistrationCreateInput 
     * @param {module:api/AgentRegistrationApi~agentRegistrationCreateCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/ApiAgentCreateHandlerOutput}
     */
    agentRegistrationCreate(orgUID, uid, opts, callback) {
      opts = opts || {};
      let postBody = opts['agentRegistrationCreateInput'];
      // verify the required parameter 'orgUID' is set
      if (orgUID === undefined || orgUID === null) {
        throw new Error("Missing the required parameter 'orgUID' when calling agentRegistrationCreate");
      }
      // verify the required parameter 'uid' is set
      if (uid === undefined || uid === null) {
        throw new Error("Missing the required parameter 'uid' when calling agentRegistrationCreate");
      }

      let pathParams = {
        'orgUID': orgUID,
        'uid': uid
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
      let returnType = ApiAgentCreateHandlerOutput;
      return this.apiClient.callApi(
        '/api/v1/org/{orgUID}/agent_registration/', 'POST',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the agentRegistrationDownloadLink operation.
     * @callback module:api/AgentRegistrationApi~agentRegistrationDownloadLinkCallback
     * @param {String} error Error message, if any.
     * @param {module:model/ApiAgentRegistrationDownloadLinkHandlerOutput} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Get a download link for this registration
     *  Create a download link for the registration  * Requires the action  *agent_registration:Load* on the agent registration  
     * @param {String} orgUID 
     * @param {String} uid 
     * @param {module:api/AgentRegistrationApi~agentRegistrationDownloadLinkCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/ApiAgentRegistrationDownloadLinkHandlerOutput}
     */
    agentRegistrationDownloadLink(orgUID, uid, callback) {
      let postBody = null;
      // verify the required parameter 'orgUID' is set
      if (orgUID === undefined || orgUID === null) {
        throw new Error("Missing the required parameter 'orgUID' when calling agentRegistrationDownloadLink");
      }
      // verify the required parameter 'uid' is set
      if (uid === undefined || uid === null) {
        throw new Error("Missing the required parameter 'uid' when calling agentRegistrationDownloadLink");
      }

      let pathParams = {
        'orgUID': orgUID,
        'uid': uid
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
      let returnType = ApiAgentRegistrationDownloadLinkHandlerOutput;
      return this.apiClient.callApi(
        '/api/v1/org/{orgUID}/agent_registration/{uid}/download_link', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the agentRegistrationGetLog operation.
     * @callback module:api/AgentRegistrationApi~agentRegistrationGetLogCallback
     * @param {String} error Error message, if any.
     * @param {Array.<module:model/DaoAgentLog>} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Get log of recent agent registration activity
     *  Get lots relating to recent agent registration activity  * Requires the action  *agent_registration:ListLog* on the agent registration  
     * @param {String} orgUID 
     * @param {String} uid 
     * @param {Object} opts Optional parameters
     * @param {Number} opts.page 
     * @param {Number} opts.pageSize 
     * @param {module:api/AgentRegistrationApi~agentRegistrationGetLogCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link Array.<module:model/DaoAgentLog>}
     */
    agentRegistrationGetLog(orgUID, uid, opts, callback) {
      opts = opts || {};
      let postBody = null;
      // verify the required parameter 'orgUID' is set
      if (orgUID === undefined || orgUID === null) {
        throw new Error("Missing the required parameter 'orgUID' when calling agentRegistrationGetLog");
      }
      // verify the required parameter 'uid' is set
      if (uid === undefined || uid === null) {
        throw new Error("Missing the required parameter 'uid' when calling agentRegistrationGetLog");
      }

      let pathParams = {
        'orgUID': orgUID,
        'uid': uid
      };
      let queryParams = {
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
      let returnType = [DaoAgentLog];
      return this.apiClient.callApi(
        '/api/v1/org/{orgUID}/agent_registration/{uid}/log', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the agentRegistrationList operation.
     * @callback module:api/AgentRegistrationApi~agentRegistrationListCallback
     * @param {String} error Error message, if any.
     * @param {Array.<module:model/AgentRegistration>} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * List agent registrations
     *  Lists the agent registrations associated with an organization  * Requires the action  *agent_registration:List* on the organization  
     * @param {String} orgUID 
     * @param {Object} opts Optional parameters
     * @param {Number} opts.page 
     * @param {Number} opts.pageSize 
     * @param {module:api/AgentRegistrationApi~agentRegistrationListCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link Array.<module:model/AgentRegistration>}
     */
    agentRegistrationList(orgUID, opts, callback) {
      opts = opts || {};
      let postBody = null;
      // verify the required parameter 'orgUID' is set
      if (orgUID === undefined || orgUID === null) {
        throw new Error("Missing the required parameter 'orgUID' when calling agentRegistrationList");
      }

      let pathParams = {
        'orgUID': orgUID
      };
      let queryParams = {
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
      let returnType = [AgentRegistration];
      return this.apiClient.callApi(
        '/api/v1/org/{orgUID}/agent_registration/', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the agentRegistrationLoad operation.
     * @callback module:api/AgentRegistrationApi~agentRegistrationLoadCallback
     * @param {String} error Error message, if any.
     * @param {module:model/AgentRegistration} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Load an agent registration
     *  Load a specified agent registration  * Requires the action  *agent_registration:Load* on the agent registration  
     * @param {String} orgUID 
     * @param {String} uid 
     * @param {module:api/AgentRegistrationApi~agentRegistrationLoadCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/AgentRegistration}
     */
    agentRegistrationLoad(orgUID, uid, callback) {
      let postBody = null;
      // verify the required parameter 'orgUID' is set
      if (orgUID === undefined || orgUID === null) {
        throw new Error("Missing the required parameter 'orgUID' when calling agentRegistrationLoad");
      }
      // verify the required parameter 'uid' is set
      if (uid === undefined || uid === null) {
        throw new Error("Missing the required parameter 'uid' when calling agentRegistrationLoad");
      }

      let pathParams = {
        'orgUID': orgUID,
        'uid': uid
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
      let returnType = AgentRegistration;
      return this.apiClient.callApi(
        '/api/v1/org/{orgUID}/agent_registration/{uid}', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the agentRegistrationUpdate operation.
     * @callback module:api/AgentRegistrationApi~agentRegistrationUpdateCallback
     * @param {String} error Error message, if any.
     * @param data This operation does not return a value.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Update an agent registration
     *  Updates an existing registration  * Requires the action  *agent_registration:Update* on the organization and the registration 
     * @param {String} orgUID The OrgUID the registration is associated with
     * @param {String} uid Agent Registration UID
     * @param {Object} opts Optional parameters
     * @param {module:model/AgentRegistrationUpdateInput} opts.agentRegistrationUpdateInput 
     * @param {module:api/AgentRegistrationApi~agentRegistrationUpdateCallback} callback The callback function, accepting three arguments: error, data, response
     */
    agentRegistrationUpdate(orgUID, uid, opts, callback) {
      opts = opts || {};
      let postBody = opts['agentRegistrationUpdateInput'];
      // verify the required parameter 'orgUID' is set
      if (orgUID === undefined || orgUID === null) {
        throw new Error("Missing the required parameter 'orgUID' when calling agentRegistrationUpdate");
      }
      // verify the required parameter 'uid' is set
      if (uid === undefined || uid === null) {
        throw new Error("Missing the required parameter 'uid' when calling agentRegistrationUpdate");
      }

      let pathParams = {
        'orgUID': orgUID,
        'uid': uid
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
        '/api/v1/org/{orgUID}/agent_registration/{uid}', 'PUT',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }


}
