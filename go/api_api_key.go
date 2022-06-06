/*
Spyderbat API UI & Public APIs

Restful APIs for use by UI & customers.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sbapi

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


// APIKeyApiService APIKeyApi service
type APIKeyApiService service

type ApiApiKeyCreateRequest struct {
	ctx context.Context
	ApiService *APIKeyApiService
	userUID string
	apiKeyCreateInput *ApiKeyCreateInput
}

func (r ApiApiKeyCreateRequest) ApiKeyCreateInput(apiKeyCreateInput ApiKeyCreateInput) ApiApiKeyCreateRequest {
	r.apiKeyCreateInput = &apiKeyCreateInput
	return r
}

func (r ApiApiKeyCreateRequest) Execute() (string, *http.Response, error) {
	return r.ApiService.ApiKeyCreateExecute(r)
}

/*
ApiKeyCreate Creates a new API key


Creates a new API key which is associated with the user

  * Requires global action *user:APIKeyCreate*


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userUID User UID
 @return ApiApiKeyCreateRequest
*/
func (a *APIKeyApiService) ApiKeyCreate(ctx context.Context, userUID string) ApiApiKeyCreateRequest {
	return ApiApiKeyCreateRequest{
		ApiService: a,
		ctx: ctx,
		userUID: userUID,
	}
}

// Execute executes the request
//  @return string
func (a *APIKeyApiService) ApiKeyCreateExecute(r ApiApiKeyCreateRequest) (string, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "APIKeyApiService.ApiKeyCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/user/{userUID}/apikey/"
	localVarPath = strings.Replace(localVarPath, "{"+"userUID"+"}", url.PathEscape(parameterToString(r.userUID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.userUID) > 64 {
		return localVarReturnValue, nil, reportError("userUID must have less than 64 elements")
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
	localVarPostBody = r.apiKeyCreateInput
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

type ApiApiKeyDeleteRequest struct {
	ctx context.Context
	ApiService *APIKeyApiService
	uid string
	userUID string
}

func (r ApiApiKeyDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiKeyDeleteExecute(r)
}

/*
ApiKeyDelete Delete an API key


Deletes a specific API key

 * Requires global action *user:APIKeyDelete*


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param uid API Key UID
 @param userUID User UID
 @return ApiApiKeyDeleteRequest
*/
func (a *APIKeyApiService) ApiKeyDelete(ctx context.Context, uid string, userUID string) ApiApiKeyDeleteRequest {
	return ApiApiKeyDeleteRequest{
		ApiService: a,
		ctx: ctx,
		uid: uid,
		userUID: userUID,
	}
}

// Execute executes the request
func (a *APIKeyApiService) ApiKeyDeleteExecute(r ApiApiKeyDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "APIKeyApiService.ApiKeyDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/user/{userUID}/apikey/{uid}"
	localVarPath = strings.Replace(localVarPath, "{"+"uid"+"}", url.PathEscape(parameterToString(r.uid, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"userUID"+"}", url.PathEscape(parameterToString(r.userUID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.uid) > 64 {
		return nil, reportError("uid must have less than 64 elements")
	}
	if strlen(r.userUID) > 64 {
		return nil, reportError("userUID must have less than 64 elements")
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

type ApiApiKeyListRequest struct {
	ctx context.Context
	ApiService *APIKeyApiService
	userUID string
}

func (r ApiApiKeyListRequest) Execute() ([]APIKey, *http.Response, error) {
	return r.ApiService.ApiKeyListExecute(r)
}

/*
ApiKeyList Lists an API key


Lists API keys associated with a user

 * Requires global action *user:APIKeyList*



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userUID User UID
 @return ApiApiKeyListRequest
*/
func (a *APIKeyApiService) ApiKeyList(ctx context.Context, userUID string) ApiApiKeyListRequest {
	return ApiApiKeyListRequest{
		ApiService: a,
		ctx: ctx,
		userUID: userUID,
	}
}

// Execute executes the request
//  @return []APIKey
func (a *APIKeyApiService) ApiKeyListExecute(r ApiApiKeyListRequest) ([]APIKey, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []APIKey
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "APIKeyApiService.ApiKeyList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/user/{userUID}/apikey/"
	localVarPath = strings.Replace(localVarPath, "{"+"userUID"+"}", url.PathEscape(parameterToString(r.userUID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.userUID) > 64 {
		return localVarReturnValue, nil, reportError("userUID must have less than 64 elements")
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

type ApiApiKeyUpdateRequest struct {
	ctx context.Context
	ApiService *APIKeyApiService
	uid string
	userUID string
	apiKeyUpdateInput *ApiKeyUpdateInput
}

func (r ApiApiKeyUpdateRequest) ApiKeyUpdateInput(apiKeyUpdateInput ApiKeyUpdateInput) ApiApiKeyUpdateRequest {
	r.apiKeyUpdateInput = &apiKeyUpdateInput
	return r
}

func (r ApiApiKeyUpdateRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiKeyUpdateExecute(r)
}

/*
ApiKeyUpdate Updates an API key


Updates a specific API Key, the only fields which can be updated are description and enabled

 * Requires global action *user:APIKeyUpdate*


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param uid API Key UID
 @param userUID User UID
 @return ApiApiKeyUpdateRequest
*/
func (a *APIKeyApiService) ApiKeyUpdate(ctx context.Context, uid string, userUID string) ApiApiKeyUpdateRequest {
	return ApiApiKeyUpdateRequest{
		ApiService: a,
		ctx: ctx,
		uid: uid,
		userUID: userUID,
	}
}

// Execute executes the request
func (a *APIKeyApiService) ApiKeyUpdateExecute(r ApiApiKeyUpdateRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "APIKeyApiService.ApiKeyUpdate")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/user/{userUID}/apikey/{uid}"
	localVarPath = strings.Replace(localVarPath, "{"+"uid"+"}", url.PathEscape(parameterToString(r.uid, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"userUID"+"}", url.PathEscape(parameterToString(r.userUID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.uid) > 64 {
		return nil, reportError("uid must have less than 64 elements")
	}
	if strlen(r.userUID) > 64 {
		return nil, reportError("userUID must have less than 64 elements")
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
	localVarPostBody = r.apiKeyUpdateInput
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