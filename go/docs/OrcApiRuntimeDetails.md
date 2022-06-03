# OrcApiRuntimeDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentArch** | Pointer to **string** |  | [optional] 
**AgentVersion** | Pointer to **string** |  | [optional] 
**BootTime** | Pointer to **float64** |  | [optional] 
**CloudImageId** | Pointer to **string** |  | [optional] 
**CloudInstanceId** | Pointer to **string** |  | [optional] 
**CloudRegion** | Pointer to **string** |  | [optional] 
**CloudTags** | Pointer to **[]string** |  | [optional] 
**CloudType** | Pointer to **string** |  | [optional] 
**CpuCores** | Pointer to **int32** |  | [optional] 
**CpuMake** | Pointer to **string** |  | [optional] 
**CpuModel** | Pointer to **string** |  | [optional] 
**Hostname** | Pointer to **string** |  | [optional] 
**IpAddresses** | Pointer to **[]string** |  | [optional] 
**MacAddresses** | Pointer to **[]string** |  | [optional] 
**OsName** | Pointer to **string** |  | [optional] 
**OsPrettyName** | Pointer to **string** |  | [optional] 
**RequestIp** | Pointer to **string** |  | [optional] 
**SrcUid** | Pointer to **string** |  | [optional] 
**Uname** | Pointer to **string** |  | [optional] 

## Methods

### NewOrcApiRuntimeDetails

`func NewOrcApiRuntimeDetails() *OrcApiRuntimeDetails`

NewOrcApiRuntimeDetails instantiates a new OrcApiRuntimeDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrcApiRuntimeDetailsWithDefaults

`func NewOrcApiRuntimeDetailsWithDefaults() *OrcApiRuntimeDetails`

NewOrcApiRuntimeDetailsWithDefaults instantiates a new OrcApiRuntimeDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentArch

`func (o *OrcApiRuntimeDetails) GetAgentArch() string`

GetAgentArch returns the AgentArch field if non-nil, zero value otherwise.

### GetAgentArchOk

`func (o *OrcApiRuntimeDetails) GetAgentArchOk() (*string, bool)`

GetAgentArchOk returns a tuple with the AgentArch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentArch

`func (o *OrcApiRuntimeDetails) SetAgentArch(v string)`

SetAgentArch sets AgentArch field to given value.

### HasAgentArch

`func (o *OrcApiRuntimeDetails) HasAgentArch() bool`

HasAgentArch returns a boolean if a field has been set.

### GetAgentVersion

`func (o *OrcApiRuntimeDetails) GetAgentVersion() string`

GetAgentVersion returns the AgentVersion field if non-nil, zero value otherwise.

### GetAgentVersionOk

`func (o *OrcApiRuntimeDetails) GetAgentVersionOk() (*string, bool)`

GetAgentVersionOk returns a tuple with the AgentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentVersion

`func (o *OrcApiRuntimeDetails) SetAgentVersion(v string)`

SetAgentVersion sets AgentVersion field to given value.

### HasAgentVersion

`func (o *OrcApiRuntimeDetails) HasAgentVersion() bool`

HasAgentVersion returns a boolean if a field has been set.

### GetBootTime

`func (o *OrcApiRuntimeDetails) GetBootTime() float64`

GetBootTime returns the BootTime field if non-nil, zero value otherwise.

### GetBootTimeOk

`func (o *OrcApiRuntimeDetails) GetBootTimeOk() (*float64, bool)`

GetBootTimeOk returns a tuple with the BootTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootTime

`func (o *OrcApiRuntimeDetails) SetBootTime(v float64)`

SetBootTime sets BootTime field to given value.

### HasBootTime

`func (o *OrcApiRuntimeDetails) HasBootTime() bool`

HasBootTime returns a boolean if a field has been set.

### GetCloudImageId

`func (o *OrcApiRuntimeDetails) GetCloudImageId() string`

GetCloudImageId returns the CloudImageId field if non-nil, zero value otherwise.

### GetCloudImageIdOk

`func (o *OrcApiRuntimeDetails) GetCloudImageIdOk() (*string, bool)`

GetCloudImageIdOk returns a tuple with the CloudImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudImageId

`func (o *OrcApiRuntimeDetails) SetCloudImageId(v string)`

SetCloudImageId sets CloudImageId field to given value.

### HasCloudImageId

`func (o *OrcApiRuntimeDetails) HasCloudImageId() bool`

HasCloudImageId returns a boolean if a field has been set.

### GetCloudInstanceId

`func (o *OrcApiRuntimeDetails) GetCloudInstanceId() string`

GetCloudInstanceId returns the CloudInstanceId field if non-nil, zero value otherwise.

