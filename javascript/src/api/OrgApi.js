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
import DaoOrgRoleResponse from '../model/DaoOrgRoleResponse';
import NotificationPolicy from '../model/NotificationPolicy';
import Org from '../model/Org';
import OrgAssignRoleInput from '../model/OrgAssignRoleInput';
import OrgInviteUsersInput from '../model/OrgInviteUsersInput';
import OrgTestNotificationTargetInput from '../model/OrgTestNotificationTargetInput';
import OrgUnassignRoleInput from '../model/OrgUnassignRoleInput';
import OrgUpdateInput from '../model/OrgUpdateInput';
import ValidationError from '../model/ValidationError';

/**
* Org service.
* @module api/OrgApi
* @version 1.0.0
*/
export default class OrgApi {

    /**
    * Constructs a new OrgApi. 
    * @alias module:api/OrgApi
    * @class
    * @param {module:ApiClient} [apiClient] Optional API client implementation to use,
    * default to {@link module:ApiClient#instance} if unspecified.
    */
    constructor(apiClient) {
        this.apiClient = apiClient || ApiClient.instance;
    }


    /**
     * Callback function to receive the result of the orgAssignRole operation.
     * @callback module:api/OrgApi~orgAssignRoleCallback
     * @param {String} error Error message, if any.
     * @param data This operation does not return a value.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Assign OrgRole
     *  Assigns a role to a particular user on an organization   * Requires the user have the action *user.AssignRole* on the organization 
     * @param {String} orgUID 
     * @param {Object} opts Optional parameters
     * @param {module:model/OrgAssignRoleInput} opts.orgAssignRoleInput 
     * @param {module:api/OrgApi~orgAssignRoleCallback} callback The callback function, accepting three arguments: error, data, response
     */
    orgAssignRole(orgUID, opts, callback) {
      opts = opts || {};
      let postBody = opts['orgAssignRoleInput'];
      // verify the required parameter 'orgUID' is set
      if (orgUID === undefined || orgUID === null) {
        throw new Error("Missing the required parameter 'orgUID' when calling orgAssignRole");
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
      let accepts = [];
      let returnType = null;
      return this.apiClient.callApi(
        '/api/v1/org/{orgUID}/assignedrole/add', 'POST',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the orgInviteUsers operation.
     * @callback module:api/OrgApi~orgInviteUsersCallback
     * @param {String} error Error message, if any.
     * @param data This operation does not return a value.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Invite users to an organization
     *  Invites users to an organization   * Requires action *org:InviteUsers* on the organization to invite users  
     * @param {String} orgUID 
     * @param {Object} opts Optional parameters
     * @param {module:model/OrgInviteUsersInput} opts.orgInviteUsersInput 
     * @param {module:api/OrgApi~orgInviteUsersCallback} callback The callback function, accepting three arguments: error, data, response
     */
    orgInviteUsers(orgUID, opts, callback) {
      opts = opts || {};
      let postBody = opts['orgInviteUsersInput'];
      // verify the required parameter 'orgUID' is set
      if (orgUID === undefined || orgUID === null) {
        throw new Error("Missing the required parameter 'orgUID' when calling orgInviteUsers");
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
      let accepts = [];
      let returnType = null;
      return this.apiClient.callApi(
        '/api/v1/org/{orgUID}/invite', 'POST',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the orgList operation.
     * @callback module:api/OrgApi~orgListCallback
     * @param {String} error Error message, if any.
     * @param {Array.<module:model/Org>} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * List organizations
     *  Lists organizations   * Will list organizations which the user has the action *org:Load* or *org:LoadExpired* on 
     * @param {Object} opts Optional parameters
     * @param {Boolean} opts.hasResourcePolicy 
     * @param {Array.<String>} opts.hasTags 
     * @param {Boolean} opts.includeExpired 
     * @param {String} opts.nameContains 
     * @param {String} opts.ownerUidEquals 
     * @param {Number} opts.page 
     * @param {Number} opts.pageSize 
     * @param {String} opts.uidEquals 
     * @param {module:api/OrgApi~orgListCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link Array.<module:model/Org>}
     */
    orgList(opts, callback) {
      opts = opts || {};
      let postBody = null;

      let pathParams = {
      };
      let queryParams = {
        'has_resource_policy': opts['hasResourcePolicy'],
        'has_tags': this.apiClient.buildCollectionParam(opts['hasTags'], 'multi'),
        'include_expired': opts['includeExpired'],
        'name_contains': opts['nameContains'],
        'owner_uid_equals': opts['ownerUidEquals'],
        'page': opts['page'],
        'page_size': opts['pageSize'],
        'uid_equals': opts['uidEquals']
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = ['apiToken'];
      let contentTypes = [];
      let accepts = ['application/json'];
      let returnType = [Org];
      return this.apiClient.callApi(
        '/api/v1/org/', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the orgListRole operation.
     * @callback module:api/OrgApi~orgListRoleCallback
     * @param {String} error Error message, if any.
     * @param {Array.<module:model/DaoOrgRoleResponse>} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * List OrgRoles
     *  Allows querying of roles on the organization   * Requires the user have the action *org:ListOrgRoles* on the organization 
     * @param {String} orgUID 
     * @param {Object} opts Optional parameters
     * @param {String} opts.roleUidEquals 
     * @param {String} opts.userEmailEquals 
     * @param {String} opts.userUidEquals 
     * @param {module:api/OrgApi~orgListRoleCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link Array.<module:model/DaoOrgRoleResponse>}
     */
    orgListRole(orgUID, opts, callback) {
      opts = opts || {};
      let postBody = null;
      // verify the required parameter 'orgUID' is set
      if (orgUID === undefined || orgUID === null) {
        throw new Error("Missing the required parameter 'orgUID' when calling orgListRole");
      }

      let pathParams = {
        'orgUID': orgUID
      };
      let queryParams = {
        'role_uid_equals': opts['roleUidEquals'],
        'user_email_equals': opts['userEmailEquals'],
        'user_uid_equals': opts['userUidEquals']
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = ['apiToken'];
      let contentTypes = [];
      let accepts = ['application/json'];
      let returnType = [DaoOrgRoleResponse];
      return this.apiClient.callApi(
        '/api/v1/org/{orgUID}/assignedrole/', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the orgLoad operation.
     * @callback module:api/OrgApi~orgLoadCallback
     * @param {String} error Error message, if any.
     * @param {module:model/Org} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Load an organization
     *  Loads an organization by UID.    * Requires action *org:Load*  * Requires action *org:LoadExpired* to load expired organizations  
     * @param {String} orgUID 
     * @param {module:api/OrgApi~orgLoadCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/Org}
     */
    orgLoad(orgUID, callback) {
      let postBody = null;
      // verify the required parameter 'orgUID' is set
      if (orgUID === undefined || orgUID === null) {
        throw new Error("Missing the required parameter 'orgUID' when calling orgLoad");
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
      let returnType = Org;
      return this.apiClient.callApi(
        '/api/v1/org/{orgUID}', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the orgLoadNotificationPolicy operation.
     * @callback module:api/OrgApi~orgLoadNotificationPolicyCallback
     * @param {String} error Error message, if any.
     * @param {module:model/NotificationPolicy} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Load Notification Policy
     *  Loads the notification policy for an organization. The notification policy defines who and how the organization is notified.     * If the content-type is application/hjson the policy will be returned as hjson  * If the content-type is application/json the policy will be returned as json    * Requires the user have the action *org:LoadNotificationPolicy* on the organization 
     * @param {String} orgUID 
     * @param {module:api/OrgApi~orgLoadNotificationPolicyCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/NotificationPolicy}
     */
    orgLoadNotificationPolicy(orgUID, callback) {
      let postBody = null;
      // verify the required parameter 'orgUID' is set
      if (orgUID === undefined || orgUID === null) {
        throw new Error("Missing the required parameter 'orgUID' when calling orgLoadNotificationPolicy");
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
      let accepts = ['application/hjson', 'application/json'];
      let returnType = NotificationPolicy;
      return this.apiClient.callApi(
        '/api/v1/org/{orgUID}/notification_policy/', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the orgTestNotificationTarget operation.
     * @callback module:api/OrgApi~orgTestNotificationTargetCallback
     * @param {String} error Error message, if any.
     * @param data This operation does not return a value.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Test Notification Target
     *  Sends a test notification to a target.   * Requires the user have the action *org:SendTestNotification* on the organization 
     * @param {String} orgUID 
     * @param {Object} opts Optional parameters
     * @param {module:model/OrgTestNotificationTargetInput} opts.orgTestNotificationTargetInput 
     * @param {module:api/OrgApi~orgTestNotificationTargetCallback} callback The callback function, accepting three arguments: error, data, response
     */
    orgTestNotificationTarget(orgUID, opts, callback) {
      opts = opts || {};
      let postBody = opts['orgTestNotificationTargetInput'];
      // verify the required parameter 'orgUID' is set
      if (orgUID === undefined || orgUID === null) {
        throw new Error("Missing the required parameter 'orgUID' when calling orgTestNotificationTarget");
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
      let accepts = [];
      let returnType = null;
      return this.apiClient.callApi(
        '/api/v1/org/{orgUID}/notification_policy/test_target', 'POST',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the orgUnassignRole operation.
     * @callback module:api/OrgApi~orgUnassignRoleCallback
     * @param {String} error Error message, if any.
     * @param data This operation does not return a value.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Unassign OrgRole
     *  Unassigns a role to a particular user on an organization   * Requires the user have the action *user.UnassignRole* on the organization 
     * @param {String} orgUID 
     * @param {Object} opts Optional parameters
     * @param {module:model/OrgUnassignRoleInput} opts.orgUnassignRoleInput 
     * @param {module:api/OrgApi~orgUnassignRoleCallback} callback The callback function, accepting three arguments: error, data, response
     */
    orgUnassignRole(orgUID, opts, callback) {
      opts = opts || {};
      let postBody = opts['orgUnassignRoleInput'];
      // verify the required parameter 'orgUID' is set
      if (orgUID === undefined || orgUID === null) {
        throw new Error("Missing the required parameter 'orgUID' when calling orgUnassignRole");
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
      let accepts = [];
      let returnType = null;
      return this.apiClient.callApi(
        '/api/v1/org/{orgUID}/assignedrole/del', 'POST',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the orgUpdate operation.
     * @callback module:api/OrgApi~orgUpdateCallback
     * @param {String} error Error message, if any.
     * @param data This operation does not return a value.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Update an organization
     *  Updates the organization    * Requires the user have the action *org:Update* on the organization 
     * @param {String} orgUID Org UID
     * @param {Object} opts Optional parameters
     * @param {module:model/OrgUpdateInput} opts.orgUpdateInput 
     * @param {module:api/OrgApi~orgUpdateCallback} callback The callback function, accepting three arguments: error, data, response
     */
    orgUpdate(orgUID, opts, callback) {
      opts = opts || {};
      let postBody = opts['orgUpdateInput'];
      // verify the required parameter 'orgUID' is set
      if (orgUID === undefined || orgUID === null) {
        throw new Error("Missing the required parameter 'orgUID' when calling orgUpdate");
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
      let returnType = null;
      return this.apiClient.callApi(
        '/api/v1/org/{orgUID}', 'PUT',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the orgUpdateNotificationPolicy operation.
     * @callback module:api/OrgApi~orgUpdateNotificationPolicyCallback
     * @param {String} error Error message, if any.
     * @param data This operation does not return a value.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Update an organization's notification policy
     *  Updates the organization's notification policy  The policy contains targets, which are named destinations, for example a list of admins, and then rules for how notifications are routed. The most basic notification is the default policy assigned to all organizations, which is to send notifications to the organization owner.   Each notification has a schema which can be used to filter how notifications are routed.   Here is an example notification policy:  ``` {     \"targets\": {       \"admins\": {         \"emails\": [           \"admin1@foo.com\",           \"admin2@foo.com\"         ]       },       \"soc\": {         \"slack\": {           \"url\": \"http://app.slack.com/XXXX\"         }       },       \"sns_example\": {         \"data\": \"UI supplied data\",         \"description\": \"Cross-Account Notification\",         \"sns\": {           \"sns_topic_arn\": \"arn:aws:sns:CUSTOMER-AWS-REGION:CUSTOMER-AWS-ACCOUNT-ID:CUSTOMER-TOPIC-NAME\",           \"cross_account_iam_role\": \"arn:aws:iam::CUSTOMER-ACCOUNT-ID:role/CUSTOMER-ROLE-NAME\"         }       }     },     \"routes\": [       {         \"target\": \"admins\",         \"expr\": {           \"schema\": \"agent_offline\"         }       },       {         \"target\": \"soc\",         \"expr\": {           \"schema\": \"spydertrace_updated\"         }       },       {         \"destination\": {           \"users\": [             \"X23hs8234lks\"           ]         }       }     ]   } ```  This policy says that any notification with the schema \"agent_offline\" is send to the admin emails, and that any notifications with the schema \"spydertrace_update\" is send to the soc slack channel, and all other notifications are sent to a user specified by their UserUID. If a destination is an explicit UserUID then the users notification policy will applied, for example to notify them by their notification type of choice.     * The following destination types are currently supported, UserUID, Email, Slack, Webhook, SNS. See the associated destination definition in the notification policy for details.   * First match for a routing rule wins, and further processsing of the notification stops.   * If a UserUID is specified the users notification policy may be used to contact that user.  * Each notification has an associated schema, which is used to provide a consistent schema for the notification.  * Expressions are optional, if no expression is specified the route matches by default  * If the content-type is application/hjson the policy will be parsed and stored as hjson, all comments will be lost on an existing hjson policy if it is uploaded as json  This is an example notification generated using a notification policy and dashboardsearch:  ``` {  \"uid\": \"lQ0Q1lKm\",  \"org_uid\": \"your_org_uid\",  \"valid_from\": \"2021-10-14T19:22:00.869159169Z\",  \"title\": \"Spyderbat: Dashboard search notification Recent interactive (shell) Processes\",  \"message\": \"Spyderbat: Dashboard search notification Recent interactive (shell) Processes\",  \"data\": {    \"dashboardsearch\": {   \"data\": {     \"createTime\": 1634237497.142,     \"createdBy\": \"user@example.com\"   },   \"description\": \"Recent interactive (shell) Processes\",   \"notify\": true,   \"notify_frequency\": 300,   \"org_uid\": \"spyderbatuid\",   \"search\": \"schema:model_process AND interactive:true\",   \"uid\": \"RTBQoR3uucyjG8ZeHJmw\"    }  },  \"schema\": \"dashboard_saved_search\",  \"records\": [    {   \"total_hits\": 0    }  ]   } ```   * Requires the user have the action *org:UpdateNotificationPolicy* on the organization 
     * @param {String} orgUID 
     * @param {module:model/NotificationPolicy} notificationPolicy The notification policy
     * @param {module:api/OrgApi~orgUpdateNotificationPolicyCallback} callback The callback function, accepting three arguments: error, data, response
     */
    orgUpdateNotificationPolicy(orgUID, notificationPolicy, callback) {
      let postBody = notificationPolicy;
      // verify the required parameter 'orgUID' is set
      if (orgUID === undefined || orgUID === null) {
        throw new Error("Missing the required parameter 'orgUID' when calling orgUpdateNotificationPolicy");
      }
      // verify the required parameter 'notificationPolicy' is set
      if (notificationPolicy === undefined || notificationPolicy === null) {
        throw new Error("Missing the required parameter 'notificationPolicy' when calling orgUpdateNotificationPolicy");
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
      let returnType = null;
      return this.apiClient.callApi(
        '/api/v1/org/{orgUID}/notification_policy', 'PUT',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }


}
