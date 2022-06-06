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
import AgentSetAgentWorkInput from '../model/AgentSetAgentWorkInput';
import AgentSetOrgWorkInput from '../model/AgentSetOrgWorkInput';
import ApiAgentWorkOutput from '../model/ApiAgentWorkOutput';
import ValidationError from '../model/ValidationError';

/**
* AgentWork service.
* @module api/AgentWorkApi
* @version 1.0.0
*/
export default class AgentWorkApi {

    /**
    * Constructs a new AgentWorkApi. 
    * @alias module:api/AgentWorkApi
    * @class
    * @param {module:ApiClient} [apiClient] Optional API client implementation to use,
    * default to {@link module:ApiClient#instance} if unspecified.
    */
    constructor(apiClient) {
        this.apiClient = apiClient || ApiClient.instance;
    }


    /**
     * Callback function to receive the result of the agentDeleteAgentWork operation.
     * @callback module:api/AgentWorkApi~agentDeleteAgentWorkCallback
     * @param {String} error Error message, if any.
     * @param data This operation does not return a value.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Delete agent work data for a specific agent
     *  Delet the work data for a specified agent.   * Requires *agent_data:Delete*  
     * @param {String} agentUID Agent UID
     * @param {String} orgUID 
     * @param {module:api/AgentWorkApi~agentDeleteAgentWorkCallback} callback The callback function, accepting three arguments: error, data, response
     */
    agentDeleteAgentWork(agentUID, orgUID, callback) {
      let postBody = null;
      // verify the required parameter 'agentUID' is set
      if (agentUID === undefined || agentUID === null) {
        throw new Error("Missing the required parameter 'agentUID' when calling agentDeleteAgentWork");
      }
      // verify the required parameter 'orgUID' is set
      if (orgUID === undefined || orgUID === null) {
        throw new Error("Missing the required parameter 'orgUID' when calling agentDeleteAgentWork");
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
      let returnType = null;
      return this.apiClient.callApi(
        '/api/v1/org/{orgUID}/agent/{agentUID}/work', 'DELETE',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the agentDeleteOrgWork operation.
     * @callback module:api/AgentWorkApi~agentDeleteOrgWorkCallback
     * @param {String} error Error message, if any.
     * @param data This operation does not return a value.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Delete agent work for an org
     *  Delete the work data for all agents for an organization.   * Requires *agent_data:Delete*  
     * @param {String} agentUID Agent UID
     * @param {String} orgUID 
     * @param {module:api/AgentWorkApi~agentDeleteOrgWorkCallback} callback The callback function, accepting three arguments: error, data, response
     */
    agentDeleteOrgWork(agentUID, orgUID, callback) {
      let postBody = null;
      // verify the required parameter 'agentUID' is set
      if (agentUID === undefined || agentUID === null) {
        throw new Error("Missing the required parameter 'agentUID' when calling agentDeleteOrgWork");
      }
      // verify the required parameter 'orgUID' is set
      if (orgUID === undefined || orgUID === null) {
        throw new Error("Missing the required parameter 'orgUID' when calling agentDeleteOrgWork");
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
      let returnType = null;
      return this.apiClient.callApi(
        '/api/v1/org/{orgUID}/agent_work', 'DELETE',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the agentGetAgentWork operation.
     * @callback module:api/AgentWorkApi~agentGetAgentWorkCallback
     * @param {String} error Error message, if any.
     * @param {module:model/ApiAgentWorkOutput} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Get agent work data for a specific agent
     *  Get the work data for a specified agent.   * Requires *agent_data:GetAgentData* 
     * @param {String} agentUID Agent UID
     * @param {String} orgUID 
     * @param {module:api/AgentWorkApi~agentGetAgentWorkCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/ApiAgentWorkOutput}
     */
    agentGetAgentWork(agentUID, orgUID, callback) {
      let postBody = null;
      // verify the required parameter 'agentUID' is set
      if (agentUID === undefined || agentUID === null) {
        throw new Error("Missing the required parameter 'agentUID' when calling agentGetAgentWork");
      }
      // verify the required parameter 'orgUID' is set
      if (orgUID === undefined || orgUID === null) {
        throw new Error("Missing the required parameter 'orgUID' when calling agentGetAgentWork");
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
      let returnType = ApiAgentWorkOutput;
      return this.apiClient.callApi(
        '/api/v1/org/{orgUID}/agent/{agentUID}/work', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the agentGetOrgWork operation.
     * @callback module:api/AgentWorkApi~agentGetOrgWorkCallback
     * @param {String} error Error message, if any.
     * @param {module:model/ApiAgentWorkOutput} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Get agent work data for the organization
     *  Get the work data for all agents associated with the organization.   * Requires *agent_data:GetOrgData* 
     * @param {String} agentUID Agent UID
     * @param {String} orgUID 
     * @param {module:api/AgentWorkApi~agentGetOrgWorkCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/ApiAgentWorkOutput}
     */
    agentGetOrgWork(agentUID, orgUID, callback) {
      let postBody = null;
      // verify the required parameter 'agentUID' is set
      if (agentUID === undefined || agentUID === null) {
        throw new Error("Missing the required parameter 'agentUID' when calling agentGetOrgWork");
      }
      // verify the required parameter 'orgUID' is set
      if (orgUID === undefined || orgUID === null) {
        throw new Error("Missing the required parameter 'orgUID' when calling agentGetOrgWork");
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
      let returnType = ApiAgentWorkOutput;
      return this.apiClient.callApi(
        '/api/v1/org/{orgUID}/agent_work', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the agentSetAgentWork operation.
     * @callback module:api/AgentWorkApi~agentSetAgentWorkCallback
     * @param {String} error Error message, if any.
     * @param data This operation does not return a value.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Set agent work data for a specific agent
     *  Set the work data for a specified agent.   * Requires *agent_data:Set*  
     * @param {String} agentUID Agent UID
     * @param {String} orgUID 
     * @param {Object} opts Optional parameters
     * @param {module:model/AgentSetAgentWorkInput} opts.agentSetAgentWorkInput 
     * @param {module:api/AgentWorkApi~agentSetAgentWorkCallback} callback The callback function, accepting three arguments: error, data, response
     */
    agentSetAgentWork(agentUID, orgUID, opts, callback) {
      opts = opts || {};
      let postBody = opts['agentSetAgentWorkInput'];
      // verify the required parameter 'agentUID' is set
      if (agentUID === undefined || agentUID === null) {
        throw new Error("Missing the required parameter 'agentUID' when calling agentSetAgentWork");
      }
      // verify the required parameter 'orgUID' is set
      if (orgUID === undefined || orgUID === null) {
        throw new Error("Missing the required parameter 'orgUID' when calling agentSetAgentWork");
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
      let contentTypes = ['application/json'];
      let accepts = ['application/json'];
      let returnType = null;
      return this.apiClient.callApi(
        '/api/v1/org/{orgUID}/agent/{agentUID}/work', 'POST',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the agentSetOrgWork operation.
     * @callback module:api/AgentWorkApi~agentSetOrgWorkCallback
     * @param {String} error Error message, if any.
     * @param data This operation does not return a value.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Set agent work data for a specific agent
     *  Set the work data for a specified agent.   * Requires *agent_data:SetData* 
     * @param {String} agentUID Agent UID
     * @param {String} orgUID 
     * @param {Object} opts Optional parameters
     * @param {module:model/AgentSetOrgWorkInput} opts.agentSetOrgWorkInput 
     * @param {module:api/AgentWorkApi~agentSetOrgWorkCallback} callback The callback function, accepting three arguments: error, data, response
     */
    agentSetOrgWork(agentUID, orgUID, opts, callback) {
      opts = opts || {};
      let postBody = opts['agentSetOrgWorkInput'];
      // verify the required parameter 'agentUID' is set
      if (agentUID === undefined || agentUID === null) {
        throw new Error("Missing the required parameter 'agentUID' when calling agentSetOrgWork");
      }
      // verify the required parameter 'orgUID' is set
      if (orgUID === undefined || orgUID === null) {
        throw new Error("Missing the required parameter 'orgUID' when calling agentSetOrgWork");
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
      let contentTypes = ['application/json'];
      let accepts = ['application/json'];
      let returnType = null;
      return this.apiClient.callApi(
        '/api/v1/org/{orgUID}/agent_work', 'POST',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }


}