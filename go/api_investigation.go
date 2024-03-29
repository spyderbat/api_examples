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


// InvestigationApiService InvestigationApi service
type InvestigationApiService service

type ApiInvestigationCreateRequest struct {
	ctx context.Context
	ApiService *InvestigationApiService
	investigationUID string
	orgUID string
	investigationCreateInput *InvestigationCreateInput
}

func (r ApiInvestigationCreateRequest) InvestigationCreateInput(investigationCreateInput InvestigationCreateInput) ApiInvestigationCreateRequest {
	r.investigationCreateInput = &investigationCreateInput
	return r
}

func (r ApiInvestigationCreateRequest) Execute() (*ApiInvestigationCreateOutput, *http.Response, error) {
	return r.ApiService.InvestigationCreateExecute(r)
}

/*
InvestigationCreate Create an investigation


Create an investigationan

 * Requires the user have the action *investigation:Create*


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param investigationUID Investigation UID
 @param orgUID Investigation OrgUID
 @return ApiInvestigationCreateRequest
*/
func (a *InvestigationApiService) InvestigationCreate(ctx context.Context, investigationUID string, orgUID string) ApiInvestigationCreateRequest {
	return ApiInvestigationCreateRequest{
		ApiService: a,
		ctx: ctx,
		investigationUID: investigationUID,
		orgUID: orgUID,
	}
}