### GetCloudInstanceIdOk

`func (o *OrcApiRuntimeDetails) GetCloudInstanceIdOk() (*string, bool)`

GetCloudInstanceIdOk returns a tuple with the CloudInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInstanceId

`func (o *OrcApiRuntimeDetails) SetCloudInstanceId(v string)`

SetCloudInstanceId sets CloudInstanceId field to given value.

### HasCloudInstanceId

`func (o *OrcApiRuntimeDetails) HasCloudInstanceId() bool`

HasCloudInstanceId returns a boolean if a field has been set.

### GetCloudRegion

`func (o *OrcApiRuntimeDetails) GetCloudRegion() string`

GetCloudRegion returns the CloudRegion field if non-nil, zero value otherwise.

### GetCloudRegionOk

`func (o *OrcApiRuntimeDetails) GetCloudRegionOk() (*string, bool)`

GetCloudRegionOk returns a tuple with the CloudRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudRegion

`func (o *OrcApiRuntimeDetails) SetCloudRegion(v string)`

SetCloudRegion sets CloudRegion field to given value.

### HasCloudRegion

`func (o *OrcApiRuntimeDetails) HasCloudRegion() bool`

HasCloudRegion returns a boolean if a field has been set.

### GetCloudTags

`func (o *OrcApiRuntimeDetails) GetCloudTags() []string`

GetCloudTags returns the CloudTags field if non-nil, zero value otherwise.

### GetCloudTagsOk

`func (o *OrcApiRuntimeDetails) GetCloudTagsOk() (*[]string, bool)`

GetCloudTagsOk returns a tuple with the CloudTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudTags

`func (o *OrcApiRuntimeDetails) SetCloudTags(v []string)`

SetCloudTags sets CloudTags field to given value.

### HasCloudTags

`func (o *OrcApiRuntimeDetails) HasCloudTags() bool`

HasCloudTags returns a boolean if a field has been set.

### GetCloudType

`func (o *OrcApiRuntimeDetails) GetCloudType() string`

GetCloudType returns the CloudType field if non-nil, zero value otherwise.

### GetCloudTypeOk

`func (o *OrcApiRuntimeDetails) GetCloudTypeOk() (*string, bool)`

GetCloudTypeOk returns a tuple with the CloudType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudType

`func (o *OrcApiRuntimeDetails) SetCloudType(v string)`

SetCloudType sets CloudType field to given value.

### HasCloudType

`func (o *OrcApiRuntimeDetails) HasCloudType() bool`

HasCloudType returns a boolean if a field has been set.

### GetCpuCores

`func (o *OrcApiRuntimeDetails) GetCpuCores() int32`

GetCpuCores returns the CpuCores field if non-nil, zero value otherwise.

### GetCpuCoresOk

`func (o *OrcApiRuntimeDetails) GetCpuCoresOk() (*int32, bool)`

GetCpuCoresOk returns a tuple with the CpuCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCores

`func (o *OrcApiRuntimeDetails) SetCpuCores(v int32)`

SetCpuCores sets CpuCores field to given value.

### HasCpuCores

`func (o *OrcApiRuntimeDetails) HasCpuCores() bool`

HasCpuCores returns a boolean if a field has been set.

### GetCpuMake

`func (o *OrcApiRuntimeDetails) GetCpuMake() string`

GetCpuMake returns the CpuMake field if non-nil, zero value otherwise.

### GetCpuMakeOk

`func (o *OrcApiRuntimeDetails) GetCpuMakeOk() (*string, bool)`

GetCpuMakeOk returns a tuple with the CpuMake field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuMake

`func (o *OrcApiRuntimeDetails) SetCpuMake(v string)`

SetCpuMake sets CpuMake field to given value.

### HasCpuMake

`func (o *OrcApiRuntimeDetails) HasCpuMake() bool`

HasCpuMake returns a boolean if a field has been set.

### GetCpuModel

`func (o *OrcApiRuntimeDetails) GetCpuModel() string`

GetCpuModel returns the CpuModel field if non-nil, zero value otherwise.

### GetCpuModelOk

`func (o *OrcApiRuntimeDetails) GetCpuModelOk() (*string, bool)`

GetCpuModelOk returns a tuple with the CpuModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuModel

`func (o *OrcApiRuntimeDetails) SetCpuModel(v string)`

SetCpuModel sets CpuModel field to given value.

### HasCpuModel

`func (o *OrcApiRuntimeDetails) HasCpuModel() bool`

HasCpuModel returns a boolean if a field has been set.

### GetHostname

