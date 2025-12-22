package v3

import (
	"encoding/json"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)

type SetCacheRuleConfigCacheRules struct {
	CacheRuleType *string `json:"CacheRuleType,omitempty" name:"CacheRuleType"`
	Value         *string `json:"Value,omitempty" name:"Value"`
	CacheEnable   *string `json:"CacheEnable,omitempty" name:"CacheEnable"`
	CacheTime     *int    `json:"CacheTime,omitempty" name:"CacheTime"`
	RespectOrigin *string `json:"RespectOrigin,omitempty" name:"RespectOrigin"`
}

type GetDomainLogsRequest struct {
	*ksyunhttp.BaseRequest
	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`
}

func (r *GetDomainLogsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetDomainLogsResponse struct {
	*ksyunhttp.BaseResponse
	GetDomainLogsResponse *string `json:"GetDomainLogsResponse" name:"GetDomainLogsResponse"`
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
	PageSize     *int    `json:"PageSize,omitempty" name:"PageSize"`
	PageNumber   *int    `json:"PageNumber,omitempty" name:"PageNumber"`
	DomainName   *string `json:"DomainName,omitempty" name:"DomainName"`
	ProjectId    *int    `json:"ProjectId,omitempty" name:"ProjectId"`
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
	PageNumber *int `json:"PageNumber" name:"PageNumber"`
	PageSize   *int `json:"PageSize" name:"PageSize"`
	TotalCount *int `json:"TotalCount" name:"TotalCount"`
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
