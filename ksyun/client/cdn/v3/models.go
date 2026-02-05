package v3

import (
	"encoding/json"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)

type SetCacheRuleConfigCacheRules struct {
	CacheRuleType *string `json:"CacheRuleType,omitempty" name:"CacheRuleType"`
	Value         *string `json:"Value,omitempty" name:"Value"`
	CacheEnable   *string `json:"CacheEnable,omitempty" name:"CacheEnable"`
	CacheTime     *int64  `json:"CacheTime,omitempty" name:"CacheTime"`
	RespectOrigin *string `json:"RespectOrigin,omitempty" name:"RespectOrigin"`
}
type SetErrorPageConfigErrorPages struct {
	ErrorHttpCode *string `json:"ErrorHttpCode,omitempty" name:"ErrorHttpCode"`
	CustomPageUrl *string `json:"CustomPageUrl,omitempty" name:"CustomPageUrl"`
}
type SetCdnBlockDomainUrlUrls struct {
	Url *string `json:"Url,omitempty" name:"Url"`
}
type SubmitRefreshCachesFiles struct {
	Url *string `json:"Url,omitempty" name:"Url"`
}
type SubmitRefreshCachesDirs struct {
	Url *string `json:"Url,omitempty" name:"Url"`
}
type SubmitPreloadCachesUrls struct {
	Url *string `json:"Url,omitempty" name:"Url"`
}
type GetRefreshOrPreloadTaskUrls struct {
	Url *string `json:"Url,omitempty" name:"Url"`
}

type GetDomainLogsRequest struct {
	*ksyunhttp.BaseRequest
	DomainId   *string `json:"DomainId,omitempty" name:"DomainId"`
	StartTime  *string `json:"StartTime,omitempty" name:"StartTime"`
	EndTime    *string `json:"EndTime,omitempty" name:"EndTime"`
	PageSize   *int64  `json:"PageSize,omitempty" name:"PageSize"`
	PageNumber *int64  `json:"PageNumber,omitempty" name:"PageNumber"`
}