// Execute executes the request
//  @return ApiInvestigationCreateOutput
func (a *InvestigationApiService) InvestigationCreateExecute(r ApiInvestigationCreateRequest) (*ApiInvestigationCreateOutput, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiInvestigationCreateOutput
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InvestigationApiService.InvestigationCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/org/{orgUID}/investigation/"
	localVarPath = strings.Replace(localVarPath, "{"+"investigationUID"+"}", url.PathEscape(parameterToString(r.investigationUID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"orgUID"+"}", url.PathEscape(parameterToString(r.orgUID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	localVarPostBody = r.investigationCreateInput
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
			return localVarReturnValue, localVarHTTPResponse, newErr
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

type ApiInvestigationDeleteRequest struct {
	ctx context.Context
	ApiService *InvestigationApiService
	investigationUID string
	orgUID string
}

func (r ApiInvestigationDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.InvestigationDeleteExecute(r)
}

/*
InvestigationDelete Delete an investigation


Deletes an investigation, by setting valid_to=now so that the investigation is virtually deleted.

 * Requires the user have the action *investigation:Delete*


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param investigationUID Investigation UID
 @param orgUID Investigation OrgUID
 @return ApiInvestigationDeleteRequest
*/
func (a *InvestigationApiService) InvestigationDelete(ctx context.Context, investigationUID string, orgUID string) ApiInvestigationDeleteRequest {
	return ApiInvestigationDeleteRequest{
		ApiService: a,
		ctx: ctx,
		investigationUID: investigationUID,
		orgUID: orgUID,
	}
}

// Execute executes the request
func (a *InvestigationApiService) InvestigationDeleteExecute(r ApiInvestigationDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InvestigationApiService.InvestigationDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/org/{orgUID}/investigation/{investigationUID}"
	localVarPath = strings.Replace(localVarPath, "{"+"investigationUID"+"}", url.PathEscape(parameterToString(r.investigationUID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"orgUID"+"}", url.PathEscape(parameterToString(r.orgUID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

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
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiInvestigationListRequest struct {
	ctx context.Context
	ApiService *InvestigationApiService
	orgUID string
}

func (r ApiInvestigationListRequest) Execute() ([]DaoInvestigation, *http.Response, error) {
	return r.ApiService.InvestigationListExecute(r)
}

/*
InvestigationList List investigations


Lists investigations

 * Will list investigations which the user has the action *investigation:Load* or *investigation:LoadExpired* on


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param orgUID
 @return ApiInvestigationListRequest
*/
func (a *InvestigationApiService) InvestigationList(ctx context.Context, orgUID string) ApiInvestigationListRequest {
	return ApiInvestigationListRequest{
		ApiService: a,
		ctx: ctx,
		orgUID: orgUID,
	}
}

// Execute executes the request
//  @return []DaoInvestigation
func (a *InvestigationApiService) InvestigationListExecute(r ApiInvestigationListRequest) ([]DaoInvestigation, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []DaoInvestigation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InvestigationApiService.InvestigationList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/org/{orgUID}/investigation/"
	localVarPath = strings.Replace(localVarPath, "{"+"orgUID"+"}", url.PathEscape(parameterToString(r.orgUID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
			return localVarReturnValue, localVarHTTPResponse, newErr
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

type ApiInvestigationListVersionsRequest struct {
	ctx context.Context
	ApiService *InvestigationApiService
	investigationUID string
	orgUID string
}

func (r ApiInvestigationListVersionsRequest) Execute() ([]DaoInvestigation, *http.Response, error) {
	return r.ApiService.InvestigationListVersionsExecute(r)
}

/*
InvestigationListVersions List Investigation Versions


Lists prior version of this investigation

 * Requires the user have the action *investigation:ListVersions*


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param investigationUID
 @param orgUID
 @return ApiInvestigationListVersionsRequest
*/
func (a *InvestigationApiService) InvestigationListVersions(ctx context.Context, investigationUID string, orgUID string) ApiInvestigationListVersionsRequest {
	return ApiInvestigationListVersionsRequest{
		ApiService: a,
		ctx: ctx,
		investigationUID: investigationUID,
		orgUID: orgUID,
	}
}

// Execute executes the request
//  @return []DaoInvestigation
func (a *InvestigationApiService) InvestigationListVersionsExecute(r ApiInvestigationListVersionsRequest) ([]DaoInvestigation, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []DaoInvestigation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InvestigationApiService.InvestigationListVersions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/org/{orgUID}/investigation/{investigationUID}/version/"
	localVarPath = strings.Replace(localVarPath, "{"+"investigationUID"+"}", url.PathEscape(parameterToString(r.investigationUID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"orgUID"+"}", url.PathEscape(parameterToString(r.orgUID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ApiInvestigationLoadRequest struct {
	ctx context.Context
	ApiService *InvestigationApiService
	investigationUID string
	orgUID string
}

func (r ApiInvestigationLoadRequest) Execute() (*DaoInvestigation, *http.Response, error) {
	return r.ApiService.InvestigationLoadExecute(r)
}

/*
InvestigationLoad Load an investigation


Loads an investigation by UID. 

 * Requires action  *investigation:Load* to load an active investigation
 * Requires action *investigation:LoadExpired* to load expired investigations



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param investigationUID
 @param orgUID
 @return ApiInvestigationLoadRequest
*/
func (a *InvestigationApiService) InvestigationLoad(ctx context.Context, investigationUID string, orgUID string) ApiInvestigationLoadRequest {
	return ApiInvestigationLoadRequest{
		ApiService: a,
		ctx: ctx,
		investigationUID: investigationUID,
		orgUID: orgUID,
	}
}

// Execute executes the request
//  @return DaoInvestigation
func (a *InvestigationApiService) InvestigationLoadExecute(r ApiInvestigationLoadRequest) (*DaoInvestigation, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DaoInvestigation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InvestigationApiService.InvestigationLoad")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/org/{orgUID}/investigation/{investigationUID}"
	localVarPath = strings.Replace(localVarPath, "{"+"investigationUID"+"}", url.PathEscape(parameterToString(r.investigationUID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"orgUID"+"}", url.PathEscape(parameterToString(r.orgUID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.investigationUID) > 64 {
		return localVarReturnValue, nil, reportError("investigationUID must have less than 64 elements")
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

type ApiInvestigationLoadVersionRequest struct {
	ctx context.Context
	ApiService *InvestigationApiService
	investigationUID string
	orgUID string
	version int32
}

func (r ApiInvestigationLoadVersionRequest) Execute() (*DaoInvestigation, *http.Response, error) {
	return r.ApiService.InvestigationLoadVersionExecute(r)
}

/*
InvestigationLoadVersion Load Investigation Version


Loads a specific version of an investigation

 * Requires the user have the action *investigation:LoadVersion*


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param investigationUID
 @param orgUID
 @param version
 @return ApiInvestigationLoadVersionRequest
*/
func (a *InvestigationApiService) InvestigationLoadVersion(ctx context.Context, investigationUID string, orgUID string, version int32) ApiInvestigationLoadVersionRequest {
	return ApiInvestigationLoadVersionRequest{
		ApiService: a,
		ctx: ctx,
		investigationUID: investigationUID,
		orgUID: orgUID,
		version: version,
	}
}

// Execute executes the request
//  @return DaoInvestigation
func (a *InvestigationApiService) InvestigationLoadVersionExecute(r ApiInvestigationLoadVersionRequest) (*DaoInvestigation, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DaoInvestigation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InvestigationApiService.InvestigationLoadVersion")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/org/{orgUID}/investigation/{investigationUID}/version/{version}"
	localVarPath = strings.Replace(localVarPath, "{"+"investigationUID"+"}", url.PathEscape(parameterToString(r.investigationUID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"orgUID"+"}", url.PathEscape(parameterToString(r.orgUID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"version"+"}", url.PathEscape(parameterToString(r.version, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ApiInvestigationUpdateRequest struct {
	ctx context.Context
	ApiService *InvestigationApiService
	investigationUID string
	orgUID string
	investigationUpdateInput *InvestigationUpdateInput
}

func (r ApiInvestigationUpdateRequest) InvestigationUpdateInput(investigationUpdateInput InvestigationUpdateInput) ApiInvestigationUpdateRequest {
	r.investigationUpdateInput = &investigationUpdateInput
	return r
}

func (r ApiInvestigationUpdateRequest) Execute() (*http.Response, error) {
	return r.ApiService.InvestigationUpdateExecute(r)
}

/*
InvestigationUpdate Update an investigation


Updates the investigationan

 * Requires the user have the action *investigation:Update*


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param investigationUID Investigation UID
 @param orgUID Investigation OrgUID
 @return ApiInvestigationUpdateRequest
*/
func (a *InvestigationApiService) InvestigationUpdate(ctx context.Context, investigationUID string, orgUID string) ApiInvestigationUpdateRequest {
	return ApiInvestigationUpdateRequest{
		ApiService: a,
		ctx: ctx,
		investigationUID: investigationUID,
		orgUID: orgUID,
	}
}

// Execute executes the request
func (a *InvestigationApiService) InvestigationUpdateExecute(r ApiInvestigationUpdateRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InvestigationApiService.InvestigationUpdate")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/org/{orgUID}/investigation/{investigationUID}"
	localVarPath = strings.Replace(localVarPath, "{"+"investigationUID"+"}", url.PathEscape(parameterToString(r.investigationUID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"orgUID"+"}", url.PathEscape(parameterToString(r.orgUID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	localVarPostBody = r.investigationUpdateInput
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
			return localVarHTTPResponse, newErr
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