`func (o *OrcApiRuntimeDetails) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *OrcApiRuntimeDetails) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *OrcApiRuntimeDetails) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *OrcApiRuntimeDetails) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetIpAddresses

`func (o *OrcApiRuntimeDetails) GetIpAddresses() []string`

GetIpAddresses returns the IpAddresses field if non-nil, zero value otherwise.

### GetIpAddressesOk

`func (o *OrcApiRuntimeDetails) GetIpAddressesOk() (*[]string, bool)`

GetIpAddressesOk returns a tuple with the IpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddresses

`func (o *OrcApiRuntimeDetails) SetIpAddresses(v []string)`

SetIpAddresses sets IpAddresses field to given value.

### HasIpAddresses

`func (o *OrcApiRuntimeDetails) HasIpAddresses() bool`

HasIpAddresses returns a boolean if a field has been set.

### GetMacAddresses

`func (o *OrcApiRuntimeDetails) GetMacAddresses() []string`

GetMacAddresses returns the MacAddresses field if non-nil, zero value otherwise.

### GetMacAddressesOk

`func (o *OrcApiRuntimeDetails) GetMacAddressesOk() (*[]string, bool)`

GetMacAddressesOk returns a tuple with the MacAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddresses

`func (o *OrcApiRuntimeDetails) SetMacAddresses(v []string)`

SetMacAddresses sets MacAddresses field to given value.

### HasMacAddresses

`func (o *OrcApiRuntimeDetails) HasMacAddresses() bool`

HasMacAddresses returns a boolean if a field has been set.

### GetOsName

`func (o *OrcApiRuntimeDetails) GetOsName() string`

GetOsName returns the OsName field if non-nil, zero value otherwise.

### GetOsNameOk

`func (o *OrcApiRuntimeDetails) GetOsNameOk() (*string, bool)`

GetOsNameOk returns a tuple with the OsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsName

`func (o *OrcApiRuntimeDetails) SetOsName(v string)`

SetOsName sets OsName field to given value.

### HasOsName

`func (o *OrcApiRuntimeDetails) HasOsName() bool`

HasOsName returns a boolean if a field has been set.

### GetOsPrettyName

`func (o *OrcApiRuntimeDetails) GetOsPrettyName() string`

GetOsPrettyName returns the OsPrettyName field if non-nil, zero value otherwise.

### GetOsPrettyNameOk

`func (o *OrcApiRuntimeDetails) GetOsPrettyNameOk() (*string, bool)`

GetOsPrettyNameOk returns a tuple with the OsPrettyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsPrettyName

`func (o *OrcApiRuntimeDetails) SetOsPrettyName(v string)`

SetOsPrettyName sets OsPrettyName field to given value.

### HasOsPrettyName

`func (o *OrcApiRuntimeDetails) HasOsPrettyName() bool`

HasOsPrettyName returns a boolean if a field has been set.

### GetRequestIp

`func (o *OrcApiRuntimeDetails) GetRequestIp() string`

GetRequestIp returns the RequestIp field if non-nil, zero value otherwise.

### GetRequestIpOk

`func (o *OrcApiRuntimeDetails) GetRequestIpOk() (*string, bool)`

GetRequestIpOk returns a tuple with the RequestIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestIp

`func (o *OrcApiRuntimeDetails) SetRequestIp(v string)`

SetRequestIp sets RequestIp field to given value.

### HasRequestIp

`func (o *OrcApiRuntimeDetails) HasRequestIp() bool`

HasRequestIp returns a boolean if a field has been set.

### GetSrcUid

`func (o *OrcApiRuntimeDetails) GetSrcUid() string`

GetSrcUid returns the SrcUid field if non-nil, zero value otherwise.

### GetSrcUidOk

`func (o *OrcApiRuntimeDetails) GetSrcUidOk() (*string, bool)`

GetSrcUidOk returns a tuple with the SrcUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcUid

`func (o *OrcApiRuntimeDetails) SetSrcUid(v string)`

SetSrcUid sets SrcUid field to given value.

### HasSrcUid

`func (o *OrcApiRuntimeDetails) HasSrcUid() bool`

HasSrcUid returns a boolean if a field has been set.

### GetUname

`func (o *OrcApiRuntimeDetails) GetUname() string`

GetUname returns the Uname field if non-nil, zero value otherwise.

### GetUnameOk

`func (o *OrcApiRuntimeDetails) GetUnameOk() (*string, bool)`

GetUnameOk returns a tuple with the Uname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUname

`func (o *OrcApiRuntimeDetails) SetUname(v string)`

SetUname sets Uname field to given value.

### HasUname

`func (o *OrcApiRuntimeDetails) HasUname() bool`

HasUname returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


