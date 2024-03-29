/*
Spyderbat API UI & Public APIs

Restful APIs for use by UI & customers.

API version: 1.0.0
Contact: apisupport@spyderbat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spyderbat_api

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


// AgentWorkApiService AgentWorkApi service
type AgentWorkApiService service

type ApiAgentDeleteAgentWorkRequest struct {
	ctx context.Context
	ApiService *AgentWorkApiService
	agentUID string
	orgUID string
}

func (r ApiAgentDeleteAgentWorkRequest) Execute() (*http.Response, error) {
	return r.ApiService.AgentDeleteAgentWorkExecute(r)
}

/*
AgentDeleteAgentWork Delete agent work data for a specific agent


Delet the work data for a specified agent.

 * Requires *agent_data:Delete* 


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param agentUID Agent UID
 @param orgUID
 @return ApiAgentDeleteAgentWorkRequest
*/
func (a *AgentWorkApiService) AgentDeleteAgentWork(ctx context.Context, agentUID string, orgUID string) ApiAgentDeleteAgentWorkRequest {
	return ApiAgentDeleteAgentWorkRequest{
		ApiService: a,
		ctx: ctx,
		agentUID: agentUID,
		orgUID: orgUID,
	}
}

// Execute executes the request
func (a *AgentWorkApiService) AgentDeleteAgentWorkExecute(r ApiAgentDeleteAgentWorkRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgentWorkApiService.AgentDeleteAgentWork")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/org/{orgUID}/agent/{agentUID}/work"
	localVarPath = strings.Replace(localVarPath, "{"+"agentUID"+"}", url.PathEscape(parameterToString(r.agentUID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"orgUID"+"}", url.PathEscape(parameterToString(r.orgUID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.agentUID) > 20 {
		return nil, reportError("agentUID must have less than 20 elements")
	}
	if strlen(r.orgUID) > 64 {
		return nil, reportError("orgUID must have less than 64 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ValidationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiAgentDeleteOrgWorkRequest struct {
	ctx context.Context
	ApiService *AgentWorkApiService
	agentUID string
	orgUID string
}

func (r ApiAgentDeleteOrgWorkRequest) Execute() (*http.Response, error) {
	return r.ApiService.AgentDeleteOrgWorkExecute(r)
}

/*
AgentDeleteOrgWork Delete agent work for an org


Delete the work data for all agents for an organization.

 * Requires *agent_data:Delete* 


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param agentUID Agent UID
 @param orgUID
 @return ApiAgentDeleteOrgWorkRequest
*/
func (a *AgentWorkApiService) AgentDeleteOrgWork(ctx context.Context, agentUID string, orgUID string) ApiAgentDeleteOrgWorkRequest {
	return ApiAgentDeleteOrgWorkRequest{
		ApiService: a,
		ctx: ctx,
		agentUID: agentUID,
		orgUID: orgUID,
	}
}

// Execute executes the request
func (a *AgentWorkApiService) AgentDeleteOrgWorkExecute(r ApiAgentDeleteOrgWorkRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgentWorkApiService.AgentDeleteOrgWork")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/org/{orgUID}/agent_work"
	localVarPath = strings.Replace(localVarPath, "{"+"agentUID"+"}", url.PathEscape(parameterToString(r.agentUID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"orgUID"+"}", url.PathEscape(parameterToString(r.orgUID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.agentUID) > 20 {
		return nil, reportError("agentUID must have less than 20 elements")
	}
	if strlen(r.orgUID) > 64 {
		return nil, reportError("orgUID must have less than 64 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ValidationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiAgentGetAgentWorkRequest struct {
	ctx context.Context
	ApiService *AgentWorkApiService
	agentUID string
	orgUID string
}

func (r ApiAgentGetAgentWorkRequest) Execute() (*ApiAgentWorkOutput, *http.Response, error) {
	return r.ApiService.AgentGetAgentWorkExecute(r)
}

/*
AgentGetAgentWork Get agent work data for a specific agent


Get the work data for a specified agent.

 * Requires *agent_data:GetAgentData*


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param agentUID Agent UID
 @param orgUID
 @return ApiAgentGetAgentWorkRequest
*/
func (a *AgentWorkApiService) AgentGetAgentWork(ctx context.Context, agentUID string, orgUID string) ApiAgentGetAgentWorkRequest {
	return ApiAgentGetAgentWorkRequest{
		ApiService: a,
		ctx: ctx,
		agentUID: agentUID,
		orgUID: orgUID,
	}
}

// Execute executes the request
//  @return ApiAgentWorkOutput
func (a *AgentWorkApiService) AgentGetAgentWorkExecute(r ApiAgentGetAgentWorkRequest) (*ApiAgentWorkOutput, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiAgentWorkOutput
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgentWorkApiService.AgentGetAgentWork")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/org/{orgUID}/agent/{agentUID}/work"
	localVarPath = strings.Replace(localVarPath, "{"+"agentUID"+"}", url.PathEscape(parameterToString(r.agentUID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"orgUID"+"}", url.PathEscape(parameterToString(r.orgUID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.agentUID) > 20 {
		return localVarReturnValue, nil, reportError("agentUID must have less than 20 elements")
	}
	if strlen(r.orgUID) > 64 {
		return localVarReturnValue, nil, reportError("orgUID must have less than 64 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ValidationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiAgentGetOrgWorkRequest struct {
	ctx context.Context
	ApiService *AgentWorkApiService
	agentUID string
	orgUID string
}

func (r ApiAgentGetOrgWorkRequest) Execute() (*ApiAgentWorkOutput, *http.Response, error) {
	return r.ApiService.AgentGetOrgWorkExecute(r)
}

/*
AgentGetOrgWork Get agent work data for the organization


Get the work data for all agents associated with the organization.

 * Requires *agent_data:GetOrgData*


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param agentUID Agent UID
 @param orgUID
 @return ApiAgentGetOrgWorkRequest
*/
func (a *AgentWorkApiService) AgentGetOrgWork(ctx context.Context, agentUID string, orgUID string) ApiAgentGetOrgWorkRequest {
	return ApiAgentGetOrgWorkRequest{
		ApiService: a,
		ctx: ctx,
		agentUID: agentUID,
		orgUID: orgUID,
	}
}

// Execute executes the request
//  @return ApiAgentWorkOutput
func (a *AgentWorkApiService) AgentGetOrgWorkExecute(r ApiAgentGetOrgWorkRequest) (*ApiAgentWorkOutput, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiAgentWorkOutput
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgentWorkApiService.AgentGetOrgWork")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/org/{orgUID}/agent_work"
	localVarPath = strings.Replace(localVarPath, "{"+"agentUID"+"}", url.PathEscape(parameterToString(r.agentUID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"orgUID"+"}", url.PathEscape(parameterToString(r.orgUID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.agentUID) > 20 {
		return localVarReturnValue, nil, reportError("agentUID must have less than 20 elements")
	}
	if strlen(r.orgUID) > 64 {
		return localVarReturnValue, nil, reportError("orgUID must have less than 64 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ValidationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiAgentSetAgentWorkRequest struct {
	ctx context.Context
	ApiService *AgentWorkApiService
	agentUID string
	orgUID string
	agentSetAgentWorkInput *AgentSetAgentWorkInput
}

func (r ApiAgentSetAgentWorkRequest) AgentSetAgentWorkInput(agentSetAgentWorkInput AgentSetAgentWorkInput) ApiAgentSetAgentWorkRequest {
	r.agentSetAgentWorkInput = &agentSetAgentWorkInput
	return r
}

func (r ApiAgentSetAgentWorkRequest) Execute() (*http.Response, error) {
	return r.ApiService.AgentSetAgentWorkExecute(r)
}

/*
AgentSetAgentWork Set agent work data for a specific agent


Set the work data for a specified agent.

 * Requires *agent_data:Set* 


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param agentUID Agent UID
 @param orgUID
 @return ApiAgentSetAgentWorkRequest
*/
func (a *AgentWorkApiService) AgentSetAgentWork(ctx context.Context, agentUID string, orgUID string) ApiAgentSetAgentWorkRequest {
	return ApiAgentSetAgentWorkRequest{
		ApiService: a,
		ctx: ctx,
		agentUID: agentUID,
		orgUID: orgUID,
	}
}

// Execute executes the request
func (a *AgentWorkApiService) AgentSetAgentWorkExecute(r ApiAgentSetAgentWorkRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgentWorkApiService.AgentSetAgentWork")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/org/{orgUID}/agent/{agentUID}/work"
	localVarPath = strings.Replace(localVarPath, "{"+"agentUID"+"}", url.PathEscape(parameterToString(r.agentUID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"orgUID"+"}", url.PathEscape(parameterToString(r.orgUID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.agentUID) > 20 {
		return nil, reportError("agentUID must have less than 20 elements")
	}
	if strlen(r.orgUID) > 64 {
		return nil, reportError("orgUID must have less than 64 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.agentSetAgentWorkInput
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ValidationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiAgentSetOrgWorkRequest struct {
	ctx context.Context
	ApiService *AgentWorkApiService
	agentUID string
	orgUID string
	agentSetOrgWorkInput *AgentSetOrgWorkInput
}

func (r ApiAgentSetOrgWorkRequest) AgentSetOrgWorkInput(agentSetOrgWorkInput AgentSetOrgWorkInput) ApiAgentSetOrgWorkRequest {
	r.agentSetOrgWorkInput = &agentSetOrgWorkInput
	return r
}

func (r ApiAgentSetOrgWorkRequest) Execute() (*http.Response, error) {
	return r.ApiService.AgentSetOrgWorkExecute(r)
}

/*
AgentSetOrgWork Set agent work data for a specific agent


Set the work data for a specified agent.

 * Requires *agent_data:SetData*


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param agentUID Agent UID
 @param orgUID
 @return ApiAgentSetOrgWorkRequest
*/
func (a *AgentWorkApiService) AgentSetOrgWork(ctx context.Context, agentUID string, orgUID string) ApiAgentSetOrgWorkRequest {
	return ApiAgentSetOrgWorkRequest{
		ApiService: a,
		ctx: ctx,
		agentUID: agentUID,
		orgUID: orgUID,
	}
}

// Execute executes the request
func (a *AgentWorkApiService) AgentSetOrgWorkExecute(r ApiAgentSetOrgWorkRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgentWorkApiService.AgentSetOrgWork")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/org/{orgUID}/agent_work"
	localVarPath = strings.Replace(localVarPath, "{"+"agentUID"+"}", url.PathEscape(parameterToString(r.agentUID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"orgUID"+"}", url.PathEscape(parameterToString(r.orgUID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.agentUID) > 20 {
		return nil, reportError("agentUID must have less than 20 elements")
	}
	if strlen(r.orgUID) > 64 {
		return nil, reportError("orgUID must have less than 64 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.agentSetOrgWorkInput
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ValidationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
