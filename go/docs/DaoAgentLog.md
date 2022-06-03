# DaoAgentLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentRegistrationUid** | Pointer to **string** |  | [optional] 
**AgentUid** | Pointer to **string** |  | [optional] 
**BytesSent** | Pointer to **int32** |  | [optional] 
**ErrUid** | Pointer to **string** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**IpAddress** | Pointer to **string** |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] 
**OrgUid** | Pointer to **string** |  | [optional] 
**RuntimeDetails** | Pointer to [**OrcApiRuntimeDetails**](OrcApiRuntimeDetails.md) |  | [optional] 
**SourceUid** | Pointer to **string** |  | [optional] 
**Time** | Pointer to **float64** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewDaoAgentLog

`func NewDaoAgentLog() *DaoAgentLog`

NewDaoAgentLog instantiates a new DaoAgentLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDaoAgentLogWithDefaults

`func NewDaoAgentLogWithDefaults() *DaoAgentLog`

NewDaoAgentLogWithDefaults instantiates a new DaoAgentLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentRegistrationUid

`func (o *DaoAgentLog) GetAgentRegistrationUid() string`

GetAgentRegistrationUid returns the AgentRegistrationUid field if non-nil, zero value otherwise.

### GetAgentRegistrationUidOk

`func (o *DaoAgentLog) GetAgentRegistrationUidOk() (*string, bool)`

GetAgentRegistrationUidOk returns a tuple with the AgentRegistrationUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentRegistrationUid

`func (o *DaoAgentLog) SetAgentRegistrationUid(v string)`

SetAgentRegistrationUid sets AgentRegistrationUid field to given value.

### HasAgentRegistrationUid

`func (o *DaoAgentLog) HasAgentRegistrationUid() bool`

HasAgentRegistrationUid returns a boolean if a field has been set.

### GetAgentUid

`func (o *DaoAgentLog) GetAgentUid() string`

GetAgentUid returns the AgentUid field if non-nil, zero value otherwise.

### GetAgentUidOk

`func (o *DaoAgentLog) GetAgentUidOk() (*string, bool)`

GetAgentUidOk returns a tuple with the AgentUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentUid

`func (o *DaoAgentLog) SetAgentUid(v string)`

SetAgentUid sets AgentUid field to given value.

### HasAgentUid

`func (o *DaoAgentLog) HasAgentUid() bool`

HasAgentUid returns a boolean if a field has been set.

### GetBytesSent

`func (o *DaoAgentLog) GetBytesSent() int32`

GetBytesSent returns the BytesSent field if non-nil, zero value otherwise.

### GetBytesSentOk

`func (o *DaoAgentLog) GetBytesSentOk() (*int32, bool)`

GetBytesSentOk returns a tuple with the BytesSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesSent

`func (o *DaoAgentLog) SetBytesSent(v int32)`

SetBytesSent sets BytesSent field to given value.

### HasBytesSent

`func (o *DaoAgentLog) HasBytesSent() bool`

HasBytesSent returns a boolean if a field has been set.

### GetErrUid

`func (o *DaoAgentLog) GetErrUid() string`

GetErrUid returns the ErrUid field if non-nil, zero value otherwise.

### GetErrUidOk

`func (o *DaoAgentLog) GetErrUidOk() (*string, bool)`

GetErrUidOk returns a tuple with the ErrUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrUid

`func (o *DaoAgentLog) SetErrUid(v string)`

SetErrUid sets ErrUid field to given value.

### HasErrUid

`func (o *DaoAgentLog) HasErrUid() bool`

HasErrUid returns a boolean if a field has been set.

### GetError

`func (o *DaoAgentLog) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *DaoAgentLog) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *DaoAgentLog) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *DaoAgentLog) HasError() bool`

HasError returns a boolean if a field has been set.

### GetIpAddress

`func (o *DaoAgentLog) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *DaoAgentLog) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *DaoAgentLog) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *DaoAgentLog) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetMsg

`func (o *DaoAgentLog) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *DaoAgentLog) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *DaoAgentLog) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *DaoAgentLog) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetOrgUid

`func (o *DaoAgentLog) GetOrgUid() string`

GetOrgUid returns the OrgUid field if non-nil, zero value otherwise.

### GetOrgUidOk

`func (o *DaoAgentLog) GetOrgUidOk() (*string, bool)`

GetOrgUidOk returns a tuple with the OrgUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgUid

`func (o *DaoAgentLog) SetOrgUid(v string)`

SetOrgUid sets OrgUid field to given value.

### HasOrgUid

`func (o *DaoAgentLog) HasOrgUid() bool`

HasOrgUid returns a boolean if a field has been set.

### GetRuntimeDetails

`func (o *DaoAgentLog) GetRuntimeDetails() OrcApiRuntimeDetails`

GetRuntimeDetails returns the RuntimeDetails field if non-nil, zero value otherwise.

### GetRuntimeDetailsOk

`func (o *DaoAgentLog) GetRuntimeDetailsOk() (*OrcApiRuntimeDetails, bool)`

GetRuntimeDetailsOk returns a tuple with the RuntimeDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeDetails

`func (o *DaoAgentLog) SetRuntimeDetails(v OrcApiRuntimeDetails)`

SetRuntimeDetails sets RuntimeDetails field to given value.

### HasRuntimeDetails

`func (o *DaoAgentLog) HasRuntimeDetails() bool`

HasRuntimeDetails returns a boolean if a field has been set.

### GetSourceUid

`func (o *DaoAgentLog) GetSourceUid() string`

GetSourceUid returns the SourceUid field if non-nil, zero value otherwise.

### GetSourceUidOk

`func (o *DaoAgentLog) GetSourceUidOk() (*string, bool)`

GetSourceUidOk returns a tuple with the SourceUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUid

`func (o *DaoAgentLog) SetSourceUid(v string)`

SetSourceUid sets SourceUid field to given value.

### HasSourceUid

`func (o *DaoAgentLog) HasSourceUid() bool`

HasSourceUid returns a boolean if a field has been set.

### GetTime

`func (o *DaoAgentLog) GetTime() float64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *DaoAgentLog) GetTimeOk() (*float64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *DaoAgentLog) SetTime(v float64)`

SetTime sets Time field to given value.

### HasTime

`func (o *DaoAgentLog) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetType

`func (o *DaoAgentLog) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DaoAgentLog) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DaoAgentLog) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DaoAgentLog) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