func (r *GetDomainLogsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetDomainLogsResponse struct {
	*ksyunhttp.BaseResponse
	DomainId   *string `json:"DomainId" name:"DomainId"`
	PageSize   *int64  `json:"PageSize" name:"PageSize"`
	PageNumber *int64  `json:"PageNumber" name:"PageNumber"`
	TotalCount *int64  `json:"TotalCount" name:"TotalCount"`
	DomainLogs []struct {
		StartTime *string `json:"StartTime" name:"StartTime"`
		EndTime   *string `json:"EndTime" name:"EndTime"`
		LogName   *string `json:"LogName" name:"LogName"`
		LogUrl    *string `json:"LogUrl" name:"LogUrl"`
		LogSize   *int64  `json:"LogSize" name:"LogSize"`
	} `json:"DomainLogs"`
}

func (r *GetDomainLogsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDomainLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetClientRequestDataRequest struct {
	*ksyunhttp.BaseRequest
	StartTime  *string `json:"StartTime,omitempty" name:"StartTime"`
	EndTime    *string `json:"EndTime,omitempty" name:"EndTime"`
	Metric     *string `json:"Metric,omitempty" name:"Metric"`
	Interval   *string `json:"Interval,omitempty" name:"Interval"`
	CdnType    *string `json:"CdnType,omitempty" name:"CdnType"`
	Domains    *string `json:"Domains,omitempty" name:"Domains"`
	Areas      *string `json:"Areas,omitempty" name:"Areas"`
	Provinces  *string `json:"Provinces,omitempty" name:"Provinces"`
	Isps       *string `json:"Isps,omitempty" name:"Isps"`
	IpType     *string `json:"IpType,omitempty" name:"IpType"`
	Schema     *string `json:"Schema,omitempty" name:"Schema"`
	ResultType *string `json:"ResultType,omitempty" name:"ResultType"`
}

func (r *GetClientRequestDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetClientRequestDataResponse struct {
	*ksyunhttp.BaseResponse
	StartTime  *string `json:"StartTime" name:"StartTime"`
	EndTime    *string `json:"EndTime" name:"EndTime"`
	Metric     *string `json:"Metric" name:"Metric"`
	Interval   *int    `json:"Interval" name:"Interval"`
	CdnType    *string `json:"CdnType" name:"CdnType"`
	Domains    *string `json:"Domains" name:"Domains"`
	Areas      *string `json:"Areas" name:"Areas"`
	Provinces  *string `json:"Provinces" name:"Provinces"`
	Isps       *string `json:"Isps" name:"Isps"`
	IpType     *string `json:"IpType" name:"IpType"`
	Schema     *string `json:"Schema" name:"Schema"`
	ResultType *string `json:"ResultType" name:"ResultType"`
	Datas      []struct {
		Condition struct {
		} `json:"Condition" name:"Condition"`
		Data []struct {
			Time *string  `json:"Time" name:"Time"`
			Flow *float64 `json:"Flow" name:"Flow"`
		} `json:"Data" name:"Data"`
	} `json:"Datas"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *GetClientRequestDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetClientRequestDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetCdnDomainsRequest struct {
	*ksyunhttp.BaseRequest
	PageSize     *int64  `json:"PageSize,omitempty" name:"PageSize"`
	PageNumber   *int64  `json:"PageNumber,omitempty" name:"PageNumber"`
	DomainName   *string `json:"DomainName,omitempty" name:"DomainName"`
	ProjectId    *int64  `json:"ProjectId,omitempty" name:"ProjectId"`
	DomainStatus *string `json:"DomainStatus,omitempty" name:"DomainStatus"`
	CdnType      *string `json:"CdnType,omitempty" name:"CdnType"`
	FuzzyMatch   *string `json:"FuzzyMatch,omitempty" name:"FuzzyMatch"`
}

func (r *GetCdnDomainsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetCdnDomainsResponse struct {
	*ksyunhttp.BaseResponse
	PageNumber *int64 `json:"PageNumber" name:"PageNumber"`
	PageSize   *int64 `json:"PageSize" name:"PageSize"`
	TotalCount *int64 `json:"TotalCount" name:"TotalCount"`
	Domains    []struct {
		DomainName      *string `json:"DomainName" name:"DomainName"`
		DomainId        *string `json:"DomainId" name:"DomainId"`
		Cname           *string `json:"Cname" name:"Cname"`
		CdnType         *string `json:"CdnType" name:"CdnType"`
		CdnSubType      *string `json:"CdnSubType" name:"CdnSubType"`
		IcpRegistration *string `json:"IcpRegistration" name:"IcpRegistration"`
		DomainStatus    *string `json:"DomainStatus" name:"DomainStatus"`
		CreatedTime     *string `json:"CreatedTime" name:"CreatedTime"`
		ModifiedTime    *string `json:"ModifiedTime" name:"ModifiedTime"`
		Description     *string `json:"Description" name:"Description"`
		Projectld       *int    `json:"Projectld" name:"Projectld"`
		ProjectName     *string `json:"ProjectName" name:"ProjectName"`
		Region          *string `json:"Region" name:"Region"`
	} `json:"Domains"`
}

func (r *GetCdnDomainsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCdnDomainsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCdnDomainRequest struct {
	*ksyunhttp.BaseRequest
	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`
}

func (r *DeleteCdnDomainRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteCdnDomainResponse struct {
	*ksyunhttp.BaseResponse
	DeleteCdnDomainResponse *string `json:"DeleteCdnDomainResponse" name:"DeleteCdnDomainResponse"`
}

func (r *DeleteCdnDomainResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCdnDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetCdnDomainBasicInfoRequest struct {
	*ksyunhttp.BaseRequest
	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`
}

func (r *GetCdnDomainBasicInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetCdnDomainBasicInfoResponse struct {
	*ksyunhttp.BaseResponse
	DomainName      *string `json:"DomainName" name:"DomainName"`
	DomainId        *string `json:"DomainId" name:"DomainId"`
	Cname           *string `json:"Cname" name:"Cname"`
	DomainStatus    *string `json:"DomainStatus" name:"DomainStatus"`
	CdnType         *string `json:"CdnType" name:"CdnType"`
	CdnSubType      *string `json:"CdnSubType" name:"CdnSubType"`
	ProjectId       *int    `json:"ProjectId" name:"ProjectId"`
	ProjectName     *string `json:"ProjectName" name:"ProjectName"`
	IcpRegistration *string `json:"IcpRegistration" name:"IcpRegistration"`
	AuditFailReason *string `json:"AuditFailReason" name:"AuditFailReason"`
	CdnProtocol     *string `json:"CdnProtocol" name:"CdnProtocol"`
	Regions         *string `json:"Regions" name:"Regions"`
	OriginType      *string `json:"OriginType" name:"OriginType"`
	OriginProtocol  *string `json:"OriginProtocol" name:"OriginProtocol"`
	Origin          *string `json:"Origin" name:"Origin"`
	OriginHttpPort  *int    `json:"OriginHttpPort" name:"OriginHttpPort"`
	OriginHttpsPort *int    `json:"OriginHttpsPort" name:"OriginHttpsPort"`
	CreatedTime     *string `json:"CreatedTime" name:"CreatedTime"`
	ModifiedTime    *string `json:"ModifiedTime" name:"ModifiedTime"`
}

func (r *GetCdnDomainBasicInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetCdnDomainBasicInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCdnDomainBasicInfoRequest struct {
	*ksyunhttp.BaseRequest
	DomainId       *string `json:"DomainId,omitempty" name:"DomainId"`
	Regions        *string `json:"Regions,omitempty" name:"Regions"`
	OriginType     *string `json:"OriginType,omitempty" name:"OriginType"`
	OriginProtocol *string `json:"OriginProtocol,omitempty" name:"OriginProtocol"`
	Origin         *string `json:"Origin,omitempty" name:"Origin"`
}

func (r *ModifyCdnDomainBasicInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifyCdnDomainBasicInfoResponse struct {
	*ksyunhttp.BaseResponse
	ModifyCdnDomainBasicInfoResponse *string `json:"ModifyCdnDomainBasicInfoResponse" name:"ModifyCdnDomainBasicInfoResponse"`
}

func (r *ModifyCdnDomainBasicInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCdnDomainBasicInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddCdnDomainRequest struct {
	*ksyunhttp.BaseRequest
	DomainName     *string `json:"DomainName,omitempty" name:"DomainName"`
	CdnType        *string `json:"CdnType,omitempty" name:"CdnType"`
	ProjectId      *string `json:"ProjectId,omitempty" name:"ProjectId"`
	CdnProtocol    *string `json:"CdnProtocol,omitempty" name:"CdnProtocol"`
	Regions        *string `json:"Regions,omitempty" name:"Regions"`
	OriginType     *string `json:"OriginType,omitempty" name:"OriginType"`
	OriginProtocol *string `json:"OriginProtocol,omitempty" name:"OriginProtocol"`
	Origin         *string `json:"Origin,omitempty" name:"Origin"`
}

func (r *AddCdnDomainRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type AddCdnDomainResponse struct {
	*ksyunhttp.BaseResponse
	DomainId     *string `json:"DomainId" name:"DomainId"`
	DomainStatus *string `json:"DomainStatus" name:"DomainStatus"`
}

func (r *AddCdnDomainResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddCdnDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetDomainConfigsRequest struct {
	*ksyunhttp.BaseRequest
	DomainId   *string `json:"DomainId,omitempty" name:"DomainId"`
	ConfigList *string `json:"ConfigList,omitempty" name:"ConfigList"`
}

func (r *GetDomainConfigsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetDomainConfigsResponse struct {
	*ksyunhttp.BaseResponse
	DomainConfigs *string `json:"DomainConfigs" name:"DomainConfigs"`
}

func (r *GetDomainConfigsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDomainConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StartStopCdnDomainRequest struct {
	*ksyunhttp.BaseRequest
	DomainId   *string `json:"DomainId,omitempty" name:"DomainId"`
	ActionType *string `json:"ActionType,omitempty" name:"ActionType"`
}

func (r *StartStopCdnDomainRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type StartStopCdnDomainResponse struct {
	*ksyunhttp.BaseResponse
	StartStopCdnDomainResponse *string `json:"StartStopCdnDomainResponse" name:"StartStopCdnDomainResponse"`
}

func (r *StartStopCdnDomainResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StartStopCdnDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetCacheRuleConfigRequest struct {
	*ksyunhttp.BaseRequest
	DomainId   *string                         `json:"DomainId,omitempty" name:"DomainId"`
	CacheRules []*SetCacheRuleConfigCacheRules `json:"CacheRules,omitempty" name:"CacheRules"`
}

func (r *SetCacheRuleConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type SetCacheRuleConfigResponse struct {
	*ksyunhttp.BaseResponse
	SetCacheRuleConfigResponse *string `json:"SetCacheRuleConfigResponse" name:"SetCacheRuleConfigResponse"`
}

func (r *SetCacheRuleConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetCacheRuleConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetBackOriginHostConfigRequest struct {
	*ksyunhttp.BaseRequest
	DomainId        *string `json:"DomainId,omitempty" name:"DomainId"`
	FollowReqDomain *string `json:"FollowReqDomain,omitempty" name:"FollowReqDomain"`
	BackOriginHost  *string `json:"BackOriginHost,omitempty" name:"BackOriginHost"`
}

func (r *SetBackOriginHostConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type SetBackOriginHostConfigResponse struct {
	*ksyunhttp.BaseResponse
	SetBackOriginHostConfigResponse *string `json:"SetBackOriginHostConfigResponse" name:"SetBackOriginHostConfigResponse"`
}

func (r *SetBackOriginHostConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetBackOriginHostConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetValidDomainListRequest struct {
	*ksyunhttp.BaseRequest
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	EndTime   *string `json:"EndTime,omitempty" name:"EndTime"`
	CdnType   *string `json:"CdnType,omitempty" name:"CdnType"`
}

func (r *GetValidDomainListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetValidDomainListResponse struct {
	*ksyunhttp.BaseResponse
	Domainlist []struct {
		Domain    *string `json:"Domain" name:"Domain"`
		DomainId  *string `json:"DomainId" name:"DomainId"`
		ProjectId *string `json:"ProjectId" name:"ProjectId"`
		CdnType   *string `json:"CdnType" name:"CdnType"`
	} `json:"Domainlist"`
}

func (r *GetValidDomainListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetValidDomainListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetDomainAuthContentRequest struct {
	*ksyunhttp.BaseRequest
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`
}

func (r *GetDomainAuthContentRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetDomainAuthContentResponse struct {
	*ksyunhttp.BaseResponse
	Content   *string `json:"Content" name:"Content"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *GetDomainAuthContentResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDomainAuthContentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetVideoSeekConfigRequest struct {
	*ksyunhttp.BaseRequest
	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`
	Enable   *string `json:"Enable,omitempty" name:"Enable"`
}

func (r *SetVideoSeekConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type SetVideoSeekConfigResponse struct {
	*ksyunhttp.BaseResponse
	SetVideoSeekConfigResponse *string `json:"SetVideoSeekConfigResponse" name:"SetVideoSeekConfigResponse"`
}

func (r *SetVideoSeekConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetVideoSeekConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetPageCompressConfigRequest struct {
	*ksyunhttp.BaseRequest
	DomainId     *string `json:"DomainId,omitempty" name:"DomainId"`
	Enable       *string `json:"Enable,omitempty" name:"Enable"`
	CompressType *string `json:"CompressType,omitempty" name:"CompressType"`
	FileType     *string `json:"FileType,omitempty" name:"FileType"`
	FileSize     *string `json:"FileSize,omitempty" name:"FileSize"`
}

func (r *SetPageCompressConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type SetPageCompressConfigResponse struct {
	*ksyunhttp.BaseResponse
	SetPageCompressConfigResponse *string `json:"SetPageCompressConfigResponse" name:"SetPageCompressConfigResponse"`
}

func (r *SetPageCompressConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetPageCompressConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetBrCompressConfigRequest struct {
	*ksyunhttp.BaseRequest
	DomainId     *string `json:"DomainId,omitempty" name:"DomainId"`
	Enable       *string `json:"Enable,omitempty" name:"Enable"`
	CompressType *string `json:"CompressType,omitempty" name:"CompressType"`
	FileType     *string `json:"FileType,omitempty" name:"FileType"`
	FileSize     *string `json:"FileSize,omitempty" name:"FileSize"`
}

func (r *SetBrCompressConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type SetBrCompressConfigResponse struct {
	*ksyunhttp.BaseResponse
	SetBrCompressConfigResponse *string `json:"SetBrCompressConfigResponse" name:"SetBrCompressConfigResponse"`
}

func (r *SetBrCompressConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetBrCompressConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetIgnoreQueryStringConfigRequest struct {
	*ksyunhttp.BaseRequest
	DomainId    *string `json:"DomainId,omitempty" name:"DomainId"`
	Enable      *string `json:"Enable,omitempty" name:"Enable"`
	HashKeyArgs *string `json:"HashKeyArgs,omitempty" name:"HashKeyArgs"`
	Type        *string `json:"Type,omitempty" name:"Type"`
}

func (r *SetIgnoreQueryStringConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type SetIgnoreQueryStringConfigResponse struct {
	*ksyunhttp.BaseResponse
	SetIgnoreQueryStringConfigResponse *string `json:"SetIgnoreQueryStringConfigResponse" name:"SetIgnoreQueryStringConfigResponse"`
}

func (r *SetIgnoreQueryStringConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetIgnoreQueryStringConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetSetOriginAdvancedConfigRequest struct {
	*ksyunhttp.BaseRequest
	DomainId              *string `json:"DomainId,omitempty" name:"DomainId"`
	Enable                *string `json:"Enable,omitempty" name:"Enable"`
	OriginType            *string `json:"OriginType,omitempty" name:"OriginType"`
	Origin                *string `json:"Origin,omitempty" name:"Origin"`
	BackupOriginType      *string `json:"BackupOriginType,omitempty" name:"BackupOriginType"`
	BackupOrigin          *string `json:"BackupOrigin,omitempty" name:"BackupOrigin"`
	OriginPolicy          *string `json:"OriginPolicy,omitempty" name:"OriginPolicy"`
	OriginPolicyBestCount *string `json:"OriginPolicyBestCount,omitempty" name:"OriginPolicyBestCount"`
}

func (r *SetSetOriginAdvancedConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type SetSetOriginAdvancedConfigResponse struct {
	*ksyunhttp.BaseResponse
	SetOriginAdvancedConfigResponse *string `json:"SetOriginAdvancedConfigResponse" name:"SetOriginAdvancedConfigResponse"`
}

func (r *SetSetOriginAdvancedConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetSetOriginAdvancedConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ValidateDomainOwnerRequest struct {
	*ksyunhttp.BaseRequest
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`
	AuthType   *string `json:"AuthType,omitempty" name:"AuthType"`
}

func (r *ValidateDomainOwnerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ValidateDomainOwnerResponse struct {
	*ksyunhttp.BaseResponse
	Content   *string `json:"Content" name:"Content"`
	Result    *string `json:"Result" name:"Result"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ValidateDomainOwnerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ValidateDomainOwnerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetHttp2OptionConfigRequest struct {
	*ksyunhttp.BaseRequest
	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`
	Enable   *string `json:"Enable,omitempty" name:"Enable"`
}

func (r *SetHttp2OptionConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type SetHttp2OptionConfigResponse struct {
	*ksyunhttp.BaseResponse
	SetHttp2OptionConfigResponse *string `json:"SetHttp2OptionConfigResponse" name:"SetHttp2OptionConfigResponse"`
}

func (r *SetHttp2OptionConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetHttp2OptionConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetReferProtectionConfigRequest struct {
	*ksyunhttp.BaseRequest
	DomainId   *string `json:"DomainId,omitempty" name:"DomainId"`
	Enable     *string `json:"Enable,omitempty" name:"Enable"`
	ReferType  *string `json:"ReferType,omitempty" name:"ReferType"`
	ReferList  *string `json:"ReferList,omitempty" name:"ReferList"`
	AllowEmpty *string `json:"AllowEmpty,omitempty" name:"AllowEmpty"`
}

func (r *SetReferProtectionConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type SetReferProtectionConfigResponse struct {
	*ksyunhttp.BaseResponse
	SetReferProtectionConfigResponse *string `json:"SetReferProtectionConfigResponse" name:"SetReferProtectionConfigResponse"`
}

func (r *SetReferProtectionConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetReferProtectionConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetIpProtectionConfigRequest struct {
	*ksyunhttp.BaseRequest
	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`
	Enable   *string `json:"Enable,omitempty" name:"Enable"`
	IpType   *string `json:"IpType,omitempty" name:"IpType"`
	IpList   *string `json:"IpList,omitempty" name:"IpList"`
}

func (r *SetIpProtectionConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type SetIpProtectionConfigResponse struct {
	*ksyunhttp.BaseResponse
	SetIpProtectionConfigResponse *string `json:"SetIpProtectionConfigResponse" name:"SetIpProtectionConfigResponse"`
}

func (r *SetIpProtectionConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetIpProtectionConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetHttpHeadersConfigRequest struct {
	*ksyunhttp.BaseRequest
	DomainId    *string `json:"DomainId,omitempty" name:"DomainId"`
	HeaderKey   *string `json:"HeaderKey,omitempty" name:"HeaderKey"`
	HeaderValue *string `json:"HeaderValue,omitempty" name:"HeaderValue"`
}

func (r *SetHttpHeadersConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type SetHttpHeadersConfigResponse struct {
	*ksyunhttp.BaseResponse
	SetHttpHeadersConfig *string `json:"SetHttpHeadersConfig" name:"SetHttpHeadersConfig"`
}

func (r *SetHttpHeadersConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetHttpHeadersConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteHttpHeadersConfigRequest struct {
	*ksyunhttp.BaseRequest
	DomainId  *string `json:"DomainId,omitempty" name:"DomainId"`
	HeaderKey *string `json:"HeaderKey,omitempty" name:"HeaderKey"`
}

func (r *DeleteHttpHeadersConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteHttpHeadersConfigResponse struct {
	*ksyunhttp.BaseResponse
	DeleteHttpHeadersConfig *string `json:"DeleteHttpHeadersConfig" name:"DeleteHttpHeadersConfig"`
}

func (r *DeleteHttpHeadersConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteHttpHeadersConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetHttpHeaderListRequest struct {
	*ksyunhttp.BaseRequest
}

func (r *GetHttpHeaderListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetHttpHeaderListResponse struct {
	*ksyunhttp.BaseResponse
	GetHttpHeaderListResponse *string `json:"GetHttpHeaderListResponse" name:"GetHttpHeaderListResponse"`
}

func (r *GetHttpHeaderListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetHttpHeaderListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetRequestAuthConfigRequest struct {
	*ksyunhttp.BaseRequest
	DomainId       *string `json:"DomainId,omitempty" name:"DomainId"`
	Enable         *string `json:"Enable,omitempty" name:"Enable"`
	AuthType       *string `json:"AuthType,omitempty" name:"AuthType"`
	Key1           *string `json:"Key1,omitempty" name:"Key1"`
	Key2           *string `json:"Key2,omitempty" name:"Key2"`
	ExpirationTime *string `json:"ExpirationTime,omitempty" name:"ExpirationTime"`
}

func (r *SetRequestAuthConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type SetRequestAuthConfigResponse struct {
	*ksyunhttp.BaseResponse
	SetRequestAuthConfigResponse *string `json:"SetRequestAuthConfigResponse" name:"SetRequestAuthConfigResponse"`
}

func (r *SetRequestAuthConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetRequestAuthConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetForceRedirectConfigRequest struct {
	*ksyunhttp.BaseRequest
	DomainId     *string `json:"DomainId,omitempty" name:"DomainId"`
	RedirectType *string `json:"RedirectType,omitempty" name:"RedirectType"`
	RedirectCode *string `json:"RedirectCode,omitempty" name:"RedirectCode"`
}

func (r *SetForceRedirectConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type SetForceRedirectConfigResponse struct {
	*ksyunhttp.BaseResponse
	SetForceRedirectConfig *string `json:"SetForceRedirectConfig" name:"SetForceRedirectConfig"`
}

func (r *SetForceRedirectConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetForceRedirectConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetErrorPageConfigRequest struct {
	*ksyunhttp.BaseRequest
	DomainId   *string                         `json:"DomainId,omitempty" name:"DomainId"`
	ErrorPages []*SetErrorPageConfigErrorPages `json:"ErrorPages,omitempty" name:"ErrorPages"`
}

func (r *SetErrorPageConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type SetErrorPageConfigResponse struct {
	*ksyunhttp.BaseResponse
	SetErrorPageConfig *string `json:"SetErrorPageConfig" name:"SetErrorPageConfig"`
}

func (r *SetErrorPageConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetErrorPageConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetTLSVersionConfigRequest struct {
	*ksyunhttp.BaseRequest
	DomainId   *string   `json:"DomainId,omitempty" name:"DomainId"`
	TLSVersion []*string `json:"TLSVersion,omitempty" name:"TLSVersion"`
}

func (r *SetTLSVersionConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type SetTLSVersionConfigResponse struct {
	*ksyunhttp.BaseResponse
	TLSVersionConfigResponse *string `json:"TLSVersionConfigResponse" name:"TLSVersionConfigResponse"`
}

func (r *SetTLSVersionConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetTLSVersionConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetBillingModeRequest struct {
	*ksyunhttp.BaseRequest
	StartTime   *string `json:"StartTime,omitempty" name:"StartTime"`
	EndTime     *string `json:"EndTime,omitempty" name:"EndTime"`
	CdnType     *string `json:"CdnType,omitempty" name:"CdnType"`
	DomainIds   *string `json:"DomainIds,omitempty" name:"DomainIds"`
	Regions     *string `json:"Regions,omitempty" name:"Regions"`
	BillingMode *string `json:"BillingMode,omitempty" name:"BillingMode"`
}

func (r *GetBillingModeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetBillingModeResponse struct {
	*ksyunhttp.BaseResponse
	GetBillingModeResponse *string `json:"GetBillingModeResponse" name:"GetBillingModeResponse"`
}

func (r *GetBillingModeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetBillingModeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetBlockUrlQuotaRequest struct {
	*ksyunhttp.BaseRequest
}

func (r *GetBlockUrlQuotaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetBlockUrlQuotaResponse struct {
	*ksyunhttp.BaseResponse
	BlockUrlQuota   *int64 `json:"BlockUrlQuota" name:"BlockUrlQuota"`
	BlockUrlSurplus *int64 `json:"BlockUrlSurplus" name:"BlockUrlSurplus"`
}

func (r *GetBlockUrlQuotaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetBlockUrlQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetBandwidthDataRequest struct {
	*ksyunhttp.BaseRequest
	StartTime    *string `json:"StartTime,omitempty" name:"StartTime"`
	EndTime      *string `json:"EndTime,omitempty" name:"EndTime"`
	CdnType      *string `json:"CdnType,omitempty" name:"CdnType"`
	DomainIds    *string `json:"DomainIds,omitempty" name:"DomainIds"`
	Regions      *string `json:"Regions,omitempty" name:"Regions"`
	ResultType   *int64  `json:"ResultType,omitempty" name:"ResultType"`
	Granularity  *int64  `json:"Granularity,omitempty" name:"Granularity"`
	DataType     *string `json:"DataType,omitempty" name:"DataType"`
	ProtocolType *string `json:"ProtocolType,omitempty" name:"ProtocolType"`
}

func (r *GetBandwidthDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetBandwidthDataResponse struct {
	*ksyunhttp.BaseResponse
	StartTime   *string `json:"StartTime" name:"StartTime"`
	EndTime     *string `json:"EndTime" name:"EndTime"`
	CdnType     *string `json:"CdnType" name:"CdnType"`
	DomainIds   *string `json:"DomainIds" name:"DomainIds"`
	Regions     *string `json:"Regions" name:"Regions"`
	ResultType  *int64  `json:"ResultType" name:"ResultType"`
	Granularity *int64  `json:"Granularity" name:"Granularity"`
	DataType    *string `json:"DataType" name:"DataType"`
	Datas       []struct {
		Time    *string `json:"Time" name:"Time"`
		Bw      *int64  `json:"Bw" name:"Bw"`
		SrcBw   *int64  `json:"SrcBw" name:"SrcBw"`
		Domains []struct {
			DomainId *string `json:"DomainId" name:"DomainId"`
			Bw       *int64  `json:"Bw" name:"Bw"`
			SrcBw    *int64  `json:"SrcBw" name:"SrcBw"`
			Regions  []struct {
				Region *string `json:"Region" name:"Region"`
				Bw     *int64  `json:"Bw" name:"Bw"`
				SrcBw  *int64  `json:"SrcBw" name:"SrcBw"`
			} `json:"Regions"`
		} `json:"Domains" name:"Domains"`
	} `json:"Datas"`
}

func (r *GetBandwidthDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetBandwidthDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetFlowDataRequest struct {
	*ksyunhttp.BaseRequest
	StartTime    *string `json:"StartTime,omitempty" name:"StartTime"`
	EndTime      *string `json:"EndTime,omitempty" name:"EndTime"`
	CdnType      *string `json:"CdnType,omitempty" name:"CdnType"`
	DomainIds    *string `json:"DomainIds,omitempty" name:"DomainIds"`
	Regions      *string `json:"Regions,omitempty" name:"Regions"`
	ResultType   *int64  `json:"ResultType,omitempty" name:"ResultType"`
	Granularity  *int64  `json:"Granularity,omitempty" name:"Granularity"`
	DataType     *string `json:"DataType,omitempty" name:"DataType"`
	ProtocolType *string `json:"ProtocolType,omitempty" name:"ProtocolType"`
}

func (r *GetFlowDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetFlowDataResponse struct {
	*ksyunhttp.BaseResponse
	StartTime   *string `json:"StartTime" name:"StartTime"`
	EndTime     *string `json:"EndTime" name:"EndTime"`
	CdnType     *string `json:"CdnType" name:"CdnType"`
	DomainIds   *string `json:"DomainIds" name:"DomainIds"`
	Regions     *string `json:"Regions" name:"Regions"`
	ResultType  *int    `json:"ResultType" name:"ResultType"`
	Granularity *int    `json:"Granularity" name:"Granularity"`
	DataType    *string `json:"DataType" name:"DataType"`
	Datas       []struct {
		Time    *string `json:"Time" name:"Time"`
		Flow    *int64  `json:"Flow" name:"Flow"`
		SrcFlow *int64  `json:"SrcFlow" name:"SrcFlow"`
		Domains []struct {
			DomainId *string `json:"DomainId" name:"DomainId"`
			Flow     *int64  `json:"Flow" name:"Flow"`
			SrcFlow  *int64  `json:"SrcFlow" name:"SrcFlow"`
			Regions  []struct {
				Region  *string `json:"Region" name:"Region"`
				Flow    *int64  `json:"Flow" name:"Flow"`
				SrcFlow *int64  `json:"SrcFlow" name:"SrcFlow"`
			} `json:"Regions"`
		} `json:"Domains" name:"Domains"`
	} `json:"Datas"`
}

func (r *GetFlowDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetFlowDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetPvDataRequest struct {
	*ksyunhttp.BaseRequest
	StartTime    *string `json:"StartTime,omitempty" name:"StartTime"`
	EndTime      *string `json:"EndTime,omitempty" name:"EndTime"`
	CdnType      *string `json:"CdnType,omitempty" name:"CdnType"`
	DomainIds    *string `json:"DomainIds,omitempty" name:"DomainIds"`
	Regions      *string `json:"Regions,omitempty" name:"Regions"`
	ResultType   *int64  `json:"ResultType,omitempty" name:"ResultType"`
	Granularity  *int64  `json:"Granularity,omitempty" name:"Granularity"`
	DataType     *string `json:"DataType,omitempty" name:"DataType"`
	ProtocolType *string `json:"ProtocolType,omitempty" name:"ProtocolType"`
}

func (r *GetPvDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetPvDataResponse struct {
	*ksyunhttp.BaseResponse
	StartTime   *string `json:"StartTime" name:"StartTime"`
	EndTime     *string `json:"EndTime" name:"EndTime"`
	CdnType     *string `json:"CdnType" name:"CdnType"`
	DomainIds   *string `json:"DomainIds" name:"DomainIds"`
	Regions     *string `json:"Regions" name:"Regions"`
	ResultType  *int64  `json:"ResultType" name:"ResultType"`
	Granularity *int64  `json:"Granularity" name:"Granularity"`
	DataType    *string `json:"DataType" name:"DataType"`
	Datas       []struct {
		Time    *string `json:"Time" name:"Time"`
		Pv      *int64  `json:"Pv" name:"Pv"`
		SrcPv   *int64  `json:"SrcPv" name:"SrcPv"`
		Domains []struct {
			DomainId *string `json:"DomainId" name:"DomainId"`
			Pv       *int64  `json:"Pv" name:"Pv"`
			Regions  []struct {
				Region *string `json:"Region" name:"Region"`
				Pv     *int64  `json:"Pv" name:"Pv"`
				SrcPv  *int64  `json:"SrcPv" name:"SrcPv"`
			} `json:"Regions"`
			SrcPv *int64 `json:"SrcPv" name:"SrcPv"`
		} `json:"Domains" name:"Domains"`
	} `json:"Datas"`
}

func (r *GetPvDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetPvDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetDomainLogServiceRequest struct {
	*ksyunhttp.BaseRequest
	ActionType  *string `json:"ActionType,omitempty" name:"ActionType"`
	DomainIds   *string `json:"DomainIds,omitempty" name:"DomainIds"`
	Granularity *string `json:"Granularity,omitempty" name:"Granularity"`
}

func (r *SetDomainLogServiceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type SetDomainLogServiceResponse struct {
	*ksyunhttp.BaseResponse
	SetDomainLogService *string `json:"SetDomainLogService" name:"SetDomainLogService"`
}

func (r *SetDomainLogServiceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetDomainLogServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetCertificateRequest struct {
	*ksyunhttp.BaseRequest
	CertificateId     *string `json:"CertificateId,omitempty" name:"CertificateId"`
	CertificateName   *string `json:"CertificateName,omitempty" name:"CertificateName"`
	ServerCertificate *string `json:"ServerCertificate,omitempty" name:"ServerCertificate"`
	PrivateKey        *string `json:"PrivateKey,omitempty" name:"PrivateKey"`
}

func (r *SetCertificateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type SetCertificateResponse struct {
	*ksyunhttp.BaseResponse
	CertificateId *string `json:"CertificateId" name:"CertificateId"`
}

func (r *SetCertificateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RemoveCertificatesRequest struct {
	*ksyunhttp.BaseRequest
	CertificateIds *string `json:"CertificateIds,omitempty" name:"CertificateIds"`
}

func (r *RemoveCertificatesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type RemoveCertificatesResponse struct {
	*ksyunhttp.BaseResponse
	RemoveCertificates *string `json:"RemoveCertificates" name:"RemoveCertificates"`
}

func (r *RemoveCertificatesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RemoveCertificatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ValidateIPRequest struct {
	*ksyunhttp.BaseRequest
	Ip *string `json:"Ip,omitempty" name:"Ip"`
}

func (r *ValidateIPRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ValidateIPResponse struct {
	*ksyunhttp.BaseResponse
	CdnIp    *string `json:"CdnIp" name:"CdnIp"`
	Isp      *string `json:"Isp" name:"Isp"`
	Region   *string `json:"Region" name:"Region"`
	Province *string `json:"Province" name:"Province"`
	City     *string `json:"City" name:"City"`
}

func (r *ValidateIPResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ValidateIPResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetCdnBlockDomainUrlRequest struct {
	*ksyunhttp.BaseRequest
	BlockType        *string                     `json:"BlockType,omitempty" name:"BlockType"`
	Urls             []*SetCdnBlockDomainUrlUrls `json:"Urls,omitempty" name:"Urls"`
	RefreshOnUnblock *string                     `json:"RefreshOnUnblock,omitempty" name:"RefreshOnUnblock"`
}

func (r *SetCdnBlockDomainUrlRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type SetCdnBlockDomainUrlResponse struct {
	*ksyunhttp.BaseResponse
	TaskId *string `json:"TaskId" name:"TaskId"`
}

func (r *SetCdnBlockDomainUrlResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetCdnBlockDomainUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SubmitRefreshCachesRequest struct {
	*ksyunhttp.BaseRequest
	Files []*SubmitRefreshCachesFiles `json:"Files,omitempty" name:"Files"`
	Dirs  []*SubmitRefreshCachesDirs  `json:"Dirs,omitempty" name:"Dirs"`
}

func (r *SubmitRefreshCachesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type SubmitRefreshCachesResponse struct {
	*ksyunhttp.BaseResponse
	RefreshTaskId *string `json:"RefreshTaskId" name:"RefreshTaskId"`
}

func (r *SubmitRefreshCachesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SubmitRefreshCachesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SubmitPreloadCachesRequest struct {
	*ksyunhttp.BaseRequest
	Urls []*SubmitPreloadCachesUrls `json:"Urls,omitempty" name:"Urls"`
}

func (r *SubmitPreloadCachesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type SubmitPreloadCachesResponse struct {
	*ksyunhttp.BaseResponse
	PreloadTaskId *string `json:"PreloadTaskId" name:"PreloadTaskId"`
}

func (r *SubmitPreloadCachesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SubmitPreloadCachesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetCertificateConfigRequest struct {
	*ksyunhttp.BaseRequest
	Enable            *string `json:"Enable,omitempty" name:"Enable"`
	DomainIds         *string `json:"DomainIds,omitempty" name:"DomainIds"`
	CertificateId     *string `json:"CertificateId,omitempty" name:"CertificateId"`
	CertificateName   *string `json:"CertificateName,omitempty" name:"CertificateName"`
	ServerCertificate *string `json:"ServerCertificate,omitempty" name:"ServerCertificate"`
	PrivateKey        *string `json:"PrivateKey,omitempty" name:"PrivateKey"`
}

func (r *SetCertificateConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type SetCertificateConfigResponse struct {
	*ksyunhttp.BaseResponse
	ConfigCertificateResponse *string `json:"ConfigCertificateResponse" name:"ConfigCertificateResponse"`
}

func (r *SetCertificateConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetCertificateConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRefreshOrPreloadTaskRequest struct {
	*ksyunhttp.BaseRequest
	StartTime  *string                        `json:"StartTime,omitempty" name:"StartTime"`
	EndTime    *string                        `json:"EndTime,omitempty" name:"EndTime"`
	TaskId     *string                        `json:"TaskId,omitempty" name:"TaskId"`
	DomainName *string                        `json:"DomainName,omitempty" name:"DomainName"`
	Urls       []*GetRefreshOrPreloadTaskUrls `json:"Urls,omitempty" name:"Urls"`
	Type       *string                        `json:"Type,omitempty" name:"Type"`
	SubType    *string                        `json:"SubType,omitempty" name:"SubType"`
	PageSize   *int64                         `json:"PageSize,omitempty" name:"PageSize"`
	PageNumber *int64                         `json:"PageNumber,omitempty" name:"PageNumber"`
}

func (r *GetRefreshOrPreloadTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetRefreshOrPreloadTaskResponse struct {
	*ksyunhttp.BaseResponse
	StartTime *string `json:"StartTime" name:"StartTime"`
	EndTime   *string `json:"EndTime" name:"EndTime"`
	Urls      []struct {
		Url *string `json:"Url" name:"Url"`
	} `json:"Urls"`
	PageSize   *int64 `json:"PageSize" name:"PageSize"`
	PageNumber *int64 `json:"PageNumber" name:"PageNumber"`
	TotalCount *int64 `json:"TotalCount" name:"TotalCount"`
	Datas      []struct {
		Type     *string  `json:"Type" name:"Type"`
		SubType  *string  `json:"SubType" name:"SubType"`
		Url      *string  `json:"Url" name:"Url"`
		Progress *float64 `json:"Progress" name:"Progress"`
		Status   *string  `json:"Status" name:"Status"`
		TaskId   *string  `json:"TaskId" name:"TaskId"`
	} `json:"Datas"`
	CreateTime *string `json:"CreateTime" name:"CreateTime"`
}

func (r *GetRefreshOrPreloadTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRefreshOrPreloadTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
