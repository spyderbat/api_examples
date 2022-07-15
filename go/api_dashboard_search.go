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


// DashboardSearchApiService DashboardSearchApi service
type DashboardSearchApiService service

type ApiDashboardSearchCreateRequest struct {
	ctx context.Context
	ApiService *DashboardSearchApiService
	dashboardSearchUID string
	orgUID string
	dashboardSearchCreateInput *DashboardSearchCreateInput
}

func (r ApiDashboardSearchCreateRequest) DashboardSearchCreateInput(dashboardSearchCreateInput DashboardSearchCreateInput) ApiDashboardSearchCreateRequest {
	r.dashboardSearchCreateInput = &dashboardSearchCreateInput
	return r
}

func (r ApiDashboardSearchCreateRequest) Execute() (string, *http.Response, error) {
	return r.ApiService.DashboardSearchCreateExecute(r)
}

/*
DashboardSearchCreate Create a dashboard search


Create a dashboard search in an org.

 * Requires action dashboard_search:Create

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param dashboardSearchUID UID for the DashboardSearch
 @param orgUID Org UID
 @return ApiDashboardSearchCreateRequest
*/
func (a *DashboardSearchApiService) DashboardSearchCreate(ctx context.Context, dashboardSearchUID string, orgUID string) ApiDashboardSearchCreateRequest {
	return ApiDashboardSearchCreateRequest{
		ApiService: a,
		ctx: ctx,
		dashboardSearchUID: dashboardSearchUID,
		orgUID: orgUID,
	}
}

