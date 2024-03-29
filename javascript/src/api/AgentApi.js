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
import Agent from '../model/Agent';

/**
* Agent service.
* @module api/AgentApi
* @version 1.0.0
*/
export default class AgentApi {

    /**
    * Constructs a new AgentApi. 
    * @alias module:api/AgentApi
    * @class
    * @param {module:ApiClient} [apiClient] Optional API client implementation to use,
    * default to {@link module:ApiClient#instance} if unspecified.
    */
    constructor(apiClient) {
        this.apiClient = apiClient || ApiClient.instance;
    }


    /**
     * Callback function to receive the result of the agentList operation.
     * @callback module:api/AgentApi~agentListCallback
     * @param {String} error Error message, if any.
     * @param {Array.<module:model/Agent>} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * List agents
     *  Lists the agents associated with an organization  * Requires the action  *agent:List* on the organization  
     * @param {String} orgUID 
     * @param {Object} opts Optional parameters
     * @param {String} opts.agentRegistrationUidEquals 
     * @param {Boolean} opts.originalAssociation 
     * @param {Number} opts.page 
     * @param {Number} opts.pageSize 
     * @param {String} opts.sourceUidEquals 
     * @param {module:api/AgentApi~agentListCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link Array.<module:model/Agent>}
     */
    agentList(orgUID, opts, callback) {
      opts = opts || {};
      let postBody = null;
      // verify the required parameter 'orgUID' is set
      if (orgUID === undefined || orgUID === null) {
        throw new Error("Missing the required parameter 'orgUID' when calling agentList");
      }

      let pathParams = {
        'orgUID': orgUID
      };
      let queryParams = {
        'agent_registration_uid_equals': opts['agentRegistrationUidEquals'],
        'original_association': opts['originalAssociation'],
        'page': opts['page'],
        'page_size': opts['pageSize'],
        'source_uid_equals': opts['sourceUidEquals']
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = ['apiToken'];
      let contentTypes = [];
      let accepts = ['application/json'];
      let returnType = [Agent];
      return this.apiClient.callApi(
        '/api/v1/org/{orgUID}/agent/', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the agentLoad operation.
     * @callback module:api/AgentApi~agentLoadCallback
     * @param {String} error Error message, if any.
     * @param {module:model/Agent} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Load an agent
     *  Load a specified agent  * Requires the action  *agent:Load* on the organization  
     * @param {String} agentUID 
     * @param {String} orgUID 
     * @param {module:api/AgentApi~agentLoadCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/Agent}
     */
    agentLoad(agentUID, orgUID, callback) {
      let postBody = null;
      // verify the required parameter 'agentUID' is set
      if (agentUID === undefined || agentUID === null) {
        throw new Error("Missing the required parameter 'agentUID' when calling agentLoad");
      }
      // verify the required parameter 'orgUID' is set
      if (orgUID === undefined || orgUID === null) {
        throw new Error("Missing the required parameter 'orgUID' when calling agentLoad");
      }

      let pathParams = {
        'agentUID': agentUID,
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
      let returnType = Agent;
      return this.apiClient.callApi(
        '/api/v1/org/{orgUID}/agent/{agentUID}', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }


}