// Execute executes the request
//  @return string
func (a *DashboardSearchApiService) DashboardSearchCreateExecute(r ApiDashboardSearchCreateRequest) (string, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DashboardSearchApiService.DashboardSearchCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/org/{orgUID}/dashboard_search/"
	localVarPath = strings.Replace(localVarPath, "{"+"dashboardSearchUID"+"}", url.PathEscape(parameterToString(r.dashboardSearchUID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"orgUID"+"}", url.PathEscape(parameterToString(r.orgUID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.orgUID) > 32 {
		return localVarReturnValue, nil, reportError("orgUID must have less than 32 elements")
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
	localVarPostBody = r.dashboardSearchCreateInput
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

type ApiDashboardSearchDeleteRequest struct {
	ctx context.Context
	ApiService *DashboardSearchApiService
	dashboardSearchUID string
	orgUID string
}

func (r ApiDashboardSearchDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.DashboardSearchDeleteExecute(r)
}

/*
DashboardSearchDelete Get a dashboard search


Get a dashboard search in an org.

 * Requires action dashboard_search:Delete

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param dashboardSearchUID
 @param orgUID
 @return ApiDashboardSearchDeleteRequest
*/
func (a *DashboardSearchApiService) DashboardSearchDelete(ctx context.Context, dashboardSearchUID string, orgUID string) ApiDashboardSearchDeleteRequest {
	return ApiDashboardSearchDeleteRequest{
		ApiService: a,
		ctx: ctx,
		dashboardSearchUID: dashboardSearchUID,
		orgUID: orgUID,
	}
}

// Execute executes the request
func (a *DashboardSearchApiService) DashboardSearchDeleteExecute(r ApiDashboardSearchDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DashboardSearchApiService.DashboardSearchDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/org/{orgUID}/dashboard_search/{dashboardSearchUID}"
	localVarPath = strings.Replace(localVarPath, "{"+"dashboardSearchUID"+"}", url.PathEscape(parameterToString(r.dashboardSearchUID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"orgUID"+"}", url.PathEscape(parameterToString(r.orgUID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.dashboardSearchUID) > 32 {
		return nil, reportError("dashboardSearchUID must have less than 32 elements")
	}
	if strlen(r.orgUID) > 32 {
		return nil, reportError("orgUID must have less than 32 elements")
	}

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

type ApiDashboardSearchGetRequest struct {
	ctx context.Context
	ApiService *DashboardSearchApiService
	dashboardSearchUID string
	orgUID string
}

func (r ApiDashboardSearchGetRequest) Execute() (*DashboardSearch, *http.Response, error) {
	return r.ApiService.DashboardSearchGetExecute(r)
}

/*
DashboardSearchGet Get a dashboard search


Get a dashboard search in an org.

 * Requires action dashboard_search:Get

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param dashboardSearchUID
 @param orgUID
 @return ApiDashboardSearchGetRequest
*/
func (a *DashboardSearchApiService) DashboardSearchGet(ctx context.Context, dashboardSearchUID string, orgUID string) ApiDashboardSearchGetRequest {
	return ApiDashboardSearchGetRequest{
		ApiService: a,
		ctx: ctx,
		dashboardSearchUID: dashboardSearchUID,
		orgUID: orgUID,
	}
}

// Execute executes the request
//  @return DashboardSearch
func (a *DashboardSearchApiService) DashboardSearchGetExecute(r ApiDashboardSearchGetRequest) (*DashboardSearch, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DashboardSearch
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DashboardSearchApiService.DashboardSearchGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/org/{orgUID}/dashboard_search/{dashboardSearchUID}"
	localVarPath = strings.Replace(localVarPath, "{"+"dashboardSearchUID"+"}", url.PathEscape(parameterToString(r.dashboardSearchUID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"orgUID"+"}", url.PathEscape(parameterToString(r.orgUID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.dashboardSearchUID) > 32 {
		return localVarReturnValue, nil, reportError("dashboardSearchUID must have less than 32 elements")
	}
	if strlen(r.orgUID) > 32 {
		return localVarReturnValue, nil, reportError("orgUID must have less than 32 elements")
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

type ApiDashboardSearchListRequest struct {
	ctx context.Context
	ApiService *DashboardSearchApiService
	orgUID string
}

func (r ApiDashboardSearchListRequest) Execute() ([]DashboardSearch, *http.Response, error) {
	return r.ApiService.DashboardSearchListExecute(r)
}

/*
DashboardSearchList List dashboard searches


Lists dashboard searches by org.

 * Requires action dashboard_search:Query

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param orgUID Org UID
 @return ApiDashboardSearchListRequest
*/
func (a *DashboardSearchApiService) DashboardSearchList(ctx context.Context, orgUID string) ApiDashboardSearchListRequest {
	return ApiDashboardSearchListRequest{
		ApiService: a,
		ctx: ctx,
		orgUID: orgUID,
	}
}

// Execute executes the request
//  @return []DashboardSearch
func (a *DashboardSearchApiService) DashboardSearchListExecute(r ApiDashboardSearchListRequest) ([]DashboardSearch, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []DashboardSearch
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DashboardSearchApiService.DashboardSearchList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/org/{orgUID}/dashboard_search/"
	localVarPath = strings.Replace(localVarPath, "{"+"orgUID"+"}", url.PathEscape(parameterToString(r.orgUID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.orgUID) > 32 {
		return localVarReturnValue, nil, reportError("orgUID must have less than 32 elements")
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

type ApiDashboardSearchUpdateRequest struct {
	ctx context.Context
	ApiService *DashboardSearchApiService
	dashboardSearchUID string
	orgUID string
	dashboardSearchUpdateInput *DashboardSearchUpdateInput
}

func (r ApiDashboardSearchUpdateRequest) DashboardSearchUpdateInput(dashboardSearchUpdateInput DashboardSearchUpdateInput) ApiDashboardSearchUpdateRequest {
	r.dashboardSearchUpdateInput = &dashboardSearchUpdateInput
	return r
}

func (r ApiDashboardSearchUpdateRequest) Execute() (*http.Response, error) {
	return r.ApiService.DashboardSearchUpdateExecute(r)
}

/*
DashboardSearchUpdate Update a dashboard search


Update a dashboard search in an org.

 * Requires action dashboard_search:Update

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param dashboardSearchUID UID for the DashboardSearch
 @param orgUID Org UID
 @return ApiDashboardSearchUpdateRequest
*/
func (a *DashboardSearchApiService) DashboardSearchUpdate(ctx context.Context, dashboardSearchUID string, orgUID string) ApiDashboardSearchUpdateRequest {
	return ApiDashboardSearchUpdateRequest{
		ApiService: a,
		ctx: ctx,
		dashboardSearchUID: dashboardSearchUID,
		orgUID: orgUID,
	}
}

// Execute executes the request
func (a *DashboardSearchApiService) DashboardSearchUpdateExecute(r ApiDashboardSearchUpdateRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DashboardSearchApiService.DashboardSearchUpdate")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/org/{orgUID}/dashboard_search/{dashboardSearchUID}"
	localVarPath = strings.Replace(localVarPath, "{"+"dashboardSearchUID"+"}", url.PathEscape(parameterToString(r.dashboardSearchUID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"orgUID"+"}", url.PathEscape(parameterToString(r.orgUID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.orgUID) > 32 {
		return nil, reportError("orgUID must have less than 32 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.dashboardSearchUpdateInput
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
